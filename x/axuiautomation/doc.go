// Package axuiautomation provides XCUIAutomation-style interfaces for macOS
// accessibility automation.
//
// This package wraps the macOS Accessibility APIs (AXUIElement, AXObserver, etc.)
// with a clean, Go-idiomatic interface inspired by Apple's XCUIAutomation framework.
//
// # Core Types
//
//   - Application: represents a running application
//   - Element: represents a UI element (button, window, etc.)
//   - ElementQuery: fluent API for finding elements
//   - Observer: event-based waiting for UI state changes
//
// # Example Usage
//
//	app, err := axuiautomation.NewApplication("com.apple.dt.Xcode")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer app.Close()
//
//	// Find and click a button
//	replayBtn := app.Windows().Buttons().ByTitle("Replay").Element(0)
//	if replayBtn.Exists() && replayBtn.IsEnabled() {
//	    replayBtn.Click()
//	}
//
//	// Wait for element with observer
//	observer, _ := app.NewObserver()
//	defer observer.Close()
//	err = observer.WaitForEnabled(replayBtn, 5*time.Minute)
//
// # Thread Safety
//
// The macOS Accessibility API must be called from the main thread. Callers
// should use runtime.LockOSThread in their main goroutine before calling
// any functions in this package. Observer-based waiting methods spin the
// CFRunLoop internally, which also requires the main thread.
//
// # Memory Ownership
//
// Elements returned from query methods (First, AllElements, Element) are
// caller-owned and must be released by calling Release when no longer needed.
// Application.Close releases the root element; callers must still release
// any elements obtained from queries. Failing to release elements leaks
// the underlying AXUIElementRef.
//
// # Accessibility Permissions
//
// This package requires accessibility permissions. Use IsProcessTrusted() to
// check if your app has the required permissions.
package axuiautomation
