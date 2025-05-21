package service

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/coldchain-traceability-system/blockchain/chaincode"
	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/models"
)

var bcClient *chaincode.BlockchainClient

// 初始化区块链服务，创建区块链客户端
func InitBlockchainService() error {
	var err error
	bcConfig := config.AppConfig.Blockchain
	bcClient, err = chaincode.NewBlockchainClient(
		bcConfig.NetworkConfigPath,
		bcConfig.ChannelID,
		bcConfig.ChaincodeName,
		bcConfig.OrgAdmin,
		bcConfig.OrgName,
	)
	if err != nil {
		log.Printf("初始化区块链客户端失败: %v", err)
		return err
	}
	log.Println("区块链客户端初始化成功")
	return nil
}

// 将产品添加到区块链
func AddProductToBlockchain(product *models.FrozenProduct) (string, error) {
	log.Printf("将产品 %s 添加到区块链", product.SKU)
	txHash, err := bcClient.AddProduct(product.SKU, product)
	if err != nil {
		log.Printf("将产品添加到区块链失败: %v", err)
		return "", err
	}

	// // 创建区块链记录
	// blockchainRecord := &models.BlockchainRecord{
	// 	TxHash:        txHash,
	// 	SKU:           product.SKU,
	// 	OperationType: "product_publish",
	// 	RawData:       models.JSONMap{"product": product},
	// }

	// 保存区块链记录到数据库 - 我们将在工厂服务中实现此功能
	// 因为它可以直接访问数据库

	return txHash, nil
}

// 更新区块链上的物流信息
func UpdateLogisticsInfo(logistics *models.LogisticsInfo) (string, error) {
	log.Printf("更新采购 ID %d 的物流信息", logistics.PurchaseID)
	txHash, err := bcClient.UpdateLogistics(logistics.LogisticsNo, logistics)
	if err != nil {
		log.Printf("更新区块链上的物流信息失败: %v", err)
		return "", err
	}

	// // 创建区块链记录
	// blockchainRecord := &models.BlockchainRecord{
	// 	TxHash:        txHash,
	// 	SKU:           "", // 我们需要从采购记录中查找 SKU
	// 	OperationType: "logistics_update",
	// 	RawData:       models.JSONMap{"logistics": logistics},
	// }

	// 保存区块链记录到数据库 - 我们将在工厂服务中实现此功能

	return txHash, nil
}

// 在区块链上转移产品的所有权
func TransferProductOwnership(sku string, fromUserID, toUserID uint64) (string, error) {
	log.Printf("将产品 %s 的所有权从用户 %d 转移到用户 %d", sku, fromUserID, toUserID)
	txHash, err := bcClient.TransferOwnership(sku, strconv.FormatUint(fromUserID, 10),
		strconv.FormatUint(toUserID, 10))
	if err != nil {
		log.Printf("转移区块链上的产品所有权失败: %v", err)
		return "", err
	}

	// // 创建区块链记录
	// blockchainRecord := &models.BlockchainRecord{
	// 	TxHash:        txHash,
	// 	SKU:           sku,
	// 	OperationType: "ownership_transfer",
	// 	RawData:       models.JSONMap{"from_user_id": fromUserID, "to_user_id": toUserID},
	// }

	// 保存区块链记录到数据库 - 我们将在工厂服务中实现此功能

	return txHash, nil
}

// 从区块链查询产品信息
func QueryProductFromBlockchain(sku string) (*models.FrozenProduct, error) {
	log.Printf("从区块链查询产品 %s", sku)
	payload, err := bcClient.QueryProduct(sku)
	if err != nil {
		log.Printf("从区块链查询产品失败: %v", err)
		return nil, err
	}

	var product models.FrozenProduct
	err = json.Unmarshal(payload, &product)
	if err != nil {
		log.Printf("反序列化产品数据失败: %v", err)
		return nil, err
	}

	return &product, nil
}
