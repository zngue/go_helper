package common_run

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/http"
	"github.com/zngue/go_helper/pkg/sign_chan"
	"gorm.io/gorm"
	"strings"
)

type RunLoad struct {
	MysqlLoad        bool
	MysqlOption      []gorm.Option
	MysqlCoonn		MysqlConnet
	RedisLoad        bool
	ConfigLoad       bool
	ConfigOption     []pkg.ConfigFn
	FnRouter         []pkg.RouterFun
	IsRegisterCenter bool
}
type MysqlConnet func(db *gorm.DB)
type RunLoadFn func(load *RunLoad) *RunLoad

func MysqlLoad(mysql bool) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.MysqlLoad = mysql
		return load
	}
}
func MysqlConn(conn  MysqlConnet) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.MysqlCoonn=conn
		return load
	}
}
func MysqlOption(optionArr ...gorm.Option) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.MysqlOption = optionArr
		return load
	}
}
func ConfigLoad(c bool) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.ConfigLoad = c
		return load
	}
}
func ConfigOption(fn ...pkg.ConfigFn) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.ConfigOption = fn
		return load
	}
}
func RedisLoad(c bool) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.RedisLoad = c
		return load
	}
}
func FnRouter(fun ...pkg.RouterFun) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.FnRouter = fun
		return load
	}
}
func IsRegisterCenter(c bool) RunLoadFn {
	return func(load *RunLoad) *RunLoad {
		load.IsRegisterCenter = c
		return load
	}
}

// MicHttp 服务请求
func MicHttp(Method, serverName string, EndPoint string, data map[string]interface{}) (string, error) {
	url := viper.GetString("micro.serviceList."+serverName) + "/"+EndPoint
	if strings.Index(url, "http") < 0 {
		url = "http://" + url
	}
	httpMicro := &http.HttpMico{
		Method: Method,
		Url:    url,
		Param:  data,
	}
	return httpMicro.Response()

}
/*
func MicHttp(Method, registerCenter string, data map[string]interface{}) (string, error) {
	url := viper.GetString("micro.serviceList."+registerCenter) + "/register"
	httpMicro := &http.HttpMico{
		Method: Method,
		Url:    url,
		Param:  data,
	}
	return httpMicro.Response()

}*/

// ServiceRegister 注册中心服务
func ServiceRegister() {
	port := viper.GetString("AppPort")
	title := viper.GetString("AppTitle")
	host := viper.GetString("AppHost")
	name := viper.GetString("AppName")
	if len(port) == 0 || len(title) == 0 || len(host) == 0 || len(name) == 0 {
		return
	}
	if port != "6006" { //中心化注册
		MicHttp(http.POST, "registerCenter","register", map[string]interface{}{
			"port":  port,
			"title": title,
			"host":  host,
			"name":  name,
		})
	}

}
func CommonGinRun(runArr ...RunLoadFn) {
	load := &RunLoad{
		MysqlLoad:        true,
		ConfigLoad:       true,
		RedisLoad:        false,
		IsRegisterCenter: false,
	}
	if len(runArr) > 0 {
		for _, fn := range runArr {
			load = fn(load)
		}
	}
	if load.ConfigLoad {
		if err := pkg.NewConfig(load.ConfigOption...); err != nil {
			logrus.Fatal(err)
			sign_chan.SignLog(err)
			return
		}
	}
	if load.MysqlLoad {
		if mysqlConn, err := pkg.NewMysql(load.MysqlOption...); err != nil {
			logrus.Fatal(err)
			return
		}else{
			if load.MysqlCoonn!=nil {
				load.MysqlCoonn(mysqlConn)
			}
		}
	}
	if load.RedisLoad {
		if _, err := pkg.NewRedis(); err != nil {
			logrus.Fatal(err)
			sign_chan.SignLog(err)
			return
		}
	}
	go func() {
		if load.IsRegisterCenter && load.MysqlLoad {
			ServiceRegister() //自动注册服务
		}
	}()
	port := viper.GetString("AppPort")
	ginRun, err := pkg.GinRun(port, load.FnRouter...)
	if err != nil {
		sign_chan.SignLog(err)
		return
	}
	var errInfo error
	go func() {
		errInfo = ginRun.ListenAndServe()
		if errInfo != nil {
			sign_chan.SignLog(errInfo)
		}
	}()
	sign_chan.SignChalNotify()
	sign_chan.ListClose(func(ctx context.Context) error {
		return ginRun.Shutdown(ctx)
	})
}
