package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/lib/redis"
)

func TestRedisLib(c *gin.Context)  {
	res, err := redis.ZRange("rank", 0, 10, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(res)
}