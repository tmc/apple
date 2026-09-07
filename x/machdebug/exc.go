package machdebug

import (
	"encoding/binary"
	"fmt"
	"time"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/x/mach"
)

// Exception subcodes for EXC_BREAKPOINT on arm64, from
// <mach/arm/exception.h>. ABI-pinned in xnu, not in Apple's documentation
// set, so not generated.
const (
	excARMBreakpoint = 1     // EXC_ARM_BREAKPOINT: a BRK or a single step
	excARMDADebug    = 0x102 // EXC_ARM_DA_DEBUG: a watchpoint
)

// maxWatchpoints is the number of hardware watchpoint register pairs this
// package will use. arm_debug_state64_t has room for 16; Apple silicon
// implements 4, and asking for more than the hardware has makes
// thread_set_state fail as a whole.
const maxWatchpoints = 4

// pending is what to do when a thread that we single-stepped stops again.
type pending struct {
	rearm   *Breakpoint // re-arm this breakpoint after the step
	deliver bool        // report the step to Wait as an EventStep
}

// serve receives exception messages until Detach closes p.done.
//
// The message plumbing is the one from examples/kernel/machexc: two port
// descriptors (thread, task), then the NDR record, the exception number, the
// code count, and that many 64-bit codes. The reply echoes the NDR record and
// appends a kern_return_t at msgh_id request+100.
func (p *Process) serve() {
	defer p.served.Done()
	for {
		select {
		case <-p.done:
			return
		default:
		}
		// A bounded receive so Detach can stop us; the exception port has no
		// other traffic, so the poll costs nothing.
		m, err := mach.Receive(p.excPort, 50*time.Millisecond)
		if err != nil {
			continue
		}
		p.handle(m)
	}
}

func (p *Process) handle(m *mach.Message) {
	if m.Header.ID != msgIDRaise || len(m.Ports) < 2 || len(m.Body) < 16 {
		return
	}
	thread := m.Ports[0]
	m.Ports[1].Deallocate() // the task port right; we already have our own

	exception := int32(binary.LittleEndian.Uint32(m.Body[8:12]))
	codeCnt := int(binary.LittleEndian.Uint32(m.Body[12:16]))
	codes := make([]int64, 0, codeCnt)
	for i := range codeCnt {
		if off := 16 + 8*i; off+8 <= len(m.Body) {
			codes = append(codes, int64(binary.LittleEndian.Uint64(m.Body[off:off+8])))
		}
	}

	s := &stop{reply: m.Header.RemotePort, thread: thread}
	copy(s.ndr[:], m.Body[:8])

	if exception != excBreakpoint {
		p.decline(s)
		return
	}
	var subcode, addr int64
	if len(codes) > 0 {
		subcode = codes[0]
	}
	if len(codes) > 1 {
		addr = codes[1]
	}
	// Every exception message carries a fresh port name for the same thread,
	// so a port name cannot identify a thread across two stops. The
	// system-wide thread id from THREAD_IDENTIFIER_INFO can.
	id, err := p.threadID(uint32(thread))
	if err != nil {
		p.decline(s)
		return
	}
	t := Thread{port: uint32(thread), id: id, p: p}

	// A thread we single-stepped is stopping again: finish the cycle before
	// anything else can interpret the address.
	p.mu.Lock()
	pend, stepping := p.pend[id]
	if stepping {
		delete(p.pend, id)
	}
	p.mu.Unlock()
	if stepping {
		p.setSingleStep(t, false)
		// Do not re-arm into a Process that is detaching: Detach has already
		// taken every patch back out.
		if pend.rearm != nil && p.alive() == nil {
			p.patch(pend.rearm)
		}
		if !pend.deliver {
			p.releaseStop(s)
			return
		}
		// A single-step exception reports no address of its own; the
		// interesting one is where the thread now stands.
		pc := uint64(addr)
		if r, err := t.Regs(); err == nil {
			pc = r.PC
		}
		s.ev = &Event{Thread: t, Addr: pc, Kind: EventStep}
		p.deliver(s)
		return
	}

	switch subcode {
	case excARMDADebug:
		if wp := p.watchpointAt(uint64(addr)); wp != nil {
			s.ev = &Event{Thread: t, Addr: wp.Addr, Kind: EventWatchpoint}
			p.deliver(s)
			return
		}
		p.decline(s)
		return

	case excARMBreakpoint:
		p.mu.Lock()
		b := p.bps[uint64(addr)]
		p.mu.Unlock()
		if b == nil {
			// Not our BRK. Declining re-raises it to whatever the target had
			// installed before we attached, which is where it belongs.
			p.decline(s)
			return
		}
		// The BRK sits where the instruction has to be, so it comes out
		// before the thread can resume.
		if err := p.unpatch(b); err != nil {
			p.decline(s)
			return
		}
		if b.once {
			p.mu.Lock()
			delete(p.bps, b.Addr)
			p.mu.Unlock()
		} else {
			// Persistent: step one instruction past the restored original,
			// then put the BRK back. Arming happens at release time so the
			// caller sees a clean, unstepping thread while it holds the
			// Event.
			s.resume = func() { p.armStep(t, pending{rearm: b}) }
		}
		s.ev = &Event{Thread: t, Addr: uint64(addr), Kind: EventBreakpoint}
		p.deliver(s)
		return
	}
	p.decline(s)
}

// deliver hands a stop to Wait, or releases it if the Process is detaching.
func (p *Process) deliver(s *stop) {
	select {
	case p.stops <- s:
	case <-p.done:
		p.releaseStop(s)
	}
}

// watchpointAt reports the watchpoint covering addr, if any.
func (p *Process) watchpointAt(addr uint64) *Watchpoint {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, w := range p.wps {
		if w != nil && addr >= w.Addr && addr < w.Addr+uint64(w.Size) {
			return w
		}
	}
	return nil
}

// decline replies KERN_FAILURE, re-raising the exception to the handler the
// target had before we attached. It is the right answer for a trap we did not
// set: swallowing one would resume the thread onto an instruction that traps
// again, forever.
func (p *Process) decline(s *stop) {
	defer s.thread.Deallocate()
	const kernFailure = 5
	reply := make([]byte, 12)
	copy(reply, s.ndr[:])
	binary.LittleEndian.PutUint32(reply[8:], kernFailure)
	mach.Send(s.reply, mach.MoveSendOnce, msgIDRaiseReply, nil, reply, 0)
}

// armStep sets the thread up to trap after one instruction.
func (p *Process) armStep(t Thread, pend pending) {
	p.mu.Lock()
	if p.pend == nil {
		p.pend = map[uint64]pending{}
	}
	p.pend[t.id] = pend
	p.mu.Unlock()
	p.setSingleStep(t, true)
}

// setSingleStep turns hardware single stepping on or off for one thread.
//
// It takes both bits: MDSCR_EL1.SS in the debug state enables the mechanism,
// and PSTATE.SS (CPSR bit 21) in the thread state says the next instruction
// is the one to step. Setting only one of them does nothing.
func (p *Process) setSingleStep(t Thread, on bool) error {
	p.mu.Lock()
	if p.ss == nil {
		p.ss = map[uint64]bool{}
	}
	if on {
		p.ss[t.id] = true
	} else {
		delete(p.ss, t.id)
	}
	p.mu.Unlock()

	if err := p.writeDebugStateTo(t.port, t.id); err != nil {
		return err
	}
	r, err := t.regs()
	if err != nil {
		return err
	}
	if on {
		r.CPSR |= cpsrSS
	} else {
		r.CPSR &^= cpsrSS
	}
	return t.setRegs(r)
}

// clearSingleStepBit turns PSTATE.SS off without touching the debug
// registers. The caller must know the thread is stopped: the write is a
// read-modify-write of the whole thread state, and on a running thread it
// would put stale registers back.
func (p *Process) clearSingleStepBit(t Thread) error {
	r, err := t.regs()
	if err != nil {
		return err
	}
	if r.CPSR&cpsrSS == 0 {
		return nil
	}
	r.CPSR &^= cpsrSS
	return t.setRegs(r)
}

// Step single-steps the thread one instruction. The completion arrives from
// [Process.Wait] as an [EventStep].
//
// The thread must be stopped — that is, the caller must be holding the
// [Event] that reported it. The step begins when the Event is released.
func (t Thread) Step() error {
	if t.p == nil {
		return fmt.Errorf("machdebug: step: zero Thread")
	}
	if err := t.p.alive(); err != nil {
		return err
	}
	t.p.mu.Lock()
	s := t.p.held
	if s == nil || s.ev == nil || s.ev.Thread.port != t.port {
		t.p.mu.Unlock()
		return fmt.Errorf("machdebug: step: thread %d is not the thread holding the current event", t.port)
	}
	prev := s.resume
	s.resume = func() {
		pend := pending{deliver: true}
		if prev != nil {
			// A persistent breakpoint already wanted this step for its
			// re-arm; take that over and report it too.
			t.p.mu.Lock()
			held := t.p.pend[t.id]
			t.p.mu.Unlock()
			pend.rearm = held.rearm
			prev()
		}
		t.p.armStep(t, pend)
	}
	t.p.mu.Unlock()
	return nil
}

// debugState is arm_debug_state64_t from <mach/arm/_structs.h>: 16 hardware
// breakpoint value/control pairs, 16 watchpoint value/control pairs, and
// MDSCR_EL1, whose bit 0 is the single-step enable.
type debugState struct {
	BVR      [16]uint64
	BCR      [16]uint64
	WVR      [16]uint64
	WCR      [16]uint64
	MDSCREL1 uint64
}

// writeDebugState pushes the current watchpoint set to every thread.
func (p *Process) writeDebugState() error {
	ts, err := p.threads()
	if err != nil {
		return err
	}
	for _, t := range ts {
		if err := p.writeDebugStateTo(t.port, t.id); err != nil {
			return err
		}
	}
	return nil
}

// writeDebugStateTo pushes the watchpoint set plus that thread's single-step
// bit to one thread.
func (p *Process) writeDebugStateTo(port uint32, id uint64) error {
	var st debugState
	p.mu.Lock()
	for i, w := range p.wps {
		if w == nil || i >= maxWatchpoints {
			continue
		}
		st.WVR[i] = w.Addr &^ 7
		bas := uint64((1<<uint(w.Size))-1) << (w.Addr & 7)
		ctrl := uint64(wcrEnable | wcrPACEL0 | (bas << wcrBASShift))
		switch w.kind {
		case WatchRead:
			ctrl |= wcrLoad
		case WatchWrite:
			ctrl |= wcrStore
		case WatchReadWrite:
			ctrl |= wcrLoad | wcrStore
		}
		st.WCR[i] = ctrl
	}
	if p.ss[id] {
		st.MDSCREL1 |= mdscrSS
	}
	p.mu.Unlock()

	cnt := kernel.Mach_msg_type_number_t(armDebugState64Count)
	kr := kernel.Thread_set_state(kernel.Thread_act_t(port),
		kernel.Thread_state_flavor_t(armDebugState64),
		kernel.Thread_state_t(unsafe.Pointer(&st)), cnt)
	if kr != kernSuccess {
		return fmt.Errorf("machdebug: thread_set_state ARM_DEBUG_STATE64: %w", kernErr(kr))
	}
	return nil
}

// threadInfo is thread_identifier_info from <mach/thread_info.h>.
type threadInfo struct {
	ID     uint64
	Handle uint64
	QAddr  uint64
}

// THREAD_IDENTIFIER_INFO and its count, the struct size in uint32 words.
// ABI-pinned in xnu, absent from Apple's documentation set.
const (
	threadIdentifierInfo      = 4
	threadIdentifierInfoCount = 6
)

// threadID reports the system-wide unique id of the thread behind a port
// name. Every exception message carries a fresh name for the same thread, so
// the id is the only handle that survives two stops.
func (p *Process) threadID(port uint32) (uint64, error) {
	var info threadInfo
	cnt := kernel.Mach_msg_type_number_t(threadIdentifierInfoCount)
	kr := kernel.Thread_info(kernel.Thread_inspect_t(port),
		kernel.Thread_flavor_t(threadIdentifierInfo),
		kernel.Thread_info_t(unsafe.Pointer(&info)), &cnt)
	if kr != kernSuccess {
		return 0, fmt.Errorf("machdebug: thread_info THREAD_IDENTIFIER_INFO: %w", kernErr(kr))
	}
	return info.ID, nil
}
