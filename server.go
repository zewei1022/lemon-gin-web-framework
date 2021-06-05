package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"net/http"
)

func RunServer()  {
	r := gin.Default()

	address := fmt.Sprintf(":%d", global.LGWF_SERVER_CONFIG.Addr)
	s := &http.Server{
		Addr: address,
		Handler: r,
		ReadTimeout: global.LGWF_SERVER_CONFIG.ReadTimeout,
		WriteTimeout: global.LGWF_SERVER_CONFIG.WriteTimeout,
		MaxHeaderBytes: global.LGWF_SERVER_CONFIG.MaxHeaderBytes,
	}

	_ = s.ListenAndServe()
}