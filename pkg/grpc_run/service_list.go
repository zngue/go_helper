package grpc_run

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/zngue/go_helper/pkg"
	"time"
)

var ServiceHostList map[string]string
var OutTime time.Duration

/*
*@Author Administrator
*@Date 23/4/2021 13:31
*@desc 获取当前本地注册服务列表
 */
func ServiceLocalList() error {
	mapString := viper.GetStringMapString("grpc.ServiceList")
	if mapString == nil {
		return errors.New("local service list is empty")
	}
	ServiceHostList = mapString
	return nil
}

/*
*@Author Administrator
*@Date 23/4/2021 13:40
*@desc
 */
func ServiceRedisList() error {
	result, err := pkg.RedisConn.HGetAll("resgter").Result()
	if err != nil {
		return err
	}
	if result != nil {
		return errors.New("redis resgter key is empty")
	}
	ServiceHostList = result
	return nil
}
