package models

import "time"

// PurchaseRecord represents a purchase record between a buyer and seller
type PurchaseRecord struct {
	PurchaseID   uint64    `json:"purchase_id" gorm:"primaryKey;autoIncrement;comment:进货ID"`
	BuyerID      uint64    `json:"buyer_id" gorm:"comment:买家ID"`
	SellerID     uint64    `json:"seller_id" gorm:"comment:卖家ID"`
	SKU          string    `json:"sku" gorm:"comment:产品SKU"`
	BatchID      uint64    `json:"batch_id" gorm:"comment:批次ID"`
	Quantity     uint      `json:"quantity" gorm:"comment:数量"`
	PurchaseTime time.Time `json:"purchase_time" gorm:"autoCreateTime;comment:进货时间"`
} 