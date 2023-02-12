package main

import (
	G "blog/server/global"
	"blog/server/initialize"
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
	initialize.Config(false)
	initialize.Others()
	G.GLOBAL_LOG = initialize.Logger()
	G.GLOBAL_DB = initialize.DataBase()
	if G.GLOBAL_DB != nil {
		db, _ := G.GLOBAL_DB.DB()
		defer db.Close()
	}

	gin.SetMode(G.GLOBAL_CONFIG.Server.Model)
	router := router.InitRouter()
	server := &http.Server{
		Addr:    G.GLOBAL_CONFIG.Server.Address,
		Handler: router,
	}

	// listen
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			G.GLOBAL_LOG.Fatal("listen on: ", zap.String("address", err.Error()))
		}
		G.GLOBAL_LOG.Info("listen on ", zap.String("address", G.GLOBAL_CONFIG.Server.Address))
	}()

	quit := make(chan os.Signal)
	// listen
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
