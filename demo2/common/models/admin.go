package models

type Admin struct {
	AdminID  uint64 `json:"admin_id" gorm:"primaryKey;autoIncrement;comment:管理员ID"`
	Username string `json:"username" gorm:"index;comment:账号"`
	Password string `json:"-" gorm:"comment:加密密码"`
	Role     string `json:"role" gorm:"type:enum('super_admin','normal_admin');default:'normal_admin';comment:角色"`
}
