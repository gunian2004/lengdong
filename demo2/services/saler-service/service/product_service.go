package service

import (
	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/models"
	"gorm.io/gorm"
)

// IProductService 定义产品服务接口
type IProductService interface {
	GetAvailableProducts() ([]ProductWithQuantity, error)
}

// ProductService 实现产品服务接口
type ProductService struct {
	db  *gorm.DB
	cfg *config.Config
}

// ProductWithQuantity 包含产品信息和对应批次的数量
type ProductWithQuantity struct {
	models.FrozenProduct
	Quantity string `json:"quantity"` // 批次数量
}

// NewProductService 创建产品服务实例
func NewProductService(db *gorm.DB, cfg *config.Config) IProductService {
	return &ProductService{db: db, cfg: cfg}
}

// GetAvailableProducts 获取可用的冷冻品列表及其数量
func (ps *ProductService) GetAvailableProducts() ([]ProductWithQuantity, error) {
	var productsWithQuantity []ProductWithQuantity

	// 查询所有可用的冷冻品及其批次数量
	err := ps.db.Table("frozen_products").
		Select("frozen_products.*, product_batches.quantity").
		Joins("JOIN product_batches ON frozen_products.sku = product_batches.sku").
		Where("product_batches.quantity > 0").
		Scan(&productsWithQuantity).Error

	return productsWithQuantity, err
}
