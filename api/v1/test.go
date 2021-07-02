package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/lib/redis"
)

func TestRedisLib(c *gin.Context)  {
	key := "test_key"

	redis.Set(key, "lemon", 3600)
	value, err := redis.Get(key)
	if err != nil {
		fmt.Println("err：", err)
	}
	fmt.Println("value：", string(value))

	exists := redis.Exists(key)
	fmt.Println("exist：", exists)
}