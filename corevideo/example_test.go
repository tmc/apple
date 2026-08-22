//go:build darwin

package corevideo_test

import (
	"fmt"

	"github.com/tmc/apple/corevideo"
)

func ExampleCVPixelBufferCreate() {
	var pixelBuffer corevideo.CVPixelBufferRef
	// 32-bit BGRA format: 'BGRA' = 0x42475241
	const pixelFormatBGRA = 0x42475241

	ret := corevideo.CVPixelBufferCreate(0, 640, 480, pixelFormatBGRA, 0, &pixelBuffer)
	if ret != 0 {
		fmt.Printf("Failed to create pixel buffer: %d\n", ret)
		return
	}
	defer corevideo.CVPixelBufferRelease(pixelBuffer)

	width := corevideo.CVPixelBufferGetWidth(pixelBuffer)
	height := corevideo.CVPixelBufferGetHeight(pixelBuffer)
	bytesPerRow := corevideo.CVPixelBufferGetBytesPerRow(pixelBuffer)
	fmt.Printf("PixelBuffer created: %dx%d, bytesPerRow: %d, return code: %d\n", width, height, bytesPerRow, ret)

	// Output:
	// PixelBuffer created: 640x480, bytesPerRow: 2560, return code: 0
}

func ExampleCVTime() {
	cvTime := corevideo.CVTime{
		TimeValue: 1000,
		TimeScale: 60,
		Flags:     1,
	}

	fmt.Printf("CVTime: Value=%d, Scale=%d, Flags=%d\n",
		cvTime.TimeValue, cvTime.TimeScale, cvTime.Flags)

	// Output:
	// CVTime: Value=1000, Scale=60, Flags=1
}

func ExampleCVTimeStamp() {
	ts := corevideo.CVTimeStamp{
		Version:        0,
		VideoTimeScale: 60,
		VideoTime:      3600,
		HostTime:       123456789,
	}

	fmt.Printf("CVTimeStamp: VideoTime=%d, Scale=%d, HostTime=%d\n",
		ts.VideoTime, ts.VideoTimeScale, ts.HostTime)

	// Output:
	// CVTimeStamp: VideoTime=3600, Scale=60, HostTime=123456789
}
