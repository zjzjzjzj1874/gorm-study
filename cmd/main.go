package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/zjzjzjzj1874/gorm-study/cmd/global"
	_ "github.com/zjzjzjzj1874/gorm-study/pkg/database"
	"github.com/zjzjzjzj1874/gorm-study/pkg/router"
)

func main() {
	global.Init() // 初始化配置文件

	gin := router.InitRouter() // 初始化路由

	if global.GlobalConfig.Debug {
		pprof.Register(gin) // pprof分析路由开启
	}
	if err := gin.Run(":8000"); err != nil {
		panic(err)
	}
}
