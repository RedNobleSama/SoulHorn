/**
* @Author: oreki
* @Date: 2021/6/7 23:23
* @Email: a912550157@gmail.com
 */

package v1

import (
	"SoulHorn/model"
	"SoulHorn/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var book model.DouBanBook

func GetBooks(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code := book.GetBooks(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
