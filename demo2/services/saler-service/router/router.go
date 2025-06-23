package router

import (
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/services/saler-service/controllers"
	"github.com/coldchain-traceability-system/services/saler-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewRouter 创建并配置 Gin 引擎
func NewRouter(cfg *config.Config) *gin.Engine {
	// 确保数据库已被初始化
	if mysql.DB == nil {
		log.Fatalf("数据库连接为空，请确保在调用 NewRouter 前已初始化数据库")
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化 Gin 引擎
	r := gin.Default()

	// 使用 CORS 中间件
	r.Use(middleware.CorsMiddleware())

	// 设置路由
	SetupRoutes(mysql.DB, cfg, r)

	return r
}

// SetupRoutes 设置所有路由
func SetupRoutes(db *gorm.DB, cfg *config.Config, r *gin.Engine) {
	// 初始化 Service 和 Controller
	purchaseService := service.NewPurchaseService(db, cfg)
	purchaseController := controllers.NewPurchaseController(purchaseService)

	logisticsService := service.NewLogisticsService(db, cfg)
	logisticsController := controllers.NewLogisticsController(logisticsService)

	// 初始化产品服务和控制器
	productService := service.NewProductService(db, cfg)
	productController := controllers.NewProductController(productService)
	r.POST("/logisticsNew", controllers.Logistics)
	// 设置路由组
	api := r.Group("/api")
	{
		// 进货相关路由
		purchase := api.Group("/purchase")
		{
			purchase.POST("", purchaseController.CreatePurchaseRecord)
			purchase.GET("", purchaseController.GetPurchaseRecords)
			purchase.PUT("/confirm/:id", purchaseController.ConfirmPurchase)
		}

		// 物流相关路由
		logistics := api.Group("/logistics")
		{
			logistics.POST("", logisticsController.AddLogisticsInfo)
			logistics.POST("/batch", logisticsController.BatchAddLogisticsInfo)
			// logistics.GET("/:batch_id", logisticsController.GetLogisticsInfoByBatchID) // 保持注释或实现它
		}

		// 产品相关路由
		products := api.Group("/products")
		{
			products.GET("/available", productController.GetAvailableProducts)
		}
	}
}

// 删除了这里重复的代码块
