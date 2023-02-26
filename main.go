package main

import (
	"blog/server/initialize"
	"blog/server/initialize/global"
	"blog/server/router"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	/*  initialize  */
	err := initialize.Config(".")
	if err != nil {
		log.Fatal("Cannot initialize config file", err)
	}
	global.GLOBAL_LOG = initialize.Logger()
	global.GLOBAL_REDIS = initialize.Redis()
	global.GLOBAL_DB = initialize.DataBase()
	if global.GLOBAL_DB != nil {
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}

	gin.SetMode(global.GLOBAL_CONFIG.Server.Model)
	router := router.InitRouter()
	server := &http.Server{
		Addr:    global.GLOBAL_CONFIG.Server.Address,
		Handler: router,
	}

	// listen
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.GLOBAL_LOG.Fatal("listen on: ", zap.String("address", err.Error()))
		}
		global.GLOBAL_LOG.Info("listen on ", zap.String("address", global.GLOBAL_CONFIG.Server.Address))
	}()

	quit := make(chan os.Signal)
	// listen
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
