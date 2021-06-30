package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/model/request"
	"github.com/zewei1022/lemon-gin-web-framework/model/response"
	"github.com/zewei1022/lemon-gin-web-framework/service"
)

func FindBook(c *gin.Context) {
	var idInfo request.IdInfo
	_ = c.ShouldBind(&idInfo)

	if book, err := service.FindBookById(idInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.BookInfo{
			ID:        book.ID,
			Name:      book.Name,
			Author:    book.Author,
			Price:     book.Price,
			Page:      book.Page,
			Isbn:      book.Isbn,
			Publisher: book.Publisher,
		}, "获取成功", c)
	}
}

func GetBookList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBind(&pageInfo)

	global.LGWF_LOGGER.Infof("GetBookList api, request info: %v", pageInfo)

	conn := global.LGWF_REDIS.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "test_key", "lemon")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	if list, total, err := service.GetBookList(pageInfo); err != nil {
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
			return
		}
	}

	if err := service.CreateBook(book); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func UpdateBook(c *gin.Context) {
	var book request.UpdateBookInfo
	_ = c.ShouldBind(&book)

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	if err := service.UpdateBook(book); err != nil {
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func DeleteBook(c *gin.Context) {
	var idInfo request.IdInfo
	_ = c.ShouldBind(&idInfo)

	if err := service.DeleteBook(idInfo.ID); err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
