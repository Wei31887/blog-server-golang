package global

import (
	"blog/server/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

const DateFormat = "2006-01-02 15:04:05"

var (
	GLOBAL_CONFIG	*config.Config
	GLOBAL_LOG		*zap.Logger
	GLOBAL_DB		*gorm.DB
	GLOBAL_REDIS   	*redis.Client 
)