package service

import (
	"log"

	"github.com/coldchain-traceability-system/common/ethereum"
	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/database/mysql"
)

// LogisticsData 表示前端传入的物流信息
type LogisticsData struct {
	SKU             string `json:"sku" binding:"required"`
	TransportMethod string `json:"transport_method" binding:""`
	Carrier         string `json:"carrier" binding:"required"`
	Temperature     string `json:"temperature" binding:"required"`
	Departure       string `json:"departure" binding:"required"`
	Destination     string `json:"destination" binding:"required"`
	OperatorName    string `json:"operator_name"` // 操作者名称
}

// SyncLogisticsToBlockchain 将物流信息同步到以太坊区块链
func SyncLogisticsToBlockchain(data LogisticsData) error {
	var product models.FrozenProduct

	// 查询指定 SKU 的产品
	if err := mysql.DB.Where("sku = ?", data.SKU).Preload("ProductBatches").First(&product).Error; err != nil {
		log.Printf("未找到对应产品: %v", err)
		return err
	}

	log.Printf("开始上链物流信息，SKU: %s，运往: %s", data.SKU, data.Destination)

	// 调用以太坊合约 AddLogisticsRecord 方法
	txHash, err := ethereum.Contract.AddLogisticsRecord(
		ethereum.Auth,
		data.SKU,
		data.TransportMethod,
		data.Carrier,
		data.Temperature,
		data.Departure,
		data.Destination,
		data.OperatorName,
	)
	if err != nil {
		log.Printf("发送物流信息到以太坊合约失败: %v", err)
		return err
	}

	log.Printf("以太坊交易哈希: %s", txHash.Hash().Hex())
	var logicst models.LogisticsInfo

	logicst.BlockchainTxHash = txHash.Hash().Hex()
	if err := mysql.DB.Save(&product).Error; err != nil {
		log.Printf("更新数据库中的物流交易哈希失败: %v", err)
		return err
	}

	log.Printf("产品 %s 的物流信息已成功上链", data.SKU)
	return nil
}
