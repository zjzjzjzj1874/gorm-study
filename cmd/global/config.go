package global

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name        string      `yaml:"name"` // 服务名
	Port        int         `yaml:"port"` // 端口
	MysqlInfo   MysqlConfig `yaml:"mysql"`
	RedisInfo   RedisConfig `yaml:"redis"`
	LogsAddress string      `yaml:"logsAddress"`
	Debug       bool        `yaml:"debug"` // 调试模式
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Pwd  string `yaml:"pwd"`
}

var (
	GlobalConfig ServerConfig // 全局配置文件
)

func Init() {
	// 实例化viper
	v := viper.New()
	v.SetConfigFile("cmd/config/local.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	// 传递给全局变量
	GlobalConfig = serverConfig
}

// 这里使用viper做解析,后面尝试	"gopkg.in/yaml.v2"来解析,err = yaml.Unmarshal(contents, &kv):contents为二进制文件,kv为一个map