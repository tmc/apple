// Command modelservice demonstrates one resident copy of model weights
// serving many client processes — the shape of Apple's own on-device
// model daemon, built from this repo's mach and iosurface bindings.
//
// The service loads its "weights" (a deterministic float32 tensor) into
// an IOSurface once, registers a bootstrap port, and serves two verbs
// over raw mach messages:
//
//   - generate: the client sends a prompt and a reply port; the service
//     streams tokens back one message at a time. Every token is derived
//     by reading a strided walk of the weight pages, so the weights are
//     load-bearing, not decoration.
//   - attach: the service hands the client the weights surface's mach
//     port. The client maps the same physical pages and verifies the
//     checksum — gigabytes of weights cross process boundaries with
//     zero bytes copied.
//
// Every process prints its phys_footprint from task_info(TASK_VM_INFO),
// the kernel's own ledger, so the "one copy" claim carries its evidence.
//
// The default mode orchestrates the whole demo: it spawns the service
// and three clients (two generating concurrently, one attaching the
// weights), then SIGKILLs a fourth client mid-stream and shows the
// service shrugging it off.
//
//	go run ./examples/mach/modelservice              # orchestrated demo
//	go run ./examples/mach/modelservice -mib 4096    # 4 GiB of weights
//
// Manual mode, separate terminals:
//
//	go run ./examples/mach/modelservice -serve -service com.tmc.modelservice
//	go run ./examples/mach/modelservice -client -service com.tmc.modelservice -prompt "hello"
//	go run ./examples/mach/modelservice -attach -service com.tmc.modelservice
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/x/mach"
)

// Message IDs on the wire. Requests carry the client's reply port as a
// port-right descriptor; replies go to that port.
const (
	msgGenerate = 1 // body: prompt bytes
	msgToken    = 2 // body: one token
	msgDone     = 3 // end of stream
	msgAttach   = 4 // body: empty
	msgWeights  = 5 // port: weights surface; body: float64 checksum
	msgError    = 6 // body: error text
)

const (
	lockReadOnly = 1 // kIOSurfaceLockReadOnly

	// task_info flavor and field offset, verified against
	// <mach/task_info.h>: TASK_VM_INFO=22, sizeof(task_vm_info_data_t)=372
	// (count 93), offsetof(phys_footprint)=144.
	taskVMInfo       = 22
	physFootprintOff = 144
	taskVMInfoCount  = 93
)

func main() {
	log.SetFlags(0)
	serve := flag.Bool("serve", false, "run as the model service")
	client := flag.Bool("client", false, "run as a generate client")
	attach := flag.Bool("attach", false, "run as a weights-attach client")
	service := flag.String("service", "", "bootstrap service name")
	prompt := flag.String("prompt", "the weather in mach kernel land", "prompt for -client")
	label := flag.String("label", "", "log prefix for this process")
	mib := flag.Int("mib", 1024, "weights size in MiB")
	tokens := flag.Int("tokens", 12, "tokens per generation")
	flag.Parse()

	if *label != "" {
		log.SetPrefix("[" + *label + "] ")
	}

	switch {
	case *serve:
		runService(*service, *mib, *tokens)
	case *client:
		runClient(*service, *prompt)
	case *attach:
		runAttach(*service, *mib)
	default:
		runDemo(*mib, *tokens)
	}
}

// physFootprint returns the kernel's ledgered footprint for this process,
// the same number Activity Monitor's "Memory" column and jetsam use.
func physFootprint() uint64 {
	var info kernel.Task_vm_info_data_t
	count := kernel.Mach_msg_type_number_t(taskVMInfoCount)
	kr := kernel.Task_info(kernel.Mach_task_self(), taskVMInfo,
		kernel.Task_info_t(unsafe.Pointer(&info)), &count)
	if kr != kernel.KERN_SUCCESS {
		return 0
	}
	raw := (*[4 * taskVMInfoCount]byte)(unsafe.Pointer(&info))
	return binary.LittleEndian.Uint64(raw[physFootprintOff : physFootprintOff+8])
}

func mibOf(b uint64) string { return fmt.Sprintf("%d MiB", b>>20) }

// ---- service ----

type server struct {
	surf     iosurface.IOSurfaceRef
	weights  []float32
	checksum float64
	tokens   int

	mu     sync.Mutex
	served int
}

func runService(service string, mib, tokens int) {
	if service == "" {
		log.Fatal("modelservice: -serve requires -service")
	}
	floats := mib << 20 / 4
	surf, err := createSurface(floats * 4)
	if err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	if rc := iosurface.IOSurfaceLock(surf, 0, nil); rc != 0 {
		log.Fatalf("modelservice: weights lock failed rc=%d", rc)
	}
	weights := surfaceFloats(surf, floats)
	sum := pattern(weights)
	iosurface.IOSurfaceUnlock(surf, 0, nil)
	log.Printf("service: weights loaded: %d MiB, checksum %.0f, footprint %s",
		mib, sum, mibOf(physFootprint()))

	recv, err := mach.NewPort()
	if err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	if err := recv.MakeSendRight(); err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	if err := mach.BootstrapRegister(service, recv); err != nil {
		log.Fatalf("modelservice: bootstrap_register %q: %v", service, err)
	}
	log.Printf("service: listening on bootstrap %q (pid %d)", service, os.Getpid())

	s := &server{surf: surf, weights: weights, checksum: sum, tokens: tokens}
	for {
		m, err := mach.Receive(recv, 0)
		if err != nil {
			log.Fatalf("modelservice: receive: %v", err)
		}
		if len(m.Ports) != 1 {
			log.Printf("service: request %d carried %d ports, want 1 — dropped", m.Header.ID, len(m.Ports))
			continue
		}
		reply := m.Ports[0]
		switch m.Header.ID {
		case msgGenerate:
			prompt := strings.TrimRight(string(m.Body), "\x00")
			go s.generate(reply, prompt)
		case msgAttach:
			go s.attach(reply)
		default:
			mach.Send(reply, mach.CopySend, msgError, nil,
				fmt.Appendf(nil, "unknown verb %d", m.Header.ID), time.Second)
			reply.Deallocate()
		}
	}
}

func (s *server) generate(reply mach.Port, prompt string) {
	defer reply.Deallocate()
	s.mu.Lock()
	s.served++
	n := s.served
	s.mu.Unlock()
	log.Printf("service: generate #%d %q (footprint %s)", n, prompt, mibOf(physFootprint()))

	seed := uint64(14695981039346656037)
	for i := range len(prompt) {
		seed = (seed ^ uint64(prompt[i])) * 1099511628211
	}
	for range s.tokens {
		seed = touchWeights(s.weights, seed)
		word := vocab[seed%uint64(len(vocab))]
		if err := mach.Send(reply, mach.CopySend, msgToken, nil, []byte(word), time.Second); err != nil {
			log.Printf("service: generate #%d: client vanished mid-stream (%v) — stream dropped, service continues", n, err)
			return
		}
		time.Sleep(60 * time.Millisecond)
	}
	if err := mach.Send(reply, mach.CopySend, msgDone, nil, nil, time.Second); err != nil {
		log.Printf("service: generate #%d: done not delivered: %v", n, err)
	}
}

func (s *server) attach(reply mach.Port) {
	defer reply.Deallocate()
	// IOSurfaceCreateMachPort mints a fresh send right per call, so
	// MoveSend hands the whole right to the client and nothing needs
	// balancing here; the service keeps its own surface reference.
	surfPort := mach.Port(iosurface.IOSurfaceCreateMachPort(s.surf))
	if surfPort == mach.PortNull {
		mach.Send(reply, mach.CopySend, msgError, nil, []byte("IOSurfaceCreateMachPort failed"), time.Second)
		return
	}
	body := make([]byte, 8)
	binary.LittleEndian.PutUint64(body, math.Float64bits(s.checksum))
	if err := mach.Send(reply, mach.CopySend, msgWeights,
		[]mach.PortRight{{Port: surfPort, Disposition: mach.MoveSend}}, body, time.Second); err != nil {
		log.Printf("service: attach: client vanished (%v)", err)
		surfPort.Deallocate()
		return
	}
	log.Printf("service: weights surface port handed out (zero bytes copied)")
}

// touchWeights folds a strided read of the weight pages into the seed —
// the stand-in decode step. It exists so generation provably reads the
// weights: unmap them and tokens stop.
func touchWeights(w []float32, seed uint64) uint64 {
	h := seed
	for range 4096 {
		h = h*6364136223846793005 + 1442695040888963407
		h ^= uint64(math.Float32bits(w[h%uint64(len(w))]))
	}
	return h
}

var vocab = []string{
	"the", "mach", "port", "crossed", "kernel", "pages", "shared", "zero",
	"copy", "tokens", "stream", "weights", "resident", "once", "unified",
	"memory", "footprint", "ledger", "surface", "right", "send", "receive",
}

// ---- clients ----

func lookupRetry(service string) mach.Port {
	deadline := time.Now().Add(10 * time.Second)
	for {
		svc, err := mach.BootstrapLookUp(service)
		if err == nil {
			return svc
		}
		if time.Now().After(deadline) {
			log.Fatalf("modelservice: bootstrap rendezvous %q: %v", service, err)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

// request sends verb+body with a fresh reply port and returns that port.
// MakeSend mints the send right the service will answer on directly from
// our receive right.
func request(service string, verb int32, body []byte) mach.Port {
	svc := lookupRetry(service)
	reply, err := mach.NewPort()
	if err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	err = mach.Send(svc, mach.CopySend, verb,
		[]mach.PortRight{{Port: reply, Disposition: mach.MakeSend}}, body, 5*time.Second)
	svc.Deallocate()
	if err != nil {
		log.Fatalf("modelservice: request: %v", err)
	}
	return reply
}

func runClient(service, prompt string) {
	if service == "" {
		log.Fatal("modelservice: -client requires -service")
	}
	reply := request(service, msgGenerate, []byte(prompt))
	defer reply.DestroyReceive()

	var words []string
	for {
		m, err := mach.Receive(reply, 30*time.Second)
		if err != nil {
			log.Fatalf("modelservice: token stream: %v", err)
		}
		switch m.Header.ID {
		case msgToken:
			words = append(words, strings.TrimRight(string(m.Body), "\x00"))
		case msgDone:
			log.Printf("client: %q -> %q (footprint %s)",
				prompt, strings.Join(words, " "), mibOf(physFootprint()))
			return
		case msgError:
			log.Fatalf("modelservice: service error: %s", strings.TrimRight(string(m.Body), "\x00"))
		}
	}
}

func runAttach(service string, mib int) {
	if service == "" {
		log.Fatal("modelservice: -attach requires -service")
	}
	before := physFootprint()
	reply := request(service, msgAttach, nil)
	defer reply.DestroyReceive()

	m, err := mach.Receive(reply, 30*time.Second)
	if err != nil {
		log.Fatalf("modelservice: attach reply: %v", err)
	}
	if m.Header.ID != msgWeights || len(m.Ports) != 1 {
		log.Fatalf("modelservice: attach reply id=%d ports=%d, want weights+1", m.Header.ID, len(m.Ports))
	}
	want := math.Float64frombits(binary.LittleEndian.Uint64(m.Body))
	surf := iosurface.IOSurfaceLookupFromMachPort(uint32(m.Ports[0]))
	if surf == 0 {
		log.Fatal("modelservice: IOSurfaceLookupFromMachPort failed")
	}
	m.Ports[0].Deallocate()

	floats := mib << 20 / 4
	if rc := iosurface.IOSurfaceLock(surf, lockReadOnly, nil); rc != 0 {
		log.Fatalf("modelservice: attach lock failed rc=%d", rc)
	}
	start := time.Now()
	got := sumFloats(surfaceFloats(surf, floats))
	elapsed := time.Since(start)
	iosurface.IOSurfaceUnlock(surf, lockReadOnly, nil)
	if math.Abs(got-want) > 1 {
		log.Fatalf("modelservice: weights checksum mismatch: got %.0f want %.0f", got, want)
	}
	gbs := float64(floats) * 4 / elapsed.Seconds() / (1 << 30)
	log.Printf("attach: verified %d MiB of weights in place at %.2f GiB/s — zero bytes copied; footprint %s -> %s",
		mib, gbs, mibOf(before), mibOf(physFootprint()))
	releaseRef(uintptr(surf))
}

// ---- orchestrated demo ----

func runDemo(mib, tokens int) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	service := fmt.Sprintf("com.tmc.modelservice.%d", os.Getpid())

	svc := exec.Command(exe, "-serve", "-service", service,
		"-mib", fmt.Sprint(mib), "-tokens", fmt.Sprint(tokens), "-label", "service")
	svc.Stderr = os.Stderr
	if err := svc.Start(); err != nil {
		log.Fatalf("modelservice: %v", err)
	}
	defer func() {
		svc.Process.Signal(syscall.SIGTERM)
		svc.Wait()
	}()

	client := func(label string, args ...string) *exec.Cmd {
		c := exec.Command(exe, append(args, "-service", service, "-label", label)...)
		c.Stderr = os.Stderr
		if err := c.Start(); err != nil {
			log.Fatalf("modelservice: %s: %v", label, err)
		}
		return c
	}

	// Two generations and a weights attach, all concurrent against the
	// one resident copy.
	c1 := client("client-1", "-client", "-prompt", "unified memory serves many masters")
	c2 := client("client-2", "-client", "-prompt", "a port right is a capability")
	c3 := client("attach-3", "-attach", "-mib", fmt.Sprint(mib))

	// A fourth client, murdered mid-stream: the service must log the
	// failed send and keep serving everyone else.
	c4 := client("victim-4", "-client", "-prompt", "doomed")
	time.Sleep(350 * time.Millisecond)
	c4.Process.Kill()
	c4.Wait()
	log.Printf("demo: SIGKILLed victim-4 mid-stream")

	for _, c := range []*exec.Cmd{c1, c2, c3} {
		if err := c.Wait(); err != nil {
			log.Fatalf("modelservice: client failed: %v", err)
		}
	}
	log.Printf("demo: all surviving clients done; one weights copy (%d MiB) served %s",
		mib, "2 generations + 1 attach + 1 dead client")
	log.Printf("demo: verify residency yourself while it runs: footprint <service-pid>")
}

// ---- shared helpers (same shapes as examples/iosurface/tensorshare) ----

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

func surfaceFloats(surf iosurface.IOSurfaceRef, n int) []float32 {
	base := iosurface.IOSurfaceGetBaseAddress(surf)
	if base == nil {
		log.Fatal("modelservice: nil base address")
	}
	return unsafe.Slice((*float32)(base), n)
}

// createSurface allocates a byte-addressable IOSurface laid out as a
// single row of 4-byte elements.
func createSurface(size int) (iosurface.IOSurfaceRef, error) {
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
	dict := corefoundation.CFDictionaryCreate(0, unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), corefoundation.CFIndex(len(keys)), nil, nil)
	ref := iosurface.IOSurfaceCreate(corefoundation.CFDictionaryRef(dict))
	releaseRef(uintptr(dict))
	for i := range keys {
		releaseRef(uintptr(keys[i]))
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
