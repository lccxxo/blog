# 阿里云OSS配置指南

## 配置说明

本系统已集成阿里云OSS对象存储服务，用于上传和存储图片、视频等文件。

## 配置步骤

### 1. 在config/config.go中配置OSS参数

修改 `config/config.go` 文件中的OSS配置：

```go
OSS: OSSConfig{
    Endpoint:        "oss-cn-hangzhou.aliyuncs.com", // OSS Endpoint，根据你的bucket地域选择
    AccessKeyID:     "your-access-key-id",           // 阿里云AccessKey ID
    AccessKeySecret: "your-access-key-secret",       // 阿里云AccessKey Secret
    BucketName:      "your-bucket-name",             // OSS Bucket名称
    Domain:          "",                             // 可选：自定义域名，如果设置了CDN域名
},
```

### 2. 如何获取Endpoint

**方法一：在阿里云OSS控制台查看**
1. 登录 [阿里云OSS控制台](https://oss.console.aliyun.com/)
2. 点击左侧"存储空间"，进入你的Bucket（例如：`zcc-mjh`）
3. 点击"概览"标签页
4. 在"访问域名"区域可以看到：
   - **外网访问Endpoint**: 例如 `oss-cn-hangzhou.aliyuncs.com`
   - **内网访问Endpoint**: 例如 `oss-cn-hangzhou-internal.aliyuncs.com`（仅同地域可用）
5. **复制外网Endpoint**（不带 `https://` 前缀），例如：`oss-cn-hangzhou.aliyuncs.com`

**方法二：根据Bucket地域确定**
根据你的OSS Bucket所在地域，Endpoint格式为：`oss-cn-<region>.aliyuncs.com`

常见地域Endpoint：
- 华东1（杭州）: `oss-cn-hangzhou.aliyuncs.com`
- 华北2（北京）: `oss-cn-beijing.aliyuncs.com`
- 华南1（深圳）: `oss-cn-shenzhen.aliyuncs.com`
- 华东2（上海）: `oss-cn-shanghai.aliyuncs.com`
- 华东5（南京-本地地域）: `oss-cn-nanjing.aliyuncs.com`
- 华北3（张家口）: `oss-cn-zhangjiakou.aliyuncs.com`
- 西南1（成都）: `oss-cn-chengdu.aliyuncs.com`
- 中国香港: `oss-cn-hongkong.aliyuncs.com`

**注意**：Endpoint只填写域名部分，不要包含 `https://` 或 `http://`

### 3. 获取AccessKey

1. 登录阿里云控制台
2. 进入"访问控制" -> "用户" -> 创建或选择用户
3. 为该用户创建AccessKey
4. 将AccessKey ID和AccessKey Secret填入配置

### 4. 创建Bucket

1. 登录阿里云OSS控制台
2. 创建Bucket，选择合适的地域
3. 设置读写权限（建议设置为"公共读"或"私有"）
4. 将Bucket名称填入配置

### 5. Domain（可选）

**Domain是用于访问文件的URL前缀，有两种选择：**

#### 选项A：不配置Domain（使用OSS默认域名）
如果你不配置Domain，系统会自动使用OSS的默认访问域名，格式为：
```
https://<BucketName>.<Endpoint>/<文件路径>
```
例如：`https://zcc-mjh.oss-cn-hangzhou.aliyuncs.com/image/2025/01/15/xxx.jpg`

**如何获取OSS默认访问域名：**
1. 登录阿里云OSS控制台
2. 进入你的Bucket（例如：`zcc-mjh`）
3. 点击"概览"标签页
4. 在"访问域名"区域可以看到完整的外网访问地址
5. 但**不需要填写**到Domain字段，系统会自动生成

#### 选项B：配置自定义域名/CDN域名（推荐）
如果你有：
- **自定义域名**：已在OSS控制台绑定到Bucket的域名
- **CDN加速域名**：通过阿里云CDN加速的域名

可以在`Domain`字段中配置：
```go
Domain: "https://cdn.example.com",  // 你的CDN域名或自定义域名
```

**如何绑定自定义域名：**
1. 在OSS控制台，进入你的Bucket
2. 点击"传输管理" -> "域名管理"
3. 点击"绑定用户域名"
4. 输入你的域名（需要在域名DNS中添加CNAME记录）
5. 绑定成功后，使用该域名访问文件

**Domain配置示例：**
```go
// 使用CDN域名（推荐，可加速访问）
Domain: "https://cdn.example.com"

// 或使用自定义域名
Domain: "https://files.example.com"

// 不使用自定义域名（留空即可）
Domain: ""
```

**建议**：
- 如果有CDN，建议配置CDN域名以加速访问
- 如果没有特殊需求，留空即可，使用OSS默认域名

## 使用环境变量（推荐）

为了安全，建议使用环境变量来配置敏感信息：

```go
OSS: OSSConfig{
    Endpoint:        os.Getenv("OSS_ENDPOINT"),
    AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
    AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
    BucketName:      os.Getenv("OSS_BUCKET_NAME"),
    Domain:          os.Getenv("OSS_DOMAIN"), // 可选
},
```

然后在启动应用前设置环境变量：
```bash
export OSS_ENDPOINT="oss-cn-hangzhou.aliyuncs.com"
export OSS_ACCESS_KEY_ID="your-access-key-id"
export OSS_ACCESS_KEY_SECRET="your-access-key-secret"
export OSS_BUCKET_NAME="your-bucket-name"
export OSS_DOMAIN="https://cdn.example.com"  # 可选
```

## 文件存储结构

上传的文件将按照以下结构存储在OSS中：
- 图片: `image/年/月/日/UUID.扩展名`
- 视频: `video/年/月/日/UUID.扩展名`

例如：
- `image/2025/01/15/123e4567-e89b-12d3-a456-426614174000.jpg`
- `video/2025/01/15/123e4567-e89b-12d3-a456-426614174000.mp4`

## API接口

### 上传图片
- **POST** `/api/upload/image`
- 需要认证
- 请求体: `multipart/form-data`，字段名: `image`
- 支持格式: jpg, jpeg, png, gif, webp, svg
- 大小限制: 5MB

### 上传视频
- **POST** `/api/upload/video`
- 需要认证
- 请求体: `multipart/form-data`，字段名: `video`
- 支持格式: mp4, avi, mov, wmv, flv, mkv, webm
- 大小限制: 100MB

### 删除文件
- **DELETE** `/api/upload/image`
- 需要认证
- 请求体: JSON `{"url": "文件URL"}`

## 注意事项

1. **安全性**: 请勿将AccessKey信息提交到代码仓库
2. **权限**: 建议使用RAM子账号，只授予OSS相关权限
3. **费用**: OSS按存储量和流量计费，请合理使用
4. **CORS**: 如果前端需要直接访问OSS，需要在OSS控制台配置CORS规则

## 故障排除

如果遇到上传失败：
1. 检查OSS配置是否正确
2. 检查AccessKey是否有相应权限
3. 检查Bucket是否存在且可访问
4. 检查网络连接是否正常
5. 查看服务器日志获取详细错误信息

