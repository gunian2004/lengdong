package controllers

import (
	"fmt"
	"net/http"

	"github.com/coldchain-traceability-system/services/user-service/service"
	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	//解析结构体
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := service.Login(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// 注册
func Register(c *gin.Context) {
	var user struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		RoleType    string `json:"role_type" binding:"required,oneof=factory distributor retailer regulator consumer"`
		ContactInfo string `json:"contact_info" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}
	fmt.Println("user:", user)
	//解析结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析结构体失败"})
		return
	}
	if err := service.CreateUser(user.Username, user.Password, user.RoleType, user.ContactInfo, user.Address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户注册成功"})
}
