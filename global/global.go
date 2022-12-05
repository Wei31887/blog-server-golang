package global

import (
	"blog/server/cache"
	"blog/server/config"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	GLOBAL_CONFIG	*config.Config
	GLOBAL_LOG		*zap.Logger
	GLOBAL_DB		*gorm.DB
	GLOBAL_CAHCHE   *cache.Cache 
)