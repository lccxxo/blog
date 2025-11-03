<template>
  <div class="comment-section">
    <div class="comment-header">
      <h3>评论 ({{ total }})</h3>
    </div>

    <!-- 发表评论表单 -->
    <el-card v-if="userStore.isLoggedIn" class="comment-form-card" shadow="never">
      <el-input
        v-model="newComment"
        type="textarea"
        :rows="4"
        placeholder="写下你的评论..."
        :disabled="submitting"
      />
      <div class="form-actions">
        <el-button 
          type="primary" 
          @click="submitComment"
          :loading="submitting"
          :disabled="!newComment.trim()"
        >
          发表评论
        </el-button>
      </div>
    </el-card>
    <el-alert v-else type="info" :closable="false" style="margin-bottom: 20px">
      <template #default>
        请先 <el-link type="primary" @click="router.push('/login')">登录</el-link> 后再发表评论
      </template>
    </el-alert>

    <!-- 评论列表 -->
    <div class="comments-list" v-loading="loading">
      <el-empty v-if="comments.length === 0 && !loading" description="暂无评论，快来抢沙发吧！" />
      
      <div v-for="comment in comments" :key="comment.id" class="comment-item">
        <div class="comment-avatar">
          <el-avatar :size="40">{{ comment.user?.username?.charAt(0).toUpperCase() }}</el-avatar>
        </div>
        <div class="comment-content">
          <div class="comment-meta">
            <span class="username">{{ comment.user?.nickname || comment.user?.username }}</span>
            <span class="date">{{ formatDate(comment.created_at) }}</span>
          </div>
          <div class="comment-text" v-if="editingCommentId !== comment.id">
            {{ comment.content }}
          </div>
          <div v-else class="comment-edit">
            <el-input
              v-model="editContent"
              type="textarea"
              :rows="3"
            />
            <div class="edit-actions">
              <el-button size="small" @click="cancelEdit">取消</el-button>
              <el-button size="small" type="primary" @click="saveEdit(comment.id)">保存</el-button>
            </div>
          </div>
          <div class="comment-actions">
            <el-button 
              text 
              size="small"
              @click="showReplyForm(comment.id)"
              v-if="userStore.isLoggedIn"
            >
              <el-icon><ChatLineRound /></el-icon>
              回复
            </el-button>
            <el-button 
              text 
              size="small"
              @click="startEdit(comment)"
              v-if="isMyComment(comment)"
            >
              <el-icon><Edit /></el-icon>
              编辑
            </el-button>
            <el-button 
              text 
              size="small"
              type="danger"
              @click="deleteComment(comment.id)"
              v-if="isMyComment(comment)"
            >
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </div>

          <!-- 回复表单 -->
          <div v-if="replyingTo === comment.id" class="reply-form">
            <el-input
              v-model="replyContent"
              type="textarea"
              :rows="3"
              :placeholder="`回复 ${comment.user?.nickname || comment.user?.username}...`"
            />
            <div class="form-actions">
              <el-button size="small" @click="cancelReply">取消</el-button>
              <el-button 
                size="small" 
                type="primary" 
                @click="submitReply(comment.id)"
                :loading="submitting"
                :disabled="!replyContent.trim()"
              >
                发表回复
              </el-button>
            </div>
          </div>

          <!-- 回复列表 -->
          <div v-if="comment.replies && comment.replies.length > 0" class="replies-list">
            <div v-for="reply in comment.replies" :key="reply.id" class="reply-item">
              <div class="comment-avatar">
                <el-avatar :size="32">{{ reply.user?.username?.charAt(0).toUpperCase() }}</el-avatar>
              </div>
              <div class="comment-content">
                <div class="comment-meta">
                  <span class="username">{{ reply.user?.nickname || reply.user?.username }}</span>
                  <span class="date">{{ formatDate(reply.created_at) }}</span>
                </div>
                <div class="comment-text">{{ reply.content }}</div>
                <div class="comment-actions">
                  <el-button 
                    text 
                    size="small"
                    type="danger"
                    @click="deleteComment(reply.id)"
                    v-if="isMyComment(reply)"
                  >
                    <el-icon><Delete /></el-icon>
                    删除
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { commentAPI } from '@/api/comment'
import { useUserStore } from '@/stores/user'

const props = defineProps({
  articleId: {
    type: Number,
    required: true
  }
})

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const submitting = ref(false)
const comments = ref([])
const total = ref(0)
const newComment = ref('')
const replyingTo = ref(null)
const replyContent = ref('')
const editingCommentId = ref(null)
const editContent = ref('')

const loadComments = async () => {
  try {
    loading.value = true
    const res = await commentAPI.getArticleComments(props.articleId)
    comments.value = res.comments || []
    total.value = res.total || 0
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const submitComment = async () => {
  try {
    submitting.value = true
    await commentAPI.createComment({
      content: newComment.value,
      article_id: props.articleId
    })
    ElMessage.success('评论发表成功')
    newComment.value = ''
    loadComments()
  } catch (error) {
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const showReplyForm = (commentId) => {
  replyingTo.value = commentId
  replyContent.value = ''
}

const cancelReply = () => {
  replyingTo.value = null
  replyContent.value = ''
}

const submitReply = async (parentId) => {
  try {
    submitting.value = true
    await commentAPI.createComment({
      content: replyContent.value,
      article_id: props.articleId,
      parent_id: parentId
    })
    ElMessage.success('回复发表成功')
    cancelReply()
    loadComments()
  } catch (error) {
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const startEdit = (comment) => {
  editingCommentId.value = comment.id
  editContent.value = comment.content
}

const cancelEdit = () => {
  editingCommentId.value = null
  editContent.value = ''
}

const saveEdit = async (commentId) => {
  try {
    await commentAPI.updateComment(commentId, { content: editContent.value })
    ElMessage.success('评论更新成功')
    cancelEdit()
    loadComments()
  } catch (error) {
    console.error(error)
  }
}

const deleteComment = async (commentId) => {
  try {
    await ElMessageBox.confirm('确定要删除这条评论吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await commentAPI.deleteComment(commentId)
    ElMessage.success('评论删除成功')
    loadComments()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

const isMyComment = (comment) => {
  return userStore.user && comment.user_id === userStore.user.id
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date
  
  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour
  
  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return `${Math.floor(diff / minute)} 分钟前`
  } else if (diff < day) {
    return `${Math.floor(diff / hour)} 小时前`
  } else if (diff < 7 * day) {
    return `${Math.floor(diff / day)} 天前`
  } else {
    return date.toLocaleDateString('zh-CN')
  }
}

onMounted(() => {
  loadComments()
})
</script>

<style scoped>
.comment-section {
  margin-top: 30px;
}

.comment-header h3 {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 24px;
  color: #1d1d1f;
  letter-spacing: -0.011em;
}

.comment-form-card {
  margin-bottom: 24px;
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
  background: rgba(255, 240, 248, 0.2);
}

.form-actions {
  margin-top: 16px;
  text-align: right;
}

.comments-list {
  margin-top: 24px;
}

.comment-item {
  display: flex;
  gap: 16px;
  padding: 24px 0;
  border-bottom: 1px solid rgba(255, 182, 193, 0.15);
  transition: background 0.2s;
  border-radius: 12px;
  padding-left: 12px;
  padding-right: 12px;
  margin-left: -12px;
  margin-right: -12px;
}

.comment-item:hover {
  background: rgba(255, 240, 248, 0.3);
}

.comment-item:last-child {
  border-bottom: none;
}

.comment-avatar {
  flex-shrink: 0;
}

.comment-content {
  flex: 1;
}

.comment-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.username {
  font-weight: 500;
  color: #1d1d1f;
  font-size: 15px;
}

.date {
  font-size: 13px;
  color: #86868b;
}

.comment-text {
  color: #1d1d1f;
  line-height: 1.8;
  margin-bottom: 12px;
  white-space: pre-wrap;
  word-wrap: break-word;
  font-size: 15px;
}

.comment-actions {
  display: flex;
  gap: 12px;
}

.comment-edit {
  margin-bottom: 12px;
}

.edit-actions {
  margin-top: 12px;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

.reply-form {
  margin-top: 16px;
  padding: 16px;
  background: rgba(255, 240, 248, 0.4);
  border-radius: 12px;
  border: 1px solid rgba(255, 182, 193, 0.2);
}

.replies-list {
  margin-top: 16px;
  padding-left: 24px;
  border-left: 2px solid rgba(255, 182, 193, 0.3);
}

.reply-item {
  display: flex;
  gap: 12px;
  padding: 16px 0;
}

.reply-item:first-child {
  padding-top: 0;
}

.reply-item:last-child {
  padding-bottom: 0;
}

:deep(.el-button--text) {
  color: #ff6b9d;
}

:deep(.el-button--text:hover) {
  color: #ff8fab;
  background: rgba(255, 240, 248, 0.3);
}

:deep(.el-button--text.is-disabled) {
  color: #c0c4cc;
}

:deep(.el-alert--info) {
  background-color: rgba(255, 240, 248, 0.5);
  border: 1px solid rgba(255, 182, 193, 0.3);
  border-radius: 12px;
}

:deep(.el-alert__title) {
  color: #1d1d1f;
}

:deep(.el-card) {
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
  background: rgba(255, 255, 255, 0.9);
}
</style>




