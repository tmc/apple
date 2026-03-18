package axuiautomation

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// Element wraps an AXUIElementRef with Go ownership semantics.
// Elements obtained from query methods (First, AllElements, Element) are
// caller-owned and must be released by calling Release. The zero value
// is not usable; use NewApplication or query methods to obtain elements.
type Element struct {
	ref AXUIElementRef
	app *Application
}

// newElement creates a new Element wrapping an AXUIElementRef.
// The element takes ownership of the ref and will release it when no longer needed.
func newElement(ref AXUIElementRef, app *Application) *Element {
	if ref == 0 {
		return nil
	}
	return &Element{ref: ref, app: app}
}

// newElementRetained creates a new Element, retaining the ref.
// Use this when the ref is owned by something else (e.g., an array).
func newElementRetained(ref AXUIElementRef, app *Application) *Element {
	if ref == 0 {
		return nil
	}
	corefoundation.CFRetain(corefoundation.CFTypeRef(ref))
	return &Element{ref: ref, app: app}
}

// Ref returns the underlying AXUIElementRef.
// The caller should not release this ref; it is owned by the Element.
func (e *Element) Ref() AXUIElementRef {
	if e == nil {
		return 0
	}
	return e.ref
}

// Release releases the underlying AXUIElementRef.
// Callers must call Release on elements returned from query methods
// (First, AllElements, Element, Parent, Children, FocusedElement, etc.)
// when the element is no longer needed. After calling Release, the
// Element must not be used.
func (e *Element) Release() {
	if e != nil && e.ref != 0 {
		corefoundation.CFRelease(corefoundation.CFTypeRef(e.ref))
		e.ref = 0
	}
}

// Exists returns true if the element reference is valid.
func (e *Element) Exists() bool {
	if e == nil || e.ref == 0 {
		return false
	}
	// Check if we can get the role attribute
	role := e.Role()
	return role != ""
}

// Role returns the element's accessibility role (e.g., "AXButton", "AXWindow").
func (e *Element) Role() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXRole")
}

// Subrole returns the element's accessibility subrole.
func (e *Element) Subrole() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXSubrole")
}

// Title returns the element's title.
func (e *Element) Title() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXTitle")
}

// Description returns the element's description.
func (e *Element) Description() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXDescription")
}

// RoleDescription returns the element's localized role description.
func (e *Element) RoleDescription() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXRoleDescription")
}

// Identifier returns the element's unique identifier.
func (e *Element) Identifier() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXIdentifier")
}

// Value returns the element's value as a string.
func (e *Element) Value() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXValue")
}

// IsEnabled returns true if the element is enabled/interactive.
func (e *Element) IsEnabled() bool {
	if e == nil || e.ref == 0 {
		return false
	}
	return getAXAttributeBool(e.ref, "AXEnabled")
}

// IsFocused returns true if the element is focused.
func (e *Element) IsFocused() bool {
	if e == nil || e.ref == 0 {
		return false
	}
	return getAXAttributeBool(e.ref, "AXFocused")
}

// IsSelected returns true if the element is selected.
func (e *Element) IsSelected() bool {
	if e == nil || e.ref == 0 {
		return false
	}
	return getAXAttributeBool(e.ref, "AXSelected")
}

// IsChecked returns true if the element's value is truthy (for checkboxes/switches).
func (e *Element) IsChecked() bool {
	if e == nil || e.ref == 0 {
		return false
	}
	return getAXAttributeBool(e.ref, "AXValue")
}

// Position returns the element's position on screen.
func (e *Element) Position() (x, y int) {
	if e == nil || e.ref == 0 {
		return 0, 0
	}
	point, ok := getAXAttributePoint(e.ref, "AXPosition")
	if !ok {
		return 0, 0
	}
	return int(point.X), int(point.Y)
}

// Size returns the element's size.
func (e *Element) Size() (width, height int) {
	if e == nil || e.ref == 0 {
		return 0, 0
	}
	size, ok := getAXAttributeSize(e.ref, "AXSize")
	if !ok {
		return 0, 0
	}
	return int(size.Width), int(size.Height)
}

// Frame returns the element's frame (position + size).
func (e *Element) Frame() Rect {
	if e == nil || e.ref == 0 {
		return Rect{}
	}
	point, _ := getAXAttributePoint(e.ref, "AXPosition")
	size, _ := getAXAttributeSize(e.ref, "AXSize")
	return Rect{
		Origin: point,
		Size:   size,
	}
}

// Center returns the center point of the element.
func (e *Element) Center() (x, y int) {
	frame := e.Frame()
	return int(frame.Origin.X + frame.Size.Width/2),
		int(frame.Origin.Y + frame.Size.Height/2)
}

// Click performs a click on the element.
// It first tries AXPress, then falls back to CGEvent if AXPress fails.
func (e *Element) Click() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	// Try AXPress first
	err := AXUIElementPerformAction(e.ref, axAttr("AXPress"))
	if int(err) == axErrorSuccess {
		return nil
	}

	// If AXPress failed with action unsupported or cannot complete, use CGEvent fallback
	if int(err) == axErrorActionUnsupported || int(err) == axErrorCannotComplete {
		x, y := e.Center()
		if x == 0 && y == 0 {
			return ErrElementNotFound
		}
		return cgEventClick(x, y)
	}

	return axErrorToGo(err)
}

// ClickAt performs a low-level CGEvent click at a specific offset relative to the element's top-left origin.
func (e *Element) ClickAt(xOffset, yOffset int) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}
	frame := e.Frame()
	// If the frame is entirely zero, the element may not be on screen or visible
	if frame.Size.Width == 0 && frame.Size.Height == 0 {
		return ErrElementNotFound
	}
	absX := int(frame.Origin.X) + xOffset
	absY := int(frame.Origin.Y) + yOffset
	return cgEventClick(absX, absY)
}

// DoubleClick performs a double-click on the element.
func (e *Element) DoubleClick() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	x, y := e.Center()
	if x == 0 && y == 0 {
		return ErrElementNotFound
	}
	return cgEventDoubleClick(x, y)
}

// Hover moves the mouse cursor to the center of the element without clicking.
func (e *Element) Hover() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	x, y := e.Center()
	if x == 0 && y == 0 {
		return ErrElementNotFound
	}
	// Use native CGWarpMouseCursorPosition
	initCGEvents()
	if !cgEventsInitialized {
		return errors.New("CGEvents not initialized")
	}
	cgWarpMouseCursorPosition(float64(x), float64(y))
	return nil
}

// PerformAction performs a named accessibility action on the element.
func (e *Element) PerformAction(action string) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}
	err := AXUIElementPerformAction(e.ref, axAttr(action))
	return axErrorToGo(err)
}

// SetValue sets the element's value.
func (e *Element) SetValue(value string) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	cfValue := corefoundation.CFStringCreateWithCString(0, value, cfStringEncodingUTF8)
	if cfValue == 0 {
		return &Error{Message: "failed to create CFString"}
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cfValue))

	err := AXUIElementSetAttributeValue(e.ref, axAttr("AXValue"), uintptr(cfValue))
	return axErrorToGo(err)
}

// Focus sets focus to this element.
func (e *Element) Focus() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	initCFBoolean()
	// Set AXFocused to true
	// Note: We need to get kCFBooleanTrue
	err := AXUIElementSetAttributeValue(e.ref, axAttr("AXFocused"), getCFBooleanTrue())
	return axErrorToGo(err)
}

// SetPosition moves the element to the specified coordinates.
func (e *Element) SetPosition(x, y float64) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	point := Point{X: x, Y: y}
	axValue := AXValueCreate(AXValueType(axValueTypeCGPoint), unsafe.Pointer(&point))
	if axValue == 0 {
		return &Error{Message: "failed to create AXValue for position"}
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(axValue)))

	err := AXUIElementSetAttributeValue(e.ref, axAttr("AXPosition"), axValue)
	return axErrorToGo(err)
}

var cfBooleanTrue uintptr
var cfBooleanFalse uintptr
var cfBooleanTrueOnce sync.Once

func getCFBooleanTrue() uintptr {
	cfBooleanTrueOnce.Do(func() {
		libCF, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}
		if sym, err := purego.Dlsym(libCF, "kCFBooleanTrue"); err == nil {
			cfBooleanTrue = *(*uintptr)(unsafe.Pointer(sym))
		}
		if sym, err := purego.Dlsym(libCF, "kCFBooleanFalse"); err == nil {
			cfBooleanFalse = *(*uintptr)(unsafe.Pointer(sym))
		}
	})
	return cfBooleanTrue
}

// Parent returns the parent element.
func (e *Element) Parent() *Element {
	if e == nil || e.ref == 0 {
		return nil
	}

	var value uintptr
	err := AXUIElementCopyAttributeValue(e.ref, axAttr("AXParent"), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return nil
	}

	return newElement(AXUIElementRef(value), e.app)
}

// Children returns the element's children.
func (e *Element) Children() []*Element {
	if e == nil || e.ref == 0 {
		return nil
	}

	refs := getAXAttributeElements(e.ref, "AXChildren")
	if len(refs) == 0 {
		return nil
	}

	result := make([]*Element, len(refs))
	for i, ref := range refs {
		result[i] = newElement(ref, e.app)
	}
	return result
}

// ChildCount returns the number of children.
func (e *Element) ChildCount() int {
	if e == nil || e.ref == 0 {
		return 0
	}

	var count int
	err := AXUIElementGetAttributeValueCount(e.ref, axAttr("AXChildren"), &count)
	if int(err) != axErrorSuccess {
		return 0
	}
	return count
}

// Query methods - start a new query rooted at this element

// Descendants returns a query for all descendants of this element.
func (e *Element) Descendants() *ElementQuery {
	return newElementQuery(e, e.app)
}

// Buttons returns a query for button elements.
func (e *Element) Buttons() *ElementQuery {
	return e.Descendants().ByRole("AXButton")
}

// Windows returns a query for window elements.
func (e *Element) Windows() *ElementQuery {
	return e.Descendants().ByRole("AXWindow")
}

// TextFields returns a query for text field elements.
func (e *Element) TextFields() *ElementQuery {
	return e.Descendants().ByRole("AXTextField")
}

// Checkboxes returns a query for checkbox elements.
func (e *Element) Checkboxes() *ElementQuery {
	return e.Descendants().ByRole("AXCheckBox")
}

// Groups returns a query for group elements.
func (e *Element) Groups() *ElementQuery {
	return e.Descendants().ByRole("AXGroup")
}

// StaticTexts returns a query for static text elements.
func (e *Element) StaticTexts() *ElementQuery {
	return e.Descendants().ByRole("AXStaticText")
}

// PopUpButtons returns a query for popup button elements.
func (e *Element) PopUpButtons() *ElementQuery {
	return e.Descendants().ByRole("AXPopUpButton")
}

// MenuItems returns a query for menu item elements.
func (e *Element) MenuItems() *ElementQuery {
	return e.Descendants().ByRole("AXMenuItem")
}

// Document returns the element's document URL (typically for windows).
func (e *Element) Document() string {
	if e == nil || e.ref == 0 {
		return ""
	}
	return getAXAttributeString(e.ref, "AXDocument")
}

// Raise raises the window to the front.
func (e *Element) Raise() error {
	return e.PerformAction("AXRaise")
}

// Application returns the parent Application for this element.
func (e *Element) Application() *Application {
	if e == nil {
		return nil
	}
	return e.app
}

// IntValue returns the element's value as an integer.
// Returns 0 if the value is not a number.
func (e *Element) IntValue() int {
	if e == nil || e.ref == 0 {
		return 0
	}
	// For most checkboxes and toggles, the value is stored as a string "0" or "1"
	value := e.Value()
	if value == "1" {
		return 1
	}
	return 0
}

// Scroll scrolls the element in the given direction by the given number of lines.
// It first tries the AXScroll action, then falls back to CGEvent scroll wheel.
func (e *Element) Scroll(direction ScrollDirection, lines int) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}
	if lines <= 0 {
		return nil
	}
	x, y := e.Center()
	return cgScrollWheel(x, y, direction, lines)
}

// ScrollToVisible scrolls through ancestor scroll areas until this element is visible.
// It walks up the tree looking for an AXScrollArea or AXOutline parent and scrolls it.
func (e *Element) ScrollToVisible() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	// Walk up looking for a scrollable ancestor.
	parent := e.Parent()
	for parent != nil {
		role := parent.Role()
		if role == "AXScrollArea" || role == "AXOutline" || role == "AXList" || role == "AXTable" {
			break
		}
		next := parent.Parent()
		parent.Release()
		parent = next
	}
	if parent == nil {
		return nil // No scrollable ancestor; assume already visible.
	}
	defer parent.Release()

	// Check if element is within the visible rect of the scroll area.
	ex, ey := e.Center()
	px, py := parent.Position()
	pw, ph := parent.Size()

	// Scroll down until the element center is within the scroll area bounds.
	maxAttempts := 20
	for i := 0; i < maxAttempts; i++ {
		if ex >= px && ex <= px+pw && ey >= py && ey <= py+ph {
			return nil
		}
		var dir ScrollDirection
		if ey < py {
			dir = ScrollUp
		} else {
			dir = ScrollDown
		}
		sx, sy := parent.Center()
		if err := cgScrollWheel(sx, sy, dir, 3); err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

// TypeText focuses the element and types the given text using keyboard events.
func (e *Element) TypeText(text string) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}
	if err := e.Focus(); err != nil {
		// Focus may fail for some elements; try clicking instead.
		if err2 := e.Click(); err2 != nil {
			return fmt.Errorf("focusing element: %w", err)
		}
		time.Sleep(50 * time.Millisecond)
	}
	return TypeText(text)
}

// SelectMenuItem clicks this element (assumed to be an AXPopUpButton or AXComboBox),
// waits for the menu to open, and clicks the menu item with the given title.
func (e *Element) SelectMenuItem(title string) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	// Click to open the popup.
	if err := e.Click(); err != nil {
		return fmt.Errorf("opening popup: %w", err)
	}
	time.Sleep(200 * time.Millisecond)

	// Search descendants for a menu item with the matching title.
	item := e.Descendants().ByRole("AXMenuItem").ByTitle(title).First()
	if item == nil {
		// Menu may be a sibling or elsewhere in the window; search from app root.
		if e.app != nil {
			item = e.app.Descendants().ByRole("AXMenuItem").ByTitle(title).First()
		}
		if item == nil {
			return fmt.Errorf("%w: menu item %q", ErrElementNotFound, title)
		}
	}
	defer item.Release()
	return item.Click()
}

// SelectValue sets the value of a combo box or other value-settable element,
// then tries to select a matching item from a popup if present.
func (e *Element) SelectValue(value string) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	role := e.Role()
	if role == "AXPopUpButton" {
		return e.SelectMenuItem(value)
	}

	// For AXComboBox: set value then confirm.
	if err := e.SetValue(value); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	// Try to find a matching menu item and click it.
	item := e.Descendants().ByRole("AXMenuItem").ByTitle(value).First()
	if item != nil {
		defer item.Release()
		return item.Click()
	}
	return nil
}

// WaitAndClick waits up to timeout for the element to exist and be enabled, then clicks it.
func (e *Element) WaitAndClick(timeout time.Duration) error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if e.Exists() && e.IsEnabled() {
			return e.Click()
		}
		time.Sleep(100 * time.Millisecond)
	}
	return ErrTimeout
}

// Screenshot captures a PNG of the element's on-screen frame using screencapture.
// Returns an error if the element has no frame or the capture fails.
func (e *Element) Screenshot() ([]byte, error) {
	if e == nil || e.ref == 0 {
		return nil, ErrInvalidElement
	}
	frame := e.Frame()
	if frame.Size.Width == 0 || frame.Size.Height == 0 {
		return nil, fmt.Errorf("element has zero-size frame")
	}
	f, err := os.CreateTemp("", "ax-screenshot-*.png")
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	f.Close()
	defer os.Remove(f.Name())

	rect := fmt.Sprintf("%d,%d,%d,%d",
		int(frame.Origin.X), int(frame.Origin.Y),
		int(frame.Size.Width), int(frame.Size.Height))
	cmd := exec.Command("screencapture", "-x", "-R", rect, "-t", "png", f.Name())
	if out, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("screencapture: %w: %s", err, out)
	}
	return os.ReadFile(f.Name())
}

// ScrollAndClick scrolls the element into view and then clicks it.
func (e *Element) ScrollAndClick() error {
	if e == nil || e.ref == 0 {
		return ErrInvalidElement
	}
	if err := e.ScrollToVisible(); err != nil {
		return err
	}
	return e.Click()
}
