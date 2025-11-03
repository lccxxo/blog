<template>
  <div class="tags-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>标签管理</span>
          <el-button v-if="userStore.isLoggedIn" type="primary" @click="showDialog()">
            <el-icon><Plus /></el-icon>
            添加标签
          </el-button>
        </div>
      </template>

      <el-table :data="tags" v-loading="loading" stripe>
        <el-table-column prop="name" label="标签名称" min-width="300" />
        <el-table-column label="创建时间" width="200">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column v-if="userStore.isLoggedIn" label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="showDialog(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog 
      v-model="dialogVisible" 
      :title="isEdit ? '编辑标签' : '添加标签'"
      width="400px"
    >
      <el-form :model="tagForm" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="tagForm.name" placeholder="请输入标签名称" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { tagAPI } from '@/api/tag'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const loading = ref(false)
const submitLoading = ref(false)
const tags = ref([])
const dialogVisible = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const formRef = ref()

const tagForm = ref({
  name: ''
})

const rules = {
  name: [
    { required: true, message: '请输入标签名称', trigger: 'blur' }
  ]
}

const loadTags = async () => {
  try {
    loading.value = true
    const res = await tagAPI.getTags()
    tags.value = res.tags || []
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

const showDialog = (tag = null) => {
  if (tag) {
    isEdit.value = true
    editId.value = tag.id
    tagForm.value = {
      name: tag.name
    }
  } else {
    isEdit.value = false
    editId.value = null
    tagForm.value = {
      name: ''
    }
  }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    submitLoading.value = true
    
    if (isEdit.value) {
      await tagAPI.updateTag(editId.value, tagForm.value)
      ElMessage.success('标签更新成功')
    } else {
      await tagAPI.createTag(tagForm.value)
      ElMessage.success('标签创建成功')
    }
    
    dialogVisible.value = false
    loadTags()
  } catch (error) {
    console.error(error)
  } finally {
    submitLoading.value = false
  }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个标签吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await tagAPI.deleteTag(id)
    ElMessage.success('删除成功')
    loadTags()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

onMounted(() => {
  loadTags()
})
</script>

<style scoped>
.tags-container {
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

:deep(.el-table--striped .el-table__body tr.el-table__row--striped.current-row td) {
  background-color: rgba(255, 182, 193, 0.15);
}

:deep(.el-dialog) {
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
}

:deep(.el-dialog__header) {
  background: rgba(255, 240, 248, 0.3);
  border-bottom: 1px solid rgba(255, 182, 193, 0.2);
}

:deep(.el-dialog__title) {
  color: #1d1d1f;
  font-weight: 600;
  font-size: 20px;
}
</style>





