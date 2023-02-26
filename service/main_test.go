package service

import (
	"blog/server/initialize"
	"blog/server/initialize/global"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	path := "../"
	initialize.Config(path)
	// initialize.Others()
	global.GLOBAL_LOG = initialize.Logger()
	global.GLOBAL_DB = initialize.DataBase()
	if global.GLOBAL_DB != nil {
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}

	os.Exit(m.Run())
}
