package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type RouterFun func(engine *gin.Engine)

/*
*@Author Administrator
*@Date 31/3/2021 11:36
*@desc
 */
func GinRun(port string, fnRouter ...RouterFun) (*http.Server, error) {
	engine := gin.New()
	gin.SetMode(gin.DebugMode)
	if len(fnRouter) > 0 {
		for _, fun := range fnRouter {
			fun(engine)
		}
	}
	server := Http(engine, port)
	return server, nil

}

/*
*@Author Administrator
*@Date 31/3/2021 11:41
*@desc
 */
func Http(engine *gin.Engine, port string) *http.Server {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	fmt.Println("|-----------------------------------|")
	fmt.Println("|            zngue gin run          |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("| Port:127.0.0.1:" + port + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "  |")
	fmt.Println("|-----------------------------------|")
	return server
}
