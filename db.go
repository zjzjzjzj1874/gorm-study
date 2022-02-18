package main

import (
	"fmt"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     string `gorm:"size:255"`       // 设置字段大小为255
	Num      int    `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address  string `gorm:"index:addr"`     // 给address字段创建名为addr的索引
	IgnoreMe int    `gorm:"-"`              // 忽略本字段
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

// Migrate 自动创建表结构
func Migrate() {
	db.AutoMigrate(&User{})
}

// Create 创建
func Create() {
	u := User{
		Name: "张三",
		Age: sql.NullInt64{
			Int64: 18,
		},
		Email:   "hell@qq.com",
		Address: "四川省成都市",
	}

	result := db.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(result.RowsAffected)
}

// Get 获取一个
func Get() {
	u := User{}
	result := db.First(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(u)
}
