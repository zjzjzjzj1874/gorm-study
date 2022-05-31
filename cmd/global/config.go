package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
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
	Endpoint      string          // mysql endpoint => "root:123456@tcp(localhost:3306)/scaffold?charset=utf8&parseTime=True&loc=Local",
	LogLevel      logger.LogLevel `json:",default=4,options=[1,2,3,4]"` // 日志等级
	SlowThreshold int             `json:",default=200"`                 // 慢sql判断条件（单位毫秒）
	LogPath       string          `json:",default=./logs/sql.log"`      // 日志文件
	Colorful      bool            `json:",optional"`                    // 彩色打印
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Pwd  string `yaml:"pwd"`
	DB   int    `yaml:"db"`
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
