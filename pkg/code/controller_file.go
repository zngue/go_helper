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
func ControllerFile(SelectTableName string) {
	ControllerList(SelectTableName)
	ControllerDetail(SelectTableName)
	ControllerDelete(SelectTableName)
	ControllerAdd(SelectTableName)
	ControllerEdit(SelectTableName)
}

/*
*@Author Administrator
*@Date 22/4/2021 15:53
*@desc
 */
func ControllerDetail(SelectTableName string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerDetailTemp, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{controller}}", SelectTableName)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath = path + "/" + controllerpath + "/" + SelectTableName
	CreateMutiDir(controllerpath)
	modelFile := controllerpath + "/detail.go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}

/*
*@Author Administrator
*@Date 22/4/2021 15:49
*@desc
 */
func ControllerList(SelectTableName string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerListTemp, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{controller}}", SelectTableName)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath = path + "/" + controllerpath + "/" + SelectTableName
	CreateMutiDir(controllerpath)
	modelFile := controllerpath + "/list.go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}

/*
*@Author Administrator
*@Date 22/4/2021 15:49
*@desc
 */
func ControllerDelete(SelectTableName string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerDeleteTemp, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{controller}}", SelectTableName)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath = path + "/" + controllerpath + "/" + SelectTableName
	CreateMutiDir(controllerpath)
	modelFile := controllerpath + "/delete.go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}
func ControllerAdd(SelectTableName string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerAddTemp, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{controller}}", SelectTableName)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath = path + "/" + controllerpath + "/" + SelectTableName
	CreateMutiDir(controllerpath)
	modelFile := controllerpath + "/add.go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}
func ControllerEdit(SelectTableName string) {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerEditTemp, "{{model}}", model)
	all = strings.ReplaceAll(all, "{{controller}}", SelectTableName)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath = path + "/" + controllerpath + "/" + SelectTableName
	CreateMutiDir(controllerpath)
	modelFile := controllerpath + "/edit.go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}
