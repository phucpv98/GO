package models

// GoCrmUser represents the user model mapped to the go_db_user table.
type GoCrmUser struct {
	UsrEmail string `gorm:"column:usr_email"`
}
