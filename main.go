package main

import (
	"blog/server/gapi"
	"blog/server/initialize"
	"blog/server/initialize/config"
	"blog/server/initialize/global"
	"blog/server/pb"
	"blog/server/server"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Running server
	// go runHttpServer(global.GLOBAL_CONFIG)
	runGrpcServer(global.GLOBAL_CONFIG)


	quit := make(chan os.Signal)
	// listen
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}

func runHttpServer(config *config.Config) {
	server := server.NewHTTPServer(config)
	err := server.RunServer(global.GLOBAL_CONFIG.Server.HTTPAddress)
		if err != nil && err != http.ErrServerClosed {
			global.GLOBAL_LOG.Fatal("listen on: ", zap.String("address", err.Error()))
		}
	global.GLOBAL_LOG.Info("listen on ", zap.String("address", global.GLOBAL_CONFIG.Server.HTTPAddress))
}

func runGrpcServer(config *config.Config) {
	server := gapi.NewGRPCServer(config)
	grpcServer := grpc.NewServer()
	pb.RegisterBlogServerServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.Server.GrpcAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %s", err)
	}
}