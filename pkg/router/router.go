package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zjzjzjzj1874/gorm-study/cmd/global"
	"github.com/zjzjzjzj1874/gorm-study/pkg/apis"
	"github.com/zjzjzjzj1874/gorm-study/pkg/apis/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	server := router.Group(global.GlobalConfig.Name)
	server.PUT("/migrate", apis.Migrate)

	{
		v0 := server.Group("/v0")
		userGroup := v0.Group("/user")
		userGroup.POST("", user.Create)
		userGroup.GET("/:id", user.Get)
		userGroup.DELETE("/:id", user.Delete)
	}

	return router
}
