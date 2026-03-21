// Command windowlist enumerates all windows, displays, and running applications
// using the ScreenCaptureKit framework and outputs the results as JSON.
//
// Usage:
//
//	windowlist              # list all windows as JSON
//	windowlist -displays    # also include display info
//	windowlist -onscreen    # only include on-screen windows
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/tmc/apple/screencapturekit"
)

type windowEntry struct {
	WindowID  uint32  `json:"window_id"`
	Title     string  `json:"title"`
	AppName   string  `json:"app_name"`
	BundleID  string  `json:"bundle_id"`
	PID       int32   `json:"pid"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	OnScreen  bool    `json:"on_screen"`
	Active    bool    `json:"active"`
	Layer     int     `json:"layer"`
}

type displayEntry struct {
	DisplayID uint32  `json:"display_id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
}

type output struct {
	Windows  []windowEntry  `json:"windows"`
	Displays []displayEntry `json:"displays,omitempty"`
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	showDisplays := flag.Bool("displays", false, "include display information")
	onScreenOnly := flag.Bool("onscreen", false, "only include on-screen windows")
	flag.Parse()

	if err := run(*showDisplays, *onScreenOnly); err != nil {
		fmt.Fprintf(os.Stderr, "windowlist: %v\n", err)
		os.Exit(1)
	}
}

func run(showDisplays, onScreenOnly bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sc := screencapturekit.GetSCShareableContentClass()
	content, err := sc.GetShareableContent(ctx)
	if err != nil {
		return fmt.Errorf("get shareable content failed (screen recording permission may be required)")
	}

	windows := content.Windows()
	var entries []windowEntry
	for _, w := range windows {
		if onScreenOnly && !w.OnScreen() {
			continue
		}
		frame := w.Frame()
		app := w.OwningApplication()
		var appName, bundleID string
		var pid int32
		if ra, ok := app.(screencapturekit.SCRunningApplication); ok && ra.ID != 0 {
			appName = ra.ApplicationName()
			bundleID = ra.BundleIdentifier()
			pid = ra.ProcessID()
		}
		entries = append(entries, windowEntry{
			WindowID: w.WindowID(),
			Title:    w.Title(),
			AppName:  appName,
			BundleID: bundleID,
			PID:      pid,
			X:        frame.Origin.X,
			Y:        frame.Origin.Y,
			Width:    frame.Size.Width,
			Height:   frame.Size.Height,
			OnScreen: w.OnScreen(),
			Active:   w.Active(),
			Layer:    w.WindowLayer(),
		})
	}

	out := output{Windows: entries}

	if showDisplays {
		displays := content.Displays()
		for _, d := range displays {
			frame := d.Frame()
			out.Displays = append(out.Displays, displayEntry{
				DisplayID: d.DisplayID(),
				X:         frame.Origin.X,
				Y:         frame.Origin.Y,
				Width:     frame.Size.Width,
				Height:    frame.Size.Height,
			})
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

