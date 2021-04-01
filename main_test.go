package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/http"
	"testing"
	"time"
)

type SybbAds struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"id"`
	NewsID    int       `gorm:"column:news_id;type:int(11)" json:"news_id"`          // 类型：1社群，2服务
	Type      int8      `gorm:"column:type;type:tinyint(2)" json:"type"`             // 跳转类型，1外链，2原生
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`         // 标题
	ImageURL  string    `gorm:"column:image_url;type:varchar(255)" json:"image_url"` // 图片链接
	URL       string    `gorm:"column:url;type:varchar(255)" json:"url"`             // 外链
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	BbType    int8      `gorm:"column:bb_type;type:tinyint(2)" json:"bb_type"` // 1-社群，2-服务，345678--待定
}

type Ads struct {
	pkg.CommonRequest
}

func MsiOne() {
	err2 := pkg.NewConfig("app", "eg/conf", "yaml")
	mysql, err := pkg.NewMysql()
	redis, err3 := pkg.NewRedis()
	pkg.GinRun("250", func(group *gin.RouterGroup) {
		group.GET("", func(context *gin.Context) {
			context.JSON(200, 250)
		})
	})
	fmt.Println(err2, mysql, redis, err, err3)
}
func MsiTwo() {
	err2 := pkg.NewConfig("app", "eg/confs", "yaml")
	mysql, err := pkg.NewMysql()
	redis, err3 := pkg.NewRedis()
	pkg.GinRun("251", func(group *gin.RouterGroup) {
		group.GET("", func(context *gin.Context) {
			context.JSON(200, 251)
		})
	})
	fmt.Println(err2, mysql, redis, err, err3)
}
func Msithree() {
	err2 := pkg.NewConfig("app", "eg/confss", "yaml")
	mysql, err := pkg.NewMysql()
	redis, err3 := pkg.NewRedis()
	pkg.GinRun("252", func(group *gin.RouterGroup) {
		group.GET("", func(context *gin.Context) {
			context.JSON(200, 252)
		})
	})
	fmt.Println(err2, mysql, redis, err, err3)
}
func TestUn(t *testing.T) {
	MsiOne()
}

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
	err2 := pkg.NewConfig("app", "eg/conf", "yaml")
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
