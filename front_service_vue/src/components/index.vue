<template>
  <!-- 上传图片组件 -->
  <div>
    <!-- 图片上传表单，阻止页面刷新 -->
    <form>
      <!-- 文件选择框，仅接受图片文件，文件变化时自动上传 -->
      <input type="file" name="image" accept="image/*" @change="handleFileChange" />
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// 使用 ref 创建响应式数据，保存上传的文件对象
const file = ref(null);

/**
 * 处理文件选择框的变化事件
 * @param {Event} event - 文件选择框的 change 事件
 */
const handleFileChange = async (event) => {
  // 获取用户选择的第一个文件
  file.value = event.target.files[0];

  // 检查是否有文件被选择
  if (!file.value) {
    alert('请选择一个文件进行上传。');
    return;
  }

  // 获取文件名并存储到 sessionStorage
  const fileName = file.value.name;
  sessionStorage.setItem('uploadedFileName', fileName);

  // 创建FormData对象用于构造请求数据
  const formData = new FormData();
  // 将文件添加到FormData中
  formData.append('image', file.value);

  // 从 localStorage 中获取 token
  const token = localStorage.getItem('token');

  // 使用try-catch处理异步请求过程中的异常
  try {
    // 发起POST请求上传文件
    const response = await fetch('http://localhost:8082/upload', {
      method: 'POST',
      body: formData,
      headers: {
        Authorization: `Bearer ${token}` // 添加 token 到请求头
      }
    });

    // 检查响应状态
    if (!response.ok) {
      throw new Error('上传失败');
    }

    // 解析响应数据
    const data = await response.json();
    console.log('成功:', data);
    alert('图片上传成功');
  } catch (error) {
    console.error('错误:', error);
    alert('图片上传失败');
    // 清除选择的文件
    file.value = null;
    // 清除文件选择框的值
    event.target.value = null;
    // 清除 sessionStorage 中的文件名
    sessionStorage.removeItem('uploadedFileName');
  }
};
</script>