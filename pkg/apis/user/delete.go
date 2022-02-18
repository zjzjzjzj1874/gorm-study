package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
	"net/http"
	"strconv"
)

func Delete(c *gin.Context) {
	var (
		err error
	)

	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_, err = database.Delete(c, uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
