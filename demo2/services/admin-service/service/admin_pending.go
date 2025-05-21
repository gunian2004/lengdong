package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/coldchain-traceability-system/common/models"

	"github.com/coldchain-traceability-system/database/mysql"
	"gorm.io/gorm"
)

// UpdateUserPending 更新用户审核状态
func UpdateUserPending(username, issuccess, auditComment string, auditorID uint64) error {
	// 去除空格和双引号
	username = strings.TrimSpace(strings.Trim(username, `"`))
	issuccess = strings.TrimSpace(strings.Trim(issuccess, `"`))
	auditComment = strings.TrimSpace(strings.Trim(auditComment, `"`))

	// 验证输入数据
	if username == "" {
		return errors.New("用户名不能为空")
	}
	if issuccess == "" {
		return errors.New("审核结果不能为空")
	}

	// 查找用户
	var user models.User
	if err := mysql.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("用户不存在")
		}
		return fmt.Errorf("查找用户失败: %w", err)
	}

	// 根据 issuccess 更新审核状态
	switch issuccess {
	case "true", "1", "approved":
		user.AuditStatus = "approved"
	case "false", "0", "rejected":
		user.AuditStatus = "rejected"
	default:
		return errors.New("无效的审核结果")
	}

	if err := mysql.DB.Save(&user).Error; err != nil {
		return fmt.Errorf("更新用户审核状态失败: %w", err)
	}

	// 记录审核日志
	auditLog := models.AuditLog{
		TargetType:   "user",
		TargetID:     fmt.Sprintf("%d", user.ID),
		AuditResult:  user.AuditStatus,
		AuditComment: auditComment,
		AuditorID:    auditorID,
		AuditTime:    time.Now(),
	}

	if err := mysql.DB.Create(&auditLog).Error; err != nil {
		return fmt.Errorf("记录审核日志失败: %w", err)
	}

	return nil
}

// UpdateProductPending 更新商品审核状态
func UpdateProductPending(productName, issuccess, auditComment string, auditorID uint64) error {
	// 去除空格和双引号
	productName = strings.TrimSpace(strings.Trim(productName, `"`))
	issuccess = strings.TrimSpace(strings.Trim(issuccess, `"`))
	auditComment = strings.TrimSpace(strings.Trim(auditComment, `"`))

	// 验证输入数据
	if productName == "" {
		return errors.New("商品名字不能为空")
	}
	if issuccess == "" {
		return errors.New("审核结果不能为空")
	}

	// 查找商品
	var product models.FrozenProduct
	if err := mysql.DB.Where("product_name = ?", productName).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("商品不存在")
		}
		return fmt.Errorf("查找商品失败: %w", err)
	}

	// 根据 issuccess 更新审核状态
	switch issuccess {
	case "true", "1", "approved":
		product.AuditStatus = "approved"
	case "false", "0", "rejected":
		product.AuditStatus = "rejected"
	default:
		return errors.New("无效的审核结果")
	}
	fmt.Println("审核结果:", product.AuditStatus)
	if err := mysql.DB.Save(&product).Error; err != nil {
		return fmt.Errorf("更新商品审核状态失败: %w", err)
	}

	// 记录审核日志
	auditLog := models.AuditLog{
		TargetType:   "product",
		TargetID:     product.SKU,
		AuditResult:  product.AuditStatus,
		AuditComment: auditComment,
		AuditorID:    auditorID,
		AuditTime:    time.Now(),
	}

	if err := mysql.DB.Create(&auditLog).Error; err != nil {
		return fmt.Errorf("记录审核日志失败: %w", err)
	}

	return nil
}
