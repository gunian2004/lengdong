package main

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/user-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	config.LoadConfig("../../")

	// 初始化数据库
	mysql.InitDB()

	// 初始 Redis 连接
	redis.InitRedis()

	// 初始化 Gin 路由器
	r := gin.Default()

	// 应用中间件
	r.Use(middleware.CorsMiddleware())

	// 配置路由
	router.NewRouter(r)

	// 启动服务器
	serverAddr := "8085"
	log.Printf("用户服务正在 %s 上启动", serverAddr)
	if err := r.Run(":8085"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
