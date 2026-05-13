//go:build darwin

package objc

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

// TestAssociateBlockWithReceiver_ReleasesOnDealloc asserts that the
// associated block's storage is released when its receiver is dealloc'd.
//
// Strategy: each iteration creates a plain NSObject, attaches a Go block
// holding a counter-decrementing closure, then releases the receiver.
// The associated block is retained by the runtime, so the block's Go
// closure stays alive across the Send. Once the receiver is dealloc'd
// the association releases the block; we observe this indirectly by
// running an autorelease pool drain and confirming no leak in the
// purego block cache (block count returns to the baseline).
func TestAssociateBlockWithReceiver_ReleasesOnDealloc(t *testing.T) {
	ensureLibObjC()
	ensureAssociation()

	const n = 64

	var live atomic.Int64

	nsobject := GetClass("NSObject")
	var keys [n]byte
	for i := 0; i < n; i++ {
		obj := Send[ID](ID(nsobject), Sel("alloc"))
		obj = Send[ID](obj, Sel("init"))

		// Allocate a sentinel captured by the block closure. When the
		// receiver is dealloc'd the association releases the block,
		// which drops purego's retain on the Go closure; the sentinel
		// becomes unreferenced and its finalizer fires.
		sentinel := new(int)
		live.Add(1)
		runtime.SetFinalizer(sentinel, func(*int) { live.Add(-1) })

		block := NewBlock(func(_ Block) { _ = sentinel })

		AssociateBlockWithReceiver(obj, &keys[i], block)
		Send[struct{}](obj, Sel("release"))
	}

	// Drive a few GC cycles + autorelease pool drains so finalizers can fire.
	deadline := time.Now().Add(2 * time.Second)
	for live.Load() > 0 && time.Now().Before(deadline) {
		Send[struct{}](Send[ID](ID(GetClass("NSAutoreleasePool")), Sel("alloc")), Sel("init"))
		runtime.GC()
		runtime.Gosched()
	}

	if got := live.Load(); got != 0 {
		t.Fatalf("expected all %d associated blocks released, %d still live", n, got)
	}
}
