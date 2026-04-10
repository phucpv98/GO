package repo

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/database"
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

type userRepository struct {
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserByEmail implements [IUserRepository].
func (up *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&models.GoCrmUser{}).RowsAffected
	// return row != NumberNull

	// user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	// if err != nil {
	// 	return false
	// }
	// fmt.Println("Get User By Email: ", user)
	// return user.UsrID != 0

	return false
}
