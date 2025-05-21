package service

import (
	"errors"
	"fmt"

	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/database/mysql"

	"gorm.io/gorm"
)

func UpdateRole(username, role string) error {
	// 检查用户名和角色是否有效
	if username == "" || role == "" {
		return errors.New("用户名和角色不能为空")
	}

	// 检查角色是否有效
	if role != "super_admin" && role != "normal_admin" {
		return errors.New("无效的角色")
	}

	// 查找现有管理员
	var admin models.Admin
	if err := mysql.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("管理员不存在")
		}
		return fmt.Errorf("查找管理员失败: %w", err)
	}

	// 更新管理员角色
	admin.Role = role
	if err := mysql.DB.Save(&admin).Error; err != nil {
		return fmt.Errorf("更新管理员角色失败: %w", err)
	}

	return nil
}
