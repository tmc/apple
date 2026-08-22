package skylight

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/private/skylight"
)

// Undocumented CGEventField values used to address a mouse event at a specific
// window. Apple documents fields 0-7 and 40 but not these; the names and values
// follow cua-driver, which in turn took them from Chromium's event handling.
//
// Field 51 is the AppKit windowNumber an NSEvent bridged from this CGEvent will
// report. Fields 91 and 92 are what the window server consults instead of
// hit-testing the global cursor position, which is what makes delivery to a
// window that is not frontmost possible at all. Field 58 groups the events of
// one gesture so a click's down and up are not treated as unrelated.
const (
	fieldWindowNumber                                  coregraphics.CGEventField = 51
	fieldClickGroupID                                  coregraphics.CGEventField = 58
	fieldWindowUnderMousePointer                       coregraphics.CGEventField = 91
	fieldWindowUnderMousePointerThatCanHandleThisEvent coregraphics.CGEventField = 92
)

// mouseEventSubtypeTouch is the subtype cua-driver stamps on synthetic mouse
// events so that targets which filter on subtype accept them.
const mouseEventSubtypeTouch = 3

// RouteMouseEventToWindow addresses event at window w of process pid, using the
// window-local point local rather than the global cursor position.
//
// A synthetic mouse event posted to the HID tap is delivered by hit-testing the
// cursor position against the window server's front-to-back order, so it lands
// on whatever is frontmost at that point -- not necessarily w, and never w when
// w is on another space. Stamping the window-local point and the routing fields
// makes the window server use the point and window id directly instead.
//
// Call this on an event built by CGEventCreateMouseEvent before posting it with
// [PostEventToPID]. Posting a routed event through the HID tap ignores the
// routing entirely.
//
// clickState is the click count: 1 for a single click, 2 for a double click.
// clickGroupID ties the events of one gesture together; use the same value for
// the down and up of a click, and a different value for the next click.
func RouteMouseEventToWindow(event coregraphics.CGEvent, w Window, pid int, local corefoundation.CGPoint, clickState, clickGroupID int64) error {
	if event == 0 {
		return fmt.Errorf("skylight: CGEvent is nil")
	}
	if err := skylight.CGEventSetWindowLocation(event, local); err != nil {
		return fmt.Errorf("skylight: CGEventSetWindowLocation: %w", err)
	}

	wid := int64(w)
	fields := []struct {
		field coregraphics.CGEventField
		value int64
	}{
		{coregraphics.KCGMouseEventClickState, clickState},
		{coregraphics.KCGMouseEventButtonNumber, 0},
		{coregraphics.KCGMouseEventSubtype, mouseEventSubtypeTouch},
		{fieldWindowNumber, wid},
		{fieldClickGroupID, clickGroupID},
		{fieldWindowUnderMousePointer, wid},
		{fieldWindowUnderMousePointerThatCanHandleThisEvent, wid},
		// The target pid is what lets a process filtering synthetic events
		// recognise the event as addressed to it rather than broadcast.
		{coregraphics.KCGEventTargetUnixProcessID, int64(pid)},
	}
	for _, f := range fields {
		if err := skylight.SLEventSetIntegerValueField(event, f.field, f.value); err != nil {
			return fmt.Errorf("skylight: SLEventSetIntegerValueField(%d): %w", f.field, err)
		}
	}
	return nil
}

// PostEventToPID delivers event to pid directly, bypassing the global HID tap
// and the hit-testing that comes with it.
//
// It posts twice, by design. SLEventPostToPid reaches Catalyst and Chromium
// targets that the public path does not, and CGEventPostToPid reaches AppKit
// targets that SkyLight drops. cua-driver does the same and calls it
// belt-and-suspenders; neither path alone covers every target.
//
// An error is returned only if the SkyLight path fails to be called at all. A
// non-zero status from it is not treated as fatal because the public path may
// still deliver, so callers that need to know an event arrived must check for
// the effect rather than trusting the return.
func PostEventToPID(pid int, event coregraphics.CGEvent) error {
	if event == 0 {
		return fmt.Errorf("skylight: CGEvent is nil")
	}
	_, err := skylight.SLEventPostToPid(int32(pid), event)
	coregraphics.CGEventPostToPid(int32(pid), event)
	if err != nil {
		return fmt.Errorf("skylight: SLEventPostToPid: %w", err)
	}
	return nil
}
