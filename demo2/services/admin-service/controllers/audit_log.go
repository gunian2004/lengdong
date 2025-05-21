package controllers

import (
	"net/http"

	"github.com/coldchain-traceability-system/services/admin-service/service"
	"github.com/gin-gonic/gin"
)

func GetAuditLog(c *gin.Context) {
	auditlog, err := service.ViewAuditLog()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"auditlog": auditlog})

}
