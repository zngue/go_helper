package sign_chan

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

var SignChan chan os.Signal

func init() {
	SignChan = make(chan os.Signal)
}
func SignChalNotify() {
	signal.Notify(SignChan, os.Interrupt)
	<-SignChan
}
func SignLog(err ...interface{}) {
	log.Println(err)
	SignChan <- os.Interrupt
}

type CloseHttp func(ctx context.Context) error

func ListClose(fns CloseHttp) {
	SignChalNotify()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := fns(ctx)
	if err != nil {
		log.Fatalln(fmt.Sprintf("强制关闭"))
	} else {
		log.Println(fmt.Sprintf("第个关闭成功"))
	}
	log.Println("服务器优雅退出")
}
