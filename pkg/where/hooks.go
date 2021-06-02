package where

import (
	"gorm.io/gorm"
	"strings"
)

/*
*@Author Administrator
*@Date 12/5/2021 12:50
*@desc
 */
func init() {

	RegsterHooks(ResiterHooksOption{
		Where: "between",
		Hooks: func(option *HooksOption) *gorm.DB {
			split := strings.Split(option.Value.String(), ",")
			return option.DB.Where(option.Field+" BETWEEN  ?  AND ?", split[0], split[1])
		},
		Action: func(option *HooksOption) bool {

			va := option.Value.String()
			split := strings.Split(va, ",")
			if va != "" && len(split) == 2 {
				return true
			}
			return false

		},
	})

}
