package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zewei1022/lemon-gin-web-framework/api/v1"
)

func InitTestRouter(Router *gin.RouterGroup)  {
	BookRouter := Router.Group("test")
	{
		BookRouter.GET("/redis", v1.TestRedisLib)
	}
}
