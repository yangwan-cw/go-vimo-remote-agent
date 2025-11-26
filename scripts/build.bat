@echo off
setlocal

echo Building go-vimo-agent.exe...

set PROJECT_NAME=go-vimo-agent
set OUTPUT_DIR=..\bin
set MAIN_FILE=..\cmd\agent\main.go

REM Create bin directory if it doesn't exist
if not exist "%OUTPUT_DIR%" (
    mkdir "%OUTPUT_DIR%"
)

REM Build the executable
go build -o "%OUTPUT_DIR%\%PROJECT_NAME%.exe" "%MAIN_FILE%"

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Build successful!
    echo Output: %OUTPUT_DIR%\%PROJECT_NAME%.exe
) else (
    echo.
    echo Build failed!
    exit /b 1
)

endlocal
