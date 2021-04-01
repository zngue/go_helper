package pkg

import (
	"fmt"
	"gorm.io/gorm"
)

type CommonRequest struct {
	Page        int  `form:"page" `      //当前页码
	IsPaginate  bool `form:"isPaginate"` //当前页码
	PageSize    int  `form:"pageSize"  ` //分页大小
	ReturnType  int  `form:"returnType"` //1 列表  2 map 列表  3 详情
	OrderMap    map[string]interface{}
	OrderString string `form:"orderString"`
	db          *gorm.DB
}

func (c *CommonRequest) OrderWhere(i interface{}) {
	if c.db == nil {
		fmt.Println("please first set db method SetDB()")
	}
	c.Paginate()
	ext := &GormExt{
		DB:          c.db,
		OrderString: c.OrderString,
		Order:       c.OrderMap,
		I:           i,
	}
	c.db = ext.Init()
	c.Paginate()
}
func (c *CommonRequest) PageLimitOffset() int {
	if c.Page == 0 {
		c.Page = 1
	}
	if c.PageSize == 0 {
		c.PageSize = 15
	}
	return (c.Page - 1) * c.PageSize
}
func (c *CommonRequest) Paginate() {
	if c.ReturnType == 3 {
		return
	}
	if c.IsPaginate {
		c.db = c.db.Offset(c.PageLimitOffset()).Limit(c.PageSize)
	} else {
		c.PageLimitOffset()
		c.db = c.db.Limit(c.PageSize)
	}
}
func (c *CommonRequest) WhereOr(ormap map[string]interface{}) {
	c.db = c.db.Or(ormap)
}
func (c *CommonRequest) SetDB(db *gorm.DB) {
	c.db = db
}
func (c *CommonRequest) GetDB() *gorm.DB {
	return c.db
}
