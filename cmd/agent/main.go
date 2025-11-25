package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/yangwan/go-vimo-agent/internal/capture"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Go-Remote Agent 运行成功！")

	config := capture.DefaultConfig()

	var capturer capture.ScreenCapturer
	var err error

	if runtime.GOOS == "windows" {
		fmt.Println("使用Windows捕获器")
		capturer, err = capture.NewWindowsCapturer(config)
	} else {
		fmt.Println(" 当前仅支持Windows，Linux版本开发中")
		return
	}
	if err != nil {
		fmt.Printf(" 初始化失败: %v\n", err)
		return
	}
	width, height, _ := capturer.GetScreenSize()
	fmt.Printf("\n屏幕尺寸: %dx%d\n\n", width, height)
	fmt.Println(" 开始捕获10帧...\n")

	os.MkdirAll("output", 0755)

	for i := 0; i < 10; i++ {
		start := time.Now()
		img, err := capturer.CaptureScreen()
		if err != nil {
			fmt.Printf(" 帧%d失败: %v\n", i+1, err)
			continue
		}

		filename := filepath.Join("output", fmt.Sprintf("frame_%03d.jpg", i+1))
		file, _ := os.Create(filename)
		jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
		file.Close()

		elapsed := time.Since(start)
		fmt.Printf("帧%d - 耗时: %v - %s\n", i+1, elapsed, filename)

		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n 完成！截图保存在 output/ 目录")

}
