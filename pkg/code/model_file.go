package code

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
	"log"
	"strings"
)

type TableList struct {
	TableName    string `json:"table_name" `
	TableComment string `json:"table_comment" `
}
type TableInfo struct {
	COLUMN_NAME    string
	IS_NULLABLE    string
	DATA_TYPE      string
	COLUMN_KEY     string
	EXTRA          string
	COLUMN_COMMENT string
}

/*
*@Author Administrator
*@Date 9/4/2021 13:38
*@desc
 */
func TableListAll(mysqldb string) []TableList {
	var tableList []TableList
	err3 := pkg.MysqlConn.Raw("select table_name from information_schema.tables where table_schema=?", mysqldb).Scan(&tableList).Error
	if err3 != nil {
		log.Fatal(err3)
	}
	return tableList
}

/*
*@Author Administrator
*@Date 9/4/2021 13:40
*@desc
 */
func GetDadatase() string {

	mysqldb := viper.GetString("db.DATABASE")
	if len(mysqldb) == 0 {
		log.Fatal("数据库不能为空")
	}
	modepath := viper.GetString("temp.modepath")
	if modepath == "" {
		log.Fatal("go mod 地址不能为空")
	}
	return mysqldb
}

/*
*@Author Administrator
*@Date 9/4/2021 13:46
*@desc
 */
func CodeModelList() {
	pkg.NewMysql()
	mysqldb := GetDadatase()
	tableList := TableListAll(mysqldb)
	fileName := new(FileNameChange)
	var SelectType int = -1
	var flag bool = false
	for {
		fmt.Println("请输入编号1单表创建，2所有表：")
		fmt.Scanln(&SelectType)
		switch SelectType {
		case 1: //单表创建
			flag = true
		case 2: //单表所有表
			flag = true
		default:
			if SelectType != 1 || SelectType != 2 {
				fmt.Println("输入错误请重新输入")
			}
			flag = false
		}
		if flag {
			break
		}
		continue
	}
	if SelectType == 1 {
		tableName := ModeCode(tableList, pkg.MysqlConn, mysqldb, fileName)
		RequestFile(tableName)
		ServiceFile(tableName)
		ControllerFile(tableName)
		RouterFile(tableName)
	} else {
		for _, table := range tableList {
			OneTable(pkg.MysqlConn, mysqldb, table.TableName, fileName)
			tableName := table.TableName
			RequestFile(tableName)
			ServiceFile(tableName)
			ControllerFile(tableName)
			RouterFile(tableName)
		}

	}

}

/*
*@Author Administrator
*@Date 8/4/2021 14:44
*@desc
 */
func ModeCode(tableList []TableList, mysql *gorm.DB, mysqldb string, fileName *FileNameChange) string {
	for i, table := range tableList {
		fmt.Println(fmt.Sprintf("%d、%s", i, table.TableName))
	}
	var SelectTableName string
	var SelectNum int = -1
	for {

		fmt.Println("请输入对应数据库编号：")
		_, err4 := fmt.Scanln(&SelectNum)
		if err4 != nil {
			fmt.Println(err4)
			continue
		}
		if SelectNum > len(tableList) || SelectNum < 0 {
			fmt.Println("编号错误，请输入有效表名编号")
			continue
		} else {
			SelectTableName = tableList[SelectNum].TableName
			break
		}
	}
	OneTable(mysql, mysqldb, SelectTableName, fileName)
	return SelectTableName
}

/*
*@Author Administrator
*@Date 9/4/2021 13:27
*@desc
 */
func OneTable(mysql *gorm.DB, mysqldb, SelectTableName string, fileName *FileNameChange) string {
	var tableInfoList []TableInfo
	err7 := mysql.Raw("select * from information_schema.columns where table_schema = ? and table_name =?", mysqldb, SelectTableName).Scan(&tableInfoList).Error
	fmt.Println(err7)
	var parameter []string
	var grpcArr []string
	for index, info := range tableInfoList {
		fileType := ""
		if info.COLUMN_NAME == "deleted_at" {
			fileType = "gorm.DeletedAt"
		} else {
			fileType = MysqlType(strings.ToUpper(info.DATA_TYPE))
		}
		file := fileName.Case2Camel(info.COLUMN_NAME)
		var pri string
		if info.COLUMN_KEY == "PRI" {
			pri = "primary_key;auto_increment;"
		}
		required := ""
		if info.IS_NULLABLE == "YES" {
			required = "required"
		}
		outFile := fileName.Lcfirst(fileName.Case2Camel(info.COLUMN_NAME))

		grpcString := fmt.Sprintf("%s %s=%d; //%s", MysqlTypeGrpc(strings.ToUpper(info.DATA_TYPE)), file, index+1, info.COLUMN_COMMENT)
		grpcArr = append(grpcArr, grpcString)
		gorm := fmt.Sprintf(`"%scolumn:%s;" form:"%s" json:"%s"  binding:"%s"`, pri, info.COLUMN_NAME, outFile, outFile, required)
		var sprintf string
		if info.COLUMN_COMMENT == "" {
			sprintf = fmt.Sprintf("%s %s `gorm:%s`", file, fileType, gorm)
		} else {
			sprintf = fmt.Sprintf("%s %s `gorm:%s`  //%s", file, fileType, gorm, info.COLUMN_COMMENT)
		}
		parameter = append(parameter, sprintf)
	}
	join := strings.Join(parameter, "\r\n\t")
	ModeFiles(SelectTableName, join)
	if viper.GetBool("temp.pbCode") {
		joinGrpc := strings.Join(grpcArr, "\r\n\t")
		ModePbFile(SelectTableName, joinGrpc)
	}
	return SelectTableName
}

func MysqlType(name string) string {
	switch name {
	case "TINYINT", "MEDIUMINT", "SMALLINT":
		return "int32"
	case "INT", "INTEGER", "BIGINT":
		return "int64"
	case "FLOAT", "DOUBLE", "DECIMAL":
		return "float64"
	case "DATE", "TIME", "YEAR", "TIMESTAMP", "CHAR", "VARCHAR", "TINYBLOB", "TINYTEXT", "BLOB", "TEXT", "MEDIUMBLOB", "MEDIUMTEXT", "LONGBLOB", "LONGTEXT":
		return "string"
	case "DATETIME":
		return "time.Time"
	}
	return ""
}
func MysqlTypeGrpc(name string) string {
	switch name {
	case "TINYINT", "MEDIUMINT", "SMALLINT":
		return "int32"
	case "INT", "INTEGER", "BIGINT":
		return "int64"
	case "FLOAT", "DOUBLE", "DECIMAL":
		return "float64"
	default:
		return "string"
	}
	return ""
}
