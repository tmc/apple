package skylight_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/x/skylight"
)

func newMouseEvent(t *testing.T, typ coregraphics.CGEventType) coregraphics.CGEventRef {
	t.Helper()
	ev := coregraphics.CGEventCreateMouseEvent(0, typ, corefoundation.CGPoint{X: 10, Y: 10}, 0)
	if ev == 0 {
		t.Skip("could not create CGEvent (no window server session?)")
	}
	t.Cleanup(func() { corefoundation.CFRelease(corefoundation.CFTypeRef(ev)) })
	return ev
}

// TestRouteMouseEventToWindow checks that the routing fields survive being
// written, by reading each one back through the public CGEvent accessor. That
// readback is the point: SLEventSetIntegerValueField returning no error only
// says the call was made, not that the field exists or kept the value.
func TestRouteMouseEventToWindow(t *testing.T) {
	const (
		wid  = 12345
		pid  = 4321
		grp  = 99
		clks = 1
	)
	ev := newMouseEvent(t, coregraphics.KCGEventLeftMouseDown)

	if err := skylight.RouteMouseEventToWindow(ev, wid, pid, corefoundation.CGPoint{X: 7, Y: 11}, clks, grp); err != nil {
		t.Fatalf("RouteMouseEventToWindow: %v", err)
	}

	for _, tt := range []struct {
		name  string
		field coregraphics.CGEventField
		want  int64
	}{
		{"clickState", coregraphics.KCGMouseEventClickState, clks},
		{"targetPID", coregraphics.KCGEventTargetUnixProcessID, pid},
		{"windowNumber", 51, wid},
		{"clickGroup", 58, grp},
		{"windowUnderPointer", 91, wid},
		{"windowUnderPointerThatCanHandle", 92, wid},
	} {
		if got := coregraphics.CGEventGetIntegerValueField(ev, tt.field); got != tt.want {
			t.Errorf("field %d (%s) = %d, want %d", tt.field, tt.name, got, tt.want)
		}
	}
}

func TestRouteMouseEventToWindowNilEvent(t *testing.T) {
	if err := skylight.RouteMouseEventToWindow(0, 1, 1, corefoundation.CGPoint{}, 1, 1); err == nil {
		t.Error("RouteMouseEventToWindow(0, ...) = nil, want error")
	}
}

// TestPostEventToPID exercises SLEventPostToPid, which was bound but never
// called. Posting to the test process itself is safe: the event is addressed at
// a pid with no event-handling run loop, so it is accepted and dropped.
func TestPostEventToPID(t *testing.T) {
	ev := newMouseEvent(t, coregraphics.KCGEventMouseMoved)
	if err := skylight.PostEventToPID(os.Getpid(), ev); err != nil {
		t.Errorf("PostEventToPID: %v", err)
	}
}

func TestPostEventToPIDNilEvent(t *testing.T) {
	if err := skylight.PostEventToPID(os.Getpid(), 0); err == nil {
		t.Error("PostEventToPID(_, 0) = nil, want error")
	}
}

// TestClickBackgroundedWindow is the end-to-end check that the routing pair
// actually delivers to a window that is not frontmost -- the claim the whole
// off-space input design rests on. It cannot run unattended: it needs a real
// target and a human to confirm the click landed.
//
// Run it as:
//
//	SKYLIGHT_TARGET_WID=<window id> SKYLIGHT_TARGET_PID=<pid> \
//	  go test ./x/skylight -run TestClickBackgroundedWindow -v
//
// with the target window visible but behind another window, or on another
// space. A pass means the SPIs accepted the sequence, not that the click
// landed; watch the target to judge that.
func TestClickBackgroundedWindow(t *testing.T) {
	widStr, pidStr := os.Getenv("SKYLIGHT_TARGET_WID"), os.Getenv("SKYLIGHT_TARGET_PID")
	if widStr == "" || pidStr == "" {
		t.Skip("set SKYLIGHT_TARGET_WID and SKYLIGHT_TARGET_PID to run")
	}
	wid, err := strconv.Atoi(widStr)
	if err != nil {
		t.Fatalf("SKYLIGHT_TARGET_WID: %v", err)
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		t.Fatalf("SKYLIGHT_TARGET_PID: %v", err)
	}

	offSpace, err := skylight.IsWindowOffSpace(skylight.Window(wid))
	if err != nil {
		t.Logf("IsWindowOffSpace: %v", err)
	} else {
		t.Logf("target window %d off-space: %v", wid, offSpace)
	}

	// Window-local coordinates of the point to click.
	local := corefoundation.CGPoint{X: 60, Y: 60}
	if v := os.Getenv("SKYLIGHT_TARGET_X"); v != "" {
		x, err := strconv.ParseFloat(v, 64)
		if err != nil {
			t.Fatalf("SKYLIGHT_TARGET_X: %v", err)
		}
		local.X = x
	}
	if v := os.Getenv("SKYLIGHT_TARGET_Y"); v != "" {
		y, err := strconv.ParseFloat(v, 64)
		if err != nil {
			t.Fatalf("SKYLIGHT_TARGET_Y: %v", err)
		}
		local.Y = y
	}
	if os.Getenv("SKYLIGHT_FOCUS_FIRST") != "" {
		if err := skylight.FocusWithoutRaise(pid, skylight.Window(wid)); err != nil {
			t.Logf("FocusWithoutRaise: %v", err)
		} else {
			t.Log("FocusWithoutRaise ok")
		}
	}
	const clickGroup = 1

	for _, step := range []struct {
		name string
		typ  coregraphics.CGEventType
	}{
		{"down", coregraphics.KCGEventLeftMouseDown},
		{"up", coregraphics.KCGEventLeftMouseUp},
	} {
		ev := newMouseEvent(t, step.typ)
		if err := skylight.RouteMouseEventToWindow(ev, skylight.Window(wid), pid, local, 1, clickGroup); err != nil {
			t.Fatalf("%s: RouteMouseEventToWindow: %v", step.name, err)
		}
		if err := skylight.PostEventToPID(pid, ev); err != nil {
			t.Fatalf("%s: PostEventToPID: %v", step.name, err)
		}
		t.Logf("%s posted to pid %d at window-local (%v, %v)", step.name, pid, local.X, local.Y)
	}
	t.Log("sequence accepted; confirm at the target whether the click landed")
}
