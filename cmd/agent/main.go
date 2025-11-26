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
	fmt.Println("Go-Remote Agent run scuccess！")

	config := capture.DefaultConfig()

	var capturer capture.ScreenCapturer
	var err error

	if runtime.GOOS == "windows" {
		fmt.Println("used Windows capture")
		capturer, err = capture.NewWindowsCapturer(config)
	} else {
		fmt.Println(" Currently, only Windows is supported. The Linux version is under development")
		return
	}
	if err != nil {
		fmt.Printf(" init error: %v\n", err)
		return
	}
	width, height, _ := capturer.GetScreenSize()
	fmt.Printf("\nscreen size: %dx%d\n\n", width, height)
	fmt.Println(" start cpatrue 10 frame...\n")

	os.MkdirAll("output", 0755)

	for i := 0; i < 10; i++ {
		start := time.Now()
		img, err := capturer.CaptureScreen()
		if err != nil {
			fmt.Printf(" frame%derror: %v\n", i+1, err)
			continue
		}

		filename := filepath.Join("output", fmt.Sprintf("frame_%03d.jpg", i+1))
		file, _ := os.Create(filename)
		jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
		file.Close()

		elapsed := time.Since(start)
		fmt.Printf("frame%d - time: %v - %s\n", i+1, elapsed, filename)

		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n done！frame saved to output/ dir")

}
