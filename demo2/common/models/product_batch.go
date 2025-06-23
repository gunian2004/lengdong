package models

import "time"

// 批次表
type ProductBatch struct {
	BatchID         uint64    `json:"batch_id" gorm:"primaryKey;autoIncrement;comment:批次ID"`
	SKU             string    `json:"sku" gorm:"comment:产品SKU;index:idx_sku,uniqueIndex;size:255"`                   // 外键，关联到 FrozenProduct.SKU
	BatchNumber     string    `json:"batch_number" gorm:"comment:批次号;uniqueIndex:idx_batch_number_sku,sku;size:255"` // 批次号和SKU联合唯一
	CurrentHolder   string    `json:"current_holder" gorm:"comment:当前持有者"`
	CurrentLocation string    `json:"current_location" gorm:"comment:当前位置"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Quantity        string    `json:"quantity" gorm:"comment:批次数量"` // 记录该批次的商品数量
}
