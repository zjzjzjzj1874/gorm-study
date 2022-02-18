package database

import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	if db.Error != nil {
		panic(db.Error)
	}
}

func Close() {
	err := db.Close()
	if err != nil {
		logrus.Errorf("close DB failure:[err:%s]", err.Error())
	}
}

func Migrate() {
	db.AutoMigrate(&User{})
}
