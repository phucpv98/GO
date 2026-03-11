package initialize

import (
	"go-ecommerce/global"
	"go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
