package controllers

import (
	"log"
	"net/http"
	"strconv" // 添加 strconv 包

	"github.com/coldchain-traceability-system/services/saler-service/service"

	"github.com/coldchain-traceability-system/common/models"
	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	purchaseService service.IPurchaseService
}

// 将参数类型从 *service.PurchaseService 修改为 service.IPurchaseService
func NewPurchaseController(ps service.IPurchaseService) *PurchaseController {
	return &PurchaseController{purchaseService: ps}
}

// CreatePurchaseRecord 处理创建进货记录的请求
func (pc *PurchaseController) CreatePurchaseRecord(c *gin.Context) {
	var req models.PurchaseRecord
	log.Println("Attempting to bind JSON for CreatePurchaseRecord")
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON in CreatePurchaseRecord: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Printf("JSON bound successfully for CreatePurchaseRecord. Request data: %+v", req)
	log.Println("Calling purchaseService.CreatePurchaseRecord")

	// 这里不需要修改，因为接口类型同样有 CreatePurchaseRecord 方法
	if err := pc.purchaseService.CreatePurchaseRecord(&req); err != nil {
		log.Printf("Error calling purchaseService.CreatePurchaseRecord: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase record"})
		return
	}

	log.Println("purchaseService.CreatePurchaseRecord completed successfully")
	c.JSON(http.StatusCreated, req)
}

// GetPurchaseRecords 处理获取进货记录列表的请求
func (pc *PurchaseController) GetPurchaseRecords(c *gin.Context) {
	buyerIDStr := c.Query("buyer_id")
	if buyerIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "buyer_id is required"})
		return
	}

	// 将 buyerID 字符串转换为 uint64
	buyerID, err := strconv.ParseUint(buyerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buyer ID format"})
		return
	}

	records, err := pc.purchaseService.GetPurchaseRecords()
	if err != nil {
		log.Printf("Error getting purchase records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get purchase records"})
		return
	}

	// 过滤指定买家的记录
	var buyerRecords []models.PurchaseRecord
	for _, record := range records {
		if record.BuyerID == buyerID {
			buyerRecords = append(buyerRecords, record)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"purchases": buyerRecords,
	})
}

// ConfirmPurchase 处理确认收货的请求
func (pc *PurchaseController) ConfirmPurchase(c *gin.Context) {
	purchaseID := c.Param("id")
	if purchaseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchase_id is required"})
		return
	}

	id, err := strconv.ParseUint(purchaseID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID format"})
		return
	}

	if err := pc.purchaseService.ConfirmPurchase(uint(id)); err != nil {
		log.Printf("Error confirming purchase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm purchase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Purchase confirmed successfully",
	})
}
