package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type ContentRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}
type IContent[RS any] interface {
	Content(ctx *gin.Context) (content *RS, err error)
}
type Content[T, RS any] struct {
	itemCallback ItemFn[T, RS]
}

func (c *Content[T, RS]) Content(ctx *gin.Context) (content *RS, err error) {
	var req ContentRequest
	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	var data *T
	data, err = models.NewDB[T]().Content(&models.ContentRequest{
		Where: map[string]any{
			"id = ?": req.Id,
		},
	})
	if data != nil {
		content = c.itemCallback(data)
	}
	return
}

func NewContent[T, RS any](itemCallback ItemFn[T, RS]) IContent[RS] {
	var t = new(Content[T, RS])
	t.itemCallback = itemCallback
	return t
}
