package initialize

import (
	"database/sql"
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/internal/po"
	"time"

	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DBName)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "InitMysqlc initialization error")
	global.Logger.Info("InitMysqlc initialization success")
	global.Mdbc = db

	// SetPool() : mở nhóm kết nối - cải thiện hiệu suất, thời gian sống của kết nối, v.v.
	// SetPoolC()
	// migrateTablesC()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB() // lấy đối tượng *sql.DB từ gorm.DB để cấu hình kết nối cơ sở dữ liệu
	if err != nil {
		fmt.Println("SetPool error: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))    // đặt thời gian tối đa mà một kết nối có thể ở trạng thái nhàn rỗi trước khi bị đóng
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)                      // đặt số lượng kết nối tối đa mà cơ sở dữ liệu có thể mở cùng một lúc
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // đặt thời gian tối đa mà một kết nối có thể tồn tại trước khi bị đóng
}

func migrateTablesC() {
	// persystent objects : 'po' folder - chứa các đối tượng liên quan đến cơ sở dữ liệu, chẳng hạn như mô hình, trình quản lý, v.v.
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		fmt.Println("Migrating tables error: ", err)
	}
}
