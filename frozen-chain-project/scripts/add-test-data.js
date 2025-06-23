// 添加测试数据的脚本
const { ethers } = require("hardhat");

async function main() {
  console.log("开始添加测试数据...");

  // 获取部署的合约
  const contractAddress = "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"; // 使用实际部署的地址
  const FrozenProduct = await ethers.getContractFactory("FrozenProduct");
  const frozenProduct = await FrozenProduct.attach(contractAddress);

  // 准备测试数据
  const testProducts = [
    {
      id: "SKU001",
      name: "冰鲜三文鱼",
      category: "海鲜",
      manufacturer: "海洋食品公司",
      productionDate: Math.floor(Date.now() / 1000) - 7 * 24 * 60 * 60, // 7天前
      expiryDate: Math.floor(Date.now() / 1000) + 30 * 24 * 60 * 60, // 30天后
      storageTemp: -18,
      batchNumber: "B20230501",
      description: "新鲜冰鲜三文鱼，适合刺身和烹饪",
    },
    {
      id: "SKU002",
      name: "冷冻牛排",
      category: "肉类",
      manufacturer: "优质肉类供应商",
      productionDate: Math.floor(Date.now() / 1000) - 15 * 24 * 60 * 60, // 15天前
      expiryDate: Math.floor(Date.now() / 1000) + 180 * 24 * 60 * 60, // 180天后
      storageTemp: -20,
      batchNumber: "B20230502",
      description: "优质安格斯牛肉，真空包装",
    },
    {
      id: "SKU003",
      name: "冷冻虾仁",
      category: "海鲜",
      manufacturer: "海洋食品公司",
      productionDate: Math.floor(Date.now() / 1000) - 10 * 24 * 60 * 60, // 10天前
      expiryDate: Math.floor(Date.now() / 1000) + 90 * 24 * 60 * 60, // 90天后
      storageTemp: -18,
      batchNumber: "B20230503",
      description: "去壳虾仁，适合炒菜和煮汤",
    },
  ];

  // 添加测试数据
  for (const product of testProducts) {
    console.log(`添加产品: ${product.name} (${product.id})`);
    try {
      const tx = await frozenProduct.addProduct(
        product.id,
        product.name,
        product.category,
        product.manufacturer,
        product.productionDate,
        product.expiryDate,
        product.storageTemp,
        product.batchNumber,
        product.description
      );
      await tx.wait();
      console.log(`- 产品 ${product.id} 添加成功`);
    } catch (error) {
      // 如果产品已存在，尝试更新
      if (error.message.includes("Product ID already exists")) {
        console.log(`- 产品 ${product.id} 已存在，尝试更新...`);
        try {
          const tx = await frozenProduct.updateProduct(
            product.id,
            product.name,
            product.category,
            product.manufacturer,
            product.productionDate,
            product.expiryDate,
            product.storageTemp,
            product.batchNumber,
            product.description
          );
          await tx.wait();
          console.log(`- 产品 ${product.id} 更新成功`);
        } catch (updateError) {
          console.error(`- 更新产品 ${product.id} 失败:`, updateError.message);
        }
      } else {
        console.error(`- 添加产品 ${product.id} 失败:`, error.message);
      }
    }
  }

  // 验证数据
  console.log("\n验证产品数据:");
  const count = await frozenProduct.getProductCount();
  console.log(`总产品数量: ${count}`);

  const ids = await frozenProduct.getAllProductIds();
  console.log(`所有产品ID: ${ids.join(", ")}`);

  console.log("\n测试数据添加完成！");
}

// 执行脚本
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("添加测试数据时出错:", error);
    process.exit(1);
  });
