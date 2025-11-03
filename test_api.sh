#!/bin/bash

# 博客系统 API 测试脚本
BASE_URL="http://localhost:8080/api"

echo "=== 测试用户注册 ==="
curl -X POST $BASE_URL/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123","nickname":"测试用户"}'

echo -e "\n\n=== 测试用户登录 ==="
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}')

echo $LOGIN_RESPONSE

# 提取 token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)
echo -e "\n获取到的 Token: $TOKEN"

echo -e "\n\n=== 测试创建分类 ==="
curl -X POST $BASE_URL/categories \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"技术分享","description":"技术相关文章"}'

echo -e "\n\n=== 测试创建标签 ==="
curl -X POST $BASE_URL/tags \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Go语言"}'

echo -e "\n\n=== 测试创建文章 ==="
curl -X POST $BASE_URL/articles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"我的第一篇博客","content":"这是文章内容","summary":"文章摘要","category_id":1,"tag_ids":[1],"is_published":true}'

echo -e "\n\n=== 测试获取文章列表 ==="
curl -X GET "$BASE_URL/articles?page=1&page_size=10"

echo -e "\n\n=== 测试获取分类列表 ==="
curl -X GET $BASE_URL/categories

echo -e "\n\n=== 测试获取标签列表 ==="
curl -X GET $BASE_URL/tags

echo -e "\n\n测试完成！"


