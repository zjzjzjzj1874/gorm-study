package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
)

func Migrate(c *gin.Context) {
	database.Migrate()
}
