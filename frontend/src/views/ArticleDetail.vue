<template>
  <div class="article-detail-container" v-loading="loading">
    <article v-if="article" class="article-content">
      <!-- Header -->
      <header class="article-header">
        <div class="article-tags">
          <span v-if="article.category" class="tag-primary">{{ article.category.name }}</span>
          <span 
            v-for="tag in article.tags" 
            :key="tag.id"
            class="tag-secondary"
          >
            {{ tag.name }}
          </span>
        </div>
        
        <h1 class="article-title">{{ article.title }}</h1>
        
        <div class="article-meta">
          <div class="meta-author">
            <div class="author-avatar-large">
              {{ article.author?.username?.charAt(0).toUpperCase() || 'U' }}
            </div>
            <div class="meta-author-info">
              <div class="author-name">{{ article.author?.nickname || article.author?.username }}</div>
              <div class="meta-details">
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
                  {{ article.view_count }} 次阅读
                </span>
              </div>
            </div>
          </div>
        </div>
      </header>

      <!-- Cover Image -->
      <div v-if="article.cover_image" class="article-cover-wrapper">
        <div class="cover-container">
          <img 
            :src="getCoverImageUrl(article.cover_image)" 
            :alt="article.title"
            class="article-cover-image"
            @error="handleImageError"
          />
        </div>
      </div>

      <!-- Content -->
      <div class="article-body">
        <MarkdownViewer :content="article.content" class="markdown-content" />
      </div>

      <!-- Actions -->
      <div v-if="isAuthor" class="article-actions">
        <button class="btn-secondary" @click="router.push(`/edit-article/${article.id}`)">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
          </svg>
          编辑文章
        </button>
        <button class="btn-danger" @click="handleDelete">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="3 6 5 6 21 6"></polyline>
            <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
          </svg>
          删除文章
        </button>
      </div>
    </article>

    <!-- Comments Section -->
    <section v-if="article" class="comments-section">
      <div class="comments-container">
        <h2 class="comments-title">评论</h2>
        <CommentSection :article-id="article.id" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { articleAPI } from '@/api/article'
import { useUserStore } from '@/stores/user'
import MarkdownViewer from '@/components/MarkdownViewer.vue'
import CommentSection from '@/components/CommentSection.vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const article = ref(null)

const isAuthor = computed(() => {
  return userStore.user && article.value && userStore.user.id === article.value.author_id
})

const loadArticle = async () => {
  try {
    loading.value = true
    const res = await articleAPI.getArticle(route.params.id)
    article.value = res.article
  } catch (error) {
    console.error(error)
    ElMessage.error('文章加载失败')
    router.push('/')
  } finally {
    loading.value = false
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
  
  return date.toLocaleDateString('zh-CN', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
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

const handleDelete = async () => {
  try {
    await ElMessageBox.confirm('确定要删除这篇文章吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await articleAPI.deleteArticle(article.value.id)
    ElMessage.success('删除成功')
    router.push('/')
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

onMounted(() => {
  loadArticle()
})
</script>

<style scoped>
.article-detail-container {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 60px 22px;
  position: relative;
}

.article-detail-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 20% 50%, rgba(255, 182, 193, 0.08) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(255, 218, 185, 0.08) 0%, transparent 50%),
    radial-gradient(circle at 40% 20%, rgba(255, 240, 248, 0.1) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

/* Article Content */
.article-content {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  padding: 60px 48px;
  border-radius: 24px;
  box-shadow: 
    0 8px 32px rgba(255, 182, 193, 0.12),
    0 4px 16px rgba(255, 182, 193, 0.08);
  border: 1.5px solid rgba(255, 182, 193, 0.2);
  position: relative;
  z-index: 1;
}

/* Header */
.article-header {
  margin-bottom: 48px;
}

.article-tags {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
  flex-wrap: wrap;
}

.tag-primary {
  padding: 6px 12px;
  font-size: 12px;
  font-weight: 500;
  color: #ff6b9d;
  background: rgba(255, 107, 157, 0.12);
  border-radius: 16px;
  letter-spacing: -0.01em;
  border: 1px solid rgba(255, 107, 157, 0.2);
}

.tag-secondary {
  padding: 6px 12px;
  font-size: 12px;
  font-weight: 400;
  color: #d4a574;
  background: rgba(255, 218, 185, 0.3);
  border-radius: 16px;
  letter-spacing: -0.01em;
  border: 1px solid rgba(255, 218, 185, 0.4);
}

.article-title {
  font-size: 48px;
  font-weight: 600;
  letter-spacing: -0.011em;
  color: #1d1d1f;
  margin-bottom: 24px;
  line-height: 1.08349;
}

.article-meta {
  display: flex;
  align-items: center;
  padding-top: 24px;
  border-top: 0.5px solid rgba(255, 182, 193, 0.2);
}

.meta-author {
  display: flex;
  align-items: center;
  gap: 16px;
}

.author-avatar-large {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ffb6c1 0%, #ffa07a 50%, #ffb347 100%);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 600;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.4);
}

.meta-author-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.author-name {
  font-size: 17px;
  font-weight: 500;
  color: #1d1d1f;
}

.meta-details {
  display: flex;
  gap: 16px;
  align-items: center;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #86868b;
  font-weight: 400;
}

.meta-item svg {
  opacity: 0.6;
}

/* Cover Image */
.article-cover-wrapper {
  margin: 48px 0;
}

.cover-container {
  width: 100%;
  border-radius: 24px;
  overflow: hidden;
  background: linear-gradient(135deg, #ffeef8 0%, #fff5e6 100%);
  box-shadow: 
    0 12px 48px rgba(255, 182, 193, 0.2),
    0 6px 24px rgba(255, 182, 193, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
  border: 2px solid rgba(255, 182, 193, 0.25);
  position: relative;
}

.cover-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.2) 0%,
    transparent 50%,
    rgba(255, 182, 193, 0.1) 100%
  );
  pointer-events: none;
  border-radius: 24px;
}

.article-cover-image {
  width: 100%;
  height: auto;
  display: block;
  object-fit: cover;
  filter: brightness(1) saturate(1.05);
  transition: filter 0.5s;
}

.cover-container:hover .article-cover-image {
  filter: brightness(1.02) saturate(1.08);
}

/* Body */
.article-body {
  margin: 48px 0;
  padding-top: 48px;
  border-top: 0.5px solid rgba(255, 182, 193, 0.2);
}

.markdown-content {
  font-size: 19px;
  line-height: 1.73684;
  letter-spacing: 0.009em;
  color: #1d1d1f;
}

/* Actions */
.article-actions {
  display: flex;
  gap: 12px;
  padding-top: 48px;
  margin-top: 48px;
  border-top: 0.5px solid rgba(255, 182, 193, 0.2);
}

.btn-secondary,
.btn-danger {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 400;
  border-radius: 980px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-secondary {
  color: #ff6b9d;
  background: #fff;
  border: 1.5px solid rgba(255, 182, 193, 0.4);
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.15);
}

.btn-secondary:hover {
  background: rgba(255, 240, 248, 0.5);
  border-color: rgba(255, 182, 193, 0.6);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(255, 182, 193, 0.25);
}

.btn-danger {
  color: #ff6b9d;
  background: #fff;
  border: 1.5px solid rgba(255, 182, 193, 0.4);
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.15);
}

.btn-danger:hover {
  background: rgba(255, 240, 248, 0.5);
  border-color: rgba(255, 107, 157, 0.6);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(255, 107, 157, 0.3);
}

/* Comments Section */
.comments-section {
  margin-top: 80px;
  padding-top: 60px;
  position: relative;
  z-index: 1;
}

.comments-container {
  max-width: 800px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  padding: 48px;
  border-radius: 24px;
  box-shadow: 
    0 8px 32px rgba(255, 182, 193, 0.12),
    0 4px 16px rgba(255, 182, 193, 0.08);
  border: 1.5px solid rgba(255, 182, 193, 0.2);
}

.comments-title {
  font-size: 32px;
  font-weight: 600;
  letter-spacing: -0.011em;
  color: #1d1d1f;
  margin-bottom: 32px;
}

/* Loading */
:deep(.el-loading-mask) {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}

/* Responsive */
@media (max-width: 768px) {
  .article-detail-container {
    padding: 0 16px;
  }

  .article-content {
    padding: 40px 0;
  }

  .article-title {
    font-size: 36px;
  }

  .article-cover-wrapper {
    margin: 32px 0;
  }

  .cover-container {
    border-radius: 12px;
  }

  .article-body {
    margin: 32px 0;
    padding-top: 32px;
  }

  .markdown-content {
    font-size: 17px;
  }

  .article-actions {
    flex-direction: column;
  }

  .btn-secondary,
  .btn-danger {
    width: 100%;
    justify-content: center;
  }
}
</style>
