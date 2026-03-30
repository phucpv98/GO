package repo

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/models"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "TipJS"
// }

// INTERFACE_VERSION
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserByEmail implements [IUserRepository].
func (u *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&models.GoCrmUser{}).RowsAffected
	return row != NumberNull
}
