package global

import (
	"go-ecommerce/pkg/logger"
	"go-ecommerce/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

/*
Config
	Redis
	Mysql
	Etc ...
*/
