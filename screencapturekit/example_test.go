//go:build darwin

package screencapturekit_test

import (
	"fmt"

	"github.com/tmc/apple/screencapturekit"
)

func ExampleSCStreamConfiguration() {
	config := screencapturekit.NewSCStreamConfiguration()
	config.SetWidth(1920)
	config.SetHeight(1080)
	config.SetShowsCursor(true)

	fmt.Printf("Width: %d\n", config.Width())
	fmt.Printf("Height: %d\n", config.Height())
	fmt.Printf("ShowsCursor: %t\n", config.ShowsCursor())

	// Output:
	// Width: 1920
	// Height: 1080
	// ShowsCursor: true
}

func ExampleSCScreenshotConfiguration() {
	config := screencapturekit.NewSCScreenshotConfiguration()
	config.SetWidth(1280)
	config.SetHeight(720)
	config.SetShowsCursor(false)

	fmt.Printf("Width: %d\n", config.Width())
	fmt.Printf("Height: %d\n", config.Height())
	fmt.Printf("ShowsCursor: %t\n", config.ShowsCursor())

	// Output:
	// Width: 1280
	// Height: 720
	// ShowsCursor: false
}
