// Package axscripttest provides an rsc.io/script bridge for macOS UI
// automation testing using the axuiautomation package. It enables
// declarative, txtar-based test scripts for accessibility-driven black-box
// testing of macOS applications.
//
// # Commands
//
//   - app-activate: bring the application to the foreground
//   - window-wait <title> [timeout]: wait for a window with the given title
//   - menu-click <item>...: click a menu item by path
//   - click role=R [title=T] [id=I]: find and click an element
//   - type <text>: type text into the focused element
//   - select <popup-title> <value>: select a value from a popup button
//   - exists role=R [title=T] [id=I]: assert an element exists
//   - value-wait <selectors> <contains> [timeout]: wait for an element's value
//   - screenshot <file>: capture the main window to a PNG file
//
// # Conditions
//
//   - [trusted]: true if the process has accessibility permissions.
//     When running inside a macgo bundle, automatically prompts and
//     waits up to 30s for the user to grant access in System Settings.
//   - [screencapture]: true if the process has screen recording
//     permissions. Also prompts and waits when inside a macgo bundle.
//   - [app-running]: true if the target application is running
//
// # Accessibility Permissions
//
// Scripts that begin with [!trusted] skip will silently pass when the test
// binary lacks accessibility permission. To trigger the system permission
// dialog, call EnsureTrusted from TestMain. This uses macgo to create an
// app bundle with a TCC identity so macOS can register and prompt for
// accessibility access:
//
//	var axprompt = flag.Bool("axprompt", false, "prompt for accessibility")
//
//	func TestMain(m *testing.M) {
//	    runtime.LockOSThread()
//	    flag.Parse()
//	    if *axprompt {
//	        if err := axscripttest.EnsureTrusted(); err != nil {
//	            log.Fatal(err)
//	        }
//	    }
//	    os.Exit(m.Run())
//	}
//
// Then run with: go test -run TestFinder -axprompt
//
// # Thread Safety
//
// The macOS Accessibility API must be called from the main OS thread.
// Use runtime.LockOSThread in TestMain before running any scripts.
//
// # Example
//
//	func TestFinder(t *testing.T) {
//	    axscripttest.RunFile(t, axscripttest.Config{App: "com.apple.finder"}, "testdata/finder.txt")
//	}
package axscripttest
