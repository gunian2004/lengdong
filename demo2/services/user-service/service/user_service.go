package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/common/utils"
	"github.com/coldchain-traceability-system/database/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 提供用户操作相关的方法
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建一个新的用户服务实例
func NewUserService() *UserService {
	return &UserService{
		db: mysql.DB,
	}
}

// CreateUser 创建一个新用户，注册一个用户
func CreateUser(username, password, roleType, contactInfo, address string) error {
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
	// 检查用户名 + 角色类型是否已存在
	var user models.User
	if err := mysql.DB.Where("username = ? AND role_type = ?", username, roleType).First(&user).Error; err == nil {
		return errors.New("该用户名和角色类型组合已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果不是“记录未找到”错误，则返回数据库错误
		return fmt.Errorf("检查用户是否存在时出错: %w", err)
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

// GetUserByID 根据 ID 获取用户
func (s *UserService) GetUserByID(id uint64) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("未找到该用户")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id uint64, updates map[string]interface{}) error {
	// 检查用户是否存在
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("未找到该用户")
		}
		return err
	}

	// 检查更新内容中是否包含密码
	if password, ok := updates["password"].(string); ok {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		updates["password"] = hashedPassword
	}

	// 执行更新
	return s.db.Model(&user).Updates(updates).Error
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint64) error {
	result := s.db.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("未找到该用户")
	}
	return result.Error
}

// Login 用户登录并返回 JWT token
func Login(username, password string) (string, *models.User, error) {
	var user models.User
	var admin models.Admin

	// 1. 尝试从普通用户表中查找
	if err := mysql.DB.Where("username = ?", username).First(&user).Error; err == nil {
		// 普通用户存在

		// 检查审核状态
		if user.AuditStatus != "approved" {
			return "", nil, errors.New("用户账户尚未通过审核")
		}

		// 验证密码
		if !utils.CheckPasswordHash(password, user.Password) {
			return "", nil, errors.New("用户名或密码错误")
		}

		// 生成 JWT token（普通用户）
		token, err := middleware.GenerateToken(user.ID, user.Username, user.RoleType)
		if err != nil {
			log.Printf("生成 token 失败: %v", err)
			return "", nil, errors.New("无法生成认证 token")
		}

		return token, &user, nil
	}

	// 2. 如果普通用户不存在，则尝试从管理员表中查找
	if err := mysql.DB.Where("username = ?", username).First(&admin).Error; err == nil {
		// 管理员存在

		// 验证密码
		if !utils.CheckPasswordHash(password, admin.Password) {
			return "", nil, errors.New("管理员账号或密码错误")
		}

		// 生成 JWT token（管理员）
		token, err := middleware.GenerateToken(uint64(admin.AdminID), admin.Username, "admin")
		if err != nil {
			log.Printf("生成 token 失败: %v", err)
			return "", nil, errors.New("无法生成认证 token")
		}

		// 构造一个通用的 User 对象用于返回
		adminUser := models.User{
			ID:       uint64(admin.AdminID),
			Username: admin.Username,
			RoleType: admin.Role,
		}

		return token, &adminUser, nil
	}

	// 3. 用户和管理员都未找到
	return "", nil, errors.New("账号或密码错误")
}

// ListUsers 分页列出用户
func (s *UserService) ListUsers(page, pageSize int, filters map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// 应用筛选条件
	query := s.db.Model(&models.User{})
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
