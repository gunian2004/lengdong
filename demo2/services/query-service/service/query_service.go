package service

import (
	"errors"
	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/database/mysql"
	bchainService "github.com/coldchain-traceability-system/services/blockchain-service/service"
	"gorm.io/gorm"
)

// QueryService provides methods for query operations
type QueryService struct {
	db *gorm.DB
}

// NewQueryService creates a new query service
func NewQueryService() *QueryService {
	return &QueryService{
		db: mysql.DB,
	}
}

// GetProductBySKU gets a product by SKU from database
func (s *QueryService) GetProductBySKU(sku string) (*models.FrozenProduct, error) {
	var product models.FrozenProduct
	if err := s.db.Where("sku = ?", sku).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// GetProductBySKUFromBlockchain gets a product by SKU from blockchain
func (s *QueryService) GetProductBySKUFromBlockchain(sku string) (*models.FrozenProduct, error) {
	return bchainService.QueryProductFromBlockchain(sku)
}

// GetProductTraceability gets the complete traceability history of a product
func (s *QueryService) GetProductTraceability(sku string) (map[string]interface{}, error) {
	// Get product
	product, err := s.GetProductBySKU(sku)
	if err != nil {
		return nil, err
	}
	
	// Get manufacturer
	var manufacturer models.User
	if err := s.db.First(&manufacturer, product.ManufacturerID).Error; err != nil {
		return nil, errors.New("manufacturer information not found")
	}
	
	// Get batches
	var batches []models.ProductBatch
	if err := s.db.Where("sku = ?", sku).Find(&batches).Error; err != nil {
		return nil, err
	}
	
	// Get purchase records for each batch
	var purchaseRecords []models.PurchaseRecord
	var batchIDs []uint64
	for _, batch := range batches {
		batchIDs = append(batchIDs, batch.BatchID)
	}
	
	if len(batchIDs) > 0 {
		if err := s.db.Where("batch_id IN ?", batchIDs).Find(&purchaseRecords).Error; err != nil {
			return nil, err
		}
	}
	
	// Get logistics info for each purchase record
	var logisticsInfos []models.LogisticsInfo
	var purchaseIDs []uint64
	for _, record := range purchaseRecords {
		purchaseIDs = append(purchaseIDs, record.PurchaseID)
	}
	
	if len(purchaseIDs) > 0 {
		if err := s.db.Where("purchase_id IN ?", purchaseIDs).Find(&logisticsInfos).Error; err != nil {
			return nil, err
		}
	}
	
	// Get blockchain records
	var blockchainRecords []models.BlockchainRecord
	if err := s.db.Where("sku = ?", sku).Find(&blockchainRecords).Error; err != nil {
		return nil, err
	}
	
	// Combine all data into a response
	return map[string]interface{}{
		"product":           product,
		"manufacturer":      manufacturer,
		"batches":           batches,
		"purchase_records":  purchaseRecords,
		"logistics_infos":   logisticsInfos,
		"blockchain_records": blockchainRecords,
	}, nil
}

// SearchProducts searches for products
func (s *QueryService) SearchProducts(query string, page, pageSize int) ([]models.FrozenProduct, int64, error) {
	var products []models.FrozenProduct
	var total int64
	
	// Search by product name, SKU, brand
	searchQuery := "%" + query + "%"
	
	// Count total
	if err := s.db.Model(&models.FrozenProduct{}).
		Where("product_name LIKE ? OR sku LIKE ? OR brand LIKE ?", searchQuery, searchQuery, searchQuery).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply pagination
	offset := (page - 1) * pageSize
	if err := s.db.Where("product_name LIKE ? OR sku LIKE ? OR brand LIKE ?", searchQuery, searchQuery, searchQuery).
		Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

// GetBatchesBySKU gets all batches for a product SKU
func (s *QueryService) GetBatchesBySKU(sku string, page, pageSize int) ([]models.ProductBatch, int64, error) {
	var batches []models.ProductBatch
	var total int64
	
	// Count total
	if err := s.db.Model(&models.ProductBatch{}).Where("sku = ?", sku).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply pagination
	offset := (page - 1) * pageSize
	if err := s.db.Where("sku = ?", sku).Offset(offset).Limit(pageSize).Find(&batches).Error; err != nil {
		return nil, 0, err
	}
	
	return batches, total, nil
}

// GetLogisticsInfoByID gets logistics info by ID
func (s *QueryService) GetLogisticsInfoByID(logisticsID uint64) (*models.LogisticsInfo, error) {
	var logistics models.LogisticsInfo
	if err := s.db.First(&logistics, logisticsID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("logistics information not found")
		}
		return nil, err
	}
	return &logistics, nil
} 