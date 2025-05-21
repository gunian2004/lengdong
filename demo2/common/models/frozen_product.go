package models

import "time"

// 厂家数据库结构体
type FrozenProduct struct {
	SKU               string         `json:"sku" gorm:"primaryKey;comment:唯一SKU编码"`
	ProductName       string         `json:"product_name" gorm:"comment:产品名称"`
	Brand             string         `json:"brand" gorm:"comment:品牌"`
	Specification     string         `json:"specification" gorm:"comment:规格/包装"`
	ProductionDate    time.Time      `json:"production_date" gorm:"comment:生产日期"`
	ExpirationDate    time.Time      `json:"expiration_date" gorm:"comment:保质期"`
	RawMaterialSource string         `json:"raw_material_source" gorm:"comment:原材料来源"`
	ProcessingMethod  string         `json:"processing_method" gorm:"comment:加工方式"`
	LogisticsTemp     string         `json:"logistics_temp" gorm:"comment:运输温度要求"`
	StorageCondition  string         `json:"storage_condition" gorm:"comment:仓储条件"`
	QCReport          string         `json:"qc_report" gorm:"comment:质检报告URL"`
	ManufacturerID    string         `json:"manufacturer_id" gorm:"comment:生产商ID"`
	BlockchainTxHash  string         `json:"blockchain_tx_hash" gorm:"comment:区块链交易哈希"`
	AuditStatus       string         `json:"audit_status" gorm:"type:enum('pending','approved','rejected');default:'pending';comment:审核状态"`
	CreatedAt         time.Time      `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	ImageURL          string         `json:"image_url" gorm:"comment:产品图片URL"`
	ProductBatches    []ProductBatch `gorm:"foreignKey:SKU;references:SKU"`
}
