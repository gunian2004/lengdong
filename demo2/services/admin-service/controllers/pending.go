package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/coldchain-traceability-system/services/admin-service/service"
	"github.com/gin-gonic/gin"
)

// AuditRequest 定义请求体结构体
type AuditRequest struct {
	Username     string `json:"username" binding:"required"`
	ProductName  string `json:"product_name" binding:"required"`
	IsSuccess    string `json:"is_success" binding:"required,oneof=true false 1 0 approved rejected"`
	AuditComment string `json:"audit_comment"`
	AuditorID    string `json:"auditor_id" binding:"required"`
}

// 角色审核
func UpdateUserPending(c *gin.Context) {
	var req struct {
		Username     string `json:"username" binding:"required"`
		IsSuccess    string `json:"is_success" binding:"required,oneof=true false 1 0 approved rejected"`
		AuditComment string `json:"audit_comment"`
		AuditorID    string `json:"auditor_id" binding:"required"`
	}
	// 使用 ShouldBindJSON 绑定请求体数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 去除空格和双引号
	req.IsSuccess = strings.TrimSpace(strings.Trim(req.IsSuccess, `"`))
	req.AuditComment = strings.TrimSpace(strings.Trim(req.AuditComment, `"`))

	fmt.Println("is_success:", req.IsSuccess, "audit_comment:", req.AuditComment, "auditor_id:", req.AuditorID)
	// 将 AuditorID 从字符串转换为 uint64
	auditorID, err := strconv.ParseUint(req.AuditorID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的审核人ID"})
		return
	}
	if err := service.UpdateUserPending(req.Username, req.IsSuccess, req.AuditComment, auditorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "审核完成"})
}

// 冷冻品审核
func UpdateProductPending(c *gin.Context) {
	// AuditRequest 定义请求体结构体

	var req struct {
		ProductName  string `json:"product_name" binding:"required"`
		IsSuccess    string `json:"is_success" binding:"required,oneof=true false 1 0 approved rejected"`
		AuditComment string `json:"audit_comment"`
		AuditorID    string `json:"auditor_id" binding:"required"`
	}
	// 使用 ShouldBindJSON 绑定请求体数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 去除空格和双引号
	req.IsSuccess = strings.TrimSpace(strings.Trim(req.IsSuccess, `"`))
	req.AuditComment = strings.TrimSpace(strings.Trim(req.AuditComment, `"`))

	fmt.Println("is_success:", req.IsSuccess, "audit_comment:", req.AuditComment, "auditor_id:", req.AuditorID)
	// 将 AuditorID 从字符串转换为 uint64
	auditorID, err := strconv.ParseUint(req.AuditorID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的审核人ID"})
		return
	}
	if err := service.UpdateProductPending(req.ProductName, req.IsSuccess, req.AuditComment, auditorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "审核完成"})
}
