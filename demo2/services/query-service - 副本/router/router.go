package router

import (
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/services/query-service/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the query service router
func SetupRouter(r *gin.Engine) {
	queryController := controllers.NewQueryController()

	// Public routes (no authentication required)
	public := r.Group("/api/v1/query")
	{
		public.GET("/products/:sku", queryController.GetProduct)
		public.GET("/products/search", queryController.SearchProducts)
	}

	// Protected routes (require authentication)
	protected := r.Group("/api/v1/query")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/products/:sku/blockchain", queryController.GetProductFromBlockchain)
		protected.GET("/products/:sku/traceability", queryController.GetProductTraceability)
		protected.GET("/products/:sku/batches", queryController.GetBatchesBySKU)
		protected.GET("/logistics/:id", queryController.GetLogisticsInfo)
	}
} 