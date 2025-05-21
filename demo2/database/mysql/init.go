package mysql

import (
	"fmt"
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error

	// 设置数据库连接
	dsn := config.AppConfig.Database.DSN()
	fmt.Println("dsn:", dsn)
	// 配置 GORM 日志
	gormConfig := &gorm.Config{}
	if config.AppConfig.Server.Mode == "debug" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	log.Println("数据库连接已建立")

	// 自动迁移数据库模型
	migrateModels()
}

// migrateModels 自动迁移数据库模型
func migrateModels() {
	log.Println("开始数据库迁移")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.FrozenProduct{},
		&models.ProductBatch{},
		&models.PurchaseRecord{},
		&models.LogisticsInfo{},
		&models.BlockchainRecord{},
		&models.AuditLog{},
	)

	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	AddForeignKeyConstraints()
	log.Println("数据库迁移完成")
}
func AddForeignKeyConstraints() error {
	// 检查外键约束是否存在
	var constraintExists bool
	err := DB.Raw("SELECT COUNT(*) > 0 FROM information_schema.KEY_COLUMN_USAGE " +
		"WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'product_batch' " +
		"AND CONSTRAINT_NAME = 'fk_sku'").Row().Scan(&constraintExists)
	if err != nil {
		return fmt.Errorf("检查外键约束是否存在失败: %w", err)
	}

	if constraintExists {
		fmt.Println("外键约束 fk_sku 已存在，跳过添加")
		return nil
	}

	// 添加外键约束
	err = DB.Exec(`
		ALTER TABLE product_batches
		ADD CONSTRAINT fk_sku
		FOREIGN KEY (sku)
		REFERENCES frozen_product(sku)
		ON DELETE CASCADE
	`).Error

	if err != nil {
		return fmt.Errorf("添加外键约束失败: %w", err)
	}

	fmt.Println("外键约束 fk_sku 添加成功")
	return nil
}
