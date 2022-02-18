package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zjzjzjzj1874/gorm-study/pkg/apis"
	"github.com/zjzjzjzj1874/gorm-study/pkg/apis/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.PUT("/migrate", apis.Migrate)

	{
		v0 := router.Group("/v0")
		userGroup := v0.Group("/user")
		userGroup.POST("", user.Create)
	}

	return router
}
