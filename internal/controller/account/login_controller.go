package account

import (
	"go-ecommerce/internal/service"
	"go-ecommerce/response"

	"github.com/gin-gonic/gin"
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
