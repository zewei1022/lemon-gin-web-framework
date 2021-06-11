package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func JsonResponse(code int, data interface{}, msg string, c *gin.Context)  {
	c.JSON(http.StatusOK, Response{code, data, msg})
}

func Ok(c *gin.Context)  {
	JsonResponse(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	JsonResponse(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	JsonResponse(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	JsonResponse(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	JsonResponse(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	JsonResponse(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	JsonResponse(ERROR, data, message, c)
}

