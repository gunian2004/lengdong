<template>
  <div class="distributor-container">
    <el-header class="header">
      <div class="header-content">
        <div class="title">冷链溯源系统-经销商页</div>
        <div class="welcome">
          欢迎，<span>{{ distributorName }}</span>
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
          <h2>可购买的冷冻品</h2>
          <el-table :data="availableProducts" style="width: 100%">
            <el-table-column prop="product_name" label="产品名称" width="180"></el-table-column>
            <el-table-column prop="brand" label="品牌" width="180"></el-table-column>
            <el-table-column prop="specification" label="规格/包装" width="180"></el-table-column>
            <el-table-column prop="production_date" label="生产日期" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.production_date) }}
              </template>
            </el-table-column>
            <el-table-column prop="expiration_date" label="保质期" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.expiration_date) }}
              </template>
            </el-table-column>
            <el-table-column prop="raw_material_source" label="原材料来源" width="180"></el-table-column>
            <el-table-column prop="processing_method" label="加工方式" width="180"></el-table-column>
            <el-table-column prop="logistics_temp" label="运输温度要求" width="180"></el-table-column>
            <el-table-column prop="storage_condition" label="仓储条件" width="180"></el-table-column>
            <el-table-column prop="manufacturer_id" label="生产商ID" width="180"></el-table-column>
            <el-table-column prop="quantity" label="可用数量" width="180"></el-table-column>
            <el-table-column prop="image_url" label="产品图片" width="180">
              <template #default="scope">
                <img 
                  :src="scope.row.image_url" 
                  alt="产品图片" 
                  v-if="scope.row.image_url" 
                  style="width: 50px; height: 50px; cursor: pointer;"
                  @click="showLargeImage(scope.row.image_url)"
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
          <h2>我的采购记录</h2>
          <el-table :data="purchaseRecords" style="width: 100%">
            <el-table-column prop="purchase_id" label="采购ID" width="120"></el-table-column>
            <el-table-column prop="buyer_id" label="买家ID" width="120"></el-table-column>
            <el-table-column prop="seller_id" label="卖家ID" width="120"></el-table-column>
            <el-table-column prop="sku" label="产品SKU" width="180"></el-table-column>
            <el-table-column prop="batch_id" label="批次ID" width="120"></el-table-column>
            <el-table-column prop="quantity" label="采购数量" width="120"></el-table-column>
            <el-table-column prop="purchase_time" label="采购时间" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.purchase_time) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
        
        <div v-else-if="activeIndex === '3'">
          <h2>物流信息管理</h2>
          <el-tabs v-model="logisticsActiveTab">
            <el-tab-pane label="添加物流信息" name="add">
              <el-form :model="logisticsForm" ref="logisticsForm" label-width="120px">
                <el-form-item label="采购ID" prop="purchase_id">
                  <el-select v-model="logisticsForm.purchase_id" placeholder="请选择采购记录">
                    <el-option 
                      v-for="item in confirmedPurchases" 
                      :key="item.purchase_id" 
                      :label="`${item.purchase_id}`" 
                      :value="item.purchase_id">
                    </el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="物流单号" prop="logistics_no">
                  <el-input v-model="logisticsForm.logistics_no" placeholder="请输入物流单号"></el-input>
                </el-form-item>
                <el-form-item label="冷库位置" prop="cold_storage_location">
                  <el-input v-model="logisticsForm.cold_storage_location" placeholder="请输入冷库位置"></el-input>
                </el-form-item>
                <el-form-item label="运输温度(℃)" prop="temperature">
                  <el-input v-model="logisticsForm.temperature" placeholder="请输入运输温度"></el-input>
                </el-form-item>
                <el-form-item label="湿度(%)" prop="humidity">
                  <el-input v-model="logisticsForm.humidity" placeholder="请输入湿度"></el-input>
                </el-form-item>
                <el-form-item label="图片URL(JSON数组)" prop="logistics_images">
                  <el-input v-model="logisticsForm.logistics_images" placeholder="请输入图片URL，多个用英文逗号分隔"></el-input>
                </el-form-item>
                <el-form-item label="区块链交易哈希" prop="blockchain_tx_hash">
                  <el-input v-model="logisticsForm.blockchain_tx_hash" placeholder="请输入区块链交易哈希"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="submitLogistics">提交物流信息</el-button>
                  <el-button type="danger" @click="clearLogisticsForm">清空</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="我的物流记录" name="list">
              <el-table :data="logisticsRecords" style="width: 100%">
                <el-table-column prop="logistics_id" label="物流ID" width="120"></el-table-column>
                <el-table-column prop="purchase_id" label="采购ID" width="120"></el-table-column>
                <el-table-column prop="logistics_no" label="物流单号" width="150"></el-table-column>
                <el-table-column prop="cold_storage_location" label="冷库位置" width="150"></el-table-column>
                <el-table-column prop="temperature" label="温度(℃)" width="120"></el-table-column>
                <el-table-column prop="humidity" label="湿度(%)" width="120"></el-table-column>
                <el-table-column prop="logistics_images" label="图片URL(JSON数组)" width="200">
                  <template #default="scope">
                    <div v-if="Array.isArray(scope.row.logistics_images)">
                      <img v-for="img in scope.row.logistics_images" :src="img" :key="img" style="width:40px;height:40px;margin-right:5px;" />
                    </div>
                    <div v-else>{{ scope.row.logistics_images }}</div>
                  </template>
                </el-table-column>
                <el-table-column prop="blockchain_tx_hash" label="区块链交易哈希" width="200"></el-table-column>
                <el-table-column prop="created_at" label="创建时间" width="180">
                  <template #default="scope">
                    {{ formatDate(scope.row.created_at) }}
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
        <el-form-item label="产品名称">
          <span>{{ purchaseForm.product_name }}</span>
        </el-form-item>
        <el-form-item label="可用数量">
          <span>{{ purchaseForm.available_quantity }}</span>
        </el-form-item>
        <el-form-item label="采购数量" prop="quantity">
          <el-input-number v-model="purchaseForm.quantity" :min="1" :max="purchaseForm.available_quantity"></el-input-number>
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
import axios from 'axios';
import { ElMessage } from 'element-plus';
import { serviceEndpoints } from '../services';

export default {
  name: 'Distributor',
  data() {
    return {
      activeIndex: '1',
      logisticsActiveTab: 'add',
      distributorName: localStorage.getItem('distributorName') || '某某经销商',
      availableProducts: [],
      purchaseRecords: [],
      logisticsRecords: [],
      purchaseDialogVisible: false,
      purchaseForm: {
        product_id: '',
        batch_id: '',
        product_name: '',
        available_quantity: 0,
        quantity: 1
      },
      logisticsForm: {
        purchase_id: '',
        logistics_no: '',
        cold_storage_location: '',
        temperature: '',
        humidity: '',
        logistics_images: '',
        blockchain_tx_hash: ''
      },
      imagePreviewVisible: false,
      previewImageUrl: ''
    };
  },
  computed: {
    // 原来是只显示confirmed状态的采购记录
    // confirmedPurchases() {
    //   return this.purchaseRecords.filter(record => record.status === 'confirmed');
    // }
    // 临时方案：直接返回所有采购记录
    confirmedPurchases() {
      return this.purchaseRecords;
    }
  },
  methods: {
    handleSelect(key) {
      this.activeIndex = key;
      if (key === '1') {
        this.fetchAvailableProducts();
      } else if (key === '2') {
        this.fetchPurchaseRecords();
      } else if (key === '3') {
        this.fetchPurchaseRecords();
        this.fetchLogisticsRecords();
      }
    },
    formatDate(dateString) {
      if (!dateString) return '无数据';
      const date = new Date(dateString);
      return date.toLocaleString();
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
        'delayed': 'danger'
      };
      return statusMap[status] || 'info';
    },
    getLogisticsStatusText(status) {
      const statusMap = {
        'in_transit': '运输中',
        'delivered': '已送达',
        'delayed': '延迟'
      };
      return statusMap[status] || '未知状态';
    },
    async fetchAvailableProducts(retryCount = 0) {
      try {
        const params = {};
        const distributorId = localStorage.getItem('manufacturer_id');
        if (distributorId) {
          params.distributor_id = distributorId;
        }
        
        const response = await axios.get(`${serviceEndpoints.distributor}/api/products/available`, { params });
        
        if (response.data && response.data.products) {
          // 确保每个产品对象都有 quantity 属性
          this.availableProducts = response.data.products.map(product => ({
            ...product,
            quantity: product.quantity || 0  // 如果没有 quantity，默认设置为 0
          }));
        } else if (response.data && Array.isArray(response.data)) {
          // 如果返回的是数组，同样确保有 quantity 属性
          this.availableProducts = response.data.map(product => ({
            ...product,
            quantity: product.quantity || 0  // 如果没有 quantity，默认设置为 0
          }));
        } else {
          this.availableProducts = [];
          ElMessage.warning('获取产品数据格式异常');
        }
      } catch (error) {
        console.error('获取可用产品失败:', error);
        
        if (retryCount < 3) {
          setTimeout(() => {
            this.fetchAvailableProducts(retryCount + 1);
          }, 1000);
        } else {
          ElMessage.error('获取可用产品失败，请检查网络连接或联系管理员');
        }
      }
    },
    async fetchPurchaseRecords() {
      try {
        const distributorId = localStorage.getItem('manufacturer_id');
        const response = await axios.get(`${serviceEndpoints.distributor}/api/purchase`, {
          params: { buyer_id: distributorId }
        });
        this.purchaseRecords = response.data.purchases || [];
      } catch (error) {
        console.error('获取采购记录失败:', error);
        ElMessage.error('获取采购记录失败');
      }
    },
    async fetchLogisticsRecords() {
      try {
        const distributorId = localStorage.getItem('manufacturer_id');
        const response = await axios.get(`${serviceEndpoints.distributor}/api/logistics`, {
          params: { distributor_id: distributorId }
        });
        console.log('物流接口返回数据:', response.data); // 调试用
        this.logisticsRecords = response.data.logistics || [];
      } catch (error) {
        console.error('获取物流记录失败:', error);
        ElMessage.error('获取物流记录失败');
      }
    },
    openPurchaseDialog(product) {
      this.purchaseForm = {
        product_id: product.sku,  // 使用产品的SKU
        batch_id: product.batch_id,
        product_name: product.product_name,
        available_quantity: product.quantity,
        quantity: 1,
        manufacturer_id: product.manufacturer_id
      };
      this.purchaseDialogVisible = true;
    },
    async submitPurchase() {
      try {
        const buyerId = parseInt(localStorage.getItem('manufacturer_id'));
        const purchaseData = {
          buyer_id: buyerId,
          seller_id: parseInt(this.purchaseForm.manufacturer_id), // 转换为数字
          sku: this.purchaseForm.product_id,
          batch_id: parseInt(this.purchaseForm.batch_id), // 转换为数字
          quantity: parseInt(this.purchaseForm.quantity)
        };
        
        const response = await axios.post(`${serviceEndpoints.distributor}/api/purchase`, purchaseData);
        
        // 修改这里的判断逻辑，检查响应状态码而不是success字段
        if (response.status >= 200 && response.status < 300) {
          ElMessage.success('采购成功');
          this.purchaseDialogVisible = false;
          await this.fetchAvailableProducts();
          await this.fetchPurchaseRecords();
        } else {
          ElMessage.error(response.data.message || '采购失败');
        }
      } catch (error) {
        console.error('采购失败:', error);
        ElMessage.error('采购失败: ' + (error.response?.data?.message || error.message));
      }
    },
    async confirmPurchase(purchaseId) {
      try {
        const response = await axios.put(`${serviceEndpoints.distributor}/api/purchase/confirm/${purchaseId}`);
        
        if (response.data.success) {
          ElMessage.success('已确认收货');
          this.fetchPurchaseRecords();
        } else {
          ElMessage.error(response.data.message || '确认收货失败');
        }
      } catch (error) {
        console.error('确认收货失败:', error);
        ElMessage.error('确认收货失败: ' + (error.response?.data?.message || error.message));
      }
    },
    async submitLogistics() {
      try {
        if (!this.logisticsForm.purchase_id) {
          ElMessage.warning('请选择采购记录');
          return;
        }
        // 处理图片URL为数组
        let images = [];
        if (this.logisticsForm.logistics_images) {
          images = this.logisticsForm.logistics_images.split(',').map(item => item.trim()).filter(Boolean);
        }
        const logisticsData = {
          purchase_id: this.logisticsForm.purchase_id,
          logistics_no: this.logisticsForm.logistics_no,
          cold_storage_location: this.logisticsForm.cold_storage_location,
          temperature: this.logisticsForm.temperature !== '' ? parseFloat(this.logisticsForm.temperature) : null,
          humidity: this.logisticsForm.humidity !== '' ? parseFloat(this.logisticsForm.humidity) : null,
          logistics_images: images,
          blockchain_tx_hash: this.logisticsForm.blockchain_tx_hash
        };
        const response = await axios.post(`${serviceEndpoints.distributor}/api/logistics`, logisticsData);
        // 只要状态码为201或200就认为成功
        if (response.status === 201 || response.status === 200) {
          ElMessage.success('物流信息已提交');
          this.clearLogisticsForm();
          this.fetchLogisticsRecords();
          this.logisticsActiveTab = 'list';
        } else {
          ElMessage.error(response.data?.message || '提交物流信息失败');
        }
      } catch (error) {
        console.error('提交物流信息失败:', error);
        ElMessage.error('提交物流信息失败: ' + (error.response?.data?.message || error.message));
      }
    },
    clearLogisticsForm() {
      if (this.$refs.logisticsForm) {
        this.$refs.logisticsForm.resetFields();
      }
    },
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('manufacturer_id');
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
    },
  },
  mounted() {
    this.fetchAvailableProducts();
    this.fetchPurchaseRecords();
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
</style>