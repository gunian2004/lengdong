package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// StringSlice 是一个自定义类型，用于在 GORM 中处理字符串切片
type StringSlice []string

// Value 返回该切片的 JSON 值，用于数据库存储
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan 从数据库中扫描 JSON 数据到 StringSlice
func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型断言失败：期望 []byte")
	}
	return json.Unmarshal(bytes, s)
}

// LogisticsInfo 表示进货记录的物流信息
type LogisticsInfo struct {
	LogisticsID         uint64      `json:"logistics_id" gorm:"primaryKey;autoIncrement;comment:物流ID"`
	PurchaseID          uint64      `json:"purchase_id" gorm:"comment:进货记录ID"`
	LogisticsNo         string      `json:"logistics_no" gorm:"index;unique;comment:物流单号"`
	ColdStorageLocation string      `json:"cold_storage_location" gorm:"comment:冷库位置"`
	Temperature         float64     `json:"temperature" gorm:"comment:温度(℃)"`
	Humidity            *float64    `json:"humidity" gorm:"comment:湿度(%)"`
	LogisticsImages     StringSlice `json:"logistics_images" gorm:"type:json;comment:图片URL(JSON数组)"`
	BlockchainTxHash    string      `json:"blockchain_tx_hash" gorm:"comment:区块链交易哈希"`
	CreatedAt           time.Time   `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
}
