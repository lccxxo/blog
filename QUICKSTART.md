# 快速入门指南

这是一个完整的博客系统，包含后端 API 服务和前端 Web 应用。

## 📋 前置要求

- **Go** >= 1.22
- **Node.js** >= 16
- **npm** >= 8

## 🚀 一键启动（推荐）

### Windows 用户

双击 `start.bat` 或在命令行运行：
```cmd
start.bat
```

### Linux/Mac 用户

```bash
chmod +x start.sh
./start.sh
```

这将自动启动：
- 后端服务：http://localhost:8080
- 前端应用：http://localhost:3000

## 📝 手动启动

### 第一步：启动后端

```bash
# 安装后端依赖
go mod download

# 启动后端服务
go run main.go
```

后端将在 `http://localhost:8080` 运行

### 第二步：启动前端

打开新的终端窗口：

```bash
# 进入前端目录
cd frontend

# 安装前端依赖（首次运行需要）
npm install

# 启动前端开发服务器
npm run dev
```

前端将在 `http://localhost:3000` 运行

## 🎯 开始使用

### 1. 注册账号

访问 http://localhost:3000 并点击"注册"按钮：

- 用户名：your_username
- 邮箱：your@email.com
- 密码：至少6位
- 昵称：显示名称（可选）

### 2. 登录

使用刚注册的账号登录系统

### 3. 创建分类和标签

在使用前建议先创建一些分类和标签：

**创建分类：**
1. 点击顶部菜单"分类"
2. 点击"添加分类"
3. 输入分类名称和描述

**创建标签：**
1. 点击顶部菜单"标签"
2. 点击"添加标签"
3. 输入标签名称

### 4. 写第一篇文章

1. 点击右上角"写文章"按钮
2. 填写文章信息：
   - 标题：必填
   - 内容：必填
   - 摘要：可选，用于列表预览
   - 封面图片：可选
   - 选择分类：单选
   - 选择标签：多选
   - 发布状态：选择"已发布"或保存为"草稿"
3. 点击"发布文章"

### 5. 管理文章

- **查看所有文章**：访问首页
- **查看我的文章**：点击"我的文章"菜单
- **编辑文章**：在文章详情页或我的文章列表中点击"编辑"
- **删除文章**：在文章详情页或我的文章列表中点击"删除"

## 📱 主要功能

### 公开功能（无需登录）

- ✅ 浏览文章列表
- ✅ 按分类筛选文章
- ✅ 按标签筛选文章
- ✅ 查看文章详情
- ✅ 查看所有分类
- ✅ 查看所有标签

### 需要登录的功能

- ✅ 创建文章
- ✅ 编辑自己的文章
- ✅ 删除自己的文章
- ✅ 管理分类
- ✅ 管理标签
- ✅ 查看个人信息

## 🔧 测试 API

项目提供了 API 测试脚本：

**Windows (PowerShell):**
```powershell
.\test_api.ps1
```

**Linux/Mac (Bash):**
```bash
chmod +x test_api.sh
./test_api.sh
```

测试脚本会自动：
1. 注册测试用户
2. 登录获取 Token
3. 创建分类和标签
4. 创建测试文章
5. 查询文章列表

## 📊 项目结构说明

```
lccxxo-blog/
├── 后端文件
│   ├── config/         配置（端口、数据库、JWT）
│   ├── controllers/    业务逻辑控制器
│   ├── database/       数据库初始化
│   ├── middlewares/    中间件（认证等）
│   ├── models/         数据模型
│   ├── routes/         路由配置
│   ├── utils/          工具函数
│   └── main.go         入口文件
│
├── 前端文件
│   └── frontend/
│       ├── src/
│       │   ├── api/      API 接口封装
│       │   ├── router/   路由配置
│       │   ├── stores/   状态管理
│       │   └── views/    页面组件
│       └── package.json
│
└── 其他文件
    ├── start.bat       Windows 启动脚本
    ├── start.sh        Linux/Mac 启动脚本
    ├── test_api.ps1    API 测试脚本 (PowerShell)
    └── test_api.sh     API 测试脚本 (Bash)
```

## ⚙️ 配置说明

### 后端配置

编辑 `config/config.go`：

```go
Server: ServerConfig{
    Port: ":8080",  // 修改后端端口
},
JWT: JWTConfig{
    Secret:     "your-secret-key",  // ⚠️ 生产环境务必修改
    ExpireTime: 24 * time.Hour,     // Token 有效期
},
```

### 前端配置

编辑 `frontend/vite.config.js`：

```js
server: {
    port: 3000,  // 修改前端端口
    proxy: {
        '/api': {
            target: 'http://localhost:8080',  // 后端地址
        }
    }
}
```

## 🐛 常见问题

### 后端无法启动

1. 检查 Go 版本：`go version`
2. 重新安装依赖：`go mod download`
3. 检查 8080 端口是否被占用

### 前端无法启动

1. 检查 Node.js 版本：`node --version`
2. 删除 `frontend/node_modules` 重新安装：
   ```bash
   cd frontend
   rm -rf node_modules
   npm install
   ```
3. 检查 3000 端口是否被占用

### 前端无法连接后端

1. 确保后端已启动（访问 http://localhost:8080/api/categories 测试）
2. 检查浏览器控制台的错误信息
3. 检查是否有 CORS 错误

### 登录后 Token 失效

JWT Token 默认 24 小时过期，可在 `config/config.go` 中修改 `ExpireTime`

## 🔒 安全建议

⚠️ **生产环境部署前请务必：**

1. 修改 JWT Secret 密钥
2. 使用更安全的数据库（MySQL/PostgreSQL）
3. 配置 HTTPS
4. 添加请求速率限制
5. 配置生产环境的 CORS 策略
6. 启用日志记录
7. 配置防火墙

## 📚 更多信息

- 完整 API 文档：查看 `README.md`
- 前端文档：查看 `frontend/README.md`
- 问题反馈：提交 Issue

## 🎉 开始你的博客之旅！

现在你可以开始使用这个博客系统了。祝你写作愉快！✨





