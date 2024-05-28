package api

import "github.com/gin-gonic/gin"

type CommonService[T any] interface {
	Add(ctx *gin.Context) (err error)
	Update(ctx *gin.Context) (err error)
	Delete(ctx *gin.Context) (err error)
	List(ctx *gin.Context) (items []*T, err error)
	ListPage(ctx *gin.Context) (list []*T, count int64, err error)
	UpdateFiled(ctx *gin.Context) (err error)
	Content(ctx *gin.Context) (content *T, err error)
	Status(ctx *gin.Context) (err error)
}
