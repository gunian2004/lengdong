package controllers

import (
	"log"
	"net/http"

	"github.com/coldchain-traceability-system/services/saler-service/service"
	"github.com/gin-gonic/gin"
)

// ProductController 处理产品相关的请求
type ProductController struct {
	productService service.IProductService
}

// NewProductController 创建产品控制器实例
func NewProductController(ps service.IProductService) *ProductController {
	return &ProductController{productService: ps}
}

// GetAvailableProducts 处理获取可用冷冻品的请求
func (pc *ProductController) GetAvailableProducts(c *gin.Context) {
	log.Println("正在获取可用冷冻品列表")

	products, err := pc.productService.GetAvailableProducts()
	if err != nil {
		log.Printf("获取可用冷冻品失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取可用冷冻品失败"})
		return
	}

	log.Printf("成功获取到 %d 个可用冷冻品", len(products))
	c.JSON(http.StatusOK, gin.H{"products": products})
}
