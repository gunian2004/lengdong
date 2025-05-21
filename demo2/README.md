# 冷链溯源系统

一个基于微服务的系统，用于跟踪和追溯冷冻食品产品在整个供应链中的流转情况，使用区块链技术确保数据的完整性和不可篡改性。

## 系统架构

该系统由以下微服务组成：

1. **用户服务**：处理用户注册、认证和资料管理。
2. **工厂服务**：管理工厂用户的商品创建、更新和批次管理。
3. **分销商服务**：管理分销商用户的产品采购和物流。
4. **零售商服务**：管理零售商用户的产品接收和销售。
5. **查询服务**：提供查询产品信息和溯源数据的 API。
6. **区块链服务**：与 Hyperledger Fabric 区块链网络交互。
7. **管理员服务**：提供系统的管理功能。

## 技术栈

- **后端**：Go (Golang)
- **Web 框架**：Gin
- **数据库**：MySQL
- **缓存**：Redis
- **区块链**：Hyperledger Fabric
- **认证**：JWT

## 前置条件

- Go 1.19+
- MySQL 8.0+
- Redis 6.0+
- Hyperledger Fabric 2.x

## 入门指南

### 1. 克隆仓库

```bash
git clone https://github.com/your-username/coldchain-traceability-system.git
cd coldchain-traceability-system
```

### 2. 配置环境

复制示例配置文件并更新为自己的设置：

```bash
cp config.example.yaml config.yaml
```

编辑 `config.yaml` 以匹配您的环境。

### 3. 设置数据库

创建一个 MySQL 数据库：

```sql
CREATE DATABASE coldchain_db;
```

应用程序将在启动时自动创建所需的表。

### 4. 启动服务

可以分别启动每个服务：

```bash
# 启动用户服务
cd services/user-service
go run main.go

# 启动工厂服务
cd services/factory-service
go run main.go

# 启动查询服务
cd services/query-service
go run main.go

# 启动区块链服务
cd services/blockchain-service
go run main.go
```

## API 文档

### 用户服务 API

- `POST /api/v1/users/register`: 注册新用户
- `POST /api/v1/users/login`: 用户登录
- `GET /api/v1/users/profile`: 获取当前用户资料
- `PUT /api/v1/users/profile`: 更新用户资料

### 工厂服务 API

- `POST /api/v1/factory/products`: 创建新产品
- `PUT /api/v1/factory/products/:sku`: 更新产品
- `GET /api/v1/factory/products/:sku`: 获取产品
- `GET /api/v1/factory/products`: 列出产品
- `POST /api/v1/factory/batches`: 创建新批次
- `GET /api/v1/factory/batches`: 列出批次

### 查询服务 API

- `GET /api/v1/query/products/:sku`: 根据 SKU 获取产品
- `GET /api/v1/query/products/search`: 搜索产品
- `GET /api/v1/query/products/:sku/blockchain`: 从区块链获取产品
- `GET /api/v1/query/products/:sku/traceability`: 获取产品的溯源数据
- `GET /api/v1/query/products/:sku/batches`: 根据 SKU 获取批次
- `GET /api/v1/query/logistics/:id`: 获取物流信息

### 区块链服务 API

- `POST /api/v1/blockchain/products`: 将产品添加到区块链
- `POST /api/v1/blockchain/logistics`: 更新区块链上的物流信息
- `POST /api/v1/blockchain/transfer`: 转移产品所有权
- `GET /api/v1/blockchain/products/:sku`: 从区块链查询产品

## 许可证

MIT License

## 贡献

欢迎贡献！请随时提交 Pull Request。