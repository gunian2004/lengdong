package controllers

import (
	"net/http"

	"github.com/coldchain-traceability-system/services/saler-service/service"
	"github.com/gin-gonic/gin"
)

func Logistics(c *gin.Context) {
	// 1. 定义接收参数的结构体
	var logisticsData service.LogisticsData

	// 2. 绑定 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&logisticsData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求数据解析失败: " + err.Error(),
		})
		return
	}

	// 3. 检查必要字段是否为空
	if logisticsData.SKU == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少必要字段: sku",
		})
		return
	}

	// 4. 调用服务层进行上链操作
	err := service.SyncLogisticsToBlockchain(logisticsData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "物流信息上链失败: " + err.Error(),
		})
		return
	}

	// 5. 成功返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "物流信息已成功上链",
	})
}
