package pkg

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type GormExt struct {
	DB          *gorm.DB
	I           interface{}
	Order       map[string]interface{}
	OrderString string
}

func (g *GormExt) Init() *gorm.DB {
	g.DB = Where(g.DB, g.I)
	g.DB = Order(g.DB, g.Order, g.OrderString)
	return g.DB
}
func whereSeparator(db *gorm.DB, where string, field string, v string) *gorm.DB {
	switch where {
	case "eq":
		db = db.Where(field+" = ? ", v)
	case "neq":
		db = db.Where(field+" != ? ", v)
	case "gt":
		db = db.Where(field+" > ? ", v)
	case "egt":
		db = db.Where(field+" >= ? ", v)
	case "lt":
		db = db.Where(field+" < ? ", v)
	case "elt":
		db = db.Where(field+" <= ? ", v)
	case "null":
		db = db.Where(field + " is null  ")
	case "notnull":
		db = db.Where(field + " is not null ")
	case "like":
		db = db.Where(field+" like ?", "%"+v+"%")
	case "in":
		arr := strings.Split(v, ",")
		db = db.Where(field+" in (?)", arr)
	case "notin":
		arr := strings.Split(v, ",")
		db = db.Where(field+" not in (?)", arr)
	}
	return db
}
func Where(dbs *gorm.DB, i interface{}) *gorm.DB {
	if dbs == nil || i == nil {
		return dbs
	}
	var db *gorm.DB
	db = dbs
	refType := reflect.TypeOf(i)
	refValue := reflect.ValueOf(i)
	for i := 0; i < refValue.NumField(); i++ {
		f := refType.Field(i)
		valueInterface := refValue.Field(i)
		if valueInterface.Kind() == reflect.Ptr {
			continue
		}
		if valueInterface.Kind() == reflect.Struct && valueInterface.Interface() != nil {
			db = Where(db, valueInterface.Interface())
		}
		value := refValue.Field(i).Interface()
		if f.Tag == "" {
			continue
		}
		where := f.Tag.Get("where")
		if where == "or" && value != nil {
			db = db.Or(value)
			continue
		}
		field := f.Tag.Get("field")
		defaults := f.Tag.Get("default")
		newVal := cast.ToString(value)
		if defaults != newVal && field != "" {
			db = whereSeparator(db, where, field, newVal)
		}
	}
	return db
}
func Order(dbs *gorm.DB, order map[string]interface{}, orderString string) *gorm.DB {
	var db *gorm.DB
	db = dbs
	if order != nil && len(orderString) > 0 {
		StringArr := strings.Split(orderString, ",")
		if len(StringArr) > 0 {
			for _, s := range StringArr {
				if val, ok := order[s]; ok {
					db = db.Order(val)
				}
			}
		}
	}
	return db
}
