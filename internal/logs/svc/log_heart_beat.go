package svc

import (
	"jsonrpcmicro/internal/logs/request"
	"jsonrpcmicro/internal/logs/response"
)

type LogService struct {
}

func (that *LogService) HeartBeat(request *request.Request, reply *response.Response) error {
	reply.State = 0
	reply.Message = "OK"
	return nil
}
