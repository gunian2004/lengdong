package service

import (
	"fmt"
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/models"
	"gorm.io/gorm"
)

// 定义物流配置结构体
type LogisticsConfig struct {
	MinTemp     float64
	MaxTemp     float64
	MinHumidity float64
	MaxHumidity float64
}

// 默认物流配置
var defaultLogisticsConfig = LogisticsConfig{
	MinTemp:     -10.0,
	MaxTemp:     30.0,
	MinHumidity: 0.0,
	MaxHumidity: 100.0,
}

// 定义 ILogisticsService 接口
type ILogisticsService interface {
	AddLogisticsInfo(info *models.LogisticsInfo) error
	BatchAddLogisticsInfo(infos []models.LogisticsInfo) error
	// GetLogisticsInfoByBatchID(batchID string) ([]models.LogisticsInfo, error)
}

type LogisticsService struct {
	db            *gorm.DB
	cfg           *config.Config
	logisticsConf LogisticsConfig
}

// 修改 NewLogisticsService 的返回类型为 ILogisticsService
func NewLogisticsService(db *gorm.DB, cfg *config.Config) ILogisticsService {
	if db == nil {
		log.Println("警告：数据库连接为空")
	}
	if cfg == nil {
		log.Println("警告：配置为空")
	}

	// 使用默认配置
	logisticsConf := defaultLogisticsConfig

	// 这里可以从 cfg 中读取物流配置，如果有的话
	// 例如：if cfg != nil && cfg.HasKey("logistics.min_temp") { ... }

	return &LogisticsService{
		db:            db,
		cfg:           cfg,
		logisticsConf: logisticsConf,
	}
}

// AddLogisticsInfo 添加单条物流信息
func (s *LogisticsService) AddLogisticsInfo(info *models.LogisticsInfo) error {
	// 添加空值检查
	if info == nil {
		return fmt.Errorf("物流信息不能为空")
	}

	// 校验温度和湿度范围
	if info.Temperature < s.logisticsConf.MinTemp || info.Temperature > s.logisticsConf.MaxTemp {
		return fmt.Errorf("temperature out of range (%v℃ to %v℃)",
			s.logisticsConf.MinTemp, s.logisticsConf.MaxTemp)
	}

	if info.Humidity != nil && (*info.Humidity < s.logisticsConf.MinHumidity ||
		*info.Humidity > s.logisticsConf.MaxHumidity) {
		return fmt.Errorf("humidity out of range (%v%% to %v%%)",
			s.logisticsConf.MinHumidity, s.logisticsConf.MaxHumidity)
	}

	// 校验物流单号唯一性
	var existing models.LogisticsInfo
	if err := s.db.Where("logistics_no = ?", info.LogisticsNo).First(&existing).Error; err == nil {
		// 如果找到了记录 (err == nil)，说明编号已存在
		return fmt.Errorf("logistics number '%s' already exists", info.LogisticsNo)
	} else if err != gorm.ErrRecordNotFound {
		// 如果是其他数据库错误，返回错误
		return fmt.Errorf("database error checking logistics number: %w", err)
	}
	// 如果 err 是 gorm.ErrRecordNotFound，说明编号不存在，可以继续

	// 验证进货记录是否存在
	var purchaseRecord models.PurchaseRecord
	if err := s.db.Where("purchase_id = ?", info.PurchaseID).First(&purchaseRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("purchase record with ID %d not found", info.PurchaseID)
		}
		return fmt.Errorf("database error checking purchase record: %w", err)
	}

	// 添加日志
	log.Printf("正在创建物流信息：%+v", info)

	return s.db.Create(info).Error
}

// BatchAddLogisticsInfo 批量添加物流信息
func (s *LogisticsService) BatchAddLogisticsInfo(infos []models.LogisticsInfo) error {
	// 添加空值检查
	if len(infos) == 0 {
		return fmt.Errorf("物流信息列表不能为空")
	}

	// 添加日志
	log.Printf("正在批量创建 %d 条物流信息", len(infos))

	return s.db.Create(&infos).Error
}

// GetLogisticsInfoByBatchID 根据批次ID获取物流信息 (如果需要实现)
// func (s *LogisticsService) GetLogisticsInfoByBatchID(batchID string) ([]models.LogisticsInfo, error) {
// 	// ... 实现逻辑 ...
// }
