package main

import (
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
	"github.com/zjzjzjzj1874/gorm-study/pkg/router"
)

func main() {
	defer database.Close() // 关闭数据库连接

	gin := router.InitRouter() // 初始化路由
	if err := gin.Run(":8000"); err != nil {
		panic(err)
	}
}
