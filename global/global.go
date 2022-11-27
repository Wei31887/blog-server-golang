package global

import (
	"blog/server/config"

	"go.uber.org/zap"
)

var (
	GLOBAL_CONFIG	*config.Config
	GLOBAL_LOG		*zap.Logger
)