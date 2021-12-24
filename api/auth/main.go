package main

import (
	"fmt"
	"jsonrpcmicro/api/auth/config"
	"jsonrpcmicro/internal/auth/svc"
	"jsonrpcmicro/utils"
	"log"
	"net/rpc/jsonrpc"
)

// 客户端调用 jsonrpc 有两个步骤
// 1. 使用 jsonrpc.Dial() 方法连接到服务，并返回一个连接 conn
// 2. 调用 conn.Call() 方法调用服务

// 定义 MathService 所需要的参数，一般是两个，int 类型
type Args struct {
	Arg1, Arg2 int
}

type Resp struct {
	Num int `json:"num"`
}

func main() {
	config := new(config.Config)
	conf := config.Init()
	client, err := utils.NewClient(conf.Etcd.Hosts)
	if err != nil {
		log.Fatalf("etcd连接错误:", err.Error())
	}
	err = client.LoadService()
	if err != nil {
		log.Fatalf("etcd service error", ":no services find")
	}
	serviceInfo := client.GetService(conf.Etcd.Key)
	fmt.Println(serviceInfo)
	// 1.
	conn, err := jsonrpc.Dial("tcp", serviceInfo.ServiceAddress)
	if err != nil {
		log.Fatal("can't not connect to")
	}
	var reply svc.UserResponse
	args := svc.LoginRequest{
		Account:  "Admin",
		Password: "123456"}

	// 调用 Add() 方法
	err = conn.Call("UserService.Login", args, &reply)
	if err != nil {
		log.Fatal("call UserService.Login error:", err)
	}
	fmt.Printf("UserServicer.Login(%s,%s)\n", args.Account, reply.Name)
}
