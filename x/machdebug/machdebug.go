// Package machdebug attaches to a Mach task and controls it: reading and
// writing its memory and registers, setting breakpoints and watchpoints, and
// restoring everything it changed on detach.
//
// It is a debugger, not a bypass. It knows nothing about what it is attached
// to — no code-signing, no entitlements, no policy, no knowledge of any
// particular target binary. Consumers own their own preconditions.
//
// Attaching to an ordinary process requires only that the target be signed
// with com.apple.security.get-task-allow, the standard debugger path: no root
// and no SIP change. Attaching to a platform binary additionally requires root
// and "csrutil enable --without debug"; that requirement belongs to the
// consumer, which must check for it at startup.
//
// # Suspension
//
// A [Process] does not suspend the target on [Attach]. Threads stop only when
// they hit a breakpoint or watchpoint. [Process.Wait] returns an [Event] for
// one such stop, and the thread that reported it stays suspended until that
// Event is released — by the next call to Wait, or by [Process.Detach].
// Other threads keep running throughout. A caller that holds an Event is
// holding exactly one thread, and holding it forever wedges only that thread.
//
// # A dead debugger wedges the target
//
// A stopped thread is blocked in the kernel waiting for a reply on our
// exception port. If this process exits without detaching, every thread
// stopped at a breakpoint stays stopped, and every breakpoint still written
// into the target's text stops the next thread to reach it — forever. Always
// defer [Process.Detach]; Attach installs a SIGINT and SIGTERM handler that
// detaches for you, but nothing can save a target from a SIGKILL.
//
// # arm64e
//
// Targets may be arm64e. Register writes go through a read-modify-write of the
// whole thread state so the pointer-authentication bits in Flags survive; see
// [Thread.SetRegs]. Setting PC to an address that did not come out of the same
// [Regs] is not supported — that needs the signed-pointer convention, and
// getting it subtly wrong faults the target on resume.
//
// # Scope
//
// arm64 only. Rosetta and x86_64 targets are out of scope, as are
// symbolication beyond what debug/macho gives for free, DWARF, and any notion
// of what the target process is.
package machdebug

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/x/mach"
)

// ErrDetached is reported by every operation on a detached [Process].
var ErrDetached = errors.New("machdebug: process detached")

// Process is a task under our control. The zero Process is not usable; get
// one from [Attach].
type Process struct {
	pid  int
	task uint32

	excPort mach.Port
	saved   []savedPorts

	stops  chan *stop    // stops from the exception server, awaiting Wait
	done   chan struct{} // closed to stop the exception server
	served sync.WaitGroup

	sigc chan os.Signal

	mu       sync.Mutex
	detached bool
	held     *stop // the stop whose thread is suspended right now
	bps      map[uint64]*Breakpoint
	wps      []*Watchpoint      // index is the debug-register slot
	pend     map[uint64]pending // threads mid single-step, by thread id
	ss       map[uint64]bool    // threads with single-step armed, by thread id
	ports    []mach.Port        // thread rights handed out by Threads, freed on Detach

	_ noCopy
}

// savedPorts is one row of the target's exception-port table as it was
// before we replaced it, from task_get_exception_ports.
type savedPorts struct {
	mask     kernel.Exception_mask_t
	handler  kernel.Exception_handler_t
	behavior kernel.Exception_behavior_t
	flavor   kernel.Thread_state_flavor_t
}

// stop is one exception message: the event it decodes to, plus what we need
// to let the reporting thread go again.
type stop struct {
	ev     *Event
	reply  mach.Port // send-once right the exception reply goes to
	ndr    [8]byte   // the request's NDR record, echoed in the reply
	thread mach.Port // the reporting thread's port right, ours to release
	resume func()    // ran just before the reply, to set up the resume
}

// Attach takes control of pid: acquires its task port, saves its current
// exception ports, and installs our own for EXC_BREAKPOINT.
//
// Every symbol this package needs is resolved before the target is touched.
// A debugger that loads a library while its target is stopped can deadlock
// against that target.
func Attach(pid int) (*Process, error) {
	if err := resolve(); err != nil {
		return nil, err
	}

	var task uint32
	if kr := taskForPID(kernel.Mach_task_self(), int32(pid), &task); kr != kernSuccess {
		return nil, fmt.Errorf("machdebug: task_for_pid %d: %w", pid, kernErr(kernel.Kern_return_t(kr)))
	}

	p := &Process{pid: pid, task: task, bps: map[uint64]*Breakpoint{}}
	p.wps = make([]*Watchpoint, 0)

	port, err := mach.NewPort()
	if err != nil {
		kernel.Mach_port_deallocate(kernel.Mach_task_self(), kernel.Mach_port_name_t(task))
		return nil, fmt.Errorf("machdebug: exception port: %w", err)
	}
	if err := port.MakeSendRight(); err != nil {
		port.DestroyReceive()
		kernel.Mach_port_deallocate(kernel.Mach_task_self(), kernel.Mach_port_name_t(task))
		return nil, fmt.Errorf("machdebug: exception port send right: %w", err)
	}
	p.excPort = port

	if err := p.saveExceptionPorts(); err != nil {
		p.teardown()
		return nil, err
	}
	// EXCEPTION_DEFAULT|MACH_EXCEPTION_CODES: the message carries thread and
	// task ports and 64-bit codes, no thread state. Same shape machexc uses.
	kr := kernel.Task_set_exception_ports(task,
		kernel.Exception_mask_t(excMaskBreakpoint), uint32(port),
		kernel.Exception_behavior_t(machExcCodes|exceptionDefault), 0)
	if kr != kernSuccess {
		p.teardown()
		return nil, fmt.Errorf("machdebug: task_set_exception_ports: %w", kernErr(kr))
	}

	p.stops = make(chan *stop, 64)
	p.done = make(chan struct{})
	p.served.Add(1)
	go p.serve()

	// A target left with a BRK written and nobody on the exception port hangs
	// forever, so an interrupt has to unwind us before it kills us.
	p.sigc = make(chan os.Signal, 1)
	signal.Notify(p.sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig, ok := <-p.sigc
		if !ok {
			return
		}
		p.Detach()
		signal.Stop(p.sigc)
		syscall.Kill(os.Getpid(), sig.(syscall.Signal))
	}()

	return p, nil
}

// Pid reports the process the [Process] is attached to.
func (p *Process) Pid() int { return p.pid }

// saveExceptionPorts records the target's EXC_BREAKPOINT handlers so Detach
// can put them back.
func (p *Process) saveExceptionPorts() error {
	// EXC_TYPES_COUNT is the table bound; ask for the whole thing.
	const max = 32
	var (
		masks     [max]kernel.Exception_mask_t
		handlers  [max]kernel.Exception_handler_t
		behaviors [max]kernel.Exception_behavior_t
		flavors   [max]kernel.Thread_state_flavor_t
	)
	cnt := kernel.Mach_msg_type_number_t(max)
	kr := kernel.Task_get_exception_ports(p.task, kernel.Exception_mask_t(excMaskBreakpoint),
		&masks[0], &cnt, &handlers[0], &behaviors[0], &flavors[0])
	if kr != kernSuccess {
		return fmt.Errorf("machdebug: task_get_exception_ports: %w", kernErr(kr))
	}
	for i := range int(cnt) {
		p.saved = append(p.saved, savedPorts{masks[i], handlers[i], behaviors[i], flavors[i]})
	}
	return nil
}

// Detach restores every patched instruction, clears every debug register,
// restores the exception ports saved at attach, resumes any thread still held
// by an unreleased [Event], and releases the task port.
//
// It is safe to call more than once, and it must succeed even after a panic:
// a process that exits with a breakpoint still written and no handler on the
// port leaves the target hung at that instruction forever.
func (p *Process) Detach() error {
	p.mu.Lock()
	if p.detached {
		p.mu.Unlock()
		return nil
	}
	p.detached = true
	bps := make([]*Breakpoint, 0, len(p.bps))
	for _, b := range p.bps {
		bps = append(bps, b)
	}
	p.bps = map[uint64]*Breakpoint{}
	p.wps = nil
	p.mu.Unlock()

	var errs []error

	// Unpatch first, and clear the debug registers and any armed single step
	// before letting anyone go: a thread released below must not run back
	// into a BRK, and must not be left stepping with nobody listening.
	for _, b := range bps {
		if err := p.unpatch(b); err != nil {
			errs = append(errs, err)
		}
	}
	p.mu.Lock()
	stepping := p.ss
	p.pend = nil
	p.ss = nil
	p.mu.Unlock()
	if ts, err := p.threads(); err == nil {
		for _, t := range ts {
			// The debug registers are a register file of their own, so
			// clearing them is safe on a thread that is running. The
			// single-step bit lives in CPSR, and putting it back means
			// reading and rewriting the whole thread state — which would
			// race a running thread and resume it on stale registers. Only
			// threads we armed need it, and those are stopped, holding on
			// our reply.
			if err := p.writeDebugStateTo(t.port, t.id); err != nil {
				errs = append(errs, err)
			}
			if !stepping[t.id] {
				continue
			}
			if err := p.clearSingleStepBit(t); err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		errs = append(errs, err)
	}

	// Let go of every thread we are holding, including stops nobody waited
	// for and any trap still in flight. Each is a thread blocked in the
	// kernel on our reply, and a reply we never send is a thread stopped
	// forever.
	if err := p.release(); err != nil {
		errs = append(errs, err)
	}
	deadline := time.Now().Add(detachGrace)
	for time.Now().Before(deadline) {
		select {
		case st := <-p.stops:
			if err := p.releaseStop(st); err != nil {
				errs = append(errs, err)
			}
			deadline = time.Now().Add(detachGrace)
		case <-time.After(10 * time.Millisecond):
		}
	}
	close(p.done)
	p.served.Wait()
	for {
		select {
		case st := <-p.stops:
			if err := p.releaseStop(st); err != nil {
				errs = append(errs, err)
			}
			continue
		default:
		}
		break
	}

	for _, s := range p.saved {
		kr := kernel.Task_set_exception_ports(p.task, s.mask, s.handler, s.behavior, s.flavor)
		// A target that exited on its own takes its exception ports with it;
		// that is not a restore failure.
		if kr != kernSuccess && kr != machSendInvalidDest {
			errs = append(errs, fmt.Errorf("machdebug: restore exception ports: %w", kernErr(kr)))
		}
	}

	signal.Stop(p.sigc)
	close(p.sigc)
	p.teardown()
	return errors.Join(errs...)
}

// detachGrace is how long Detach waits for a trap already in flight to
// arrive so it can be replied to. A trap that arrives after the exception
// server is gone leaves its thread stopped for good.
const detachGrace = 250 * time.Millisecond

// teardown releases the ports this Process owns.
func (p *Process) teardown() {
	for _, t := range p.ports {
		t.Deallocate()
	}
	p.ports = nil
	if p.excPort != 0 {
		p.excPort.Deallocate() // the send right we inserted
		p.excPort.DestroyReceive()
		p.excPort = 0
	}
	if p.task != 0 {
		kernel.Mach_port_deallocate(kernel.Mach_task_self(), kernel.Mach_port_name_t(p.task))
		p.task = 0
	}
}

// alive reports an error if the Process has been detached.
func (p *Process) alive() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.detached {
		return ErrDetached
	}
	return nil
}

// Threads returns the target's current threads.
//
// The port rights behind the returned threads are owned by the Process and
// released by [Process.Detach].
func (p *Process) Threads() ([]Thread, error) {
	if err := p.alive(); err != nil {
		return nil, err
	}
	return p.threads()
}

// threads is Threads without the detached check, for Detach's own restore
// path, which runs after the Process is marked detached.
func (p *Process) threads() ([]Thread, error) {
	var list kernel.Thread_act_array_t
	var cnt kernel.Mach_msg_type_number_t
	if kr := kernel.Task_threads(kernel.Task_inspect_t(p.task), &list, &cnt); kr != kernSuccess {
		return nil, fmt.Errorf("machdebug: task_threads: %w", kernErr(kr))
	}
	ports := unsafe.Slice((*uint32)(unsafe.Pointer(list)), int(cnt))
	ts := make([]Thread, 0, cnt)
	p.mu.Lock()
	p.mu.Unlock()
	for _, port := range ports {
		id, err := p.threadID(port)
		if err != nil {
			continue
		}
		ts = append(ts, Thread{port: port, id: id, p: p})
		p.mu.Lock()
		p.ports = append(p.ports, mach.Port(port))
		p.mu.Unlock()
	}
	kernel.Mach_vm_deallocate(kernel.Mach_task_self(),
		kernel.Mach_vm_address_t(uintptr(unsafe.Pointer(list))),
		kernel.Mach_vm_size_t(uintptr(cnt)*unsafe.Sizeof(uint32(0))))
	return ts, nil
}

// Breakpoint is an armed breakpoint. Clear disarms it.
type Breakpoint struct {
	// Addr is the address the breakpoint was armed at.
	Addr uint64

	p    *Process
	orig [4]byte // the instruction the BRK replaced
	once bool    // disarm on the first hit instead of re-arming

	mu    sync.Mutex
	armed bool

	_ noCopy
}

// SetBreakpoint arms a breakpoint at addr, saving the instruction it replaces.
//
// A breakpoint is persistent: it fires every time the address is executed,
// and the resume path single-steps past the restored instruction before
// re-arming. Use [Process.SetOneShot] for a breakpoint that should fire once.
func (p *Process) SetBreakpoint(addr uint64) (*Breakpoint, error) {
	return p.setBreakpoint(addr, false)
}

// SetOneShot arms a breakpoint at addr that disarms itself when it fires,
// restoring the original instruction without re-arming it. It needs no
// single-step cycle, which makes it both simpler and cheaper than a
// persistent breakpoint.
//
// This is the primitive a temporary return-site breakpoint wants: arm it on
// the LR read at function entry, take the one stop, and it is gone.
func (p *Process) SetOneShot(addr uint64) (*Breakpoint, error) {
	return p.setBreakpoint(addr, true)
}

func (p *Process) setBreakpoint(addr uint64, once bool) (*Breakpoint, error) {
	if err := p.alive(); err != nil {
		return nil, err
	}
	if addr%4 != 0 {
		return nil, fmt.Errorf("machdebug: breakpoint at %#x: address is not instruction-aligned", addr)
	}
	p.mu.Lock()
	if _, dup := p.bps[addr]; dup {
		p.mu.Unlock()
		return nil, fmt.Errorf("machdebug: breakpoint at %#x: already set", addr)
	}
	p.mu.Unlock()

	b := &Breakpoint{Addr: addr, p: p, once: once}
	if _, err := p.ReadAt(b.orig[:], int64(addr)); err != nil {
		return nil, fmt.Errorf("machdebug: breakpoint at %#x: read original: %w", addr, err)
	}
	if err := p.patch(b); err != nil {
		return nil, err
	}
	p.mu.Lock()
	p.bps[addr] = b
	p.mu.Unlock()
	return b, nil
}

// patch writes BRK #0 over the breakpoint's address.
func (p *Process) patch(b *Breakpoint) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.armed {
		return nil
	}
	var brk [4]byte
	binary.LittleEndian.PutUint32(brk[:], brkZero)
	// Text is read-only and shared; the write must take the VM_PROT_COPY
	// path so the patch is private to this task.
	if err := p.writeProtected(b.Addr, brk[:]); err != nil {
		return fmt.Errorf("machdebug: arm breakpoint at %#x: %w", b.Addr, err)
	}
	b.armed = true
	return nil
}

// unpatch puts the original instruction back.
func (p *Process) unpatch(b *Breakpoint) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.armed {
		return nil
	}
	if err := p.writeProtected(b.Addr, b.orig[:]); err != nil {
		return fmt.Errorf("machdebug: disarm breakpoint at %#x: %w", b.Addr, err)
	}
	b.armed = false
	return nil
}

// Clear disarms the breakpoint, restoring the original instruction. It is
// safe to call on a one-shot that has already fired.
func (b *Breakpoint) Clear() error {
	if b.p == nil {
		return errors.New("machdebug: clear: breakpoint was never set")
	}
	if err := b.p.alive(); err != nil {
		return err
	}
	b.p.mu.Lock()
	delete(b.p.bps, b.Addr)
	b.p.mu.Unlock()
	return b.p.unpatch(b)
}

// Watch says which accesses a watchpoint traps.
type Watch int

// Watch kinds.
const (
	WatchRead Watch = 1 + iota
	WatchWrite
	WatchReadWrite
)

// Watchpoint is an armed watchpoint. Clear disarms it.
type Watchpoint struct {
	// Addr and Size describe the watched region.
	Addr uint64
	Size int

	p    *Process
	slot int
	kind Watch

	_ noCopy
}

// SetWatchpoint traps accesses of kind w to the size bytes at addr, using the
// hardware debug registers. size must be a power of two no larger than 8, and
// addr must be size-aligned. The number of watchpoints is a small hardware
// limit; SetWatchpoint reports an error when they are exhausted.
func (p *Process) SetWatchpoint(addr uint64, size int, w Watch) (*Watchpoint, error) {
	if err := p.alive(); err != nil {
		return nil, err
	}
	switch size {
	case 1, 2, 4, 8:
	default:
		return nil, fmt.Errorf("machdebug: watchpoint size %d: must be 1, 2, 4 or 8", size)
	}
	if addr%uint64(size) != 0 {
		return nil, fmt.Errorf("machdebug: watchpoint at %#x: address is not %d-byte aligned", addr, size)
	}
	switch w {
	case WatchRead, WatchWrite, WatchReadWrite:
	default:
		return nil, fmt.Errorf("machdebug: watchpoint kind %d: unknown", int(w))
	}

	wp := &Watchpoint{Addr: addr, Size: size, p: p, kind: w}
	p.mu.Lock()
	slot := -1
	for i, cur := range p.wps {
		if cur == nil {
			slot = i
			break
		}
	}
	if slot < 0 {
		if len(p.wps) >= maxWatchpoints {
			p.mu.Unlock()
			return nil, fmt.Errorf("machdebug: watchpoint at %#x: all %d debug registers in use", addr, maxWatchpoints)
		}
		p.wps = append(p.wps, nil)
		slot = len(p.wps) - 1
	}
	wp.slot = slot
	p.wps[slot] = wp
	p.mu.Unlock()

	if err := p.writeDebugState(); err != nil {
		p.mu.Lock()
		p.wps[slot] = nil
		p.mu.Unlock()
		return nil, err
	}
	return wp, nil
}

// Clear disarms the watchpoint and releases its debug register.
func (w *Watchpoint) Clear() error {
	if w.p == nil {
		return errors.New("machdebug: clear: watchpoint was never set")
	}
	if err := w.p.alive(); err != nil {
		return err
	}
	w.p.mu.Lock()
	if w.slot < len(w.p.wps) && w.p.wps[w.slot] == w {
		w.p.wps[w.slot] = nil
	}
	w.p.mu.Unlock()
	return w.p.writeDebugState()
}

// EventKind says why a thread stopped.
type EventKind int

// Event kinds.
const (
	// EventBreakpoint is a breakpoint hit. Addr is the breakpoint address.
	EventBreakpoint EventKind = 1 + iota
	// EventWatchpoint is a watchpoint hit. Addr is the watched address.
	EventWatchpoint
	// EventStep is a single-step completion, reported only to callers that
	// asked for one with [Thread.Step].
	EventStep
)

// String names the kind.
func (k EventKind) String() string {
	switch k {
	case EventBreakpoint:
		return "breakpoint"
	case EventWatchpoint:
		return "watchpoint"
	case EventStep:
		return "step"
	}
	return fmt.Sprintf("EventKind(%d)", int(k))
}

// Event is one stop. The reporting thread stays suspended until the Event is
// released by the next call to [Process.Wait] or by [Process.Detach].
type Event struct {
	Thread Thread
	Addr   uint64
	Kind   EventKind
}

// Wait returns the next stop, releasing the thread held by the previous
// Event. It blocks until a thread stops or the process detaches.
func (p *Process) Wait() (*Event, error) {
	return p.WaitTimeout(0)
}

// WaitTimeout is [Process.Wait] with a deadline: it reports
// [os.ErrDeadlineExceeded] if no thread stops within d. A zero d blocks.
func (p *Process) WaitTimeout(d time.Duration) (*Event, error) {
	if err := p.alive(); err != nil {
		return nil, err
	}
	if err := p.release(); err != nil {
		return nil, err
	}
	var timeout <-chan time.Time
	if d > 0 {
		t := time.NewTimer(d)
		defer t.Stop()
		timeout = t.C
	}
	select {
	case s := <-p.stops:
		p.mu.Lock()
		p.held = s
		p.mu.Unlock()
		return s.ev, nil
	case <-p.done:
		return nil, ErrDetached
	case <-timeout:
		return nil, os.ErrDeadlineExceeded
	}
}

// release replies to the held stop, letting its thread run again.
func (p *Process) release() error {
	p.mu.Lock()
	s := p.held
	p.held = nil
	p.mu.Unlock()
	if s == nil {
		return nil
	}
	return p.releaseStop(s)
}

// releaseStop replies KERN_SUCCESS to one stop.
func (p *Process) releaseStop(s *stop) error {
	defer s.thread.Deallocate()
	// A resume action re-arms a breakpoint or starts a single step. Neither
	// is something a detaching Process should be setting up.
	p.mu.Lock()
	detached := p.detached
	p.mu.Unlock()
	if s.resume != nil && !detached {
		s.resume()
	}

	// KERN_SUCCESS, not KERN_FAILURE: KERN_SUCCESS says the exception is
	// handled and the thread should resume. KERN_FAILURE re-raises to the
	// host handler, which is right for a crash reporter and fatal for a
	// debugger — it turns every breakpoint into a SIGTRAP that kills the
	// target.
	reply := make([]byte, 12)
	copy(reply, s.ndr[:])
	binary.LittleEndian.PutUint32(reply[8:], kernSuccess)
	if err := mach.Send(s.reply, mach.MoveSendOnce, msgIDRaiseReply, nil, reply, 0); err != nil {
		return fmt.Errorf("machdebug: exception reply: %w", err)
	}
	return nil
}

// Thread is one thread of the target task.
type Thread struct {
	port uint32
	id   uint64 // the system-wide thread id; port names are per-message
	p    *Process
}

// Port reports the thread's Mach port name in this task's namespace.
func (t Thread) Port() uint32 { return t.port }

// Regs reads the thread's full register state.
func (t Thread) Regs() (*Regs, error) {
	if t.p == nil {
		return nil, errors.New("machdebug: regs: zero Thread")
	}
	if err := t.p.alive(); err != nil {
		return nil, err
	}
	return t.regs()
}

// regs is Regs without the detached check, for this package's own paths,
// including Detach's restore.
func (t Thread) regs() (*Regs, error) {
	var r Regs
	cnt := kernel.Mach_msg_type_number_t(armThreadState64Count)
	kr := kernel.Thread_get_state(kernel.Thread_read_t(t.port),
		kernel.Thread_state_flavor_t(armThreadState64),
		kernel.Thread_state_t(unsafe.Pointer(&r)), &cnt)
	if kr != kernSuccess {
		return nil, fmt.Errorf("machdebug: thread_get_state: %w", kernErr(kr))
	}
	return &r, nil
}

// SetRegs writes the thread's register state.
//
// Always read with [Thread.Regs], mutate the fields you mean, and write the
// same value back. A synthesized Regs loses the pointer-authentication state
// in Flags and faults the thread on resume.
//
// Redirecting PC is not supported: on arm64e a PC that did not come out of
// the target's own state needs the signed-pointer convention, and SetRegs
// reports an error rather than getting that subtly wrong. Writing back the
// PC that Regs returned is fine.
func (t Thread) SetRegs(r *Regs) error {
	if t.p == nil {
		return errors.New("machdebug: setregs: zero Thread")
	}
	if r == nil {
		return errors.New("machdebug: setregs: nil Regs")
	}
	if err := t.p.alive(); err != nil {
		return err
	}
	cur, err := t.Regs()
	if err != nil {
		return err
	}
	if cur.PC != r.PC {
		return fmt.Errorf("machdebug: setregs: pc %#x differs from the thread's %#x: pc redirection is not supported", r.PC, cur.PC)
	}
	return t.setRegs(r)
}

// setRegs is SetRegs without the PC guard, for this package's own resume
// paths, which move PC only by values the thread itself produced.
func (t Thread) setRegs(r *Regs) error {
	w := *r
	cnt := kernel.Mach_msg_type_number_t(armThreadState64Count)
	kr := kernel.Thread_set_state(kernel.Thread_act_t(t.port),
		kernel.Thread_state_flavor_t(armThreadState64),
		kernel.Thread_state_t(unsafe.Pointer(&w)), cnt)
	if kr != kernSuccess {
		return fmt.Errorf("machdebug: thread_set_state: %w", kernErr(kr))
	}
	return nil
}

// noCopy makes go vet's copylocks check reject copying a value that owns
// target-side state.
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// kernErr turns a kern_return_t into an error, nil on KERN_SUCCESS.
func kernErr(kr kernel.Kern_return_t) error {
	if kr == kernSuccess {
		return nil
	}
	return fmt.Errorf("kern_return %d (%#x)", kr, uint32(kr))
}
