package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/util"
)

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
type Common[T any] struct {
	common CommonService[T]
}
type RouterConst string

const Add RouterConst = "add"
const Update RouterConst = "update"
const Delete RouterConst = "delete"
const List RouterConst = "list"
const ListPage RouterConst = "listPage"
const UpdateFiled RouterConst = "updateFiled"
const Content RouterConst = "content"
const Status RouterConst = "status"

func NewCommon[T any](common CommonService[T]) CommonApi {
	data := new(Common[T])
	data.common = common
	return data
}

type RouterFn func(router *gin.RouterGroup)

func Router[T any](routerName string, data CommonService[T], api *gin.RouterGroup, uris ...RouterConst) (router *gin.RouterGroup) {
	var apiData = NewCommon[T](data)
	router = api.Group(routerName)
	if len(uris) > 0 {
		if util.InArray(Content, uris) {
			router.GET("content", apiData.Content())
		}
		if util.InArray(List, uris) {
			router.GET("list", apiData.List())
		}
		if util.InArray(ListPage, uris) {
			router.GET("listPage", apiData.ListPage())
		}
		if util.InArray(UpdateFiled, uris) {
			router.POST("updateFiled", apiData.UpdateFiled())
		}
		if util.InArray(Status, uris) {
			router.POST("status", apiData.Status())
		}
		if util.InArray(Add, uris) {
			router.POST("add", apiData.Add())
		}
		if util.InArray(Update, uris) {
			router.POST("update", apiData.Update())
		}
		if util.InArray(Delete, uris) {
			router.POST("delete", apiData.Delete())
		}
	} else {
		router.GET("content", apiData.Content())
		router.GET("list", apiData.List())
		router.GET("listPage", apiData.ListPage())
		router.POST("updateFiled", apiData.UpdateFiled())
		router.POST("status", apiData.Status())
		router.POST("add", apiData.Add())
		router.POST("update", apiData.Update())
		router.POST("delete", apiData.Delete())
	}

	return
}

func (c *Common[T]) Content() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var content, err = c.common.Content(ctx)
		DataWithErr(ctx, err, content)
	}
}

func (c *Common[T]) Status() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Status(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) UpdateFiled() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.UpdateFiled(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Add(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Update(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Delete(ctx)
		DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, err = c.common.List(ctx)
		DataWithErr(ctx, err, list)
	}
}

func (c *Common[T]) ListPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, count, err = c.common.ListPage(ctx)
		DataWithErr(ctx, err, gin.H{
			"list":  list,
			"count": count,
		})
	}
}
