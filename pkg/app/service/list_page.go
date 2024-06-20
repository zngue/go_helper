package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type ListPage[T, R, RS any] struct {
	callback     ListRequestFn[R]
	itemCallback ItemFn[T, RS]
}

func (l *ListPage[T, R, RS]) ListPage(ctx *gin.Context) (items []*RS, count int64, err error) {
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
	dataList, count, err = models.NewDB[T]().ListPage(listReq)
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

type IListPage[RS any] interface {
	ListPage(ctx *gin.Context) (items []*RS, count int64, err error)
}

type ListCallback[T, R, RS any] struct {
	Fn           ListRequestFn[R]
	ItemCallback ItemFn[T, RS]
}

func NewListPage[T, R, RS any](fn ListRequestFn[R], itemCallback ItemFn[T, RS]) IListPage[RS] {
	var list = new(ListPage[T, R, RS])
	list.callback = fn
	list.itemCallback = itemCallback
	return list
}
