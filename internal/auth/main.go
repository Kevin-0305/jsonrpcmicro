// package main

// import (
// 	"log"
// 	"net"
// 	"net/rpc"
// 	"net/rpc/jsonrpc"
// )

// // Golang RPC 的实现需要 5 个步骤
// // 1. 定义一个服务结构
// // 2. 为这个服务结构定义几个服务方法，每个方法接受两个参数和返回 error 类型
// // 3. 使用 rpc.Register() 方法注册 「服务结构」 的实例
// // 4. 监听套接字
// // 5. 为每一个套接字调用 jsonrpc.ServerConn(conn) 方法

// // 1
// // 定义一个服务，因为提供的是数学服务，所以就叫 MathService
// type MathService struct {
// }

// // 定义 MathService 所需要的参数，一般是两个，int 类型
// type Args struct {
// 	Arg1, Arg2 int
// }

// type Resp struct {
// 	Num int `json:"num"`
// }

// // 2.
// // 实现加法服务，加法需要两个参数
// // 所有的 jsonrpc 方法只有两个参数，第一个参数用于接收所有参数，
// // 第二个参数用于处理返回结果，是一个指针
// // 所有的 jsonrpc 都只有一个返回值，error,用于指示是否发生错误
// func (that *MathService) Add(args Args, reply *Resp) error {
// 	reply.Num = args.Arg1 + args.Arg2
// 	return nil
// }

// // 实现减法服务
// func (that *MathService) Sub(args Args, reply *Resp) error {
// 	reply.Num = args.Arg1 - args.Arg2
// 	return nil
// }

// // 实现乘法服务
// func (that *MathService) Mul(args Args, reply *Resp) error {
// 	reply.Num = args.Arg1 * args.Arg2
// 	return nil
// }

// // 实现除法服务
// func (that *MathService) Div(args Args, reply *Resp) error {
// 	reply.Num = args.Arg1 / args.Arg2
// 	return nil
// }

// func main() {
// 	// 3.
// 	rpc.Register(new(MathService))
// 	// 4.
// 	sock, err := net.Listen("tcp", ":8080")
// 	log.Println("listen at :8080")
// 	if err != nil {
// 		log.Fatal("listen error:", err)
// 	}

// 	for {
// 		conn, err := sock.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		// 5.
// 		go jsonrpc.ServeConn(conn)
// 	}

// }
package main

import (
	"fmt"
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/auth/config"
	"jsonrpcmicro/internal/auth/initialize"
	"jsonrpcmicro/internal/auth/svc"
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
	initialize.InitLog(&config.Conf)
	global.Log.Error("test")
	if global.DB != nil {
		initialize.InitMysql(global.DB)
	}
	auth := svc.CheckAuthority()
	if auth {
		fmt.Println("True")
	}

	//jsonrpc 服务注册
	rpc.Register(new(svc.UserService))

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
