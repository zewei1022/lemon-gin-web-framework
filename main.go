package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/lib/mongodb"
	"github.com/zewei1022/lemon-gin-web-framework/lib/redis"
	"github.com/zewei1022/lemon-gin-web-framework/router"
	"github.com/zewei1022/lemon-gin-web-framework/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	InitConfig()
	InitLogger()
	InitDb()
	InitRedis()
	InitMongodb()
	RunServer()
}

func InitDb() {
	switch global.LGWF_CONFIG.Database.Type {
	case "mysql":
		m := global.LGWF_CONFIG.Database
		dsn := m.Username + ":" + m.Password + "@tcp(" + m.Hostname + ")/" + m.DbName + "?" + m.Config
		fmt.Printf("db dsn is : %s\n", dsn)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
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

func InitRedis()  {
	redis.Initialize(global.LGWF_CONFIG.Redis)
}

func InitMongodb()  {
	mongodb.Initialize(global.LGWF_CONFIG.Mongodb)
}

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	var customConfig *config.Config
	err = viper.Unmarshal(&customConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	global.LGWF_CONFIG.Server = customConfig.Server
	global.LGWF_CONFIG.Database = customConfig.Database
	global.LGWF_CONFIG.JWT = customConfig.JWT
	global.LGWF_CONFIG.Zap = customConfig.Zap
	global.LGWF_CONFIG.Redis = customConfig.Redis
	global.LGWF_CONFIG.Mongodb = customConfig.Mongodb
}

func InitLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	global.LGWF_LOGGER = sugarLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	timeObj := time.Now()
	logDirector := global.LGWF_CONFIG.Zap.Director
	fullLogFileName := fmt.Sprintf("%s/%d-%d-%d.log", logDirector, timeObj.Year(), timeObj.Month(), timeObj.Day())

	if ok, _ := utils.PathExists(logDirector); !ok {
		_ = utils.CreateDir(logDirector)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fullLogFileName,
		MaxSize:    global.LGWF_CONFIG.Zap.MaxSize,
		MaxBackups: global.LGWF_CONFIG.Zap.MaxBackups,
		MaxAge:     global.LGWF_CONFIG.Zap.MaxAge,
		Compress:   global.LGWF_CONFIG.Zap.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func RunServer() {
	r := router.Init()

	address := global.LGWF_CONFIG.Server.Addr
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    time.Duration(global.LGWF_CONFIG.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.LGWF_CONFIG.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: global.LGWF_CONFIG.Server.MaxHeaderBytes,
	}

	fmt.Printf("Listening and serving HTTP on %s\n", address)

	global.LGWF_LOGGER.Error(s.ListenAndServe().Error())
}
