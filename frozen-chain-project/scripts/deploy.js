// 部署FrozenProduct合约的脚本
const hre = require("hardhat");

async function main() {
  console.log("开始部署FrozenProduct合约...");

  // 获取合约工厂
  const FrozenProduct = await hre.ethers.getContractFactory("FrozenProduct");

  // 部署合约
  const frozenProduct = await FrozenProduct.deploy();

  // 等待部署完成
  await frozenProduct.waitForDeployment();

  // 获取合约地址
  const frozenProductAddress = await frozenProduct.getAddress();

  console.log(`FrozenProduct合约已部署到地址: ${frozenProductAddress}`);
}

// 执行部署
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
