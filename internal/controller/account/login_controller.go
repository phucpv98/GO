package account

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/models"
	"go-ecommerce/internal/service"
	"go-ecommerce/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Management controller Login User
var LoginController = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamsInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CodeSuccess, "Login success")
}

func (c *cUserLogin) Register(ctx *gin.Context) {
	var params models.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamsInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}
