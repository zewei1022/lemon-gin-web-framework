package global

import (
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	LGWF_DB              *gorm.DB
	LGWF_LOGGER          *zap.SugaredLogger
	LGWF_CONFIG          config.Config
)
