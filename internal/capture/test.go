package capture

import "fmt"

func Test123() error {
	fmt.Printf("capture test")
	return nil
}

type TestConfig struct {
	DisplayIndex int
	Quality      int
	MaxFPS       int
}
