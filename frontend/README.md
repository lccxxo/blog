# 博客系统前端

基于 Vue 3 + Vite + Element Plus 的博客系统前端应用

## 技术栈

- Vue 3 - 渐进式 JavaScript 框架
- Vite - 下一代前端构建工具
- Vue Router - 官方路由管理器
- Pinia - Vue 状态管理
- Element Plus - Vue 3 组件库
- Axios - HTTP 客户端

## 功能特性

- ✅ 用户注册和登录
- ✅ 文章浏览和搜索
- ✅ 文章创建和编辑
- ✅ 分类和标签管理
- ✅ 个人文章管理
- ✅ 响应式设计

## 快速开始

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

应用将在 `http://localhost:3000` 启动

### 构建生产版本

```bash
npm run build
```

### 预览生产构建

```bash
npm run preview
```

## 项目结构

```
frontend/
├── src/
│   ├── api/           # API 接口
│   ├── assets/        # 静态资源
│   ├── components/    # 组件
│   ├── router/        # 路由配置
│   ├── stores/        # 状态管理
│   ├── views/         # 页面视图
│   ├── App.vue        # 根组件
│   └── main.js        # 入口文件
├── index.html         # HTML 模板
├── vite.config.js     # Vite 配置
└── package.json       # 依赖配置
```

## 页面说明

### 公开页面
- **首页** (`/`) - 文章列表，支持分类和标签筛选
- **文章详情** (`/article/:id`) - 查看文章详细内容
- **分类列表** (`/categories`) - 查看所有分类
- **标签列表** (`/tags`) - 查看所有标签
- **登录** (`/login`) - 用户登录
- **注册** (`/register`) - 用户注册

### 需要登录的页面
- **创建文章** (`/create-article`) - 发布新文章
- **编辑文章** (`/edit-article/:id`) - 编辑已有文章
- **我的文章** (`/my-articles`) - 管理个人文章
- **个人信息** (`/profile`) - 查看个人资料

## API 代理配置

开发环境下，所有 `/api` 请求会被代理到 `http://localhost:8080`

可以在 `vite.config.js` 中修改代理配置

## 环境要求

- Node.js >= 16
- npm >= 8

## 注意事项

- 确保后端服务已在 8080 端口启动
- 登录后 Token 会保存在 localStorage 中
- 路由使用了权限守卫，未登录用户无法访问需要认证的页面





