package input

import (
	"fmt"
	"os"
	"sync"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// CGEvent mouse and keyboard event types.
const (
	EventNull           = 0
	EventLeftMouseDown  = 1
	EventLeftMouseUp    = 2
	EventRightMouseDown = 3
	EventRightMouseUp   = 4
	EventMouseMoved     = 5
	EventKeyDown        = 10
	EventKeyUp          = 11
)

// HIDEventTap is the location for posting events to the HID system.
const HIDEventTap = 0

// Common virtual key codes.
const (
	KeyReturn = 36
	KeyTab    = 48
	KeySpace  = 49
	KeyDelete = 51
	KeyEscape = 53
)

var (
	cgEventCreateMouseEvent         func(source uintptr, mouseType uint32, mouseCursorPosition corefoundation.CGPoint, mouseButton uint32) uintptr
	cgEventCreateKeyboardEvent      func(source uintptr, virtualKey uint16, keyDown bool) uintptr
	cgEventPost                     func(tap uint32, event uintptr)
	cgEventPostToPid                func(pid int32, event uintptr)
	cgEventSetFlags                 func(event uintptr, flags uint64)
	cgEventKeyboardSetUnicodeString func(event uintptr, stringLength uint64, unicodeString *uint16)

	initOnce sync.Once
	initErr  error
)

func ensureInit() error {
	initOnce.Do(func() {
		appServices, err := purego.Dlopen("/System/Library/Frameworks/ApplicationServices.framework/ApplicationServices", purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			initErr = fmt.Errorf("load ApplicationServices: %w", err)
			return
		}
		purego.RegisterLibFunc(&cgEventCreateMouseEvent, appServices, "CGEventCreateMouseEvent")
		purego.RegisterLibFunc(&cgEventCreateKeyboardEvent, appServices, "CGEventCreateKeyboardEvent")

		coreGraphics, err := purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			initErr = fmt.Errorf("load CoreGraphics: %w", err)
			return
		}
		purego.RegisterLibFunc(&cgEventPost, coreGraphics, "CGEventPost")
		purego.RegisterLibFunc(&cgEventPostToPid, coreGraphics, "CGEventPostToPid")
		purego.RegisterLibFunc(&cgEventSetFlags, coreGraphics, "CGEventSetFlags")
		purego.RegisterLibFunc(&cgEventKeyboardSetUnicodeString, coreGraphics, "CGEventKeyboardSetUnicodeString")
	})
	return initErr
}

// CreateMouseEvent creates a mouse event at the given position.
// The source parameter is typically 0 (no source). mouseType should be one of
// the Event* mouse constants. mouseButton is 0 for left, 1 for right.
func CreateMouseEvent(source uintptr, mouseType uint32, position corefoundation.CGPoint, mouseButton uint32) (uintptr, error) {
	if err := ensureInit(); err != nil {
		return 0, err
	}
	return cgEventCreateMouseEvent(source, mouseType, position, mouseButton), nil
}

// CreateKeyboardEvent creates a keyboard event. virtualKey is the macOS
// virtual key code (e.g., KeyReturn=36, KeyTab=48).
func CreateKeyboardEvent(source uintptr, virtualKey uint16, keyDown bool) (uintptr, error) {
	if err := ensureInit(); err != nil {
		return 0, err
	}
	return cgEventCreateKeyboardEvent(source, virtualKey, keyDown), nil
}

// PostToSelf posts an event to this process. This is useful for directing
// keystrokes to a VM window owned by the current process rather than
// whatever app the user has focused.
func PostToSelf(event uintptr) error {
	if err := ensureInit(); err != nil {
		return err
	}
	cgEventPostToPid(int32(os.Getpid()), event)
	return nil
}

// Post posts an event to the given HID event tap location.
// Use HIDEventTap for the system HID tap.
func Post(tap uint32, event uintptr) error {
	if err := ensureInit(); err != nil {
		return err
	}
	cgEventPost(tap, event)
	return nil
}

// SetFlags sets the modifier flags (shift, control, etc.) on an event.
func SetFlags(event uintptr, flags uint64) error {
	if err := ensureInit(); err != nil {
		return err
	}
	cgEventSetFlags(event, flags)
	return nil
}

// SetUnicodeString sets the Unicode string on a keyboard event.
// This allows typing arbitrary characters that don't have a direct keycode.
func SetUnicodeString(event uintptr, s string) error {
	if err := ensureInit(); err != nil {
		return err
	}
	if len(s) == 0 {
		return nil
	}
	runes := []rune(s)
	utf16 := make([]uint16, len(runes))
	for i, r := range runes {
		utf16[i] = uint16(r)
	}
	cgEventKeyboardSetUnicodeString(event, uint64(len(utf16)), &utf16[0])
	return nil
}

// TypeCharacter types a single character by creating key down/up events with
// the character set as a Unicode string. Events are posted to the current
// process.
func TypeCharacter(char rune) error {
	if err := ensureInit(); err != nil {
		return err
	}
	eventDown, err := CreateKeyboardEvent(0, 0, true)
	if err != nil {
		return err
	}
	if eventDown == 0 {
		return fmt.Errorf("create key down event")
	}
	SetUnicodeString(eventDown, string(char))
	PostToSelf(eventDown)

	eventUp, err := CreateKeyboardEvent(0, 0, false)
	if err != nil {
		return err
	}
	if eventUp == 0 {
		return fmt.Errorf("create key up event")
	}
	SetUnicodeString(eventUp, string(char))
	PostToSelf(eventUp)
	return nil
}
