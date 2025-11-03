<template>
  <div class="markdown-viewer" v-html="renderedContent"></div>
</template>

<script setup>
import { computed } from 'vue'
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'

const props = defineProps({
  content: {
    type: String,
    default: ''
  }
})

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

// 自定义渲染器，处理图片URL
const renderer = new marked.Renderer()
const originalImageRenderer = renderer.image.bind(renderer)

renderer.image = function(href, title, text) {
  // 如果是相对路径，添加完整的URL
  if (href && href.startsWith('/uploads')) {
    href = `http://localhost:8080${href}`
  }
  return originalImageRenderer(href, title, text)
}

marked.use({ renderer })

const renderedContent = computed(() => {
  return marked(props.content || '')
})
</script>

<style scoped>
.markdown-viewer {
  line-height: 1.73684;
  letter-spacing: 0.009em;
  color: #1d1d1f;
}

/* Markdown 样式 - 苹果风格 */
.markdown-viewer :deep(h1) {
  font-size: 40px;
  font-weight: 600;
  letter-spacing: -0.011em;
  margin: 48px 0 24px;
  line-height: 1.08349;
  color: #1d1d1f;
}

.markdown-viewer :deep(h2) {
  font-size: 32px;
  font-weight: 600;
  letter-spacing: -0.011em;
  margin: 40px 0 20px;
  line-height: 1.0625;
  color: #1d1d1f;
}

.markdown-viewer :deep(h3) {
  font-size: 24px;
  font-weight: 600;
  letter-spacing: -0.011em;
  margin: 32px 0 16px;
  line-height: 1.16667;
  color: #1d1d1f;
}

.markdown-viewer :deep(h4) {
  font-size: 21px;
  font-weight: 600;
  letter-spacing: -0.011em;
  margin: 28px 0 14px;
  line-height: 1.19048;
  color: #1d1d1f;
}

.markdown-viewer :deep(p) {
  margin: 24px 0;
  line-height: 1.73684;
  letter-spacing: 0.009em;
  color: #1d1d1f;
}

.markdown-viewer :deep(code) {
  background: rgba(255, 182, 193, 0.15);
  padding: 3px 8px;
  border-radius: 6px;
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Fira Code', 'Droid Sans Mono', 'Source Code Pro', monospace;
  font-size: 0.85em;
  color: #ff6b9d;
  font-weight: 500;
  border: 1px solid rgba(255, 182, 193, 0.2);
}

.markdown-viewer :deep(pre) {
  background: rgba(255, 240, 248, 0.5);
  padding: 24px;
  border-radius: 12px;
  overflow-x: auto;
  margin: 32px 0;
  border: 1px solid rgba(255, 182, 193, 0.3);
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.1);
}

.markdown-viewer :deep(pre code) {
  background: transparent;
  padding: 0;
  border-radius: 0;
  color: #1d1d1f;
  font-weight: 400;
}

.markdown-viewer :deep(blockquote) {
  border-left: 3px solid #ff6b9d;
  padding: 16px 24px;
  margin: 32px 0;
  background: rgba(255, 182, 193, 0.1);
  border-radius: 0 8px 8px 0;
  color: #1d1d1f;
  font-style: normal;
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.08);
}

.markdown-viewer :deep(ul),
.markdown-viewer :deep(ol) {
  padding-left: 1.5em;
  margin: 24px 0;
}

.markdown-viewer :deep(li) {
  margin: 12px 0;
  line-height: 1.73684;
  letter-spacing: 0.009em;
}

.markdown-viewer :deep(a) {
  color: #ff6b9d;
  text-decoration: none;
  font-weight: 500;
  transition: opacity 0.2s;
}

.markdown-viewer :deep(a:hover) {
  opacity: 0.7;
  text-decoration: underline;
  color: #ff8fab;
}

.markdown-viewer :deep(img) {
  max-width: 100%;
  height: auto;
  border-radius: 12px;
  margin: 32px 0;
  box-shadow: 0 4px 20px rgba(255, 182, 193, 0.2);
  border: 1px solid rgba(255, 182, 193, 0.2);
}

.markdown-viewer :deep(table) {
  border-collapse: collapse;
  width: 100%;
  margin: 32px 0;
  display: block;
  overflow-x: auto;
  border-radius: 12px;
  overflow: hidden;
  border: 0.5px solid rgba(0, 0, 0, 0.1);
}

.markdown-viewer :deep(table th),
.markdown-viewer :deep(table td) {
  border: 0.5px solid rgba(0, 0, 0, 0.08);
  padding: 12px 16px;
}

.markdown-viewer :deep(table th) {
  background: rgba(255, 240, 248, 0.6);
  font-weight: 600;
  color: #1d1d1f;
  border-color: rgba(255, 182, 193, 0.3);
}

.markdown-viewer :deep(table tr:nth-child(even)) {
  background: rgba(0, 0, 0, 0.02);
}

.markdown-viewer :deep(hr) {
  border: none;
  border-top: 0.5px solid rgba(0, 0, 0, 0.1);
  margin: 48px 0;
}

.markdown-viewer :deep(strong) {
  font-weight: 600;
  color: #1d1d1f;
}

.markdown-viewer :deep(em) {
  font-style: italic;
  color: #1d1d1f;
}

.markdown-viewer :deep(del) {
  text-decoration: line-through;
  opacity: 0.6;
}

.markdown-viewer :deep(video) {
  max-width: 100%;
  height: auto;
  border-radius: 12px;
  margin: 32px 0;
  box-shadow: 0 4px 20px rgba(255, 182, 193, 0.2);
  border: 1px solid rgba(255, 182, 193, 0.2);
  background: #000;
}

.markdown-viewer :deep(video)::-webkit-media-controls-panel {
  background: rgba(255, 182, 193, 0.1);
}
</style>



