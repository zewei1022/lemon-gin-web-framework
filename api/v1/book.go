package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zewei1022/lemon-gin-web-framework/model/request"
	"github.com/zewei1022/lemon-gin-web-framework/model/response"
	"github.com/zewei1022/lemon-gin-web-framework/service"
)

func GetBookList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBind(&pageInfo)

	if err, list, total := service.GetBookList(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func CreateBook(c *gin.Context) {
 	var book request.CreateBookInfo
	_ = c.ShouldBind(&book)

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.FailWithMessage(err.Error(), c)
		}
	}

	if err := service.CreateBook(book); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
