<template>
  <div class="my-articles-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>我的文章</span>
          <el-button type="primary" @click="router.push('/create-article')">
            <el-icon><Plus /></el-icon>
            创建文章
          </el-button>
        </div>
      </template>

      <el-table :data="articles" v-loading="loading" stripe>
        <el-table-column label="封面" width="120">
          <template #default="{ row }">
            <el-image 
              v-if="row.cover_image"
              :src="getCoverImageUrl(row.cover_image)" 
              fit="cover"
              class="table-cover-image"
              :preview-src-list="[getCoverImageUrl(row.cover_image)]"
            >
              <template #error>
                <div class="small-image-error">
                  <el-icon><Picture /></el-icon>
                </div>
              </template>
            </el-image>
            <span v-else style="color: #c0c4cc;">无封面</span>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题" min-width="200" />
        <el-table-column label="分类" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.category" size="small">{{ row.category.name }}</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="标签" width="200">
          <template #default="{ row }">
            <el-tag 
              v-for="tag in row.tags" 
              :key="tag.id"
              size="small"
              type="info"
              style="margin-right: 5px"
            >
              {{ tag.name }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_published ? 'success' : 'info'" size="small">
              {{ row.is_published ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="浏览量" width="100">
          <template #default="{ row }">
            {{ row.view_count }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="router.push(`/article/${row.id}`)">查看</el-button>
            <el-button size="small" type="primary" @click="router.push(`/edit-article/${row.id}`)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-if="total > 0"
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.page_size"
        :total="total"
        layout="prev, pager, next, total"
        @current-change="loadArticles"
        class="pagination"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { articleAPI } from '@/api/article'

const router = useRouter()
const loading = ref(false)
const articles = ref([])
const total = ref(0)

const pagination = ref({
  page: 1,
  page_size: 10
})

const loadArticles = async () => {
  try {
    loading.value = true
    const res = await articleAPI.getMyArticles(pagination.value)
    articles.value = res.articles || []
    total.value = res.total || 0
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return date.toLocaleString('zh-CN')
}

// 获取封面图片URL
const getCoverImageUrl = (url) => {
  if (!url) return ''
  // 如果是相对路径，添加完整的URL
  if (url.startsWith('/uploads')) {
    return `http://localhost:8080${url}`
  }
  return url
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这篇文章吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await articleAPI.deleteArticle(id)
    ElMessage.success('删除成功')
    loadArticles()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

onMounted(() => {
  loadArticles()
})
</script>

<style scoped>
.my-articles-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 22px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 24px;
  font-weight: 600;
  color: #1d1d1f;
  letter-spacing: -0.011em;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 32px;
}

.table-cover-image {
  width: 80px;
  height: 60px;
  border-radius: 8px;
  cursor: pointer;
  border: 1px solid rgba(255, 182, 193, 0.2);
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.1);
}

.small-image-error {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 80px;
  height: 60px;
  background: rgba(255, 240, 248, 0.5);
  border-radius: 8px;
  font-size: 24px;
  color: rgba(255, 182, 193, 0.6);
  border: 1px solid rgba(255, 182, 193, 0.2);
}

:deep(.el-table) {
  border: 1px solid rgba(255, 182, 193, 0.2);
  border-radius: 12px;
  overflow: hidden;
}

:deep(.el-table th) {
  background: rgba(255, 240, 248, 0.4);
  border-color: rgba(255, 182, 193, 0.2);
  color: #1d1d1f;
  font-weight: 600;
}

:deep(.el-table td) {
  border-color: rgba(255, 182, 193, 0.15);
}

:deep(.el-table--striped .el-table__body tr.el-table__row--striped td) {
  background: rgba(255, 240, 248, 0.2);
}

:deep(.el-table__row:hover) {
  background: rgba(255, 240, 248, 0.3);
}
</style>



