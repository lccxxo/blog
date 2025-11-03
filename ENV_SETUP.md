# 环境变量配置指南

## 问题说明

GitHub的推送保护功能检测到代码中包含敏感信息（阿里云AccessKey），因此阻止了推送。为了安全，所有敏感配置都应该通过环境变量设置，而不是硬编码在代码中。

## 解决方案

### 1. 设置环境变量

#### Windows PowerShell

```powershell
$env:OSS_ENDPOINT="oss-cn-beijing.aliyuncs.com"
$env:OSS_ACCESS_KEY_ID="你的AccessKey ID"
$env:OSS_ACCESS_KEY_SECRET="你的AccessKey Secret"
$env:OSS_BUCKET_NAME="zcc-mjh"
$env:OSS_DOMAIN=""
```

#### Windows CMD

```cmd
set OSS_ENDPOINT=oss-cn-beijing.aliyuncs.com
set OSS_ACCESS_KEY_ID=你的AccessKey ID
set OSS_ACCESS_KEY_SECRET=你的AccessKey Secret
set OSS_BUCKET_NAME=zcc-mjh
set OSS_DOMAIN=
```

#### Linux/Mac

```bash
export OSS_ENDPOINT="oss-cn-beijing.aliyuncs.com"
export OSS_ACCESS_KEY_ID="你的AccessKey ID"
export OSS_ACCESS_KEY_SECRET="你的AccessKey Secret"
export OSS_BUCKET_NAME="zcc-mjh"
export OSS_DOMAIN=""
```

### 2. 永久设置环境变量（推荐）

#### Windows

1. 右键"此电脑" -> 属性
2. 高级系统设置 -> 环境变量
3. 在"用户变量"中点击"新建"
4. 添加以下变量：
   - `OSS_ENDPOINT` = `oss-cn-beijing.aliyuncs.com`
   - `OSS_ACCESS_KEY_ID` = `你的AccessKey ID`
   - `OSS_ACCESS_KEY_SECRET` = `你的AccessKey Secret`
   - `OSS_BUCKET_NAME` = `zcc-mjh`
   - `OSS_DOMAIN` = (留空)

#### Linux/Mac

编辑 `~/.bashrc` 或 `~/.zshrc`，添加：

```bash
export OSS_ENDPOINT="oss-cn-beijing.aliyuncs.com"
export OSS_ACCESS_KEY_ID="你的AccessKey ID"
export OSS_ACCESS_KEY_SECRET="你的AccessKey Secret"
export OSS_BUCKET_NAME="zcc-mjh"
export OSS_DOMAIN=""
```

然后运行 `source ~/.bashrc` 或 `source ~/.zshrc`

### 3. 从Git历史中移除敏感信息

由于敏感信息已经提交到了Git历史，需要清理历史记录：

```bash
# 方法1: 使用 git filter-repo (推荐)
git filter-repo --path config/config.go --invert-paths

# 方法2: 如果已经推送，需要强制推送（谨慎使用）
git push --force
```

或者，如果这是私有仓库，可以直接在GitHub上重置提交：

1. 在GitHub上删除包含敏感信息的提交
2. 或者创建一个新的分支，重新提交代码

### 4. 验证配置

运行应用后，检查日志：

```
OSS初始化成功
```

如果没有设置环境变量或配置不完整，会显示：

```
警告: OSS初始化失败: OSS配置不完整，请检查配置 (将使用本地存储)
```

## 重要提示

1. **永远不要**将敏感信息硬编码在代码中
2. **永远不要**将包含敏感信息的文件提交到Git仓库
3. 如果敏感信息已经泄露，应立即在阿里云控制台中重新生成AccessKey
4. 建议使用RAM子账号，只授予必要的OSS权限

## 下一步

1. 立即在阿里云控制台重新生成AccessKey（因为已泄露）
2. 设置环境变量
3. 从Git历史中移除敏感信息
4. 重新提交代码

