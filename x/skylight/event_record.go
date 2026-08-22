package skylight

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
)

// EventRecord is the 248-byte record that SkyLight passes between the window
// server and its clients. Apple does not document it, so it is described here
// by hand rather than in private/skylight, which holds only generated bindings.
// Nothing generated needs it: SLPSPostEventRecordTo takes the record as a byte
// slice, so the layout is this package's concern alone.
//
// Its size is load-bearing. Offset 0x04 holds the record's own declared length,
// and _SLSEventRecordLength is a constant function returning it (mov w0, #0xf8;
// ret at 0x186d83290), so buffers passed to SLPSPostEventRecordTo must be
// exactly 248 bytes. Use [EventRecord.Bytes] rather than reslicing by hand.
//
// Field provenance, all read from the arm64e dyld shared cache of macOS 15.
// Offsets not listed here rest on structural inference only:
//
//	0x04 DeclaredLength   _SLSEventRecordLength returns 0xf8
//	0x08 EventType        decodeEventRecordForPostTo masks off the low bit and
//	                      compares against kCGEventKeyDown, which is what makes
//	                      this a type field rather than half of an opcode
//	                      (and w8, w8, #0xfffffffe; cmp w8, #0xa)
//	0x38 EventFlags       decodeEventRecordForPost ORs it with the result of
//	                      _CGXCurrentEventFlags and stores it back
//	0x3c WindowID         the window id the focus recipe stamps
//	0x48 SessionField     written from the session object when Attributes bit 2
//	                      is clear
//	0x54 Attributes       bit 2 gates SessionField above
//	0x90 KeyCode          uint16; SLPSPostEventRecordTo refuses a key event
//	                      whose value here is 0x7f
//	0xe0 AppendixPtr      _SLSEventRecordSetBaseAttributes stores the result of
//	                      _CGSEventAppendixCreate here, and
//	                      _SLSEventRecordGetContextID dereferences it as a
//	                      null-checked pointer
//
// The tail past the common header is variant by event type -- a keyboard record
// carries keyboard payload from roughly 0x88 on -- so a field there is only
// meaningful for the matching type.
//
// ActivationState is the one field carried on faith. yabai and cua-driver write
// 0x01 to focus and 0x02 to defocus and the recipe works, but SkyLight itself
// never reads the byte: scanning its __text for byte and halfword accesses at
// +0x8a finds eighteen sites, none in event-record code, and the word loads at
// +0x88 in both decode functions are on the session object rather than the
// record. The value is carried opaquely to the receiving process and
// interpreted there, so naming it means disassembling AppKit or HIToolbox.
type EventRecord struct {
	MajorVersion   uint16
	MinorVersion   uint16
	DeclaredLength uint32
	EventType      uint32
	SubtypeFlags   uint32
	LocationX      float64
	LocationY      float64
	WinLocationX   float64
	WinLocationY   float64
	EventTime      uint64
	EventFlags     uint32
	WindowID       uint32
	ConnectionID   uint32
	_              [4]byte // 0x44
	SessionField   uint32  // 0x48
	_              [8]byte // 0x4c
	Attributes     uint32  // 0x54
	_              [50]byte
	// ActivationState is 0x01 to focus and 0x02 to defocus. See the type
	// documentation: SkyLight does not read it.
	ActivationState uint8   // 0x8a
	_               [5]byte // 0x8b
	KeyCode         uint16  // 0x90
	_               [78]byte
	AppendixPtr     uintptr // 0xe0
	_               [16]byte
}

// eventRecordLength is the size an EventRecord declares at offset 0x04.
const eventRecordLength = 248

// Bytes exposes r as the raw buffer SLPSPostEventRecordTo expects. The result
// aliases r, so a write through r after the call is visible through the slice.
//
// The array length is written out rather than derived so that a change to
// EventRecord fails to compile here instead of silently truncating the buffer.
func (r *EventRecord) Bytes() []byte {
	var _ [1]struct{} = [unsafe.Sizeof(*r) - eventRecordLength + 1]struct{}{}
	return (*[eventRecordLength]byte)(unsafe.Pointer(r))[:]
}

// eventRecordOffsets are the offsets within the opaque __CGEvent at which the
// record pointer has been observed. __CGEvent begins with a CFRuntimeBase and a
// uint32, which puts the pointer at 24 on arm64; the rest are probed in case
// padding or an added field moves it.
var eventRecordOffsets = []uintptr{24, 32, 16, 40, 48}

// EventRecordFromCGEvent returns the EventRecord backing event.
//
// __CGEvent is opaque, so the record is found by probing the offsets at which
// its pointer has been observed and accepting the first candidate that declares
// itself 248 bytes long. That check is what makes the result trustworthy: an
// unrelated pointer field is very unlikely to address memory whose second word
// is exactly 248. The record is owned by event and stays valid only as long as
// the caller holds a reference to it.
//
// An error is returned rather than a possibly-bad pointer when no candidate
// qualifies, which is the expected result if the layout changes.
func EventRecordFromCGEvent(event coregraphics.CGEvent) (*EventRecord, error) {
	if event == 0 {
		return nil, fmt.Errorf("skylight: CGEvent is nil")
	}
	base := uintptr(event)
	for _, off := range eventRecordOffsets {
		ptr := *(*unsafe.Pointer)(unsafe.Pointer(base + off))
		// Reject null and the unmapped first page before dereferencing.
		if ptr == nil || uintptr(ptr) < 0x10000 {
			continue
		}
		if rec := (*EventRecord)(ptr); rec.DeclaredLength == eventRecordLength {
			return rec, nil
		}
	}
	return nil, fmt.Errorf("skylight: no %d-byte event record found in CGEvent %#x", eventRecordLength, base)
}
