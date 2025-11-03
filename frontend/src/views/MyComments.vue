<template>
  <div class="my-comments-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>我的评论</span>
        </div>
      </template>

      <el-table :data="comments" v-loading="loading" stripe>
        <el-table-column label="评论内容" min-width="300">
          <template #default="{ row }">
            <div class="comment-content">{{ row.content }}</div>
          </template>
        </el-table-column>
        <el-table-column label="所属文章" min-width="200">
          <template #default="{ row }">
            <el-link 
              v-if="row.article" 
              type="primary" 
              @click="router.push(`/article/${row.article.id}`)"
            >
              {{ row.article.title }}
            </el-link>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="评论时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button 
              size="small" 
              @click="viewArticle(row.article?.id)"
              :disabled="!row.article"
            >
              查看文章
            </el-button>
            <el-button 
              size="small" 
              type="danger" 
              @click="handleDelete(row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-if="total > 0"
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.page_size"
        :total="total"
        layout="prev, pager, next, total"
        @current-change="loadComments"
        class="pagination"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { commentAPI } from '@/api/comment'

const router = useRouter()
const loading = ref(false)
const comments = ref([])
const total = ref(0)

const pagination = ref({
  page: 1,
  page_size: 10
})

const loadComments = async () => {
  try {
    loading.value = true
    const res = await commentAPI.getMyComments(pagination.value)
    comments.value = res.comments || []
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

const viewArticle = (articleId) => {
  if (articleId) {
    router.push(`/article/${articleId}`)
  }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这条评论吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await commentAPI.deleteComment(id)
    ElMessage.success('删除成功')
    loadComments()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

onMounted(() => {
  loadComments()
})
</script>

<style scoped>
.my-comments-container {
  max-width: 1200px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 18px;
  font-weight: bold;
}

.comment-content {
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}
</style>





