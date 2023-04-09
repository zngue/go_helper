package db_resources

import (
	"errors"
	"gorm.io/gorm"
)

type DBResource[T any] struct {
	Resource *gorm.DB
	Model    *T
}
type Page struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
type WhereRequest struct {
	Where map[string]any
}

// PageLimit 写个分页查询的结构体
func (p *Page) PageLimit(db *gorm.DB) (tx *gorm.DB) {
	if p.Page == -1 {
		return db
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 20
	}
	db = db.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize)
	return db
}

type ListRequest struct {
	*Page
	Where map[string]any
	Order []string
	Fn    func(db *gorm.DB) *gorm.DB
}

// ListPage 获取列表带分页
func (d *DBResource[T]) ListPage(request *ListRequest) (dataList []*T, count int64, err error) {
	db := d.Resource.Model(d.Model)
	if request.Where != nil { //条件
		db = d.Where(request.Where, db)
	}
	if request.Order != nil { //排序
		db = d.Order(request.Order, db)
	}
	if request.Page != nil && request.Page.Page != -1 { //-1不分页
		err = db.Count(&count).Error
		if err != nil {
			return
		}
	}
	if request.Page != nil {
		db = request.Page.PageLimit(db)
	}
	if request.Fn != nil {
		db = request.Fn(db)
	}
	err = db.Find(&dataList).Error
	return
}

// Update 更新数据
func (d *DBResource[T]) Update(whereData, data map[string]any) (err error) {
	conn := d.Resource.Model(d.Model)
	if whereData != nil {
		conn = d.Where(whereData, conn)
	}
	err = conn.Updates(data).Error
	return
}

// Detail 获取详情
func (d *DBResource[T]) Detail(whereData map[string]any) (data *T, err error) {
	conn := d.Resource.Model(d.Model)
	if whereData != nil {
		conn = d.Where(whereData, conn)
	}
	err = conn.First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("操作失败")
	}
	return
}

// Delete 删除数据
func (d *DBResource[T]) Delete(whereData map[string]any) (err error) {
	conn := d.Resource.Model(d.Model)
	if whereData != nil {
		conn = d.Where(whereData, conn)
	}
	err = conn.Delete(d.Model).Error
	return
}

// Where 设置where条件
func (d *DBResource[T]) Where(where map[string]any, db *gorm.DB) *gorm.DB {
	if where != nil && len(where) > 0 {
		for k, v := range where {
			db = db.Where(k, v)
		}
	}
	return db
}

// Order 设置order 排序
func (d *DBResource[T]) Order(order []string, db *gorm.DB) *gorm.DB {
	if order != nil && len(order) > 0 {
		for _, v := range order {
			db = db.Order(v)
		}
	}
	return db
}

// NewDBResource 实例化 DBResource
func NewDBResource[T any](db *gorm.DB) *DBResource[T] {
	model := new(T)
	return &DBResource[T]{Resource: db, Model: model}
}
