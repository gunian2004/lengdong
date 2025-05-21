package controllers

import (
	"net/http"
	"strconv"

	"github.com/coldchain-traceability-system/common/utils"
	"github.com/coldchain-traceability-system/services/query-service/service"
	"github.com/gin-gonic/gin"
)

// QueryController handles query HTTP requests
type QueryController struct {
	queryService *service.QueryService
}

// NewQueryController creates a new query controller
func NewQueryController() *QueryController {
	return &QueryController{
		queryService: service.NewQueryService(),
	}
}

// GetProduct gets a product by SKU
func (c *QueryController) GetProduct(ctx *gin.Context) {
	// Get SKU from URL
	sku := ctx.Param("sku")
	if sku == "" {
		utils.BadRequestResponse(ctx, "Product SKU is required")
		return
	}
	
	// Get product
	product, err := c.queryService.GetProductBySKU(sku)
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, product)
}

// GetProductFromBlockchain gets a product by SKU from blockchain
func (c *QueryController) GetProductFromBlockchain(ctx *gin.Context) {
	// Get SKU from URL
	sku := ctx.Param("sku")
	if sku == "" {
		utils.BadRequestResponse(ctx, "Product SKU is required")
		return
	}
	
	// Get product from blockchain
	product, err := c.queryService.GetProductBySKUFromBlockchain(sku)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to query product from blockchain: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, product)
}

// GetProductTraceability gets the complete traceability history of a product
func (c *QueryController) GetProductTraceability(ctx *gin.Context) {
	// Get SKU from URL
	sku := ctx.Param("sku")
	if sku == "" {
		utils.BadRequestResponse(ctx, "Product SKU is required")
		return
	}
	
	// Get product traceability
	traceData, err := c.queryService.GetProductTraceability(sku)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get product traceability: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, traceData)
}

// SearchProducts searches for products
func (c *QueryController) SearchProducts(ctx *gin.Context) {
	// Get search query
	query := ctx.Query("q")
	if query == "" {
		utils.BadRequestResponse(ctx, "Search query is required")
		return
	}
	
	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	
	// Search products
	products, total, err := c.queryService.SearchProducts(query, page, pageSize)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to search products: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, gin.H{
		"products":    products,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// GetBatchesBySKU gets all batches for a product SKU
func (c *QueryController) GetBatchesBySKU(ctx *gin.Context) {
	// Get SKU from URL
	sku := ctx.Param("sku")
	if sku == "" {
		utils.BadRequestResponse(ctx, "Product SKU is required")
		return
	}
	
	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	
	// Get batches
	batches, total, err := c.queryService.GetBatchesBySKU(sku, page, pageSize)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get batches: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, gin.H{
		"batches":     batches,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// GetLogisticsInfo gets logistics info by ID
func (c *QueryController) GetLogisticsInfo(ctx *gin.Context) {
	// Get logistics ID from URL
	logisticsID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(ctx, "Invalid logistics ID")
		return
	}
	
	// Get logistics info
	logistics, err := c.queryService.GetLogisticsInfoByID(logisticsID)
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, logistics)
} 