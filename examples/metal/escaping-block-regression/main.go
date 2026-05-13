// escaping-block-regression demonstrates an ObjC API pattern that broke
// before applegen learned to skip `defer block.Release()` for escaping
// blocks.
//
// MTLCommandBuffer.AddCompletedHandler retains the supplied block and
// invokes it on GPU completion, possibly after the calling Go function
// has returned. The pre-fix generator emitted:
//
//	func (o MTLCommandBufferObject) AddCompletedHandler(block ...) {
//	    _block0 := objc.NewBlock(...)
//	    defer _block0.Release()    // <-- decrements refcount before Metal fires the block
//	    objc.Send(...)
//	}
//
// In practice Metal performs a synchronous `[block copy]` inside the
// addCompletedHandler: call so the defer leaves refcount >= 1 and the
// handler fires. But the pattern was a latent use-after-free for any
// framework that defers the copy (e.g. AVAudioEngine.installTapOnBus:),
// and even for Metal the dispose-on-zero-refcount removes the Go closure
// from purego's cache, leaving a window for crashes under GC pressure.
//
// Post-fix (appledocs c09228c541) the generator omits the defer for
// selectors matching install*/set*Block:/set*Handler:/add*Handler:/
// subscribe*/register*/observe*. This program submits many command
// buffers with completion handlers, forces GC between submission and
// completion, and asserts that every handler fired.
package main

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"

	"github.com/tmc/apple/metal"
)

const iterations = 200

func main() {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.ID == 0 {
		fmt.Fprintln(os.Stderr, "metal: no default device available; skipping")
		os.Exit(0)
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		fmt.Fprintln(os.Stderr, "metal: failed to create command queue")
		os.Exit(1)
	}

	var fired int64
	for i := 0; i < iterations; i++ {
		buf := queue.CommandBuffer()
		buf.AddCompletedHandler(func(metal.MTLCommandBuffer) {
			atomic.AddInt64(&fired, 1)
		})
		buf.Commit()
		runtime.GC()
		buf.WaitUntilCompleted()
	}

	got := atomic.LoadInt64(&fired)
	if got != iterations {
		fmt.Fprintf(os.Stderr, "regression: handler fired %d/%d times\n", got, iterations)
		os.Exit(1)
	}
	fmt.Printf("ok: %d/%d MTLCommandBuffer completion handlers fired\n", got, iterations)
}
