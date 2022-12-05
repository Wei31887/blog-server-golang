package initialize

import (
	"blog/server/cache"
	"blog/server/global"
	"time"
)


func Others() {
	// cahce
	global.GLOBAL_CAHCHE = cache.New(
		time.Duration(global.GLOBAL_CONFIG.Cache.Expire) * time.Second, 
		time.Duration(global.GLOBAL_CONFIG.Cache.Clearup) * time.Second,
	)
}