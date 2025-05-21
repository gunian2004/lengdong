package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONMap is a map type for storing JSON data
type JSONMap map[string]interface{}

// Value implements the driver.Valuer interface for JSONMap
func (m JSONMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan implements the sql.Scanner interface for JSONMap
func (m *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}

// BlockchainRecord represents a record in the blockchain
type BlockchainRecord struct {
	TxHash        string    `json:"tx_hash" gorm:"primaryKey;comment:交易哈希"`
	SKU           string    `json:"sku" gorm:"comment:产品SKU"`
	OperationType string    `json:"operation_type" gorm:"type:enum('product_publish','logistics_update','ownership_transfer');comment:操作类型"`
	Timestamp     time.Time `json:"timestamp" gorm:"autoCreateTime;comment:上链时间"`
	RawData       JSONMap   `json:"raw_data" gorm:"type:json;comment:原始数据"`
} 