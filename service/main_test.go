package service

import (
	G "blog/server/global"
	"blog/server/initialize"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	path := "../"
	initialize.Config(path)
	// initialize.Others()
	G.GLOBAL_LOG = initialize.Logger()
	G.GLOBAL_DB = initialize.DataBase()
	if G.GLOBAL_DB != nil {
		db, _ := G.GLOBAL_DB.DB()
		defer db.Close()
	}

	os.Exit(m.Run())
}