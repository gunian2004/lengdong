<template>
  <div class="auth-container">
    <el-card class="auth-card">
      <div class="logo">冷冻品溯源登录页面</div>
      <div class="toggle-auth" @click="toggleAuth">
        <el-button v-if="isLogin" type="text">没有账号？去注册</el-button>
        <el-button v-else type="text">已有账号？去登录</el-button>
      </div>
      <el-form v-if="isLogin" ref="loginForm" :model="loginForm" label-width="80px">
        <el-form-item label="账号名">
          <el-input v-model="loginForm.username" placeholder="请输入账号" required></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input type="password" v-model="loginForm.password" placeholder="请输入密码" required></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
      <el-form v-else ref="registerForm" :model="registerForm" label-width="80px">
        <el-form-item label="账号名">
          <el-input v-model="registerForm.username" placeholder="请输入账号名" required></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input type="password" v-model="registerForm.password" placeholder="请输入密码" required></el-input>
        </el-form-item>
        <el-form-item label="身份">
          <el-select v-model="registerForm.role_type" placeholder="请选择身份" required>
            <el-option label="生产厂家" value="factory"></el-option>
            <el-option label="经销商" value="distributor"></el-option>
            <el-option label="零售商" value="retailer"></el-option>
            <el-option label="监管者" value="regulator"></el-option>
            <el-option label="消费者" value="consumer"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="联系方式">
          <el-input v-model="registerForm.contact_info" placeholder="请输入联系方式" required></el-input>
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="registerForm.address" placeholder="请输入地址" required></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import axios from 'axios';
import { ElMessage } from 'element-plus';

export default {
  data() {
    return {
      isLogin: true,
      loginForm: {
        username: '',
        password: '',
      },
      registerForm: {
        username: '',
        password: '',
        role_type: '',
        contact_info: '',
        address: '',
      },
    };
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:8085/login', {
          username: this.loginForm.username,
          password: this.loginForm.password,
        });

        // 存储 token 和用户角色到 localStorage
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('role_type', response.data.user.role_type);
        localStorage.setItem('factoryName', response.data.user.username); // 存储厂家用户名
        localStorage.setItem('distributorName',response.data.user.username)
        localStorage.setItem('current_location', response.data.user.current_location); // 存储当前位置
        localStorage.setItem('manufacturer_id', response.data.user.id);
        ElMessage.success('登录成功');
        console.log(response.data);

        // 根据角色类型跳转到对应的页面
        this.redirectBasedOnRole(response.data.user.role_type);
      } catch (error) {
        let errorMessage = '登录失败';
        if (error.response && error.response.data && error.response.data.error) {
          errorMessage = error.response.data.error;
        }
        ElMessage.error(errorMessage);
        console.error(error);
      }
    },
    async handleRegister() {
      try {
        const response = await axios.post('http://localhost:8085/register', {
          username: this.registerForm.username,
          password: this.registerForm.password,
          role_type: this.registerForm.role_type,
          contact_info: this.registerForm.contact_info,
          address: this.registerForm.address,
        });

        ElMessage.success('注册成功');
        console.log(response.data);

        // 注册成功后，切换到登录并清除注册表单信息
        this.isLogin = true;
        this.registerForm = {
          username: '',
          password: '',
          role_type: '',
          contact_info: '',
          address: '',
        };
      } catch (error) {
        let errorMessage = '注册失败';
        if (error.response && error.response.data && error.response.data.error) {
          errorMessage = error.response.data.error;
        }
        ElMessage.error(errorMessage);
        console.error('注册失败的详细信息:', error);
      }
    },
    toggleAuth() {
      this.isLogin = !this.isLogin;
    },
    redirectBasedOnRole(roleType) {
      switch (roleType) {
        case 'factory':
          this.$router.push('/factory'); // 假设工厂页面的路由是 /factory
          break;
        case 'distributor':
          this.$router.push('/distributor'); // 假设经销商页面的路由是 /distributor
          break;
        case 'retailer':
          this.$router.push('/retailer'); // 假设零售商页面的路由是 /retailer
          break;
        case 'regulator':
          this.$router.push('/regulator'); // 假设监管者页面的路由是 /regulator
          break;
        case 'consumer':
          this.$router.push('/consumer'); // 假设消费者页面的路由是 /consumer
          break;
        case 'super_admin':
          this.$router.push('/dashboard'); //
          break;
        case 'normal_admin':
          this.$router.push('/dashboard'); //
          break;
        default:
          ElMessage.error('未知角色类型');
      }
    },
  },
};
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-image: url('https://example.com/background.jpg'); /* 替换为你的背景图片 */
  background-size: cover;
  background-position: center;
}

.auth-card {
  width: 400px;
  padding: 20px;
  background: white;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.logo {
  font-size: 1.5em;
  margin-bottom: 20px;
}

.toggle-auth {
  margin-bottom: 20px;
}

.el-form-item {
  margin-bottom: 15px;
}

.el-button {
  width: 100%;
}
</style>