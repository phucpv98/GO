package implements

import (
	"context"
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/internal/database"
	"go-ecommerce/internal/models"
	"go-ecommerce/internal/utils"
	"go-ecommerce/internal/utils/crypto"
	"go-ecommerce/internal/utils/random"
	"go-ecommerce/response"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
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

func (s *UserLogin) Register(ctx context.Context, in *models.RegisterInput) (codeResult int, err error) {
	// 1. Hash Email
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType: %s\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// 2. Check email exists
	userFound, err := s.r.CheckUserBaseExist(ctx, hashKey)
	if err != nil {
		return response.ErrorCodeUserHasExisted, err
	}
	if userFound > 0 {
		return response.ErrorCodeUserHasExisted, fmt.Errorf("user has already registered")
	}

	// 3. Check userKey exists
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed:: ", err)
		return response.ErrorInvalidOTP, err
	case otpFound != "":
		return response.ErrorCodeOtpNotExists, fmt.Errorf("")
	}

	// 4. Generate OTP (when userKey is not exists)
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("Otp is :::%d\n", otpNew)

	// 5. Save OTP -> to Redis
	err = global.Rdb.Set(ctx, userKey, strconv.Itoa(otpNew), 5*time.Minute).Err()
	if err != nil {
		return response.ErrorInvalidOTP, err
	}

	return otpNew, nil

	// 6. Send OTP -> to Email/Mobile
	// switch in.VerifyType {
	// case "EMAIL":
	// 	err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
	// 	if err != nil {
	// 		return response.ErrorSendEmailOTP, err
	// 	}

	// 	// 7. save OTP to MySQL
	// 	result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
	// 		VerifyOtp:     strconv.Itoa(otpNew),
	// 		VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
	// 		VerifyKey:     in.VerifyKey,
	// 		VerifyKeyHash: hashKey,
	// 	})

	// 	if err != nil {
	// 		return response.ErrorSendEmailOTP, err
	// 	}

	// 	// 8. getlastId
	// 	lastIdVerifyuser, err := result.LastInsertId()
	// 	if err != nil {
	// 		return response.ErrorSendEmailOTP, err
	// 	}
	// 	fmt.Printf("lastIdVerifyuser: %d\n", lastIdVerifyuser)
	// 	return response.CodeSuccess, nil
	// case "MOBILE":
	// 	// handle mobile ...
	// 	return response.CodeSuccess, nil
	// }
}

func (s *UserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *UserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
