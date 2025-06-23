package controllers

import (
	"net/http"
	"strconv"

	"github.com/coldchain-traceability-system/services/query-service/service"
	"github.com/gin-gonic/gin"
)

// QueryController 处理查询相关的请求
type QueryController struct {
	queryService service.IQueryService
}

// NewQueryController 创建一个新的查询控制器实例
func NewQueryController(qs service.IQueryService) *QueryController {
	return &QueryController{
		queryService: qs,
	}
}

// GetProductBySKU 处理根据 SKU 码查询冷冻品信息的请求
func (qc *QueryController) GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")
	if sku == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SKU 码不能为空"})
		return
	}

	result, err := qc.queryService.GetProductBySKU(sku)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetProductList 处理查询冷冻品列表的请求
func (qc *QueryController) GetProductList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的页码参数"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的每页数量参数"})
		return
	}

	filters := make(map[string]interface{})
	for key, values := range c.Request.URL.Query() {
		if key != "page" && key != "page_size" {
			filters[key] = values[0]
		}
	}

	result, err := qc.queryService.GetProductList(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": result})
}
