package pkg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

/*
*@Author Administrator
*@Date 31/3/2021 11:06
*@desc
 */
func NewConfig(Name, path, ext string) error {

	config := &Config{
		Name: Name,
		Path: path,
		Ext:  ext,
	}
	return config.InitConfig()
}

type Config struct {
	Name string //文件名称
	Path string //文件路径
	Ext  string //文件后缀
}

func (c *Config) InitConfig() error {
	viper.AddConfigPath(c.Path)
	viper.SetConfigName(c.Name)
	viper.SetConfigType(c.Ext)
	// 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	slice := viper.GetStringSlice("LoadConfig")
	for _, FileName := range slice {
		viper.SetConfigName(FileName)
		viper.SetConfigType(c.Ext)
		err := viper.MergeInConfig()
		if err != nil {
			return err
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name)
	})
	return nil
}
