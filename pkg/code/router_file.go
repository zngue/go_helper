package code

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

/*
*@Author Administrator
*@Date 9/4/2021 12:17
*@desc
 */
func RouterFile(SelectTableName string) {
	f2 := new(FileNameChange)
	model := f2.Case2Camel(SelectTableName)
	router := f2.Lcfirst(model)
	all := strings.ReplaceAll(routerTemp, "{{model}}", model)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	routerpath := viper.GetString("temp.routerpath")
	routerpath = path + "/" + routerpath
	CreateMutiDir(routerpath)
	modelFile := routerpath + "/" + SelectTableName + ".go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	all = strings.ReplaceAll(all, "{{table}}", SelectTableName)
	all = strings.ReplaceAll(all, "{{router}}", router)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}
