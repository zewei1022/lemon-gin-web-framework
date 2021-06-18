package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zewei1022/lemon-gin-web-framework/api/v1"
)

func InitBookRouter(Router *gin.RouterGroup)  {
	BookRouter := Router.Group("book")
	{
		BookRouter.GET("/findBook", v1.FindBook)
		BookRouter.GET("/getBookList", v1.GetBookList)
		BookRouter.POST("/createBook", v1.CreateBook)
		BookRouter.POST("/updateBook", v1.UpdateBook)
		BookRouter.POST("/deleteBook", v1.DeleteBook)
	}
}