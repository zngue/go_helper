package where

import (
	"gorm.io/gorm"
	"strings"
)

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
