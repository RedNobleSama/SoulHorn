/**
* @Author: oreki
* @Date: 2021/6/6 22:30
* @Email: a912550157@gmail.com
 */

package model

import (
	"SoulHorn/utils/errmsg"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	//GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required, min=4, max=12" label:"用户名"`
	Password string `gorm:"type:varchar(50);not null" json:"password" validate:"required, min=6, max=20" label:"密码"`
	Role     int    `gorm:"type:int;Default: 2" json:"role" validate:"required, gte=2" label:"角色"`
	Email    string `gorm:"type:varchar(20);not null" json:"email" validate:"required, email" lable:"邮箱"`
}

// CheckUser 查询用户是否存在
func (u User)CheckUser(args ...string) int {
	db.Select("id").Where("username = ?", args[0]).First(&u)
	if u.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreatUser 新增用户
func (u User) CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

