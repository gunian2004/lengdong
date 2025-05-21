package controllers

import (

	"log" // 导入 log 包
	"net/http"

	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/services/saler-service/service"
	"github.com/gin-gonic/gin"
)

type LogisticsController struct {
	// 将类型从 *service.LogisticsService 修改为 service.ILogisticsService
	logisticsService service.ILogisticsService
}

// 将参数类型从 *service.LogisticsService 修改为 service.ILogisticsService
func NewLogisticsController(ls service.ILogisticsService) *LogisticsController {
	return &LogisticsController{logisticsService: ls}
}

// AddLogisticsInfo 处理添加物流信息的请求
func (lc *LogisticsController) AddLogisticsInfo(c *gin.Context) {
	var req models.LogisticsInfo
	log.Println("Attempting to bind JSON for AddLogisticsInfo") // 添加日志
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON in AddLogisticsInfo: %v", err) // 添加错误日志
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Printf("JSON bound successfully for AddLogisticsInfo. Request data: %+v", req) // 添加日志，打印请求数据
	log.Println("Calling logisticsService.AddLogisticsInfo")                           // 添加日志

	// 调用服务层方法
	if err := lc.logisticsService.AddLogisticsInfo(&req); err != nil {
		log.Printf("Error calling logisticsService.AddLogisticsInfo: %v", err) // 添加错误日志
		// 保持原始的错误响应，但现在我们有了更详细的服务器日志
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add logistics info"})
		return
	}

	log.Println("logisticsService.AddLogisticsInfo completed successfully") // 添加日志
	c.JSON(http.StatusCreated, req)
}

// BatchAddLogisticsInfo 处理批量添加物流信息的请求
func (lc *LogisticsController) BatchAddLogisticsInfo(c *gin.Context) {
	var reqs []models.LogisticsInfo
	if err := c.ShouldBindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := lc.logisticsService.BatchAddLogisticsInfo(reqs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to batch add logistics info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Batch logistics info added successfully"})
}

// GetLogisticsInfoByBatchID 处理根据批次ID获取物流信息的请求 (如果需要实现)
// func (lc *LogisticsController) GetLogisticsInfoByBatchID(c *gin.Context) {
// 	// ... 实现逻辑 ...
// }
