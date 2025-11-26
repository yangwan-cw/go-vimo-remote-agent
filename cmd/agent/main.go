package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/yangwan/go-vimo-agent/internal/capture"

	"github.com/yangwan/go-vimo-agent/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志器（开发模式）
	if err := logger.Init(true); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	logger.Info("Go-Remote Agent started successfully")

	config := capture.DefaultConfig()

	var capturer capture.ScreenCapturer
	var err error

	if runtime.GOOS == "windows" {
		logger.Info("Using Windows capture")
		capturer, err = capture.NewWindowsCapturer(config)
	} else {
		logger.Warn("Currently, only Windows is supported. The Linux version is under development")
		return
	}
	if err != nil {
		logger.Error("Failed to initialize capturer", zap.Error(err))
		return
	}
	width, height, _ := capturer.GetScreenSize()
	logger.Info("Screen size detected", zap.Int("width", width), zap.Int("height", height))
	logger.Info("Starting to capture 10 frames")

	os.MkdirAll("output", 0755)

	for i := 0; i < 10; i++ {
		start := time.Now()
		img, err := capturer.CaptureScreen()
		if err != nil {
			logger.Error("Failed to capture frame", zap.Int("frame", i+1), zap.Error(err))
			continue
		}

		filename := filepath.Join("output", fmt.Sprintf("frame_%03d.jpg", i+1))
		file, _ := os.Create(filename)
		jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
		file.Close()

		elapsed := time.Since(start)
		logger.Info("Frame captured",
			zap.Int("frame", i+1),
			zap.Duration("time", elapsed),
			zap.String("filename", filename))

		time.Sleep(1 * time.Second)
	}
	logger.Info("All frames captured successfully", zap.String("output_dir", "output/"))

}
