#!/bin/bash
set -e

echo "========================================"
echo "构建博客系统"
echo "========================================"
echo ""

# 获取脚本所在目录（项目根目录）
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

echo "[1/3] 构建前端..."
cd frontend
npm install --silent
npm run build
cd "$SCRIPT_DIR"
echo "  ✓ 前端构建完成 (frontend/dist/)"

echo "[2/3] 编译后端..."
go build -ldflags="-s -w" -o blog .
echo "  ✓ 后端编译完成 (blog)"

echo "[3/3] 复制前端产物到 dist/..."
# systemd 服务从 /opt/blog 运行时，前端文件在 frontend/dist/ 下

echo ""
echo "========================================"
echo "构建完成！"
echo "========================================"
echo "二进制文件: $SCRIPT_DIR/blog"
echo "前端产物:   $SCRIPT_DIR/frontend/dist/"
echo ""
echo "运行方式:"
echo "  直接运行:  ./run.sh"
echo "  安装服务:  sudo ./install.sh"
