package models

type CreateUser struct {
	// 用户名
	UserName string `json:"userName"`
	// 邮箱
	Email string `json:"email"`
	// 地址
	Address string `json:"address"`
}
