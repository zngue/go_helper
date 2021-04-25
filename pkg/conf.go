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

var configDefult = Config{
	Name: "app",
	Path: "conf",
	Ext:  "yaml",
}

type ConfigFn func(config *Config) *Config

/*
*@Author Administrator
*@Date 25/4/2021 14:36
*@desc 配置默认读取应用名字

 */
func Name(name string) ConfigFn {
	return func(config *Config) *Config {
		config.Name = name
		return config
	}
}

/*
*@Author Administrator
*@Date 25/4/2021 14:39
*@desc 配置路径
 */
func Path(path string) ConfigFn {
	return func(config *Config) *Config {
		config.Path = path
		return config
	}
}

/*
*@Author Administrator
*@Date 25/4/2021 14:40
*@desc 文件后缀
 */
func Ext(ext string) ConfigFn {
	return func(config *Config) *Config {
		config.Ext = ext
		return config
	}
}

func NewConfig(configList ...ConfigFn) error {

	if len(configList) > 0 {
		for _, fn := range configList {
			fn(&configDefult)
		}
	}
	return configDefult.InitConfig()
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
	mapString := viper.GetStringMapString("LoadConfig")
	for fileName, ext := range mapString {
		if ext == "" {
			ext = c.Ext
		}
		viper.SetConfigName(fileName)
		viper.SetConfigType(ext)
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
