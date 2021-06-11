package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	InitConfig()
	InitDb()
	RunServer()
}

func InitDb()  {
	switch global.LGWF_DATABASE_CONFIG.Type {
	case "mysql":
		m := global.LGWF_DATABASE_CONFIG
		dsn := m.Username + ":" + m.Password + "@tcp(" + m.Hostname + ")/" + m.DbName + "?" + m.Config
		fmt.Printf("db dsn is : %s\n", dsn)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn, // DSN data source name
			DefaultStringSize: 256, // string 类型字段的默认长度
			DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		global.LGWF_DB = db
	}
}

func InitConfig() {
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
	global.LGWF_DATABASE_CONFIG = config.Database
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
