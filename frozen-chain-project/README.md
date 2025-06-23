# 冷冻品信息上链系统

这是一个基于区块链技术的冷冻品信息管理系统，使用 Hardhat 和 React 开发。

## 项目结构

```
frozen-chain-project/
├── contracts/         # 智能合约
├── scripts/           # 部署和辅助脚本
├── test/              # 测试文件
├── frontend/          # 前端应用
└── hardhat.config.js  # Hardhat配置
```

## 安装依赖

```bash
# 安装后端依赖
npm install

# 安装前端依赖
cd frontend
npm install
cd ..
```

## 编译智能合约

```bash
npx hardhat compile
```

## 运行本地区块链

```bash
npx hardhat node
```

## 部署智能合约

在新的终端窗口中运行：

```bash
npx hardhat run scripts/deploy.js --network localhost
```

部署后，记下合约地址，并更新前端代码中的合约地址。

## 复制 ABI 到前端

```bash
node scripts/copy-abi.js
```

## 启动前端应用

```bash
cd frontend
npm start
```

## 使用 MetaMask 连接

1. 安装 MetaMask 浏览器扩展
2. 创建或导入账户
3. 添加本地网络：
   - 网络名称：Hardhat 本地
   - RPC URL：http://127.0.0.1:8545/
   - 链 ID：1337
   - 货币符号：ETH
4. 导入 Hardhat 提供的私钥以获取测试账户

## 功能说明

- 管理员（合约部署者）可以添加和更新冷冻品信息
- 任何用户都可以查询冷冻品信息
- 所有冷冻品信息都存储在区块链上，确保数据的不可篡改性和透明性
