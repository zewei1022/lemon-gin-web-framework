package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/lib/redis"
)

func TestRedisLib(c *gin.Context)  {
	key := "test_key"

	redis.Set(key, "lemon", 3600)
	value, _ := redis.Get(key)

	fmt.Println("value：", string(value))

	exists, _ := redis.Exists(key)
	fmt.Println("exist：", exists)

	redis.Expire(key, 1800)

	second, _ := redis.Ttl(key)
	fmt.Println("ttl：", second)

	delCount, _ := redis.Del("del1", "del2", "del3")
	fmt.Println("delete count：", delCount)
}