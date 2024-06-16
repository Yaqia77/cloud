package main

import (
	"cloud/config/mysql"
	"cloud/config/redis"
	"cloud/router"
	"fmt"
)

func main() {
	// 创建 MySQL 配置实例
	mysqlConfig := mysql.NewMySQL(
		mysql.MaxOpenConns(10),
		mysql.MaxIdleConns(5),
	)

	// 从环境变量加载配置
	cfg := redis.LoadConfig()

	// 初始化 Redis 客户端
	redis.InitRedis(cfg)

	// 测试 Redis 连接
	val, err := redis.Client.Ping(redis.Ctx).Result()
	if err != nil {
		fmt.Println("Redis 连接失败:", err)
	} else {
		fmt.Println("Redis 连接成功:", val)
	}
	// 初始化数据库连接
	mysqlConfig.InitDB()

	var r = router.SetupRouter()
	r.Run(":8080")

}
