package service

import (
	"fmt"
	"go-ecommerce/internal/repo"
	"go-ecommerce/internal/utils/crypto"
	"go-ecommerce/internal/utils/random"
	"go-ecommerce/internal/utils/sendto"
	"go-ecommerce/response"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	authRepo repo.IAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, authRepo repo.IAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

// GetUserByEmail implements [IUserService].
func (us *userService) Register(email string, purpose string) int {
	// 0. Hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email is :::%s\n", hashEmail)

	// 1. check email exists
	if us.userRepo.GetUserByEmail(hashEmail) {
		return response.ErrorCodeUserHasExisted
	}

	// 2. New OTP -> ...
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("Otp is :::%d\n", otp)

	// 3. Save OTP in Redis with expiration time
	err := us.authRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrorInvalidOTP
	}

	// 4. Send Email OTP
	// err = sendto.SendTextEmailOtp([]string{email}, "account@email.com", strconv.Itoa(otp))
	// if err != nil {
	// 	return response.ErrorSendEmailOTP
	// }

	err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, "otp-auth.html")
	if err != nil {
		return response.ErrorSendEmailOTP
	}

	return response.CodeSuccess
}
