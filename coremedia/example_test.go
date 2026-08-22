//go:build darwin

package coremedia_test

import (
	"fmt"

	"github.com/tmc/apple/coremedia"
)

func ExampleCMTime() {
	time := coremedia.CMTimeMake(300, 60)
	seconds := coremedia.CMTimeGetSeconds(time)
	fmt.Printf("Value: %d, Timescale: %d\n", time.Value(), time.Timescale())
	fmt.Printf("Seconds: %.1f\n", seconds)

	// Output:
	// Value: 300, Timescale: 60
	// Seconds: 5.0
}

func ExampleCMTimeRange() {
	start := coremedia.CMTimeMake(0, 60)
	duration := coremedia.CMTimeMake(120, 60)
	timeRange := coremedia.CMTimeRangeMake(start, duration)

	fmt.Printf("Start: %.1fs, Duration: %.1fs\n",
		coremedia.CMTimeGetSeconds(timeRange.Start()),
		coremedia.CMTimeGetSeconds(timeRange.Duration()),
	)

	// Output:
	// Start: 0.0s, Duration: 2.0s
}

func ExampleCMVideoDimensions() {
	dims := coremedia.CMVideoDimensions{
		Width:  1920,
		Height: 1080,
	}

	fmt.Printf("Video Dimensions: %dx%d\n", dims.Width, dims.Height)

	// Output:
	// Video Dimensions: 1920x1080
}
