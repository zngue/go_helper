package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/app/service"
	"github.com/zngue/go_helper/pkg/util"
)

type Rn func(api.CommonApi) *DataInfo
type DataInfo struct {
	Fn   gin.HandlerFunc
	Path string
}

func DataApiContent(data api.CommonApi) *DataInfo {
	return

}

func Router[T any](routerName string, data service.IService[T], apiRouter *gin.RouterGroup, uris ...api.RouterConst) (router *gin.RouterGroup) {
	var apiData = NewCommon[T](data)
	router = apiRouter.Group(routerName)
	if len(uris) > 0 {
		if util.InArray(api.Content, uris) {
			router.GET("content", apiData.Content())
		}
		if util.InArray(api.List, uris) {
			router.GET("list", apiData.List())
		}
		if util.InArray(api.ListPage, uris) {
			router.GET("listPage", apiData.ListPage())
		}
		if util.InArray(api.UpdateFiled, uris) {
			router.POST("updateFiled", apiData.UpdateFiled())
		}
		if util.InArray(api.Status, uris) {
			router.POST("status", apiData.Status())
		}
		if util.InArray(api.Add, uris) {
			router.POST("add", apiData.Add())
		}
		if util.InArray(api.Update, uris) {
			router.POST("update", apiData.Update())
		}
		if util.InArray(api.Delete, uris) {
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

type Common[T any] struct {
	common service.IService[T]
}

func NewCommon[T any](common service.IService[T]) api.CommonApi {
	data := new(Common[T])
	data.common = common
	return data
}
func (c *Common[T]) Content() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var content, err = c.common.Content(ctx)
		api.DataWithErr(ctx, err, content)
	}
}

func (c *Common[T]) Status() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Status(ctx)
		api.DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) UpdateFiled() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.UpdateField(ctx)
		api.DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Add(ctx)
		api.DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Update(ctx)
		api.DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err = c.common.Delete(ctx)
		api.DataWithErr(ctx, err, nil)
	}
}

func (c *Common[T]) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, err = c.common.List(ctx)
		api.DataWithErr(ctx, err, list)
	}
}

func (c *Common[T]) ListPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var list, count, err = c.common.ListPage(ctx)
		api.DataWithErr(ctx, err, gin.H{
			"list":  list,
			"count": count,
		})
	}
}
