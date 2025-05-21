package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/coldchain-traceability-system/common/models"

	"github.com/coldchain-traceability-system/database/mysql"
	"golang.org/x/crypto/bcrypt"
)

// AddUser 添加用户
func AddUser(username, password, roleType, contactInfo, address string) error {
	// 去除空格和双引号
	username = strings.TrimSpace(strings.Trim(username, `"`))
	password = strings.TrimSpace(strings.Trim(password, `"`))
	roleType = strings.TrimSpace(strings.Trim(roleType, `"`))
	contactInfo = strings.TrimSpace(strings.Trim(contactInfo, `"`))
	address = strings.TrimSpace(strings.Trim(address, `"`))

	// 验证输入数据
	if username == "" || password == "" || roleType == "" {
		return errors.New("用户名、密码和角色类型不能为空")
	}

	// 检查角色类型是否有效
	if roleType != "factory" && roleType != "distributor" && roleType != "retailer" && roleType != "regulator" && roleType != "consumer" {
		return errors.New("无效的角色类型")
	}

	// 检查用户名是否已存在
	var user models.User
	if err := mysql.DB.Where("username = ?", username).First(&user).Error; err == nil {
		return errors.New("用户名已存在")
	}

	// 哈希加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码哈希加密失败: %w", err)
	}

	// 创建新用户
	newUser := models.User{
		Username:    username,
		Password:    string(hashedPassword),
		RoleType:    roleType,
		ContactInfo: contactInfo,
		Address:     address,
		AuditStatus: "pending", // 默认审核状态为 pending
	}

	if err := mysql.DB.Create(&newUser).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	return nil
}
