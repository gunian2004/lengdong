package service

import (
	"errors"

	"github.com/coldchain-traceability-system/common/models"
	"gorm.io/gorm"
)

// IQueryService 定义查询服务接口
type IQueryService interface {
	GetProductBySKU(sku string) (map[string]interface{}, error)
	GetProductList(page, pageSize int, filters map[string]interface{}) ([]map[string]interface{}, error)
}

// QueryService 实现 IQueryService 接口
type QueryService struct {
	db *gorm.DB
}

// NewQueryService 创建一个新的查询服务实例
func NewQueryService(db *gorm.DB) IQueryService {
	return &QueryService{
		db: db,
	}
}

// GetProductBySKU 根据 SKU 码查询冷冻品信息
func (s *QueryService) GetProductBySKU(sku string) (map[string]interface{}, error) {
	if sku == "" {
		return nil, errors.New("SKU 码不能为空")
	}

	var product models.FrozenProduct
	if err := s.db.Where("sku = ?", sku).Preload("ProductBatches").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("未找到该 SKU 码对应的冷冻品信息")
		}
		return nil, err
	}

	result := make(map[string]interface{})
	result["product"] = product

	if len(product.ProductBatches) > 0 {
		result["current_holder"] = product.ProductBatches[0].CurrentHolder
		result["current_location"] = product.ProductBatches[0].CurrentLocation
		result["quantity"] = product.ProductBatches[0].Quantity
	}

	return result, nil
}

// GetProductList 查询冷冻品列表，支持分页和筛选
func (s *QueryService) GetProductList(page, pageSize int, filters map[string]interface{}) ([]map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var products []models.FrozenProduct
	query := s.db.Preload("ProductBatches")

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, product := range products {
		item := make(map[string]interface{})
		item["product"] = product

		if len(product.ProductBatches) > 0 {
			item["current_holder"] = product.ProductBatches[0].CurrentHolder
			item["current_location"] = product.ProductBatches[0].CurrentLocation
			item["quantity"] = product.ProductBatches[0].Quantity
		}

		result = append(result, item)
	}

	return result, nil
}
