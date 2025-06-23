<template>
  <div class="factory-container">
    <el-card class="factory-card">
      <div class="logo">生产厂家页面</div>
      <div class="content">
        <p>欢迎，生产厂家！</p>
        <p>这里可以展示与生产厂家相关的功能和数据。</p>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'Retailer',
};
</script>

<style scoped>
.factory-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-image: url('https://example.com/background.jpg'); /* 替换为你的背景图片 */
  background-size: cover;
  background-position: center;
}

.factory-card {
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

.content {
  margin-top: 20px;
}
</style>