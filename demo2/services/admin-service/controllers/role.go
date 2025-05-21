package controllers

import (
	"net/http"

	"github.com/coldchain-traceability-system/services/admin-service/service"
	"github.com/gin-gonic/gin"
)

func UpdateRole(c *gin.Context) {
	// 从前端获取用户的信息
	var admin struct {
		Username string `json:"username" binding:"required"`
		Role     string `json:"role" binding:"required,oneof=super_admin normal_admin"`
	}
	// 使用 ShouldBind 绑定请求体数据
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 验证输入数据
	if admin.Username == "" || admin.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名和角色不能为空"})
		return
	}

	if err := service.UpdateRole(admin.Username, admin.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "修改身份信息成功"})
}
