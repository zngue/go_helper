package http

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/zngue/go_helper/pkg"
	"io/ioutil"
	"strings"
	"time"
)

const POST = "POST"
const GET = "GET"

type HttpFn func(http *ZngueHttp) *ZngueHttp
type serverhost struct {
	ServiceName string
	Host        string
}
type HttpMico struct {
	Method    string
	Url       string
	Host      string
	EndPoint  string
	ServiceId string
	Body      []byte
	Param     map[string]interface{}
	Header    map[string]interface{}
	fn        HttpFn
	Timeout   time.Duration
}

func (m *HttpMico) DoRequest() (*ZngueHttp, error) {
	var httpRequest *ZngueHttp
	isMicro := viper.GetBool("micro.isMicro")
	if isMicro {
		url, err := m.BuildUrl()
		if err != nil {
			return nil, err
		}
		m.Url = url
	}
	if m.Url == "" {
		return nil, errors.New("url is empty")
	}
	switch m.Method {
	case POST:
		httpRequest = Post(m.Url)
	case GET:
		httpRequest = Get(m.Url)
	}
	if m.Body != nil {
		httpRequest.body = m.Body
	}
	if m.Param != nil {
		for key, val := range m.Param {
			httpRequest.Param(key, cast.ToString(val))
		}
	}
	if m.Header != nil {
		for key, val := range m.Header {
			httpRequest.Header(key, cast.ToString(val))
		}
	}
	if m.fn != nil {
		httpRequest = m.fn(httpRequest)
	}
	return httpRequest, nil
}

func (m *HttpMico) Response() (string, error) {
	req, err := m.DoRequest()
	if err != nil {
		return "", err
	}
	response, errResp := req.Response()
	if errResp != nil {
		return "", errResp
	}
	all, errs := ioutil.ReadAll(response.Body)
	if errs != nil {
		return "", errs
	}
	return string(all), nil
}
func (m *HttpMico) Formaturl(url string) string {
	if strings.Index(url, "http") < 0 {
		url = "http://" + url
	}
	last := url[(len(url) - 1):]
	if last != "/" {
		url += "/"
	}
	return url
}

//读取并绑定数据
func (m *HttpMico) Bind(i interface{}) error {
	str, err := m.Response()
	if err != nil {
		return err
	}
	err2 := json.Unmarshal([]byte(str), i)
	if err2 != nil {
		return err2
	}
	return nil
}

//redis数据读取
func (m *HttpMico) GetResiger() (string, error) {
	redisDBNum := viper.GetString("micro.redisDB")
	DBNUM := viper.GetString("REDIS.DBNUM")
	if pkg.RedisConn == nil {
		return "", errors.New("RedisConn is nil")
	}
	var rds *redis.StringSliceCmd
	if redisDBNum == DBNUM {
		rds = pkg.RedisConn.SMembers("register")
	} else {
		pipeline := pkg.RedisConn.Pipeline()
		pipeline.Do("select", redisDBNum)
		rds = pipeline.SMembers("register")
		pipeline.Do("select", DBNUM)
		_, err := pipeline.Exec()
		if err != nil {
			return "", err
		}
	}
	register, err := rds.Result()
	if err != nil {
		return "", err
	}
	for _, val := range register {
		server := serverhost{}
		err := json.Unmarshal([]byte(val), &server)
		if err != nil {
			return "", err
		}
		if server.ServiceName == m.ServiceId {
			return m.Formaturl(server.Host) + strings.Trim(m.EndPoint, " "), nil
		}
	}
	return "", errors.New("ServiceId is not exist")
}

//本地数据请求
func (m *HttpMico) GetLocal() (string, error) {
	serviceList := viper.GetStringMapString("micro.serviceList")
	if len(serviceList) == 0 {
		return "", errors.New("configure micro.serviceList not exist")
	}
	for id, host := range serviceList {
		if m.ServiceId == id {
			return m.Formaturl(host) + strings.Trim(m.EndPoint, " "), nil
		}
	}
	return "", errors.New("ServiceId is not exist")
}

//创建url
func (m *HttpMico) BuildUrl() (string, error) {
	if m.ServiceId == "" {
		return "", errors.New("ServiceId is empty")
	}
	isLocal := viper.GetBool("micro.isLocal")
	if isLocal {
		return m.GetLocal()
	} else {
		return m.GetResiger()
	}
}
