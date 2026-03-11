package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
