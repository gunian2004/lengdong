<template>
  <div class="distributor-container">
    <el-header class="header">
      <div class="header-content">
        <div class="title">冷链溯源系统-经销商页</div>
        <div class="welcome">
          欢迎，<span>{{ distributorName }}</span>
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
      <el-aside width="200px">
        <el-menu
            :default-active="activeIndex"
            class="el-menu-vertical-demo"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
        >
          <el-menu-item index="1">
            <i class="el-icon-menu"></i>
            <span>可购买的冷冻品</span>
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-shopping-cart-full"></i>
            <span>我的采购记录</span>
          </el-menu-item>
          <el-menu-item index="3">
            <i class="el-icon-truck"></i>
            <span>物流信息管理</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <div v-if="activeIndex === '1'">
          <h2>可购买的冷冻品 (数据来源: 区块链)</h2>
          <el-table :data="availableProducts" style="width: 100%">
            <el-table-column prop="sku" label="SKU码" width="180">
              <template #default="scope">
                <el-tooltip content="点击复制SKU" placement="top">
                  <span class="sku-copy" @click="copySku(scope.row.sku)">{{ scope.row.sku }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column prop="productName" label="产品名称" width="180"></el-table-column>
            <el-table-column prop="brand" label="品牌" width="180"></el-table-column>
            <el-table-column prop="specification" label="规格/包装" width="180"></el-table-column>
            <el-table-column prop="productionDate" label="生产日期" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.productionDate) }}
              </template>
            </el-table-column>
            <el-table-column prop="expirationDate" label="保质期" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.expirationDate) }}
              </template>
            </el-table-column>
            <el-table-column prop="rawMaterialSource" label="原材料来源" width="180"></el-table-column>
            <el-table-column prop="processingMethod" label="加工方式" width="180"></el-table-column>
            <el-table-column prop="logisticsTemp" label="运输温度要求" width="180"></el-table-column>
            <el-table-column prop="storageCondition" label="仓储条件" width="180"></el-table-column>
            <el-table-column prop="manufacturerID" label="生产商ID" width="180"></el-table-column>
            <el-table-column prop="quantity" label="可用数量" width="180"></el-table-column>
            <el-table-column prop="batchNumber" label="批次号" width="180"></el-table-column>
            <el-table-column prop="imageURL" label="产品图片" width="180">
              <template #default="scope">
                <img
                    :src="scope.row.imageURL"
                    alt="产品图片"
                    v-if="scope.row.imageURL"
                    style="width: 50px; height: 50px; cursor: pointer;"
                    @click="showLargeImage(scope.row.imageURL)"
                >
                <span v-else>图片不可用</span>
              </template>
            </el-table-column>
            <el-table-column label="操作" width="180">
              <template #default="scope">
                <el-button type="primary" @click="openPurchaseDialog(scope.row)">采购</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <div v-else-if="activeIndex === '2'">
          <h2>我的采购记录 (数据来源: 区块链)</h2>
          <el-table :data="purchaseRecords" style="width: 100%">
            <el-table-column prop="purchaseId" label="采购ID" width="180"></el-table-column>
            <el-table-column prop="sku" label="SKU码" width="180">
              <template #default="scope">
                <el-tooltip content="点击复制SKU" placement="top">
                  <span class="sku-copy" @click="copySku(scope.row.sku)">{{ scope.row.sku }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column prop="productName" label="产品名称" width="180"></el-table-column>
            <el-table-column prop="quantity" label="采购数量" width="180"></el-table-column>
            <el-table-column prop="purchaseDate" label="采购日期" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.purchaseDate) }}
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="180">
              <template #default="scope">
                <el-tag :type="getStatusType(scope.row.status)">
                  {{ getStatusText(scope.row.status) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="sellerId" label="卖家ID" width="180"></el-table-column>
            <el-table-column label="操作" width="180">
              <template #default="scope">
                <el-button
                    type="success"
                    @click="confirmPurchase(scope.row.purchaseId)"
                    v-if="scope.row.status !== 'confirmed'"
                >确认收货</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <div v-else-if="activeIndex === '3'">
          <h2>物流信息管理 (数据来源: 区块链)</h2>
          <el-tabs v-model="logisticsActiveTab">
            <el-tab-pane label="添加物流信息" name="add">
              <el-form :model="logisticsForm" ref="logisticsForm" label-width="120px">
                <el-form-item label="SKU码" prop="sku">
                  <el-input
                      v-model="logisticsForm.sku"
                      placeholder="请输入SKU码"
                      @keyup.enter="validateSku"
                      :rules="[{ required: true, message: '请输入SKU码', trigger: 'blur' }]"
                  />
                </el-form-item>
                <el-form-item label="运输方式" prop="transportMethod">
                  <el-input
                      v-model="logisticsForm.transportMethod"
                      placeholder="请输入运输方式"
                      :rules="[{ required: true, message: '请输入运输方式', trigger: 'blur' }]"
                  />
                </el-form-item>
                <el-form-item label="运输公司" prop="carrier">
                  <el-input v-model="logisticsForm.carrier" placeholder="请输入运输公司"></el-input>
                </el-form-item>
                <el-form-item label="运输温度" prop="temperature">
                  <el-input v-model="logisticsForm.temperature" placeholder="请输入运输温度"></el-input>
                </el-form-item>
                <el-form-item label="出发地" prop="departure">
                  <el-input v-model="logisticsForm.departure" placeholder="请输入出发地"></el-input>
                </el-form-item>
                <el-form-item label="目的地" prop="destination">
                  <el-input v-model="logisticsForm.destination" placeholder="请输入目的地"></el-input>
                </el-form-item>
                <el-form-item label="预计到达时间" prop="estimatedArrival">
                  <el-date-picker
                      v-model="logisticsForm.estimatedArrival"
                      type="datetime"
                      placeholder="选择日期时间"
                      :rules="[{ required: true, message: '请选择预计到达时间', trigger: 'change' }]"
                  />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="submitLogistics">提交物流信息</el-button>
                  <el-button type="danger" @click="clearLogisticsForm">清空</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="我的物流记录" name="list">
              <div v-if="logisticsLoading" class="loading-container">
                <el-loading></el-loading>
                <p>加载物流记录中...</p>
              </div>
              <el-table
                  v-else
                  :data="allLogisticsData"
                  style="width: 100%"
                  empty-text="暂无物流记录"
              >
<!--                <el-table-column prop="sku" label="SKU码" width="120">-->
<!--                  <template #default="scope">-->
<!--                    <el-tooltip content="点击复制SKU" placement="top">-->
<!--                      <span class="sku-copy" @click="copySku(scope.row.sku)">{{ scope.row.sku }}</span>-->
<!--                    </el-tooltip>-->
<!--                  </template>-->
<!--                </el-table-column>-->
<!--                <el-table-column prop="productName" label="产品名称" width="120">-->
<!--                  <template #default="scope">-->
<!--                    {{ scope.row.productName || '未知产品' }}-->
<!--                  </template>-->
<!--                </el-table-column>-->
                <el-table-column prop="logistics.transportMethod" label="运输方式" width="120"></el-table-column>
                <el-table-column prop="logistics.carrier" label="运输公司" width="120"></el-table-column>
                <el-table-column prop="logistics.temperature" label="运输温度" width="120"></el-table-column>
                <el-table-column prop="logistics.departure" label="出发地" width="120"></el-table-column>
                <el-table-column prop="logistics.destination" label="目的地" width="120"></el-table-column>
                <el-table-column prop="logistics.timestamp" label="操作时间" width="180">
                  <template #default="scope">
                    {{ formatDate(scope.row.logistics.timestamp) }}
                  </template>
                </el-table-column>
                <el-table-column prop="logistics.operatorName" label="操作方" width="120"></el-table-column>
                <el-table-column prop="logistics.status" label="状态" width="120">
                  <template #default="scope">
                    <el-tag :type="getLogisticsStatusType(scope.row.logistics.status)">
                      {{ getLogisticsStatusText(scope.row.logistics.status) }}
                    </el-tag>
                  </template>
                </el-table-column>
              </el-table>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-main>
    </el-container>

    <!-- 图片预览对话框 -->
    <el-dialog
        v-model="imagePreviewVisible"
        title="商品图片预览"
        width="50%"
        :before-close="handleCloseImagePreview"
    >
      <div style="text-align: center;">
        <img :src="previewImageUrl" style="max-width: 100%; max-height: 70vh;" alt="商品大图">
      </div>
    </el-dialog>

    <!-- 采购对话框 -->
    <el-dialog title="采购商品" v-model="purchaseDialogVisible" width="30%">
      <el-form :model="purchaseForm" ref="purchaseForm" label-width="100px">
        <el-form-item label="SKU码">
          <span>{{ purchaseForm.sku }}</span>
        </el-form-item>
        <el-form-item label="产品名称">
          <span>{{ purchaseForm.productName }}</span>
        </el-form-item>
        <el-form-item label="可用数量">
          <span>{{ purchaseForm.availableQuantity }}</span>
        </el-form-item>
        <el-form-item label="采购数量" prop="quantity">
          <el-input-number v-model="purchaseForm.quantity" :min="1" :max="purchaseForm.availableQuantity"></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="purchaseDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitPurchase">确认采购</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ElMessage } from "element-plus";
import blockchainService from "../utils/blockchainService";

export default {
  name: 'Distributor',
  data() {
    return {
      activeIndex: '1',
      logisticsActiveTab: 'add',
      distributorName: localStorage.getItem('distributorName') || '某某经销商',
      availableProducts: [],
      purchaseRecords: [],
      allLogisticsData: [], // 所有冷冻品和物流信息
      purchaseDialogVisible: false,
      imagePreviewVisible: false,
      previewImageUrl: '',
      blockchainConnected: false,
      loading: false,
      logisticsLoading: false,
      errorState: {
        products: false,
        purchases: false,
        logistics: false
      },
      purchaseForm: {
        sku: '',
        productName: '',
        availableQuantity: 0,
        quantity: 1
      },
      logisticsForm: {
        sku: '',
        transportMethod: '',
        carrier: '',
        temperature: '',
        departure: '',
        destination: '',
        estimatedArrival: ''
      }
    };
  },
  computed: {
    confirmedPurchases() {
      return this.purchaseRecords.filter(record => record.status === 'confirmed');
    }
  },
  methods: {
    // 连接区块链
    async connectBlockchain() {
      try {
        this.loading = true;
        const connected = await blockchainService.init();
        if (connected) {
          this.blockchainConnected = true;
          ElMessage.success("区块链连接成功");
          await this.fetchAvailableProducts();
          await this.fetchPurchaseRecords();
          await this.fetchAllLogisticsData();
        } else {
          ElMessage.error("区块链连接失败，请安装MetaMask钱包");
        }
      } catch (error) {
        console.error("区块链连接错误:", error);
        ElMessage.error("区块链连接错误: " + error.message);
      } finally {
        this.loading = false;
      }
    },

    handleSelect(key) {
      this.activeIndex = key;
      if (key === '1') {
        this.fetchAvailableProducts();
      } else if (key === '2') {
        this.fetchPurchaseRecords();
      } else if (key === '3') {
        this.fetchAllLogisticsData();
      }
    },

    formatDate(timestamp) {
      if (!timestamp || isNaN(timestamp)) return '无数据';
      return new Date(timestamp * 1000).toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      });
    },

    getStatusType(status) {
      const statusMap = {
        'pending': 'warning',
        'confirmed': 'success',
        'cancelled': 'danger'
      };
      return statusMap[status] || 'info';
    },

    getStatusText(status) {
      const statusMap = {
        'pending': '待确认',
        'confirmed': '已确认',
        'cancelled': '已取消'
      };
      return statusMap[status] || '未知状态';
    },

    getLogisticsStatusType(status) {
      const statusMap = {
        'in_transit': 'warning',
        'delivered': 'success',
        'delayed': 'danger',
        'pending': 'info'
      };
      return statusMap[status] || 'info';
    },

    getLogisticsStatusText(status) {
      const statusMap = {
        'in_transit': '运输中',
        'delivered': '已送达',
        'delayed': '延迟',
        'pending': '待处理'
      };
      return statusMap[status] || '未知状态';
    },

    copySku(sku) {
      navigator.clipboard.writeText(sku).then(() => {
        ElMessage.success('SKU已复制到剪贴板');
      }).catch(err => {
        console.error('复制失败: ', err);
        ElMessage.error('复制失败');
      });
    },

    validateSku() {
      const skuPattern = /^[A-Za-z0-9]{8,20}$/;
      if (this.logisticsForm.sku && !skuPattern.test(this.logisticsForm.sku)) {
        ElMessage.warning('SKU格式应为8-20位字母和数字');
      }
    },

    // 从区块链获取可用产品
    async fetchAvailableProducts() {
      this.loading = true;
      this.errorState.products = false;

      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const productsData = await blockchainService.getAllAvailableProducts();
        this.availableProducts = productsData.map(product => ({
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
          batchNumber: product.batchNumber,
          quantity: product.quantity,
          currentHolder: product.currentHolder,
          currentLocation: product.currentLocation,
          isValid: product.isValid
        }));

        ElMessage.success(`成功获取${this.availableProducts.length}个产品`);
      } catch (error) {
        console.error('获取可用产品失败:', error);
        this.errorState.products = true;
        this.availableProducts = [];
        ElMessage.error('获取可用产品失败: ' + error.message);
      } finally {
        this.loading = false;
      }
    },

    // 从区块链获取采购记录
    async fetchPurchaseRecords() {
      this.loading = true;
      this.errorState.purchases = false;

      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const distributorAddress = await blockchainService.getAccountAddress();
        const { purchases } = await blockchainService.getBuyerCompletePurchaseHistory(distributorAddress);
        this.purchaseRecords = purchases.map(purchase => ({
          purchaseId: purchase.purchaseId,
          sku: purchase.sku,
          productName: purchase.productName,
          quantity: purchase.quantity,
          purchaseDate: purchase.purchaseDate,
          status: purchase.status,
          sellerId: purchase.sellerId,
          buyerId: purchase.buyerId,
          isValid: purchase.isValid
        }));

        ElMessage.success(`成功获取${this.purchaseRecords.length}条采购记录`);
      } catch (error) {
        console.error('获取采购记录失败:', error);
        this.errorState.purchases = true;
        this.purchaseRecords = [];
        ElMessage.error('获取采购记录失败: ' + error.message);
      } finally {
        this.loading = false;
      }
    },

    // 从区块链获取所有冷冻品和物流信息
    async fetchAllLogisticsData() {
      this.logisticsLoading = true;
      this.errorState.logistics = false;
      this.allLogisticsData = [];

      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        // 获取所有采购记录的SKU
        const allSkus = this.purchaseRecords.map(record => record.sku);
        if (allSkus.length === 0) {
          ElMessage.info("暂无采购记录，无法获取物流信息");
          return;
        }

        // 创建SKU到产品名称的映射，用于数据处理
        const skuToProductName = {};
        this.purchaseRecords.forEach(record => {
          skuToProductName[record.sku] = record.productName;
        });

        // 获取所有冷冻品的完整信息
        const allProductsInfo = await Promise.all(
            allSkus.map(sku => blockchainService.getCompleteProductInfo(sku))
        );

        // 处理数据格式
        const processedData = [];
        allProductsInfo.forEach(productInfo => {
          if (productInfo && productInfo.logisticsRecords && productInfo.logisticsRecords.length > 0) {
            // 遍历该产品的所有物流记录
            productInfo.logisticsRecords.forEach(logistics => {
              // 从映射中获取产品名称，如果没有则使用SKU
              const productName = skuToProductName[productInfo.sku] || productInfo.productName || productInfo.sku;

              processedData.push({
                sku: productInfo.sku,
                productName: productName,
                logistics: {
                  ...logistics,
                  status: logistics.isValid ? 'in_transit' : 'delivered',
                  timestamp: logistics.timestamp || Date.now() / 1000
                }
              });
            });
          }
        });

        if (processedData.length > 0) {
          // 按时间戳降序排列，最新的记录在前
          this.allLogisticsData = processedData.sort((a, b) => b.logistics.timestamp - a.logistics.timestamp);
          ElMessage.success(`成功获取${this.allLogisticsData.length}条物流记录`);
        } else {
          ElMessage.info('暂无物流记录');
        }
      } catch (error) {
        console.error('获取物流记录失败:', error);
        this.errorState.logistics = true;
        ElMessage.error('获取物流记录失败: ' + error.message);
      } finally {
        this.logisticsLoading = false;
      }
    },

    openPurchaseDialog(product) {
      this.purchaseForm = {
        sku: product.sku,
        productName: product.productName,
        availableQuantity: product.quantity,
        quantity: 1
      };
      this.purchaseDialogVisible = true;
    },

    // 提交采购请求到区块链
    async submitPurchase() {
      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const distributorAddress = await blockchainService.getAccountAddress();
        const productData = this.availableProducts.find(item => item.sku === this.purchaseForm.sku);

        if (!productData) {
          ElMessage.error('产品信息不存在');
          return;
        }

        const purchaseData = {
          sku: this.purchaseForm.sku,
          productName: this.purchaseForm.productName,
          quantity: this.purchaseForm.quantity,
          sellerId: productData.currentHolder,
          buyerId: distributorAddress
        };

        const txResult = await blockchainService.createPurchaseRecord(purchaseData);

        if (txResult) {
          ElMessage.success('采购申请已提交到区块链');
          this.purchaseDialogVisible = false;
          this.fetchAvailableProducts();
          this.fetchPurchaseRecords();
          this.fetchAllLogisticsData();
        } else {
          ElMessage.error('采购失败');
        }
      } catch (error) {
        console.error('采购失败:', error);
        ElMessage.error('采购失败: ' + error.message);
      }
    },

    // 确认收货并更新到区块链
    async confirmPurchase(purchaseId) {
      try {
        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        const txResult = await blockchainService.confirmPurchase(purchaseId);

        if (txResult) {
          ElMessage.success('已确认收货并更新到区块链');
          this.fetchPurchaseRecords();
          this.fetchAllLogisticsData();
        } else {
          ElMessage.error('确认收货失败');
        }
      } catch (error) {
        console.error('确认收货失败:', error);
        ElMessage.error('确认收货失败: ' + error.message);
      }
    },

    // 提交物流信息到区块链
    async submitLogistics() {
      try {
        if (!this.logisticsForm.sku) {
          ElMessage.warning('请输入SKU码');
          return;
        }

        const skuPattern = /^[A-Za-z0-9]{8,20}$/;
        if (!skuPattern.test(this.logisticsForm.sku)) {
          ElMessage.warning('SKU格式应为8-20位字母和数字');
          return;
        }

        if (!this.logisticsForm.transportMethod) {
          ElMessage.warning('请输入运输方式');
          return;
        }

        if (!this.logisticsForm.estimatedArrival) {
          ElMessage.warning('请选择预计到达时间');
          return;
        }

        const purchaseRecord = this.confirmedPurchases.find(item => item.sku === this.logisticsForm.sku);
        if (!purchaseRecord) {
          ElMessage.warning('未找到该SKU的已确认采购记录');
          return;
        }

        const logisticsData = {
          sku: this.logisticsForm.sku,
          transportMethod: this.logisticsForm.transportMethod,
          carrier: this.logisticsForm.carrier,
          temperature: this.logisticsForm.temperature,
          departure: this.logisticsForm.departure,
          destination: this.logisticsForm.destination,
          estimatedArrival: Math.floor(new Date(this.logisticsForm.estimatedArrival).getTime() / 1000),
          operatorName: this.distributorName,
          purchaseId: purchaseRecord.purchaseId
        };

        if (!this.blockchainConnected) {
          await this.connectBlockchain();
          if (!this.blockchainConnected) return;
        }

        await blockchainService.addLogisticsRecord(logisticsData);
        ElMessage.success('物流信息已成功上链');
        this.clearLogisticsForm();
        this.fetchAllLogisticsData();
        this.logisticsActiveTab = 'list';
      } catch (error) {
        console.error('提交物流信息失败:', error);
        ElMessage.error('提交物流信息失败: ' + (error.message || '请检查输入信息'));
      }
    },

    clearLogisticsForm() {
      if (this.$refs.logisticsForm) {
        this.$refs.logisticsForm.resetFields();
      }
      this.logisticsForm = {
        sku: '',
        transportMethod: '',
        carrier: '',
        temperature: '',
        departure: '',
        destination: '',
        estimatedArrival: ''
      };
    },

    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('userId');
      localStorage.removeItem('distributorName');
      this.$router.push('/');
    },

    showLargeImage(url) {
      this.previewImageUrl = url;
      this.imagePreviewVisible = true;
    },

    handleCloseImagePreview() {
      this.imagePreviewVisible = false;
      this.previewImageUrl = '';
    }
  },
  mounted() {
    this.connectBlockchain();
  }
};
</script>

<style scoped>
.distributor-container {
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

.el-aside {
  background-color: #545c64;
  color: #fff;
}

.el-main {
  padding: 20px;
  overflow-y: auto;
}

.el-menu-vertical-demo {
  height: 100%;
}

.el-form-item {
  margin-bottom: 15px;
}

.el-button {
  margin-top: 10px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
}

.el-dialog__body {
  padding: 20px;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
}

.sku-copy {
  cursor: pointer;
  color: #409EFF;
  text-decoration: underline;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
  color: #909399;
}
</style>