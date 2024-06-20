package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type Delete[T any] struct {
}

func (d *Delete[T]) Delete(ctx *gin.Context) (err error) {
	var req = new(DeleteRequest)
	if err = ctx.ShouldBindJSON(req); err != nil {
		return
	}
	err = models.NewDB[T]().Delete(req.Condition)
	return
}

type DeleteRequest struct {
	Condition map[string]any `json:"condition"`
}
type IDelete interface {
	Delete(ctx *gin.Context) (err error)
}

func NewDelete[T any]() IDelete {
	return new(Delete[T])
}
