<template>
  <div class="admin-dashboard">
    <el-header class="header">
      <div class="header-content">
        <div class="title">后台管理系统</div>
        <div class="welcome">
          欢迎，<span>{{ role_type }}</span><span>{{ factoryName }}</span>
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
          <el-submenu :index="item.index" v-for="item in menuItems" :key="item.index">
            <template #title>
              <i :class="item.icon"></i>
              <span>{{ item.title }}</span>
            </template>
            <el-menu-item :index="subItem.index" v-for="subItem in item.subItems" :key="subItem.index">
              {{ subItem.title }}
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <div class="table-container">
          <!-- 冷冻品表格 -->
          <template v-if="activeIndex === '1-1'">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="product_name" label="产品名称" width="180" />
              <el-table-column prop="brand" label="品牌" width="180" />
              <el-table-column prop="specification" label="规格/包装" width="180" />
              <el-table-column prop="production_date" label="生产日期" width="180" />
              <el-table-column prop="expiration_date" label="保质期" width="180" />
              <el-table-column prop="raw_material_source" label="原材料来源" width="180" />
              <el-table-column prop="processing_method" label="加工方式" width="180" />
              <el-table-column prop="logistics_temp" label="运输温度要求" width="180" />
              <el-table-column prop="storage_condition" label="仓储条件" width="180" />
              <el-table-column prop="qc_report" label="质检报告URL" width="180">
                <template #default="scope">
                  <a :href="scope.row.qc_report" target="_blank" v-if="scope.row.qc_report">
                    查看质检报告
                  </a>
                  <span v-else>质检报告不可用</span>
                </template>
              </el-table-column>
              <el-table-column prop="manufacturer_id" label="生产商ID" width="180" />
              <el-table-column prop="blockchain_tx_hash" label="区块链交易哈希" width="180" />
              <el-table-column prop="audit_status" label="审核状态" width="180" />
              <el-table-column prop="image_url" label="产品图片URL" width="180">
                <template #default="scope">
                  <img :src="scope.row.image_url" alt="产品图片" v-if="scope.row.image_url" style="width: 50px; height: 50px;" />
                  <span v-else>图片不可用</span>
                </template>
              </el-table-column>
              <el-table-column prop="sku" label="SKU" width="180" />
              <el-table-column prop="batch_id" label="批次ID" width="180" />
              <el-table-column prop="batch_number" label="批次号" width="180" />
              <el-table-column prop="current_holder" label="当前持有者" width="180" />
              <el-table-column prop="current_location" label="当前位置" width="180" />
              <el-table-column prop="created_at" label="创建时间" width="180" />
              <el-table-column prop="quantity" label="数量" width="180" />
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[10]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            />
          </template>
          <!-- 用户信息表格 -->
          <template v-else-if="['1-2', '1-3', '1-4', '1-5', '1-6'].includes(activeIndex)">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="id" label="用户ID" width="180" />
              <el-table-column prop="username" label="用户名" width="180" />
              <el-table-column prop="role_type" label="角色类型" width="180" />
              <el-table-column prop="contact_info" label="联系方式" width="180" />
              <el-table-column prop="address" label="地址" width="180" />
              <el-table-column prop="audit_status" label="审核状态" width="180" />
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[10]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            />
          </template>
          <!-- 添加管理员表单 -->
          <template v-else-if="activeIndex === '2-1'">
            <el-form
                :model="adminForm"
                ref="adminForm"
                label-width="120px"
            >
              <el-form-item label="用户名" prop="username">
                <el-input v-model="adminForm.username" placeholder="请输入用户名" />
              </el-form-item>
              <el-form-item label="密码" prop="password">
                <el-input v-model="adminForm.password" placeholder="请输入密码" type="password" />
              </el-form-item>
              <el-form-item label="角色" prop="role">
                <el-select v-model="adminForm.role" placeholder="请选择角色">
                  <el-option label="超级管理员" value="super_admin" />
                  <el-option label="普通管理员" value="normal_admin" />
                </el-select>
              </el-form-item>
              <div class="form-footer">
                <el-button @click="addAdminDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitAdminForm">提交</el-button>
              </div>
            </el-form>
          </template>
          <!-- 权限分配表格 -->
          <template v-else-if="activeIndex === '2-2'">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="admin_id" label="管理员ID" width="180" />
              <el-table-column prop="username" label="管理员" width="180" />
              <el-table-column label="权限" width="180">
                <template #default="scope">
                  <el-select v-model="scope.row.role" placeholder="请选择角色" @change="updateRole(scope.row)">
                    <el-option label="超级管理员" value="super_admin" />
                    <el-option label="普通管理员" value="normal_admin" />
                  </el-select>
                </template>
              </el-table-column>
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[10]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            />
          </template>
          <!-- 用户审核表格 -->
          <template v-else-if="activeIndex === '2-3'">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="id" label="用户ID" width="180" />
              <el-table-column prop="username" label="用户名" width="180" />
              <el-table-column prop="role_type" label="角色类型" width="180" />
              <el-table-column prop="contact_info" label="联系方式" width="180" />
              <el-table-column prop="address" label="地址" width="180" />
              <el-table-column prop="audit_status" label="审核状态" width="180" />
              <el-table-column label="操作" width="180">
                <template #default="scope">
                  <div class="button-group">
                    <el-button class="audit-button" type="success" @click="approveItem(scope.row, 1)">审核通过</el-button>
                    <el-button class="audit-button" type="danger" @click="approveItem(scope.row, 0)">审核不通过</el-button>
                  </div>
                </template>
              </el-table-column>
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[10]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            />
          </template>
          <!-- 产品审核表格 -->
          <template v-else-if="activeIndex === '2-4'">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="product_name" label="产品名称" width="180" />
              <el-table-column prop="brand" label="品牌" width="180" />
              <el-table-column prop="specification" label="规格/包装" width="180" />
              <el-table-column prop="production_date" label="生产日期" width="180" />
              <el-table-column prop="expiration_date" label="保质期" width="180" />
              <el-table-column prop="raw_material_source" label="原材料来源" width="180" />
              <el-table-column prop="processing_method" label="加工方式" width="180" />
              <el-table-column prop="logistics_temp" label="运输温度要求" width="180" />
              <el-table-column prop="storage_condition" label="仓储条件" width="180" />
              <el-table-column prop="qc_report" label="质检报告URL" width="180">
                <template #default="scope">
                  <a :href="scope.row.qc_report" target="_blank" v-if="scope.row.qc_report">
                    查看质检报告
                  </a>
                  <span v-else>质检报告不可用</span>
                </template>
              </el-table-column>
              <el-table-column prop="manufacturer_id" label="生产商ID" width="180" />
              <el-table-column prop="blockchain_tx_hash" label="区块链交易哈希" width="180" />
              <el-table-column prop="audit_status" label="审核状态" width="180" />
              <el-table-column prop="image_url" label="产品图片URL" width="180">
                <template #default="scope">
                  <img :src="scope.row.image_url" alt="产品图片" v-if="scope.row.image_url" style="width: 50px; height: 50px;" />
                  <span v-else>图片不可用</span>
                </template>
              </el-table-column>
              <el-table-column prop="sku" label="SKU" width="180" />
              <el-table-column prop="batch_id" label="批次ID" width="180" />
              <el-table-column prop="batch_number" label="批次号" width="180" />
              <el-table-column prop="current_holder" label="当前持有者" width="180" />
              <el-table-column prop="current_location" label="当前位置" width="180" />
              <el-table-column prop="created_at" label="创建时间" width="180" />
              <el-table-column prop="quantity" label="数量" width="180" />
              <el-table-column label="操作" width="180">
                <template #default="scope">
                  <div class="button-group">
                    <el-button class="audit-button" type="success" @click="approveItem(scope.row, 1)">审核通过</el-button>
                    <el-button class="audit-button" type="danger" @click="approveItem(scope.row, 0)">审核不通过</el-button>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </template>
          <!-- 审核记录表格 -->
          <template v-else-if="activeIndex === '2-5'">
            <el-input
                class="input-fixed-top"
                v-model="searchQuery"
                placeholder="请输入搜索内容"
                style="width: 300px;"
            />
            <el-table
                :data="filteredData"
                style="width: 100%; table-layout: fixed;"
                border
                stripe
            >
              <el-table-column prop="audit_id" label="审核ID" width="180" />
              <el-table-column prop="target_type" label="审核类型" width="180" />
              <el-table-column prop="target_id" label="目标ID" width="180" />
              <el-table-column prop="audit_result" label="审核结果" width="180" />
              <el-table-column prop="audit_comment" label="审核意见" width="180" />
              <el-table-column prop="auditor_id" label="审核人ID" width="180" />
              <el-table-column prop="audit_time" label="审核时间" width="180">
                <template #default="scope">
                  {{ new Date(scope.row.audit_time).toLocaleString() }}
                </template>
              </el-table-column>
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[10]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
            />
          </template>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import axios from 'axios';
import { ElMessage } from 'element-plus';

export default {
  data() {
    return {
      factoryName: localStorage.getItem('factoryName') || '某某厂家',
      role_type: this.translate(localStorage.getItem('role_type') || '管理员'),
      activeIndex: '',
      menuItems: [
        {
          index: '1',
          title: '查看信息',
          icon: 'el-icon-menu',
          subItems: [
            { index: '1-1', title: '查看冷冻品', method: this.fetchProduct },
            { index: '1-2', title: '查看供应商', method: this.fetchFactory },
            { index: '1-3', title: '查看经销商', method: this.fetchDistributor },
            { index: '1-4', title: '查看零售商', method: this.fetchRetailer },
            { index: '1-5', title: '查看监管机构', method: this.fetchRegulator },
            { index: '1-6', title: '查看消费者', method: this.fetchConsumer },
          ],
        },
        {
          index: '2',
          title: '人员管理',
          icon: 'el-icon-document',
          subItems: [
            { index: '2-1', title: '添加管理员', method: this.openAddAdminDialog },
            { index: '2-2', title: '权限分配', method: this.fetchAdmins },
            { index: '2-3', title: '用户审核', method: this.fetchPendingUsers },
            { index: '2-4', title: '产品审核', method: this.fetchPendingProducts },
            { index: '2-5', title: '审核记录', method: this.fetchAuditLogs },
          ],
        },
      ],
      tableData: [],
      adminForm: {
        username: '',
        password: '',
        role: ''
      },
      currentPage: 1,
      pageSize: 10,
      total: 0,
      searchQuery: ''
    };
  },
  methods: {
    handleSelect(index) {
      if (this.activeIndex === index) {
        this.activeIndex = '';
        this.tableData = [];
      } else {
        this.activeIndex = index;
        const selectedItem = this.findMenuItemByIndex(index);
        if (selectedItem && selectedItem.method) {
          selectedItem.method();
        }
      }
    },
    findMenuItemByIndex(index) {
      for (const item of this.menuItems) {
        for (const subItem of item.subItems) {
          if (subItem.index === index) {
            return subItem;
          }
        }
      }
      return null;
    },
    async fetchProduct() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_product', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.product) {
          this.total = response.data.product.length;
          this.tableData = response.data.product.map(product => {
            const batches = product.ProductBatches || [];
            return {
              ...product,
              ...batches[0],
              audit_status: this.translate(product.audit_status),
              current_holder: this.translate(product.current_holder)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching product data:', error);
        this.tableData = [];
      }
    },
    async fetchFactory() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_factory', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.user) {
          this.total = response.data.user.length;
          this.tableData = response.data.user.map(user => {
            return {
              ...user,
              audit_status: this.translate(user.audit_status),
              role_type: this.translate(user.role_type)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching factory data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async fetchDistributor() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_distributor', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.distributor) {
          this.total = response.data.distributor.length;
          this.tableData = response.data.distributor.map(distributor => {
            return {
              ...distributor,
              audit_status: this.translate(distributor.audit_status),
              role_type: this.translate(distributor.role_type)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching distributor data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async fetchRetailer() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_retailer', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.retailer) {
          this.total = response.data.retailer.length;
          this.tableData = response.data.retailer.map(retailer => {
            return {
              ...retailer,
              audit_status: this.translate(retailer.audit_status),
              role_type: this.translate(retailer.role_type)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching retailer data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async fetchRegulator() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_regulator', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.regulator) {
          this.total = response.data.regulator.length;
          this.tableData = response.data.regulator.map(regulator => {
            return {
              ...regulator,
              audit_status: this.translate(regulator.audit_status),
              role_type: this.translate(regulator.role_type)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching regulator data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async fetchConsumer() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_consumer', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.consumer) {
          this.total = response.data.consumer.length;
          this.tableData = response.data.consumer.map(consumer => {
            return {
              ...consumer,
              audit_status: this.translate(consumer.audit_status),
              role_type: this.translate(consumer.role_type)
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching consumer data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async submitAdminForm() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.post('http://localhost:8080/add_admin', {
          username: this.adminForm.username,
          password: this.adminForm.password,
          role: this.adminForm.role
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        ElMessage.success('管理员添加成功');
        this.adminForm = {
          username: '',
          password: '',
          role: ''
        };
        this.activeIndex = '';
      } catch (error) {
        console.error('Error adding admin:', error);
        ElMessage.error('管理员添加失败');
      }
    },
    async fetchAdmins() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_admin', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.admin) {
          this.total = response.data.admin.length;
          this.tableData = response.data.admin;
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching admins data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async updateRole(admin) {
      try {
        const token = localStorage.getItem('token');
        const currentRole = this.translate(localStorage.getItem('role_type') || '管理员');
        if (currentRole !== '超级管理员') {
          ElMessage.error('您无权限修改权限');
          return;
        }
        const response = await axios.post('http://localhost:8080/update_role', {
          username: admin.username,
          role: admin.role
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        ElMessage.success('权限更新成功');
      } catch (error) {
        console.error('Error updating role:', error);
        ElMessage.error('权限更新失败');
      }
    },
    async fetchPendingUsers() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_user', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.user) {
          this.total = response.data.user.length;
          this.tableData = response.data.user
              .filter(user => user.audit_status === 'pending')
              .map(user => {
                return {
                  ...user,
                  audit_status: this.translate(user.audit_status),
                  role_type: this.translate(user.role_type)
                };
              });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching pending users data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    async fetchPendingProducts() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_product', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.product) {
          this.total = response.data.product.length;
          this.tableData = response.data.product
              .filter(product => product.audit_status === 'pending')
              .map(product => {
                const batches = product.ProductBatches || [];
                return {
                  ...product,
                  ...batches[0],
                  audit_status: this.translate(product.audit_status),
                  current_holder: this.translate(product.current_holder)
                };
              });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching product data:', error);
        this.tableData = [];
      }
    },
    async approveItem(item, isApproved) {
      try {
        const token = localStorage.getItem('token');
        let url = '';
        let data = {};
        let auditType = '';

        if (this.activeIndex === '2-3') {
          url = 'http://localhost:8080/update_user_pending';
          data = {
            username: item.username,
            is_success: isApproved ? 'approved' : 'rejected',
            audit_comment: '审核意见',
            auditor_id: '1'
          };
          auditType = '用户';
        } else if (this.activeIndex === '2-4') {
          url = 'http://localhost:8080/update_product_pending';
          data = {
            product_name: item.product_name,
            is_success: isApproved ? 'approved' : 'rejected',
            audit_comment: '审核意见',
            auditor_id: '1'
          };
          auditType = '产品';
        }

        const response = await axios.post(url, data, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        ElMessage.success(`${auditType}审核成功`);
        if (this.activeIndex === '2-3') {
          await this.fetchPendingUsers();
        } else if (this.activeIndex === '2-4') {
          await this.fetchPendingProducts();
        }
      } catch (error) {
        let errorMessage = '审核失败';
        if (error.response && error.response.data && error.response.data.error) {
          errorMessage = error.response.data.error;
        }
        ElMessage.error(errorMessage);
        console.error(error);
      }
    },
    async fetchAuditLogs() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/view_audit_log', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log('Response data:', response.data);
        if (response.data && response.data.auditlog) {
          this.total = response.data.auditlog.length;
          this.tableData = response.data.auditlog.map(auditlog => {
            return {
              ...auditlog,
              target_type: this.translate(auditlog.target_type),
              audit_result: this.translate(auditlog.audit_result),
              audit_time: new Date(auditlog.audit_time).toLocaleString()
            };
          });
        } else {
          this.tableData = [];
          this.total = 0;
        }
      } catch (error) {
        console.error('Error fetching audit logs data:', error);
        this.tableData = [];
        this.total = 0;
      }
    },
    handleSizeChange(newSize) {
      this.pageSize = newSize;
      this.fetchData();
    },
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.fetchData();
    },
    fetchData() {
      switch (this.activeIndex) {
        case '1-1':
          this.fetchProduct();
          break;
        case '1-2':
          this.fetchFactory();
          break;
        case '1-3':
          this.fetchDistributor();
          break;
        case '1-4':
          this.fetchRetailer();
          break;
        case '1-5':
          this.fetchRegulator();
          break;
        case '1-6':
          this.fetchConsumer();
          break;
        case '2-2':
          this.fetchAdmins();
          break;
        case '2-3':
          this.fetchPendingUsers();
          break;
        case '2-4':
          this.fetchPendingProducts();
          break;
        case '2-5':
          this.fetchAuditLogs();
          break;
        default:
          this.tableData = [];
          this.total = 0;
      }
    },
    translate(text) {
      const translations = {
        'super_admin': '超级管理员',
        'normal_admin': '普通管理员',
        'user': '用户',
        'product': '产品',
        'approved': '已批准',
        'rejected': '已拒绝',
        'pending': '待审核',
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
    },
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('role_type');
      localStorage.removeItem('factoryName');
      localStorage.removeItem('current_location');
      localStorage.removeItem('manufacturer_id');
      this.$router.push('/');
    },
    openAddAdminDialog() {
      this.addAdminDialogVisible = true;
    }
  },
  computed: {
    filteredData() {
      if (!this.searchQuery) {
        return this.tableData;
      }
      const query = this.searchQuery.toLowerCase();
      return this.tableData.filter(item => {
        return Object.values(item).some(value => {
          return String(value).toLowerCase().includes(query);
        });
      });
    }
  },
  created() {
    // 页面加载时初始化
    this.fetchData();
  }
};
</script>

<style scoped>
.admin-dashboard {
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
</style>