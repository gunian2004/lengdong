package models

import "time"

//审核记录表
type AuditLog struct {
	AuditID      uint64    `json:"audit_id" gorm:"primaryKey;autoIncrement;comment:审核ID"`
	TargetType   string    `json:"target_type" gorm:"type:enum('user','product');comment:审核类型"`
	TargetID     string    `json:"target_id" gorm:"comment:目标ID"`
	AuditResult  string    `json:"audit_result" gorm:"type:enum('approved','rejected');comment:审核结果"`
	AuditComment string    `json:"audit_comment" gorm:"comment:审核意见"`
	AuditorID    uint64    `json:"auditor_id" gorm:"comment:审核人ID"`
	AuditTime    time.Time `json:"audit_time" gorm:"autoCreateTime;comment:审核时间"`
}
