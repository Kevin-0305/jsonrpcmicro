package main

import (
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/logs/config"
	"jsonrpcmicro/internal/logs/initialize"
	"jsonrpcmicro/internal/logs/svc"
	"jsonrpcmicro/utils"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/google/uuid"
)

func main() {
	config.Conf = config.Init()
	global.DB = initialize.GormMysql()
	global.REDIS = initialize.Redis()

	if global.DB != nil {
		initialize.InitMysql(global.DB)
	}

	//jsonrpc 服务注册
	rpc.Register(new(svc.LogSystemService))
	rpc.Register(new(svc.LogService))

	ListenOn := config.Conf.RpcServerConf.ListenOn
	sock, err := net.Listen("tcp", ":"+ListenOn)
	log.Println("listen at :" + ListenOn)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	//etcd 服务注册
	serviceId := uuid.New().String() // 服务ID
	service, err := utils.NewService(config.Conf.Etcd.Hosts)
	if err != nil {
		log.Fatal("etcd connect error:", err)
	}

	go func() {
		err := service.RegService(serviceId, config.Conf.Etcd.Key, config.Conf.RpcServerConf.ServiceAddress+":"+config.Conf.RpcServerConf.ListenOn)
		if err != nil {
			log.Fatal("etcd connect error:", err)
		}
	}()

	for {
		conn, err := sock.Accept()
		if err != nil {
			continue
		}
		// 5.
		go jsonrpc.ServeConn(conn)
	}

}
