package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
)

func Get(c *gin.Context) {
	var (
		err error
	)

	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	u, err := database.Get(c, uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, u)
}
