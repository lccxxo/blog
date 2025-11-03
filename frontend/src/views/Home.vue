<template>
  <div class="home-container">
    <!-- Hero Section -->
    <section class="hero-section">
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
          <p class="hero-subtitle">把每一天写成诗，把每一刻都记住</p>
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
            </div>
            <h2 class="article-title">{{ article.title }}</h2>
            <p class="article-summary">
              {{ article.summary || article.content.substring(0, 120) + '...' }}
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
import { ref, onMounted } from 'vue'
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
}

.hero-image-container {
  position: relative;
  width: 100%;
  height: 80vh;
  max-height: 800px;
  min-height: 500px;
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

/* Filters - 温馨风格 */
.filters-section {
  padding: 60px 22px;
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
  padding: 80px 22px;
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
  .hero-image-container {
    height: 70vh;
    min-height: 400px;
    padding: 5vh 5%;
  }
  
  .hero-title {
    font-size: 36px;
  }

  .hero-subtitle {
    font-size: 18px;
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
