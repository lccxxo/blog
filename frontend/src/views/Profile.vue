<template>
  <div class="profile-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>个人信息</span>
        </div>
      </template>
      <div class="profile-content" v-if="userStore.user">
        <el-descriptions :column="1" border>
          <el-descriptions-item label="用户名">
            {{ userStore.user.username }}
          </el-descriptions-item>
          <el-descriptions-item label="邮箱">
            {{ userStore.user.email }}
          </el-descriptions-item>
          <el-descriptions-item label="昵称">
            {{ userStore.user.nickname || '-' }}
          </el-descriptions-item>
          <el-descriptions-item label="注册时间">
            {{ formatDate(userStore.user.created_at) }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  return date.toLocaleString('zh-CN')
}
</script>

<style scoped>
.profile-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 40px 22px;
}

.card-header {
  text-align: center;
  font-size: 28px;
  font-weight: 600;
  color: #1d1d1f;
  letter-spacing: -0.011em;
  padding: 20px 0;
}

.profile-content {
  padding: 30px 0;
}

:deep(.el-descriptions__label) {
  color: #1d1d1f;
  font-weight: 500;
  font-size: 15px;
}

:deep(.el-descriptions__content) {
  color: #1d1d1f;
  font-size: 15px;
}

:deep(.el-descriptions__border) {
  border-color: rgba(255, 182, 193, 0.2);
}

:deep(.el-descriptions__label.is-bordered-label) {
  background: rgba(255, 240, 248, 0.3);
  border-color: rgba(255, 182, 193, 0.2);
}

:deep(.el-descriptions__content.is-bordered-content) {
  border-color: rgba(255, 182, 193, 0.2);
}
</style>





