# 博客系统 API 测试脚本 (PowerShell)
$BASE_URL = "http://localhost:8080/api"

Write-Host "=== 测试用户注册 ===" -ForegroundColor Green
$registerBody = @{
    username = "testuser"
    email = "test@example.com"
    password = "password123"
    nickname = "测试用户"
} | ConvertTo-Json

Invoke-RestMethod -Uri "$BASE_URL/auth/register" -Method Post -Body $registerBody -ContentType "application/json"

Write-Host "`n=== 测试用户登录 ===" -ForegroundColor Green
$loginBody = @{
    username = "testuser"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Uri "$BASE_URL/auth/login" -Method Post -Body $loginBody -ContentType "application/json"
$token = $loginResponse.token
Write-Host "获取到的 Token: $token"

$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

Write-Host "`n=== 测试创建分类 ===" -ForegroundColor Green
$categoryBody = @{
    name = "技术分享"
    description = "技术相关文章"
} | ConvertTo-Json

Invoke-RestMethod -Uri "$BASE_URL/categories" -Method Post -Body $categoryBody -Headers $headers

Write-Host "`n=== 测试创建标签 ===" -ForegroundColor Green
$tagBody = @{
    name = "Go语言"
} | ConvertTo-Json

Invoke-RestMethod -Uri "$BASE_URL/tags" -Method Post -Body $tagBody -Headers $headers

Write-Host "`n=== 测试创建文章 ===" -ForegroundColor Green
$articleBody = @{
    title = "我的第一篇博客"
    content = "这是文章内容"
    summary = "文章摘要"
    category_id = 1
    tag_ids = @(1)
    is_published = $true
} | ConvertTo-Json

Invoke-RestMethod -Uri "$BASE_URL/articles" -Method Post -Body $articleBody -Headers $headers

Write-Host "`n=== 测试获取文章列表 ===" -ForegroundColor Green
Invoke-RestMethod -Uri "$BASE_URL/articles?page=1&page_size=10" -Method Get

Write-Host "`n=== 测试获取分类列表 ===" -ForegroundColor Green
Invoke-RestMethod -Uri "$BASE_URL/categories" -Method Get

Write-Host "`n=== 测试获取标签列表 ===" -ForegroundColor Green
Invoke-RestMethod -Uri "$BASE_URL/tags" -Method Get

Write-Host "`n测试完成！" -ForegroundColor Green


