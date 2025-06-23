# admin_service_vue

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

## 区块链集成

本项目已集成区块链功能，可以从区块链网络获取冷冻品信息。

### 集成步骤

1. 安装依赖：

   ```bash
   npm install ethers@6.14.3
   ```

2. 配置区块链连接：

   - 在 `src/utils/blockchainService.js` 中配置合约地址
   - 默认连接到本地 Hardhat 网络 (http://127.0.0.1:8545/)
   - 需要 MetaMask 钱包支持

3. 使用方法：
   - 在消费者页面可以切换数据源（API/区块链）
   - 选择"区块链数据"后，点击"连接区块链"按钮
   - 授权 MetaMask 连接后即可查看区块链上的冷冻品信息

### 区块链合约

区块链合约位于 `frozen-chain-project` 目录，包括：

- `FrozenProduct.sol`: 冷冻品信息智能合约
- 提供添加、更新、查询冷冻品信息的功能
- 所有数据存储在区块链上，确保数据不可篡改

### 本地开发

1. 启动本地区块链：

   ```bash
   cd ../frozen-chain-project
   npx hardhat node
   ```

2. 部署合约：

   ```bash
   # 在新的终端窗口中运行
   cd frozen-chain-project
   npx hardhat run scripts/deploy.js --network localhost
   ```

3. 添加测试数据：

   ```bash
   # 部署后运行
   npx hardhat run scripts/add-test-data.js --network localhost
   ```

4. 复制合约地址到 `blockchainService.js` 中的 `contractAddress` 变量

   - 当前部署地址为: `0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512`

5. 启动前端应用：
   ```bash
   npm run dev
   ```

### 使用 MetaMask 连接

1. 安装 MetaMask 浏览器扩展
2. 添加本地网络：
   - 网络名称：Hardhat 本地
   - RPC URL：http://127.0.0.1:8545/
   - 链 ID：31337
   - 货币符号：ETH
3. 导入 Hardhat 提供的私钥以获取测试账户
   - 测试账户私钥: `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`
