package main

import (
	"blog/server/global"
	"blog/server/initialize"
	"net/http"

	"go.uber.org/zap"
)

func main(){
	/*  initialize  */
	// initialize config
	initialize.InitializeConfig() 
	// initialize logger
	global.GLOBAL_LOG = initialize.InitializeLogger()
	// initialize databse
	global.GLOBAL_DB = initialize.InitializeDataBase()
	if global.GLOBAL_DB != nil {
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}

	// simpleHttpGet("www.google.com")
	// simpleHttpGet("http://www.google.com")
}


// test for logger
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		global.GLOBAL_LOG.Error(
			"Error fetching URL ...", 
			zap.String("url", url), 
			zap.Error(err))
	} else {
		global.GLOBAL_LOG.Info(
			"Success! ", 
			zap.String("code", resp.Status), 
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}

