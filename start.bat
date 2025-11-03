@echo off
chcp 65001 >nul
echo ========================================
echo 启动博客系统
echo ========================================
echo.

REM 检查前端依赖是否已安装
if not exist "frontend\node_modules" (
    echo [!] 检测到前端依赖未安装
    echo [*] 正在安装前端依赖，请稍候...
    cd frontend
    call npm install
    cd ..
    echo [√] 前端依赖安装完成
    echo.
)

echo [1/2] 启动后端服务...
start "博客后端服务" cmd /k "go run main.go"
timeout /t 3 /nobreak >nul

echo [2/2] 启动前端服务...
start "博客前端服务" cmd /k "cd frontend && npm run dev"

echo.
echo ========================================
echo 服务启动完成！
echo ========================================
echo 后端: http://localhost:8080
echo 前端: http://localhost:3000
echo.
echo 提示: 如果前端启动失败，请先运行:
echo   cd frontend
echo   npm install
echo.
echo 按任意键关闭此窗口...
pause >nul

