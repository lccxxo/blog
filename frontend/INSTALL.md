# 前端安装说明

## 首次运行前端项目

由于前端依赖包体积较大，需要先安装依赖：

### 方法一：使用 npm（推荐）

```bash
# 进入前端目录
cd frontend

# 安装依赖（首次运行或 package.json 更新后）
npm install

# 启动开发服务器
npm run dev
```

### 方法二：使用 cnpm（国内用户推荐，速度更快）

```bash
# 安装 cnpm（如果还没安装）
npm install -g cnpm --registry=https://registry.npmmirror.com

# 进入前端目录
cd frontend

# 使用 cnpm 安装依赖
cnpm install

# 启动开发服务器
npm run dev
```

### 方法三：使用 yarn

```bash
# 安装 yarn（如果还没安装）
npm install -g yarn

# 进入前端目录
cd frontend

# 安装依赖
yarn install

# 启动开发服务器
yarn dev
```

## 常见问题

### 1. 安装速度慢

**解决方案：** 使用国内镜像源

```bash
# 临时使用淘宝镜像
npm install --registry=https://registry.npmmirror.com

# 或永久设置
npm config set registry https://registry.npmmirror.com
```

### 2. 安装失败

**解决方案：** 清除缓存后重新安装

```bash
# 删除 node_modules 和 package-lock.json
rm -rf node_modules package-lock.json

# 清除 npm 缓存
npm cache clean --force

# 重新安装
npm install
```

### 3. 端口被占用

**解决方案：** 修改端口

编辑 `vite.config.js`：
```js
server: {
    port: 3001,  // 改成其他端口
    ...
}
```

## 依赖说明

本项目主要依赖：

- **Vue 3** - 前端框架
- **Vite** - 构建工具
- **Element Plus** - UI 组件库
- **Vue Router** - 路由管理
- **Pinia** - 状态管理
- **Axios** - HTTP 客户端
- **marked** - Markdown 解析器
- **highlight.js** - 代码高亮

## 首次安装预计时间

- 使用 npm：5-10 分钟（取决于网络速度）
- 使用 cnpm：2-5 分钟
- 使用 yarn：3-7 分钟

## 安装完成后

依赖安装完成后，项目目录下会生成 `node_modules` 文件夹（约 200MB+），这是正常的。

之后每次启动项目只需要：

```bash
cd frontend
npm run dev
```

无需再次安装依赖，除非 `package.json` 有更新。





