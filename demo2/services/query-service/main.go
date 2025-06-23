package main

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/query-service/controllers"
	"github.com/coldchain-traceability-system/services/query-service/router"
	"github.com/coldchain-traceability-system/services/query-service/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	config.LoadConfig("../../")
	// 初始化数据库连接
	mysql.InitDB()
	// 初始化 Redis
	redis.InitRedis()

	// 创建查询服务实例
	queryService := service.NewQueryService(mysql.DB)

	// 创建查询控制器实例
	queryController := controllers.NewQueryController(queryService)

	// 创建 Gin 引擎
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// 配置路由
	router.NewRouter(r, queryController)

	// 启动服务器
	serverAddr := "8083"
	log.Printf("查询服务正在 %s 上启动", serverAddr)
	if err := r.Run(":" + serverAddr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
