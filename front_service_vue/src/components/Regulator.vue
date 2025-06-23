<template>
  <div class="regulator-container">
    <el-header class="header">
      <div class="header-content">
        <div class="title">冷链溯源系统-监管者页</div>
        <div class="welcome">
          欢迎，<span>{{ regulatorName }}</span>
        </div>
        <div class="blockchain-status">
          <el-tag :type="blockchainConnected ? 'success' : 'danger'">
            {{ blockchainConnected ? "区块链已连接" : "区块链未连接" }}
          </el-tag>
        </div>
        <el-button type="primary" @click="logout">退出登录</el-button>
      </div>
    </el-header>
    <el-container>
      <el-main>
        <!-- 区块链连接按钮 -->
        <el-button
            v-if="!blockchainConnected"
            type="primary"
            @click="connectBlockchain"
        >
          连接区块链
        </el-button>

        <!-- 添加搜索输入框和按钮 -->
        <el-input
            v-model="searchSku"
            placeholder="请输入 SKU 码进行搜索"
            @keyup.enter="searchBySku"
        ></el-input>
        <el-button @click="searchBySku">搜索</el-button>

        <h2>所有冷冻品信息 (数据来源: 区块链)</h2>

        <div v-if="loading" class="loading-container">
          <el-loading></el-loading>
          <p>加载中...</p>
        </div>

        <el-table
            v-else
            :data="products"
            style="width: 100%"
            class="custom-table"
        >
          <!-- SKU列，添加点击复制功能 -->
          <el-table-column
              label="SKU码"
              width="180"
          >
            <template #default="scope">
              <span
                  class="copyable-sku"
                  @click="copyToClipboard(scope.row.product.sku)"
              >
                {{ scope.row.product.sku }}
              </span>
              <el-tooltip
                  :content="copyStatus === scope.row.product.sku ? '已复制!' : '点击复制'"
                  placement="top"
              >
                <template #content>
                  <span>{{ copyStatus === scope.row.product.sku ? '已复制!' : '点击复制' }}</span>
                </template>
              </el-tooltip>
            </template>
          </el-table-column>

          <el-table-column
              prop="product.productName"
              label="产品名称"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.brand"
              label="品牌"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.specification"
              label="规格/包装"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.productionDate"
              label="生产日期"
              width="180"
              :formatter="formatDate"
          ></el-table-column>
          <el-table-column
              prop="product.expirationDate"
              label="保质期"
              width="180"
              :formatter="formatDate"
          ></el-table-column>
          <el-table-column
              prop="product.rawMaterialSource"
              label="原材料来源"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.processingMethod"
              label="加工方式"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.logisticsTemp"
              label="运输温度要求"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="product.storageCondition"
              label="仓储条件"
              width="180"
          ></el-table-column>
          <el-table-column label="质检报告URL" width="180">
            <template #default="scope">
              <a
                  :href="scope.row.product.qcReport"
                  target="_blank"
                  v-if="scope.row.product.qcReport"
              >
                查看质检报告
              </a>
              <span v-else>质检报告不可用</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="product.manufacturerID"
              label="生产商ID"
              width="180"
          ></el-table-column>
          <el-table-column label="产品图片URL" width="180">
            <template #default="scope">
              <img
                  :src="scope.row.product.imageURL"
                  alt="产品图片"
                  v-if="scope.row.product.imageURL"
                  style="width: 50px; height: 50px"
              />
              <span v-else>图片不可用</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="currentHolder"
              label="当前交易哈希"
              width="180"
              :formatter="formatAddress"
          ></el-table-column>
          <el-table-column
              prop="currentLocation"
              label="当前位置"
              width="180"
          ></el-table-column>
          <el-table-column
              prop="quantity"
              label="数量"
              width="180"
          ></el-table-column>
          <!-- 区块链特有字段 - 批次号 -->
          <el-table-column
              prop="batchNumber"
              label="批次号"
              width="180"
          ></el-table-column>
          <el-table-column
              label="查看物流"
              width="100"
          >
            <template #default="scope">
              <el-button size="small" @click="viewLogistics(scope.row.product.sku)">
                查看
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <!-- 物流详情对话框 -->
        <el-dialog
            v-model="logisticsDialogVisible"
            title="物流信息详情"
            width="80%"
            :before-close="handleCloseLogisticsDialog"
        >
          <el-table :data="currentLogisticsRecords" style="width: 100%">
            <el-table-column
                prop="transportMethod"
                label="运输方式"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="carrier"
                label="运输公司"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="temperature"
                label="运输温度"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="departure"
                label="出发地"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="destination"
                label="目的地"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="operatorName"
                label="操作方"
                width="180"
            ></el-table-column>
            <el-table-column
                prop="timestamp"
                label="操作时间"
                width="180"
                :formatter="formatLogisticsDate"
            ></el-table-column>
          </el-table>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { ElMessage } from "element-plus";
import blockchainService from "../utils/blockchainService";

export default {
  data() {
    return {
      regulatorName: localStorage.getItem("regulatorName") || "监管者",
      products: [],
      searchSku: "",
      blockchainConnected: false,
      loading: false,
      logisticsDialogVisible: false,
      currentLogisticsRecords: [],
      copyStatus: null, // 用于跟踪复制状态
    };
  },
  methods: {
    // 格式化日期显示
    formatDate(row, column) {
      const date = row[column.property];
      return date
          ? new Date(date).toLocaleDateString("zh-CN", {
            year: "numeric",
            month: "2-digit",
            day: "2-digit",
            hour: "2-digit",
            minute: "2-digit"
          })
          : "无数据";
    },

    // 格式化物流记录时间
    formatLogisticsDate(row) {
      const timestamp = row.timestamp || Math.floor(new Date().getTime() / 1000);
      return new Date(timestamp * 1000).toLocaleDateString("zh-CN", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit"
      });
    },

    // 格式化区块链地址显示（缩写处理）
    formatAddress(row) {
      const address = row.currentHolder;
      if (!address) return "无数据";
      return `${address.slice(0, 6)}...${address.slice(-6)}`;
    },

    // 复制文本到剪贴板
    copyToClipboard(text) {
      navigator.clipboard.writeText(text)
          .then(() => {
            this.copyStatus = text;
            ElMessage.success('SKU已复制到剪贴板');

            // 3秒后重置复制状态
            setTimeout(() => {
              this.copyStatus = null;
            }, 3000);
          })
          .catch(err => {
            console.error('复制失败:', err);
            ElMessage.error('复制失败，请手动复制');
          });
    },

    // 连接区块链
    async connectBlockchain() {
      try {
        this.loading = true;
        const connected = await blockchainService.init();
        if (connected) {
          this.blockchainConnected = true;
          ElMessage.success("区块链连接成功");
          this.fetchProducts();
        } else {
          ElMessage.error("区块链连接失败，请安装MetaMask钱包");
        }
      } catch (error) {
        ElMessage.error("区块链连接错误: " + error.message);
        console.error(error);
      } finally {
        this.loading = false;
      }
    },

    // 从区块链获取所有产品
    async fetchProducts() {
      this.loading = true;
      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const products = await blockchainService.getAllAvailableProducts();
        this.products = products.map(item => ({
          // 转换数据格式以适配表格
          product: {
            sku: item.sku,
            productName: item.productName,
            brand: item.brand,
            specification: item.specification,
            productionDate: item.productionDate,
            expirationDate: item.expirationDate,
            rawMaterialSource: item.rawMaterialSource,
            processingMethod: item.processingMethod,
            logisticsTemp: item.logisticsTemp,
            storageCondition: item.storageCondition,
            qcReport: item.qcReport,
            manufacturerID: item.manufacturerID,
            imageURL: item.imageURL,
          },
          batchNumber: item.batchNumber,
          currentHolder: item.currentHolder,
          currentLocation: item.currentLocation,
          quantity: item.quantity,
        }));
      } catch (error) {
        ElMessage.error("获取区块链产品失败: " + error.message);
        console.error(error);
      } finally {
        this.loading = false;
      }
    },

    // 搜索产品
    async searchBySku() {
      if (!this.searchSku) {
        ElMessage.warning("请输入 SKU 码");
        return;
      }

      this.loading = true;
      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const { product, logisticsRecords } = await blockchainService.getCompleteProductInfo(this.searchSku);
        if (product) {
          this.products = [{
            product: {
              sku: product.sku,
              productName: product.productName,
              brand: product.brand,
              specification: product.specification,
              productionDate: product.productionDate,
              expirationDate: product.expirationDate,
              rawMaterialSource: product.rawMaterialSource,
              processingMethod: product.processingMethod,
              logisticsTemp: product.logisticsTemp,
              storageCondition: product.storageCondition,
              qcReport: product.qcReport,
              manufacturerID: product.manufacturerID,
              imageURL: product.imageURL,
            },
            batchNumber: product.batchNumber,
            currentHolder: product.currentHolder,
            currentLocation: product.currentLocation,
            quantity: product.quantity,
          }];

          // 处理物流记录，添加时间戳
          this.currentLogisticsRecords = logisticsRecords.map(rec => ({
            ...rec,
            timestamp: rec.productionDate ? Math.floor(new Date(rec.productionDate).getTime() / 1000) : Date.now() / 1000
          }));
        } else {
          this.products = [];
          ElMessage.warning("未找到该产品");
        }
      } catch (error) {
        ElMessage.error("查询区块链产品失败: " + error.message);
        console.error(error);
      } finally {
        this.loading = false;
      }
    },

    // 查看物流信息
    async viewLogistics(sku) {
      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const { logisticsRecords } = await blockchainService.getCompleteProductInfo(sku);
        // 处理物流记录，添加时间戳
        this.currentLogisticsRecords = logisticsRecords.map(rec => ({
          ...rec,
          timestamp: rec.productionDate ? Math.floor(new Date(rec.productionDate).getTime() / 1000) : Date.now() / 1000
        }));
        this.logisticsDialogVisible = true;
      } catch (error) {
        ElMessage.error("获取物流信息失败: " + error.message);
        console.error(error);
      }
    },

    // 关闭物流对话框
    handleCloseLogisticsDialog() {
      this.logisticsDialogVisible = false;
    },

    logout() {
      localStorage.removeItem("token");
      localStorage.removeItem("role_type");
      localStorage.removeItem("regulatorName");
      this.$router.push("/");
    },
  },
  mounted() {
    this.fetchProducts();
  },
};
</script>

<style scoped>
.regulator-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.header {
  background-color: #409eff;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.title {
  font-size: 1.5em;
}

.welcome {
  font-size: 1.2em;
}

.blockchain-status {
  margin: 0 15px;
}

.el-main {
  padding: 20px;
}

.el-input {
  width: 300px;
  margin-right: 10px;
}

.el-table {
  margin-top: 20px;
}

.custom-table {
  background-color: white;
  border-collapse: collapse;
}

.custom-table th,
.custom-table td {
  border: 1px solid #bdc3c7;
  padding: 8px;
  text-align: left;
}

.custom-table th {
  background-color: #ecf0f1;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
}

/* 添加SKU复制样式 */
.copyable-sku {
  color: #409eff;
  text-decoration: underline;
  cursor: pointer;
}

.copyable-sku:hover {
  color: #1890ff;
}
</style>