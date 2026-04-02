package implements

import (
	"context"
	"go-ecommerce/internal/database"
)

type UserLogin struct {
	// Implement the IUserLogin interface
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *UserLogin {
	return &UserLogin{
		r: r,
	}
}

func (s *UserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *UserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *UserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *UserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
