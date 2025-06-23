package models

type User struct {
	GSModel
	ID          uint64 `json:"id" gorm:"primaryKey;autoIncrement;comment:用户ID"`
	Username    string `json:"username" gorm:"index;unique;comment:用户名"`
	Password    string `json:"-" gorm:"comment:加密密码"`
	RoleType    string `json:"role_type" gorm:"type:enum('factory','distributor','retailer','regulator','consumer');comment:角色类型"`
	ContactInfo string `json:"contact_info" gorm:"comment:联系方式"`
	Address     string `json:"address" gorm:"comment:地址"`
	AuditStatus string `json:"audit_status" gorm:"type:enum('pending','approved','rejected');default:'pending';comment:审核状态"`
}
