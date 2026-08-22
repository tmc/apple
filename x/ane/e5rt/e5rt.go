//go:build darwin

package e5rt

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// FrameworkPath is the Espresso framework the e5rt_* symbols are exported from.
const FrameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"

// Symbols lists the e5rt_* entry points this package resolves, in the order of
// the five steps: compile, load, bind, dispatch. Each name has been observed to
// resolve on macOS 26.x; see the package documentation for what that does and
// does not establish.
var Symbols = []string{
	// Compile.
	"e5rt_e5_compiler_create_with_config",
	"e5rt_e5_compiler_compile",
	"e5rt_e5_compiler_is_new_compile_required",

	// Load.
	"e5rt_program_library_create",
	"e5rt_program_library_retain_program_function",
	"e5rt_program_function_load_for_execution",
	"e5rt_precompiled_compute_op_create_options_create_with_program_function",
	"e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options",

	// Bind.
	"e5rt_buffer_object_alloc",
	"e5rt_buffer_object_get_data_ptr",
	"e5rt_io_port_bind_buffer_object",
	"e5rt_execution_stream_operation_retain_input_port",
	"e5rt_execution_stream_operation_retain_output_port",

	// Dispatch.
	"e5rt_execution_stream_create",
	"e5rt_execution_stream_operation_prepare_op_for_encode",
	"e5rt_execution_stream_encode_operation",
	"e5rt_execution_stream_execute_sync",
	"e5rt_execution_stream_submit_async",
	"e5rt_execution_stream_reset",
}

// A Lib is the loaded Espresso framework with its e5rt_* entry points resolved.
// The zero value is not usable; obtain one from [Open]. A Lib is safe for
// concurrent use.
type Lib struct {
	handle uintptr
	syms   map[string]uintptr
}

var (
	openOnce sync.Once
	openLib  *Lib
	openErr  error
)

// Open loads the Espresso framework and resolves every name in [Symbols].
// It reports an error if the framework cannot be loaded, and reports the first
// unresolved symbol as an error while still returning the Lib, so a caller that
// needs only part of the route can proceed. The framework is loaded at most
// once per process.
func Open() (*Lib, error) {
	openOnce.Do(func() {
		h, err := purego.Dlopen(FrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			openErr = fmt.Errorf("load Espresso framework: %w", err)
			return
		}
		l := &Lib{handle: h, syms: make(map[string]uintptr, len(Symbols))}
		for _, name := range Symbols {
			sym, err := purego.Dlsym(h, name)
			if err != nil || sym == 0 {
				if openErr == nil {
					openErr = fmt.Errorf("resolve %s: %w", name, err)
				}
				continue
			}
			l.syms[name] = sym
		}
		openLib = l
	})
	return openLib, openErr
}

// Sym returns the address of a resolved e5rt_* entry point. Use it for symbols
// this package does not wrap, whose calling convention the caller must supply.
func (l *Lib) Sym(name string) (uintptr, error) {
	if l == nil {
		return 0, fmt.Errorf("e5rt: library not open")
	}
	sym, ok := l.syms[name]
	if !ok {
		return 0, fmt.Errorf("e5rt: symbol %s unresolved", name)
	}
	return sym, nil
}

// Resolved reports the names from [Symbols] that resolved.
func (l *Lib) Resolved() []string {
	if l == nil {
		return nil
	}
	names := make([]string, 0, len(l.syms))
	for _, name := range Symbols {
		if _, ok := l.syms[name]; ok {
			names = append(names, name)
		}
	}
	return names
}

// Status is the int64 error code every e5rt_* entry point returns. Zero is
// success. The meaning of nonzero values is not documented and has not been
// recovered; [Status.Err] reports the raw code.
type Status int64

// Err reports a non-nil error for any nonzero status.
func (s Status) Err(op string) error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("e5rt: %s: status %d", op, int64(s))
}

// call invokes a resolved entry point and returns its int64 status.
func (l *Lib) call(name string, args ...uintptr) (Status, error) {
	sym, err := l.Sym(name)
	if err != nil {
		return 0, err
	}
	r, _, _ := purego.SyscallN(sym, args...)
	return Status(int64(r)), nil
}

// callErr invokes a resolved entry point and folds its status into the error.
func (l *Lib) callErr(name string, args ...uintptr) error {
	st, err := l.call(name, args...)
	if err != nil {
		return err
	}
	return st.Err(name)
}

// Passing the address of a local through a uintptr is unsafe: the conversion
// hides the pointer from the compiler, which may then keep the value on the
// stack, where a stack growth during the call relocates it and leaves the
// callee writing into stale memory. The unsafe.Pointer rules exempt only
// conversions in the argument list of a recognized syscall.Syscall-shaped call,
// and purego.SyscallN is not one. newOut and cstring therefore force their
// storage to the heap, and every caller must runtime.KeepAlive it across the
// call so the collector cannot reclaim it while only a uintptr refers to it.

// newOut allocates an out-parameter cell on the heap. It is deliberately not
// inlined: returning the pointer from a non-inlinable function is what forces
// the escape.
//
//go:noinline
func newOut() *uintptr {
	return new(uintptr)
}

// cstring returns a NUL-terminated heap copy of s and its address. The caller
// must keep the returned slice alive across the call.
//
//go:noinline
func cstring(s string) ([]byte, uintptr) {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return b, uintptr(unsafe.Pointer(&b[0]))
}

// CompilerCreateWithConfig creates a compiler from an options handle and
// returns it.
//
// UNVERIFIED: the signature int64_t(void **compiler_out, void *config) is taken
// from the source paper's listing 6.1, not confirmed against the binary. config
// may be zero; the paper describes compiler options as a string-keyed
// dictionary of std::any values rather than a fixed-layout struct, and does not
// give a constructor for one.
func (l *Lib) CompilerCreateWithConfig(config uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_create_with_config", uintptr(unsafe.Pointer(out)), config)
	runtime.KeepAlive(out)
	return *out, err
}

// CompilerCompile compiles the network description at modelPath and returns the
// resulting program library. The paper describes modelPath as accepting the
// .espresso.net netplist representation alongside .mil; that has not been
// verified here.
//
// UNVERIFIED: the signature
// int64_t(compiler, const char *model_path, void *options, void **library_out)
// is taken from the paper, not confirmed against the binary.
func (l *Lib) CompilerCompile(compiler uintptr, modelPath string, options uintptr) (uintptr, error) {
	path, p := cstring(modelPath)
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_compile", compiler, p, options, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(path)
	runtime.KeepAlive(out)
	return *out, err
}

// ProgramLibraryRetainProgramFunction retains the callable function a compiled
// program exposes under fnName.
//
// UNVERIFIED: the signature
// int64_t(library, const char *fn_name, void **function_out) is taken from the
// paper, not confirmed against the binary.
func (l *Lib) ProgramLibraryRetainProgramFunction(library uintptr, fnName string) (uintptr, error) {
	name, p := cstring(fnName)
	out := newOut()
	err := l.callErr("e5rt_program_library_retain_program_function", library, p, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(name)
	runtime.KeepAlive(out)
	return *out, err
}

// ProgramFunctionLoadForExecution prepares a retained program function for
// execution on the device.
//
// UNVERIFIED: the single-argument signature int64_t(function) is taken from the
// paper's chapter 4 and 5 listings. The chapter 6 listings do not call it at
// all, so its place in the direct-route sequence is inferred from the paper's
// phase table rather than from a call site.
func (l *Lib) ProgramFunctionLoadForExecution(function uintptr) error {
	return l.callErr("e5rt_program_function_load_for_execution", function)
}

// BufferObjectAlloc allocates a buffer object of nbytes and returns it.
//
// UNVERIFIED: the signature int64_t(void **buf_out, size_t nbytes, int type) is
// taken from the paper, not confirmed against the binary. The paper passes zero
// for typ and does not say what other values mean.
func (l *Lib) BufferObjectAlloc(nbytes uintptr, typ int) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_buffer_object_alloc", uintptr(unsafe.Pointer(out)), nbytes, uintptr(typ))
	runtime.KeepAlive(out)
	return *out, err
}

// BufferObjectGetDataPtr returns the host address of a buffer object's storage.
//
// UNVERIFIED: the two-argument shape (buf, &ptr) appears only inside a comment
// in the paper's listings, never in a call. Both the argument order and the
// argument count are guesses.
func (l *Lib) BufferObjectGetDataPtr(buf uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_buffer_object_get_data_ptr", buf, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// OperationRetainInputPort retains the named input port of an operation.
//
// UNVERIFIED: the signature
// int64_t(op, const char *port_name, void **port_out) is taken from the paper,
// not confirmed against the binary. It is at least consistent across every
// listing in which it appears.
func (l *Lib) OperationRetainInputPort(op uintptr, portName string) (uintptr, error) {
	name, p := cstring(portName)
	out := newOut()
	err := l.callErr("e5rt_execution_stream_operation_retain_input_port", op, p, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(name)
	runtime.KeepAlive(out)
	return *out, err
}

// OperationRetainOutputPort retains the named output port of an operation.
//
// UNVERIFIED: same claimed shape as [Lib.OperationRetainInputPort], and the
// paper never calls it inside a chapter 6 listing.
func (l *Lib) OperationRetainOutputPort(op uintptr, portName string) (uintptr, error) {
	name, p := cstring(portName)
	out := newOut()
	err := l.callErr("e5rt_execution_stream_operation_retain_output_port", op, p, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(name)
	runtime.KeepAlive(out)
	return *out, err
}

// IOPortBindBufferObject binds a buffer object to a retained I/O port.
//
// UNVERIFIED: the signature int64_t(port, buffer_object) is taken from the
// paper. The paper further asserts that binding the same buffer to an input and
// an output port aliases it so state stays resident on the engine across steps;
// that behavior has not been observed here.
func (l *Lib) IOPortBindBufferObject(port, buf uintptr) error {
	return l.callErr("e5rt_io_port_bind_buffer_object", port, buf)
}

// ExecutionStreamCreate creates an execution stream.
//
// UNVERIFIED: the signature int64_t(void **stream_out) is taken from the paper.
func (l *Lib) ExecutionStreamCreate() (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_execution_stream_create", uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// PrepareOpForEncode prepares an operation before each encode.
//
// UNVERIFIED: the signature int64_t(op) is taken from the paper, which places
// this call inside the hot loop ahead of every encode.
func (l *Lib) PrepareOpForEncode(op uintptr) error {
	return l.callErr("e5rt_execution_stream_operation_prepare_op_for_encode", op)
}

// EncodeOperation encodes a prepared operation into a stream.
//
// UNVERIFIED: the signature int64_t(stream, op) is taken from the paper.
func (l *Lib) EncodeOperation(stream, op uintptr) error {
	return l.callErr("e5rt_execution_stream_encode_operation", stream, op)
}

// ExecuteSync submits an encoded stream and blocks until it completes.
//
// UNVERIFIED: the signature int64_t(stream) is taken from the paper, as is the
// claim that this form blocks.
func (l *Lib) ExecuteSync(stream uintptr) error {
	return l.callErr("e5rt_execution_stream_execute_sync", stream)
}

// ExecutionStreamReset resets a stream for reuse after execution.
//
// UNVERIFIED: the signature int64_t(stream) is taken from the paper.
func (l *Lib) ExecutionStreamReset(stream uintptr) error {
	return l.callErr("e5rt_execution_stream_reset", stream)
}

// Deliberately unwrapped, reachable through [Lib.Sym]:
//
//   - e5rt_e5_compiler_is_new_compile_required and e5rt_program_library_create
//     are named only in the paper's phase table, with no call site and no
//     argument list anywhere in the text.
//
//   - e5rt_precompiled_compute_op_create_options_create_with_program_function
//     and
//     e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options
//     are called with contradictory argument orders in different listings of
//     the same paper: out-parameter first in the chapter 6 listings, last
//     elsewhere. Wrapping either order would encode a coin flip as an API.
//     TODO: recover the true order from the binary and wrap them; without these
//     two the load phase of the route cannot be completed from Go.
//
//   - e5rt_execution_stream_submit_async is named but never called. The paper
//     describes three submission forms (one synchronous, a lightweight
//     asynchronous one returning submit and complete identifiers, and a full
//     asynchronous one taking a timeout) but gives only this one asynchronous
//     symbol, so which form it is remains unknown.
//     TODO: determine which submission form this symbol implements.
