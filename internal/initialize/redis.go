package initialize

import (
	"context"
	"fmt"
	"go-ecommerce/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background() // Context: một kiểu dữ liệu trong Go được sử dụng để truyền thông tin về thời gian sống của một tác vụ, hủy bỏ tác vụ, hoặc truyền dữ liệu giữa các goroutine. Nó giúp quản lý tài nguyên và kiểm soát luồng thực thi trong ứng dụng.

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port), // địa chỉ và cổng của Redis server
		Password: r.Password,                           // no password set
		DB:       r.Database,                           // use default DB
		PoolSize: 10,                                   // xác định số lượng kết nối tối đa trong pool kết nối Redis, giúp cải thiện hiệu suất khi có nhiều yêu cầu đến Redis cùng lúc
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis Initialization Error: ", zap.Error(err))
	}

	fmt.Println("Redis Initialization Success")
	global.Rdb = rdb
}
