# 登录权限问题诊断指南

## 🔍 问题诊断步骤

### 第一步：检查前端依赖是否已安装

前端依赖未安装会导致 Pinia store 无法正常工作。

```bash
# 进入前端目录
cd frontend

# 检查是否存在 node_modules 文件夹
# 如果不存在，执行：
npm install
```

### 第二步：检查后端是否正常运行

```bash
# 确保后端正在运行
go run main.go

# 应该看到输出：
# Database initialized successfully
# 服务器启动在端口 :8080
```

### 第三步：测试登录接口

打开浏览器控制台（F12），然后尝试登录，观察以下信息：

#### 1. 检查登录请求

在 Network（网络）标签中查看登录请求：
- 请求地址：`http://localhost:3000/api/auth/login`
- 请求方法：POST
- 状态码：应该是 200

#### 2. 检查响应数据

点击请求，查看 Response（响应）：
```json
{
  "message": "登录成功",
  "token": "eyJhbGciOiJIUzI1NiIs...",  // ← 应该有这个
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户"
  }
}
```

#### 3. 检查 LocalStorage

在 Application（应用）标签中：
- 展开 Local Storage → http://localhost:3000
- 应该看到：
  - `token`: "eyJhbGciOiJIUzI1NiIs..."
  - `user`: "{"id":1,"username":"testuser",...}"

### 第四步：检查后续请求

登录成功后，访问需要认证的页面（如创建文章），在 Network 中查看请求头：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

如果没有这个请求头，说明 token 没有被正确发送。

## 🐛 常见问题和解决方案

### 问题 1：登录后立即跳转到登录页

**原因**：前端依赖未安装，Pinia store 无法工作

**解决方案**：
```bash
cd frontend
npm install
# 重新启动前端
npm run dev
```

### 问题 2：显示"未登录或登录已过期"

**原因**：Token 过期或无效

**解决方案**：
1. 清除浏览器 LocalStorage
2. 重新注册和登录
3. 检查系统时间是否正确

### 问题 3：显示"token格式错误"

**原因**：请求头格式不正确

**解决方案**：
检查 `frontend/src/api/request.js` 中的请求拦截器：
```javascript
config.headers.Authorization = `Bearer ${userStore.token}`
```

### 问题 4：显示"无效的token"

**原因**：JWT 密钥不匹配或 token 损坏

**解决方案**：
1. 重启后端服务
2. 清除浏览器 LocalStorage
3. 重新登录

### 问题 5：CORS 错误

**原因**：跨域配置问题

**检查**：确保后端 `routes/routes.go` 中有 CORS 配置：
```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))
```

## 🔧 手动测试登录流程

### 使用 PowerShell 测试

```powershell
# 1. 注册用户
$body = @{
    username = "testuser"
    email = "test@example.com"
    password = "password123"
    nickname = "测试用户"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/auth/register" -Method Post -Body $body -ContentType "application/json"

# 2. 登录获取 token
$loginBody = @{
    username = "testuser"
    password = "password123"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8080/api/auth/login" -Method Post -Body $loginBody -ContentType "application/json"

# 3. 查看 token
$response.token

# 4. 使用 token 访问受保护的接口
$headers = @{
    "Authorization" = "Bearer $($response.token)"
    "Content-Type" = "application/json"
}

Invoke-RestMethod -Uri "http://localhost:8080/api/profile" -Headers $headers
```

### 使用浏览器控制台测试

```javascript
// 在浏览器控制台（F12）执行

// 1. 登录
fetch('/api/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    username: 'testuser',
    password: 'password123'
  })
})
.then(r => r.json())
.then(data => {
  console.log('登录响应:', data)
  // 保存 token
  localStorage.setItem('token', data.token)
  localStorage.setItem('user', JSON.stringify(data.user))
})

// 2. 测试受保护的接口
fetch('/api/profile', {
  headers: {
    'Authorization': `Bearer ${localStorage.getItem('token')}`
  }
})
.then(r => r.json())
.then(data => console.log('个人信息:', data))
```

## 🎯 快速修复步骤

如果以上都不行，尝试完全重置：

```bash
# 1. 停止所有服务

# 2. 删除数据库（重新开始）
rm blog.db

# 3. 清除浏览器数据
# 打开浏览器控制台（F12）
# Application → Local Storage → 右键删除全部

# 4. 重新启动后端
go run main.go

# 5. 重新启动前端
cd frontend
npm run dev

# 6. 重新注册账号
# 访问 http://localhost:3000/register

# 7. 登录
# 访问 http://localhost:3000/login
```

## 📋 检查清单

在提问前，请确认：

- [ ] 后端正在运行（端口 8080）
- [ ] 前端正在运行（端口 3000）
- [ ] 前端依赖已安装（存在 frontend/node_modules）
- [ ] 浏览器控制台没有错误
- [ ] 登录请求返回了 token
- [ ] LocalStorage 中保存了 token 和 user
- [ ] 后续请求带有 Authorization 请求头

## 💡 提供详细信息

如果问题仍未解决，请提供：

1. **浏览器控制台错误信息**（F12 → Console）
2. **网络请求详情**（F12 → Network → 点击失败的请求）
3. **后端日志输出**
4. **登录响应数据**
5. **LocalStorage 内容**

这样我才能更准确地帮您定位问题！





