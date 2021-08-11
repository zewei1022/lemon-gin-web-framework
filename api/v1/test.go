package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/lib/mongodb"
	"github.com/zewei1022/lemon-gin-web-framework/lib/redis"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func TestRedisLib(c *gin.Context)  {
	res, err := redis.ZRank("rank", "a2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func TestMongodbLib(c *gin.Context)  {
	collection := mongodb.Cli.Database("testing").Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _ := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID
	fmt.Println("id：", id)
}