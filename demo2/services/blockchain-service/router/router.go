package router

import (
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/services/blockchain-service/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the blockchain service router
func SetupRouter(r *gin.Engine) {
	blockchainController := controllers.NewBlockchainController()

	// All blockchain routes require authentication
	blockchain := r.Group("/api/v1/blockchain")
	blockchain.Use(middleware.JWTAuthMiddleware())
	{
		blockchain.POST("/products", blockchainController.AddProductToBlockchain)
		blockchain.POST("/logistics", blockchainController.UpdateLogisticsInfo)
		blockchain.POST("/transfer", blockchainController.TransferProductOwnership)
		blockchain.GET("/products/:sku", blockchainController.QueryProductFromBlockchain)
	}
} 