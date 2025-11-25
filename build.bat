@echo off
echo 正在构建 go-vimo-remote-agent...
echo.

:: 设置环境变量
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

:: 使用静态链接标志构建（去除-H windowsgui以便看到控制台输出）
go build -ldflags="-s -w -extldflags=-static" -o go-vimo-agent.exe ./cmd/agent/

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ✅ 构建成功！
    echo 📦 可执行文件: go-vimo-agent.exe
    echo 📏 文件大小:
    dir go-vimo-agent.exe | find "go-vimo-agent.exe"
    echo.
    echo 💡 提示：
    echo    - 此exe已静态链接，可在其他Windows电脑上运行
    echo    - 如需隐藏控制台窗口，添加 -H windowsgui 到 ldflags
    echo.
) else (
    echo.
    echo ❌ 构建失败！
    echo.
    echo 可能的原因：
    echo 1. 缺少C编译器（需要 MinGW-w64 或 TDM-GCC）
    echo 2. 依赖包未下载（运行: go mod download）
    echo.
    echo 安装MinGW-w64:
    echo https://github.com/niXman/mingw-builds-binaries/releases
)

pause
