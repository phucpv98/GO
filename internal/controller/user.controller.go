package controller

import (
	"go-ecommerce/internal/service"
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
	email := c.Query("email")
	purpose := c.Query("purpose")

	errCode := uc.userService.Register(email, purpose)
	response.SuccessResponse(c, errCode, nil)
}
