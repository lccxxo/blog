<template>
  <div class="categories-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>分类管理</span>
          <el-button v-if="userStore.isLoggedIn" type="primary" @click="showDialog()">
            <el-icon><Plus /></el-icon>
            添加分类
          </el-button>
        </div>
      </template>

      <el-table :data="categories" v-loading="loading" stripe>
        <el-table-column prop="name" label="分类名称" min-width="200" />
        <el-table-column prop="description" label="描述" min-width="300" />
        <el-table-column label="创建时间" width="180">
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
      :title="isEdit ? '编辑分类' : '添加分类'"
      width="500px"
    >
      <el-form :model="categoryForm" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="categoryForm.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input 
            v-model="categoryForm.description" 
            type="textarea"
            :rows="4"
            placeholder="请输入分类描述" 
          />
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
import { categoryAPI } from '@/api/category'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const loading = ref(false)
const submitLoading = ref(false)
const categories = ref([])
const dialogVisible = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const formRef = ref()

const categoryForm = ref({
  name: '',
  description: ''
})

const rules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' }
  ]
}

const loadCategories = async () => {
  try {
    loading.value = true
    const res = await categoryAPI.getCategories()
    categories.value = res.categories || []
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

const showDialog = (category = null) => {
  if (category) {
    isEdit.value = true
    editId.value = category.id
    categoryForm.value = {
      name: category.name,
      description: category.description || ''
    }
  } else {
    isEdit.value = false
    editId.value = null
    categoryForm.value = {
      name: '',
      description: ''
    }
  }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    submitLoading.value = true
    
    if (isEdit.value) {
      await categoryAPI.updateCategory(editId.value, categoryForm.value)
      ElMessage.success('分类更新成功')
    } else {
      await categoryAPI.createCategory(categoryForm.value)
      ElMessage.success('分类创建成功')
    }
    
    dialogVisible.value = false
    loadCategories()
  } catch (error) {
    console.error(error)
  } finally {
    submitLoading.value = false
  }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个分类吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await categoryAPI.deleteCategory(id)
    ElMessage.success('删除成功')
    loadCategories()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

onMounted(() => {
  loadCategories()
})
</script>

<style scoped>
.categories-container {
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





