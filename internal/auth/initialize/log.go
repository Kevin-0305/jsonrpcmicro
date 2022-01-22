package initialize

import (
	"fmt"
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/auth/config"
	"jsonrpcmicro/internal/auth/request"
	"jsonrpcmicro/internal/auth/response"
	"jsonrpcmicro/middleware"
	"jsonrpcmicro/utils"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func InitLog(conf *config.Config) {
	conn := ConnectLogService(conf)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			var args request.Request
			var reply *response.Response
			err := conn.Call("LogService.HeartBeat", args, &reply)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func ConnectLogService(conf *config.Config) *rpc.Client {
	client, err := utils.NewClient(conf.Log.Hosts)
	err = client.LoadService()
	if err != nil {
		log.Println("etcd service error", ":no services find")
	}
	serviceInfo := client.GetService(conf.Log.Key)
	conn, err := jsonrpc.Dial("tcp", serviceInfo.ServiceAddress)
	if err != nil {
		log.Println("can't not connect to")
	}
	global.Log = middleware.Log(conn)
	return conn
}
