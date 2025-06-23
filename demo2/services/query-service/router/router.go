package router

import (
	"github.com/coldchain-traceability-system/services/query-service/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, queryController *controllers.QueryController) {
	// 添加新的路由规则
	r.GET("/products", queryController.GetProductList)
	r.GET("/products/:sku", queryController.GetProductBySKU)
}
