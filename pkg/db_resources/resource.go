package db_resources

import (
	"encoding/json"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
	"time"
)

type Resource[T any] struct {
	db    *gorm.DB
	Model *T
}
type Page struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
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

// NewResource 新建一个结构体泛型 Resource
func NewResource[T any](db *gorm.DB) *Resource[T] {
	model := new(T)
	return &Resource[T]{
		db:    db,
		Model: model,
	}
}

// Update 更新数据
func (d *Resource[T]) Update(where, data map[string]any) (err error) {
	db := d.db.Model(d.Model)
	db = d.Where(where, db)
	err = db.Updates(data).Error
	return
}

// Create 新建数据
func (d *Resource[T]) Create(data *T) (err error) {
	err = d.db.Model(d.Model).Create(&data).Error
	return
}

// Where where条件
func (d *Resource[T]) Where(where map[string]any, db *gorm.DB) *gorm.DB {
	if where != nil && len(where) > 0 {
		for key, val := range where {
			db = db.Where(key, val)
		}
	}
	return db
}

// Order 排序
func (d *Resource[T]) Order(order []string, db *gorm.DB) *gorm.DB {
	if order != nil && len(order) > 0 {
		for _, val := range order {
			db = db.Order(val)
		}
	}
	return db
}

// Del 删除数据
func (d *Resource[T]) Del(where map[string]any) (err error) {
	db := d.db.Model(d.Model)
	db = d.Where(where, db)
	err = db.Delete(d.Model).Error
	return
}

// Content 查询单条数据
func (d *Resource[T]) Content(request *Request) (data *T, err error) {
	db := d.db.Model(d.Model)
	db = d.Order(request.Order, db)
	if request.Common != nil {
		if request.Common.Where != nil {
			db = d.Where(request.Common.Where, db)
		}
		if request.Common.Fn != nil {
			db = request.Common.Fn(db)
		}
		if request.Common.Select != nil {
			db = db.Select(request.Common.Select)
		}
		if request.Common.Order != nil {
			db = d.Order(request.Common.Order, db)
		}
		//Fn 不为nil
		if request.Common.Fn != nil {
			db = request.Fn(db)
		}
	}
	err = db.First(data).Error
	return
}

// Helper request *Request
func (d *Resource[T]) Helper(db *gorm.DB, request *Request) *gorm.DB {
	if request.Common != nil {
		if request.Common.Where != nil {
			db = d.Where(request.Common.Where, db)
		}
		if request.Common.Fn != nil {
			db = request.Common.Fn(db)
		}
		if request.Common.Select != nil {
			db = db.Select(request.Common.Select)
		}
		if request.Common.Order != nil {
			db = d.Order(request.Common.Order, db)
		}
		//Fn 不为nil
		if request.Common.Fn != nil {
			db = request.Fn(db)
		}
	}
	//分页设置
	if request.IsSingle { //单条数据
		return db
	}
	if request.Page != nil {
		if request.Page.Page != -1 {
			db = request.Page.PageLimit(db)
		}
	}
	return db

}

// List 查询多条数据
func (d *Resource[T]) List(request *Request) (data []*T, err error) {
	db := d.db.Model(d.Model)
	db = d.Helper(db, request)
	err = db.Find(&data).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

// ListPage 查询多条数据 带分页 count
func (d *Resource[T]) ListPage(request *Request) (dataList []*T, count int64, err error) {
	db := d.db.Model(d.Model)
	db = d.Helper(db, request)
	if request.Page != nil { // 分页
		if request.Page.Page != -1 { // 不查询全部的时候统计总数
			err = db.Count(&count).Error
			if err != nil {
				return
			}
		}
	}
	err = db.Find(&dataList).Error
	if err == gorm.ErrRecordNotFound { // 没有数据的时候不报错
		err = nil
	}
	return
}

// Count 查询总数
func (d *Resource[T]) Count(request *Request) (count int64, err error) {
	db := d.db.Model(d.Model)
	db = d.Helper(db, request)
	err = db.Count(&count).Error
	return
}

// Conn 获取数据库连接
func (d *Resource[T]) Conn(data *Request) *gorm.DB {
	db := d.db.Model(d.Model)
	db = d.Helper(db, data)
	return db
}

func Data[T any]() *Resource[T] {
	return NewResource[T](pkg.MysqlConn)
}

type CacheOption struct {
	// 缓存key
	Key        string
	ExpireTime time.Duration
	// 缓存数据FN
	CacheFn CacheFn
}

func DataCache(option *CacheOption, v any) error {
	return CacheCommon(option.Key, v, option.ExpireTime, option.CacheFn)
}

type CacheFn func() (err error, i any)

func CacheCommon(key string, v any, expireTime time.Duration, fn CacheFn) (err error) {
	var (
		redisValue string
		data       interface{}
	)
	redisValue, _ = pkg.RedisConn.Get(key).Result()

	if redisValue != "" {
		err = json.Unmarshal([]byte(redisValue), v)
		if err != nil {
			return
		}
		return
	}
	err, data = fn()
	if err != nil {
		return
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = pkg.RedisConn.Set(key, string(marshal), expireTime).Err()
	if err != nil {
		return
	}
	return
}
