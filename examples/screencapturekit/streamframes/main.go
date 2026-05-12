// Command streamframes captures screen frames with SCStreamOutput.
//
// Usage:
//
//	streamframes                  # capture five frames from the main display
//	streamframes -duration 5s     # capture for up to five seconds
//	streamframes -max-frames 60   # stop after 60 frames
//	streamframes -display-id 1    # capture a specific display
//	streamframes -list            # list available displays
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/screencapturekit"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	displayID := flag.Uint("display-id", 0, "capture a specific display by ID")
	duration := flag.Duration("duration", 2*time.Second, "maximum capture duration")
	maxFrames := flag.Uint64("max-frames", 5, "maximum screen frames to capture; zero captures until duration expires")
	fps := flag.Int("fps", 30, "requested frames per second")
	list := flag.Bool("list", false, "list available displays")
	flag.Parse()

	if err := run(uint32(*displayID), *duration, *maxFrames, *fps, *list); err != nil {
		fmt.Fprintf(os.Stderr, "streamframes: %v\n", err)
		os.Exit(1)
	}
}

func run(displayID uint32, duration time.Duration, maxFrames uint64, fps int, list bool) error {
	if fps <= 0 {
		return fmt.Errorf("fps must be positive")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sc := screencapturekit.GetSCShareableContentClass()
	content, err := sc.GetShareableContent(ctx)
	if err != nil {
		return fmt.Errorf("get shareable content failed (screen recording permission may be required)")
	}

	if list {
		return listDisplays(content)
	}

	display, ok := findDisplay(content, displayID)
	if !ok {
		if displayID != 0 {
			return fmt.Errorf("display %d not found", displayID)
		}
		return fmt.Errorf("no displays available")
	}

	filter := screencapturekit.NewContentFilterWithDisplayExcludingWindows(display, nil)
	config := screencapturekit.NewSCStreamConfiguration()
	config.SetWidth(uintptr(display.Width()))
	config.SetHeight(uintptr(display.Height()))
	config.SetShowsCursor(true)
	config.SetQueueDepth(3)
	config.SetCapturesAudio(false)
	config.SetMinimumFrameInterval(coremedia.CMTimeMake(1, int32(fps)))

	done := make(chan struct{})
	var once sync.Once
	var frames atomic.Uint64
	finish := func() {
		once.Do(func() {
			close(done)
		})
	}

	output := screencapturekit.NewSCStreamOutput(screencapturekit.SCStreamOutputConfig{
		StreamDidOutputSampleBufferOfType: func(stream screencapturekit.SCStream, sampleBuffer coremedia.CMSampleBufferRef, type_ screencapturekit.SCStreamOutputType) {
			if type_ != screencapturekit.SCStreamOutputTypeScreen {
				return
			}
			n := frames.Add(1)
			if n <= 5 {
				fmt.Fprintf(os.Stderr, "frame %d sampleBuffer=0x%x type=%s\n", n, sampleBuffer, type_)
			}
			if maxFrames != 0 && n >= maxFrames {
				finish()
			}
		},
	})
	delegate := screencapturekit.NewSCStreamDelegate(screencapturekit.SCStreamDelegateConfig{})
	stream := screencapturekit.NewStreamWithFilterConfigurationDelegate(filter, config, delegate)
	queue := dispatch.QueueCreate("github.com.tmc.apple.examples.screencapturekit.streamframes")

	if _, err := stream.AddStreamOutputTypeSampleHandlerQueueError(output, screencapturekit.SCStreamOutputTypeScreen, queue); err != nil {
		return fmt.Errorf("add stream output: %w", err)
	}

	if err := stream.StartCapture(ctx); err != nil {
		return fmt.Errorf("start capture failed (screen recording permission may be required): %w", err)
	}

	timer := time.NewTimer(duration)
	select {
	case <-done:
	case <-timer.C:
	}
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopCancel()
	if err := stream.StopCapture(stopCtx); err != nil {
		return fmt.Errorf("stop capture: %w", err)
	}
	if _, err := stream.RemoveStreamOutputTypeError(output, screencapturekit.SCStreamOutputTypeScreen); err != nil {
		return fmt.Errorf("remove stream output: %w", err)
	}

	runtime.KeepAlive(output)
	runtime.KeepAlive(delegate)
	runtime.KeepAlive(stream)

	fmt.Printf("captured %d screen frames from display %d\n", frames.Load(), display.DisplayID())
	return nil
}

func listDisplays(content *screencapturekit.SCShareableContent) error {
	fmt.Println("Displays:")
	for _, d := range content.Displays() {
		frame := d.Frame()
		fmt.Printf("  id=%d  %dx%d at (%.0f,%.0f)\n",
			d.DisplayID(), d.Width(), d.Height(), frame.Origin.X, frame.Origin.Y)
	}
	return nil
}

func findDisplay(content *screencapturekit.SCShareableContent, id uint32) (screencapturekit.SCDisplay, bool) {
	displays := content.Displays()
	if id == 0 && len(displays) > 0 {
		return displays[0], true
	}
	for _, d := range displays {
		if d.DisplayID() == id {
			return d, true
		}
	}
	return screencapturekit.SCDisplay{}, false
}
