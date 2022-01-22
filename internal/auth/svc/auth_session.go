package svc

import (
	"encoding/json"
	"fmt"
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/auth/model"
	"jsonrpcmicro/utils"
	"time"
)

func AddSession(user model.AuthUser) string {
	parseDuration, _ := time.ParseDuration("24h")
	sessionID := utils.RandString(16)
	authoritys := GetPolicyPathByAuthorityId(string(user.ID))
	fmt.Println(authoritys)
	session := model.Session{
		SessionID:  sessionID,
		UserID:     user.ID,
		AuthGroup:  user.AuthGroupId,
		NickName:   user.NickName,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Expires:    time.Now().Add(parseDuration),
		Authoritys: authoritys,
	}
	value, err := json.Marshal(session)
	if err != nil {
		fmt.Println("json marshal err,%s", err.Error())
	}
	global.REDIS.Set(session.SessionID, value, parseDuration)
	return sessionID
}

func RefreshSession() {

}

// func ParseSession(sessionID string) model.AuthUser {
// 	return nil
// }
