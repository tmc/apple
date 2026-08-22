// Command iogpucensus reports the Objective-C class identity of the objects the
// public Metal API vends.
//
// design/iogpu-spec.md proposes using private/iogpu — the IOGPUMetal* classes
// behind Metal — as a measurement instrument. Every workstream in that spec
// assumes a public MTLDevice, MTLBuffer, or MTLTexture can be treated as its
// corresponding IOGPUMetal* class. That assumption is load-bearing and
// unverified: a census that reports selectors responding on the wrong object
// certifies nothing.
//
// This command answers the public-object class-identity half of the census and
// performs a read-only selector-presence census. It does not invoke private
// selectors, so it remains useful before private/iogpu compiles.
// It prints the full superclass chain of each object the public API returns and
// says whether the IOGPUMetal* class the spec expects appears anywhere in it.
//
// A "no" here invalidates the corresponding workstream before anyone builds on
// it, which is the point of running it early.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/iogpu"
)

// expectation pairs an object obtained from the public API with the private
// class design/iogpu-spec.md assumes it to be.
type expectation struct {
	what  string // what the public API call returns
	want  string // the IOGPUMetal* class the spec assumes
	group string // the capability group in the spec that depends on it
	id    objc.ID
}

type selectorExpectation struct {
	what, selector string
	id             objc.ID
	want           bool
}

func main() {
	log := flag.Bool("v", false, "also print the class of every intermediate object")
	probeSurface := flag.Bool("probe-ane-surface", false, "invoke the read-only _aneIOSurface probe in this process")
	flag.Parse()
	_ = log

	// Metal objects must be created and inspected from one thread; nothing here
	// submits work, but the device is process-wide state and pinning keeps the
	// class lookups on a single thread for reproducibility.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Printf("iogpucensus: class identity of public Metal objects\n")
	fmt.Printf("host: %s\n\n", hostVersion())

	dev := metal.MTLCreateSystemDefaultDevice()
	if dev.ID == 0 {
		fmt.Fprintln(os.Stderr, "iogpucensus: no Metal device; nothing to census")
		os.Exit(1)
	}

	const bufBytes = 4 << 10
	buf := dev.NewBufferWithLengthOptions(bufBytes, metal.MTLResourceStorageModeShared)

	exps := []expectation{
		{what: "MTLCreateSystemDefaultDevice()", want: "IOGPUMetalDevice", group: "C, E (accounting, submit path)", id: dev.ID},
		{what: "device.newBufferWithLength:options:", want: "IOGPUMetalBuffer", group: "A, B (ANE eligibility, aliasing)", id: buf.GetID()},
	}
	selectors := []selectorExpectation{
		{what: "MTLBuffer private ANE surface selector", selector: "_aneIOSurface", id: buf.GetID(), want: true},
		{what: "deliberately absent selector control", selector: "iogpucensus_selector_absent:", id: buf.GetID(), want: false},
	}

	var missing int
	for _, e := range exps {
		if !report(e) {
			missing++
		}
	}

	var selectorMismatches int
	for _, e := range selectors {
		if !reportSelector(e) {
			selectorMismatches++
		}
	}
	if *probeSurface {
		reportANESurface(buf.GetID())
	}

	fmt.Println()
	switch {
	case missing == 0 && selectorMismatches == 0:
		fmt.Println("VERDICT: every expected IOGPUMetal* class appears in its object's chain.")
		fmt.Println("The class-identity assumption in design/iogpu-spec.md holds on this host.")
		fmt.Println("VERDICT: selector presence controls passed; selector semantics and returned-surface aliasing remain unmeasured.")
	default:
		fmt.Printf("VERDICT: %d class and %d selector expectations FAILED.\n", missing, selectorMismatches)
		fmt.Println("The workstreams named beside each failure rest on a class relationship")
		fmt.Println("that does not exist on this host. Fix the spec before building on it.")
	}
	if missing > 0 || selectorMismatches > 0 {
		os.Exit(2)
	}
}

func reportANESurface(id objc.ID) {
	// This is deliberately opt-in: the private selector is a read-only probe,
	// but it still belongs in a disposable process and is never a product path.
	b := iogpu.IOGPUMetalBufferFromID(id)
	surface, err := b.AneIOSurface()
	if err != nil {
		fmt.Printf("ane surface: refused: %v\n", err)
		return
	}
	if surface == 0 {
		fmt.Println("ane surface: returned null")
		return
	}
	fmt.Printf("ane surface: returned non-null id=%d\n", iosurface.IOSurfaceGetID(surface))
}

func reportSelector(e selectorExpectation) bool {
	got := objc.RespondsToSelector(e.id, objc.Sel(e.selector))
	ok := got == e.want
	fmt.Printf("selector %s (%s): got=%t want=%t", e.selector, e.what, got, e.want)
	if ok {
		fmt.Println(" — PASS")
	} else {
		fmt.Println(" — FAIL")
	}
	return ok
}

// report prints the superclass chain for e and returns whether e.want is in it.
func report(e expectation) bool {
	fmt.Printf("%s\n", e.what)
	fmt.Printf("  expects: %s   (spec group %s)\n", e.want, e.group)

	chain := superclassChain(e.id)
	if len(chain) == 0 {
		fmt.Printf("  chain:   <could not read class>\n")
		fmt.Printf("  RESULT:  FAIL — no class for object\n\n")
		return false
	}

	fmt.Printf("  chain:   ")
	for i, name := range chain {
		if i > 0 {
			fmt.Printf(" -> ")
		}
		fmt.Printf("%s", name)
	}
	fmt.Println()

	for _, name := range chain {
		if name == e.want {
			fmt.Printf("  RESULT:  ok — %s is in the chain\n\n", e.want)
			return true
		}
	}
	fmt.Printf("  RESULT:  FAIL — %s is NOT in the chain; the concrete class is %s\n\n", e.want, chain[0])
	return false
}

// superclassChain returns the class of id followed by each of its superclasses,
// most derived first. It returns nil if the object has no readable class.
func superclassChain(id objc.ID) []string {
	cls := objectivec.Object_getClass(objectivec.Object{ID: id})
	var names []string
	// The chain is short (single-digit); the bound is a guard against a cycle in
	// a runtime that has been tampered with, not an expected case.
	for i := 0; cls != 0 && i < 32; i++ {
		names = append(names, className(cls))
		cls = objectivec.Class_getSuperclass(cls)
	}
	return names
}

func className(cls objc.Class) string {
	p := objectivec.Class_getName(cls)
	if p == nil {
		return "<unnamed>"
	}
	return goString(p)
}

// goString converts a NUL-terminated C string to a Go string.
func goString(p *byte) string {
	if p == nil {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Add(unsafe.Pointer(p), n)) != 0 {
		n++
	}
	return string(unsafe.Slice(p, n))
}

// hostVersion reads the OS version from the running system rather than from a
// written-down constant: an environment note in this tree was wrong once
// because it was typed from memory.
func hostVersion() string {
	product, err := syscall.Sysctl("kern.osproductversion")
	if err != nil {
		product = "?"
	}
	release, err := syscall.Sysctl("kern.osrelease")
	if err != nil {
		release = "?"
	}
	return fmt.Sprintf("macOS %s (Darwin %s) %s", product, release, runtime.GOARCH)
}
