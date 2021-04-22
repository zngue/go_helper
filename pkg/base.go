package pkg

import (
	"errors"
	"github.com/zngue/go_helper/pkg/where"
	"gorm.io/gorm"
)

var ResgterWhereHooks []where.ResiterHooksOption

type CommonRequest struct {
	Page        int  `form:"page" `      //当前页码
	IsPaginate  bool `form:"isPaginate"` //当前页码
	PageSize    int  `form:"pageSize"  ` //分页大小
	ReturnType  int  `form:"returnType"` //1 列表  2 map 列表  3 详情
	OrderMap    map[string]interface{}
	OrderString string `form:"orderString"`
	db          *gorm.DB
	Actions     int `form:"action"` //默认 1 修改 2 列表 3 单条数据 4删除 5添加
	Delete      int `form:"delete"` //是否包含删除数据
	Data        interface{}
	Error       error
}

func (c *CommonRequest) Init(db *gorm.DB, i interface{}) (tx *gorm.DB) {
	c.SetDB(db)
	c.OrderWhere(i)
	c.Action()
	return c.db
}

func (c *CommonRequest) Action() {
	switch c.Actions {
	case 1:
		if c.Data == nil {
			c.Error = errors.New("data is nil can not update")
			return
		}
		c.db = c.db.Updates(c.Data)
	case 2:
		c.db = c.db.Find(c.Data)
	case 3:
		c.db = c.db.First(c.Data)
	case 4:
		c.db = c.db.Delete(c.Data)
	case 5:
		c.db = c.db.Create(c.Data)
	}
}

func (c *CommonRequest) OrderWhere(i interface{}) {
	if c.db == nil {
		c.Error = errors.New("please first set db method SetDB()")
		return
	}
	c.db = where.NewGorm().Where(c.db, i)
	if c.OrderMap != nil && c.OrderString != "" {
		c.db = Order(c.db, c.OrderMap, c.OrderString)
	}
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
func (c *CommonRequest) SetDB(db *gorm.DB) {
	c.db = db
}
func (c *CommonRequest) GetDB() *gorm.DB {
	return c.db
}
