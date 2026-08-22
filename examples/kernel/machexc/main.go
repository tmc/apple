// Command machexc demonstrates a Mach-exception crash reporter: it catches
// EXC_BAD_ACCESS at the task level, prints the faulting thread's full
// register state with a symbolized fault PC, then lets the crash proceed to
// Go's normal signal handling.
//
// Run the A/B:
//
//	go run . -mode=report   # Mach report with registers, then Go's traceback
//	go run . -mode=stock    # stock Go: the traceback alone
//
// The design rule is observe-and-re-raise, never swallow: the handler
// replies KERN_FAILURE to the exception message, so the kernel falls
// through to the host-level handler and converts the exception to the
// SIGSEGV Go's runtime expects. Both arms end in the same Go crash; the
// report arm just knows more first.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/x/mach"
)

// Exception constants from <mach/exception_types.h> and
// <mach/arm/thread_status.h>. ABI-pinned in xnu; not present in Apple's
// documentation set, so not generated.
const (
	excBadAccess           = 1        // EXC_BAD_ACCESS
	excMaskBadAccess       = 1 << 1   // EXC_MASK_BAD_ACCESS
	exceptionDefault       = 1        // EXCEPTION_DEFAULT: message carries thread+task ports, no state
	machExcCodes     int32 = -1 << 31 // MACH_EXCEPTION_CODES: 64-bit codes, msgh_id 2405
	armThreadState64       = 6        // ARM_THREAD_STATE64
	armState64Count        = 68       // ARM_THREAD_STATE64_COUNT (uint32 words)
	kernFailure            = 5        // KERN_FAILURE
	msgIDRaise             = 2405     // mach_exception_raise request
	msgIDRaiseReply        = 2505     // its reply (request + 100)
)

// armThreadState mirrors arm_thread_state64_t from
// <mach/arm/_structs.h> (non-ptrauth layout).
type armThreadState struct {
	X     [29]uint64
	FP    uint64
	LR    uint64
	SP    uint64
	PC    uint64
	CPSR  uint32
	Flags uint32
}

func main() {
	mode := flag.String("mode", "report", "report (Mach handler) or stock (Go signal handling only)")
	flag.Parse()

	if *mode == "report" {
		if err := installHandler(); err != nil {
			fmt.Fprintln(os.Stderr, "install handler:", err)
			os.Exit(1)
		}
		fmt.Println("mach exception handler installed; crashing a foreign callee...")
	} else {
		fmt.Println("stock Go signal handling; crashing a foreign callee...")
	}
	time.Sleep(100 * time.Millisecond)
	crashForeign()
}

// crashForeign faults inside libSystem's memcpy — a foreign frame, so
// stock Go's report has no useful PC to show.
func crashForeign() {
	memcpy, err := purego.Dlsym(purego.RTLD_DEFAULT, "memcpy")
	if err != nil {
		panic(err)
	}
	var fn func(dst, src uintptr, n uint64) uintptr
	purego.RegisterFunc(&fn, memcpy)
	src := make([]byte, 16)
	fn(0xdead, uintptr(unsafe.Pointer(&src[0])), 16)
}

func installHandler() error {
	port, err := mach.NewPort()
	if err != nil {
		return err
	}
	if err := port.MakeSendRight(); err != nil {
		return err
	}
	kr := kernel.Task_set_exception_ports(
		kernel.Mach_task_self(),
		kernel.Exception_mask_t(excMaskBadAccess),
		uint32(port),
		kernel.Exception_behavior_t(machExcCodes|exceptionDefault),
		0, // THREAD_STATE_NONE: EXCEPTION_DEFAULT carries no state in the message
	)
	if kr != 0 {
		return fmt.Errorf("task_set_exception_ports: 0x%x", uint32(kr))
	}

	ready := make(chan struct{})
	go serve(port, ready)
	<-ready
	return nil
}

// serve blocks on the exception port. The faulting thread is suspended by
// the kernel until we reply; replying KERN_FAILURE re-raises rather than
// swallows.
func serve(port mach.Port, ready chan<- struct{}) {
	close(ready)
	for {
		m, err := mach.Receive(port, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "exception receive:", err)
			return
		}
		if m.Header.ID != msgIDRaise || len(m.Ports) < 2 || len(m.Body) < 16 {
			fmt.Fprintf(os.Stderr, "unexpected exception message id=%d ports=%d\n", m.Header.ID, len(m.Ports))
			continue
		}
		thread := m.Ports[0]
		exception := int32(binary.LittleEndian.Uint32(m.Body[8:12]))
		codeCnt := binary.LittleEndian.Uint32(m.Body[12:16])
		var codes []int64
		for i := range int(codeCnt) {
			off := 16 + 8*i
			if off+8 <= len(m.Body) {
				codes = append(codes, int64(binary.LittleEndian.Uint64(m.Body[off:off+8])))
			}
		}

		report(thread, exception, codes)

		// Observe-and-re-raise: KERN_FAILURE tells the kernel this handler
		// did not resolve the exception, so delivery falls through to the
		// host handler and thence to the SIGSEGV Go's runtime expects.
		reply := make([]byte, 12)
		copy(reply, m.Body[:8]) // echo the NDR record
		binary.LittleEndian.PutUint32(reply[8:], kernFailure)
		if err := mach.Send(m.Header.RemotePort, mach.MoveSendOnce, msgIDRaiseReply, nil, reply, 0); err != nil {
			fmt.Fprintln(os.Stderr, "exception reply:", err)
		}
	}
}

func report(thread mach.Port, exception int32, codes []int64) {
	var state armThreadState
	cnt := kernel.Mach_msg_type_number_t(armState64Count)
	kr := kernel.Thread_get_state(uint32(thread),
		kernel.Thread_state_flavor_t(armThreadState64),
		kernel.Thread_state_t(unsafe.Pointer(&state)), &cnt)

	fmt.Println("\n=== mach exception report (the part stock Go cannot print) ===")
	name := "?"
	if exception == excBadAccess {
		name = "EXC_BAD_ACCESS"
	}
	fmt.Printf("exception: %s (%d)  codes: %#x\n", name, exception, codes)
	if len(codes) >= 2 {
		fmt.Printf("fault address: %#x\n", codes[1])
	}
	if kr != 0 {
		fmt.Printf("thread_get_state failed: 0x%x\n", uint32(kr))
	} else {
		fmt.Printf("pc:  %#016x  %s\n", state.PC, symbolize(uintptr(state.PC)))
		fmt.Printf("lr:  %#016x  %s\n", state.LR, symbolize(uintptr(state.LR)))
		fmt.Printf("sp:  %#016x  fp: %#016x  cpsr: %#08x\n", state.SP, state.FP, state.CPSR)
		for i := 0; i < 29; i += 4 {
			for j := i; j < i+4 && j < 29; j++ {
				fmt.Printf("x%-2d %#016x  ", j, state.X[j])
			}
			fmt.Println()
		}
	}
	fmt.Println("=== end report; re-raising to Go's signal handling ===")
}

// dlInfo mirrors Dl_info from <dlfcn.h>.
type dlInfo struct {
	FName *byte
	FBase uintptr
	SName *byte
	SAddr uintptr
}

var dladdr func(addr uintptr, info *dlInfo) int32

func symbolize(addr uintptr) string {
	if dladdr == nil {
		sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "dladdr")
		if err != nil {
			return ""
		}
		purego.RegisterFunc(&dladdr, sym)
	}
	var info dlInfo
	if dladdr(addr, &info) == 0 || info.SName == nil {
		return ""
	}
	return fmt.Sprintf("%s + %#x  (%s)", cstr(info.SName), addr-info.SAddr, cstr(info.FName))
}

func cstr(p *byte) string {
	if p == nil {
		return ""
	}
	var b []byte
	for ptr := unsafe.Pointer(p); *(*byte)(ptr) != 0; ptr = unsafe.Add(ptr, 1) {
		b = append(b, *(*byte)(ptr))
	}
	return string(b)
}
