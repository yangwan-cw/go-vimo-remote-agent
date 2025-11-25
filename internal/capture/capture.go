package capture

import "image"

// 屏幕捕获器接口
type ScreenCapturer interface {
	CaptureScreen() (image.Image, error)

	GetScreenSize() (width, height int, err error)

	Close() error
}

// 配置结构体: 捕获时的性能
type CaptureConfig struct {
	DisplayIndex int
	Quality      int
	MaxFPS       int
}

// 配置结构其配置构造工程
func DefaultConfig() *CaptureConfig {
	return &CaptureConfig{
		DisplayIndex: 0,
		Quality:      80,
		MaxFPS:       30,
	}
}
