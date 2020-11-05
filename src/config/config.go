package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var (
	defaultConfigPath = *flag.String("config", "config/users.json", "customize config path")
)

// Database 数据库连接配置
type Database struct {
	DSN string `json:"dsn"`
}

type configuration struct {
	Databases     []Database `json:"databases"`
	PublicUrl     string     `json:"publicUrl"`
	ListenAddress string     `json:"listenAddress"`
}

// LoadCustomize 加载自定义路径配置
func LoadCustomize(pathName string) configuration {
	var result configuration

	buffer, err := ioutil.ReadFile(pathName)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(buffer, &result)
	if err != nil {
		panic(err.Error())
	}

	return result
}

// LoadDefault 加载环境变量默认配置
func LoadDefault() configuration {
	return LoadCustomize(defaultConfigPath)
}
