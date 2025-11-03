<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <div class="header-icon">
            <svg width="32" height="32" viewBox="0 0 24 24" fill="#ff6b9d">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </div>
          <h2>欢迎回来</h2>
          <p class="header-subtitle">登录你的账号</p>
        </div>
      </template>
      <el-form :model="loginForm" :rules="rules" ref="formRef" label-width="80px" class="login-form">
        <el-form-item label="用户名" prop="username">
          <el-input 
            v-model="loginForm.username" 
            placeholder="请输入用户名"
            size="large"
          >
            <template #prefix>
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#ff6b9d" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input 
            v-model="loginForm.password" 
            type="password" 
            placeholder="请输入密码" 
            show-password
            size="large"
          >
            <template #prefix>
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#ff6b9d" stroke-width="2">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
              </svg>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button 
            type="primary" 
            @click="handleLogin" 
            :loading="loading" 
            size="large"
            class="submit-button"
          >
            登录
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button 
            text 
            @click="router.push('/register')" 
            class="link-button"
          >
            还没有账号？<span class="link-text">立即注册</span>
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { authAPI } from '@/api/auth'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const formRef = ref()
const loading = ref(false)

const loginForm = ref({
  username: '',
  password: ''
})

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少6位', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  try {
    await formRef.value.validate()
    loading.value = true
    const res = await authAPI.login(loginForm.value)
    userStore.setToken(res.token)
    userStore.setUser(res.user)
    ElMessage.success('登录成功')
    router.push('/')
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  min-height: calc(100vh - 140px);
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #ffeef8 0%, #fff5e6 50%, #ffe5e5 100%);
  padding: 40px 20px;
}

.login-card {
  width: 100%;
  max-width: 480px;
  border-radius: 24px;
  border: 1.5px solid rgba(255, 182, 193, 0.3);
  box-shadow: 0 8px 32px rgba(255, 182, 193, 0.2);
  overflow: hidden;
}

.card-header {
  text-align: center;
  padding: 40px 20px 30px;
  background: linear-gradient(135deg, rgba(255, 182, 193, 0.1) 0%, rgba(255, 218, 185, 0.1) 100%);
}

.header-icon {
  display: flex;
  justify-content: center;
  margin-bottom: 16px;
}

.card-header h2 {
  font-size: 32px;
  font-weight: 600;
  color: #1d1d1f;
  margin: 0 0 8px 0;
  letter-spacing: -0.011em;
}

.header-subtitle {
  font-size: 15px;
  color: #86868b;
  margin: 0;
  font-weight: 400;
}

.login-form {
  padding: 30px 40px 40px;
}

.submit-button {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  border-radius: 980px;
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border: none;
  box-shadow: 0 4px 16px rgba(255, 107, 157, 0.3);
  transition: all 0.3s;
}

.submit-button:hover {
  background: linear-gradient(135deg, #ff8fab 0%, #ff6b9d 100%);
  box-shadow: 0 6px 20px rgba(255, 107, 157, 0.4);
  transform: translateY(-1px);
}

.link-button {
  width: 100%;
  color: #86868b;
  font-size: 14px;
}

.link-text {
  color: #ff6b9d;
  font-weight: 500;
  transition: color 0.2s;
}

.link-button:hover .link-text {
  color: #ff8fab;
}

:deep(.el-form-item__label) {
  color: #1d1d1f;
  font-weight: 500;
  font-size: 15px;
}

:deep(.el-input__wrapper) {
  border-radius: 12px;
  padding: 12px 16px;
}

:deep(.el-input__prefix) {
  display: flex;
  align-items: center;
  margin-right: 8px;
}

:deep(.el-form-item) {
  margin-bottom: 24px;
}

:deep(.el-form-item__error) {
  color: #ff6b9d;
  font-size: 13px;
}
</style>
