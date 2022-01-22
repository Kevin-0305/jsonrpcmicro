package main

import (
	"jsonrpcmicro/api/config"
	"jsonrpcmicro/api/initialize"
	"jsonrpcmicro/global"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	Router := initialize.Routers()
	global.REDIS = initialize.Redis("127.0.0.1:6379")
	global.ApiConfig = config.Init()
	address := "0.0.0.0:12000"
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	log.Fatal("server run success on ", zap.String("address", address))
	log.Fatal(s.ListenAndServe())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	s.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

	return s
}

type server interface {
	ListenAndServe() error
}
