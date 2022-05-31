package database

import (
	"context"
	"time"

	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/gorm-study/pkg/models"
	"gorm.io/gorm"
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
	if err := gormDB.Create(u).Error; err != nil {
		logrus.WithContext(ctx).Errorf("creare user failure:[params:%+v,err:%s]", params, err.Error())
		return nil, err
	}
	return u, nil
}

// Get 查询
func Get(ctx context.Context, uid uint) (*User, error) {
	var (
		user = &User{}
	)
	if err := gormDB.First(user, uid).Error; err != nil {
		logrus.WithContext(ctx).Errorf("get failure:[id:%d,err:%s]", uid, err.Error())
		return nil, err
	}
	return user, nil
}

// Delete 删除
func Delete(ctx context.Context, uid uint) (int64, error) {
	result := gormDB.Delete(&User{}, uid)
	if err := result.Error; err != nil {
		logrus.WithContext(ctx).Errorf("delete failure:[err:%s]", err.Error())
		return 0, err
	}
	return result.RowsAffected, nil
}
