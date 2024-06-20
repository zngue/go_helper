package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type ListRequest[R any] struct {
	Page *models.Page
	Data *R
}
type ListRequestFn[R any] func(req *models.ListRequest, data *ListRequest[R])

type ItemFn[T, RS any] func(data *T) *RS

type List[T, R, RS any] struct {
	callback     ListRequestFn[R]
	itemCallback ItemFn[T, RS]
}

func (l *List[T, R, RS]) List(ctx *gin.Context) (items []*RS, err error) {
	if l.itemCallback == nil {
		err = errors.New("itemCallback is nil")
		return
	}
	var req = new(ListRequest[R])
	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	listReq := new(models.ListRequest)
	if req.Page != nil {
		page := models.NewPage(req.Page.Page, req.Page.PageSize)
		req.Page = page
	}
	listReq.Order = []string{"id desc"}
	if l.callback != nil {
		l.callback(listReq, req)
	}
	var dataList []*T
	dataList, err = models.NewDB[T]().List(listReq)
	if err != nil {
		return
	}
	if len(dataList) > 0 {
		for _, item := range dataList {
			var val = l.itemCallback(item)
			if val != nil {
				items = append(items, val)
			}
		}
	}
	return
}

type IList[T any] interface {
	List(ctx *gin.Context) (items []*T, err error)
}

// NewList T model 数据模型  R request 请求数据 RS response 响应数据
func NewList[T, R, RS any](fn ListRequestFn[R], itemCallback ItemFn[T, RS]) IList[RS] {
	var list = new(List[T, R, RS])
	list.callback = fn
	list.itemCallback = itemCallback
	return list
}
