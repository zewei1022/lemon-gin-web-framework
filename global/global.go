package global

import (
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"gorm.io/gorm"
)

var (
	LGWF_DB              *gorm.DB
	LGWF_SERVER_CONFIG   config.Server
	LGWF_DATABASE_CONFIG config.Database
	LGWF_JWT_CONFIG      config.JWT
)
