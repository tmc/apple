// Command screenshot captures a screenshot of a display or window using the
// ScreenCaptureKit framework and saves it as a PNG file.
//
// Usage:
//
//	screenshot                    # capture the main display
//	screenshot -o out.png         # save to file
//	screenshot -window-id 12345   # capture a specific window
//	screenshot -display-id 1      # capture a specific display
//	screenshot -list              # list available windows and displays
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/screencapturekit"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	output := flag.String("o", "screenshot.png", "output file path")
	windowID := flag.Uint("window-id", 0, "capture a specific window by ID")
	displayID := flag.Uint("display-id", 0, "capture a specific display by ID")
	list := flag.Bool("list", false, "list available windows and displays")
	flag.Parse()

	if err := run(*output, uint32(*windowID), uint32(*displayID), *list); err != nil {
		fmt.Fprintf(os.Stderr, "screenshot: %v\n", err)
		os.Exit(1)
	}
}

func run(output string, windowID, displayID uint32, list bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sc := screencapturekit.GetSCShareableContentClass()
	content, err := sc.GetShareableContent(ctx)
	if err != nil {
		return fmt.Errorf("get shareable content failed (screen recording permission may be required)")
	}

	if list {
		return listContent(content)
	}

	var filter screencapturekit.SCContentFilter
	if windowID != 0 {
		window, ok := findWindow(content, windowID)
		if !ok {
			return fmt.Errorf("window %d not found", windowID)
		}
		filter = screencapturekit.NewContentFilterWithDesktopIndependentWindow(window)
	} else {
		display, ok := findDisplay(content, displayID)
		if !ok {
			if displayID != 0 {
				return fmt.Errorf("display %d not found", displayID)
			}
			return fmt.Errorf("no displays available")
		}
		filter = screencapturekit.NewContentFilterWithDisplayExcludingWindows(display, nil)
	}

	config := screencapturekit.NewSCStreamConfiguration()
	mgr := screencapturekit.GetSCScreenshotManagerClass()

	cgImage, err := mgr.CaptureImageWithFilterConfiguration(ctx, filter, config)
	if err != nil {
		return fmt.Errorf("capture failed (screen recording permission may be required)")
	}
	if cgImage == 0 {
		return fmt.Errorf("capture returned nil image")
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cgImage))

	return saveCGImage(cgImage, output)
}

func listContent(content *screencapturekit.SCShareableContent) error {
	fmt.Println("Displays:")
	for _, d := range content.Displays() {
		frame := d.Frame()
		fmt.Printf("  id=%d  %dx%d at (%.0f,%.0f)\n",
			d.DisplayID(), d.Width(), d.Height(), frame.Origin.X, frame.Origin.Y)
	}
	fmt.Println("\nWindows:")
	for _, w := range content.Windows() {
		frame := w.Frame()
		app := w.OwningApplication()
		var appName string
		if ra, ok := app.(screencapturekit.SCRunningApplication); ok && ra.ID != 0 {
			appName = ra.ApplicationName()
		}
		title := w.Title()
		if title == "" {
			title = "(untitled)"
		}
		onScreen := ""
		if w.OnScreen() {
			onScreen = " [on-screen]"
		}
		fmt.Printf("  id=%d  %q  app=%q  %.0fx%.0f%s\n",
			w.WindowID(), title, appName, frame.Size.Width, frame.Size.Height, onScreen)
	}
	return nil
}

func findWindow(content *screencapturekit.SCShareableContent, id uint32) (screencapturekit.SCWindow, bool) {
	for _, w := range content.Windows() {
		if w.WindowID() == id {
			return w, true
		}
	}
	return screencapturekit.SCWindow{}, false
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

func saveCGImage(cgImage coregraphics.CGImageRef, output string) error {
	ciImage := coreimage.NewImageWithCGImage(cgImage)
	if ciImage.ID == 0 {
		return fmt.Errorf("create CIImage from CGImage")
	}

	ciCtx := coreimage.NewCIContext()
	colorSpace := coregraphics.CGColorSpaceCreateDeviceRGB()

	absPath, err := filepath.Abs(output)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}
	url := foundation.NewURLFileURLWithPath(absPath)

	// kCIFormatRGBA8 = 24
	ok, writeErr := ciCtx.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(ciImage, url, 24, colorSpace, nil)
	if writeErr != nil {
		return fmt.Errorf("write png: %w", writeErr)
	}
	if !ok {
		return fmt.Errorf("write png failed")
	}

	fmt.Fprintf(os.Stderr, "screenshot saved to %s\n", absPath)
	return nil
}
