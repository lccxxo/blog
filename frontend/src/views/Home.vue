<template>
  <div class="home-container">
    <!-- Hero Section -->
    <section class="hero-section">
      <!-- 恋爱日记标题图片 -->
      <div class="diary-title">
        <img 
          src="/diary-title.png" 
          alt="恋爱日记"
          class="diary-title-img"
          @error="handleDiaryTitleError"
        />
      </div>
      <!-- 温馨装饰小图片 -->
      <div class="decorative-images">
        <!-- 标题左侧装饰 -->
        <img 
          src="/decorations/baker-bear.png" 
          alt=""
          class="decor-img decor-left"
          @error="handleDecorError"
        />
        <!-- 标题右侧装饰 -->
        <img 
          src="/decorations/v-sign-bear.png" 
          alt=""
          class="decor-img decor-right"
          @error="handleDecorError"
        />
        <!-- 左上角装饰 -->
        <img 
          src="/decorations/look-bear.png" 
          alt=""
          class="decor-img decor-top-left"
          @error="handleDecorError"
        />
        <!-- 右上角装饰 -->
        <img 
          src="/decorations/sleep-bear.png" 
          alt=""
          class="decor-img decor-top-right"
          @error="handleDecorError"
        />
        <!-- 左下角装饰 -->
        <img 
          src="/decorations/motorcycle-bear.png" 
          alt=""
          class="decor-img decor-bottom-left"
          @error="handleDecorError"
        />
        <!-- 中部左侧装饰 -->
        <img 
          src="/decorations/flower-bear.png" 
          alt=""
          class="decor-img decor-mid-left"
          @error="handleDecorError"
        />
        <!-- 中部右侧装饰 -->
        <img 
          src="/decorations/pink-bear.png" 
          alt=""
          class="decor-img decor-mid-right"
          @error="handleDecorError"
        />
      </div>
      <div class="hero-image-container" :class="{ 'hero-fallback': !heroImageLoaded }">
        <!-- 气泡装饰 -->
        <div class="bubbles">
          <div class="bubble bubble-1"></div>
          <div class="bubble bubble-2"></div>
          <div class="bubble bubble-3"></div>
          <div class="bubble bubble-4"></div>
          <div class="bubble bubble-5"></div>
          <div class="bubble bubble-6"></div>
          <div class="bubble bubble-7"></div>
          <div class="bubble bubble-8"></div>
        </div>
        <img 
          v-if="heroImageLoaded"
          :src="heroImageSrc" 
          alt="欢迎来到我们的博客"
          class="hero-image"
          @error="handleHeroImageError"
          @load="handleHeroImageLoad"
        />
        <div class="hero-overlay"></div>
        <div class="hero-content">
          <h1 class="hero-title">记录我们</h1>
          <p class="hero-subtitle">把每一天写成诗 把每一刻都记住</p>
        </div>
        <!-- 恋爱纪念日 -->
        <div class="anniversary-card">
          <div class="anniversary-icon">
            <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
            </svg>
          </div>
          <div class="anniversary-info">
            <div class="anniversary-years">第 {{ anniversaryYears }} 年</div>
            <div class="anniversary-countdown">
              <div v-if="countdown.days > 0 || countdown.hours > 0 || countdown.minutes > 0 || countdown.seconds > 0" class="countdown-text">
                距离下一个纪念日还有
              </div>
              <div v-else class="countdown-text">今天就是我们的纪念日！</div>
              <div v-if="countdown.days > 0 || countdown.hours > 0 || countdown.minutes > 0 || countdown.seconds > 0" class="countdown-time">
                <span v-if="countdown.days > 0" class="time-unit">
                  <span class="time-value">{{ countdown.days }}</span>
                  <span class="time-label">天</span>
                </span>
                <span v-if="countdown.hours > 0" class="time-unit">
                  <span class="time-value">{{ countdown.hours }}</span>
                  <span class="time-label">时</span>
                </span>
                <span v-if="countdown.minutes > 0" class="time-unit">
                  <span class="time-value">{{ countdown.minutes }}</span>
                  <span class="time-label">分</span>
                </span>
                <span class="time-unit">
                  <span class="time-value">{{ countdown.seconds }}</span>
                  <span class="time-label">秒</span>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Filters -->
    <section class="filters-section">
      <div class="filters-container">
        <div class="filter-group">
          <select 
            v-model="filters.category_id" 
            @change="loadArticles"
            class="apple-select"
          >
            <option value="">所有分类</option>
            <option 
              v-for="category in categories" 
              :key="category.id"
              :value="category.id"
            >
              {{ category.name }}
            </option>
          </select>
          
          <select 
            v-model="filters.tag_id" 
            @change="loadArticles"
            class="apple-select"
          >
            <option value="">所有标签</option>
            <option 
              v-for="tag in tags" 
              :key="tag.id"
              :value="tag.id"
            >
              {{ tag.name }}
            </option>
          </select>
        </div>
      </div>
    </section>

    <!-- Articles Grid -->
    <section class="articles-section">
      <!-- View Toggle -->
      <div class="view-toggle-container">
        <button 
          class="view-toggle-btn" 
          :class="{ active: viewMode === 'grid' }"
          @click="viewMode = 'grid'"
          title="网格视图"
        >
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="3" width="7" height="7"></rect>
            <rect x="14" y="3" width="7" height="7"></rect>
            <rect x="14" y="14" width="7" height="7"></rect>
            <rect x="3" y="14" width="7" height="7"></rect>
          </svg>
        </button>
        <button 
          class="view-toggle-btn" 
          :class="{ active: viewMode === 'list' }"
          @click="viewMode = 'list'"
          title="列表视图"
        >
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="8" y1="6" x2="21" y2="6"></line>
            <line x1="8" y1="12" x2="21" y2="12"></line>
            <line x1="8" y1="18" x2="21" y2="18"></line>
            <line x1="3" y1="6" x2="3.01" y2="6"></line>
            <line x1="3" y1="12" x2="3.01" y2="12"></line>
            <line x1="3" y1="18" x2="3.01" y2="18"></line>
          </svg>
        </button>
      </div>
      <div v-if="loading" class="loading-container">
        <div class="loading-spinner"></div>
        <p class="loading-text">加载中...</p>
      </div>

      <div v-else-if="articles.length === 0" class="empty-state">
        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" opacity="0.3">
          <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path>
          <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path>
        </svg>
        <p>暂无文章</p>
      </div>

      <div v-else :class="['articles-container', viewMode === 'list' ? 'articles-list' : 'articles-grid']">
        <article
          v-for="article in articles"
          :key="article.id"
          class="article-card"
          @click="router.push(`/article/${article.id}`)"
        >
          <div v-if="article.cover_image" class="article-cover">
            <img 
              :src="getCoverImageUrl(article.cover_image)" 
              :alt="article.title"
              class="cover-image"
              @error="handleImageError"
            />
            <div class="cover-overlay"></div>
          </div>
          <div class="article-body">
            <div class="article-tags">
              <span v-if="article.category" class="tag-primary">{{ article.category.name }}</span>
              <span 
                v-for="tag in article.tags?.slice(0, 2)" 
                :key="tag.id"
                class="tag-secondary"
              >
                {{ tag.name }}
              </span>
              <span v-if="hasVideo(article.content)" class="tag-video" title="包含视频">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polygon points="23 7 16 12 23 17 23 7"></polygon>
                  <rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect>
                </svg>
                视频
              </span>
            </div>
            <h2 class="article-title">{{ article.title }}</h2>
            <p class="article-summary">
              {{ article.summary || getPlainTextSummary(article.content) }}
            </p>
            <div class="article-footer">
              <div class="article-author">
                <div class="author-avatar">
                  {{ article.author?.username?.charAt(0).toUpperCase() || 'U' }}
                </div>
                <span class="author-name">
                  {{ article.author?.nickname || article.author?.username }}
                </span>
              </div>
              <div class="article-meta">
                <span class="meta-item">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="10"></circle>
                    <polyline points="12 6 12 12 16 14"></polyline>
                  </svg>
                  {{ formatDate(article.created_at) }}
                </span>
                <span class="meta-item">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                    <circle cx="12" cy="12" r="3"></circle>
                  </svg>
                  {{ article.view_count }}
                </span>
              </div>
            </div>
          </div>
        </article>
      </div>

      <!-- Pagination -->
      <div v-if="total > pagination.page_size" class="pagination-container">
        <button 
          class="pagination-btn"
          :disabled="pagination.page === 1"
          @click="goToPage(pagination.page - 1)"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15 18 9 12 15 6"></polyline>
          </svg>
          上一页
        </button>
        
        <div class="pagination-info">
          第 {{ pagination.page }} / {{ Math.ceil(total / pagination.page_size) }} 页
        </div>
        
        <button 
          class="pagination-btn"
          :disabled="pagination.page >= Math.ceil(total / pagination.page_size)"
          @click="goToPage(pagination.page + 1)"
        >
          下一页
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="9 18 15 12 9 6"></polyline>
          </svg>
        </button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { articleAPI } from '@/api/article'
import { categoryAPI } from '@/api/category'
import { tagAPI } from '@/api/tag'

const router = useRouter()
const loading = ref(false)
const articles = ref([])
const categories = ref([])
const tags = ref([])
const total = ref(0)

const pagination = ref({
  page: 1,
  page_size: 12
})

const filters = ref({
  category_id: '',
  tag_id: ''
})

const viewMode = ref('list') // 'grid' 或 'list'，默认横向列表

// 恋爱纪念日配置（请修改为你的纪念日）
const anniversaryDate = new Date('2025-10-25T01:00:00') // 2025年10月25日凌晨1点
const countdown = ref({
  days: 0,
  hours: 0,
  minutes: 0,
  seconds: 0
})
const anniversaryYears = ref(0)
let countdownTimer = null

// 计算纪念日相关数据
const calculateAnniversary = () => {
  const now = new Date()
  const thisYear = now.getFullYear()
  
  // 计算年份数
  const yearsDiff = thisYear - anniversaryDate.getFullYear()
  const thisYearAnniversary = new Date(thisYear, anniversaryDate.getMonth(), anniversaryDate.getDate())
  const nextYearAnniversary = new Date(thisYear + 1, anniversaryDate.getMonth(), anniversaryDate.getDate())
  
  // 如果今年的纪念日还没到，年份数减1
  if (now < thisYearAnniversary) {
    anniversaryYears.value = Math.max(0, yearsDiff - 1)
  } else {
    anniversaryYears.value = yearsDiff
  }
  
  // 计算倒计时（到下一个纪念日）
  let targetDate = thisYearAnniversary
  if (now >= thisYearAnniversary) {
    targetDate = nextYearAnniversary
    anniversaryYears.value = yearsDiff + 1
  }
  
  const diff = targetDate - now
  
  if (diff <= 0) {
    // 今天就是纪念日
    countdown.value = { days: 0, hours: 0, minutes: 0, seconds: 0 }
  } else {
    countdown.value.days = Math.floor(diff / (1000 * 60 * 60 * 24))
    countdown.value.hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
    countdown.value.minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
    countdown.value.seconds = Math.floor((diff % (1000 * 60)) / 1000)
  }
}

const loadArticles = async () => {
  try {
    loading.value = true
    const params = { ...pagination.value }
    if (filters.value.category_id) params.category_id = filters.value.category_id
    if (filters.value.tag_id) params.tag_id = filters.value.tag_id
    
    const res = await articleAPI.getArticles(params)
    articles.value = res.articles || []
    total.value = res.total || 0
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadCategories = async () => {
  try {
    const res = await categoryAPI.getCategories()
    categories.value = res.categories || []
  } catch (error) {
    console.error(error)
  }
}

const loadTags = async () => {
  try {
    const res = await tagAPI.getTags()
    tags.value = res.tags || []
  } catch (error) {
    console.error(error)
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  if (days === 0) return '今天'
  if (days === 1) return '昨天'
  if (days < 7) return `${days}天前`
  if (days < 30) return `${Math.floor(days / 7)}周前`
  
  return date.toLocaleDateString('zh-CN', { month: 'short', day: 'numeric' })
}

// 检查文章内容是否包含视频
const hasVideo = (content) => {
  if (!content) return false
  // 检查是否包含video标签
  return /<video[\s\S]*?>/i.test(content) || /<source[\s\S]*?>/i.test(content)
}

// 从Markdown/HTML内容中提取纯文本摘要
const getPlainTextSummary = (content) => {
  if (!content) return ''
  
  // 移除HTML标签（包括video、img等）
  let text = content.replace(/<[^>]*>/g, '')
  
  // 移除Markdown代码块
  text = text.replace(/```[\s\S]*?```/g, '')
  text = text.replace(/`[^`]+`/g, '')
  
  // 移除Markdown链接语法 [text](url)
  text = text.replace(/\[([^\]]+)\]\([^\)]+\)/g, '$1')
  
  // 移除Markdown图片语法 ![alt](url) 或 ![alt](url "title")
  text = text.replace(/!\[([^\]]*)\]\([^\)]+\)/g, '')
  
  // 移除图片URL（包括各种格式）
  // 匹配类似 http://..., https://..., /uploads/..., 文件名.jpg 等
  text = text.replace(/https?:\/\/[^\s]+/g, '') // HTTP/HTTPS URL
  text = text.replace(/\/uploads\/[^\s]+/g, '') // 相对路径
  text = text.replace(/[^\s]+\.(jpg|jpeg|png|gif|webp|svg|bmp|ico)(\?[^\s]*)?/gi, '') // 图片文件名
  
  // 移除仅包含文件名的行（如：1000001097.jpg）
  text = text.replace(/^[\w\-]+\.(jpg|jpeg|png|gif|webp|svg|bmp|ico)\s*$/gim, '')
  
  // 移除Markdown标题标记
  text = text.replace(/^#{1,6}\s+/gm, '')
  
  // 移除Markdown列表标记
  text = text.replace(/^[\*\-\+]\s+/gm, '')
  text = text.replace(/^\d+\.\s+/gm, '')
  
  // 移除多余的空白字符和换行
  text = text.replace(/\s+/g, ' ').trim()
  
  // 截取前120个字符
  if (text.length > 120) {
    text = text.substring(0, 120) + '...'
  }
  
  return text || '暂无摘要'
}

const getCoverImageUrl = (url) => {
  if (!url) return ''
  if (url.startsWith('/uploads')) {
    return `http://localhost:8080${url}`
  }
  return url
}

const handleImageError = (e) => {
  e.target.style.display = 'none'
}

// Hero图片处理
const heroImageLoaded = ref(false)
const heroImageSrc = ref('/hero-banner.jpg')

const handleHeroImageError = () => {
  // 如果图片加载失败，使用渐变背景
  heroImageLoaded.value = false
}

const handleHeroImageLoad = () => {
  // 图片加载成功
  heroImageLoaded.value = true
}

const handleDiaryTitleError = (e) => {
  // 如果恋爱日记标题图片加载失败，隐藏标题容器
  if (e.target) {
    e.target.style.display = 'none'
  }
}

const handleDecorError = (e) => {
  // 如果装饰图片加载失败，隐藏该图片
  if (e.target) {
    e.target.style.display = 'none'
  }
}

// 尝试加载图片（检查图片是否存在）
const checkHeroImage = () => {
  const img = new Image()
  img.onload = () => {
    heroImageLoaded.value = true
  }
  img.onerror = () => {
    heroImageLoaded.value = false
  }
  img.src = heroImageSrc.value
}

const goToPage = (page) => {
  pagination.value.page = page
  loadArticles()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(() => {
  loadArticles()
  loadCategories()
  loadTags()
  checkHeroImage()
  
  // 初始化纪念日计算
  calculateAnniversary()
  // 每秒更新倒计时
  countdownTimer = setInterval(() => {
    calculateAnniversary()
  }, 1000)
})

onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
})
</script>

<style scoped>
.home-container {
  width: 100%;
  max-width: 1440px;
  margin: 0 auto;
  padding: 0;
}

/* Hero Section - 温馨风格 */
.hero-section {
  width: 100%;
  margin-bottom: 0;
  overflow: hidden;
  position: relative;
}

/* 恋爱日记标题图片 */
.diary-title {
  position: absolute;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 10;
  width: 100%;
  text-align: center;
  pointer-events: none;
}

.diary-title-img {
  max-width: 500px;
  width: auto;
  height: auto;
  max-height: 150px;
  object-fit: contain;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.15));
  animation: gentleFloat 4s ease-in-out infinite;
}

@keyframes gentleFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
  }
}

/* 温馨装饰小图片 */
.decorative-images {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 8;
  pointer-events: none;
}

.decor-img {
  position: absolute;
  width: auto;
  height: auto;
  max-width: 120px;
  max-height: 120px;
  object-fit: contain;
  filter: drop-shadow(0 2px 8px rgba(0, 0, 0, 0.1));
  opacity: 0.85;
  animation: decorFloat 5s ease-in-out infinite;
}

/* 标题左侧装饰 */
.decor-left {
  top: 50%;
  left: 14%;
  transform: translateY(-50%) translateX(0px);
  max-width: 130px;
  max-height: 130px;
  animation: decorFloat 4s ease-in-out infinite;
  animation-delay: 0.5s;
}

/* 标题右侧装饰 */
.decor-right {
  top: 55%;
  right: 15%;
  transform: translateY(-50%) translateX(20px);
  max-width: 140px;
  max-height: 140px;
  animation: decorFloat 4.5s ease-in-out infinite;
  animation-delay: 1s;
}

/* 左上角装饰 */
.decor-top-left {
  top: 80px;
  left: 8%;
  max-width: 90px;
  max-height: 90px;
  animation: decorFloat 5s ease-in-out infinite;
  animation-delay: 0.3s;
}

/* 右上角装饰 */
.decor-top-right {
  top: 80px;
  right: 8%;
  max-width: 90px;
  max-height: 90px;
  animation: decorFloat 4.8s ease-in-out infinite;
  animation-delay: 0.7s;
}

/* 左下角装饰 */
.decor-bottom-left {
  bottom: 120px;
  left: 8%;
  max-width: 110px;
  max-height: 110px;
  animation: decorFloat 5.2s ease-in-out infinite;
  animation-delay: 1.2s;
}

/* 中部左侧装饰 */
.decor-mid-left {
  top: 35%;
  left: 5%;
  max-width: 100px;
  max-height: 100px;
  animation: decorFloatMidLeft 4.5s ease-in-out infinite;
  animation-delay: 0.8s;
}

/* 中部右侧装饰 */
.decor-mid-right {
  top: 35%;
  right: 5%;
  max-width: 100px;
  max-height: 100px;
  animation: decorFloatMidRight 5s ease-in-out infinite;
  animation-delay: 1.5s;
}

@keyframes decorFloat {
  0%, 100% {
    transform: translateY(0) translateX(0) rotate(0deg);
  }
  25% {
    transform: translateY(-6px) translateX(3px) rotate(2deg);
  }
  50% {
    transform: translateY(-10px) translateX(0) rotate(0deg);
  }
  75% {
    transform: translateY(-6px) translateX(-3px) rotate(-2deg);
  }
}

/* 为不同位置的装饰保留各自的transform */
.decor-left {
  animation-name: decorFloatLeft;
}

.decor-right {
  animation-name: decorFloatRight;
}

@keyframes decorFloatLeft {
  0%, 100% {
    transform: translateY(-50%) translateX(-20px) rotate(0deg);
  }
  50% {
    transform: translateY(-50%) translateX(-25px) rotate(3deg);
  }
}

@keyframes decorFloatRight {
  0%, 100% {
    transform: translateY(-50%) translateX(20px) rotate(0deg);
  }
  50% {
    transform: translateY(-50%) translateX(25px) rotate(-3deg);
  }
}

.decor-mid-left {
  animation-name: decorFloatMidLeft;
}

.decor-mid-right {
  animation-name: decorFloatMidRight;
}

@keyframes decorFloatMidLeft {
  0%, 100% {
    transform: translateY(0) translateX(0) rotate(0deg);
  }
  25% {
    transform: translateY(-8px) translateX(5px) rotate(3deg);
  }
  50% {
    transform: translateY(-12px) translateX(0) rotate(0deg);
  }
  75% {
    transform: translateY(-8px) translateX(-5px) rotate(-3deg);
  }
}

@keyframes decorFloatMidRight {
  0%, 100% {
    transform: translateY(0) translateX(0) rotate(0deg);
  }
  25% {
    transform: translateY(-8px) translateX(-5px) rotate(-3deg);
  }
  50% {
    transform: translateY(-12px) translateX(0) rotate(0deg);
  }
  75% {
    transform: translateY(-8px) translateX(5px) rotate(3deg);
  }
}

.hero-image-container {
  position: relative;
  width: 100%;
  height: 90vh;
  max-height: 900px;
  min-height: 600px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #ffeef8 0%, #fff5e6 50%, #ffe5e5 100%);
  padding: 10vh 10%;
}

.hero-image-container.hero-fallback {
  background: linear-gradient(135deg, #ffeef8 0%, #fff5e6 50%, #ffe5e5 100%);
}

.hero-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
  object-position: center;
  display: block;
  position: relative;
  z-index: 1;
}

/* 气泡装饰 */
.bubbles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  z-index: 0;
  pointer-events: none;
}

.bubble {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(
    135deg,
    rgba(255, 182, 193, 0.35) 0%,
    rgba(255, 218, 185, 0.28) 50%,
    rgba(255, 240, 248, 0.38) 100%
  );
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  border: 3px solid rgba(255, 255, 255, 0.6);
  box-shadow: 
    0 12px 48px rgba(255, 182, 193, 0.4),
    0 6px 24px rgba(255, 107, 157, 0.3),
    inset 0 3px 12px rgba(255, 255, 255, 0.6),
    inset 0 -3px 12px rgba(255, 182, 193, 0.2);
  animation: float 15s ease-in-out infinite;
}

.bubble::before {
  content: '';
  position: absolute;
  top: 18%;
  left: 20%;
  width: 35%;
  height: 35%;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.7);
  filter: blur(6px);
  box-shadow: 0 0 20px rgba(255, 255, 255, 0.8);
}

.bubble-1 {
  width: 100px;
  height: 100px;
  top: 10%;
  left: 8%;
  animation-delay: 0s;
  animation-duration: 12s;
}

.bubble-2 {
  width: 80px;
  height: 80px;
  top: 20%;
  right: 12%;
  animation-delay: 2s;
  animation-duration: 18s;
}

.bubble-3 {
  width: 120px;
  height: 120px;
  bottom: 15%;
  left: 10%;
  animation-delay: 4s;
  animation-duration: 20s;
}

.bubble-4 {
  width: 90px;
  height: 90px;
  bottom: 25%;
  right: 8%;
  animation-delay: 1s;
  animation-duration: 16s;
}

.bubble-5 {
  width: 70px;
  height: 70px;
  top: 50%;
  left: 5%;
  animation-delay: 3s;
  animation-duration: 14s;
}

.bubble-6 {
  width: 110px;
  height: 110px;
  top: 30%;
  right: 5%;
  animation-delay: 5s;
  animation-duration: 22s;
}

.bubble-7 {
  width: 85px;
  height: 85px;
  bottom: 40%;
  left: 15%;
  animation-delay: 2.5s;
  animation-duration: 17s;
}

.bubble-8 {
  width: 75px;
  height: 75px;
  top: 60%;
  right: 15%;
  animation-delay: 4.5s;
  animation-duration: 19s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0) translateX(0) scale(1);
    opacity: 0.85;
  }
  25% {
    transform: translateY(-30px) translateX(20px) scale(1.05);
    opacity: 0.95;
  }
  50% {
    transform: translateY(-20px) translateX(-15px) scale(0.95);
    opacity: 0.9;
  }
  75% {
    transform: translateY(-40px) translateX(10px) scale(1.02);
    opacity: 0.92;
  }
}

.hero-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    to bottom,
    rgba(255, 240, 248, 0.3) 0%,
    rgba(255, 245, 230, 0.2) 50%,
    rgba(255, 229, 229, 0.4) 100%
  );
  pointer-events: none;
  z-index: 2;
}

.hero-content {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
  z-index: 3;
  max-width: 900px;
  padding: 0 24px;
}

.hero-title {
  font-size: 64px;
  font-weight: 600;
  letter-spacing: -0.02em;
  color: #fff;
  margin-bottom: 20px;
  line-height: 1.1;
  text-shadow: 0 2px 20px rgba(0, 0, 0, 0.15);
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.1));
}

.hero-subtitle {
  font-size: 24px;
  font-weight: 400;
  color: rgba(255, 255, 255, 0.95);
  letter-spacing: 0.01em;
  line-height: 1.5;
  text-shadow: 0 2px 12px rgba(0, 0, 0, 0.2);
}

/* 恋爱纪念日卡片 */
.anniversary-card {
  position: absolute;
  bottom: 30px;
  right: 30px;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 18px;
  padding: 16px 20px;
  box-shadow: 
    0 8px 32px rgba(255, 182, 193, 0.25),
    0 4px 16px rgba(255, 182, 193, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
  border: 1.5px solid rgba(255, 182, 193, 0.35);
  z-index: 4;
  min-width: 220px;
  max-width: 240px;
  animation: floatIn 1s ease-out;
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.anniversary-card:hover {
  transform: translateY(-4px) scale(1.02);
  box-shadow: 
    0 16px 56px rgba(255, 182, 193, 0.4),
    0 8px 28px rgba(255, 182, 193, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.9);
}

@keyframes floatIn {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.9);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.anniversary-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 42px;
  height: 42px;
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border-radius: 50%;
  margin: 0 auto 12px;
  box-shadow: 
    0 6px 20px rgba(255, 107, 157, 0.35),
    inset 0 2px 6px rgba(255, 255, 255, 0.3);
  color: #fff;
  animation: heartbeat 2s ease-in-out infinite;
}

.anniversary-icon svg {
  width: 22px;
  height: 22px;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

@keyframes heartbeat {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

.anniversary-icon svg {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.anniversary-info {
  text-align: center;
}

.anniversary-years {
  font-size: 22px;
  font-weight: 600;
  color: #ff6b9d;
  margin-bottom: 10px;
  letter-spacing: -0.02em;
  text-shadow: 0 2px 6px rgba(255, 107, 157, 0.2);
}

.anniversary-countdown {
  font-size: 11px;
  color: #86868b;
  line-height: 1.5;
}

.countdown-text {
  margin-bottom: 6px;
  font-weight: 500;
  color: #8b5a3c;
  font-size: 11px;
}

.countdown-time {
  display: flex;
  justify-content: center;
  gap: 6px;
  margin-top: 6px;
  flex-wrap: wrap;
}

.time-unit {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 32px;
}

.time-value {
  font-size: 18px;
  font-weight: 600;
  color: #ff6b9d;
  line-height: 1;
  margin-bottom: 2px;
  text-shadow: 0 2px 4px rgba(255, 107, 157, 0.2);
}

.time-label {
  font-size: 10px;
  color: #86868b;
  font-weight: 400;
  letter-spacing: 0.3px;
}

/* Filters - 温馨风格 */
.filters-section {
  padding: 30px 22px;
  background: transparent;
  border-bottom: none;
  position: relative;
}

.filters-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 30% 50%, rgba(255, 182, 193, 0.06) 0%, transparent 60%),
    radial-gradient(circle at 70% 30%, rgba(255, 218, 185, 0.05) 0%, transparent 60%);
  pointer-events: none;
}

.filters-container {
  max-width: 1200px;
  margin: 0 auto;
}

.filter-group {
  display: flex;
  gap: 16px;
  justify-content: center;
  position: relative;
  z-index: 1;
}

.apple-select {
  padding: 12px 24px;
  font-size: 15px;
  color: #8b5a3c;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 182, 193, 0.3);
  border-radius: 980px;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg width='12' height='8' viewBox='0 0 12 8' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1L6 6L11 1' stroke='%23ffb6c1' stroke-width='2' stroke-linecap='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 18px center;
  padding-right: 48px;
  box-shadow: 
    0 4px 16px rgba(255, 182, 193, 0.12),
    0 2px 8px rgba(255, 182, 193, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.apple-select:hover {
  border-color: rgba(255, 182, 193, 0.5);
  background: rgba(255, 255, 255, 0.95);
  box-shadow: 
    0 6px 24px rgba(255, 182, 193, 0.2),
    0 3px 12px rgba(255, 182, 193, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.6);
  transform: translateY(-2px);
}

.apple-select:focus {
  outline: none;
  border-color: #ff6b9d;
  background: rgba(255, 255, 255, 1);
  box-shadow: 
    0 0 0 4px rgba(255, 107, 157, 0.15),
    0 6px 24px rgba(255, 182, 193, 0.25),
    inset 0 1px 0 rgba(255, 255, 255, 0.7);
}

/* Articles Section - 温馨风格 */
.articles-section {
  padding: 20px 22px 80px 22px;
  background: transparent;
  position: relative;
}

.loading-container,
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120px 22px;
  color: #86868b;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(255, 182, 193, 0.2);
  border-top-color: #ffb6c1;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.loading-text {
  font-size: 17px;
  color: #86868b;
}

.empty-state p {
  margin-top: 16px;
  font-size: 17px;
  color: #86868b;
}

/* Articles Grid */
.articles-container {
  max-width: 1200px;
  margin: 0 auto;
}

.articles-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 32px;
}

.articles-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.articles-list .article-card {
  display: flex;
  flex-direction: row;
  max-width: 100%;
  height: auto;
}

.articles-list .article-cover {
  width: 320px;
  height: 200px;
  flex-shrink: 0;
  border-radius: 20px 0 0 20px;
  margin-right: 0;
}

.articles-list .article-body {
  flex: 1;
  padding: 28px 32px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.articles-list .article-title {
  font-size: 28px;
  margin-bottom: 12px;
  -webkit-line-clamp: 2;
  line-clamp: 2;
}

.articles-list .article-summary {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 16px;
  -webkit-line-clamp: 3;
  line-clamp: 3;
}

.articles-list .article-footer {
  margin-top: auto;
  padding-top: 0;
}

/* View Toggle */
.view-toggle-container {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-bottom: 32px;
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  padding: 0 22px;
}

.view-toggle-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  border-radius: 12px;
  border: 2px solid rgba(255, 182, 193, 0.3);
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  color: #8b5a3c;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  box-shadow: 
    0 2px 8px rgba(255, 182, 193, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.view-toggle-btn:hover {
  border-color: rgba(255, 182, 193, 0.5);
  background: rgba(255, 255, 255, 1);
  transform: translateY(-2px);
  box-shadow: 
    0 4px 12px rgba(255, 182, 193, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.6);
}

.view-toggle-btn.active {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border-color: #ff6b9d;
  color: #fff;
  box-shadow: 
    0 4px 16px rgba(255, 107, 157, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
}

.view-toggle-btn.active:hover {
  background: linear-gradient(135deg, #ff8fab 0%, #ff6b9d 100%);
  box-shadow: 
    0 6px 20px rgba(255, 107, 157, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
}

.article-card {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  border: 2px solid rgba(255, 182, 193, 0.25);
  position: relative;
  box-shadow: 
    0 8px 32px rgba(255, 182, 193, 0.12),
    0 4px 16px rgba(255, 182, 193, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.6);
}

.article-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    135deg, 
    rgba(255, 182, 193, 0.08) 0%, 
    rgba(255, 218, 185, 0.06) 50%,
    rgba(255, 240, 248, 0.08) 100%
  );
  opacity: 0;
  transition: opacity 0.5s;
  pointer-events: none;
  border-radius: 24px;
}

.article-card::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(
    circle,
    rgba(255, 255, 255, 0.3) 0%,
    transparent 70%
  );
  opacity: 0;
  transition: opacity 0.5s;
  pointer-events: none;
}

.article-card:hover {
  transform: translateY(-12px) scale(1.01);
  box-shadow: 
    0 24px 64px rgba(255, 182, 193, 0.25),
    0 12px 32px rgba(255, 182, 193, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
  border-color: rgba(255, 182, 193, 0.4);
}

.article-card:hover::before {
  opacity: 1;
}

.article-card:hover::after {
  opacity: 1;
  animation: shimmer 2s ease-in-out infinite;
}

@keyframes shimmer {
  0%, 100% { transform: translate(-50%, -50%) rotate(0deg); }
  50% { transform: translate(-50%, -50%) rotate(180deg); }
}

.article-cover {
  width: 100%;
  height: 240px;
  overflow: hidden;
  position: relative;
  background: linear-gradient(135deg, #ffeef8 0%, #fff5e6 50%, #ffe5e5 100%);
}

.cover-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.8s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  filter: brightness(1) saturate(1.05);
}

.article-card:hover .cover-image {
  transform: scale(1.08);
  filter: brightness(1.05) saturate(1.1);
}

.cover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    to bottom, 
    rgba(255, 240, 248, 0.15) 0%, 
    transparent 40%,
    rgba(255, 182, 193, 0.1) 100%
  );
  pointer-events: none;
  transition: opacity 0.5s;
}

.article-card:hover .cover-overlay {
  opacity: 0.8;
}

.article-body {
  padding: 28px;
  position: relative;
  z-index: 1;
}

.article-tags {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
  flex-wrap: wrap;
}

.tag-primary {
  padding: 8px 16px;
  font-size: 12px;
  font-weight: 500;
  color: #ff6b9d;
  background: linear-gradient(135deg, rgba(255, 107, 157, 0.15) 0%, rgba(255, 182, 193, 0.12) 100%);
  border-radius: 20px;
  letter-spacing: -0.01em;
  border: 1.5px solid rgba(255, 107, 157, 0.25);
  box-shadow: 
    0 2px 8px rgba(255, 107, 157, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  transition: all 0.3s;
}

.tag-primary:hover {
  transform: translateY(-1px);
  box-shadow: 
    0 4px 12px rgba(255, 107, 157, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.tag-secondary {
  padding: 8px 16px;
  font-size: 12px;
  font-weight: 400;
  color: #d4a574;
  background: linear-gradient(135deg, rgba(255, 218, 185, 0.35) 0%, rgba(255, 240, 248, 0.25) 100%);
  border-radius: 20px;
  letter-spacing: -0.01em;
  border: 1.5px solid rgba(255, 218, 185, 0.4);
  box-shadow: 
    0 2px 8px rgba(255, 218, 185, 0.12),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  transition: all 0.3s;
}

.tag-secondary:hover {
  transform: translateY(-1px);
  box-shadow: 
    0 4px 12px rgba(255, 218, 185, 0.18),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.tag-video {
  padding: 8px 16px;
  font-size: 12px;
  font-weight: 500;
  color: #ff6b9d;
  background: linear-gradient(135deg, rgba(255, 107, 157, 0.15) 0%, rgba(255, 182, 193, 0.12) 100%);
  border-radius: 20px;
  letter-spacing: -0.01em;
  border: 1.5px solid rgba(255, 107, 157, 0.25);
  box-shadow: 
    0 2px 8px rgba(255, 107, 157, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  transition: all 0.3s;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  flex-shrink: 0;
}

.tag-video:hover {
  transform: translateY(-1px);
  box-shadow: 
    0 4px 12px rgba(255, 107, 157, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.tag-video svg {
  flex-shrink: 0;
}

.article-title {
  font-size: 24px;
  font-weight: 600;
  letter-spacing: -0.011em;
  color: #1d1d1f;
  margin-bottom: 12px;
  line-height: 1.16667;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-summary {
  font-size: 15px;
  line-height: 1.47;
  color: #86868b;
  margin-bottom: 20px;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 182, 193, 0.2);
  position: relative;
}

.article-footer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 20%;
  right: 20%;
  height: 1px;
  background: linear-gradient(
    to right,
    transparent 0%,
    rgba(255, 182, 193, 0.4) 50%,
    transparent 100%
  );
}

.article-author {
  display: flex;
  align-items: center;
  gap: 8px;
}

.author-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ffb6c1 0%, #ffa07a 50%, #ffb347 100%);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
  box-shadow: 
    0 4px 12px rgba(255, 182, 193, 0.4),
    0 2px 6px rgba(255, 182, 193, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  transition: all 0.3s;
  border: 2px solid rgba(255, 255, 255, 0.5);
}

.author-avatar:hover {
  transform: scale(1.1);
  box-shadow: 
    0 6px 16px rgba(255, 182, 193, 0.5),
    0 3px 8px rgba(255, 182, 193, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
}

.author-name {
  font-size: 13px;
  font-weight: 400;
  color: #1d1d1f;
}

.article-meta {
  display: flex;
  gap: 12px;
  align-items: center;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #86868b;
}

.meta-item svg {
  opacity: 0.6;
}

/* Pagination */
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 24px;
  margin-top: 80px;
  padding: 40px 0;
}

.pagination-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 12px 28px;
  font-size: 15px;
  font-weight: 500;
  color: #8b5a3c;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 182, 193, 0.3);
  border-radius: 980px;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  box-shadow: 
    0 4px 16px rgba(255, 182, 193, 0.12),
    0 2px 8px rgba(255, 182, 193, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.pagination-btn:hover:not(:disabled) {
  background: rgba(255, 255, 255, 1);
  border-color: rgba(255, 182, 193, 0.5);
  transform: translateY(-3px);
  box-shadow: 
    0 8px 24px rgba(255, 182, 193, 0.2),
    0 4px 12px rgba(255, 182, 193, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.7);
}

.pagination-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.pagination-info {
  font-size: 14px;
  color: #86868b;
  font-weight: 400;
}

/* Responsive */
@media (max-width: 768px) {
  .diary-title {
    top: 15px;
  }

  .diary-title-img {
    max-width: 280px;
    max-height: 100px;
  }

  /* 移动端装饰图片调整 */
  .decor-img {
    max-width: 70px;
    max-height: 70px;
    opacity: 0.75;
  }

  .decor-left,
  .decor-right {
    max-width: 75px;
    max-height: 75px;
    top: 45%;
  }

  .decor-left {
    left: 10%;
    transform: translateY(-50%) translateX(0px);
  }

  .decor-right {
    right: 10%;
    transform: translateY(-50%) translateX(10px);
    top: 52%;
    max-width: 90px;
    max-height: 90px;
  }

  .decor-top-left,
  .decor-top-right {
    max-width: 50px;
    max-height: 50px;
    top: 60px;
  }

  .decor-top-left {
    left: 5%;
  }

  .decor-top-right {
    right: 5%;
  }

  .decor-bottom-left {
    bottom: 100px;
    left: 5%;
    max-width: 65px;
    max-height: 65px;
  }

  .decor-mid-left,
  .decor-mid-right {
    top: 35%;
    max-width: 60px;
    max-height: 60px;
  }

  .decor-mid-left {
    left: 3%;
  }

  .decor-mid-right {
    right: 3%;
  }

  .hero-image-container {
    height: 85vh;
    min-height: 500px;
    padding: 5vh 5%;
  }
  
  .hero-title {
    font-size: 36px;
  }

  .hero-subtitle {
    font-size: 18px;
  }

  .anniversary-card {
    bottom: 15px;
    right: 15px;
    padding: 12px 16px;
    min-width: 180px;
    max-width: 200px;
    background: rgba(255, 255, 255, 0.88);
    opacity: 0.9;
  }

  .anniversary-icon {
    width: 36px;
    height: 36px;
    margin-bottom: 8px;
  }

  .anniversary-icon svg {
    width: 18px;
    height: 18px;
  }

  .anniversary-years {
    font-size: 18px;
    margin-bottom: 8px;
  }

  .anniversary-countdown {
    font-size: 10px;
  }

  .countdown-text {
    font-size: 10px;
    margin-bottom: 4px;
  }

  .time-value {
    font-size: 16px;
  }

  .time-label {
    font-size: 9px;
  }

  .countdown-time {
    gap: 4px;
    margin-top: 4px;
  }

  .time-unit {
    min-width: 28px;
  }

  .articles-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }

  .filter-group {
    flex-direction: column;
  }
}
</style>
