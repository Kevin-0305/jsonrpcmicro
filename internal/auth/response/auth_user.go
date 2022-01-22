package response

import "github.com/gofrs/uuid"

type LoginResponse struct {
	SessionID string `json:"sessionId"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Info      string `json:"info"`
}

type AuthUserResponse struct {
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                // 用户UUID
	Account   string    `json:"account" gorm:"comment:用户登录名"`              // 用户登录名
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称"
	AuthGroup int       `json:"authGroup" gorm:"foreignKey:AuthGroupId;comment:用户分组"`
}
