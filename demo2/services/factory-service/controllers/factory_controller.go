// controllers/product.go
package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/coldchain-traceability-system/services/factory-service/service"
	"github.com/gin-gonic/gin"
)

// CreateProduct 处理添加冷冻品的请求
func CreateProduct(c *gin.Context) {
	var request struct {
		Product service.FrozenProductBinding `json:"product" binding:"required"`
		Batch   service.ProductBatchBinding  `json:"batch" binding:"required"`
	}

	// 绑定请求体到 request 结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求体解析错误: " + err.Error()})
		return
	}
	// 调用 AddFrozenProduct 函数
	if err := service.AddFrozenProduct(request.Product, request.Batch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加冷冻品失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "冷冻品添加成功"})
}

// 查看冷冻品
func ViewProduct(c *gin.Context) {
	manufacturer_id := c.Query("manufacturer_id")
	product, err := service.ViewProduct(manufacturer_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"product": product})
}
func UploadHandler(c *gin.Context) {
	// 从请求中获取文件
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "无法获取上传的文件",
		})
		fmt.Println("无法获取上传的文件", err)
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "无法打开上传的文件",
		})
		fmt.Println("无法打开上传的文件", err)
		return
	}
	defer src.Close()

	// 定义目标路径
	staticDir := "./static"
	imagesDir := "./static/Images"

	// 检查 static 文件夹是否存在，不存在则创建
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		if err := os.Mkdir(staticDir, os.ModePerm); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "无法创建 static 文件夹",
			})
			fmt.Println("无法创建 static 文件夹", err)
			return
		}
	}

	// 检查 Images 文件夹是否存在，不存在则创建
	if _, err := os.Stat(imagesDir); os.IsNotExist(err) {
		if err := os.Mkdir(imagesDir, os.ModePerm); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "无法创建 Images 文件夹",
			})
			fmt.Println("无法创建 Images 文件夹", err)
			return
		}
	}

	// 构建目标文件路径
	dstPath := fmt.Sprintf("%s/%s", imagesDir, file.Filename)

	// 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "无法创建目标文件",
		})
		fmt.Println("无法创建目标文件", err)
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "保存文件时发生错误",
		})
		fmt.Println("保存文件时发生错误", err)
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("文件 %s 成功上传至 %s", file.Filename, dstPath),
	})
}

// 上链
func Chain(c *gin.Context) {
	// 从 URL 查询参数中获取 sku
	sku := c.Query("sku")
	if sku == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少参数: sku"})
		return
	}

	// 调用 service 层的 SyncToBlockchain 函数
	err := service.SyncToBlockchain(sku)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上链失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "产品已成功上链"})
}
