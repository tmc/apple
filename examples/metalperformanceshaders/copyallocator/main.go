//go:build !mpsfixedcopyallocator

// Command copyallocator is a failing regression gate for two applegen defects
// that together make the MPS fallback copy allocator unusable from Go.
//
// It is expected to FAIL on the tree it was written against (release-v0.7.0,
// 2026-08-16). It fails by crashing a child process with SIGSEGV inside
// Metal Performance Shaders. Do not "fix" this example by weakening it: it
// is the acceptance test for the generator fix, and an example that passes
// today would prove nothing.
//
// # What the API is
//
// MPSImage.framework/Headers/MPSImageKernel.h:109 declares
//
//	typedef id<MTLTexture> __nonnull NS_RETURNS_RETAINED (^MPSCopyAllocator)(
//	    MPSKernel * __nonnull filter,
//	    id<MTLCommandBuffer> __nonnull commandBuffer,
//	    id<MTLTexture> __nonnull sourceTexture);
//
//	- (BOOL) encodeToCommandBuffer:(id<MTLCommandBuffer>)commandBuffer
//	                inPlaceTexture:(__strong id<MTLTexture> * __nonnull)texture
//	         fallbackCopyAllocator:(MPSCopyAllocator __nullable)copyAllocator;
//
// The kernel first tries to filter the texture in place. When that is not
// possible it calls the allocator, which must RETURN a fresh texture at +1,
// filters into it, releases the old texture and writes the new one back
// through the inPlaceTexture pointer. Both directions matter: the block
// returns a value, and the texture argument is a pointer to the caller's
// variable, not the texture itself.
//
// # Defect 1: the block's return value was folded into its name
//
// metalperformanceshaders/blocks.gen.go:158 emits
//
//	type MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler = func(*MPSKernel, metal.MTLCommandBuffer, metal.MTLTexture)
//
// The leading "MTLTexture" in the generated name is the block's RETURN type,
// which the generator consumed as part of the identifier and then dropped.
// Worse, blocks.gen.go:171 binds the block's first PARAMETER to a variable
// called `result`, so the return slot is not merely missing, it is
// misattributed. The Go handler has no way to hand a texture back.
//
// Observed consequence: the handler runs with fully valid arguments and then
// SIGSEGVs in objc_msgSend the instant it returns, because MPS messages
// whatever garbage the trampoline left in x0 as the returned texture.
//
// A correct fix must give the handler an `objc.ID` (or metal.MTLTexture)
// return AND honour NS_RETURNS_RETAINED, handing the texture back at +1.
//
// # Defect 2: the __strong id* out-parameter was passed by value
//
// metalperformanceshaders/mps_unary_image_kernel.gen.go:228 emits
//
//	func (u MPSUnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(
//	    commandBuffer metal.MTLCommandBuffer,
//	    texture metal.MTLTexture,             // <-- Apple's type is id<MTLTexture> *
//	    copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool
//
// so the texture object is passed where MPS expects the ADDRESS of a texture
// variable. MPS dereferences the object pointer as an `id *` and writes
// through it. Observed: a deterministic SIGSEGV inside the encode call
// itself, before the allocator is ever entered. Proven by A/B — a
// hand-written objc.Send passing unsafe.Pointer(&texID) runs cleanly 5/5 and
// MPS rewrites the variable as documented.
//
// Fixing only one of the two still crashes. Defect 2 kills the call before
// the allocator runs; defect 1 kills it on the way out of the allocator.
//
// # Why this program has two processes
//
// The failure is a SIGSEGV raised on a cgo/ObjC stack. Go's recover() cannot
// catch it — the process dies. So the assertions run in a child process and
// the parent scores it. The parent requires a positive marker line from the
// child: an empty log, a missing binary (exit 127) or a child that dies
// before printing scores as FAIL, never as success.
//
// # What the example asserts
//
// The five things a working copy allocator must do, and where each stands
// against the SHIPPED generated API:
//
//  1. the allocator is invoked                     — expressible (marker line)
//  2. its three arguments are valid objects        — expressible (each is
//     messaged with -description and -class, not merely nil-checked)
//  3. the texture the allocator returns is the one
//     MPS filters into, i.e. the caller's texture
//     variable is REPLACED by it                   — INEXPRESSIBLE: the handler
//     cannot return (defect 1) and the caller passes no variable (defect 2)
//  4. encode returns true and the output holds the
//     filtered image                               — INEXPRESSIBLE in the
//     allocator path for the same reason: the filtered result lands in a
//     texture this program can never be handed
//  5. NS_RETURNS_RETAINED refcount correctness     — NOT CHECKED, and
//     deliberately not faked. -retainCount is unreliable under ARC/taggedptr
//     and MPS may retain the texture again internally; there is no sound
//     in-process check. The right proof is a leaks/heap run against a fixed
//     binding, out of scope here.
//
// Assertions 3 and 4 are written out as the code one WOULD write, in
// fixed.go, behind the `mpsfixedcopyallocator` build tag so that this
// package still builds today. They are not silently dropped. Once the
// generator emits the corrected signatures, that file compiles and
//
//	go run -tags mpsfixedcopyallocator ./examples/metalperformanceshaders/copyallocator
//
// runs all four checkable assertions. Note that the fix necessarily changes
// the shipped signature, so this file's call site will stop compiling at
// that point — that compile error is the last step of the gate, and the
// maintainer landing the fix should delete the today-path below and promote
// fixed.go.
//
// Usage:
//
//	copyallocator          # parent: runs the child and scores it
//	copyallocator -child   # the crashing case; not meant to be run directly
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
)

func init() {
	runtime.LockOSThread()
}

// Marker lines. The parent scores the child on these alone, so each one has
// to be printed from a point in the child that proves the thing it claims.
const (
	markerBegin     = "COPYALLOC-CHILD-BEGIN"
	markerSetup     = "COPYALLOC-SETUP-OK"
	markerAllocator = "COPYALLOC-ALLOCATOR-ENTERED"
	markerArgs      = "COPYALLOC-ARGS-VALID"
	markerAllocated = "COPYALLOC-ALLOCATOR-RETURNED-TEXTURE"
	markerReturned  = "COPYALLOC-ENCODE-RETURNED"
)

const (
	width  = 64
	height = 64
	// kernel is the AreaMax window. MPSImageGaussianBlur always succeeds in
	// place on Apple silicon and so never calls the allocator, which would
	// make this example vacuous. A 7x7 AreaMax does force the fallback.
	kernel = 7
)

var child = flag.Bool("child", false, "run the crashing case in this process")

func main() {
	flag.Parse()
	if *child {
		runChild()
		return
	}
	os.Exit(runParent())
}

// runParent re-execs this program with -child, then scores the result. It
// deliberately treats "no output" and "did not start" as failures.
func runParent() int {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "copyallocator: cannot find own executable: %v\n", err)
		return 1
	}
	cmd := exec.Command(exe, "-child")
	out, runErr := cmd.CombinedOutput()
	log := string(out)

	fmt.Printf("--- child output (%d bytes) ---\n", len(out))
	fmt.Print(digest(log))
	fmt.Println("--- end child output ---")

	status := describeExit(runErr)
	fmt.Printf("child exit: %s\n", status)

	// A positive marker is required. Without it the child never reached its
	// own first statement, so nothing below is meaningful.
	if !strings.Contains(log, markerBegin) {
		fmt.Println("FAIL: the child printed no start marker; it did not run (missing binary, exec failure, or an early abort)")
		return 1
	}

	got := map[string]bool{}
	for _, m := range []string{markerSetup, markerAllocator, markerArgs, markerReturned} {
		got[m] = strings.Contains(log, m)
	}
	for _, m := range []string{markerSetup, markerAllocator, markerArgs, markerReturned} {
		fmt.Printf("  %-32s %v\n", m, got[m])
	}

	// Two different defects both stop the child after setup and before the
	// allocator, so the marker gap alone does not name one. They are told
	// apart by HOW the child died: defect 1 is a Go panic raised in
	// objc.NewBlock while the block is still being built, on the Go side of
	// the boundary, before any Objective-C runs. Defect 2 is a fault inside
	// MPS after the call crosses over. Reporting the marker gap as defect 2
	// unconditionally was wrong: on this tree it is defect 1 that fires, and
	// defect 2 is currently UNMEASURED because the call never reaches MPS.
	blockPanic := strings.Contains(log, "panic: objc:") &&
		strings.Contains(log, "objc.NewBlock")

	switch {
	case !got[markerSetup]:
		fmt.Println("FAIL: the child died during Metal/MPS setup, before exercising the copy allocator")
	case !got[markerAllocator] && blockPanic:
		fmt.Println("FAIL (expected on this tree): the child panicked in objc.NewBlock while")
		fmt.Println("      constructing the fallback-allocator block, before any Objective-C ran.")
		fmt.Println("      This is DEFECT 1: the generated block trampoline takes an interface")
		fmt.Println("      parameter, and encodeType cannot encode one, so the constructor is")
		fmt.Println("      unreachable for every caller. See blocks.gen.go.")
		fmt.Println("      DEFECT 2 is UNMEASURED by this run: the call never reached MPS.")
	case !got[markerAllocator]:
		fmt.Println("FAIL: MPS crashed inside")
		fmt.Println("      encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator: before the")
		fmt.Println("      fallback allocator was entered. This is DEFECT 2: the generated wrapper")
		fmt.Println("      passes the texture object where Apple declares __strong id<MTLTexture> *,")
		fmt.Println("      so MPS dereferences an object pointer as an id* and writes through it.")
		fmt.Println("      See mps_unary_image_kernel.gen.go.")
	case !got[markerReturned]:
		fmt.Println("FAIL (expected once defect 2 alone is fixed): the allocator was entered with")
		fmt.Println("      valid arguments and then the process died on the way out. This is")
		fmt.Println("      DEFECT 1: the generated block type has no return value, so MPS reads")
		fmt.Println("      whatever is in x0 as the returned texture and messages it.")
		fmt.Println("      See blocks.gen.go:158.")
	case runErr != nil:
		fmt.Println("FAIL: the child reached the end of the encode call but exited non-zero")
	default:
		// Reaching here means the crash is gone, but assertions 3 and 4 are
		// still unexpressed against the shipped signature, so this is not a
		// pass either.
		fmt.Println("FAIL: the child survived, but assertions 3 and 4 (the allocator's texture is")
		fmt.Println("      actually used, and the filtered result is numerically correct) cannot be")
		fmt.Println("      expressed against the shipped signature. Land the generator fix and run")
		fmt.Println("      with -tags mpsfixedcopyallocator.")
	}
	fmt.Println()
	fmt.Println("This example is a gate for a v0.7.0 blocker. It passes only when the generator")
	fmt.Println("emits (a) a block type that returns a texture at +1 and (b) an encode wrapper")
	fmt.Println("taking *metal.MTLTexture, and fixed.go is promoted over the path in main.go.")
	return 1
}

// digest trims a Go fatal-signal dump down to the part that identifies the
// fault: the child's own lines, the SIGSEGV header, and any stack frame in
// this repository. Scoring is done on the full log, never on this.
// COPYALLOC_FULL_LOG=1 prints everything.
func digest(log string) string {
	if os.Getenv("COPYALLOC_FULL_LOG") != "" {
		if !strings.HasSuffix(log, "\n") && log != "" {
			return log + "\n"
		}
		return log
	}
	var b strings.Builder
	elided := 0
	for _, line := range strings.Split(strings.TrimRight(log, "\n"), "\n") {
		keep := strings.HasPrefix(line, "COPYALLOC-") ||
			strings.HasPrefix(line, "child:") ||
			strings.HasPrefix(line, "SIGSEGV") ||
			strings.HasPrefix(line, "signal ") ||
			strings.HasPrefix(line, "fatal error") ||
			strings.Contains(line, "github.com/tmc/apple/")
		if keep {
			fmt.Fprintf(&b, "%s\n", line)
			continue
		}
		elided++
	}
	if elided > 0 {
		fmt.Fprintf(&b, "[%d further lines of runtime traceback elided; set COPYALLOC_FULL_LOG=1 to see them]\n", elided)
	}
	return b.String()
}

func describeExit(err error) string {
	if err == nil {
		return "0"
	}
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		if ws, ok := ee.ProcessState.Sys().(syscall.WaitStatus); ok && ws.Signaled() {
			return fmt.Sprintf("killed by signal %v (%s)", ws.Signal(), ws.Signal())
		}
		return fmt.Sprintf("exit status %d", ee.ExitCode())
	}
	return err.Error()
}

// runChild exercises the shipped API. On the tree this was written against
// it does not return: MPS faults inside the encode call.
func runChild() {
	fmt.Println(markerBegin)
	os.Stdout.Sync()

	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		fmt.Println("child: no Metal device available")
		os.Exit(1)
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		fmt.Println("child: could not create a command queue")
		os.Exit(1)
	}
	filter := mps.NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device, kernel, kernel)
	if filter.GetID() == 0 {
		fmt.Println("child: could not create MPSImageAreaMax")
		os.Exit(1)
	}
	texture, err := newTexture(device)
	if err != nil {
		fmt.Printf("child: %v\n", err)
		os.Exit(1)
	}

	// An impulse: AreaMax spreads the peak over the whole window, so the
	// output is checkable without knowing MPS's internals.
	src := make([]float32, width*height)
	src[(height/2)*width+width/2] = 1
	region := metal.MTLRegion{Size: metal.MTLSize{Width: width, Height: height, Depth: 1}}
	texture.ReplaceRegionMipmapLevelWithBytesBytesPerRow(region, 0, unsafe.Pointer(&src[0]), width*4)
	runtime.KeepAlive(src)

	buf := queue.CommandBuffer()
	if buf.GetID() == 0 {
		fmt.Println("child: could not create a command buffer")
		os.Exit(1)
	}
	fmt.Printf("%s device=%s kernel=%dx%d\n", markerSetup, device.Name(), kernel, kernel)
	os.Stdout.Sync()

	var allocated objc.ID
	allocator := func(k *mps.MPSKernel, cb metal.MTLCommandBuffer, source metal.MTLTexture) metal.MTLTexture {
		// Assertion 1: the allocator ran at all.
		fmt.Println(markerAllocator)
		os.Stdout.Sync()

		// Assertion 2: the three arguments are live objects. Nil-checking
		// proves nothing here — a stale or garbage pointer is non-nil — so
		// each one is messaged and its class and description printed.
		var kernelID objc.ID
		if k != nil {
			kernelID = k.ID
		}
		ok := true
		for _, a := range []struct {
			name string
			id   objc.ID
		}{
			{"filter", kernelID},
			{"commandBuffer", cb.GetID()},
			{"sourceTexture", source.GetID()},
		} {
			if a.id == 0 {
				fmt.Printf("child: allocator argument %s is nil\n", a.name)
				ok = false
				continue
			}
			fmt.Printf("child: allocator argument %s: class=%s description=%s\n",
				a.name, describe(objc.Send[objc.ID](a.id, objc.Sel("class"))), describe(a.id))
		}
		if ok {
			fmt.Println(markerArgs)
		}
		os.Stdout.Sync()

		// Assertion 3, now expressible: RULE 1 LANDED (appledocs f007343a289),
		// so the handler type carries its `metal.MTLTexture` return and the
		// allocator can hand a texture back. Before that commit this block
		// could not be written at all and control fell off the end of the
		// function, leaving MPS to message whatever was in x0.
		replacement, err := newTexture(device)
		if err != nil {
			fmt.Printf("child: allocator could not create a replacement texture: %v\n", err)
			os.Stdout.Sync()
			return metal.MTLTextureObjectFromID(0)
		}
		allocated = replacement.GetID()
		// NS_RETURNS_RETAINED: MPS takes ownership of a +1 texture and
		// releases the one it replaces.
		objc.Send[objc.ID](allocated, objc.Sel("retain"))
		fmt.Printf("%s %#x\n", markerAllocated, uintptr(allocated))
		os.Stdout.Sync()
		return replacement
	}

	// DEFECT 2 is right here: Apple's second parameter is
	// `__strong id<MTLTexture> *texture`, the address of a texture variable.
	// The shipped wrapper takes the texture by value, so this call hands MPS
	// an object pointer to write through. It faults before `allocator` runs.
	got := filter.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(buf, texture, allocator)
	fmt.Printf("%s %v\n", markerReturned, got)
	os.Stdout.Sync()

	buf.Commit()
	buf.WaitUntilCompleted()

	// Assertions 3 and 4 would be scored here. They are not scored, and the
	// parent knows it: `texture` is still this program's original object,
	// because the shipped signature gave MPS nowhere to write a replacement
	// and gave the allocator no way to produce one. Reading pixels out of it
	// would be a check of the in-place path, not of the fallback path, and
	// printing "ok" for it would be a lie.
	fmt.Println("child: reached the end of the encode path")
	os.Exit(0)
}

// describe returns an object's -description as a Go string. It is used to
// prove that a pointer handed to the allocator is a live object rather than
// merely non-nil.
func describe(id objc.ID) string {
	if id == 0 {
		return "<nil>"
	}
	s := objc.Send[objc.ID](id, objc.Sel("description"))
	if s == 0 {
		return "<no description>"
	}
	return foundation.NSStringFromID(s).UTF8String()
}

// newTexture returns a shared single-channel float texture MPS can read and
// write.
func newTexture(device metal.MTLDeviceObject) (metal.MTLTexture, error) {
	desc := metal.GetMTLTextureDescriptorClass().Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(
		metal.MTLPixelFormatR32Float, width, height, false)
	if desc.GetID() == 0 {
		return nil, fmt.Errorf("could not create a texture descriptor")
	}
	desc.SetUsage(metal.MTLTextureUsageShaderRead | metal.MTLTextureUsageShaderWrite)
	desc.SetStorageMode(metal.MTLStorageModeShared)
	tex := device.NewTextureWithDescriptor(desc)
	if tex.GetID() == 0 {
		return nil, fmt.Errorf("could not create a %dx%d texture", width, height)
	}
	return tex, nil
}
