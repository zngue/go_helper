package http

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var defaultSetting = ZngueHttpSettings{
	UserAgent:        "ZngueServer",
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	Gzip:             true,
	DumpBody:         true,
}

type FilterChain func(next Filter) Filter

type Filter func(ctx context.Context, req *ZngueHttp) (*http.Response, error)

var defaultCookieJar http.CookieJar
var settingMutex sync.Mutex

// it will be the last filter and execute request.Do
var doRequestFilter = func(ctx context.Context, req *ZngueHttp) (*http.Response, error) {
	return req.doRequest(ctx)
}

// createDefaultCookie creates a global cookiejar to store cookies.
func createDefaultCookie() {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultCookieJar, _ = cookiejar.New(nil)
}

// SetDefaultSetting overwrites default settings
func SetDefaultSetting(setting ZngueHttpSettings) {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultSetting = setting
}

// NewZngueRequest returns *ZngueHttp with specific method
func NewZngueRequest(rawurl, method string) *ZngueHttp {
	var resp http.Response
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Println("Httplib:", err)
	}
	req := http.Request{
		URL:        u,
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	return &ZngueHttp{
		url:     rawurl,
		req:     &req,
		params:  map[string][]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
}

// Get returns *ZngueHttp with GET method.
func Get(url string) *ZngueHttp {
	return NewZngueRequest(url, "GET")
}

// Post returns *ZngueHttp with POST method.
func Post(url string) *ZngueHttp {
	return NewZngueRequest(url, "POST")
}

// Put returns *ZngueHttp with PUT method.
func Put(url string) *ZngueHttp {
	return NewZngueRequest(url, "PUT")
}

// Delete returns *ZngueHttp DELETE method.
func Delete(url string) *ZngueHttp {
	return NewZngueRequest(url, "DELETE")
}

// Head returns *ZngueHttp with HEAD method.
func Head(url string) *ZngueHttp {
	return NewZngueRequest(url, "HEAD")
}

// ZngueHttpSettings is the http.Client setting
type ZngueHttpSettings struct {
	ShowDebug        bool
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TLSClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	Transport        http.RoundTripper
	CheckRedirect    func(req *http.Request, via []*http.Request) error
	EnableCookie     bool
	Gzip             bool
	DumpBody         bool
	Retries          int // if set to -1 means will retry forever
	RetryDelay       time.Duration
	FilterChains     []FilterChain
}

// ZngueHttp provides more useful methods than http.Request for requesting a url.
type ZngueHttp struct {
	url     string
	req     *http.Request
	params  map[string][]string
	files   map[string]string
	setting ZngueHttpSettings
	resp    *http.Response
	body    []byte
	dump    []byte
}

// GetRequest returns the request object
func (b *ZngueHttp) GetRequest() *http.Request {
	return b.req
}

// Setting changes request settings
func (b *ZngueHttp) Setting(setting ZngueHttpSettings) *ZngueHttp {
	b.setting = setting
	return b
}

// SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.
func (b *ZngueHttp) SetBasicAuth(username, password string) *ZngueHttp {
	b.req.SetBasicAuth(username, password)
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *ZngueHttp) SetEnableCookie(enable bool) *ZngueHttp {
	b.setting.EnableCookie = enable
	return b
}

// SetUserAgent sets User-Agent header field
func (b *ZngueHttp) SetUserAgent(useragent string) *ZngueHttp {
	b.setting.UserAgent = useragent
	return b
}

// Debug sets show debug or not when executing request.
func (b *ZngueHttp) Debug(isdebug bool) *ZngueHttp {
	b.setting.ShowDebug = isdebug
	return b
}

// Retries sets Retries times.
// default is 0 (never retry)
// -1 retry indefinitely (forever)
// Other numbers specify the exact retry amount
func (b *ZngueHttp) Retries(times int) *ZngueHttp {
	b.setting.Retries = times
	return b
}

// RetryDelay sets the time to sleep between reconnection attempts
func (b *ZngueHttp) RetryDelay(delay time.Duration) *ZngueHttp {
	b.setting.RetryDelay = delay
	return b
}

// DumpBody sets the DumbBody field
func (b *ZngueHttp) DumpBody(isdump bool) *ZngueHttp {
	b.setting.DumpBody = isdump
	return b
}

// DumpRequest returns the DumpRequest
func (b *ZngueHttp) DumpRequest() []byte {
	return b.dump
}

// SetTimeout sets connect time out and read-write time out for ZngueRequest.
func (b *ZngueHttp) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *ZngueHttp {
	b.setting.ConnectTimeout = connectTimeout
	b.setting.ReadWriteTimeout = readWriteTimeout
	return b
}

// SetTLSClientConfig sets TLS connection configuration if visiting HTTPS url.
func (b *ZngueHttp) SetTLSClientConfig(config *tls.Config) *ZngueHttp {
	b.setting.TLSClientConfig = config
	return b
}

// Header adds header item string in request.
func (b *ZngueHttp) Header(key, value string) *ZngueHttp {
	b.req.Header.Set(key, value)
	return b
}

// SetHost set the request host
func (b *ZngueHttp) SetHost(host string) *ZngueHttp {
	b.req.Host = host
	return b
}

// SetProtocolVersion sets the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *ZngueHttp) SetProtocolVersion(vers string) *ZngueHttp {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}

	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		b.req.Proto = vers
		b.req.ProtoMajor = major
		b.req.ProtoMinor = minor
	}

	return b
}

// SetCookie adds a cookie to the request.
func (b *ZngueHttp) SetCookie(cookie *http.Cookie) *ZngueHttp {
	b.req.Header.Add("Cookie", cookie.String())
	return b
}

// SetTransport sets the transport field
func (b *ZngueHttp) SetTransport(transport http.RoundTripper) *ZngueHttp {
	b.setting.Transport = transport
	return b
}

// SetProxy sets the HTTP proxy
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (b *ZngueHttp) SetProxy(proxy func(*http.Request) (*url.URL, error)) *ZngueHttp {
	b.setting.Proxy = proxy
	return b
}

// SetCheckRedirect specifies the policy for handling redirects.
//
// If CheckRedirect is nil, the Client uses its default policy,
// which is to stop after 10 consecutive requests.
func (b *ZngueHttp) SetCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) *ZngueHttp {
	b.setting.CheckRedirect = redirect
	return b
}

// SetFilters will use the filter as the invocation filters
func (b *ZngueHttp) SetFilters(fcs ...FilterChain) *ZngueHttp {
	b.setting.FilterChains = fcs
	return b
}

// AddFilters adds filter
func (b *ZngueHttp) AddFilters(fcs ...FilterChain) *ZngueHttp {
	b.setting.FilterChains = append(b.setting.FilterChains, fcs...)
	return b
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *ZngueHttp) Param(key, value string) *ZngueHttp {
	if param, ok := b.params[key]; ok {
		b.params[key] = append(param, value)
	} else {
		b.params[key] = []string{value}
	}
	return b
}

// PostFile adds a post file to the request
func (b *ZngueHttp) PostFile(formname, filename string) *ZngueHttp {
	b.files[formname] = filename
	return b
}

// Body adds request raw body.
// Supports string and []byte.
func (b *ZngueHttp) Body(data interface{}) *ZngueHttp {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	}
	return b
}

// XMLBody adds the request raw body encoded in XML.
func (b *ZngueHttp) XMLBody(obj interface{}) (*ZngueHttp, error) {
	if b.req.Body == nil && obj != nil {
		byts, err := xml.Marshal(obj)
		if err != nil {
			return b, err
		}
		b.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		b.req.ContentLength = int64(len(byts))
		b.req.Header.Set("Content-Type", "application/xml")
	}
	return b, nil
}

// YAMLBody adds the request raw body encoded in YAML.
func (b *ZngueHttp) YAMLBody(obj interface{}) (*ZngueHttp, error) {
	if b.req.Body == nil && obj != nil {
		byts, err := yaml.Marshal(obj)
		if err != nil {
			return b, err
		}
		b.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		b.req.ContentLength = int64(len(byts))
		b.req.Header.Set("Content-Type", "application/x+yaml")
	}
	return b, nil
}

// JSONBody adds the request raw body encoded in JSON.
func (b *ZngueHttp) JSONBody(obj interface{}) (*ZngueHttp, error) {
	if b.req.Body == nil && obj != nil {
		byts, err := json.Marshal(obj)
		if err != nil {
			return b, err
		}
		b.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		b.req.ContentLength = int64(len(byts))
		b.req.Header.Set("Content-Type", "application/json")
	}
	return b, nil
}

func (b *ZngueHttp) buildURL(paramBody string) {
	// build GET url with query string
	if b.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Contains(b.url, "?") {
			b.url += "&" + paramBody
		} else {
			b.url = b.url + "?" + paramBody
		}
		return
	}

	// build POST/PUT/PATCH url and body
	if (b.req.Method == "POST" || b.req.Method == "PUT" || b.req.Method == "PATCH" || b.req.Method == "DELETE") && b.req.Body == nil {
		// with files
		if len(b.files) > 0 {
			pr, pw := io.Pipe()
			bodyWriter := multipart.NewWriter(pw)
			go func() {
				for formname, filename := range b.files {
					fileWriter, err := bodyWriter.CreateFormFile(formname, filename)
					if err != nil {
						log.Println("Httplib:", err)
					}
					fh, err := os.Open(filename)
					if err != nil {
						log.Println("Httplib:", err)
					}
					// iocopy
					_, err = io.Copy(fileWriter, fh)
					fh.Close()
					if err != nil {
						log.Println("Httplib:", err)
					}
				}
				for k, v := range b.params {
					for _, vv := range v {
						bodyWriter.WriteField(k, vv)
					}
				}
				bodyWriter.Close()
				pw.Close()
			}()
			b.Header("Content-Type", bodyWriter.FormDataContentType())
			b.req.Body = ioutil.NopCloser(pr)
			b.Header("Transfer-Encoding", "chunked")
			return
		}

		// with params
		if len(paramBody) > 0 {
			b.Header("Content-Type", "application/x-www-form-urlencoded")
			b.Body(paramBody)
		}
	}
}

func (b *ZngueHttp) getResponse() (*http.Response, error) {
	if b.resp.StatusCode != 0 {
		return b.resp, nil
	}
	resp, err := b.DoRequest()
	if err != nil {
		return nil, err
	}
	b.resp = resp
	return resp, nil
}

// DoRequest executes client.Do
func (b *ZngueHttp) DoRequest() (resp *http.Response, err error) {
	return b.DoRequestWithCtx(context.Background())
}

func (b *ZngueHttp) DoRequestWithCtx(ctx context.Context) (resp *http.Response, err error) {

	root := doRequestFilter
	if len(b.setting.FilterChains) > 0 {
		for i := len(b.setting.FilterChains) - 1; i >= 0; i-- {
			root = b.setting.FilterChains[i](root)
		}
	}
	return root(ctx, b)
}

func (b *ZngueHttp) doRequest(ctx context.Context) (resp *http.Response, err error) {
	var paramBody string
	if len(b.params) > 0 {
		var buf bytes.Buffer
		for k, v := range b.params {
			for _, vv := range v {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(vv))
				buf.WriteByte('&')
			}
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	b.buildURL(paramBody)
	urlParsed, err := url.Parse(b.url)
	if err != nil {
		return nil, err
	}

	b.req.URL = urlParsed

	trans := b.setting.Transport

	if trans == nil {
		// create default transport
		trans = &http.Transport{
			TLSClientConfig:     b.setting.TLSClientConfig,
			Proxy:               b.setting.Proxy,
			Dial:                TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout),
			MaxIdleConnsPerHost: 100,
		}
	} else {
		// if b.transport is *http.Transport then set the settings.
		if t, ok := trans.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = b.setting.TLSClientConfig
			}
			if t.Proxy == nil {
				t.Proxy = b.setting.Proxy
			}
			if t.Dial == nil {
				t.Dial = TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout)
			}
		}
	}

	var jar http.CookieJar
	if b.setting.EnableCookie {
		if defaultCookieJar == nil {
			createDefaultCookie()
		}
		jar = defaultCookieJar
	}

	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	if b.setting.UserAgent != "" && b.req.Header.Get("User-Agent") == "" {
		b.req.Header.Set("User-Agent", b.setting.UserAgent)
	}

	if b.setting.CheckRedirect != nil {
		client.CheckRedirect = b.setting.CheckRedirect
	}

	if b.setting.ShowDebug {
		dump, err := httputil.DumpRequest(b.req, b.setting.DumpBody)
		if err != nil {
			log.Println(err.Error())
		}
		b.dump = dump
	}
	// retries default value is 0, it will run once.
	// retries equal to -1, it will run forever until success
	// retries is setted, it will retries fixed times.
	// Sleeps for a 400ms between calls to reduce spam
	for i := 0; b.setting.Retries == -1 || i <= b.setting.Retries; i++ {
		resp, err = client.Do(b.req)
		if err == nil {
			break
		}
		time.Sleep(b.setting.RetryDelay)
	}
	return resp, err
}

// String returns the body string in response.
// Calls Response inner.
func (b *ZngueHttp) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Bytes returns the body []byte in response.
// Calls Response inner.
func (b *ZngueHttp) Bytes() ([]byte, error) {
	if b.body != nil {
		return b.body, nil
	}
	resp, err := b.getResponse()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if b.setting.Gzip && resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		b.body, err = ioutil.ReadAll(reader)
		return b.body, err
	}
	b.body, err = ioutil.ReadAll(resp.Body)
	return b.body, err
}

// ToFile saves the body data in response to one file.
// Calls Response inner.
func (b *ZngueHttp) ToFile(filename string) error {
	resp, err := b.getResponse()
	if err != nil {
		return err
	}
	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()
	err = pathExistAndMkdir(filename)
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// Check if the file directory exists. If it doesn't then it's created
func pathExistAndMkdir(filename string) (err error) {
	filename = path.Dir(filename)
	_, err = os.Stat(filename)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(filename, os.ModePerm)
		if err == nil {
			return nil
		}
	}
	return err
}

// ToJSON returns the map that marshals from the body bytes as json in response.
// Calls Response inner.
func (b *ZngueHttp) ToJSON(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// ToXML returns the map that marshals from the body bytes as xml in response .
// Calls Response inner.
func (b *ZngueHttp) ToXML(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// ToYAML returns the map that marshals from the body bytes as yaml in response .
// Calls Response inner.
func (b *ZngueHttp) ToYAML(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// Response executes request client gets response manually.
func (b *ZngueHttp) Response() (*http.Response, error) {
	return b.getResponse()
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}
