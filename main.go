package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"net/http"
	"time"
)

func main() {
	InitConfigByViper()
	RunServer()
}

func InitConfigByViper() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	var config *config.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err.Error())
	}

	global.LGWF_SERVER_CONFIG = config.Server
}

func RunServer()  {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", global.LGWF_SERVER_CONFIG.Addr),
		Handler: r,
		ReadTimeout: time.Duration(global.LGWF_SERVER_CONFIG.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(global.LGWF_SERVER_CONFIG.WriteTimeout) * time.Second,
		MaxHeaderBytes: global.LGWF_SERVER_CONFIG.MaxHeaderBytes,
	}

	_ = s.ListenAndServe()
}
