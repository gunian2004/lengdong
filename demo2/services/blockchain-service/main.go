package main

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/blockchain-service/router"
	"github.com/coldchain-traceability-system/services/blockchain-service/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	config.LoadConfig(".")

	// 初始化数据库
	mysql.InitDB()

	// 初始化 Redis
	redis.InitRedis()

	// 初始化区块链服务
	if err := service.InitBlockchainService(); err != nil {
		log.Fatalf("区块链服务初始化失败: %v", err)
	}

	// 设置 Gin 运行模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 初始化 Gin 路由器
	r := gin.Default()

	// 应用中间件
	r.Use(middleware.CorsMiddleware())

	// 配置路由
	router.SetupRouter(r)

	// 启动服务
	serverAddr := "8081"
	log.Printf("区块链服务启动在 %s", serverAddr)
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
