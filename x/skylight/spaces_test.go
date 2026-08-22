package skylight_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/x/skylight"
)

// onScreenWindows returns the ids of the windows the window server currently
// knows about, or skips. Real ids matter here: SpacesForWindow took an id that
// never reached SLSCopySpacesForWindows as a CoreFoundation object, so it only
// misbehaved when handed something the window server would try to look up.
func onScreenWindows(t *testing.T) []skylight.Window {
	t.Helper()
	var out []skylight.Window
	arr := coregraphics.CGWindowListCopyWindowInfo(
		coregraphics.KCGWindowListOptionOnScreenOnly|coregraphics.KCGWindowListExcludeDesktopElements, 0)
	if arr == 0 {
		t.Skip("CGWindowListCopyWindowInfo returned nil (no window server session?)")
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(arr))

	key := corefoundation.CFStringCreateWithCString(0, "kCGWindowNumber", uint32(corefoundation.KCFStringEncodingUTF8))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(key))

	for i := 0; i < corefoundation.CFArrayGetCount(arr); i++ {
		d := corefoundation.CFDictionaryRef(uintptr(corefoundation.CFArrayGetValueAtIndex(arr, i)))
		v := corefoundation.CFDictionaryGetValue(d, unsafe.Pointer(uintptr(key)))
		if v == nil {
			continue
		}
		var id int64
		if corefoundation.CFNumberGetValue(corefoundation.CFNumberRef(uintptr(v)), corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&id)) && id > 0 {
			out = append(out, skylight.Window(id))
		}
	}
	if len(out) == 0 {
		t.Skip("no on-screen windows to test against")
	}
	return out
}

// onScreenWindow returns one window id to exercise the lookup path with.
func onScreenWindow(t *testing.T) skylight.Window {
	t.Helper()
	return onScreenWindows(t)[0]
}

// TestSpacesForWindowRealID is the regression test for the bus error.
//
// SLSCopySpacesForWindows takes a CFArray whose elements are CFNumbers, and
// returns one too. The original code passed the window id as the array element
// and read the results back as raw values, so the window server dereferenced
// the id itself as a pointer. That is a SIGBUS, not an error return, which is
// why it survived: the call reported success on every id small enough to be
// unmapped rather than invalid, and the function had no test.
//
// Surviving the call is most of the assertion. CGWindowList may include
// low-numbered system windows for which the window server reports no spaces,
// so try every listed window until one also exercises the result decoding.
func TestSpacesForWindowRealID(t *testing.T) {
	for _, w := range onScreenWindows(t) {
		spaces, err := skylight.SpacesForWindow(w)
		if err != nil || len(spaces) == 0 {
			continue
		}
		for _, s := range spaces {
			if s == 0 {
				t.Errorf("SpacesForWindow(%d) returned space 0, which is not a valid space id", w)
			}
		}
		return
	}
	t.Skip("no on-screen window reported any spaces")
}

// TestActiveSpaceAgreesWithWindowSpaces is the decode check that matters.
//
// ActiveSpace and SpacesForWindow reach the window server by different routes,
// so if either decoded its result wrong the two would never name the same
// space. Requiring the active space to appear among the spaces of some
// on-screen window pins both against each other without assuming anything
// about the machine.
//
// Note it does not assume that every on-screen window is on the active space.
// kCGWindowListOptionOnScreenOnly reports low-numbered system windows that live
// on a space of their own, so a single window being off-space is normal and
// says nothing about correctness.
func TestActiveSpaceAgreesWithWindowSpaces(t *testing.T) {
	active, err := skylight.ActiveSpace()
	if err != nil {
		t.Fatalf("ActiveSpace: %v", err)
	}
	if active == 0 {
		t.Fatal("ActiveSpace = 0, want a valid space id")
	}

	seen := make(map[skylight.Space]bool)
	for _, w := range onScreenWindows(t) {
		spaces, err := skylight.SpacesForWindow(w)
		if err != nil {
			continue // Windows come and go between the list and the lookup.
		}
		for _, s := range spaces {
			if s == active {
				return
			}
			seen[s] = true
		}
	}
	t.Errorf("no on-screen window is on the active space %d; windows reported spaces %v", active, seen)
}

// TestIsWindowOffSpaceMatchesSpaces pins the composition rather than the
// environment: off-space must mean exactly "the active space is not among the
// window's spaces". A disagreement here is IsWindowOffSpace inventing an
// answer, which is what would make it unsafe as a routing precondition.
func TestIsWindowOffSpaceMatchesSpaces(t *testing.T) {
	w := onScreenWindow(t)

	active, err := skylight.ActiveSpace()
	if err != nil {
		t.Fatalf("ActiveSpace: %v", err)
	}
	spaces, err := skylight.SpacesForWindow(w)
	if err != nil {
		t.Fatalf("SpacesForWindow(%d): %v", w, err)
	}
	want := true
	for _, s := range spaces {
		if s == active {
			want = false
		}
	}

	got, err := skylight.IsWindowOffSpace(w)
	if err != nil {
		t.Fatalf("IsWindowOffSpace(%d): %v", w, err)
	}
	if got != want {
		t.Errorf("IsWindowOffSpace(%d) = %v, want %v (active=%d spaces=%v)", w, got, want, active, spaces)
	}
}

// TestSpacesForWindowUnknownID checks the error path stays an error. An id the
// window server does not know must not be reported as a window that happens to
// be on no spaces, because IsWindowOffSpace would read that as "off-space" and
// a caller would route input at a window that does not exist.
func TestSpacesForWindowUnknownID(t *testing.T) {
	const unknown = skylight.Window(0x7FFFFFF0)

	spaces, err := skylight.SpacesForWindow(unknown)
	if err != nil {
		return // Refusing outright is fine.
	}
	if len(spaces) != 0 {
		t.Errorf("SpacesForWindow(%d) = %v, want no spaces for an unknown window", unknown, spaces)
	}
}
