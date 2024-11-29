package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model

	RoleId       uint `json:"roleId"`
	PermissionId uint `json:"permissionId"`

	Role       *Role       `json:"role" gorm:"foreignKey:RoleId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permission *Permission `json:"permission" gorm:"foreignKey:PermissionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
