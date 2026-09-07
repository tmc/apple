//go:build darwin

package lldb

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// cell is the opaque storage for one SB value. The underlying C++ object is at
// most 16 bytes (a smart pointer), but the field is padded to 32 bytes so it
// exceeds 16: purego only routes a struct return through the indirect-result
// register (arm64 x8, or a hidden memory pointer on amd64) when the return type
// is larger than 16 bytes, which is what the C++ ABI requires for these
// non-trivially-copyable classes.
type cell struct{ _, _, _, _ uintptr }

// object owns one constructed SB value. Wrapper copies share this object,
// including the guard, rather than copying a C++ smart pointer.
type object struct {
	cell    *cell
	once    sync.Once
	destroy func(unsafe.Pointer)
}

func own(c cell, destroy func(unsafe.Pointer)) *object {
	return &object{cell: &c, destroy: destroy}
}

func (o *object) ptr() unsafe.Pointer { return unsafe.Pointer(o.cell) }

func (o *object) close() {
	if o == nil {
		return
	}
	o.once.Do(func() {
		o.destroy(o.ptr())
	})
}

// SB API function pointers, resolved once in load.
var (
	fnDestroyDebugger func(self unsafe.Pointer)
	fnDebuggerDtor    func(self unsafe.Pointer)
	fnListenerDtor    func(self unsafe.Pointer)
	fnTargetDtor      func(self unsafe.Pointer)
	fnBreakpointDtor  func(self unsafe.Pointer)
	fnProcessDtor     func(self unsafe.Pointer)
	fnThreadDtor      func(self unsafe.Pointer)
	fnFrameDtor       func(self unsafe.Pointer)
	fnValueDtor       func(self unsafe.Pointer)
	fnErrorDtor       func(self unsafe.Pointer)

	fnInitialize func()
	fnCreate     func() cell
	fnSetAsync   func(self unsafe.Pointer, async bool)
	fnCreateTgt  func(self unsafe.Pointer, triple *byte) cell
	fnCreateTgtA func(self unsafe.Pointer, file, arch *byte) cell
	fnListener   func(self unsafe.Pointer) cell

	fnAttachName func(self, listener unsafe.Pointer, name *byte, waitFor bool, err unsafe.Pointer) cell
	fnLaunch     func(self unsafe.Pointer, argv, envp **byte, cwd *byte) cell
	fnTriple     func(self unsafe.Pointer) unsafe.Pointer
	fnBreakName  func(self unsafe.Pointer, name, module *byte) cell
	fnEval       func(self unsafe.Pointer, expr *byte) cell

	fnContinue   func(self unsafe.Pointer) cell
	fnState      func(self unsafe.Pointer) int32
	fnNumThreads func(self unsafe.Pointer) uint32
	fnThreadAt   func(self unsafe.Pointer, idx uint64) cell
	fnProcValid  func(self unsafe.Pointer) bool
	fnKill       func(self unsafe.Pointer) cell
	fnInterrupt  func(self unsafe.Pointer)
	fnDetach     func(self unsafe.Pointer) cell
	fnBpValid    func(self unsafe.Pointer) bool

	fnStopReason func(self unsafe.Pointer) int32
	fnFrameAt    func(self unsafe.Pointer, idx uint32) cell
	fnStepOut    func(self, frame unsafe.Pointer)

	fnFindReg  func(self unsafe.Pointer, name *byte) cell
	fnFuncName func(self unsafe.Pointer) unsafe.Pointer
	fnGetValue func(self unsafe.Pointer) unsafe.Pointer
	fnSetCStr  func(self unsafe.Pointer, val *byte) bool
	fnObjDesc  func(self unsafe.Pointer) unsafe.Pointer
	fnAsUnsig  func(self unsafe.Pointer, fail uint64) uint64
	fnValValid func(self unsafe.Pointer) bool

	fnErrCtor    func(self unsafe.Pointer)
	fnErrSuccess func(self unsafe.Pointer) bool
	fnErrCStr    func(self unsafe.Pointer) unsafe.Pointer
)

var (
	loadOnce sync.Once
	loadErr  error
)

// symbols maps each function pointer to its mangled C++ symbol.
func symbols() []struct {
	p    any
	name string
} {
	return []struct {
		p    any
		name string
	}{
		{&fnDestroyDebugger, "_ZN4lldb10SBDebugger7DestroyERS0_"},
		{&fnDebuggerDtor, "_ZN4lldb10SBDebuggerD1Ev"},
		{&fnListenerDtor, "_ZN4lldb10SBListenerD1Ev"},
		{&fnTargetDtor, "_ZN4lldb8SBTargetD1Ev"},
		{&fnBreakpointDtor, "_ZN4lldb12SBBreakpointD1Ev"},
		{&fnProcessDtor, "_ZN4lldb9SBProcessD1Ev"},
		{&fnThreadDtor, "_ZN4lldb8SBThreadD1Ev"},
		{&fnFrameDtor, "_ZN4lldb7SBFrameD1Ev"},
		{&fnValueDtor, "_ZN4lldb7SBValueD1Ev"},
		{&fnErrorDtor, "_ZN4lldb7SBErrorD1Ev"},

		{&fnInitialize, "_ZN4lldb10SBDebugger10InitializeEv"},
		{&fnCreate, "_ZN4lldb10SBDebugger6CreateEv"},
		{&fnSetAsync, "_ZN4lldb10SBDebugger8SetAsyncEb"},
		{&fnCreateTgt, "_ZN4lldb10SBDebugger12CreateTargetEPKc"},
		{&fnCreateTgtA, "_ZN4lldb10SBDebugger27CreateTargetWithFileAndArchEPKcS2_"},
		{&fnListener, "_ZN4lldb10SBDebugger11GetListenerEv"},

		{&fnAttachName, "_ZN4lldb8SBTarget23AttachToProcessWithNameERNS_10SBListenerEPKcbRNS_7SBErrorE"},
		{&fnLaunch, "_ZN4lldb8SBTarget12LaunchSimpleEPPKcS3_S2_"},
		{&fnTriple, "_ZN4lldb8SBTarget9GetTripleEv"},
		{&fnBreakName, "_ZN4lldb8SBTarget22BreakpointCreateByNameEPKcS2_"},
		{&fnEval, "_ZN4lldb8SBTarget18EvaluateExpressionEPKc"},

		{&fnContinue, "_ZN4lldb9SBProcess8ContinueEv"},
		{&fnState, "_ZN4lldb9SBProcess8GetStateEv"},
		{&fnNumThreads, "_ZN4lldb9SBProcess13GetNumThreadsEv"},
		{&fnThreadAt, "_ZN4lldb9SBProcess16GetThreadAtIndexEm"},
		{&fnProcValid, "_ZNK4lldb9SBProcess7IsValidEv"},
		{&fnKill, "_ZN4lldb9SBProcess4KillEv"},
		{&fnInterrupt, "_ZN4lldb9SBProcess18SendAsyncInterruptEv"},
		{&fnDetach, "_ZN4lldb9SBProcess6DetachEv"},
		{&fnBpValid, "_ZNK4lldb12SBBreakpoint7IsValidEv"},

		{&fnStopReason, "_ZN4lldb8SBThread13GetStopReasonEv"},
		{&fnFrameAt, "_ZN4lldb8SBThread15GetFrameAtIndexEj"},
		{&fnStepOut, "_ZN4lldb8SBThread14StepOutOfFrameERNS_7SBFrameE"},

		{&fnFindReg, "_ZN4lldb7SBFrame12FindRegisterEPKc"},
		{&fnFuncName, "_ZNK4lldb7SBFrame15GetFunctionNameEv"},

		{&fnGetValue, "_ZN4lldb7SBValue8GetValueEv"},
		{&fnSetCStr, "_ZN4lldb7SBValue19SetValueFromCStringEPKc"},
		{&fnObjDesc, "_ZN4lldb7SBValue20GetObjectDescriptionEv"},
		{&fnAsUnsig, "_ZN4lldb7SBValue18GetValueAsUnsignedEy"},
		{&fnValValid, "_ZN4lldb7SBValue7IsValidEv"},

		{&fnErrCtor, "_ZN4lldb7SBErrorC1Ev"},
		{&fnErrSuccess, "_ZNK4lldb7SBError7SuccessEv"},
		{&fnErrCStr, "_ZNK4lldb7SBError10GetCStringEv"},
	}
}

// load opens LLDB.framework and resolves every SB symbol exactly once.
func load() error {
	loadOnce.Do(func() {
		path, err := frameworkPath()
		if err != nil {
			loadErr = err
			return
		}
		h, err := purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			loadErr = fmt.Errorf("lldb: dlopen %s: %w", path, err)
			return
		}
		for _, s := range symbols() {
			if err := register(s.p, h, s.name); err != nil {
				loadErr = err
				return
			}
		}
		fnInitialize()
	})
	return loadErr
}

func register(fptr any, h uintptr, name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("lldb: register %s: %v", name, r)
		}
	}()
	purego.RegisterLibFunc(fptr, h, name)
	return nil
}

// frameworkPath returns the path to the LLDB.framework mach-o binary.
func frameworkPath() (string, error) {
	var candidates []string
	if out, err := exec.Command("lldb", "-P").Output(); err == nil {
		p := strings.TrimSpace(string(out))
		if i := strings.Index(p, "LLDB.framework"); i >= 0 {
			base := p[:i+len("LLDB.framework")]
			candidates = append(candidates,
				filepath.Join(base, "Versions/A/LLDB"),
				filepath.Join(base, "LLDB"))
		}
	}
	candidates = append(candidates,
		"/Library/Developer/CommandLineTools/Library/PrivateFrameworks/LLDB.framework/Versions/A/LLDB",
		"/Applications/Xcode.app/Contents/SharedFrameworks/LLDB.framework/Versions/A/LLDB",
		"/Applications/Xcode-beta.app/Contents/SharedFrameworks/LLDB.framework/Versions/A/LLDB",
	)
	for _, c := range candidates {
		if fileExists(c) {
			return c, nil
		}
	}
	return "", errors.New("lldb: LLDB.framework not found (install Xcode or the Command Line Tools)")
}

// State is an SBProcess execution state (lldb::StateType).
type State int32

const (
	StateInvalid   State = 0
	StateUnloaded  State = 1
	StateConnected State = 2
	StateAttaching State = 3
	StateLaunching State = 4
	StateStopped   State = 5
	StateRunning   State = 6
	StateStepping  State = 7
	StateCrashed   State = 8
	StateDetached  State = 9
	StateExited    State = 10
	StateSuspended State = 11
)

func (s State) String() string {
	switch s {
	case StateInvalid:
		return "invalid"
	case StateUnloaded:
		return "unloaded"
	case StateConnected:
		return "connected"
	case StateAttaching:
		return "attaching"
	case StateLaunching:
		return "launching"
	case StateStopped:
		return "stopped"
	case StateRunning:
		return "running"
	case StateStepping:
		return "stepping"
	case StateCrashed:
		return "crashed"
	case StateDetached:
		return "detached"
	case StateExited:
		return "exited"
	case StateSuspended:
		return "suspended"
	}
	return fmt.Sprintf("State(%d)", int32(s))
}

// StopReason is why a thread stopped (lldb::StopReason).
type StopReason int32

const (
	StopReasonInvalid    StopReason = 0
	StopReasonNone       StopReason = 1
	StopReasonTrace      StopReason = 2
	StopReasonBreakpoint StopReason = 3
	StopReasonWatchpoint StopReason = 4
	StopReasonSignal     StopReason = 5
	StopReasonException  StopReason = 6
	StopReasonExec       StopReason = 7
	StopReasonPlandone   StopReason = 8
)

// Debugger wraps lldb::SBDebugger.
type Debugger struct{ c *object }

// Close destroys the debugger and releases its native handle. Close all child
// handles and detach or kill debugged processes first. Close is idempotent and
// safe on a zero value; copies share the same lifetime.
func (d Debugger) Close() { d.c.close() }

// Create initializes LLDB (once) and returns a new debugger.
func Create() (Debugger, error) {
	if err := load(); err != nil {
		return Debugger{}, err
	}
	return Debugger{c: own(fnCreate(), destroyDebugger)}, nil
}

// SetAsync sets asynchronous execution mode.
func (d Debugger) SetAsync(async bool) { fnSetAsync(d.c.ptr(), async) }

// CreateTarget creates a target for file, or an empty target if file is "".
func (d Debugger) CreateTarget(file string) Target {
	return Target{c: own(fnCreateTgt(d.c.ptr(), cstr(file)), fnTargetDtor)}
}

// CreateTargetWithArch creates a target for the executable at file. arch may be
// "" to use the host architecture.
func (d Debugger) CreateTargetWithArch(file, arch string) Target {
	var archp *byte
	if arch != "" {
		archp = cstr(arch)
	}
	return Target{c: own(fnCreateTgtA(d.c.ptr(), cstr(file), archp), fnTargetDtor)}
}

// Listener returns the debugger's listener.
func (d Debugger) Listener() Listener { return Listener{c: own(fnListener(d.c.ptr()), fnListenerDtor)} }

// Listener wraps lldb::SBListener.
type Listener struct{ c *object }

// Close releases the native listener handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (l Listener) Close() { l.c.close() }

// Target wraps lldb::SBTarget.
type Target struct{ c *object }

// Close releases the native target handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (t Target) Close() { t.c.close() }

// AttachToProcessWithName attaches to a process by executable path. The returned
// Error reports why attachment failed when the process is not valid.
func (t Target) AttachToProcessWithName(l Listener, name string, waitFor bool) (Process, Error) {
	e := newError()
	p := fnAttachName(t.c.ptr(), l.c.ptr(), cstr(name), waitFor, e.c.ptr())
	return Process{c: own(p, fnProcessDtor)}, e
}

// LaunchSimple launches the target's executable and returns the process. argv
// and envp may be nil; cwd may be "". With asynchronous mode disabled the call
// returns once the process stops (for example at a breakpoint).
func (t Target) LaunchSimple(argv, envp []string, cwd string) Process {
	var cwdp *byte
	if cwd != "" {
		cwdp = cstr(cwd)
	}
	return Process{c: own(fnLaunch(t.c.ptr(), cstrv(argv), cstrv(envp), cwdp), fnProcessDtor)}
}

// Triple returns the target architecture triple.
func (t Target) Triple() string { return goString(fnTriple(t.c.ptr())) }

// BreakpointCreateByName sets a breakpoint on the named symbol.
func (t Target) BreakpointCreateByName(name string) Breakpoint {
	return Breakpoint{c: own(fnBreakName(t.c.ptr(), cstr(name), nil), fnBreakpointDtor)}
}

// EvaluateExpression evaluates expr in the target and returns its value.
func (t Target) EvaluateExpression(expr string) Value {
	return Value{c: own(fnEval(t.c.ptr(), cstr(expr)), fnValueDtor)}
}

// Breakpoint wraps lldb::SBBreakpoint.
type Breakpoint struct{ c *object }

// Close releases the native breakpoint handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (b Breakpoint) Close() { b.c.close() }

// IsValid reports whether the breakpoint was created.
func (b Breakpoint) IsValid() bool { return fnBpValid(b.c.ptr()) }

// Process wraps lldb::SBProcess.
type Process struct{ c *object }

// Close releases the native process handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
// Close does not detach or kill the process; call Detach or Kill first.
func (p Process) Close() { p.c.close() }

// IsValid reports whether the process is attached and usable.
func (p Process) IsValid() bool { return fnProcValid(p.c.ptr()) }

// Continue resumes the process.
func (p Process) Continue() Error { return Error{c: own(fnContinue(p.c.ptr()), fnErrorDtor)} }

// SendAsyncInterrupt requests a stop without waiting for synchronous execution.
func (p Process) SendAsyncInterrupt() { fnInterrupt(p.c.ptr()) }

// Detach releases the process from debugger control and resumes it.
func (p Process) Detach() Error { return Error{c: own(fnDetach(p.c.ptr()), fnErrorDtor)} }

// Kill terminates the process.
func (p Process) Kill() Error { return Error{c: own(fnKill(p.c.ptr()), fnErrorDtor)} }

// State returns the current process state.
func (p Process) State() State { return State(fnState(p.c.ptr())) }

// NumThreads returns the number of threads.
func (p Process) NumThreads() uint32 { return fnNumThreads(p.c.ptr()) }

// ThreadAtIndex returns the thread at index i.
func (p Process) ThreadAtIndex(i uint32) Thread {
	return Thread{c: own(fnThreadAt(p.c.ptr(), uint64(i)), fnThreadDtor)}
}

// Thread wraps lldb::SBThread.
type Thread struct{ c *object }

// Close releases the native thread handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (t Thread) Close() { t.c.close() }

// StopReason returns why the thread stopped.
func (t Thread) StopReason() StopReason { return StopReason(fnStopReason(t.c.ptr())) }

// FrameAtIndex returns the stack frame at index i (0 is the top frame).
func (t Thread) FrameAtIndex(i uint32) Frame {
	return Frame{c: own(fnFrameAt(t.c.ptr(), i), fnFrameDtor)}
}

// StepOutOfFrame steps out of frame f.
func (t Thread) StepOutOfFrame(f Frame) { fnStepOut(t.c.ptr(), f.c.ptr()) }

// Frame wraps lldb::SBFrame.
type Frame struct{ c *object }

// Close releases the native frame handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (f Frame) Close() { f.c.close() }

// FindRegister returns the named register as a value.
func (f Frame) FindRegister(name string) Value {
	return Value{c: own(fnFindReg(f.c.ptr(), cstr(name)), fnValueDtor)}
}

// FunctionName returns the frame's function name, or "".
func (f Frame) FunctionName() string { return goString(fnFuncName(f.c.ptr())) }

// Value wraps lldb::SBValue.
type Value struct{ c *object }

// Close releases the native value handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (v Value) Close() { v.c.close() }

// IsValid reports whether the value is usable.
func (v Value) IsValid() bool { return fnValValid(v.c.ptr()) }

// Value returns the value's textual form (e.g. a register's hex value), or "".
func (v Value) Value() string { return goString(fnGetValue(v.c.ptr())) }

// SetValueFromCString sets the value from s, returning whether it succeeded.
func (v Value) SetValueFromCString(s string) bool { return fnSetCStr(v.c.ptr(), cstr(s)) }

// ObjectDescription returns the value's object description (Obj-C -description), or "".
func (v Value) ObjectDescription() string { return goString(fnObjDesc(v.c.ptr())) }

// ValueAsUnsigned returns the value as an unsigned integer, or fail on error.
func (v Value) ValueAsUnsigned(fail uint64) uint64 { return fnAsUnsig(v.c.ptr(), fail) }

// Error wraps lldb::SBError.
type Error struct{ c *object }

// Close releases the native error handle. It is safe to call more than once
// or on a zero value. Copies share the same handle and are all closed together.
func (e Error) Close() { e.c.close() }

// newError returns a constructed, empty SBError after load has succeeded.
func newError() Error {
	e := Error{c: own(cell{}, fnErrorDtor)}
	fnErrCtor(e.c.ptr())
	return e
}

// Success reports whether the error represents success.
func (e Error) Success() bool { return fnErrSuccess(e.c.ptr()) }

// String returns the error description.
func (e Error) String() string { return goString(fnErrCStr(e.c.ptr())) }

// destroyDebugger removes LLDB's global debugger registration before releasing
// the SB wrapper. The SB destructor alone only drops the wrapper's reference.
func destroyDebugger(p unsafe.Pointer) {
	fnDestroyDebugger(p)
	fnDebuggerDtor(p)
}
