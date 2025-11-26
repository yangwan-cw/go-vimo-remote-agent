package capture

import (
	"errors"
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
)

type WindowsCapturer struct {
	config *CaptureConfig

	displayIndex int

	bounds image.Rectangle
}

func NewWindowsCapturer(config *CaptureConfig) (*WindowsCapturer, error) {
	n := screenshot.NumActiveDisplays() //

	if config.DisplayIndex >= n {
		return nil, errors.New("DisplayIndex out of range")
	}

	bounds := screenshot.GetDisplayBounds(config.DisplayIndex)

	fmt.Printf("Windows captrue init \n")
	fmt.Printf("resolution: %dx%d\n", bounds.Dx(), bounds.Dy())

	return &WindowsCapturer{
		config:       config,
		displayIndex: config.DisplayIndex,
		bounds:       bounds,
	}, nil
}

func (wc *WindowsCapturer) CaptureScreen() (image.Image, error) {
	return screenshot.CaptureDisplay(wc.displayIndex)
}

func (wc *WindowsCapturer) GetScreenSize() (width, height int, err error) {
	return wc.bounds.Dx(), wc.bounds.Dy(), nil
}

func (wc *WindowsCapturer) Close() error {
	return nil
}
