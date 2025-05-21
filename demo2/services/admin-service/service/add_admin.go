package service

import (
	"github.com/coldchain-traceability-system/common/models"

	"errors"
	"fmt"

	"github.com/coldchain-traceability-system/database/mysql"
	"golang.org/x/crypto/bcrypt"
)

// AddAdmin 添加管理员
func AddAdmin(username, role, password string) error {
	// 检查用户名和角色是否有效
	if username == "" || role == "" {
		return errors.New("用户名和角色不能为空")
	}

	// 检查角色是否有效
	if role != "super_admin" && role != "normal_admin" {
		return errors.New("无效的角色")
	}

	// 检查用户名是否已存在
	var admin models.Admin
	if err := mysql.DB.Where("username = ?", username).First(&admin).Error; err == nil {
		return errors.New("用户名已存在")
	}
	// 哈希加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码哈希加密失败: %w", err)
	}
	// 创建新管理员
	newAdmin := models.Admin{
		Username: username,
		Role:     role,
		Password: string(hashedPassword),
	}

	if err := mysql.DB.Create(&newAdmin).Error; err != nil {
		return fmt.Errorf("创建管理员失败: %w", err)
	}

	return nil
}
