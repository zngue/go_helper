package code

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

/*
*@Author Administrator
*@Date 9/4/2021 12:51
*@desc
 */
func ModeFiles(SelectTableName, st string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(modelTemp, "{{struct}}", st)
	path := viper.GetString("temp.path")
	modelpath := viper.GetString("temp.modelpath")
	modelpath = path + "/" + modelpath
	CreateMutiDir(modelpath)
	modelFile := modelpath + "/" + SelectTableName + ".go"
	all = strings.ReplaceAll(all, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{tableName}}", SelectTableName)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}

/*
*@Author Administrator
*@Date 13/4/2021 14:32
*@desc
 */
func ModePbFile(SelectTableName string, st string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(pbTemp, "{{service}}", st)
	path := viper.GetString("temp.path")
	pbpath := viper.GetString("temp.pbpath")
	pbpath = path + "/" + pbpath
	CreateMutiDir(pbpath)
	modelFile := pbpath + "/" + SelectTableName + ".proto"
	all = strings.ReplaceAll(all, "{{model}}", model)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}
