package model

import (
	"jsonrpcmicro/global"

	"github.com/gofrs/uuid"
)

type AuthUser struct {
	global.BaseModel
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                // 用户UUID
	Username    string    `json:"userName" gorm:"comment:用户登录名"`             // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                  // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称"
	AuthGroup   AuthGroup `json:"authGroup" gorm:"foreignKey:AuthGroupId;comment:用户分组"`
	AuthGroupId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}
