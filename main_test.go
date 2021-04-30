package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/config"
	"github.com/zngue/go_helper/pkg/http"
	"github.com/zngue/go_helper/pkg/where"
	"gorm.io/gorm"
	"strings"
	"testing"
)

type UserInfo func(user string) string

func maisn(info UserInfo) string {
	usera := info("1")
	return usera
}
func TestNs(t *testing.T) {

	s := maisn(func(user string) string {

		return user + "1223"
	})
	fmt.Println(s)
}
func TestHttp(t *testing.T) {
	err2 := pkg.NewConfig()
	mysql, err := pkg.NewMysql()
	redis, err3 := pkg.NewRedis()
	mico := http.HttpMico{
		Method:    http.GET,
		ServiceId: "sy:api",
		EndPoint:  "api/bangbang/groupownet/home",
	}
	url, err2 := mico.Response()

	fmt.Println(url, err2)

	fmt.Println(err2, mysql, redis, err, err3, mico)
}

func TestWhere(t *testing.T) {
	where.RegsterHooks(where.ResiterHooksOption{
		Hooks: func(option *where.HooksOption) *gorm.DB {
			s := option.Value.String()
			sList := strings.Split(s, ",")
			if len(sList) == 2 && sList[0] != "" && sList[1] != "" {
				return option.DB.Where(option.Field+" >= ? ", sList[0]).Where(option.Field+" <= ?", sList[1])
			}
			return option.DB
		},
		Action: func(option *where.HooksOption) bool {
			if cast.ToString(option.Value.Interface()) != option.Default {
				return true
			}
			return false
		},
		Where: "between",
	})

}

/*
*@Author Administrator
*@Date 29/4/2021 11:24
*@desc
 */
func TestOss(t *testing.T) {

	config.NewConfig(config.Path("eg/conf"))

	run, _ := pkg.GinRun("7898", func(engine *gin.Engine) {

	})
	run.ListenAndServe()

}
