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
import { ElMessage } from 'element-plus';
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
      },
      currentPage: 1,
      pageSize: 8,
    };
  },
  computed: {
    paginatedProducts() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.products.slice(start, end);
    },
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
        this.products = response.data.product;
        console.log(this.products);
      } catch (error) {
        ElMessage.error('获取产品失败');
        console.error(error);
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
          },
          batch: {
            current_holder: localStorage.getItem('factoryName'),
            current_location: localStorage.getItem('current_location') || '位置获取失败',
            quantity: this.productForm.quantity,
          },
        };
        console.log("提交的数据", productData);
        const response = await axios.post('http://localhost:8082/products', productData, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        ElMessage.success('产品添加成功');
        this.clearForm(); // 提交成功后清空表单
        this.fetchProducts();
      } catch (error) {
        let errorMessage = '产品添加失败';
        if (error.response && error.response.data && error.response.data.error) {
          errorMessage = error.response.data.error;
        }
        ElMessage.error(errorMessage);
        console.error(error);
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
      };
      sessionStorage.removeItem('product_name');
      sessionStorage.removeItem('brand');
      sessionStorage.removeItem('specification');
      sessionStorage.removeItem('production_date');
      sessionStorage.removeItem('expiration_date');
      sessionStorage.removeItem('raw_material_source');
      sessionStorage.removeItem('processing_method');
      sessionStorage.removeItem('logistics_temp');
      sessionStorage.removeItem('storage_condition');
      sessionStorage.removeItem('qc_report');
      sessionStorage.removeItem('image_url');
      sessionStorage.removeItem('quantity');
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
        callback(new Error('请输入有效的数字'));
      } else {
        callback();
      }
    },
    validateTempMax(rule, value, callback) {
      if (!Number.isInteger(value)) {
        callback(new Error('请输入有效的数字'));
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
      console.log('切换到第', page, '页');
      this.currentPage = page;
    }
  },
  mounted() {
    this.fetchProducts();
  },
};
</script>

<style scoped>
.factory-container {
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

.temp-form-item {
  display: flex;
  align-items: center;
}

.temp-range {
  display: flex;
  align-items: center;
}

.temp-input {
  width: 120px;
  margin-right: 8px;
}

.temp-separator {
  margin: 0 8px;
}

.el-main {
  padding: 20px;
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
</style>    