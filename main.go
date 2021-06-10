package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/router"
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
	r := router.Init()

	address := fmt.Sprintf(":%d", global.LGWF_SERVER_CONFIG.Addr)
	s := &http.Server{
		Addr: address,
		Handler: r,
		ReadTimeout: time.Duration(global.LGWF_SERVER_CONFIG.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(global.LGWF_SERVER_CONFIG.WriteTimeout) * time.Second,
		MaxHeaderBytes: global.LGWF_SERVER_CONFIG.MaxHeaderBytes,
	}

	fmt.Printf("Listening and serving HTTP on %s\n", address)

	_ = s.ListenAndServe()
}
