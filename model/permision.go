package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Api         string `json:"api" gorm:"unique"`
	Description string `json:"description"`
	Action      bool   `json:"action"`

	RolePermissions []RolePermission `json:"rolePermissions" gorm:"foreignKey:PermissionId;"`
}
