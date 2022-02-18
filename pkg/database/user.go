package database

import (
	"context"
	"github.com/zjzjzjzj1874/gorm-study/pkg/models"
	"time"

	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
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

// Create 新增
func Create(ctx context.Context, params *models.CreateUser) (*User, error) {
	u := &User{
		Name:    params.UserName,
		Email:   params.Email,
		Address: params.Address,
	}
	if err := db.Create(u).Error; err != nil {
		logrus.WithContext(ctx).Errorf("creare user failure:[params:%+v,err:%s]", params, err.Error())
		return nil, err
	}
	return u, nil
}

// First 查询
func First(ctx context.Context, user *User) (*User, error) {
	if err := db.First(user).Error; err != nil {
		logrus.WithContext(ctx).Errorf("first failure:[err:%s]", err.Error())
		return nil, err
	}
	return user, nil
}

// Delete 删除
func Delete(ctx context.Context, user *User) (int64, error) {
	result := db.Delete(user)
	if err := result.Error; err != nil {
		logrus.WithContext(ctx).Errorf("delete failure:[err:%s]", err.Error())
		return 0, err
	}
	return result.RowsAffected, nil
}
