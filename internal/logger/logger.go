package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init 初始化日志器
func Init(isDev bool) error {
	var config zap.Config

	if isDev {
		// 开发环境：输出到控制台，格式化更易读
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// 生产环境：输出 JSON 格式
		config = zap.NewProductionConfig()
	}

	// 配置输出路径
	config.OutputPaths = []string{"stdout", "logs/app.log"}
	config.ErrorOutputPaths = []string{"stderr", "logs/error.log"}

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

// Sync 刷新日志缓冲区
func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}

// InitWithFile 初始化日志器并将日志写入指定文件
func InitWithFile(isDev bool, logFile, errorFile string) error {
	var config zap.Config

	if isDev {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	// 确保日志目录存在
	os.MkdirAll("logs", 0755)

	config.OutputPaths = []string{"stdout", logFile}
	config.ErrorOutputPaths = []string{"stderr", errorFile}

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

// Info 记录信息级别日志
func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

// Error 记录错误级别日志
func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

// Debug 记录调试级别日志
func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

// Warn 记录警告级别日志
func Warn(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

// Fatal 记录致命错误日志并退出程序
func Fatal(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}
