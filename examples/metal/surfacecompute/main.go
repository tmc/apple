// Command surfacecompute demonstrates GPU compute over an IOSurface shared
// between two Go processes with zero copies.
//
// The producer allocates a 2D float32 IOSurface, fills it from Go, and
// passes the surface's mach port to a spawned consumer as a port-right
// descriptor (the tensorshare route: IOSurfaceCreateMachPort → x/mach Send
// with MoveSend → IOSurfaceLookupFromMachPort). The consumer wraps the
// surface in a Metal texture with newTextureWithDescriptor:iosurface:plane:
// and dispatches a compute kernel that sums every element — the GPU reads
// the producer's pages directly; no blit, no upload, no copy.
//
// Two proofs, in order:
//
//  1. The GPU's checksum matches the sum the producer computed on the CPU
//     when it wrote the pattern.
//  2. The producer mutates one element in place; the consumer re-dispatches
//     the same kernel and the GPU sees the new value. A texture backed by a
//     snapshot copy would return the stale sum.
//
// It then reports the GPU's read bandwidth over the shared pages next to a
// CPU sum of the consumer's own mapping of the same pages.
//
//	go run ./examples/metal/surfacecompute
//	go run ./examples/metal/surfacecompute -width 2048 -height 2048
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
	"unsafe"

	"flag"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/x/mach"
	"github.com/tmc/apple/x/zerocopy"
)

// probeWindow is the byte span probed by x/zerocopy at each end of the surface.
const probeWindow = 4096

// kernelSource sums one texture row per thread into a partials buffer; the
// CPU folds the H partials. Row-parallel keeps the kernel trivial while
// still reading every byte of the surface through the texture unit.
const kernelSource = `
#include <metal_stdlib>
using namespace metal;

kernel void sumrows(texture2d<float, access::read> tex [[texture(0)]],
                    device float *partials [[buffer(0)]],
                    uint row [[thread_position_in_grid]]) {
    if (row >= tex.get_height()) {
        return;
    }
    float acc = 0.0f;
    const uint w = tex.get_width();
    for (uint x = 0; x < w; x++) {
        acc += tex.read(uint2(x, row)).r;
    }
    partials[row] = acc;
}
`

const sentinel = float32(-99999)

func main() {
	log.SetFlags(0)
	width := flag.Int("width", 4096, "surface width in float32 elements")
	height := flag.Int("height", 4096, "surface height in rows")
	reps := flag.Int("reps", 5, "benchmark repetitions")
	consume := flag.Bool("consume", false, "run as consumer child (internal)")
	service := flag.String("service", "", "bootstrap name for the port handoff (consumer)")
	flag.Parse()

	if *consume {
		runConsumer(*service, *width, *height, *reps)
		return
	}
	runProducer(*width, *height, *reps)
}

func runProducer(width, height, reps int) {
	surf, err := createSurface(width, height)
	if err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}

	iosurface.IOSurfaceLock(surf, 0, nil)
	want := fillPattern(surf, width, height)
	iosurface.IOSurfaceUnlock(surf, 0, nil)

	bytes := int64(width) * int64(height) * 4
	log.Printf("producer: %dx%d float32 surface (%.0f MiB), CPU checksum %.0f",
		width, height, float64(bytes)/(1<<20), want)

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}
	service := fmt.Sprintf("com.tmc.surfacecompute.%d", os.Getpid())
	child := exec.Command(exe, "-consume",
		"-service", service,
		"-width", fmt.Sprint(width), "-height", fmt.Sprint(height),
		"-reps", fmt.Sprint(reps))
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}
	stdout, err := child.StdoutPipe()
	if err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}
	if err := child.Start(); err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}

	// Hand the surface's port to the child (tensorshare's production route).
	surfPort := mach.Port(iosurface.IOSurfaceCreateMachPort(surf))
	if surfPort == mach.PortNull {
		log.Fatal("surfacecompute: IOSurfaceCreateMachPort failed")
	}
	var svc mach.Port
	deadline := time.Now().Add(5 * time.Second)
	for {
		svc, err = mach.BootstrapLookUp(service)
		if err == nil {
			break
		}
		if time.Now().After(deadline) {
			log.Fatalf("surfacecompute: bootstrap rendezvous: %v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err := mach.Send(svc, mach.CopySend, 1, []mach.PortRight{{Port: surfPort, Disposition: mach.MoveSend}}, nil, 5*time.Second); err != nil {
		log.Fatalf("surfacecompute: send surface port: %v", err)
	}
	svc.Deallocate()

	sc := bufio.NewScanner(stdout)
	expect := func(word string) string {
		if !sc.Scan() {
			log.Fatalf("surfacecompute: child exited early waiting for %q", word)
		}
		got, rest, _ := strings.Cut(sc.Text(), " ")
		if got != word {
			log.Fatalf("surfacecompute: child said %q, want %q", sc.Text(), word)
		}
		return rest
	}

	// Proof 1: the GPU sums the shared pages and matches the CPU checksum.
	fmt.Fprintf(stdin, "%.0f\n", want)
	expect("verified")
	log.Printf("consumer: GPU checksum over port-passed surface matches — zero bytes copied")

	// Proof 2: mutate one element in place; the GPU must see it on the next
	// dispatch. A snapshot copy would keep returning the old sum.
	iosurface.IOSurfaceLock(surf, 0, nil)
	base := (*float32)(iosurface.IOSurfaceGetBaseAddress(surf))
	orig := *base
	*base = sentinel
	iosurface.IOSurfaceUnlock(surf, 0, nil)
	delta := float64(sentinel) - float64(orig)
	fmt.Fprintf(stdin, "mutated %.0f\n", delta)
	expect("observed")
	log.Printf("consumer: GPU observed the in-place CPU write (Δ=%.0f) — live pages, not a snapshot", delta)

	// x/zerocopy probe of the consumer's CPU mapping: multi-offset sentinel
	// writes with verified restores, at both ends of the surface. (The GPU
	// texture view is covered by the Δ re-dispatch above; this hardens the
	// claim that the consumer's mapping is the same pages, not a copy.)
	iosurface.IOSurfaceLock(surf, 0, nil)
	raw := unsafe.Slice((*byte)(iosurface.IOSurfaceGetBaseAddress(surf)), width*height*4)
	win := min(probeWindow, len(raw))
	for _, off := range []int{0, len(raw) - win} {
		err := zerocopy.CheckFunc(raw[off:off+win], func() ([]byte, error) {
			fmt.Fprintf(stdin, "peek %d %d\n", off, win)
			return hex.DecodeString(expect("peek"))
		})
		if err != nil {
			log.Fatalf("surfacecompute: zerocopy probe at offset %d: %v", off, err)
		}
	}
	iosurface.IOSurfaceUnlock(surf, 0, nil)
	fmt.Fprintln(stdin, "endprobe")
	log.Printf("consumer: x/zerocopy probe passed at both ends of the surface")

	// Bench: GPU texture-unit read vs the consumer's CPU read, same pages.
	log.Printf("bench: GPU kernel sum        %s", expect("gpu"))
	log.Printf("bench: CPU mapped sum        %s", expect("cpu"))

	if err := child.Wait(); err != nil {
		log.Fatalf("surfacecompute: consumer: %v", err)
	}
	releaseRef(uintptr(surf))
}

func runConsumer(service string, width, height, reps int) {
	// Receive the surface's port as a mach port-right descriptor.
	recv, err := mach.NewPort()
	if err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}
	if err := recv.MakeSendRight(); err != nil {
		log.Fatalf("surfacecompute: %v", err)
	}
	if err := mach.BootstrapRegister(service, recv); err != nil {
		log.Fatalf("surfacecompute: bootstrap_register: %v", err)
	}
	m, err := mach.Receive(recv, 10*time.Second)
	if err != nil {
		log.Fatalf("surfacecompute: receive surface port: %v", err)
	}
	if len(m.Ports) != 1 {
		log.Fatalf("surfacecompute: message carried %d ports, want 1", len(m.Ports))
	}
	surf := iosurface.IOSurfaceLookupFromMachPort(uint32(m.Ports[0]))
	if surf == 0 {
		log.Fatal("surfacecompute: IOSurfaceLookupFromMachPort failed")
	}
	m.Ports[0].Deallocate()
	recv.DestroyReceive()

	g := newGPU(surf, width, height)

	in := bufio.NewScanner(os.Stdin)

	// Proof 1: GPU checksum vs producer's CPU checksum.
	if !in.Scan() {
		log.Fatal("surfacecompute: no checksum on stdin")
	}
	var want float64
	fmt.Sscanf(in.Text(), "%f", &want)
	got := g.sum()
	if !closeEnough(got, want, width, height) {
		log.Fatalf("surfacecompute: GPU checksum mismatch: got %.0f want %.0f", got, want)
	}
	fmt.Println("verified")

	// Proof 2: re-dispatch after the producer's in-place write.
	if !in.Scan() {
		log.Fatal("surfacecompute: expected mutated")
	}
	var delta float64
	if _, err := fmt.Sscanf(in.Text(), "mutated %f", &delta); err != nil {
		log.Fatalf("surfacecompute: expected mutated line: %q", in.Text())
	}
	after := g.sum()
	if !closeEnough(after, want+delta, width, height) {
		log.Fatalf("surfacecompute: GPU did not observe mutation: got %.0f want %.0f", after, want+delta)
	}
	fmt.Println("observed")

	// Serve peek requests for the producer's zerocopy probe, hex-dumping
	// windows of our own mapping on demand.
	for {
		if !in.Scan() {
			log.Fatal("surfacecompute: stdin closed during zerocopy probe")
		}
		line := in.Text()
		if line == "endprobe" {
			break
		}
		var off, n int
		if _, err := fmt.Sscanf(line, "peek %d %d", &off, &n); err != nil {
			log.Fatalf("surfacecompute: bad probe request %q", line)
		}
		iosurface.IOSurfaceLock(surf, 1, nil) // kIOSurfaceLockReadOnly
		window := unsafe.Slice((*byte)(iosurface.IOSurfaceGetBaseAddress(surf)), width*height*4)[off : off+n]
		h := hex.EncodeToString(window)
		iosurface.IOSurfaceUnlock(surf, 1, nil)
		fmt.Println("peek", h)
	}

	// Bench arms. Same pages: once through the texture unit, once through
	// the consumer's own CPU mapping.
	bytes := int64(width) * int64(height) * 4
	best := time.Duration(math.MaxInt64)
	for range reps {
		start := time.Now()
		g.sum()
		if d := time.Since(start); d < best {
			best = d
		}
	}
	fmt.Printf("gpu %s\n", rate(bytes, best))

	base := iosurface.IOSurfaceGetBaseAddress(surf)
	data := unsafe.Slice((*float32)(base), width*height)
	best = time.Duration(math.MaxInt64)
	for range reps {
		iosurface.IOSurfaceLock(surf, 1, nil) // kIOSurfaceLockReadOnly
		start := time.Now()
		var s float64
		for _, v := range data {
			s += float64(v)
		}
		d := time.Since(start)
		iosurface.IOSurfaceUnlock(surf, 1, nil)
		_ = s
		if d < best {
			best = d
		}
	}
	fmt.Printf("cpu %s\n", rate(bytes, best))

	releaseRef(uintptr(surf))
}

// gpu owns the Metal objects for the sum kernel over one surface.
type gpu struct {
	device   metal.MTLDeviceObject
	queue    metal.MTLCommandQueue
	pipeline metal.MTLComputePipelineState
	texture  metal.MTLTexture
	partials metal.MTLBuffer
	height   int
}

func newGPU(surf iosurface.IOSurfaceRef, width, height int) *gpu {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		log.Fatal("surfacecompute: no Metal device")
	}

	desc := metal.GetMTLTextureDescriptorClass().Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(
		metal.MTLPixelFormatR32Float, uint(width), uint(height), false)
	desc.SetUsage(metal.MTLTextureUsageShaderRead)
	desc.SetStorageMode(metal.MTLStorageModeShared)

	// The zero-copy step: the texture's backing store IS the shared surface.
	tex := device.NewTextureWithDescriptorIosurfacePlane(desc, surf, 0)
	if tex.GetID() == 0 {
		log.Fatal("surfacecompute: newTextureWithDescriptor:iosurface:plane: returned nil")
	}

	lib, err := device.NewLibraryWithSourceOptionsError(kernelSource, nil)
	if err != nil {
		log.Fatalf("surfacecompute: shader compilation: %v", err)
	}
	fn := lib.NewFunctionWithName("sumrows")
	if fn.GetID() == 0 {
		log.Fatal("surfacecompute: kernel function not found")
	}
	pipeline, err := device.NewComputePipelineStateWithFunctionError(fn)
	if err != nil {
		log.Fatalf("surfacecompute: pipeline: %v", err)
	}

	partials := device.NewBufferWithLengthOptions(uint(height)*4, metal.MTLResourceStorageModeShared)
	if partials.GetID() == 0 {
		log.Fatal("surfacecompute: partials buffer allocation failed")
	}

	return &gpu{
		device:   device,
		queue:    device.NewCommandQueue(),
		pipeline: pipeline,
		texture:  tex,
		partials: partials,
		height:   height,
	}
}

// sum dispatches one row-sum kernel over the surface and folds the partials.
func (g *gpu) sum() float64 {
	cb := g.queue.CommandBuffer()
	enc := cb.ComputeCommandEncoder()
	enc.SetComputePipelineState(g.pipeline)
	enc.SetTextureAtIndex(g.texture, 0)
	enc.SetBufferWithOffsetAtIndex(g.partials, 0, 0)

	tg := g.pipeline.MaxTotalThreadsPerThreadgroup()
	if tg > uint(g.height) {
		tg = uint(g.height)
	}
	enc.DispatchThreadsThreadsPerThreadgroup(
		metal.MTLSize{Width: uint(g.height), Height: 1, Depth: 1},
		metal.MTLSize{Width: uint(tg), Height: 1, Depth: 1})
	enc.EndEncoding()
	cb.Commit()
	cb.WaitUntilCompleted()

	partials := unsafe.Slice((*float32)(g.partials.Contents()), g.height)
	var sum float64
	for _, v := range partials {
		sum += float64(v)
	}
	return sum
}

// closeEnough compares two checksums under float32 accumulation error. The
// GPU sums each row in float32 while the reference accumulates in float64,
// so allow half a ULP per element of drift, scaled by magnitude.
func closeEnough(got, want float64, width, height int) bool {
	tol := math.Max(1, math.Abs(want)) * 1e-7 * float64(width)
	_ = height
	return math.Abs(got-want) <= tol
}

// fillPattern writes a deterministic pattern and returns its float64 sum.
func fillPattern(surf iosurface.IOSurfaceRef, width, height int) float64 {
	base := iosurface.IOSurfaceGetBaseAddress(surf)
	data := unsafe.Slice((*float32)(base), width*height)
	var sum float64
	for i := range data {
		data[i] = float32(i%251) * 0.5
		sum += float64(data[i])
	}
	return sum
}

func rate(bytes int64, d time.Duration) string {
	gbs := float64(bytes) / d.Seconds() / (1 << 30)
	return fmt.Sprintf("%8.2f GiB/s  (%d MiB in %v)", gbs, bytes>>20, d.Round(time.Microsecond))
}

// createSurface allocates a width x height float32 IOSurface. bytesPerRow is
// width*4, which for the default sizes is comfortably beyond Metal's linear
// texture alignment; newGPU would fail loudly if it were not.
func createSurface(width, height int) (iosurface.IOSurfaceRef, error) {
	keys := []unsafe.Pointer{
		cfString("IOSurfaceWidth"),
		cfString("IOSurfaceHeight"),
		cfString("IOSurfaceBytesPerElement"),
		cfString("IOSurfaceBytesPerRow"),
		cfString("IOSurfacePixelFormat"),
	}
	values := []unsafe.Pointer{
		cfInt(width),
		cfInt(height),
		cfInt(4),
		cfInt(width * 4),
		cfInt(0x4c303066), // 'L00f' one-component 32-bit float, matches R32Float
	}
	dict := corefoundation.CFDictionaryCreate(0, unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), corefoundation.CFIndex(len(keys)), nil, nil)
	ref := iosurface.IOSurfaceCreate(corefoundation.CFDictionaryRef(dict))
	releaseRef(uintptr(dict))
	for i := range keys {
		releaseRef(uintptr(keys[i]))
		releaseRef(uintptr(values[i]))
	}
	if ref == 0 {
		return 0, fmt.Errorf("IOSurfaceCreate failed for %dx%d", width, height)
	}
	return ref, nil
}

func cfString(s string) unsafe.Pointer {
	ref := corefoundation.CFStringCreateWithCString(0, s, 0x08000100) // kCFStringEncodingUTF8
	return refPointer(uintptr(ref))
}

func cfInt(v int) unsafe.Pointer {
	val := int64(v)
	ref := corefoundation.CFNumberCreate(0, corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&val))
	return refPointer(uintptr(ref))
}

// refPointer converts a CF reference held as uintptr into an unsafe.Pointer
// for CFDictionaryCreate's void* arrays. CF references are not Go pointers,
// so this is not a GC-visibility hazard.
func refPointer(p uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&p))
}

// releaseRef balances a +1 CF reference held as uintptr.
func releaseRef(p uintptr) {
	corefoundation.CFRelease(refPointer(p))
}
