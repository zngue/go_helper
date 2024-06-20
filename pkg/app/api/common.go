package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/app/service"
)

type FnApi func(api.CommonApi) *DataInfo

type DataInfo struct {
	Fn     gin.HandlerFunc
	Path   string
	Method Method
}

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

func DataApiContent() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.Content(),
			Path:   "content",
			Method: GET,
		}
	}
}
func DataApiList() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.List(),
			Path:   "list",
			Method: GET,
		}
	}
}
func DataApiListPage() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.Content(),
			Path:   "listPage",
			Method: GET,
		}
	}
}

func DataApiAdd() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.Add(),
			Path:   "add",
			Method: POST,
		}
	}
}
func DataApiUpdate() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.Update(),
			Path:   "update",
			Method: POST,
		}
	}
}
func DataApiUpdateField() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.UpdateFiled(),
			Path:   "updateField",
			Method: POST,
		}
	}
}
func DataApiDelete() FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     commonApi.Delete(),
			Path:   "delete",
			Method: POST,
		}
	}
}
func DataApiCommon(fn gin.HandlerFunc, Path string, method Method) FnApi {
	return func(commonApi api.CommonApi) *DataInfo {
		return &DataInfo{
			Fn:     fn,
			Path:   Path,
			Method: method,
		}
	}
}

func Router[T any](routerName string, dataServer service.IService[T], apiRouter *gin.RouterGroup, fns ...FnApi) (router *gin.RouterGroup) {
	var dataApi = NewCommon[T](dataServer)
	router = apiRouter.Group(routerName)
	var fnUri []*DataInfo
	for _, fn := range fns {
		info := fn(dataApi)
		if info != nil {
			fnUri = append(fnUri, info)
		}
	}
	if len(fnUri) > 0 {
		for _, infoFn := range fnUri {
			if infoFn.Method == POST {
				router.POST(infoFn.Path, infoFn.Fn)
			} else {
				router.GET(infoFn.Path, infoFn.Fn)
			}
		}
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
