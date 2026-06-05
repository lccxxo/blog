#!/bin/bash
set -e

echo "========================================"
echo "安装博客系统为系统服务"
echo "========================================"
echo ""

# 检查 root 权限
if [ "$(id -u)" -ne 0 ]; then
    echo "错误: 请使用 sudo 运行此脚本"
    echo "用法: sudo ./install.sh"
    exit 1
fi

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
APP_DIR="/opt/blog"

echo "[1/5] 构建项目..."
bash "$SCRIPT_DIR/build.sh"

echo ""
echo "[2/5] 安装到 $APP_DIR ..."
mkdir -p "$APP_DIR"
# 复制二进制文件
cp "$SCRIPT_DIR/blog" "$APP_DIR/blog"
chmod +x "$APP_DIR/blog"
# 复制前端构建产物
if [ -d "$SCRIPT_DIR/frontend/dist" ]; then
    rm -rf "$APP_DIR/frontend/dist" 2>/dev/null || true
    mkdir -p "$APP_DIR/frontend"
    cp -r "$SCRIPT_DIR/frontend/dist" "$APP_DIR/frontend/dist"
fi
# 复制上传目录（如果存在）
if [ -d "$SCRIPT_DIR/uploads" ]; then
    cp -r "$SCRIPT_DIR/uploads" "$APP_DIR/uploads"
fi
echo "  ✓ 文件已复制到 $APP_DIR"

echo ""
echo "[3/5] 配置 systemd 服务..."
cp "$SCRIPT_DIR/blog.service" /etc/systemd/system/blog.service
systemctl daemon-reload
echo "  ✓ 服务文件已安装"

echo ""
echo "[4/5] 设置开机自启..."
systemctl enable blog
echo "  ✓ 已设置开机自启"

echo ""
echo "[5/5] 启动服务..."
systemctl start blog
sleep 1

# 检查服务状态
if systemctl is-active --quiet blog; then
    echo "  ✓ 服务已成功启动!"
else
    echo "  ✗ 服务启动失败，查看日志: journalctl -u blog -n 20"
    exit 1
fi

echo ""
echo "========================================"
echo "安装完成！"
echo "========================================"
echo ""
echo "服务已运行在: http://localhost:3000"
echo ""
echo "常用命令:"
echo "  查看状态:    systemctl status blog"
echo "  停止服务:    systemctl stop blog"
echo "  启动服务:    systemctl start blog"
echo "  重启服务:    systemctl restart blog"
echo "  查看日志:    journalctl -u blog -f"
echo "  禁用自启:    systemctl disable blog"
echo "  卸载:       sudo rm -rf /opt/blog && sudo rm /etc/systemd/system/blog.service && sudo systemctl daemon-reload"
