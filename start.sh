#!/bin/bash

echo "========================================"
echo "启动博客系统 (开发模式)"
echo "========================================"
echo ""

echo "[1/2] 启动后端服务..."
go run main.go &
BACKEND_PID=$!
sleep 2

echo "[2/2] 启动前端开发服务器..."
cd frontend
npm run dev &
FRONTEND_PID=$!

echo ""
echo "========================================"
echo "服务启动完成！"
echo "========================================"
echo "后端API: http://localhost:3000"
echo "前端开发: http://localhost:5173"
echo ""
echo "按 Ctrl+C 停止所有服务"

# 等待用户中断
trap "kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit" INT
wait
