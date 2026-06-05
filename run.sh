#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# 检查二进制文件是否存在
if [ ! -f "./blog" ]; then
    echo "二进制文件不存在，正在构建..."
    bash "$SCRIPT_DIR/build.sh"
fi

echo "========================================"
echo "启动博客系统"
echo "========================================"
echo "访问地址: http://localhost:3000"
echo "按 Ctrl+C 停止服务"
echo ""

exec ./blog
