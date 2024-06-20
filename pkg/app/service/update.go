package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type Update[T, R any] struct {
	fn UpdateToMap[R]
}
type UpdateToMap[R any] func(data *R) (whereMap, updateMap map[string]any)

func (u *Update[T, R]) Update(ctx *gin.Context) (err error) {
	if u.fn == nil {
		err = errors.New("fn is nil")
		return
	}
	var req = new(R)
	if err = ctx.ShouldBind(req); err != nil {
		return
	}
	var where, data = u.fn(req)
	err = models.NewDB[T]().Update(where, data)
	return
}

type IUpdate interface {
	Update(ctx *gin.Context) (err error)
}

func NewUpdate[T, R any](callback UpdateToMap[R]) IUpdate {
	var t = new(Update[T, R])
	t.fn = callback
	return t
}
