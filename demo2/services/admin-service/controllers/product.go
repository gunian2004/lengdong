package controllers

import (
	"github.com/coldchain-traceability-system/services/admin-service/service"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	product, err := service.ViewProduct()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"product": product})
}
