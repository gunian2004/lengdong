package main

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/ethereum"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/factory-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	config.LoadConfig("../../")

	// 初始化数据库
	mysql.InitDB()

	// 初始化 Redis
	redis.InitRedis()

	// 初始化以太坊模块
	ethereum.InitEthereum()

	// 设置 Gin 运行模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 初始化 Gin 路由器
	r := gin.Default()
	//加载静态资源文件
	r.Static("/static", "./static")

	// 应用中间件
	r.Use(middleware.CorsMiddleware())

	// 配置路由
	router.NewRouter(r)

	// 启动服务
	serverAddr := "8082"
	log.Printf("工厂服务启动在 %s", serverAddr)
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
