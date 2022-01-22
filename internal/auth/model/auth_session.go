package model

import (
	"time"
)

type Session struct {
	SessionID  string       `json:"sessionId"`
	UserID     uint         `json:"userId"`
	AuthGroup  uint         `json:"authGroup"`
	NickName   string       `json:"nickName"`
	CreateTime time.Time    `json:"createTime"`
	UpdateTime time.Time    `json:"updateTime"`
	Expires    time.Time    `json:"expires"`
	Authoritys []CasbinInfo `json:"authoritys"`
}

// type Authority struct {
// 	Path   string `json:"path"`
// 	Method string `json:"method"`
// }
