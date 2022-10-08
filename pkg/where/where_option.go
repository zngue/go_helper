package where

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

func OptionWhereInit(db *gorm.DB, i interface{}) *gorm.DB {
	refType := reflect.TypeOf(i)
	refValue := reflect.ValueOf(i)
	for j := 0; j < refValue.NumField(); j++ {
		f := refType.Field(j)
		valueInterface := refValue.Field(j)
		if &valueInterface == nil {
			continue
		}
		if valueInterface.Kind() == reflect.Ptr {
			continue
		}
		value := valueInterface.Interface()
		if valueInterface.Kind() == reflect.Struct && value != nil {
			db = OptionWhereInit(db, value)
		}
		if f.Tag == "" {
			continue
		}
		if value == nil {
			continue
		}
		field := f.Tag.Get("field")
		defaults := f.Tag.Get("default")
		where := f.Tag.Get("where")
		if field == "" || where == "" {
			continue
		}
		if defaults != valueInterface.String() {
			whereMaps := fmt.Sprintf("%s %s ?", field, where)
			db = db.Where(whereMaps, whereMaps)
		}
	}
	return db
}
