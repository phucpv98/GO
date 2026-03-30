package repo

import (
	"fmt"
	"go-ecommerce/global"
	"time"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type authRepository struct{}

// AddOTP implements [IAuthRepository].
func (a *authRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("otp:%s", email)
	return global.Rdb.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IAuthRepository {
	return &authRepository{}
}
