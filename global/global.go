package global

import (
	"database/sql"
	"go-ecommerce/pkg/logger"
	"go-ecommerce/pkg/setting"

	"github.com/redis/go-redis/v9"
	kafka "github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)

/*
Config
	Redis
	Mysql
	Etc ...
*/
