package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type DeleteFn[R any] func(req *R) (whereMap map[string]any)
type Delete[T, R any] struct {
	callback DeleteFn[R]
}

func (d *Delete[T, R]) Delete(ctx *gin.Context) (err error) {
	var req = new(R)
	if err = ctx.ShouldBind(req); err != nil {
		return
	}
	var where = d.callback(req) // 回调函数
	err = models.NewDB[T]().Delete(where)
	return
}

type IDelete interface {
	Delete(ctx *gin.Context) (err error)
}

func NewDelete[T, R any](callback DeleteFn[R]) IDelete {
	var t = new(Delete[T, R])
	t.callback = callback
	return t
}
