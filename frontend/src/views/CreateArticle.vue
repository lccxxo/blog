<template>
  <div class="create-article-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>{{ isEdit ? '编辑文章' : '创建文章' }}</span>
        </div>
      </template>
      <el-form :model="articleForm" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="文章标题" prop="title">
          <el-input v-model="articleForm.title" placeholder="请输入文章标题" />
        </el-form-item>
        <el-form-item label="文章摘要" prop="summary">
          <el-input 
            v-model="articleForm.summary" 
            type="textarea" 
            :rows="3"
            placeholder="请输入文章摘要（可选）" 
          />
        </el-form-item>
        <el-form-item label="文章内容" prop="content">
          <MarkdownEditor v-model="articleForm.content" />
        </el-form-item>
        <el-form-item label="封面图片" prop="cover_image">
          <div class="cover-image-upload">
            <el-input 
              v-model="articleForm.cover_image" 
              placeholder="请输入封面图片URL或上传图片（可选）"
              style="margin-bottom: 10px;"
            />
            <div class="upload-actions">
              <el-upload
                ref="coverUploadRef"
                :auto-upload="false"
                :show-file-list="false"
                :on-change="handleCoverImageSelect"
                accept="image/*"
              >
                <el-button size="small" :loading="coverUploading">
                  <el-icon><Upload /></el-icon>
                  上传封面图片
                </el-button>
              </el-upload>
              <el-button 
                v-if="articleForm.cover_image" 
                size="small" 
                type="info"
                @click="previewCover = true"
              >
                <el-icon><View /></el-icon>
                预览
              </el-button>
            </div>
            <el-image
              v-if="articleForm.cover_image"
              :src="getCoverImageUrl(articleForm.cover_image)"
              fit="cover"
              class="cover-preview"
              :preview-src-list="[getCoverImageUrl(articleForm.cover_image)]"
            />
          </div>
        </el-form-item>
        <el-form-item label="选择分类" prop="category_id">
          <el-select v-model="articleForm.category_id" placeholder="请选择分类" style="width: 100%">
            <el-option
              v-for="category in categories"
              :key="category.id"
              :label="category.name"
              :value="category.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="选择标签" prop="tag_ids">
          <el-select 
            v-model="articleForm.tag_ids" 
            multiple 
            placeholder="请选择标签"
            style="width: 100%"
          >
            <el-option
              v-for="tag in tags"
              :key="tag.id"
              :label="tag.name"
              :value="tag.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="发布状态">
          <el-switch 
            v-model="articleForm.is_published"
            active-text="已发布"
            inactive-text="草稿"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" :loading="loading">
            {{ isEdit ? '保存修改' : '发布文章' }}
          </el-button>
          <el-button @click="router.back()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { articleAPI } from '@/api/article'
import { categoryAPI } from '@/api/category'
import { tagAPI } from '@/api/tag'
import { uploadAPI } from '@/api/upload'
import MarkdownEditor from '@/components/MarkdownEditor.vue'

const router = useRouter()
const route = useRoute()
const formRef = ref()
const loading = ref(false)
const categories = ref([])
const tags = ref([])
const coverUploadRef = ref()
const coverUploading = ref(false)
const previewCover = ref(false)

const isEdit = ref(!!route.params.id)

const articleForm = ref({
  title: '',
  content: '',
  summary: '',
  cover_image: '',
  category_id: null,
  tag_ids: [],
  is_published: false
})

const rules = {
  title: [
    { required: true, message: '请输入文章标题', trigger: 'blur' }
  ],
  content: [
    { required: true, message: '请输入文章内容', trigger: 'blur' }
  ]
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

const loadArticle = async () => {
  if (!isEdit.value) return
  
  try {
    loading.value = true
    const res = await articleAPI.getArticle(route.params.id)
    const article = res.article
    articleForm.value = {
      title: article.title,
      content: article.content,
      summary: article.summary || '',
      cover_image: article.cover_image || '',
      category_id: article.category_id || null,
      tag_ids: article.tags?.map(tag => tag.id) || [],
      is_published: article.is_published
    }
  } catch (error) {
    console.error(error)
    ElMessage.error('文章加载失败')
    router.back()
  } finally {
    loading.value = false
  }
}

const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    loading.value = true
    
    if (isEdit.value) {
      await articleAPI.updateArticle(route.params.id, articleForm.value)
      ElMessage.success('文章更新成功')
    } else {
      await articleAPI.createArticle(articleForm.value)
      ElMessage.success('文章创建成功')
    }
    router.push('/my-articles')
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

// 处理封面图片选择
const handleCoverImageSelect = async (file) => {
  if (!file || !file.raw) return

  // 验证文件类型
  const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp']
  if (!allowedTypes.includes(file.raw.type)) {
    ElMessage.error('不支持的图片格式，请上传 JPG、PNG、GIF 或 WEBP 格式')
    return
  }

  // 文件大小限制已移除

  try {
    coverUploading.value = true
    const res = await uploadAPI.uploadImage(file.raw)
    articleForm.value.cover_image = res.url
    ElMessage.success('封面图片上传成功')
  } catch (error) {
    console.error('封面图片上传失败:', error)
    ElMessage.error(error.response?.data?.error || '封面图片上传失败')
  } finally {
    coverUploading.value = false
  }
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

onMounted(() => {
  loadCategories()
  loadTags()
  loadArticle()
})
</script>

<style scoped>
.create-article-container {
  max-width: 900px;
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

.cover-image-upload {
  width: 100%;
}

.upload-actions {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  flex-wrap: wrap;
}

.cover-preview {
  width: 100%;
  max-width: 400px;
  max-height: 300px;
  border-radius: 12px;
  margin-top: 16px;
  border: 1.5px solid rgba(255, 182, 193, 0.3);
  box-shadow: 0 4px 12px rgba(255, 182, 193, 0.15);
}

:deep(.el-form-item__label) {
  color: #1d1d1f;
  font-weight: 500;
  font-size: 15px;
}

:deep(.el-form-item) {
  margin-bottom: 28px;
}

:deep(.el-form-item__error) {
  color: #ff6b9d;
  font-size: 13px;
}

:deep(.el-button) {
  border-radius: 980px;
}

:deep(.el-button--primary) {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8fab 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(255, 107, 157, 0.3);
}

:deep(.el-button--primary:hover) {
  background: linear-gradient(135deg, #ff8fab 0%, #ff6b9d 100%);
  box-shadow: 0 6px 16px rgba(255, 107, 157, 0.4);
}

:deep(.el-button:not(.el-button--primary)) {
  border: 1.5px solid rgba(255, 182, 193, 0.4);
  color: #ff6b9d;
  background: #fff;
}

:deep(.el-button:not(.el-button--primary):hover) {
  border-color: rgba(255, 182, 193, 0.6);
  background: rgba(255, 240, 248, 0.5);
  color: #ff8fab;
}

:deep(.el-switch.is-checked .el-switch__core) {
  background-color: #ff6b9d;
}

:deep(.el-switch__core) {
  border-color: rgba(255, 182, 193, 0.5);
}

:deep(.el-upload) {
  border-radius: 980px;
}

:deep(.el-upload .el-button) {
  background: rgba(255, 240, 248, 0.5);
  border-color: rgba(255, 182, 193, 0.4);
  color: #ff6b9d;
}

:deep(.el-upload .el-button:hover) {
  background: rgba(255, 240, 248, 0.8);
  border-color: rgba(255, 182, 193, 0.6);
}
</style>

