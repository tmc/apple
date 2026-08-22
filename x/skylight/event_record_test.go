package skylight_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/x/skylight"
)

// TestEventRecordFromCGEvent checks the probe against records the system builds
// itself, which is the only way to tell a correct offset from a lucky one: the
// fields we set through the public CGEvent API must show up where the struct
// says they are.
func TestEventRecordFromCGEvent(t *testing.T) {
	tests := []struct {
		name      string
		create    func() coregraphics.CGEventRef
		wantType  uint32
		wantFlags bool
	}{
		{
			name:     "keyboard down",
			create:   func() coregraphics.CGEventRef { return coregraphics.CGEventCreateKeyboardEvent(0, 0, true) },
			wantType: uint32(coregraphics.KCGEventKeyDown),
		},
		{
			name:     "keyboard up",
			create:   func() coregraphics.CGEventRef { return coregraphics.CGEventCreateKeyboardEvent(0, 0, false) },
			wantType: uint32(coregraphics.KCGEventKeyUp),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := tt.create()
			if ev == 0 {
				t.Skip("could not create CGEvent (no window server session?)")
			}
			defer corefoundation.CFRelease(corefoundation.CFTypeRef(ev))

			rec, err := skylight.EventRecordFromCGEvent(ev)
			if err != nil {
				t.Fatalf("EventRecordFromCGEvent: %v", err)
			}
			if rec.DeclaredLength != 248 {
				t.Errorf("DeclaredLength = %d, want 248", rec.DeclaredLength)
			}
			if rec.EventType != tt.wantType {
				t.Errorf("EventType = %d, want %d (CGEventGetType agrees: %d)",
					rec.EventType, tt.wantType, coregraphics.CGEventGetType(ev))
			}
		})
	}
}

// TestEventRecordFromCGEventNil pins the error path: a bad event must produce an
// error, never a pointer the caller would go on to dereference.
func TestEventRecordFromCGEventNil(t *testing.T) {
	if rec, err := skylight.EventRecordFromCGEvent(0); err == nil {
		t.Errorf("EventRecordFromCGEvent(0) = %v, want error", rec)
	}
}

// TestEventRecordTailBytes reports what a real record holds around offset 0x8a,
// the byte yabai and cua-driver write 0x01/0x02 to for focus/defocus and the one
// field of the focus record still unexplained.
//
// It does not settle that question, and the reason is worth recording. A
// keyboard record holds keyboard payload here (observed: 0x88 reads
// 01 00 00 00 fc 00 61 00, with 0x8a itself zero), so the tail is variant by
// event type rather than common to all records. Only a kCGSEventAppActive
// (type 0x0D) record would show the focus field, and those are built by the
// window server, not by CGEventCreate -- so this route cannot reach it.
// Settling 0x8a needs the static route: the encoder reached through
// _CGSEncodeEventRecord from _SLPSPostEventRecordTo.
//
// The value here is the negative result plus the confirmation that extraction
// works on records we did not construct. It is a probe, not an assertion: it
// fails only if the extraction itself breaks. Run with -v to read the bytes.
func TestEventRecordTailBytes(t *testing.T) {
	ev := coregraphics.CGEventCreateKeyboardEvent(0, 0, true)
	if ev == 0 {
		t.Skip("could not create CGEvent")
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(ev))

	rec, err := skylight.EventRecordFromCGEvent(ev)
	if err != nil {
		t.Fatalf("EventRecordFromCGEvent: %v", err)
	}

	raw := (*[248]byte)(unsafe.Pointer(rec))
	for off := 0x80; off < 0x90; off += 8 {
		t.Logf("0x%02x: % 02x", off, raw[off:off+8])
	}
	t.Logf("ActivationState (0x8a) = %#02x", rec.ActivationState)
}
