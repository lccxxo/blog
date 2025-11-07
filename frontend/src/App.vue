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
              <span class="nav-icon" v-if="item.icon">
                <component :is="item.icon" />
              </span>
              <span>{{ item.label }}</span>
            </a>
            <template v-if="userStore.isLoggedIn">
              <a 
                v-for="item in authMenuItems" 
                :key="item.path"
                :class="['nav-link', { active: activeMenu === item.path }]"
                @click="router.push(item.path)"
              >
                <span class="nav-icon" v-if="item.icon">
                  <component :is="item.icon" />
                </span>
                <span>{{ item.label }}</span>
              </a>
            </template>
          </div>

          <div class="nav-actions">
            <template v-if="userStore.isLoggedIn">
              <!-- 通知图标和面板 -->
              <div class="notification-wrapper">
                <div class="notification-icon" @click="toggleNotifications">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
                    <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
                  </svg>
                  <span v-if="unreadCount > 0" class="badge">{{ unreadCount > 99 ? '99+' : unreadCount }}</span>
                </div>
                
                <!-- 通知面板 -->
                <div v-if="showNotifications" class="notifications-panel" @click.stop>
                  <div class="panel-header">
                    <h3>通知</h3>
                    <div class="header-actions">
                      <button @click="markAllAsRead" class="mark-all-btn">全部已读</button>
                      <button @click="showNotifications = false" class="close-btn">×</button>
                    </div>
                  </div>
                  <div class="notifications-list" v-loading="loadingNotifications">
                    <div v-if="notifications.length === 0" class="empty-state">
                      <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                        <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
                        <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
                      </svg>
                      <p>暂无通知</p>
                    </div>
                    <div 
                      v-for="notification in notifications" 
                      :key="notification.id"
                      :class="['notification-item', { unread: !notification.is_read }]"
                      @click="handleNotificationClick(notification)"
                    >
                      <div class="notification-avatar">
                        <UserAvatar
                          :avatar="notification.from_user?.avatar"
                          :username="notification.from_user?.username"
                          :nickname="notification.from_user?.nickname"
                          size="small"
                        />
                      </div>
                      <div class="notification-content">
                        <p class="notification-text">{{ notification.content }}</p>
                        <span class="notification-time">{{ formatTime(notification.created_at) }}</span>
                      </div>
                      <button class="delete-btn" @click.stop="deleteNotification(notification.id)">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <line x1="18" y1="6" x2="6" y2="18"></line>
                          <line x1="6" y1="6" x2="18" y2="18"></line>
                        </svg>
                      </button>
                    </div>
                  </div>
                </div>
              </div>
              
              <button class="btn-primary" @click="router.push('/create-article')">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="12" y1="5" x2="12" y2="19"></line>
                  <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
                写文章
              </button>
              <div class="user-dropdown">
                <div class="user-avatar-clickable" @click.stop="showUserMenu = !showUserMenu">
                  <UserAvatar
                    :avatar="userStore.user?.avatar"
                    :username="userStore.user?.username"
                    :nickname="userStore.user?.nickname"
                    size="medium"
                  />
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
import { computed, ref, onMounted, onUnmounted, h, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import { notificationAPI } from '@/api/notification'
import UserAvatar from '@/components/UserAvatar.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const isScrolled = ref(false)
const showUserMenu = ref(false)
const showNotifications = ref(false)
const notifications = ref([])
const unreadCount = ref(0)
const loadingNotifications = ref(false)

const activeMenu = computed(() => route.path)

// 三丽鸥风格 Hello Kitty 图标
// 首页 - 经典蝴蝶结 Kitty（左耳蝴蝶结）
const HelloKittyHomeIcon = () => h('svg', {
  width: 20,
  height: 20,
  viewBox: '0 0 24 24',
  fill: 'none',
  xmlns: 'http://www.w3.org/2000/svg',
  class: 'hello-kitty-icon'
}, [
  // 右耳
  h('circle', { cx: 16.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 16.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 左耳 + 蝴蝶结
  h('circle', { cx: 7.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 7.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 蝴蝶结（在左耳上）
  h('path', { d: 'M 6 4 Q 6.5 3 7.5 3.5 Q 8.5 4 8.5 5 Q 8.5 6 7.5 6.5 Q 6.5 7 6 6.5 Q 5.5 6 5.5 5 Q 5.5 4 6 4 Z', fill: '#FF1493' }),
  h('circle', { cx: 7.5, cy: 5, r: 0.6, fill: '#FFFFFF' }),
  h('path', { d: 'M 6.5 4.5 L 8.5 4.5', stroke: '#FFFFFF', 'stroke-width': 0.5 }),
  // 头部
  h('circle', { cx: 12, cy: 12, r: 7, fill: '#FFFFFF', stroke: 'none' }),
  // 左眼（大而圆）
  h('ellipse', { cx: 9.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 10, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 右眼
  h('ellipse', { cx: 14.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 15, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 黄色小鼻子
  h('ellipse', { cx: 12, cy: 13.5, rx: 0.8, ry: 0.6, fill: '#FFD700' }),
  // 没有嘴巴（Hello Kitty 的特征）
])

// 分类 - 戴帽子 Kitty
const HelloKittyCategoryIcon = () => h('svg', {
  width: 20,
  height: 20,
  viewBox: '0 0 24 24',
  fill: 'none',
  xmlns: 'http://www.w3.org/2000/svg',
  class: 'hello-kitty-icon'
}, [
  // 右耳
  h('circle', { cx: 16.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 16.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 左耳
  h('circle', { cx: 7.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 7.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 头部
  h('circle', { cx: 12, cy: 12, r: 7, fill: '#FFFFFF', stroke: 'none' }),
  // 帽子
  h('ellipse', { cx: 12, cy: 5.5, rx: 4.5, ry: 2.5, fill: '#FFD700' }),
  h('ellipse', { cx: 12, cy: 4.8, rx: 3.5, ry: 1.8, fill: '#FFE135' }),
  h('circle', { cx: 12, cy: 4.5, r: 0.5, fill: '#FFFFFF' }),
  // 左眼
  h('ellipse', { cx: 9.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 10, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 右眼
  h('ellipse', { cx: 14.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 15, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 黄色小鼻子
  h('ellipse', { cx: 12, cy: 13.5, rx: 0.8, ry: 0.6, fill: '#FFD700' }),
])

// 标签 - 带星星 Kitty（右耳蝴蝶结）
const HelloKittyTagIcon = () => h('svg', {
  width: 20,
  height: 20,
  viewBox: '0 0 24 24',
  fill: 'none',
  xmlns: 'http://www.w3.org/2000/svg',
  class: 'hello-kitty-icon'
}, [
  // 右耳 + 蝴蝶结
  h('circle', { cx: 16.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 16.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  h('path', { d: 'M 15.5 4 Q 16 3 17 3.5 Q 18 4 18 5 Q 18 6 17 6.5 Q 16 7 15.5 6.5 Q 15 6 15 5 Q 15 4 15.5 4 Z', fill: '#FF69B4' }),
  h('circle', { cx: 17, cy: 5, r: 0.6, fill: '#FFFFFF' }),
  h('path', { d: 'M 16 4.5 L 18 4.5', stroke: '#FFFFFF', 'stroke-width': 0.5 }),
  // 左耳
  h('circle', { cx: 7.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 7.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 星星装饰
  h('path', { d: 'M 12 2 L 12.8 4.2 L 15 4 L 13.2 5.8 L 15.5 7.5 L 12.5 7 L 12 9 L 11.5 7 L 8.5 7.5 L 10.8 5.8 L 9 4 L 11.2 4.2 Z', fill: '#FFD700' }),
  // 头部
  h('circle', { cx: 12, cy: 12, r: 7, fill: '#FFFFFF', stroke: 'none' }),
  // 左眼
  h('ellipse', { cx: 9.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 10, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 右眼
  h('ellipse', { cx: 14.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 15, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 黄色小鼻子
  h('ellipse', { cx: 12, cy: 13.5, rx: 0.8, ry: 0.6, fill: '#FFD700' }),
])

// 我的文章 - 拿笔 Kitty
const HelloKittyArticleIcon = () => h('svg', {
  width: 20,
  height: 20,
  viewBox: '0 0 24 24',
  fill: 'none',
  xmlns: 'http://www.w3.org/2000/svg',
  class: 'hello-kitty-icon'
}, [
  // 右耳
  h('circle', { cx: 16.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 16.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 左耳
  h('circle', { cx: 7.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 7.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 头部
  h('circle', { cx: 12, cy: 12, r: 7, fill: '#FFFFFF', stroke: 'none' }),
  // 笔（在右侧）
  h('line', { x1: 17.5, y1: 10, x2: 19, y2: 8.5, stroke: '#8B4513', 'stroke-width': 2.5, 'stroke-linecap': 'round' }),
  h('line', { x1: 17.5, y1: 10.5, x2: 19.5, y2: 8.5, stroke: '#FF6B9D', 'stroke-width': 1.8, 'stroke-linecap': 'round' }),
  h('circle', { cx: 19.2, cy: 8.2, r: 0.5, fill: '#FFB6C1' }),
  // 左眼
  h('ellipse', { cx: 9.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 10, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 右眼
  h('ellipse', { cx: 14.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 15, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 黄色小鼻子
  h('ellipse', { cx: 12, cy: 13.5, rx: 0.8, ry: 0.6, fill: '#FFD700' }),
])

// 我的评论 - 带对话气泡 Kitty
const HelloKittyCommentIcon = () => h('svg', {
  width: 20,
  height: 20,
  viewBox: '0 0 24 24',
  fill: 'none',
  xmlns: 'http://www.w3.org/2000/svg',
  class: 'hello-kitty-icon'
}, [
  // 右耳
  h('circle', { cx: 16.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 16.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 左耳
  h('circle', { cx: 7.5, cy: 5, r: 3.5, fill: '#FFFFFF', stroke: 'none' }),
  h('circle', { cx: 7.5, cy: 5, r: 2.5, fill: '#FFB6C1', opacity: 0.3 }),
  // 头部
  h('circle', { cx: 12, cy: 12, r: 7, fill: '#FFFFFF', stroke: 'none' }),
  // 对话气泡（右侧）
  h('ellipse', { cx: 17.5, cy: 16, rx: 4, ry: 3.5, fill: '#FFE4E1', stroke: '#FFB6C1', 'stroke-width': 1.2 }),
  h('path', { d: 'M 15 16.5 L 13 18.5 L 14 17 Z', fill: '#FFE4E1' }),
  h('circle', { cx: 16.5, cy: 15.2, r: 0.5, fill: '#FFB6C1' }),
  h('circle', { cx: 18, cy: 16, r: 0.5, fill: '#FFB6C1' }),
  h('circle', { cx: 16.5, cy: 16.8, r: 0.5, fill: '#FFB6C1' }),
  // 左眼
  h('ellipse', { cx: 9.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 10, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 右眼
  h('ellipse', { cx: 14.5, cy: 11, rx: 2.2, ry: 2.8, fill: '#000000' }),
  h('circle', { cx: 15, cy: 10.5, r: 0.8, fill: '#FFFFFF' }),
  // 黄色小鼻子
  h('ellipse', { cx: 12, cy: 13.5, rx: 0.8, ry: 0.6, fill: '#FFD700' }),
])

const menuItems = [
  { path: '/', label: '首页', icon: HelloKittyHomeIcon },
  { path: '/categories', label: '分类', icon: HelloKittyCategoryIcon },
  { path: '/tags', label: '标签', icon: HelloKittyTagIcon }
]

const authMenuItems = [
  { path: '/my-articles', label: '我的文章', icon: HelloKittyArticleIcon },
  { path: '/my-comments', label: '我的评论', icon: HelloKittyCommentIcon }
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

// 切换通知面板
const toggleNotifications = () => {
  console.log('切换通知面板, 当前状态:', showNotifications.value)
  showNotifications.value = !showNotifications.value
  console.log('新状态:', showNotifications.value)
  if (showNotifications.value) {
    loadNotifications()
  }
}

const handleClickOutside = (e) => {
  if (!e.target.closest('.user-dropdown')) {
    showUserMenu.value = false
  }
  if (!e.target.closest('.notification-wrapper')) {
    showNotifications.value = false
  }
}

// 加载通知列表
const loadNotifications = async () => {
  if (!userStore.isLoggedIn) {
    console.log('用户未登录，跳过加载通知')
    return
  }
  
  try {
    loadingNotifications.value = true
    console.log('开始加载通知...')
    const res = await notificationAPI.getNotifications({ page: 1, page_size: 20 })
    notifications.value = res.notifications || []
    console.log('通知加载成功:', notifications.value)
  } catch (error) {
    console.error('加载通知失败:', error)
    ElMessage.error('加载通知失败: ' + (error.response?.data?.error || error.message))
  } finally {
    loadingNotifications.value = false
  }
}

// 加载未读数量
const loadUnreadCount = async () => {
  if (!userStore.isLoggedIn) {
    console.log('用户未登录，跳过加载未读数')
    return
  }
  
  try {
    console.log('开始加载未读数量...')
    const res = await notificationAPI.getUnreadCount()
    unreadCount.value = res.count || 0
    console.log('未读数量:', unreadCount.value)
  } catch (error) {
    console.error('加载未读数量失败:', error)
  }
}

// 标记所有为已读
const markAllAsRead = async () => {
  try {
    await notificationAPI.markAllAsRead()
    notifications.value.forEach(n => n.is_read = true)
    unreadCount.value = 0
    ElMessage.success('已全部标记为已读')
  } catch (error) {
    console.error('标记失败:', error)
  }
}

// 删除通知
const deleteNotification = async (id) => {
  try {
    await notificationAPI.deleteNotification(id)
    notifications.value = notifications.value.filter(n => n.id !== id)
    loadUnreadCount()
    ElMessage.success('通知已删除')
  } catch (error) {
    console.error('删除失败:', error)
  }
}

// 点击通知
const handleNotificationClick = async (notification) => {
  console.log('点击通知，完整数据:', JSON.stringify(notification, null, 2))
  
  // 标记为已读
  if (!notification.is_read) {
    try {
      await notificationAPI.markAsRead(notification.id)
      notification.is_read = true
      unreadCount.value = Math.max(0, unreadCount.value - 1)
    } catch (error) {
      console.error('标记通知为已读失败:', error)
    }
  }
  
  // 跳转到对应文章和评论
  const articleId = notification.article_id || notification.article?.id
  console.log('文章ID:', articleId, '通知对象:', notification)
  
  if (articleId) {
    showNotifications.value = false
    // 如果有评论ID，添加到URL参数中
    const commentId = notification.comment_id || notification.comment?.id
    console.log('准备跳转 - 文章ID:', articleId, '评论ID:', commentId)
    
    if (commentId) {
      router.push({
        path: `/article/${articleId}`,
        query: { comment: String(commentId) }
      }).catch(err => {
        console.error('路由跳转失败:', err)
        ElMessage.error('跳转失败，请重试')
      })
    } else {
      router.push(`/article/${articleId}`).catch(err => {
        console.error('路由跳转失败:', err)
        ElMessage.error('跳转失败，请重试')
      })
    }
  } else {
    console.error('通知中没有文章ID:', notification)
    ElMessage.error('通知数据异常，无法跳转')
  }
}

// 格式化时间
const formatTime = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date
  
  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour
  
  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return `${Math.floor(diff / minute)}分钟前`
  } else if (diff < day) {
    return `${Math.floor(diff / hour)}小时前`
  } else if (diff < 7 * day) {
    return `${Math.floor(diff / day)}天前`
  } else {
    return date.toLocaleDateString('zh-CN')
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  document.addEventListener('click', handleClickOutside)
  
  // 如果已登录，加载通知
  if (userStore.isLoggedIn) {
    loadNotifications()
    loadUnreadCount()
    
    // 每30秒刷新未读数
    setInterval(loadUnreadCount, 30000)
  }
})

// 监听登录状态变化
watch(() => userStore.isLoggedIn, (isLoggedIn) => {
  if (isLoggedIn) {
    loadNotifications()
    loadUnreadCount()
  } else {
    notifications.value = []
    unreadCount.value = 0
  }
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
  color: #ffb6c1;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s;
  line-height: 72px;
  position: relative;
  white-space: nowrap;
  display: flex;
  align-items: center;
  gap: 8px;
}

.nav-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.3s;
}

.nav-link:hover .nav-icon {
  transform: scale(1.15) rotate(5deg);
}

.nav-link.active .nav-icon {
  transform: scale(1.1);
}

.nav-link:hover {
  color: #ff8fab;
  opacity: 1;
}

.nav-link.active {
  color: #ff6b9d;
  font-weight: 500;
}

.nav-link.active::after {
  content: '';
  position: absolute;
  bottom: 16px;
  left: 20px;
  right: 20px;
  height: 2.5px;
  background: linear-gradient(90deg, #ff6b9d 0%, #ff8fab 100%);
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(255, 107, 157, 0.3);
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

/* 通知包装器 */
.notification-wrapper {
  position: relative;
}

/* 通知图标样式 */
.notification-icon {
  position: relative;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 12px;
  transition: all 0.3s ease;
  color: #1d1d1f;
}

.notification-icon:hover {
  background: rgba(255, 182, 193, 0.15);
  color: #ff6b9d;
}

.notification-icon .badge {
  position: absolute;
  top: 6px;
  right: 6px;
  background: linear-gradient(135deg, #ff6b9d, #ff8fab);
  color: white;
  font-size: 10px;
  font-weight: 600;
  padding: 2px 5px;
  border-radius: 10px;
  min-width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(255, 107, 157, 0.4);
}

/* 通知面板样式 */
.notifications-panel {
  position: absolute;
  top: calc(100% + 10px);
  right: 0;
  width: 380px;
  max-height: 500px;
  background: rgba(255, 255, 255, 0.98);
  backdrop-filter: blur(20px);
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.12);
  border: 1px solid rgba(255, 182, 193, 0.2);
  z-index: 1001;
  overflow: hidden;
}

.panel-header {
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid rgba(255, 182, 193, 0.15);
  background: rgba(255, 240, 248, 0.3);
}

.panel-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #1d1d1f;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.mark-all-btn {
  padding: 4px 12px;
  font-size: 12px;
  color: #ff6b9d;
  background: transparent;
  border: 1px solid rgba(255, 107, 157, 0.3);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.mark-all-btn:hover {
  background: rgba(255, 182, 193, 0.15);
  border-color: #ff6b9d;
}

.close-btn {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: #86868b;
  background: transparent;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.close-btn:hover {
  background: rgba(255, 182, 193, 0.15);
  color: #ff6b9d;
}

.notifications-list {
  max-height: 420px;
  overflow-y: auto;
}

.empty-state {
  padding: 60px 20px;
  text-align: center;
  color: #86868b;
}

.empty-state svg {
  color: rgba(255, 182, 193, 0.5);
  margin-bottom: 16px;
}

.empty-state p {
  margin: 0;
  font-size: 14px;
}

.notification-item {
  padding: 16px 20px;
  display: flex;
  gap: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  border-bottom: 1px solid rgba(255, 182, 193, 0.1);
  position: relative;
}

.notification-item:hover {
  background: rgba(255, 240, 248, 0.5);
}

.notification-item.unread {
  background: rgba(255, 240, 248, 0.3);
}

.notification-item.unread::before {
  content: '';
  position: absolute;
  left: 8px;
  top: 50%;
  transform: translateY(-50%);
  width: 6px;
  height: 6px;
  background: #ff6b9d;
  border-radius: 50%;
}

.notification-avatar {
  flex-shrink: 0;
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-text {
  margin: 0 0 4px 0;
  font-size: 14px;
  color: #1d1d1f;
  line-height: 1.5;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
}

.notification-time {
  font-size: 12px;
  color: #86868b;
}

.delete-btn {
  flex-shrink: 0;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  color: #86868b;
  opacity: 0;
  transition: all 0.3s ease;
}

.notification-item:hover .delete-btn {
  opacity: 1;
}

.delete-btn:hover {
  background: rgba(255, 107, 157, 0.15);
  color: #ff6b9d;
}

@media (max-width: 768px) {
  .notifications-panel {
    width: calc(100vw - 32px);
    right: 16px;
  }
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
    gap: 6px;
  }
  
  .nav-icon {
    width: 16px;
    height: 16px;
  }
}

.hello-kitty-icon {
  transition: all 0.3s;
  filter: drop-shadow(0 2px 4px rgba(255, 182, 193, 0.3));
}

.nav-link:hover .hello-kitty-icon {
  filter: drop-shadow(0 3px 6px rgba(255, 182, 193, 0.5));
}

.nav-link.active .hello-kitty-icon {
  filter: drop-shadow(0 3px 8px rgba(255, 107, 157, 0.6));
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

