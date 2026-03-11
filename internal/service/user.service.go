package service

import (
	"go-ecommerce/internal/repo"
	"go-ecommerce/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetUserByEmail implements [IUserService].
func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExisted
	}

	return response.ErrorCodeSuccess
}
