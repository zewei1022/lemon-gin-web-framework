package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	Router := gin.Default()

	gin.SetMode("debug")

	apiV1Group := Router.Group("/api/v1/")
	{
		InitBookRouter(apiV1Group)
	}

	return Router
}