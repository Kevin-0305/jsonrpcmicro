package auth

import (
	"fmt"
	"jsonrpcmicro/api/response"
	"jsonrpcmicro/global"
	authRequest "jsonrpcmicro/internal/auth/request"
	authResponse "jsonrpcmicro/internal/auth/response"
	"jsonrpcmicro/utils"
	"log"
	"net/rpc/jsonrpc"

	"github.com/gin-gonic/gin"
)

// 客户端调用 jsonrpc 有两个步骤
// 1. 使用 jsonrpc.Dial() 方法连接到服务，并返回一个连接 conn
// 2. 调用 conn.Call() 方法调用服务

// 定义 MathService 所需要的参数，一般是两个，int 类型

func Login(c *gin.Context) {

	// config := new(config.Config)
	// conf := config.Init()
	etcd := global.ApiConfig.FindEtcdSvc("Auth")
	client, err := utils.NewClient(etcd.Hosts)
	if err != nil {
		log.Fatalf("etcd连接错误:", err.Error())
	}
	err = client.LoadService()
	if err != nil {
		log.Fatalf("etcd service error", ":no services find")
	}
	serviceInfo := client.GetService(etcd.Key)

	// 1.
	conn, err := jsonrpc.Dial("tcp", serviceInfo.ServiceAddress)
	if err != nil {
		log.Fatal("can't not connect to")
	}
	var args authRequest.LoginRequest
	var reply authResponse.LoginResponse
	_ = c.ShouldBindJSON(&args)
	fmt.Println(args)
	// 调用 Add() 方法
	err = conn.Call("UserService.Login", args, &reply)
	if err != nil {
		log.Fatal("call UserService.Login error:", err)
	}
	fmt.Printf("UserServicer.Login(%d,%s)\n", reply.SessionID, reply.Name)

	response.OkWithData(reply, c)
}
