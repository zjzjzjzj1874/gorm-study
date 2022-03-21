package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/gorm-study/pkg/database"
	"github.com/zjzjzjzj1874/gorm-study/pkg/models"
)

func Create(c *gin.Context) {
	var (
		err    error
		params *models.CreateUser
	)

	if err = c.BindJSON(&params); err != nil {
		logrus.WithContext(c).Warnf("create user params err:%s", err.Error())
		return
	}

	u, err := database.Create(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, u)
}
