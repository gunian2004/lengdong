package router

import (
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/services/factory-service/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {

	// 商品添加路由
	r.POST("/products", middleware.JWTAuthMiddleware(), controllers.CreateProduct)
	//商品查看路由
	r.GET("/view_product", middleware.JWTAuthMiddleware(), controllers.ViewProduct)
	r.POST("/upload", middleware.JWTAuthMiddleware(), controllers.UploadHandler)
	r.POST("/chain", middleware.JWTAuthMiddleware(), controllers.Chain)

}
