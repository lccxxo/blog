<template>
  <div class="markdown-editor">
    <el-tabs v-model="activeTab" class="editor-tabs">
      <el-tab-pane label="编辑" name="edit">
        <el-input
          v-model="content"
          type="textarea"
          :rows="20"
          placeholder="支持 Markdown 语法..."
          @input="handleInput"
        />
      </el-tab-pane>
      <el-tab-pane label="预览" name="preview">
        <div class="markdown-preview" v-html="renderedContent"></div>
      </el-tab-pane>
    </el-tabs>
    
    <div class="markdown-toolbar">
      <el-space wrap>
        <el-button size="small" @click="insertMarkdown('# ', '')">H1</el-button>
        <el-button size="small" @click="insertMarkdown('## ', '')">H2</el-button>
        <el-button size="small" @click="insertMarkdown('### ', '')">H3</el-button>
        <el-divider direction="vertical" />
        <el-button size="small" @click="insertMarkdown('**', '**')">
          <strong>粗体</strong>
        </el-button>
        <el-button size="small" @click="insertMarkdown('*', '*')">
          <em>斜体</em>
        </el-button>
        <el-button size="small" @click="insertMarkdown('~~', '~~')">
          <s>删除线</s>
        </el-button>
        <el-divider direction="vertical" />
        <el-button size="small" @click="insertMarkdown('[', '](url)')">链接</el-button>
        <el-button size="small" @click="insertMarkdown('![', '](url)')">图片</el-button>
        <el-upload
          ref="uploadRef"
          :auto-upload="false"
          :show-file-list="false"
          :on-change="handleImageSelect"
          accept="image/*"
        >
          <el-button size="small" type="primary" :loading="uploading">
            <el-icon><Upload /></el-icon>
            上传图片
          </el-button>
        </el-upload>
        <el-upload
          ref="videoUploadRef"
          :auto-upload="false"
          :show-file-list="false"
          :on-change="handleVideoSelect"
          accept="video/*"
        >
          <el-button size="small" type="primary" :loading="videoUploading">
            <el-icon><VideoPlay /></el-icon>
            上传视频
          </el-button>
        </el-upload>
        <el-button size="small" @click="insertMarkdown('`', '`')">代码</el-button>
        <el-button size="small" @click="insertMarkdown('```\n', '\n```')">代码块</el-button>
        <el-divider direction="vertical" />
        <el-button size="small" @click="insertMarkdown('> ', '')">引用</el-button>
        <el-button size="small" @click="insertMarkdown('- ', '')">列表</el-button>
        <el-button size="small" @click="insertMarkdown('1. ', '')">数字列表</el-button>
      </el-space>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'
import { ElMessage } from 'element-plus'
import { VideoPlay } from '@element-plus/icons-vue'
import { uploadAPI } from '@/api/upload'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const activeTab = ref('edit')
const content = ref(props.modelValue)
const uploading = ref(false)
const videoUploading = ref(false)
const uploadRef = ref()
const videoUploadRef = ref()

// 配置 marked
marked.setOptions({
  highlight: function(code, lang) {
    if (lang && hljs.getLanguage(lang)) {
      return hljs.highlight(code, { language: lang }).value
    }
    return hljs.highlightAuto(code).value
  },
  breaks: true,
  gfm: true
})

// 配置marked以允许HTML（用于视频标签）
marked.use({
  mangle: false,
  headerIds: false
})

const renderedContent = computed(() => {
  return marked(content.value || '')
})

const handleInput = () => {
  emit('update:modelValue', content.value)
}

const insertMarkdown = (before, after) => {
  const textarea = document.querySelector('.markdown-editor textarea')
  if (!textarea) return

  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  const selectedText = content.value.substring(start, end)
  const newText = content.value.substring(0, start) + before + selectedText + after + content.value.substring(end)
  
  content.value = newText
  emit('update:modelValue', content.value)
  
  // 设置光标位置
  setTimeout(() => {
    textarea.focus()
    const newCursorPos = start + before.length + selectedText.length
    textarea.setSelectionRange(newCursorPos, newCursorPos)
  }, 0)
}

watch(() => props.modelValue, (newVal) => {
  content.value = newVal
})

// 处理图片选择
const handleImageSelect = async (file) => {
  if (!file || !file.raw) return

  // 验证文件类型
  const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp', 'image/svg+xml']
  if (!allowedTypes.includes(file.raw.type)) {
    ElMessage.error('不支持的图片格式，请上传 JPG、PNG、GIF、WEBP 或 SVG 格式')
    return
  }

  // 文件大小限制已移除

  try {
    uploading.value = true
    const res = await uploadAPI.uploadImage(file.raw, (percent) => {
      console.log(`上传进度: ${percent}%`)
    })

    // 插入图片Markdown语法
    const imageUrl = res.url
    const imageName = file.name
    const imageMarkdown = `![${imageName}](${imageUrl})`
    
    const textarea = document.querySelector('.markdown-editor textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      content.value = content.value.substring(0, start) + imageMarkdown + content.value.substring(end)
      emit('update:modelValue', content.value)
      
      // 设置光标位置
      setTimeout(() => {
        textarea.focus()
        const newCursorPos = start + imageMarkdown.length
        textarea.setSelectionRange(newCursorPos, newCursorPos)
      }, 0)
    }

    ElMessage.success('图片上传成功')
  } catch (error) {
    console.error('图片上传失败:', error)
    ElMessage.error(error.response?.data?.error || '图片上传失败')
  } finally {
    uploading.value = false
  }
}

// 处理视频选择
const handleVideoSelect = async (file) => {
  if (!file || !file.raw) return

  // 验证文件类型
  const allowedTypes = ['video/mp4', 'video/avi', 'video/quicktime', 'video/x-ms-wmv', 'video/x-flv', 'video/x-matroska', 'video/webm']
  if (!allowedTypes.includes(file.raw.type)) {
    ElMessage.error('不支持的视频格式，请上传 MP4、AVI、MOV、WMV、FLV、MKV 或 WEBM 格式')
    return
  }

  // 文件大小限制已移除

  try {
    videoUploading.value = true
    const res = await uploadAPI.uploadVideo(file.raw, (percent) => {
      console.log(`视频上传进度: ${percent}%`)
    })

    // 插入视频HTML标签（Markdown本身不支持视频，使用HTML）
    const videoUrl = res.url
    const videoName = file.name
    const videoType = file.raw.type || 'video/mp4'
    const videoHtml = `<video controls width="100%">\n  <source src="${videoUrl}" type="${videoType}">\n  您的浏览器不支持视频标签。\n</video>\n\n`
    
    const textarea = document.querySelector('.markdown-editor textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      content.value = content.value.substring(0, start) + videoHtml + content.value.substring(end)
      emit('update:modelValue', content.value)
      
      // 设置光标位置
      setTimeout(() => {
        textarea.focus()
        const newCursorPos = start + videoHtml.length
        textarea.setSelectionRange(newCursorPos, newCursorPos)
      }, 0)
    }

    ElMessage.success('视频上传成功')
  } catch (error) {
    console.error('视频上传失败:', error)
    ElMessage.error(error.response?.data?.error || '视频上传失败')
  } finally {
    videoUploading.value = false
  }
}
</script>

<style scoped>
.markdown-editor {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}

.editor-tabs {
  margin-bottom: 0;
}

.markdown-preview {
  min-height: 480px;
  padding: 20px;
  background: #fff;
  overflow-y: auto;
  max-height: 600px;
}

.markdown-toolbar {
  padding: 10px;
  background: #f5f7fa;
  border-top: 1px solid #dcdfe6;
}

/* Markdown 样式 */
.markdown-preview :deep(h1) {
  font-size: 2em;
  margin: 0.67em 0;
  border-bottom: 1px solid #eaecef;
  padding-bottom: 0.3em;
}

.markdown-preview :deep(h2) {
  font-size: 1.5em;
  margin: 0.75em 0;
  border-bottom: 1px solid #eaecef;
  padding-bottom: 0.3em;
}

.markdown-preview :deep(h3) {
  font-size: 1.25em;
  margin: 0.83em 0;
}

.markdown-preview :deep(h4) {
  font-size: 1em;
  margin: 1em 0;
}

.markdown-preview :deep(p) {
  margin: 1em 0;
  line-height: 1.8;
}

.markdown-preview :deep(code) {
  background: #f3f4f5;
  padding: 2px 6px;
  border-radius: 3px;
  font-family: 'Consolas', 'Monaco', monospace;
  font-size: 0.9em;
}

.markdown-preview :deep(pre) {
  background: #f6f8fa;
  padding: 16px;
  border-radius: 6px;
  overflow-x: auto;
  margin: 1em 0;
}

.markdown-preview :deep(pre code) {
  background: transparent;
  padding: 0;
  border-radius: 0;
}

.markdown-preview :deep(blockquote) {
  border-left: 4px solid #dfe2e5;
  padding-left: 16px;
  margin: 1em 0;
  color: #6a737d;
}

.markdown-preview :deep(ul),
.markdown-preview :deep(ol) {
  padding-left: 2em;
  margin: 1em 0;
}

.markdown-preview :deep(li) {
  margin: 0.5em 0;
}

.markdown-preview :deep(a) {
  color: #409eff;
  text-decoration: none;
}

.markdown-preview :deep(a:hover) {
  text-decoration: underline;
}

.markdown-preview :deep(img) {
  max-width: 100%;
  height: auto;
}

.markdown-preview :deep(table) {
  border-collapse: collapse;
  width: 100%;
  margin: 1em 0;
}

.markdown-preview :deep(table th),
.markdown-preview :deep(table td) {
  border: 1px solid #dfe2e5;
  padding: 8px 12px;
}

.markdown-preview :deep(table th) {
  background: #f6f8fa;
  font-weight: bold;
}

.markdown-preview :deep(hr) {
  border: none;
  border-top: 1px solid #eaecef;
  margin: 2em 0;
}
</style>



