package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/models"
)

type IService[T any] interface {
	IAdd
	IUpdate
	IUpdateField
	IList[T]
	IListPage[T]
	IContent[T]
	IStatus
	IDelete
}
type ModelService struct {
}
type Service struct {
}
type AddRequest struct {
}

func AddRequestDeal(req *AddRequest) *ModelService {
	return &ModelService{}
}
func (s *Service) Add(ctx *gin.Context) (err error) {
	return NewAdd[ModelService, AddRequest](AddRequestDeal).Add(ctx)
}

type UpdateRequest struct {
	Id int `json:"id" form:"id"`
}

func UpdateDataDeal(req *UpdateRequest) (whereMap, updateMap map[string]any) {
	whereMap = map[string]any{"id": req.Id}
	updateMap = map[string]any{"name": "test"}
	return
}
func (s *Service) Update(ctx *gin.Context) (err error) {
	return NewUpdate[ModelService](UpdateDataDeal).Update(ctx)
}

func (s *Service) UpdateField(ctx *gin.Context) (err error) {
	return NewUpdateField[ModelService]().UpdateField(ctx)
}

type ItemService struct {
}
type ListServiceRequest struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

func (s *Service) List(ctx *gin.Context) (items []*ItemService, err error) {
	return NewList(ListServiceWhere, ModelChangeItem).List(ctx)
}

func ListServiceWhere(modelRequest *models.ListRequest, req *ListRequest[ListServiceRequest]) {
	var where = make(map[string]any)
	if req.Data != nil {
		if req.Data.Id > 0 {
			where["id = ?"] = req.Data.Id
		}
		if req.Data.Name != "" {
			where["name like ?"] = fmt.Sprintf("%s%%", req.Data.Name)
		}
	}
	modelRequest.Where = where
	if req.Page != nil {
		modelRequest.Page = req.Page
	}
}
func (s *Service) ListPage(ctx *gin.Context) (items []*ItemService, count int64, err error) {
	return NewListPage(ListServiceWhere, ModelChangeItem).ListPage(ctx)
}

type ContentServiceRequest struct {
}

func (s *Service) Content(ctx *gin.Context) (content *ItemService, err error) {
	return NewContent[ModelService, ItemService](ModelChangeItem).Content(ctx)
}
func ModelChangeItem(req *ModelService) (data *ItemService) {
	return data
}

func (s *Service) Status(ctx *gin.Context) (err error) {
	return NewStatus[ModelService]().Status(ctx)
}

type DelRequest struct {
}

func DelWhere(req *DelRequest) (whereMap map[string]any) {

	return
}
func (s *Service) Delete(ctx *gin.Context) (err error) {

	return NewDelete[ModelService, DelRequest](DelWhere).Delete(ctx)
}

func NewService() IService[ItemService] {

	return new(Service)
}
