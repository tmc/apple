//go:build darwin

package objc

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// Generated bindings hand Objective-C a pointer derived from a Go []byte —
// BytesPointer(b) — and the slice itself never appears at the call site. The
// backing array must survive until objc_msgSend returns or the callee reads
// freed memory. Send is where that is arranged, so it is tested here rather
// than at each of the hundreds of call sites.
//
// gcInsideCall replaces one of the pre-registered msgSendN stubs so that Go
// code runs at the exact moment the process is inside objc_msgSend, which is
// the only window in which the collection could happen.
func gcInsideCall(freed *atomic.Bool, observed *bool) func(id, sel, a1, a2, a3 uintptr) uintptr {
	return func(id, sel, a1, a2, a3 uintptr) uintptr {
		for i := 0; i < 5; i++ {
			runtime.GC()
			time.Sleep(5 * time.Millisecond)
		}
		*observed = freed.Load()
		return 0
	}
}

func TestSendKeepsSliceAliveThroughDerivedPointer(t *testing.T) {
	if objcMsgSendAddr == 0 {
		t.Skip("libobjc unavailable; fast path not registered")
	}
	orig := msgSend3
	defer func() { msgSend3 = orig }()

	var freed atomic.Bool
	var collectedDuringCall bool
	msgSend3 = gcInsideCall(&freed, &collectedDuringCall)

	func() {
		b := make([]byte, 64)
		runtime.SetFinalizer(&b[0], func(*byte) { freed.Store(true) })
		Send[struct{}](ID(1), SEL(2), BytesPointer(b), uint(len(b)), uintptr(0))
	}()

	if collectedDuringCall {
		t.Error("GC reclaimed the []byte backing array while Objective-C still held the pointer; Send stopped keeping its arguments alive across the call")
	}
}

// TestKeepAliveProbeCanDetectCollection is the positive control for the test
// above. It passes a uintptr, which the collector does not trace, so the
// backing array must be reclaimed. If this reports no collection, the probe
// cannot observe the failure at all and the test above proves nothing.
func TestKeepAliveProbeCanDetectCollection(t *testing.T) {
	if objcMsgSendAddr == 0 {
		t.Skip("libobjc unavailable; fast path not registered")
	}
	orig := msgSend3
	defer func() { msgSend3 = orig }()

	var freed atomic.Bool
	var collectedDuringCall bool
	msgSend3 = gcInsideCall(&freed, &collectedDuringCall)

	func() {
		b := make([]byte, 64)
		runtime.SetFinalizer(&b[0], func(*byte) { freed.Store(true) })
		Send[struct{}](ID(1), SEL(2), uintptr(unsafe.Pointer(unsafe.SliceData(b))), uint(len(b)), uintptr(0))
	}()

	if !collectedDuringCall {
		t.Error("probe never observed a collection, so TestSendKeepsSliceAliveThroughDerivedPointer is vacuous")
	}
}
