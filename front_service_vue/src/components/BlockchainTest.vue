<template>
  <div class="blockchain-test-container">
    <h1>区块链连接测试</h1>

    <div class="connection-status">
      <p>
        连接状态:
        <span :class="blockchainConnected ? 'connected' : 'disconnected'">
          {{ blockchainConnected ? "已连接" : "未连接" }}
        </span>
      </p>
      <button @click="connectBlockchain" :disabled="blockchainConnected">
        连接区块链
      </button>
    </div>

    <div class="product-list" v-if="blockchainConnected">
      <h2>区块链产品列表</h2>
      <div v-if="loading" class="loading">加载中...</div>
      <div v-else-if="products.length === 0" class="no-data">暂无产品数据</div>
      <table v-else>
        <thead>
          <tr>
            <th>产品ID</th>
            <th>名称</th>
            <th>类别</th>
            <th>生产商</th>
            <th>生产日期</th>
            <th>过期日期</th>
            <th>存储温度</th>
            <th>批次号</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in products" :key="product.product.sku">
            <td>{{ product.product.sku }}</td>
            <td>{{ product.product.product_name }}</td>
            <td>{{ product.product.specification }}</td>
            <td>{{ product.product.manufacturer_id }}</td>
            <td>{{ formatDate(product.product.production_date) }}</td>
            <td>{{ formatDate(product.product.expiration_date) }}</td>
            <td>{{ product.product.logistics_temp }}</td>
            <td>{{ product.batch_number }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="debug-info">
      <h3>调试信息</h3>
      <pre>{{ debugInfo }}</pre>
    </div>
  </div>
</template>

<script>
import blockchainService from "../utils/blockchainService";

export default {
  name: "BlockchainTest",
  data() {
    return {
      blockchainConnected: false,
      products: [],
      loading: false,
      debugInfo: "",
    };
  },
  methods: {
    async connectBlockchain() {
      try {
        this.loading = true;
        this.debugInfo = "正在连接区块链...";

        const connected = await blockchainService.init();
        this.blockchainConnected = connected;

        if (connected) {
          this.debugInfo += "\n连接成功！";
          await this.loadProducts();
        } else {
          this.debugInfo += "\n连接失败，请确保安装了MetaMask并授权连接。";
        }
      } catch (error) {
        this.debugInfo += `\n连接出错: ${error.message}`;
        console.error("区块链连接错误:", error);
      } finally {
        this.loading = false;
      }
    },

    async loadProducts() {
      try {
        this.loading = true;
        this.debugInfo += "\n正在加载产品数据...";

        const products = await blockchainService.getAllProducts();
        this.products = products;

        this.debugInfo += `\n加载了 ${products.length} 个产品`;
      } catch (error) {
        this.debugInfo += `\n加载产品出错: ${error.message}`;
        console.error("加载产品错误:", error);
      } finally {
        this.loading = false;
      }
    },

    formatDate(dateString) {
      try {
        return new Date(dateString).toLocaleDateString("zh-CN");
      } catch (e) {
        return dateString;
      }
    },
  },
};
</script>

<style scoped>
.blockchain-test-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.connection-status {
  margin-bottom: 20px;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.connected {
  color: green;
  font-weight: bold;
}

.disconnected {
  color: red;
  font-weight: bold;
}

button {
  background-color: #4caf50;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

.loading {
  text-align: center;
  padding: 20px;
  font-style: italic;
}

.no-data {
  text-align: center;
  padding: 20px;
  color: #888;
}

.debug-info {
  margin-top: 30px;
  padding: 15px;
  background-color: #f8f8f8;
  border: 1px solid #ddd;
  border-radius: 4px;
}

pre {
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
