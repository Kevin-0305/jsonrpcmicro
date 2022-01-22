package middleware

import (
	"fmt"
	"jsonrpcmicro/internal/logs/request"
	"jsonrpcmicro/internal/logs/response"
	"log"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

type LogHook struct {
	Conn *rpc.Client
}

func (h *LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.PanicLevel,
	}
}

func (h *LogHook) Fire(entry *logrus.Entry) error {
	var args request.LogSystemRequest
	var reply response.Response
	args.File = entry.Caller.File
	args.Message = entry.Message
	args.Level = entry.Level.String()
	args.Line = entry.Caller.Line
	args.OccurTime = entry.Time
	// 调用 Add() 方法
	err := h.Conn.Call("LogSystemService.Create", args, &reply)
	if err != nil {
		log.Fatal("call UserService.Login error:", err)
	}
	log.Println("")
	fmt.Printf("Log send(%d,%s)\n", reply.State, reply.Message)

	return nil
}

func Log(conn *rpc.Client) *logrus.Logger {
	log := logrus.New()
	log.AddHook(&LogHook{Conn: conn})
	log.SetLevel(logrus.TraceLevel)
	log.SetReportCaller(true)
	return log
}

var LogRpcConn *rpc.Client
