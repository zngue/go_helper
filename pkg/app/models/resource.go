package models

import (
	"errors"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type Page struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

//分页处理

// PageHandle Page 分页处理
func (p *Page) PageHandle(db *gorm.DB) *gorm.DB {
	if p.Page == -1 { //不分页
		return db
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	offset := (p.Page - 1) * p.PageSize
	return db.Offset(int(offset)).Limit(int(p.PageSize))
}
func NewPage(pagePageSize ...int) *Page {
	var page, pageSize int
	if len(pagePageSize) == 0 {
		page = -1
	}
	if len(pagePageSize) > 0 {
		page = pagePageSize[0]
	}
	if len(pagePageSize) > 1 {
		pageSize = pagePageSize[1]
	}
	return &Page{Page: page, PageSize: pageSize}
}

type Fn func(db *gorm.DB) *gorm.DB

type DB[T any] struct {
	Source *gorm.DB
	Model  *T
}
type ListRequest struct {
	*Page
	Where  map[string]any
	Order  []string
	Select any
	Fn     Fn
}
type ContentRequest struct {
	Where  map[string]any
	Select any
	Fn     Fn
	Order  []string
}

// Content 获取单条数据
func (d *DB[T]) Content(data *ContentRequest) (resData *T, err error) {
	db := d.Source.Model(d.Model)
	db = d.ListHelper(db, &ListRequest{
		Where:  data.Where,
		Order:  data.Order,
		Select: data.Select,
		Fn:     data.Fn,
	})
	err = db.First(&resData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("数据不存在")
	}
	return
}

type ListFn func(req *ListRequest) *ListRequest

func ListPagaFn(pagePageSize ...int) ListFn {
	var page = NewPage(pagePageSize...)
	return func(req *ListRequest) *ListRequest {
		req.Page = page
		return req
	}
}
func ListWhereFn(where map[string]any) ListFn {
	return func(req *ListRequest) *ListRequest {
		req.Where = where
		return req
	}
}
func ListOrderFn(order []string) ListFn {
	return func(req *ListRequest) *ListRequest {
		req.Order = order
		return req
	}
}
func ListSelectFn(selects any) ListFn {
	return func(req *ListRequest) *ListRequest {
		req.Select = selects
		return req
	}
}
func ListReqFn(fn Fn) ListFn {
	return func(req *ListRequest) *ListRequest {
		req.Fn = fn
		return req
	}
}
func (d *DB[T]) ListPageFn(fns ...ListFn) (list []*T, count int64, err error) {
	var data = new(ListRequest)
	if len(fns) > 0 {
		for _, v := range fns {
			data = v(data)
		}
	}
	return d.ListPage(data)
}

// ListFn 获取列表
func (d *DB[T]) ListFn(fns ...ListFn) (list []*T, err error) {
	var data = new(ListRequest)
	if len(fns) > 0 {
		for _, v := range fns {
			data = v(data)
		}
	}
	return d.List(data)
}

// List 获取列表
func (d *DB[T]) List(data *ListRequest) (list []*T, err error) {
	if data.Page == nil {
		data.Page = new(Page)
	}
	db := d.Source.Model(d.Model)
	db = d.ListHelper(db, data)
	if data.Page.Page != -1 {
		db = data.Page.PageHandle(db)
	}
	err = db.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

// ListHelper ListRequest
func (d *DB[T]) ListHelper(db *gorm.DB, data *ListRequest) *gorm.DB {

	if data.Where != nil && len(data.Where) > 0 {
		db = d.Where(data.Where, db)
	}
	if data.Order != nil && len(data.Order) > 0 {
		db = d.Order(data.Order, db)
	}
	if data.Select != nil {
		db = d.Select(db, data.Select)
	}
	if data.Fn != nil {
		db = data.Fn(db)
	}
	return db
}

// ListPage 获取列表带分页
func (d *DB[T]) ListPage(data *ListRequest) (list []*T, count int64, err error) {
	db := d.Source.Model(d.Model)
	db = d.ListHelper(db, data)
	if data.Page != nil && data.Page.Page != -1 {
		err = db.Count(&count).Error
		if err != nil {
			return
		}
		db = data.Page.PageHandle(db)
	}
	err = db.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

// Select 设置查询字段
func (d *DB[T]) Select(db *gorm.DB, data any) *gorm.DB {
	if data != nil {
		db = db.Select(data)
	}
	return db
}

// Add 新增
func (d *DB[T]) Add(data *T) (err error) {
	db := d.Source.Model(d.Model)
	err = db.Create(data).Error
	return
}

// AddMore 批量新增
func (d *DB[T]) AddMore(data []*T) (err error) {
	db := d.Source.Model(d.Model)
	err = db.Create(data).Error
	return
}

// Where map[string]any where条件
func (d *DB[T]) Where(data map[string]any, db *gorm.DB) *gorm.DB {

	if data != nil && len(data) > 0 {
		for k, v := range data {
			db = db.Where(k, v)
		}
	}
	return db
}

// Order []string 排序条件
func (d *DB[T]) Order(data []string, db *gorm.DB) *gorm.DB {
	if data != nil && len(data) > 0 {
		for _, v := range data {
			db = db.Order(v)
		}
	}
	return db
}

// Update 更新 where map  data map
func (d *DB[T]) Update(where, data map[string]any) (err error) {
	db := d.Source.Model(d.Model)
	if where == nil || len(where) == 0 {
		err = errors.New("where条件不能为空")
		return
	}
	db = d.Where(where, db)
	err = db.Updates(data).Error
	return
}

// Delete 删除
func (d *DB[T]) Delete(where map[string]any) (err error) {
	db := d.Source.Model(d.Model)
	if where == nil || len(where) == 0 {
		err = errors.New("where条件不能为空")
		return
	}
	db = d.Where(where, db)
	err = db.Delete(d.Model).Error
	return
}

type DelRequest struct {
	Where map[string]any
	Fn    Fn
}

// DeleteFn 关联删除
func (d *DB[T]) DeleteFn(data *DelRequest) (err error) {
	db := d.Source.Model(d.Model)
	if data.Where == nil || len(data.Where) == 0 {
		err = errors.New("where条件不能为空")
		return
	}
	db = d.Where(data.Where, db)
	if data.Fn != nil {
		db = data.Fn(db)
	}
	err = db.Delete(d.Model).Error
	return
}

// NewDB 实例化 DB
func NewDB[T any]() *DB[T] {
	model := new(T)
	return &DB[T]{
		Source: pkg.MysqlConn,
		Model:  model,
	}
}
