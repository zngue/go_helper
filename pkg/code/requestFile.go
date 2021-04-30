package code

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

/*
*@Author Administrator
*@Date 8/4/2021 15:00
*@desc
 */

var temp string

func RequestFile(SelectTableName string) {

	path := viper.GetString("temp.path")
	requestpath := viper.GetString("temp.requestpath")
	requestpath = path + "/" + requestpath
	CreateMutiDir(requestpath)
	modelFile := requestpath + "/" + SelectTableName + ".go"
	tableModel := new(FileNameChange).Case2Camel(SelectTableName)
	requestModelName := tableModel + "Request"
	temp = strings.ReplaceAll(requestTemp, "{{Request}}", requestModelName)
	temp = strings.ReplaceAll(temp, "{{model}}", tableModel)
	f, _ := os.Create(modelFile)
	f.Write([]byte(temp))
}
