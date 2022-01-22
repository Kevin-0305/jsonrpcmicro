package svc

import (
	"errors"
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/auth/model"
	"jsonrpcmicro/internal/auth/request"
	"jsonrpcmicro/internal/auth/response"
	"jsonrpcmicro/utils"
	"log"

	"gorm.io/gorm"
)

type UserService struct {
}

func (that *UserService) Login(loginRequest request.LoginRequest, reply *response.LoginResponse) error {
	//global.Log.Error("test")
	var user model.AuthUser
	err := global.DB.Where("account =?", loginRequest.Account).First(&user).Error
	if err != nil {
		log.Println("账号不存在", err.Error())
		reply.Status = 1
		reply.Info = "账号不存在"
		return nil
	}
	if loginRequest.Password != user.Password {
		reply.Status = 1
		reply.Info = "密码错误"
		return nil
	}
	var auths []model.CasbinInfo
	auth := model.CasbinInfo{
		Path:   "/auth/Login",
		Method: "POST",
	}
	auths = append(auths, auth)
	sessionID := AddSession(user)
	reply.SessionID = sessionID
	reply.Name = user.NickName
	reply.Status = 0
	reply.Info = "登录成功"
	return nil
}

func CreateUser(createUser request.AuthUserRequest, reply *response.AuthUserResponse) {
	var user model.AuthUser
	if !errors.Is(global.DB.Where("account = ?", createUser.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册

	}
	user.Password = utils.MD5V([]byte(createUser.Password))
}
