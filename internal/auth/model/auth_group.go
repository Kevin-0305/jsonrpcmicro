package model

import (
	"jsonrpcmicro/global"
)

type AuthGroup struct {
	global.BaseModel
	GroupName string    `json:"groupName" gorm:"comment:组名"` // 用户登录名
	Authority Authority `json:"groupName" gorm:"comment:组名"`
}
