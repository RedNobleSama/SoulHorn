/**
* @Author: oreki
* @Date: 2021/6/7 0:04
* @Email: a912550157@gmail.com
 */

package v1

import (
	"SoulHorn/model"
	"SoulHorn/utils/errmsg"
	"SoulHorn/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	user model.User
	msg string
	code int
)



func AddUser(c *gin.Context) {
	_ = c.ShouldBind(&user)
	msg, code = validator.Validate(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"message": msg,
		})
		c.Abort()
	}

	code = user.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		user.CreateUser(&user)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data": user,
		"message": errmsg.GetErrMsg(code),
	})
}