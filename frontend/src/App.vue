<template>
  <div id="app">
    <div class="layout-container">
      <!-- 苹果风格导航栏 -->
      <nav class="apple-nav" :class="{ 'nav-scrolled': isScrolled }">
        <div class="nav-container">
          <div class="nav-brand" @click="router.push('/')">
            <div class="brand-icon">
              <svg viewBox="0 0 24 24" fill="currentColor" stroke="none">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
              </svg>
            </div>
            <span class="brand-text">lccxxo</span>
          </div>
          
          <div class="nav-menu">
            <a 
              v-for="item in menuItems" 
              :key="item.path"
              :class="['nav-link', { active: activeMenu === item.path }]"
              @click="router.push(item.path)"
            >
              {{ item.label }}
            </a>
            <template v-if="userStore.isLoggedIn">
              <a 
                v-for="item in authMenuItems" 
                :key="item.path"
                :class="['nav-link', { active: activeMenu === item.path }]"
                @click="router.push(item.path)"
              >
                {{ item.label }}
              </a>
            </template>
          </div>

          <div class="nav-actions">
            <template v-if="userStore.isLoggedIn">
              <button class="btn-primary" @click="router.push('/create-article')">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="12" y1="5" x2="12" y2="19"></line>
                  <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
                写文章
              </button>
              <div class="user-dropdown">
                <div class="user-avatar" @click="showUserMenu = !showUserMenu">
                  {{ userStore.user?.username?.charAt(0).toUpperCase() }}
                </div>
                <div v-if="showUserMenu" class="dropdown-menu" @click.stop>
                  <div class="dropdown-item" @click="router.push('/profile'); showUserMenu = false">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                      <circle cx="12" cy="7" r="4"></circle>
                    </svg>
                    个人信息
                  </div>
                  <div class="dropdown-divider"></div>
                  <div class="dropdown-item" @click="handleLogout">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
                      <polyline points="16 17 21 12 16 7"></polyline>
                      <line x1="21" y1="12" x2="9" y2="12"></line>
                    </svg>
                    退出登录
                  </div>
                </div>
              </div>
            </template>
            <template v-else>
              <a class="nav-link auth-link" @click="router.push('/login')">登录</a>
              <button class="btn-primary" @click="router.push('/register')">注册</button>
            </template>
          </div>
        </div>
      </nav>

      <!-- 主内容区域 -->
      <main class="main-content">
        <router-view />
      </main>

      <!-- 页脚 -->
      <footer class="apple-footer">
        <div class="footer-content">
          <p>© 2025 mjh&zcc's home. All rights reserved.</p>
        </div>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const isScrolled = ref(false)
const showUserMenu = ref(false)

const activeMenu = computed(() => route.path)

const menuItems = [
  { path: '/', label: '首页' },
  { path: '/categories', label: '分类' },
  { path: '/tags', label: '标签' }
]

const authMenuItems = [
  { path: '/my-articles', label: '我的文章' },
  { path: '/my-comments', label: '我的评论' }
]

const handleScroll = () => {
  isScrolled.value = window.scrollY > 10
}

const handleLogout = () => {
  userStore.logout()
  ElMessage.success('已退出登录')
  router.push('/')
  showUserMenu.value = false
}

const handleClickOutside = (e) => {
  if (!e.target.closest('.user-dropdown')) {
    showUserMenu.value = false
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.layout-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: #fff;
}

/* 苹果风格导航栏 - 更大气的设计 */
.apple-nav {
  position: sticky;
  top: 0;
  z-index: 1000;
  background: rgba(255, 250, 249, 0.85);
  backdrop-filter: saturate(180%) blur(30px);
  -webkit-backdrop-filter: saturate(180%) blur(30px);
  border-bottom: 0.5px solid rgba(255, 182, 193, 0.2);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  height: 72px;
  box-shadow: 0 1px 20px rgba(255, 182, 193, 0.08);
}

.apple-nav.nav-scrolled {
  background: rgba(255, 250, 249, 0.95);
  box-shadow: 0 2px 24px rgba(255, 182, 193, 0.12);
}

.nav-container {
  max-width: 1440px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 44px;
}

.nav-brand {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  user-select: none;
  transition: opacity 0.2s;
  padding: 4px 0;
}

.nav-brand:hover {
  opacity: 0.7;
}

.brand-icon {
  width: 32px;
  height: 32px;
  color: #ff6b9d;
}

.brand-text {
  font-size: 24px;
  font-weight: 600;
  letter-spacing: -0.022em;
  color: #1d1d1f;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
  justify-content: center;
}

.nav-link {
  padding: 0 20px;
  font-size: 15px;
  font-weight: 400;
  color: #1d1d1f;
  text-decoration: none;
  cursor: pointer;
  transition: opacity 0.2s;
  line-height: 72px;
  position: relative;
  white-space: nowrap;
}

.nav-link:hover {
  opacity: 0.65;
}

.nav-link.active {
  color: #1d1d1f;
  font-weight: 500;
}

.nav-link.active::after {
  content: '';
  position: absolute;
  bottom: 16px;
  left: 20px;
  right: 20px;
  height: 3px;
  background: #1d1d1f;
  border-radius: 2px;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.btn-primary {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 24px;
  font-size: 15px;
  font-weight: 400;
  color: #fff;
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border: none;
  border-radius: 980px;
  cursor: pointer;
  transition: all 0.2s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  white-space: nowrap;
  min-height: 44px;
  box-shadow: 0 2px 12px rgba(255, 107, 157, 0.3);
}

.btn-primary:hover {
  background: linear-gradient(135deg, #ff8fab 0%, #ff6b9d 100%);
  transform: scale(1.02);
  box-shadow: 0 4px 16px rgba(255, 107, 157, 0.4);
}

.btn-primary:active {
  transform: scale(0.98);
}

.auth-link {
  padding: 0 12px;
  font-size: 15px;
  color: #ff6b9d;
  font-weight: 400;
  cursor: pointer;
  transition: opacity 0.2s;
  line-height: 72px;
}

.auth-link:hover {
  opacity: 0.7;
}

.user-dropdown {
  position: relative;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ffb6c1 0%, #ffa07a 50%, #ffb347 100%);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s;
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.4);
}

.user-avatar:hover {
  transform: scale(1.05);
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 12px);
  right: 0;
  min-width: 200px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: saturate(180%) blur(20px);
  -webkit-backdrop-filter: saturate(180%) blur(20px);
  border-radius: 12px;
  border: 1px solid rgba(255, 182, 193, 0.3);
  box-shadow: 0 4px 20px rgba(255, 182, 193, 0.2);
  padding: 10px 0;
  z-index: 1001;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 20px;
  font-size: 15px;
  color: #1d1d1f;
  cursor: pointer;
  transition: background 0.15s;
}

.dropdown-item:hover {
  background: rgba(255, 182, 193, 0.15);
}

.dropdown-item svg {
  opacity: 0.6;
}

.dropdown-divider {
  height: 0.5px;
  background: rgba(0, 0, 0, 0.1);
  margin: 4px 0;
}

.main-content {
  flex: 1;
  background: linear-gradient(180deg, #fffaf9 0%, #fff5f5 50%, #ffeef8 100%);
  min-height: calc(100vh - 72px - 80px);
  position: relative;
}

.main-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    radial-gradient(circle at 20% 50%, rgba(255, 182, 193, 0.08) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(255, 218, 185, 0.08) 0%, transparent 50%),
    radial-gradient(circle at 40% 20%, rgba(255, 240, 248, 0.1) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

.apple-footer {
  background: linear-gradient(180deg, rgba(255, 250, 249, 0.95) 0%, rgba(255, 240, 248, 0.9) 100%);
  backdrop-filter: blur(20px);
  border-top: 0.5px solid rgba(255, 182, 193, 0.2);
  padding: 50px 0;
  box-shadow: 0 -2px 20px rgba(255, 182, 193, 0.08);
  margin-top: 80px;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 22px;
  text-align: center;
}

.footer-content p {
  font-size: 12px;
  color: #86868b;
  letter-spacing: -0.01em;
}

@media (max-width: 1024px) {
  .nav-container {
    padding: 0 24px;
  }
  
  .brand-text {
    font-size: 20px;
  }
  
  .brand-icon {
    width: 28px;
    height: 28px;
  }
  
  .nav-link {
    padding: 0 16px;
    font-size: 14px;
  }
}

@media (max-width: 768px) {
  .apple-nav {
    height: 64px;
  }
  
  .nav-container {
    padding: 0 20px;
  }
  
  .nav-menu {
    display: none;
  }
  
  .nav-link.auth-link {
    display: none;
  }
  
  .brand-text {
    font-size: 18px;
  }
  
  .brand-icon {
    width: 24px;
    height: 24px;
  }
  
  .btn-primary {
    padding: 8px 20px;
    font-size: 14px;
    min-height: 40px;
  }
  
  .user-avatar {
    width: 36px;
    height: 36px;
    font-size: 14px;
  }
}
</style>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'SF Pro Text', 'Helvetica Neue', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #1d1d1f;
  background: #fff;
}

#app {
  min-height: 100vh;
}

/* 全局滚动条样式 */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: rgba(255, 240, 248, 0.5);
}

::-webkit-scrollbar-thumb {
  background: rgba(255, 182, 193, 0.5);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 182, 193, 0.7);
}

/* Element Plus 全局粉色主题 */
:deep(.el-button--primary) {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border-color: #ff6b9d;
  box-shadow: 0 2px 12px rgba(255, 107, 157, 0.3);
}

:deep(.el-button--primary:hover) {
  background: linear-gradient(135deg, #ff8fab 0%, #ff6b9d 100%);
  border-color: #ff8fab;
  box-shadow: 0 4px 16px rgba(255, 107, 157, 0.4);
}

:deep(.el-input__wrapper) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.3) inset;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.5) inset;
}

:deep(.el-input.is-focus .el-input__wrapper) {
  box-shadow: 0 0 0 1px #ff6b9d inset, 0 0 0 3px rgba(255, 107, 157, 0.1);
}

:deep(.el-textarea__inner) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.3) inset;
}

:deep(.el-textarea__inner:hover) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.5) inset;
}

:deep(.el-textarea.is-focus .el-textarea__inner) {
  box-shadow: 0 0 0 1px #ff6b9d inset, 0 0 0 3px rgba(255, 107, 157, 0.1);
}

:deep(.el-select .el-input__wrapper) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.3) inset;
}

:deep(.el-select:hover .el-input__wrapper) {
  box-shadow: 0 0 0 1px rgba(255, 182, 193, 0.5) inset;
}

:deep(.el-select.is-focus .el-input__wrapper) {
  box-shadow: 0 0 0 1px #ff6b9d inset, 0 0 0 3px rgba(255, 107, 157, 0.1);
}

:deep(.el-card) {
  border: 1px solid rgba(255, 182, 193, 0.15);
  box-shadow: 0 2px 12px rgba(255, 182, 193, 0.1);
}

:deep(.el-card__header) {
  border-bottom: 1px solid rgba(255, 182, 193, 0.15);
  background: rgba(255, 240, 248, 0.3);
}

:deep(.el-link--primary) {
  color: #ff6b9d;
}

:deep(.el-link--primary:hover) {
  color: #ff8fab;
}

:deep(.el-tag--primary) {
  background: rgba(255, 107, 157, 0.12);
  border-color: rgba(255, 107, 157, 0.2);
  color: #ff6b9d;
}

:deep(.el-tag--info) {
  background: rgba(255, 218, 185, 0.3);
  border-color: rgba(255, 218, 185, 0.4);
  color: #d4a574;
}

:deep(.el-tabs__item.is-active) {
  color: #ff6b9d;
}

:deep(.el-tabs__active-bar) {
  background-color: #ff6b9d;
}

:deep(.el-tabs__item:hover) {
  color: #ff8fab;
}

:deep(.el-upload) {
  border: 1.5px dashed rgba(255, 182, 193, 0.4);
  border-radius: 8px;
  background: rgba(255, 240, 248, 0.3);
}

:deep(.el-upload:hover) {
  border-color: rgba(255, 182, 193, 0.6);
  background: rgba(255, 240, 248, 0.5);
}

:deep(.el-image__inner) {
  border-radius: 12px;
}

:deep(.el-avatar) {
  background: linear-gradient(135deg, #ffb6c1 0%, #ffa07a 50%, #ffb347 100%);
}

:deep(.el-switch.is-checked .el-switch__core) {
  background-color: #ff6b9d;
  border-color: #ff6b9d;
}

:deep(.el-switch__core) {
  border-color: rgba(255, 182, 193, 0.5);
}

:deep(.el-switch.is-checked .el-switch__core::after) {
  background-color: #fff;
}

:deep(.el-message-box) {
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
}

:deep(.el-message-box__header) {
  background: rgba(255, 240, 248, 0.3);
  border-bottom: 1px solid rgba(255, 182, 193, 0.2);
}

:deep(.el-message-box__title) {
  color: #1d1d1f;
  font-weight: 600;
}

:deep(.el-message-box__btns .el-button--primary) {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border: none;
}

:deep(.el-pagination .el-pager li.is-active) {
  background-color: #ff6b9d;
  color: #fff;
}

:deep(.el-pagination .el-pager li:hover) {
  color: #ff6b9d;
}

:deep(.el-pagination button:hover) {
  color: #ff6b9d;
}

:deep(.el-pagination .btn-next:hover),
:deep(.el-pagination .btn-prev:hover) {
  color: #ff6b9d;
}
</style>

