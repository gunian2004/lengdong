package main

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/ethereum"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/saler-service/router"
	"github.com/gin-gonic/gin"
)

// 123
func main() {
	// 加载配置文件
	config.LoadConfig("../../")

	// 初始化数据库
	mysql.InitDB()

	// 初始化 Redis
	redis.InitRedis()
	//初始化以太坊
	ethereum.InitEthereum()
	// 设置 Gin 运行模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 创建路由引擎
	r := router.NewRouter(&config.AppConfig)

	// 启动服务
	serverAddr := ":8084"
	log.Printf("经销商服务启动在 %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
