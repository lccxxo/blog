# 登录功能测试脚本
$BASE_URL = "http://localhost:8080/api"

Write-Host "========================================" -ForegroundColor Green
Write-Host "博客系统登录功能测试" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
Write-Host ""

# 测试1: 注册用户
Write-Host "[测试 1/4] 注册用户..." -ForegroundColor Cyan
$registerBody = @{
    username = "debuguser"
    email = "debug@example.com"
    password = "debug123"
    nickname = "调试用户"
} | ConvertTo-Json

try {
    $registerResponse = Invoke-RestMethod -Uri "$BASE_URL/auth/register" -Method Post -Body $registerBody -ContentType "application/json" -ErrorAction Stop
    Write-Host "✓ 注册成功" -ForegroundColor Green
    Write-Host "  用户ID: $($registerResponse.user.id)" -ForegroundColor Gray
    Write-Host "  用户名: $($registerResponse.user.username)" -ForegroundColor Gray
} catch {
    if ($_.Exception.Response.StatusCode -eq 400) {
        Write-Host "! 用户已存在，继续测试..." -ForegroundColor Yellow
    } else {
        Write-Host "✗ 注册失败: $($_.Exception.Message)" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""

# 测试2: 登录获取token
Write-Host "[测试 2/4] 用户登录..." -ForegroundColor Cyan
$loginBody = @{
    username = "debuguser"
    password = "debug123"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-RestMethod -Uri "$BASE_URL/auth/login" -Method Post -Body $loginBody -ContentType "application/json" -ErrorAction Stop
    $token = $loginResponse.token
    
    Write-Host "✓ 登录成功" -ForegroundColor Green
    Write-Host "  Token (前50字符): $($token.Substring(0, [Math]::Min(50, $token.Length)))..." -ForegroundColor Gray
    Write-Host "  用户名: $($loginResponse.user.username)" -ForegroundColor Gray
    Write-Host "  邮箱: $($loginResponse.user.email)" -ForegroundColor Gray
} catch {
    Write-Host "✗ 登录失败: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host ""

# 测试3: 使用token访问受保护的接口
Write-Host "[测试 3/4] 访问受保护的接口 (获取个人信息)..." -ForegroundColor Cyan
$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

try {
    $profileResponse = Invoke-RestMethod -Uri "$BASE_URL/profile" -Headers $headers -ErrorAction Stop
    Write-Host "✓ 成功访问受保护的接口" -ForegroundColor Green
    Write-Host "  用户ID: $($profileResponse.user.id)" -ForegroundColor Gray
    Write-Host "  用户名: $($profileResponse.user.username)" -ForegroundColor Gray
    Write-Host "  邮箱: $($profileResponse.user.email)" -ForegroundColor Gray
} catch {
    Write-Host "✗ 访问失败: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "  可能原因: token无效或已过期" -ForegroundColor Yellow
    exit 1
}

Write-Host ""

# 测试4: 测试错误的token
Write-Host "[测试 4/4] 测试无效token (应该返回401)..." -ForegroundColor Cyan
$badHeaders = @{
    "Authorization" = "Bearer invalidtoken123"
    "Content-Type" = "application/json"
}

try {
    Invoke-RestMethod -Uri "$BASE_URL/profile" -Headers $badHeaders -ErrorAction Stop
    Write-Host "✗ 测试失败: 应该拒绝无效token" -ForegroundColor Red
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Host "✓ 正确拒绝了无效token (401 Unauthorized)" -ForegroundColor Green
    } else {
        Write-Host "? 返回了意外的状态码: $($_.Exception.Response.StatusCode)" -ForegroundColor Yellow
    }
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Green
Write-Host "测试完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
Write-Host ""
Write-Host "如果所有测试都通过，说明后端认证功能正常。" -ForegroundColor Cyan
Write-Host "如果前端仍有问题，请检查:" -ForegroundColor Cyan
Write-Host "  1. 前端依赖是否已安装 (npm install)" -ForegroundColor Gray
Write-Host "  2. 浏览器控制台是否有错误" -ForegroundColor Gray
Write-Host "  3. LocalStorage 中是否保存了 token" -ForegroundColor Gray
Write-Host ""





