package initialize

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/database"
	"go-ecommerce/internal/service"
	"go-ecommerce/internal/service/implements"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)

	// User Service Interface
	service.InitUserLogin(implements.NewUserLoginImpl(queries))
}
