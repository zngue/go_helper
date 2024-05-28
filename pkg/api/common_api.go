package api

import "github.com/gin-gonic/gin"

type CommonApi interface {
	Content() gin.HandlerFunc
	Status() gin.HandlerFunc
	UpdateFiled() gin.HandlerFunc
	Add() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	List() gin.HandlerFunc
	ListPage() gin.HandlerFunc
}
type Common[T any, C any] struct {
	common CommonService[T, C]
}

func NewCommon[T any, C any](common CommonService[T, C]) CommonApi {
	data := new(Common[T, C])
	data.common = common
	return data
}

type RouterFn func(router *gin.RouterGroup)

func Router[T any, C any](routerName string, data CommonService[T, C], api *gin.RouterGroup, fns ...RouterFn) {
	var apiData = NewCommon[T, C](data)
	var router = api.Group(routerName)
	router.GET("content", apiData.Content())
	router.GET("list", apiData.List())
	router.GET("listPage", apiData.ListPage())
	router.POST("updateFiled", apiData.UpdateFiled())
	router.POST("status", apiData.Status())
	router.POST("add", apiData.Add())
	router.POST("update", apiData.Update())
	router.POST("delete", apiData.Delete())
	if len(fns) > 0 {
		for _, fn := range fns {
			fn(router)
		}
	}
	return
}

func (c *Common[T, C]) Content() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var content, err = c.common.Content(ctx)
		DataWithErr(ctx, err, content)
	}
}

func (c *Common[T, C]) Status() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Status(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T, C]) UpdateFiled() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.UpdateFiled(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T, C]) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Add(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T, C]) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Update(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T, C]) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Delete(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T, C]) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, err = c.common.List(ctx)
		DataWithErr(ctx, err, list)
	}
}

func (c *Common[T, C]) ListPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, count, err = c.common.ListPage(ctx)
		DataWithErr(ctx, err, gin.H{
			"list":  list,
			"count": count,
		})
	}
}
