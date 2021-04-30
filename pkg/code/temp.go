package code

var serviceTemp string = `package service

import (
	"{{path}}/model"
	"{{path}}/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type {{model}}Service struct {
	
}
type {{model}}ServiceInterface interface {
	List(req *request.{{model}}Request) (*[]model.{{model}},error)
	Detail(req *request.{{model}}Request) (*model.{{model}},error)
	Delete(req *request.{{model}}Request) (err error)
	Add(req *request.{{model}}Request) (err error)
	Edit(req *request.{{model}}Request) (err error)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func ( *{{model}}Service ) InitModelDB(req *request.{{model}}Request)(tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.{{model}}{})
	return req.Common(db)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *{{model}}Service) List(req *request.{{model}}Request) (*[]model.{{model}},error) {
	req.Actions=2
	var list []model.{{model}}
	req.Data=&list
	err := r.InitModelDB(req).Error
	return &list,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *{{model}}Service) Detail(req *request.{{model}}Request) (*model.{{model}},error) {
	req.Actions=3
	var one model.{{model}}
	req.Data=&one
	err := r.InitModelDB(req).Error
	return &one,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *{{model}}Service) Add(req *request.{{model}}Request) (err error) {
	req.Actions=5
	//添加数据请自行处理
	err = r.InitModelDB(req).Error
	return err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *{{model}}Service) Edit(req *request.{{model}}Request) (err error){
	req.Actions=1
	//req.Data= [...] 修改数据请自行处理好
	err = r.InitModelDB(req).Error
	return err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *{{model}}Service) Delete(req *request.{{model}}Request) (err error){
	req.Actions=4
	err = r.InitModelDB(req).Error
	return  err
}
/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func New{{model}}Service() {{model}}ServiceInterface {
	return new({{model}}Service)
}
`
var requestTemp string = "package request\nimport (\n\t\"github.com/zngue/go_helper/pkg\"\n\t\"gorm.io/gorm\"\n)\ntype {{model}}Request struct {\n\tpkg.CommonRequest\n\tID int `form:\"id\" field:\"id\" where:\"eq\" default:\"0\"`\n}\nfunc (r *{{model}}Request) Common(db *gorm.DB) *gorm.DB {\n\ttx := r.Init(db, *r)\n\treturn tx\n}"

var controllerListTemp string = `package {{controller}}

import (
	"github.com/gin-gonic/gin"
	"{{path}}/request"
	"{{path}}/service"
	"github.com/zngue/go_helper/pkg/response"
)
func  List(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	res, err := service.New{{model}}Service().List(&req)
	response.HttpSuccessWithError(ctx,err,res)
	return
}
`

var controllerDetailTemp string = `package {{controller}}

import (
	"github.com/gin-gonic/gin"
	"{{path}}/request"
	"{{path}}/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func  Detail(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	res,err := service.New{{model}}Service().Detail(&req)
	response.HttpSuccessWithError(ctx,err,res)
	return
}
`
var controllerDeleteTemp string = `package {{controller}}

import (
	"github.com/gin-gonic/gin"
	"{{path}}/request"
	"{{path}}/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func Delete(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	err := service.New{{model}}Service().Delete(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
`

var controllerEditTemp string = `package {{controller}}

import (
	"github.com/gin-gonic/gin"
	"{{path}}/request"
	"{{path}}/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func  Edit(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	if err := ctx.Request.ParseForm(); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	postForm := ctx.Request.PostForm
	data := make(map[string]interface{})
	for key, val := range postForm {
		data[key]=val
	}
	req.Data=data
	err := service.New{{model}}Service().Edit(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
`

var controllerAddTemp string = `package {{controller}}

import (
	"github.com/gin-gonic/gin"
	"{{path}}/model"
	"{{path}}/request"
	"{{path}}/service"
	"github.com/zngue/go_helper/pkg/response"
)
/*
*@Author Administrator
*@desc Auto_Code
 */
func Add(ctx *gin.Context) {
	var req request.{{model}}Request
	var data model.{{model}}
	if err := ctx.ShouldBind(&data); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	req.Data=&data
	err := service.New{{model}}Service().Add(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
`

var controllerTemp string = `package controller

import (
	"github.com/gin-gonic/gin"
	"{{path}}/request"
	"{{path}}/service"
	"{{path}}/model"
	"github.com/zngue/go_helper/pkg/response"
)










/*
*@Author Administrator
*@Date 9/4/2021 11:42
*@desc  修改
 */
func ( *{{model}} ) Edit(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	if err := ctx.Request.ParseForm(); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	postForm := ctx.Request.PostForm
	data := make(map[string]interface{})
	for key, val := range postForm {
		data[key]=val
	}
	req.Data=data
	err := service.New{{model}}Service().Edit(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
/*
*@Author Administrator
*@Date 9/4/2021 11:42
*@desc  删除
 */
func ( *{{model}} ) Delete(ctx *gin.Context) {
	var req request.{{model}}Request
	if err := ctx.ShouldBind(&req); err != nil {
		response.HttpParameterError(ctx,err)
		return
	}
	err := service.New{{model}}Service().Delete(&req)
	response.HttpSuccessWithError(ctx,err,nil)
}
`
var routerTemp string = `package router

import (
	"github.com/gin-gonic/gin"
	"{{path}}/controller/{{table}}"
)

/*
*@Author Administrator
*@Date 9/4/2021 11:48
*@desc
 */
func {{model}}Router(group *gin.RouterGroup)  {
	{{model}}Group := group.Group("{{router}}")
	{
		{{model}}Group.GET("list",{{table}}.List)
		{{model}}Group.GET("detail",{{table}}.Detail)
		{{model}}Group.POST("edit",{{table}}.Edit)
		{{model}}Group.POST("delete",{{table}}.Delete)
		{{model}}Group.POST("add",{{table}}.Add)
	}
}
`
var modelTemp string = `package model


type {{model}} struct {
	{{struct}}
}
func (m *{{model}}) TableName() string  {
	return "{{tableName}}"
}`

var pbTemp = `syntax="proto3";
option go_package=".;pb";
service {{model}}Service{
  rpc List({{model}}) returns ({{model}}ResponseList) {} //列表数据
  rpc Add({{model}}) returns ({{model}}Response) {} //添加数据
  rpc Detail({{model}}) returns ({{model}}Response) {} //单挑数据获取
  rpc Edit({{model}}) returns ({{model}}Response) {}//编辑数据
  rpc Delete({{model}}) returns ({{model}}Response) {}//删除数据
}
message {{model}}Response{
  int64  Code=1;
  string Message=2;
  {{model}} Data=3;
}
message {{model}}ResponseList{
  repeated {{model}} Data=3;
}
message {{model}}{
  {{service}}
}`
