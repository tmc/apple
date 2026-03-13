package axuiautomation

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// Application represents a running macOS application for accessibility
// automation. It holds the root AXUIElementRef for the process and serves
// as the entry point for element queries and observers. Call Close when
// done to release the root element and cached resources.
type Application struct {
	bundleID string
	pid      int32
	root     *Element
	strCache *stringCache
}

// NewApplication creates a new Application by bundle ID.
// It looks up the running process by bundle ID and creates an accessibility element for it.
func NewApplication(bundleID string) (*Application, error) {
	// Find PID by bundle ID using pgrep
	pid, err := findPIDByBundleID(bundleID)
	if err != nil {
		return nil, ErrNotRunning
	}

	app := NewApplicationFromPID(pid)
	if app != nil {
		app.bundleID = bundleID
	}
	return app, nil
}

// NewApplicationFromPID creates a new Application from a process ID.
func NewApplicationFromPID(pid int32) *Application {
	ref := AXUIElementCreateApplication(pid)
	if ref == 0 {
		return nil
	}

	app := &Application{
		pid:      pid,
		strCache: newStringCache(),
	}
	app.root = newElement(ref, app)

	return app
}

// findPIDByBundleID finds the PID of a running application by its bundle ID or name.
// It uses "lsappinfo list" which enumerates all running GUI apps. It supports:
// 1. Exact numeric PID
// 2. Exact Bundle ID (e.g. "com.apple.Safari")
// 3. Exact localized name (e.g. "Safari")
// 4. Case-insensitive substring match of the name (e.g. "safari" or "antigrav")
func findPIDByBundleID(query string) (int32, error) {
	// 1. Exact numeric PID
	if pid, err := strconv.ParseInt(query, 10, 32); err == nil {
		return int32(pid), nil
	}

	// Dump all applications via lsappinfo list
	cmd := exec.Command("lsappinfo", "list")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("lsappinfo failed: %w", err)
	}

	type appInfo struct {
		Name     string
		BundleID string
		PID      int32
	}
	var apps []appInfo
	var cur appInfo

	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, ") \"") && strings.Contains(line, "ASN:") {
			if cur.BundleID != "" || cur.PID != 0 {
				apps = append(apps, cur)
			}
			cur = appInfo{}
			s := line[strings.Index(line, "\"")+1:]
			cur.Name = s[:strings.Index(s, "\"")]
		} else if strings.HasPrefix(line, "bundleID=") {
			id := strings.Trim(strings.TrimPrefix(line, "bundleID="), `"`)
			if id != "[ NULL ]" {
				cur.BundleID = id
			}
		} else if strings.HasPrefix(line, "pid = ") {
			rest := strings.TrimPrefix(line, "pid = ")
			if i := strings.IndexAny(rest, " \t"); i > 0 {
				rest = rest[:i]
			}
			if p, err := strconv.ParseInt(rest, 10, 32); err == nil {
				cur.PID = int32(p)
			}
		}
	}
	if cur.BundleID != "" || cur.PID != 0 {
		apps = append(apps, cur)
	}

	// Priority 1: Exact Bundle ID match
	for _, app := range apps {
		if strings.EqualFold(app.BundleID, query) {
			return app.PID, nil
		}
	}

	// Priority 2: Exact Name match
	for _, app := range apps {
		if strings.EqualFold(app.Name, query) {
			return app.PID, nil
		}
	}

	// Priority 3: Substring Name match
	queryLower := strings.ToLower(query)
	for _, app := range apps {
		if strings.Contains(strings.ToLower(app.Name), queryLower) {
			return app.PID, nil
		}
	}

	// Priority 4: Substring Bundle ID match
	for _, app := range apps {
		if strings.Contains(strings.ToLower(app.BundleID), queryLower) {
			return app.PID, nil
		}
	}

	return 0, ErrNotRunning
}

// PID returns the application's process ID.
func (a *Application) PID() int32 {
	return a.pid
}

// BundleID returns the application's bundle ID.
func (a *Application) BundleID() string {
	return a.bundleID
}

// Root returns the root accessibility element for the application.
func (a *Application) Root() *Element {
	return a.root
}

// Close releases all resources associated with the application.
func (a *Application) Close() {
	if a.root != nil {
		a.root.Release()
		a.root = nil
	}
	if a.strCache != nil {
		a.strCache.release()
	}
}

// IsRunning returns true if the application is still running.
func (a *Application) IsRunning() bool {
	if a.root == nil {
		return false
	}
	// Try to get a basic attribute to verify the app is accessible
	return a.root.Exists()
}

// Activate brings the application to the foreground.
func (a *Application) Activate() error {
	// Use AppleScript to activate the application
	script := `tell application "System Events" to set frontmost of (first process whose unix id is ` +
		strconv.Itoa(int(a.pid)) + `) to true`

	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

// Windows returns a query for all windows of the application.
func (a *Application) Windows() *ElementQuery {
	if a.root == nil {
		return newElementQuery(nil, a)
	}
	return newElementQuery(a.root, a).ByRole("AXWindow")
}

// WindowList returns all windows using the AXWindows attribute directly,
// which is faster and more reliable than traversing AXChildren.
func (a *Application) WindowList() []*Element {
	if a.root == nil || a.root.ref == 0 {
		return nil
	}
	refs := getAXAttributeElements(a.root.ref, "AXWindows")
	els := make([]*Element, 0, len(refs))
	for _, ref := range refs {
		if ref != 0 {
			els = append(els, newElement(ref, a))
		}
	}
	return els
}

// MainWindow returns the application's main window (first window).
func (a *Application) MainWindow() *Element {
	return a.Windows().First()
}

// WindowByTitle returns the first window with the given title.
func (a *Application) WindowByTitle(title string) *Element {
	return a.Windows().ByTitle(title).First()
}

// WindowByTitleContains returns the first window whose title contains the given substring.
func (a *Application) WindowByTitleContains(substr string) *Element {
	return a.Windows().ByTitleContains(substr).First()
}

// FocusedElement returns the currently focused element.
func (a *Application) FocusedElement() *Element {
	if a.root == nil || a.root.ref == 0 {
		return nil
	}

	var value uintptr
	err := AXUIElementCopyAttributeValue(a.root.ref, axAttr("AXFocusedUIElement"), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return nil
	}

	return newElement(AXUIElementRef(value), a)
}

// MenuBar returns the application's menu bar element.
// It spins the CFRunLoop briefly on each attempt so that AX IPC replies are
// delivered even in CLI processes that do not run a persistent event loop.
func (a *Application) MenuBar() *Element {
	if a.root == nil || a.root.ref == 0 {
		return nil
	}

	for range 8 {
		var value uintptr
		err := AXUIElementCopyAttributeValue(a.root.ref, axAttr("AXMenuBar"), &value)
		if int(err) == axErrorSuccess && value != 0 {
			return newElement(AXUIElementRef(value), a)
		}
		SpinRunLoop(100 * time.Millisecond)
	}
	return nil
}

// Dialogs returns a query for all dialog windows.
func (a *Application) Dialogs() *ElementQuery {
	if a.root == nil {
		return newElementQuery(nil, a)
	}
	return newElementQuery(a.root, a).ByRole("AXDialog")
}

// Sheets returns a query for all sheet windows.
func (a *Application) Sheets() *ElementQuery {
	if a.root == nil {
		return newElementQuery(nil, a)
	}
	return newElementQuery(a.root, a).ByRole("AXSheet")
}

// NewObserver creates a new observer for this application.
func (a *Application) NewObserver() (*Observer, error) {
	return NewObserver(a)
}

// Descendants returns a query for all descendants of the application root.
func (a *Application) Descendants() *ElementQuery {
	if a.root == nil {
		return newElementQuery(nil, a)
	}
	return newElementQuery(a.root, a)
}

// Buttons returns a query for all buttons in the application.
func (a *Application) Buttons() *ElementQuery {
	return a.Descendants().ByRole("AXButton")
}

// TextFields returns a query for all text fields in the application.
func (a *Application) TextFields() *ElementQuery {
	return a.Descendants().ByRole("AXTextField")
}

// Checkboxes returns a query for all checkboxes in the application.
func (a *Application) Checkboxes() *ElementQuery {
	return a.Descendants().ByRole("AXCheckBox")
}

// IsProcessTrusted checks if the current process has accessibility permissions.
func IsProcessTrusted() bool {
	return AXIsProcessTrusted()
}

// PromptForAccessibility triggers the system accessibility permission prompt.
// Returns true if already trusted, false otherwise.
func PromptForAccessibility() bool {
	return isTrustedFresh(true)
}

// CheckAccessibilityAccess performs a diagnostic check to see if accessibility API is working.
// Returns the AX error code (0 = success, -25211 = API disabled/no permission).
func CheckAccessibilityAccess(pid int32) (int, string) {
	ref := AXUIElementCreateApplication(pid)
	if ref == 0 {
		return -1, "failed to create AXUIElement"
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(ref)))

	// Try to get the AXChildren attribute - this will fail if we don't have permission
	var value uintptr
	err := AXUIElementCopyAttributeValue(ref, axAttr("AXChildren"), &value)
	// Convert to signed int32 to get proper error codes
	code := int(int32(err))
	if code == 0 && value != 0 {
		corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))
		return 0, "OK"
	}

	switch code {
	case -25211:
		return code, "API disabled (no accessibility permission)"
	case -25202:
		return code, "invalid UI element"
	case -25204:
		return code, "cannot complete"
	case -25205:
		return code, "attribute unsupported"
	case -25212:
		return code, "no value"
	default:
		return code, fmt.Sprintf("error code %d", code)
	}
}

// WaitForWindow waits up to timeout for a window whose title contains the given substring to appear.
func (a *Application) WaitForWindow(title string, timeout time.Duration) (*Element, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if w := a.WindowByTitleContains(title); w != nil {
			return w, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil, ErrTimeout
}

// ClickMenuItem clicks a menu item by its path (e.g., ["File", "Export..."]).
func (a *Application) ClickMenuItem(path []string) error {
	if len(path) == 0 {
		return ErrElementNotFound
	}

	menuBar := a.MenuBar()
	if menuBar == nil {
		return &Error{Message: "menu bar not found"}
	}

	// Navigate through the menu hierarchy
	current := menuBar
	for i, itemName := range path {
		// Find the menu item
		var menuItem *Element
		children := current.Children()
		for _, child := range children {
			role := child.Role()
			title := child.Title()
			if (role == "AXMenuBarItem" || role == "AXMenuItem" || role == "AXMenu") && title == itemName {
				menuItem = child
				break
			}
		}

		if menuItem == nil {
			return fmt.Errorf("menu item %q not found", itemName)
		}

		// Click the menu item to open it
		if err := menuItem.Click(); err != nil {
			return fmt.Errorf("failed to click menu item %q: %w", itemName, err)
		}

		// For non-leaf items, wait for menu to expand and find the submenu
		if i < len(path)-1 {
			time.Sleep(200 * time.Millisecond)

			// Find the submenu that opened
			children = menuItem.Children()
			for _, child := range children {
				if child.Role() == "AXMenu" {
					current = child
					break
				}
			}
		}
	}

	return nil
}
