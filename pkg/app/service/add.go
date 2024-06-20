package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type Add[T, R any] struct {
	callback AddFn[T, R]
}
type AddFn[T, R any] func(req *R) *T

func (a *Add[T, R]) Add(ctx *gin.Context) (err error) {
	if a.callback == nil {
		err = errors.New("callback is nil")
		return
	}
	var req = new(R)
	if err = ctx.ShouldBindJSON(req); err != nil {
		return
	}
	data := a.callback(req)
	err = models.NewDB[T]().Add(data)
	return
}

type IAdd interface {
	Add(ctx *gin.Context) (err error)
}

func NewAdd[T, R any](callback AddFn[T, R]) IAdd {
	var t = new(Add[T, R])
	t.callback = callback
	return t
}
