package model

import (
	"time"

	"gorm.io/gorm"
)

type AuthGroup struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`              // 删除时间
	GroupName string         `json:"groupName" gorm:"comment:组名"` // 用户登录名
}
