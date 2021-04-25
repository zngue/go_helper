package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	pb "github.com/zngue/go_helper/eg/pbmodel"
	"github.com/zngue/go_helper/eg/temp"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/grpc_run"
	"github.com/zngue/go_helper/pkg/http"
	"github.com/zngue/go_helper/pkg/sign_chan"
	"github.com/zngue/go_helper/pkg/where"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"strings"
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
	err2 := pkg.NewConfig(pkg.Path("eg/conf"))
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
	err2 := pkg.NewConfig()
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
	err2 := pkg.NewConfig()
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

func TestUserInfoHttp(t *testing.T) {
	http, err := pkg.GinRun("3378", func(group *gin.RouterGroup) {

	})
	if err != nil {
		sign_chan.SignLog(err)
	}
	go func() {
		cerr := pkg.NewConfig(pkg.Path("eg/conf"))
		if cerr != nil {
			sign_chan.SignLog(cerr)
		}
		_, merr := pkg.NewMysql()
		if merr != nil {
			sign_chan.SignLog(merr)
		}
	}()
	go func() {
		err2 := http.ListenAndServe()
		if err2 != nil {
			sign_chan.SignLog(err2)
		}
	}()
	sign_chan.ListClose(func(ctx context.Context) error {
		return http.Shutdown(ctx)
	})

}

func TestMysql(t *testing.T) {
	pkg.NewConfig(pkg.Path("eg/conf"))
	pkg.NewMysql()
	var req temp.Request
	//req.Data=&temp.ArticleList{}
	model := pkg.MysqlConn.Model(&temp.ArticleList{})
	req.Actions = 1
	req.IsPaginate = true
	req.PageSize = 100
	err := req.Common(model).Error
	fmt.Println(err)
}

func TestGrpcService(t *testing.T) {
	pkg.NewConfig(pkg.Path("eg/conf"))
	err2 := grpc_run.ServiceLocalList()
	fmt.Println(err2)
	err := grpc_run.ServiceRegister("sy:user", func(server *grpc.Server) {
		pb.RegisterAmpArticleServiceServer(server, new(pb.UnimplementedAmpArticleServiceServer))
	})
	fmt.Println(err)
}

func TestGrpcClient(t *testing.T) {
	pkg.NewConfig(pkg.Path("eg/conf"))
	err2 := grpc_run.ServiceLocalList()
	errgrpx := grpc_run.ClientRegister("sy:user", func(conn *grpc.ClientConn, ctx context.Context) (err error) {
		in := pb.AmpArticle{
			Id: 1,
		}
		client := pb.NewAmpArticleServiceClient(conn)
		list, err4 := client.List(ctx, &in)
		fmt.Println(list, err4)
		return nil
	})
	fmt.Println(err2, errgrpx)
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
