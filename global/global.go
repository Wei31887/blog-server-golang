package global

import (
	"blog/server/cache"
	"blog/server/config"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

const DateFormat = "2006-01-02 15:04:05"

var (
	GLOBAL_CONFIG	*config.Config
	GLOBAL_LOG		*zap.Logger
	GLOBAL_DB		*gorm.DB
	GLOBAL_CAHCHE   *cache.Cache 
)