package model

import "jsonrpcmicro/global"

type Authority struct {
	global.BaseModel
	AuthApi     AuthApi `json:"authApi" gorm:"foreignKey:AuthApiId;comment:apiID"`
	AuthApiId   int
	AuthGroup   AuthGroup `json:"authGroup" gorm:"foreignKey:AuthGroupId;comment:分组ID"`
	AuthGroupId int
}
