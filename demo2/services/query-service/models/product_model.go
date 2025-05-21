package models

import "time"

// 冷冻品链路信息
type ProductChain struct {
	SKU         string       `json:"sku" gorm:"primaryKey;comment:唯一SKU编码"`
	ProductName string       `json:"product_name" gorm:"comment:产品名称"`
	Brand       string       `json:"brand" gorm:"comment:品牌"`
	Batches     []BatchInfo  `json:"batches" gorm:"comment:批次信息"`
	Logistics   []Logistics  `json:"logistics" gorm:"comment:物流信息"`
	Blockchain  []Blockchain `json:"blockchain" gorm:"comment:区块链记录"`
}

// 批次信息
type BatchInfo struct {
	BatchNumber     string `json:"batch_number" gorm:"comment:批次号"`
	CurrentHolder   string `json:"current_holder" gorm:"type:enum('factory','distributor','retailer');comment:当前持有者"`
	CurrentLocation string `json:"current_location" gorm:"comment:当前位置"`
	CreatedAt       string `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}

// 物流信息
type Logistics struct {
	LogisticsNo         string   `json:"logistics_no" gorm:"index;unique;comment:物流单号"`
	ColdStorageLocation string   `json:"cold_storage_location" gorm:"comment:冷库位置"`
	Temperature         float64  `json:"temperature" gorm:"comment:温度(℃)"`
	Humidity            *float64 `json:"humidity,omitempty" gorm:"comment:湿度(%)"`
	LogisticsImages     []string `json:"logistics_images" gorm:"comment:图片URL(JSON数组)"`
	CreatedAt           string   `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}

// 区块链记录
type Blockchain struct {
	TxHash        string                 `json:"tx_hash" gorm:"primaryKey;comment:交易哈希"`
	OperationType string                 `json:"operation_type" gorm:"type:enum('product_publish','logistics_update','ownership_transfer');comment:操作类型"`
	Timestamp     string                 `json:"timestamp" gorm:"autoCreateTime;comment:上链时间"`
	RawData       map[string]interface{} `json:"raw_data" gorm:"comment:原始数据"`
}

// FrozenProduct 冻品信息表的数据模型
type FrozenProduct struct {
	SKU               string    `json:"sku" gorm:"primaryKey;comment:唯一SKU编码"`
	ProductName       string    `json:"product_name" gorm:"comment:产品名称"`
	Brand             string    `json:"brand" gorm:"comment:品牌"`
	Specification     string    `json:"specification" gorm:"comment:规格/包装"`
	ProductionDate    time.Time `json:"production_date" gorm:"comment:生产日期"`
	ExpirationDate    time.Time `json:"expiration_date" gorm:"comment:保质期"`
	RawMaterialSource string    `json:"raw_material_source" gorm:"comment:原材料来源"`
	ProcessingMethod  string    `json:"processing_method" gorm:"comment:加工方式"`
	LogisticsTemp     string    `json:"logistics_temp" gorm:"comment:运输温度要求"`
	StorageCondition  string    `json:"storage_condition" gorm:"comment:仓储条件"`
	QCReport          string    `json:"qc_report" gorm:"comment:质检报告URL"`
	ManufacturerID    uint64    `json:"manufacturer_id" gorm:"comment:生产商ID"`
	BlockchainTxHash  string    `json:"blockchain_tx_hash" gorm:"comment:区块链交易哈希"`
	AuditStatus       string    `json:"audit_status" gorm:"type:enum('pending','published','offline');default:'pending';comment:审核状态"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}

// ProductBatch 表示批次表的数据模型
type ProductBatch struct {
	BatchID         uint64    `json:"batch_id" gorm:"primaryKey;autoIncrement;comment:批次ID"`
	SKU             string    `json:"sku" gorm:"comment:产品SKU"`
	BatchNumber     string    `json:"batch_number" gorm:"comment:批次号"`
	CurrentHolder   string    `json:"current_holder" gorm:"type:enum('factory','distributor','retailer');comment:当前持有者"`
	CurrentLocation string    `json:"current_location" gorm:"comment:当前位置"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}

// LogisticsInfo 表示物流信息表的数据模型
type LogisticsInfo struct {
	LogisticsID         uint64    `json:"logistics_id" gorm:"primaryKey;autoIncrement;comment:物流ID"`
	PurchaseID          uint64    `json:"purchase_id" gorm:"comment:进货记录ID"`
	LogisticsNo         string    `json:"logistics_no" gorm:"index;unique;comment:物流单号"`
	ColdStorageLocation string    `json:"cold_storage_location" gorm:"comment:冷库位置"`
	Temperature         float64   `json:"temperature" gorm:"comment:温度(℃)"`
	Humidity            *float64  `json:"humidity,omitempty" gorm:"comment:湿度(%)"`
	LogisticsImages     []string  `json:"logistics_images" gorm:"comment:图片URL(JSON数组)"`
	BlockchainTxHash    string    `json:"blockchain_tx_hash" gorm:"comment:区块链交易哈希"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}

// BlockchainRecord 表示区块链记录表的数据模型
type BlockchainRecord struct {
	TxHash        string                 `json:"tx_hash" gorm:"primaryKey;comment:交易哈希"`
	SKU           string                 `json:"sku" gorm:"comment:产品SKU"`
	OperationType string                 `json:"operation_type" gorm:"type:enum('product_publish','logistics_update','ownership_transfer');comment:操作类型"`
	Timestamp     time.Time              `json:"timestamp" gorm:"autoCreateTime;comment:上链时间"`
	RawData       map[string]interface{} `json:"raw_data" gorm:"comment:原始数据"`
}
