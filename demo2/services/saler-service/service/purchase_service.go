package service

import (
	"errors"
	"fmt" // 确保导入 fmt 包
	"strconv"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/models"

	"gorm.io/gorm"
)

// 定义采购配置结构体
type PurchaseConfig struct {
	ValidBuyerRoles  map[string]bool
	ValidSellerRoles map[string]bool
}

// 默认采购配置
var defaultPurchaseConfig = PurchaseConfig{
	ValidBuyerRoles: map[string]bool{
		"distributor": true,
		"retailer":    true,
	},
	ValidSellerRoles: map[string]bool{
		"factory":     true,
		"distributor": true,
	},
}

// IPurchaseService 接口定义
type IPurchaseService interface {
	CreatePurchaseRecord(record *models.PurchaseRecord) error
	GetPurchaseRecords() ([]models.PurchaseRecord, error)
	ConfirmPurchase(id uint /*, userID uint */) error
}

// PurchaseService 处理进货记录相关逻辑
type PurchaseService struct {
	db           *gorm.DB
	cfg          *config.Config
	purchaseConf PurchaseConfig
}

// NewPurchaseService 创建服务实例
func NewPurchaseService(db *gorm.DB, cfg *config.Config) IPurchaseService {
	// 使用默认配置
	purchaseConf := defaultPurchaseConfig

	// 这里可以从 cfg 中读取采购配置，如果有的话
	// 例如：if cfg != nil && cfg.HasKey("purchase.valid_buyer_roles") { ... }

	return &PurchaseService{
		db:           db,
		cfg:          cfg,
		purchaseConf: purchaseConf,
	}
}

// CreatePurchaseRecord 创建进货记录
func (ps *PurchaseService) CreatePurchaseRecord(record *models.PurchaseRecord) error {
	// 开始事务
	tx := ps.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询产品批次
	var batch models.ProductBatch
	if err := tx.Where("sku = ?", record.SKU).First(&batch).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to query product batch: %w", err)
	}

	// 将批次ID赋值给采购记录
	record.BatchID = batch.BatchID

	// 将批次数量转换为uint进行比较
	batchQuantity, err := strconv.ParseUint(batch.Quantity, 10, 32)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("invalid batch quantity format: %w", err)
	}

	// 检查库存是否足够
	if uint(batchQuantity) < record.Quantity {
		tx.Rollback()
		return errors.New("insufficient stock")
	}

	// 计算新的库存数量
	newQuantity := uint(batchQuantity) - record.Quantity

	// 更新批次库存
	if err := tx.Model(&batch).Update("quantity", strconv.FormatUint(uint64(newQuantity), 10)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update batch quantity: %w", err)
	}

	// 创建采购记录
	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create purchase record: %w", err)
	}

	// 提交事务
	return tx.Commit().Error
}

// GetPurchaseRecords 获取进货记录
func (ps *PurchaseService) GetPurchaseRecords() ([]models.PurchaseRecord, error) {
	var records []models.PurchaseRecord
	fmt.Println("Attempting to get purchase records from DB...")

	result := ps.db.Find(&records)
	if result.Error != nil {
		fmt.Printf("Error getting purchase records from DB: %v\n", result.Error)
		return nil, result.Error
	}

	fmt.Printf("Successfully retrieved %d purchase records\n", len(records))
	return records, nil
}

// ConfirmPurchase 确认进货记录
func (ps *PurchaseService) ConfirmPurchase(id uint) error {
	tx := ps.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var record models.PurchaseRecord
	if err := tx.First(&record, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("进货记录不存在")
		}
		return err
	}

	// 更新状态为已确认
	if err := tx.Model(&record).Update("status", "confirmed").Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// --- 移除以下重复的定义 ---
// type IPurchaseService interface { ... }
// type PurchaseService struct { ... }
// func NewPurchaseService(db *gorm.DB) IPurchaseService { ... }
