package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// StringSlice is a custom type for handling string slices in GORM
type StringSlice []string

// Value returns the JSON value of the slice for database storage
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan scans JSON data into a StringSlice
func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

// LogisticsInfo represents logistics information for a purchase record
type LogisticsInfo struct {
	LogisticsID           uint64      `json:"logistics_id" gorm:"primaryKey;autoIncrement;comment:物流ID"`
	PurchaseID            uint64      `json:"purchase_id" gorm:"comment:进货记录ID"`
	LogisticsNo           string      `json:"logistics_no" gorm:"index;unique;comment:物流单号"`
	ColdStorageLocation   string      `json:"cold_storage_location" gorm:"comment:冷库位置"`
	Temperature           float64     `json:"temperature" gorm:"comment:温度(℃)"`
	Humidity              *float64    `json:"humidity" gorm:"comment:湿度(%)"`
	LogisticsImages       StringSlice `json:"logistics_images" gorm:"type:json;comment:图片URL(JSON数组)"`
	BlockchainTxHash      string      `json:"blockchain_tx_hash" gorm:"comment:区块链交易哈希"`
	CreatedAt             time.Time   `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
} 