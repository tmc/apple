// Command tensorshare demonstrates zero-copy tensor handoff between two Go
// processes via IOSurface.
//
// The parent process allocates a float32 IOSurface, fills it with a
// deterministic pattern, and spawns this same binary as a consumer child.
// The child finds the surface with IOSurfaceLookup, maps it, and verifies
// the checksum without a single byte being copied between the processes.
// The parent then mutates one element in place and the child observes the
// new value on re-read — the same physical pages, mapped twice.
//
// It then benchmarks the handoff against the conventional route: streaming
// the same buffer through a pipe.
//
//	go run ./examples/iosurface/tensorshare
//	go run ./examples/iosurface/tensorshare -floats $((64<<20)) -hold
//
// With -hold, both processes pause after verification and print the vmmap
// invocations that show the shared region mapped into each address space.
//
// By default the surface's mach port crosses to the child as a port-right
// descriptor in a mach message (IOSurfaceCreateMachPort → x/mach Send with
// MoveSend → IOSurfaceLookupFromMachPort), rendezvousing over a bootstrap
// name — the production route. -transport=global keeps the original
// deprecated IOSurfaceIsGlobal/IOSurfaceLookup path for comparison.
package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/x/mach"
	"github.com/tmc/apple/x/zerocopy"
)

const (
	lockReadOnly = 1 // kIOSurfaceLockReadOnly
	sentinel     = float32(-12345.5)
	probeWindow  = 4096 // bytes probed by x/zerocopy at each end of the surface
)

func main() {
	log.SetFlags(0)
	floats := flag.Int("floats", 16<<20, "float32 elements in the tensor")
	consume := flag.Bool("consume", false, "run as consumer child (internal)")
	pipeArm := flag.Bool("pipe", false, "run as pipe-benchmark child (internal)")
	surfID := flag.Uint("id", 0, "IOSurface ID to look up (consumer, global transport)")
	service := flag.String("service", "", "bootstrap name for port handoff (consumer, port transport)")
	transport := flag.String("transport", "port", "surface handoff: port (mach port descriptor) or global (deprecated IOSurfaceIsGlobal)")
	reps := flag.Int("reps", 5, "benchmark repetitions")
	hold := flag.Bool("hold", false, "pause after verification for vmmap inspection")
	flag.Parse()

	switch {
	case *consume:
		runConsumer(uint32(*surfID), *service, *floats, *reps, *hold)
	case *pipeArm:
		runPipeChild(*floats, *reps)
	default:
		runProducer(*transport, *floats, *reps, *hold)
	}
}

// pattern fills data deterministically and returns its sum.
func pattern(data []float32) float64 {
	var sum float64
	for i := range data {
		data[i] = float32(i%251) * 0.5
		sum += float64(data[i])
	}
	return sum
}

func sumFloats(data []float32) float64 {
	var sum float64
	for _, v := range data {
		sum += float64(v)
	}
	return sum
}

func runProducer(transport string, floats, reps int, hold bool) {
	global := transport == "global"
	surf, err := createSurface(floats*4, global)
	if err != nil {
		log.Fatalf("tensorshare: %v", err)
	}

	if rc := iosurface.IOSurfaceLock(surf, 0, nil); rc != 0 {
		log.Fatalf("tensorshare: producer lock failed rc=%d", rc)
	}
	data := surfaceFloats(surf, floats)
	want := pattern(data)
	iosurface.IOSurfaceUnlock(surf, 0, nil)

	bytes := int64(floats) * 4
	log.Printf("producer: surface %d float32 (%.0f MiB), checksum %.0f, transport=%s",
		floats, float64(bytes)/(1<<20), want, transport)

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("tensorshare: %v", err)
	}
	args := []string{"-consume", "-floats", fmt.Sprint(floats), "-reps", fmt.Sprint(reps)}
	service := fmt.Sprintf("com.tmc.tensorshare.%d", os.Getpid())
	if global {
		args = append(args, "-id", fmt.Sprint(iosurface.IOSurfaceGetID(surf)))
	} else {
		args = append(args, "-service", service)
	}
	if hold {
		args = append(args, "-hold")
	}
	child := exec.Command(exe, args...)
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		log.Fatalf("tensorshare: %v", err)
	}
	stdout, err := child.StdoutPipe()
	if err != nil {
		log.Fatalf("tensorshare: %v", err)
	}
	if err := child.Start(); err != nil {
		log.Fatalf("tensorshare: %v", err)
	}

	if !global {
		// The child registers a receive right under the service name; look
		// it up (retrying while the child starts) and send the surface's
		// port as a port-right descriptor. MoveSend transfers our right, so
		// there is nothing to balance afterward — the right now lives in
		// the message, then in the child.
		surfPort := mach.Port(iosurface.IOSurfaceCreateMachPort(surf))
		if surfPort == mach.PortNull {
			log.Fatal("tensorshare: IOSurfaceCreateMachPort failed")
		}
		var svc mach.Port
		deadline := time.Now().Add(5 * time.Second)
		for {
			svc, err = mach.BootstrapLookUp(service)
			if err == nil {
				break
			}
			if time.Now().After(deadline) {
				log.Fatalf("tensorshare: bootstrap rendezvous: %v", err)
			}
			time.Sleep(10 * time.Millisecond)
		}
		if err := mach.Send(svc, mach.CopySend, 1, []mach.PortRight{{Port: surfPort, Disposition: mach.MoveSend}}, nil, 5*time.Second); err != nil {
			log.Fatalf("tensorshare: send surface port: %v", err)
		}
		svc.Deallocate()
		log.Printf("producer: surface port sent via mach message (bootstrap %q)", service)
	}

	sc := bufio.NewScanner(stdout)
	expect := func(word string) string {
		if !sc.Scan() {
			log.Fatalf("tensorshare: child exited early waiting for %q", word)
		}
		got, rest, _ := strings.Cut(sc.Text(), " ")
		if got != word {
			log.Fatalf("tensorshare: child said %q, want %q", sc.Text(), word)
		}
		return rest
	}

	// Phase 1: child verifies the checksum against its own mapping.
	fmt.Fprintf(stdin, "%.0f\n", want)
	expect("verified")
	how := "port-passed surface (IOSurfaceLookupFromMachPort)"
	if global {
		how = "global-ID lookup (IOSurfaceLookup)"
	}
	log.Printf("consumer: checksum verified via %s — zero bytes sent", how)

	// Phase 2: mutate one element in place; child must observe it.
	if rc := iosurface.IOSurfaceLock(surf, 0, nil); rc != 0 {
		log.Fatalf("tensorshare: producer relock failed rc=%d", rc)
	}
	// data aliases the surface, so capture the original value before the
	// in-place write or the checksum adjustment below cancels itself out.
	orig0 := float64(surfaceFloats(surf, floats)[0])
	surfaceFloats(surf, floats)[0] = sentinel
	iosurface.IOSurfaceUnlock(surf, 0, nil)
	fmt.Fprintln(stdin, "mutated")
	expect("observed")
	log.Printf("consumer: observed in-place mutation (element 0 = %v) — shared pages, not a copy", sentinel)

	// Phase 2b: the x/zerocopy probe. Phase 2 proves one write traveled;
	// this proves the alias at both ends of the surface with restores
	// verified too — a consumer that lazily snapshotted the pages could
	// pass phase 2 by luck and would fail here.
	if rc := iosurface.IOSurfaceLock(surf, 0, nil); rc != 0 {
		log.Fatalf("tensorshare: producer probe lock failed rc=%d", rc)
	}
	raw := unsafe.Slice((*byte)(iosurface.IOSurfaceGetBaseAddress(surf)), floats*4)
	win := min(probeWindow, len(raw))
	for _, off := range []int{0, len(raw) - win} {
		err := zerocopy.CheckFunc(raw[off:off+win], func() ([]byte, error) {
			fmt.Fprintf(stdin, "peek %d %d\n", off, win)
			return hex.DecodeString(expect("peek"))
		})
		if err != nil {
			log.Fatalf("tensorshare: zerocopy probe at offset %d: %v", off, err)
		}
	}
	iosurface.IOSurfaceUnlock(surf, 0, nil)
	fmt.Fprintln(stdin, "endprobe")
	log.Printf("consumer: x/zerocopy probe passed at both ends of the surface — aliased pages, restores observed")

	if hold {
		log.Printf("hold: inspect the shared mapping while both processes are alive:")
		log.Printf("  vmmap %d | grep -i iosurface", os.Getpid())
		log.Printf("  vmmap %d | grep -i iosurface", child.Process.Pid)
		log.Printf("press enter to continue")
		fmt.Scanln()
		fmt.Fprintln(stdin, "release")
	}

	// Phase 3: child benchmarks reading the mapped surface in place.
	sharedLine := expect("shared")
	log.Printf("bench: mapped in-place read   %s", sharedLine)
	if err := child.Wait(); err != nil {
		log.Fatalf("tensorshare: consumer: %v", err)
	}

	// Pipe arm: same bytes, conventional copy through a pipe.
	pipeGBs, err := benchPipe(exe, surf, floats, reps, want-orig0+float64(sentinel))
	if err != nil {
		log.Fatalf("tensorshare: pipe arm: %v", err)
	}
	log.Printf("bench: pipe stream + read     %s", pipeGBs)

	releaseRef(uintptr(surf))
}

func runConsumer(id uint32, service string, floats, reps int, hold bool) {
	var surf iosurface.IOSurfaceRef
	if service != "" {
		// Port transport: publish a receive right, then receive the
		// surface's port as a descriptor in a mach message.
		recv, err := mach.NewPort()
		if err != nil {
			log.Fatalf("tensorshare: %v", err)
		}
		if err := recv.MakeSendRight(); err != nil {
			log.Fatalf("tensorshare: %v", err)
		}
		if err := mach.BootstrapRegister(service, recv); err != nil {
			log.Fatalf("tensorshare: bootstrap_register: %v", err)
		}
		m, err := mach.Receive(recv, 10*time.Second)
		if err != nil {
			log.Fatalf("tensorshare: receive surface port: %v", err)
		}
		if len(m.Ports) != 1 {
			log.Fatalf("tensorshare: message carried %d ports, want 1", len(m.Ports))
		}
		surf = iosurface.IOSurfaceLookupFromMachPort(uint32(m.Ports[0]))
		if surf == 0 {
			log.Fatal("tensorshare: IOSurfaceLookupFromMachPort failed")
		}
		// The lookup retains the surface; the port right itself is ours to
		// balance, and the receive right has served its purpose.
		m.Ports[0].Deallocate()
		recv.DestroyReceive()
	} else {
		surf = iosurface.IOSurfaceLookup(id)
		if surf == 0 {
			log.Fatalf("tensorshare: IOSurfaceLookup(%d) failed — surface not global or producer gone", id)
		}
	}
	in := bufio.NewScanner(os.Stdin)

	// Phase 1: verify checksum against our own mapping of the same pages.
	if !in.Scan() {
		log.Fatal("tensorshare: no checksum on stdin")
	}
	var want float64
	fmt.Sscanf(in.Text(), "%f", &want)
	if rc := iosurface.IOSurfaceLock(surf, lockReadOnly, nil); rc != 0 {
		log.Fatalf("tensorshare: consumer lock failed rc=%d", rc)
	}
	data := surfaceFloats(surf, floats)
	got := sumFloats(data)
	iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)
	if math.Abs(got-want) > 1 {
		log.Fatalf("tensorshare: checksum mismatch: got %.0f want %.0f", got, want)
	}
	fmt.Println("verified")

	// Phase 2: observe the producer's in-place mutation.
	if !in.Scan() || in.Text() != "mutated" {
		log.Fatal("tensorshare: expected mutated")
	}
	iosurface.IOSurfaceLock(surf, lockReadOnly, nil)
	v := surfaceFloats(surf, floats)[0]
	iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)
	if v != sentinel {
		log.Fatalf("tensorshare: mutation not visible: element 0 = %v", v)
	}
	fmt.Println("observed")

	// Phase 2b: serve peek requests for the parent's zerocopy probe,
	// hex-dumping windows of our own mapping on demand.
	for {
		if !in.Scan() {
			log.Fatal("tensorshare: stdin closed during zerocopy probe")
		}
		line := in.Text()
		if line == "endprobe" {
			break
		}
		var off, n int
		if _, err := fmt.Sscanf(line, "peek %d %d", &off, &n); err != nil {
			log.Fatalf("tensorshare: bad probe request %q", line)
		}
		iosurface.IOSurfaceLock(surf, lockReadOnly, nil)
		window := unsafe.Slice((*byte)(iosurface.IOSurfaceGetBaseAddress(surf)), floats*4)[off : off+n]
		h := hex.EncodeToString(window)
		iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)
		fmt.Println("peek", h)
	}

	if hold {
		if !in.Scan() || in.Text() != "release" {
			log.Fatal("tensorshare: expected release")
		}
	}

	// Phase 3: bench in-place reads of the shared mapping.
	best := time.Duration(math.MaxInt64)
	for range reps {
		iosurface.IOSurfaceLock(surf, lockReadOnly, nil)
		start := time.Now()
		sumFloats(surfaceFloats(surf, floats))
		if d := time.Since(start); d < best {
			best = d
		}
		iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)
	}
	fmt.Printf("shared %s\n", rate(int64(floats)*4, best))
	releaseRef(uintptr(surf))
}

// benchPipe streams the surface contents through a pipe to a child that
// reads and sums them, timing the whole conventional path.
func benchPipe(exe string, surf iosurface.IOSurfaceRef, floats, reps int, want float64) (string, error) {
	child := exec.Command(exe, "-pipe", "-floats", fmt.Sprint(floats), "-reps", fmt.Sprint(reps))
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		return "", err
	}
	stdout, err := child.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := child.Start(); err != nil {
		return "", err
	}

	iosurface.IOSurfaceLock(surf, lockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surf)
	src := unsafe.Slice((*byte)(base), floats*4)
	w := bufio.NewWriterSize(stdin, 1<<20)
	for range reps {
		if _, err := w.Write(src); err != nil {
			return "", err
		}
	}
	w.Flush()
	stdin.Close()
	iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)

	sc := bufio.NewScanner(stdout)
	if !sc.Scan() {
		return "", fmt.Errorf("pipe child exited early")
	}
	line := sc.Text()
	if err := child.Wait(); err != nil {
		return "", err
	}
	sumStr, rest, _ := strings.Cut(strings.TrimSpace(line), " ")
	var got float64
	fmt.Sscanf(sumStr, "%f", &got)
	if math.Abs(got-want) > 1 {
		return "", fmt.Errorf("pipe checksum mismatch: got %.0f want %.0f", got, want)
	}
	return rest, nil
}

func runPipeChild(floats, reps int) {
	buf := make([]float32, floats)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&buf[0])), floats*4)
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	best := time.Duration(math.MaxInt64)
	var sum float64
	for range reps {
		start := time.Now()
		if _, err := readFull(r, raw); err != nil {
			log.Fatalf("tensorshare: pipe read: %v", err)
		}
		sum = sumFloats(buf)
		if d := time.Since(start); d < best {
			best = d
		}
	}
	fmt.Printf("%.0f %s\n", sum, rate(int64(floats)*4, best))
}

func readFull(r *bufio.Reader, p []byte) (int, error) {
	n := 0
	for n < len(p) {
		m, err := r.Read(p[n:])
		n += m
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func rate(bytes int64, d time.Duration) string {
	gbs := float64(bytes) / d.Seconds() / (1 << 30)
	return fmt.Sprintf("%8.2f GiB/s  (%d MiB in %v)", gbs, bytes>>20, d.Round(time.Microsecond))
}

func surfaceFloats(surf iosurface.IOSurfaceRef, n int) []float32 {
	base := iosurface.IOSurfaceGetBaseAddress(surf)
	if base == nil {
		log.Fatal("tensorshare: nil base address")
	}
	return unsafe.Slice((*float32)(base), n)
}

// createSurface allocates a byte-addressable IOSurface of at least size
// bytes, laid out as a single row of 4-byte elements. With global set the
// surface is discoverable by ID via the deprecated IOSurfaceIsGlobal; the
// port transport needs no such property.
func createSurface(size int, global bool) (iosurface.IOSurfaceRef, error) {
	width := size / 4
	keys := []unsafe.Pointer{
		cfString("IOSurfaceWidth"),
		cfString("IOSurfaceHeight"),
		cfString("IOSurfaceBytesPerElement"),
		cfString("IOSurfaceBytesPerRow"),
		cfString("IOSurfaceAllocSize"),
		cfString("IOSurfacePixelFormat"),
	}
	values := []unsafe.Pointer{
		cfInt(width),
		cfInt(1),
		cfInt(4),
		cfInt(size),
		cfInt(size),
		cfInt(0),
	}
	if global {
		keys = append(keys, cfString("IOSurfaceIsGlobal"))
		values = append(values, cfBoolTrue())
	}
	dict := corefoundation.CFDictionaryCreate(0, unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), corefoundation.CFIndex(len(keys)), nil, nil)
	ref := iosurface.IOSurfaceCreate(corefoundation.CFDictionaryRef(dict))
	releaseRef(uintptr(dict))
	for i := range keys {
		releaseRef(uintptr(keys[i]))
		// values[6] is kCFBooleanTrue, a constant; releasing a retained
		// reference to it is still balanced since cfBoolTrue retains.
		releaseRef(uintptr(values[i]))
	}
	if ref == 0 {
		return 0, fmt.Errorf("IOSurfaceCreate failed for %d bytes", size)
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

func cfBoolTrue() unsafe.Pointer {
	ref := corefoundation.CFRetain(refPointer(uintptr(corefoundation.KCFBooleanTrue)))
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
