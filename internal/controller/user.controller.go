package controller

import (
	"fmt"
	"go-ecommerce/internal/service"
	"go-ecommerce/internal/vo"
	"go-ecommerce/response"

	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetInfoUser(c *gin.Context) {
// 	// response.SuccessResponse(c, 20001, []string{"felix", "m1", "m2"})

// 	response.SuccessResponse(c, 20001, uc.userService.GetInfoUser())
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrorCodeParamsInvalid, err.Error())
		return
	}

	fmt.Printf("Email params: %s", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
