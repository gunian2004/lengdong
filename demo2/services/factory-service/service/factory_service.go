package service

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/coldchain-traceability-system/common/ethereum"
	"github.com/coldchain-traceability-system/common/models"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/mozillazg/go-pinyin"
)

// 厂家数据绑定结构体
type FrozenProductBinding struct {
	ProductName       string    `json:"product_name" gorm:"comment:产品名称" binding:"required"`
	Brand             string    `json:"brand" gorm:"comment:品牌" binding:"required"`
	Specification     string    `json:"specification" gorm:"comment:规格/包装" binding:"required"`
	ProductionDate    time.Time `json:"production_date" gorm:"comment:生产日期" binding:"required"`
	ExpirationDate    time.Time `json:"expiration_date" gorm:"comment:保质期" binding:"required"`
	RawMaterialSource string    `json:"raw_material_source" gorm:"comment:原材料来源" binding:"required"`
	ProcessingMethod  string    `json:"processing_method" gorm:"comment:加工方式" binding:"required"`
	LogisticsTemp     string    `json:"logistics_temp" gorm:"comment:运输温度要求" binding:"required"`
	StorageCondition  string    `json:"storage_condition" gorm:"comment:仓储条件" binding:"required"`
	QCReport          string    `json:"qc_report" gorm:"comment:质检报告URL" binding:"required"`
	ManufacturerID    string    `json:"manufacturer_id" gorm:"comment:生产商ID" binding:"required"`
	ImageURL          string    `json:"image_url" gorm:"comment:产品图片URL" binding:"required"`
}

// 批次表绑定
type ProductBatchBinding struct {
	CurrentHolder   string `json:"current_holder" gorm:"comment:当前持有者" binding:"required"`
	CurrentLocation string `json:"current_location" gorm:"comment:当前位置" binding:"required"`
	Quantity        string `json:"quantity" gorm:"comment:批次数量" binding:"required"` // 记录该批次的商品数量
}

func getInitial(s string) string {
	// 去除首尾空格和引号
	s = strings.TrimSpace(strings.Trim(s, "\""))

	if len(s) == 0 {
		return "X" // 默认返回 "X" 避免空字符串
	}

	// 尝试将整个字符串转换为拼音首字母
	args := pinyin.NewArgs()
	args.Heteronym = true // 启用多音字模式
	pinyinSlice := pinyin.Pinyin(s, args)

	var initials []rune
	for _, pinyinChars := range pinyinSlice {
		if len(pinyinChars) > 0 {
			initials = append(initials, rune(pinyinChars[0][0]))
		}
	}

	// 如果没有提取到首字母，返回默认值 "X"
	if len(initials) == 0 {
		fmt.Printf("无法解析首字母: %q, 使用默认值 X\n", s)
		return "X" // 默认返回 "X"
	}

	// 将首字母拼接成字符串并转换为大写
	initialStr := strings.ToUpper(string(initials))

	return initialStr
}

// generateSKU 生成SKU码
func GenerateSKU(product FrozenProductBinding, batchNumber string) string {
	// 获取品牌首字母
	brandInitial := getInitial(product.Brand)
	if brandInitial == "" {
		fmt.Println("品牌首字母为空，使用默认值 X")
		brandInitial = "X" // 默认值
	}

	// 获取产品名称首字母
	productNameInitial := getInitial(product.ProductName)
	if productNameInitial == "" {
		fmt.Println("产品名称首字母为空，使用默认值 X")
		productNameInitial = "X" // 默认值
	}

	// 获取生产日期
	productionDate := product.ProductionDate.Format("20060102")
	if productionDate == "" {
		fmt.Println("生产日期为空，使用默认值 00000000")
		productionDate = "00000000" // 默认值
	}

	// 获取原材料来源首字母
	rawMaterialInitial := getInitial(product.RawMaterialSource)
	if rawMaterialInitial == "" {
		fmt.Println("原材料来源首字母为空，使用默认值 X")
		rawMaterialInitial = "X" // 默认值
	}

	// 获取生产商ID
	manufacturerID := product.ManufacturerID

	// 生成SKU码
	sku := fmt.Sprintf("%s%s%s%s%s",
		brandInitial,
		productNameInitial,
		productionDate,
		rawMaterialInitial,
		manufacturerID,
	)

	// 检查SKU是否唯一
	var existingProduct models.FrozenProduct
	if err := mysql.DB.Where("sku = ?", sku).First(&existingProduct).Error; err == nil {
		fmt.Printf("Duplicate SKU detected: %s, retrying...\n", sku) // 添加日志
		return GenerateSKU(product, batchNumber)                     // 递归生成新的SKU码
	}

	return sku
}

// 批次号获取
func generateBatchNumber() string {
	timestamp := time.Now().UnixNano() // 获取当前时间的纳秒级时间戳
	return fmt.Sprintf("Batch%d", timestamp)
}

// AddFrozenProduct 添加冷冻品
func AddFrozenProduct(product FrozenProductBinding, batch ProductBatchBinding) error {
	tx := mysql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 生成批次号
	batchNumber := generateBatchNumber()

	// 生成SKU码
	sku := GenerateSKU(product, batchNumber)
	fmt.Sprintln(sku)

	// 创建FrozenProduct
	frozenProduct := models.FrozenProduct{
		SKU:               sku,
		ProductName:       product.ProductName,
		Brand:             product.Brand,
		Specification:     product.Specification,
		ProductionDate:    product.ProductionDate,
		ExpirationDate:    product.ExpirationDate,
		RawMaterialSource: product.RawMaterialSource,
		ProcessingMethod:  product.ProcessingMethod,
		LogisticsTemp:     product.LogisticsTemp,
		StorageCondition:  product.StorageCondition,
		QCReport:          product.QCReport,
		ManufacturerID:    product.ManufacturerID,
		ImageURL:          "http://localhost:8082/static/Images/" + product.ImageURL,
		AuditStatus:       "pending", // 默认审核状态
	}
	if err := tx.Create(&frozenProduct).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建ProductBatch
	productBatch := models.ProductBatch{
		SKU:             frozenProduct.SKU,
		BatchNumber:     batchNumber, // 使用相同的批次号
		CurrentHolder:   batch.CurrentHolder,
		CurrentLocation: batch.CurrentLocation,
		Quantity:        batch.Quantity,
	}
	if err := tx.Create(&productBatch).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

type Product struct {
	models.FrozenProduct
	models.ProductBatch
}

func ViewProduct(manufacturerID string) ([]Product, error) {
	if manufacturerID == "" {
		return nil, errors.New("生产商ID不能为空")
	}

	var frozenProducts []models.FrozenProduct

	// 查询指定 ManufacturerID 的产品，并预加载其批次信息
	if err := mysql.DB.
		Where("manufacturer_id = ?", manufacturerID).
		Preload("ProductBatches").
		Find(&frozenProducts).Error; err != nil {
		return nil, fmt.Errorf("数据库查询出错: %w", err)
	}

	// 检查查询结果是否为空
	if len(frozenProducts) == 0 {
		return nil, errors.New("没有找到该生产商相关的冷冻品信息")
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

// SyncToBlockchain 将审核通过的产品信息同步到以太坊区块链

func SyncToBlockchain(sku string) error {
	var product models.FrozenProduct

	// 查询指定 SKU 的产品
	if err := mysql.DB.Where("sku = ?", sku).Preload("ProductBatches").First(&product).Error; err != nil {
		log.Printf("未找到对应产品: %v", err)
		return err
	}

	// 检查审核状态
	if product.AuditStatus != "approved" {
		log.Printf("产品 %s 审核状态不是 approved，当前状态: %s", sku, product.AuditStatus)
		return errors.New("只有审核通过的产品才能上链")
	}

	// 确保至少有一个批次
	if len(product.ProductBatches) == 0 {
		log.Printf("产品 %s 没有对应的批次信息", sku)
		return errors.New("没有可上链的批次信息")
	}

	// 获取第一个批次
	batch := product.ProductBatches[0]

	// 转换时间为 big.Int
	productionDate := big.NewInt(product.ProductionDate.Unix())
	expirationDate := big.NewInt(product.ExpirationDate.Unix())
	quantity := big.NewInt(0)
	if _, ok := quantity.SetString(batch.Quantity, 10); !ok {
		return errors.New("批次数量转换失败")
	}

	log.Println("开始上链操作，SKU:", sku)

	// ✅ 调用以太坊合约 AddProduct 方法（注意参数顺序）
	txHash, err := ethereum.Contract.AddProduct(
		ethereum.Auth,
		product.SKU,
		product.ProductName,
		product.Brand,
		product.Specification,
		productionDate,
		expirationDate,
		product.RawMaterialSource,
		product.ProcessingMethod,
		product.LogisticsTemp,
		product.StorageCondition,
		product.QCReport,
		product.ManufacturerID,
		product.ImageURL,
		batch.BatchNumber,
		quantity,
		batch.CurrentLocation,
	)
	if err != nil {
		log.Printf("发送冷冻品信息到以太坊合约失败: %v", err)
		return err
	}

	log.Printf("以太坊交易哈希: %s", txHash.Hash().Hex())

	// 更新数据库中的以太坊交易哈希
	product.BlockchainTxHash = txHash.Hash().Hex()
	if err := mysql.DB.Save(&product).Error; err != nil {
		log.Printf("更新数据库中的以太坊交易哈希失败: %v", err)
		return err
	}

	log.Printf("产品 %s 成功上链", sku)
	return nil
}
