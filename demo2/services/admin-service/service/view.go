package service

import (
	"errors"
	"fmt"

	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/database/mysql"
)

type Product struct {
	models.FrozenProduct
	models.ProductBatch
}

// 查找冷冻品
func ViewProduct() ([]Product, error) {
	var frozenProducts []models.FrozenProduct

	// 使用 GORM 的关联查询功能
	if err := mysql.DB.Preload("ProductBatches").Find(&frozenProducts).Error; err != nil {
		return nil, fmt.Errorf("数据库查询出错: %w", err)
	}

	// 检查查询结果是否为空
	if len(frozenProducts) == 0 {
		return nil, errors.New("没有找到冷冻品信息")
	}

	var products []Product
	for _, fp := range frozenProducts {
		for _, pb := range fp.ProductBatches {
			products = append(products, Product{
				FrozenProduct: fp,
				ProductBatch:  pb,
			})
		}
	}

	return products, nil
}

// 查看厂家信息
func ViewFactory() ([]models.User, error) {
	var users []models.User

	// 查询所有用户身份为 factory 的用户信息
	if err := mysql.DB.Where("role_type = ?", "factory").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询厂家信息失败: %w", err)
	}

	// 检查查询结果是否为空
	if len(users) == 0 {
		return nil, errors.New("没有找到厂家信息")
	}

	return users, nil
}

// 查看经销商信息
func ViewDistributor() ([]models.User, error) {
	var users []models.User
	//查询所有经销商的用户信息
	if err := mysql.DB.Where("role_type = ?", "distributor").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询经销商信息失败: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("没有找到经销商信息")
	}
	return users, nil

}

// 查看零售商信息
func ViewRetailer() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Where("role_type = ?", "retailer").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询零售商信息失败: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("没有找到零售商信息")
	}
	return users, nil
}

// 查看监管机构信息
func ViewRegulator() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Where("role_type = ?", "regulator").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询监管机构信息失败: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("没有找到监管机构信息")
	}
	return users, nil
}

// 查看消费者信息
func ViewConsumer() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Where("role_type = ?", "consumer").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询消费者信息失败: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("没有找到消费者信息")
	}
	return users, nil
}
func ViewUser() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询用户信息失败: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("没有找到用户信息")
	}
	return users, nil
}

// 查看管理员信息
func ViewAdmin() ([]models.Admin, error) {
	var admins []models.Admin
	if err := mysql.DB.Find(&admins).Error; err != nil {
		return nil, fmt.Errorf("查询管理员信息失败: %w", err)
	}
	if len(admins) == 0 {
		return nil, errors.New("没有找到管理员信息")
	}
	return admins, nil
}

// 查看审核记录表
func ViewAuditLog() ([]models.AuditLog, error) {
	var auditLogs []models.AuditLog
	if err := mysql.DB.Find(&auditLogs).Error; err != nil {
		return nil, fmt.Errorf("查询审核记录表失败: %w", err)
	}
	if len(auditLogs) == 0 {
		return nil, errors.New("没有找到审核记录表信息")
	}
	return auditLogs, nil
}
