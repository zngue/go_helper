package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type UpdateField[T any] struct {
}
type UpdateFieldRequest struct {
	Data      map[string]any `json:"data" binding:"required"`
	Condition map[string]any `json:"condition" binding:"required"`
}

func (u *UpdateField[T]) UpdateField(ctx *gin.Context) (err error) {
	var req = new(UpdateFieldRequest)
	if err = ctx.ShouldBindJSON(req); err != nil {
		return
	}
	err = models.NewDB[T]().Update(req.Condition, req.Data)
	return
}

type IUpdateField interface {
	UpdateField(ctx *gin.Context) (err error)
}

func NewUpdateField[T any]() IUpdateField {
	return new(UpdateField[T])

}
