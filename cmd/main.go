package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
	"github.com/zjzjzjzj1874/gorm-study/pkg/router"
)

func main() {
	defer database.Close() // 关闭数据库连接

	gin := router.InitRouter() // 初始化路由

	pprof.Register(gin) // pprof分析路由开启

	if err := gin.Run(":8000"); err != nil {
		panic(err)
	}
}
