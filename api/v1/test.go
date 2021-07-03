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

	renameRes, _ := redis.Rename("test1", "test2")
	fmt.Println("rename result：", renameRes)

	renameNxRes, _ := redis.RenameNx("test2", "test2")
	fmt.Println("renamenx result：", renameNxRes)

	keysResult, _ := redis.Keys("test*")
	fmt.Printf("keys result: %v", keysResult)
}