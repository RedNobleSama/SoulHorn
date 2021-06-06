/**
* @Author: oreki
* @Date: 2021/6/5 11:08
* @Email: a912550157@gmail.com
 */

package routes

import (
	v1 "SoulHorn/api/v1"
	"SoulHorn/utils"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//无需登录调用的接口
	NoAuthRouter := router.Group("api/v1")
	{
		NoAuthRouter.POST("user/add", v1.AddUser)
	}

	return router
}
