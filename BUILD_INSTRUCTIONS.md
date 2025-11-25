# 构建说明

## 问题说明

如果构建的 `go-vimo-agent.exe` 在其他电脑上无法运行，通常是因为：

1. **动态链接依赖**：exe依赖了MinGW的DLL文件（如libgcc_s_seh-1.dll等）
2. **CGO依赖**：项目使用了需要CGO的库（github.com/kbinani/screenshot）

## 解决方案

### 方法1：使用提供的构建脚本（推荐）

直接运行 `build.bat`：

```cmd
build.bat
```

脚本会自动：
- 设置正确的环境变量
- 使用静态链接标志 `-extldflags=-static`
- 生成可在任何Windows电脑运行的exe

### 方法2：手动构建命令

```cmd
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w -extldflags=-static" -o go-vimo-agent.exe ./cmd/agent/
```

### ldflags 说明

- `-s`：去除符号表，减小文件大小
- `-w`：去除DWARF调试信息，减小文件大小
- `-extldflags=-static`：**关键**，静态链接所有C库依赖

### 可选：隐藏控制台窗口

如果需要程序运行时不显示控制台窗口，添加 `-H windowsgui`：

```cmd
go build -ldflags="-s -w -H windowsgui -extldflags=-static" -o go-vimo-agent.exe ./cmd/agent/
```

## 前置要求

### 必须安装C编译器

由于使用了CGO，必须安装以下任一C编译器：

1. **MinGW-w64**（推荐）
   - 下载：https://github.com/niXman/mingw-builds-binaries/releases
   - 安装后添加 `bin` 目录到系统PATH

2. **TDM-GCC**
   - 下载：https://jmeubank.github.io/tdm-gcc/

验证安装：
```cmd
gcc --version
```

### 依赖包

首次构建前运行：
```cmd
go mod download
```

## 验证构建

构建成功后，检查文件大小和依赖：

```cmd
# 查看文件信息
dir go-vimo-agent.exe

# 在当前电脑测试运行
go-vimo-agent.exe
```

## 跨机器部署

生成的 `go-vimo-agent.exe` 应该能在任何 Windows 7+ (64位) 系统上运行，无需额外DLL文件。

如果仍然无法运行，可能需要：
- Visual C++ 运行库：https://aka.ms/vs/17/release/vc_redist.x64.exe

## 故障排除

### 错误：gcc: command not found
安装MinGW-w64并添加到PATH

### 错误：undefined reference
依赖包未正确下载，运行 `go mod download`

### exe可以在本机运行，但在其他电脑报错
- 检查是否使用了 `-extldflags=-static` 标志
- 确认目标电脑是64位Windows系统（如需32位，设置 GOARCH=386）
