package initialize

import (
	"fmt"
	"go-ecommerce/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	n := global.Config.Logger
	fmt.Println("Load configuration mysql ", m.User, m.Password, m.Host, m.Port, m.DBName)
	fmt.Println("Load configuration logger ", n.Log_level, n.File_log_name, n.Max_size, n.Max_age, n.Max_backups, n.Compress)
	InitLogger()
	// InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()

	r := InitRouter()
	r.Run() // listens on 0.0.0.0:8080 by default
}
