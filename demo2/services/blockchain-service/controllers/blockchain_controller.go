package controllers

import (
	"net/http"

	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/common/utils"
	"github.com/coldchain-traceability-system/services/blockchain-service/service"
	"github.com/gin-gonic/gin"
)

// BlockchainController handles blockchain HTTP requests
type BlockchainController struct{}

// NewBlockchainController creates a new blockchain controller
func NewBlockchainController() *BlockchainController {
	return &BlockchainController{}
}

// AddProductToBlockchain adds a product to the blockchain
func (c *BlockchainController) AddProductToBlockchain(ctx *gin.Context) {
	var product models.FrozenProduct
	if err := ctx.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request body: "+err.Error())
		return
	}
	
	txHash, err := service.AddProductToBlockchain(&product)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to add product to blockchain: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, gin.H{
		"tx_hash": txHash,
		"message": "Product added to blockchain successfully",
	})
}

// UpdateLogisticsInfo updates logistics information on the blockchain
func (c *BlockchainController) UpdateLogisticsInfo(ctx *gin.Context) {
	var logistics models.LogisticsInfo
	if err := ctx.ShouldBindJSON(&logistics); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request body: "+err.Error())
		return
	}
	
	txHash, err := service.UpdateLogisticsInfo(&logistics)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update logistics info on blockchain: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, gin.H{
		"tx_hash": txHash,
		"message": "Logistics information updated on blockchain successfully",
	})
}

// TransferProductOwnership transfers ownership of a product on the blockchain
func (c *BlockchainController) TransferProductOwnership(ctx *gin.Context) {
	var transferReq struct {
		SKU       string `json:"sku" binding:"required"`
		FromUser  uint64 `json:"from_user" binding:"required"`
		ToUser    uint64 `json:"to_user" binding:"required"`
	}
	
	if err := ctx.ShouldBindJSON(&transferReq); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request body: "+err.Error())
		return
	}
	
	txHash, err := service.TransferProductOwnership(transferReq.SKU, transferReq.FromUser, transferReq.ToUser)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to transfer product ownership on blockchain: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, gin.H{
		"tx_hash": txHash,
		"message": "Product ownership transferred on blockchain successfully",
	})
}

// QueryProductFromBlockchain queries product information from the blockchain
func (c *BlockchainController) QueryProductFromBlockchain(ctx *gin.Context) {
	sku := ctx.Param("sku")
	if sku == "" {
		utils.BadRequestResponse(ctx, "Product SKU is required")
		return
	}
	
	product, err := service.QueryProductFromBlockchain(sku)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to query product from blockchain: "+err.Error())
		return
	}
	
	utils.SuccessResponse(ctx, product)
} 