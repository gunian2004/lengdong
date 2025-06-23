// 将合约ABI复制到前端目录的脚本
const fs = require("fs");
const path = require("path");

async function main() {
  // 合约名称
  const contractName = "FrozenProduct";

  try {
    // 检查artifacts目录是否存在
    const artifactsDir = path.join(__dirname, "..", "artifacts", "contracts");
    if (!fs.existsSync(artifactsDir)) {
      console.log("正在编译合约...");
      // 如果artifacts目录不存在，先编译合约
      const { execSync } = require("child_process");
      execSync("npx hardhat compile", { stdio: "inherit" });
    }

    // 源文件路径
    const sourcePath = path.join(
      __dirname,
      "..",
      "artifacts",
      "contracts",
      `${contractName}.sol`,
      `${contractName}.json`
    );

    // 检查源文件是否存在
    if (!fs.existsSync(sourcePath)) {
      console.error(`错误: 找不到ABI文件: ${sourcePath}`);
      console.log("请确保合约已编译，并且合约名称正确");
      return;
    }

    // 目标文件夹路径
    const targetDir = path.join(
      __dirname,
      "..",
      "frontend",
      "src",
      "contracts"
    );

    // 目标文件路径
    const targetPath = path.join(targetDir, `${contractName}.json`);

    // 确保目标目录存在
    if (!fs.existsSync(targetDir)) {
      fs.mkdirSync(targetDir, { recursive: true });
    }

    // 读取合约JSON文件
    const contractArtifact = require(sourcePath);

    // 只提取需要的部分（ABI和网络信息）
    const minifiedArtifact = {
      abi: contractArtifact.abi,
      networks: contractArtifact.networks || {},
    };

    // 写入目标文件
    fs.writeFileSync(targetPath, JSON.stringify(minifiedArtifact, null, 2));

    console.log(
      `已成功将 ${contractName} 合约的ABI复制到前端目录: ${targetPath}`
    );

    // 同时复制到Vue前端项目
    const vueTargetDir = path.join(
      __dirname,
      "..",
      "..",
      "front_service_vue",
      "src",
      "utils"
    );
    if (fs.existsSync(path.join(__dirname, "..", "..", "front_service_vue"))) {
      if (!fs.existsSync(vueTargetDir)) {
        fs.mkdirSync(vueTargetDir, { recursive: true });
      }

      // 创建一个简化版本的ABI文件，直接包含在blockchainService.js中
      console.log("正在更新Vue前端的blockchainService.js中的ABI...");

      // 检查blockchainService.js是否存在
      const blockchainServicePath = path.join(
        vueTargetDir,
        "blockchainService.js"
      );
      if (fs.existsSync(blockchainServicePath)) {
        let content = fs.readFileSync(blockchainServicePath, "utf8");

        // 替换ABI部分
        const abiString = JSON.stringify(contractArtifact.abi, null, 2);
        content = content.replace(
          /const FrozenProductABI = \{[\s\S]*?abi: \[([\s\S]*?)\]\s*\};/m,
          `const FrozenProductABI = {\n  abi: ${abiString}\n};`
        );

        fs.writeFileSync(blockchainServicePath, content);
        console.log(`已更新Vue前端的blockchainService.js中的ABI`);
      }
    }
  } catch (error) {
    console.error("复制ABI时出错:", error);
  }
}

// 执行脚本
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
