package controllers

import (
	"net/http"

	"github.com/coldchain-traceability-system/services/admin-service/service"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		RoleType    string `json:"role_type" binding:"required,oneof=factory distributor retailer regulator consumer"`
		ContactInfo string `json:"contact_info" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}
	//解析结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddUser(user.Username, user.Password, user.RoleType, user.ContactInfo, user.Address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加用户成功"})
}

// 获取厂家用户信息
func GetFactory(c *gin.Context) {
	user, err := service.ViewFactory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// 获取经销商用户信息
func GetDistributor(c *gin.Context) {
	distributor, err := service.ViewDistributor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"distributor": distributor})
}

// 获取零售商信息
func GetRetailer(c *gin.Context) {
	retailer, err := service.ViewRetailer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"retailer": retailer})
}

// 获取监管机构信息
func GetRegulator(c *gin.Context) {
	regulator, err := service.ViewRegulator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"regulator": regulator})
}

// 获取消费者路由
func GetConsumer(c *gin.Context) {
	consumer, err := service.ViewConsumer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"consumer": consumer})
}

// 获取用户路由
func GetUser(c *gin.Context) {
	user, err := service.ViewUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// 获取管理员路由
func GetAdmin(c *gin.Context) {
	admin, err := service.ViewAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}
