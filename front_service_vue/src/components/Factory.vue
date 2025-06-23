<template>
  <div class="factory-container">
    <el-header class="header">
      <div class="header-content">
        <div class="title">冷链溯源系统-生产厂家页</div>
        <div class="welcome">
          欢迎，<span>{{ factoryName }}</span>
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
            <span>我的冷冻品</span>
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-plus"></i>
            <span>添加冷冻品</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <div v-if="activeIndex === '1'">
          <h2>我的冷冻品</h2>
          <el-table :data="paginatedProducts" style="width: 100%">
            <el-table-column prop="product_name" label="产品名称" width="180"></el-table-column>
            <el-table-column prop="brand" label="品牌" width="180"></el-table-column>
            <el-table-column prop="specification" label="规格/包装" width="180"></el-table-column>
            <el-table-column prop="production_date" label="生产日期" width="180"></el-table-column>
            <el-table-column prop="expiration_date" label="保质期" width="180"></el-table-column>
            <el-table-column prop="raw_material_source" label="原材料来源" width="180"></el-table-column>
            <el-table-column prop="processing_method" label="加工方式" width="180"></el-table-column>
            <el-table-column prop="logistics_temp" label="运输温度要求" width="180"></el-table-column>
            <el-table-column prop="storage_condition" label="仓储条件" width="180"></el-table-column>
            <el-table-column prop="qc_report" label="质检报告" width="180">
              <template #default="scope">
                <a :href="scope.row.qc_report" target="_blank" v-if="scope.row.qc_report">
                  查看质检报告
                </a>
                <span v-else>质检报告不可用</span>
              </template>
            </el-table-column>
            <el-table-column prop="manufacturer_id" label="生产商ID" width="180"></el-table-column>
            <el-table-column prop="image_url" label="产品图片" width="180">
              <template #default="scope">
                <el-image
                    style="width: 50px; height: 50px"
                    :src="scope.row.image_url"
                    :preview-src-list="[scope.row.image_url]"
                    fit="cover"
                    :preview-teleported="true"
                />
              </template>
            </el-table-column>
            <el-table-column prop="current_holder" label="当前持有者" width="180"></el-table-column>
            <el-table-column prop="current_location" label="当前位置" width="180"></el-table-column>
            <el-table-column prop="quantity" label="数量" width="180"></el-table-column>
            <!-- 新增SKU列 -->
            <el-table-column prop="sku" label="SKU" width="180"></el-table-column>
            <!-- 审核状态列 -->
            <el-table-column prop="audit_status" label="审核状态" width="180">
              <template #default="scope">
                <el-tag :type="getStatusType(scope.row.audit_status)">
                  {{ getStatusText(scope.row.audit_status) }}
                </el-tag>
              </template>
            </el-table-column>
            <!-- 上链状态列 -->
            <el-table-column label="上链状态" width="180">
              <template #default="scope">
                <el-tag :type="getChainStatusType(scope.row.sku)">
                  {{ getChainStatusText(scope.row.sku) }}
                </el-tag>
              </template>
            </el-table-column>
            <!-- 区块链交易哈希列 -->
            <el-table-column prop="blockchain_tx_hash" label="区块链交易哈希" width="300">
              <template #default="scope">
                <span v-if="scope.row.blockchain_tx_hash">
                  <el-tooltip content="点击复制" placement="top">
                    <span class="tx-hash" @click="copyToClipboard(scope.row.blockchain_tx_hash)">{{ scope.row.blockchain_tx_hash.slice(0, 10) }}...</span>
                  </el-tooltip>
                </span>
                <span v-else>无</span>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[8, 16, 24]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="products.length"
          />
        </div>
        <div v-else-if="activeIndex === '2'">
          <h2>添加冷冻品</h2>
          <el-form :model="productForm" ref="productForm" label-width="120px">
            <el-form-item label="产品名称" prop="product_name">
              <el-input v-model="productForm.product_name" placeholder="请输入产品名称" required></el-input>
            </el-form-item>
            <el-form-item label="品牌" prop="brand">
              <el-input v-model="productForm.brand" placeholder="请输入品牌" required></el-input>
            </el-form-item>
            <el-form-item label="规格/包装" prop="specification">
              <el-input v-model="productForm.specification" placeholder="请输入规格/包装" required></el-input>
            </el-form-item>
            <el-form-item label="生产日期" prop="production_date">
              <el-date-picker v-model="productForm.production_date" type="date" placeholder="选择日期" required></el-date-picker>
            </el-form-item>
            <el-form-item label="保质期" prop="expiration_date">
              <el-date-picker v-model="productForm.expiration_date" type="date" placeholder="选择日期" required></el-date-picker>
            </el-form-item>
            <el-form-item label="原材料来源" prop="raw_material_source">
              <el-input v-model="productForm.raw_material_source" placeholder="请输入原材料来源" required></el-input>
            </el-form-item>
            <el-form-item label="加工方式" prop="processing_method">
              <el-input v-model="productForm.processing_method" placeholder="请输入加工方式" required></el-input>
            </el-form-item>
            <el-form-item label="运输温度要求" prop="logistics_temp" class="temp-form-item">
              <div class="temp-range">
                <el-input
                    v-model.number="productForm.logistics_temp_min"
                    placeholder="下限温度"
                    class="temp-input"
                    @input="updateTemperatureRange"
                    :rules="[
                        { required: true, message: '请输入下限温度', trigger: 'blur' },
                        { validator: validateTempMin, trigger: 'blur' }
                     ]"
                >
                  <template #suffix>°C</template>
                </el-input>
                <span class="temp-separator">-</span>
                <el-input
                    v-model.number="productForm.logistics_temp_max"
                    placeholder="上限温度"
                    class="temp-input"
                    @input="updateTemperatureRange"
                    :rules="[
          { required: true, message: '请输入上限温度', trigger: 'blur' },
          { validator: validateTempMax, trigger: 'blur' }
        ]"
                >
                  <template #suffix>°C</template>
                </el-input>
              </div>
            </el-form-item>
            <el-form-item label="仓储条件" prop="storage_condition">
              <el-input v-model="productForm.storage_condition" placeholder="请输入仓储条件" required></el-input>
            </el-form-item>
            <el-form-item label="质检报告" prop="qc_report">
              <el-input v-model="productForm.qc_report" placeholder="请输入质检报告URL"></el-input>
            </el-form-item>
            <el-form-item label="产品图片" prop="image_url">
              <image-upload @upload-success="handleImageUploadSuccess" />
            </el-form-item>
            <el-form-item label="数量" prop="quantity">
              <el-input v-model="productForm.quantity" placeholder="请输入数量" required></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitProduct">提交</el-button>
              <el-button type="danger" @click="clearForm">清空</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import axios from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import ImageUpload from './index.vue';

export default {
  components: {
    ImageUpload,
  },
  data() {
    return {
      activeIndex: '1',
      factoryName: localStorage.getItem('factoryName') || '某某厂家',
      products: [],
      productForm: {
        product_name: '',
        brand: '',
        specification: '',
        production_date: '',
        expiration_date: '',
        raw_material_source: '',
        processing_method: '',
        logistics_temp: '',
        storage_condition: '',
        qc_report: '',
        image_url: '',
        quantity: '',
        logistics_temp_min: null,
        logistics_temp_max: null,
        audit_status: 'pending',
      },
      currentPage: 1,
      pageSize: 8,
      loadingMap: {},       // 加载状态映射
      chainCalledMap: {},   // 记录已调用接口的SKU
      chainStatusMap: {}    // 记录上链状态
    };
  },
  computed: {
    paginatedProducts() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.products.slice(start, end);
    }
  },
  methods: {
    handleSelect(index) {
      this.activeIndex = index;
    },
    async fetchProducts() {
      try {
        const manufacturer_id = localStorage.getItem('manufacturer_id') || '0';
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8082/view_product', {
          params: {
            manufacturer_id: manufacturer_id
          },
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('获取产品列表响应:', response.data);

        if (response.data && response.data.product) {
          // 合并产品和批次数据，确保获取完整信息
          this.products = response.data.product.map(product => {
            const batches = product.ProductBatches || [];
            return {
              ...product,
              ...batches[0],
              audit_status: product.audit_status, // 直接使用后端状态
              current_holder: this.translate(product.current_holder)
            };
          });
        } else {
          this.products = [];
        }

        // 强制刷新上链状态
        this.chainStatusMap = {};
        this.checkAndCallChainAPI();
      } catch (error) {
        ElMessage.error('获取产品失败');
        console.error('获取产品列表错误:', error);
        this.products = [];
      }
    },
    async submitProduct() {
      if (!this.productForm.logistics_temp) {
        ElMessage.error('请输入完整的运输温度要求');
        return;
      }
      try {
        const token = localStorage.getItem('token');
        const productionDate = new Date(this.productForm.production_date).toISOString();
        const expirationDate = new Date(this.productForm.expiration_date).toISOString();

        const productData = {
          product: {
            product_name: this.productForm.product_name,
            brand: this.productForm.brand,
            specification: this.productForm.specification,
            production_date: productionDate,
            expiration_date: expirationDate,
            raw_material_source: this.productForm.raw_material_source,
            processing_method: this.productForm.processing_method,
            logistics_temp: this.productForm.logistics_temp,
            storage_condition: this.productForm.storage_condition,
            qc_report: this.productForm.qc_report,
            manufacturer_id: localStorage.getItem('manufacturer_id'),
            image_url: sessionStorage.getItem('uploadedFileName'),
            audit_status: 'pending',
          },
          batch: {
            current_holder: localStorage.getItem('factoryName'),
            current_location: localStorage.getItem('current_location') || '位置获取失败',
            quantity: this.productForm.quantity,
          },
        };
        console.log("提交的产品数据:", productData);
        const response = await axios.post('http://localhost:8082/products', productData, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        ElMessage.success('产品添加成功');
        this.clearForm(); // 提交成功后清空表单
        this.fetchProducts(); // 刷新产品列表
      } catch (error) {
        let errorMessage = '产品添加失败';
        if (error.response && error.response.data && error.response.data.error) {
          errorMessage = error.response.data.error;
        }
        ElMessage.error(errorMessage);
        console.error('添加产品错误:', error);
      }
    },
    clearForm() {
      this.productForm = {
        product_name: '',
        brand: '',
        specification: '',
        production_date: '',
        expiration_date: '',
        raw_material_source: '',
        processing_method: '',
        logistics_temp: '',
        storage_condition: '',
        qc_report: '',
        image_url: '',
        quantity: '',
        logistics_temp_min: null,
        logistics_temp_max: null,
        audit_status: 'pending',
      };
      sessionStorage.removeItem('uploadedFileName');
    },
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('role_type');
      localStorage.removeItem('factoryName');
      localStorage.removeItem('current_location');
      localStorage.removeItem('manufacturer_id');
      this.$router.push('/');
    },
    updateTemperatureRange() {
      if (this.productForm.logistics_temp_min !== null && this.productForm.logistics_temp_max !== null) {
        this.productForm.logistics_temp = `${this.productForm.logistics_temp_min}°C - ${this.productForm.logistics_temp_max}°C`;
      } else {
        this.productForm.logistics_temp = '';
      }
    },
    validateTempMin(rule, value, callback) {
      if (!Number.isInteger(value)) {
        callback(new Error('请输入有效的整数'));
      } else {
        callback();
      }
    },
    validateTempMax(rule, value, callback) {
      if (!Number.isInteger(value)) {
        callback(new Error('请输入有效的整数'));
      } else if (this.productForm.logistics_temp_min !== null && value < this.productForm.logistics_temp_min) {
        callback(new Error('上限温度不能低于下限温度'));
      } else {
        callback();
      }
    },
    handleSizeChange(size) {
      this.pageSize = size;
      this.currentPage = 1;
    },
    handleCurrentChange(page) {
      this.currentPage = page;
    },
    getStatusType(status) {
      const typeMap = {
        'pending': 'warning',
        'approved': 'success',
        'rejected': 'danger',
      };
      return typeMap[status] || 'info';
    },
    getStatusText(status) {
      const textMap = {
        'pending': '待审核',
        'approved': '已上链',
        'rejected': '已拒绝',
      };
      return textMap[status] || status;
    },
    getChainStatusType(sku) {
      const status = this.chainStatusMap[sku];
      if (status === 'loading') return 'info';
      if (status === 'success') return 'success';
      if (status === 'failed') return 'danger';
      return 'info';
    },
    getChainStatusText(sku) {
      const status = this.chainStatusMap[sku];
      if (status === 'loading') return '上链中...';
      if (status === 'success') return '上链成功';
      if (status === 'failed') return '上链失败';

      // 检查后端是否已记录区块链交易哈希
      const product = this.products.find(p => p.sku === sku);
      if (product && product.blockchain_tx_hash) {
        return '上链成功';
      }

      return '未上链';
    },
    handleImageUploadSuccess(fileName) {
      this.productForm.image_url = fileName;
      sessionStorage.setItem('uploadedFileName', fileName);
    },
    checkAndCallChainAPI() {
      console.log('检查并调用上链API...');

      // 遍历产品列表，检查需要上链的产品
      for (const product of this.products) {
        const sku = product.sku;

        // 跳过SKU为空的产品
        if (!sku) {
          console.warn('产品SKU为空，跳过调用:', product);
          this.chainStatusMap[sku] = 'failed';
          continue;
        }

        // 直接使用后端返回的审核状态
        const auditStatus = product.audit_status;

        // 检查是否是已上链状态（后端返回'approved'表示已上链）
        const isApproved = auditStatus === 'approved';
        const hasTxHash = product.blockchain_tx_hash;

        // 更新上链状态显示
        this.updateChainStatus(sku, auditStatus, hasTxHash);

        // 只处理审核状态为"已上链"且未调用过接口、没有交易哈希的产品
        if (isApproved && !this.chainCalledMap[sku] && !hasTxHash) {
          console.log(`产品 ${sku} 状态为已上链，准备调用上链API`);
          this.callChainAPI(sku);
        } else if (isApproved && hasTxHash) {
          // 已有交易哈希，直接标记为上链成功
          this.chainStatusMap[sku] = 'success';
          this.chainCalledMap[sku] = true;
          console.log(`产品 ${sku} 已有交易哈希，已上链`);
        } else {
          console.log(`产品 ${sku} 不需要上链: 状态=${auditStatus}, 是否已调用=${this.chainCalledMap[sku]}, 交易哈希=${hasTxHash ? '存在' : '不存在'}`);
        }
      }
    },
    async callChainAPI(sku) {
      // 标记为已调用
      this.chainCalledMap[sku] = true;
      this.chainStatusMap[sku] = 'loading';
      this.loadingMap[sku] = true;

      try {
        const token = localStorage.getItem('token');
        if (!token) {
          this.chainStatusMap[sku] = 'failed';
          this.loadingMap[sku] = false;
          ElMessage.error('请先登录以进行上链操作');
          return;
        }

        console.log(`调用产品 ${sku} 的上链API`);

        // 调用后端/chain接口
        const response = await axios.post(
            'http://localhost:8082/chain',
            {},
            {
              params: { sku },
              headers: {
                Authorization: `Bearer ${token}`
              }
            }
        );

        console.log(`产品 ${sku} 上链API响应:`, response.data);

        if (response.data.success) {
          // 更新上链状态和区块链交易哈希
          this.updateProductChainStatus(sku, response.data.txHash);
          this.chainStatusMap[sku] = 'success';
          ElMessage.success(`产品 ${sku} 已自动上链`);
        } else {
          this.chainStatusMap[sku] = 'failed';
          ElMessage.error(`产品 ${sku} 上链失败: ${response.data.message || '未知错误'}`);
        }
      } catch (error) {
        this.chainStatusMap[sku] = 'failed';
        console.error(`产品 ${sku} 上链接口调用失败`, error);
        ElMessage.error(`产品 ${sku} 上链失败: ${error.message}`);
      } finally {
        this.loadingMap[sku] = false;
      }
    },
    updateProductChainStatus(sku, txHash) {
      // 找到对应的产品并更新区块链交易哈希
      const productIndex = this.products.findIndex(p => p.sku === sku);
      if (productIndex !== -1) {
        this.products[productIndex].blockchain_tx_hash = txHash;
      }
    },
    updateChainStatus(sku, auditStatus, hasTxHash) {
      // 根据审核状态和交易哈希设置上链状态
      if (auditStatus === 'approved') {
        if (hasTxHash) {
          this.chainStatusMap[sku] = 'success';
        } else {
          // 已审核但无交易哈希，可能是上链中或失败
          this.chainStatusMap[sku] = this.chainCalledMap[sku] ? 'failed' : 'unknown';
        }
      } else {
        this.chainStatusMap[sku] = 'not_approved';
      }
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        ElMessage.success('交易哈希已复制到剪贴板');
      }).catch(err => {
        console.error('无法复制: ', err);
        ElMessage.error('复制失败');
      });
    },
    translate(text) {
      const translations = {
        'super_admin': '超级管理员',
        'normal_admin': '普通管理员',
        'user': '用户',
        'product': '产品',
        'factory': '工厂',
        'distributor': '经销商',
        'retailer': '零售商',
        'regulator': '监管机构',
        'consumer': '消费者',
        'Total': '总计',
        'page': '页',
        'Go to': '跳转到'
      };
      return translations[text] || text;
    }
  },
  mounted() {
    this.fetchProducts();

    // 添加定时刷新，确保状态变化能被检测到
    this.refreshInterval = setInterval(() => {
      this.fetchProducts();
    }, 10000); // 每10秒刷新一次
  },
  beforeDestroy() {
    // 清理定时器
    if (this.refreshInterval) {
      clearInterval(this.refreshInterval);
    }
  },
  watch: {
    // 监听产品列表变化，确保状态变化能被检测到
    products: {
      handler: 'checkAndCallChainAPI',
      deep: true,
      immediate: true
    }
  }
};
</script>

<style scoped>
.factory-container {
  height: 100vh;
}

.header {
  background-color: #409EFF;
  color: white;
  text-align: center;
  line-height: 60px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.title {
  font-size: 1.5em;
  font-weight: bold;
}

.welcome {
  display: flex;
  align-items: center;
  gap: 10px;
}

.el-aside {
  background-color: #545c64;
  color: #fff;
}

.el-menu-vertical-demo {
  border-right: none;
}

.el-main {
  background-color: #f5f7fa;
  padding: 20px;
}

.table-container {
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
  padding: 20px;
  height: calc(100vh - 140px);
  overflow: auto;
}

.input-fixed-top {
  margin-bottom: 20px;
}

.el-table {
  margin-top: 10px;
}

.el-table th {
  background-color: #f5f7fa;
  font-weight: bold;
}

.el-pagination {
  margin-top: 20px;
  text-align: right;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 10px;
}

.audit-button {
  padding: 8px 12px;
}

.form-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

/* 表格单元格样式优化 */
.el-table__cell {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 确保表格数据与表头对齐的关键样式 */
.el-table--scrollable-x .el-table__body-wrapper {
  overflow-x: auto;
}

/* 优化滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 温度输入框样式 */
.temp-form-item {
  display: flex;
  align-items: center;
}

.temp-range {
  display: flex;
  align-items: center;
  width: 100%;
}

.temp-input {
  flex: 1;
  margin: 0 5px;
}

.temp-separator {
  padding: 0 5px;
}

/* 交易哈希样式 */
.tx-hash {
  cursor: pointer;
  color: #409EFF;
  text-decoration: underline;
}
</style>