//go:build darwin

package e5rt

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/objc"
)

// FrameworkPath is the Espresso framework the e5rt_* symbols are exported from.
const FrameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"

// Symbols lists the e5rt_* entry points this package resolves, in the order of
// the five steps: compile, load, bind, dispatch. Each name has been observed to
// resolve on macOS 26.x; see the package documentation for what that does and
// does not establish. Espresso exports many more e5rt_* names than this;
// [Lib.Lookup] reaches them.
var Symbols = []string{
	// Compile.
	"e5rt_e5_compiler_config_options_create",
	"e5rt_e5_compiler_config_options_set_cache_bundle_location",
	"e5rt_e5_compiler_config_options_release",
	"e5rt_e5_compiler_create_with_config",
	"e5rt_e5_compiler_release",
	"e5rt_e5_compiler_options_create",
	"e5rt_e5_compiler_options_set_compute_device_types_mask",
	"e5rt_e5_compiler_options_get_compute_device_types_mask",
	"e5rt_e5_compiler_options_set_force_recompilation",
	"e5rt_e5_compiler_options_set_segmenter",
	"e5rt_e5_compiler_options_set_custom_ane_compiler_options",
	"e5rt_e5_compiler_options_release",
	"e5rt_e5_compiler_compile",
	"e5rt_e5_compiler_is_new_compile_required",

	// Load.
	"e5rt_program_library_create",
	"e5rt_program_library_retain_program_function",
	"e5rt_program_library_release",
	"e5rt_program_function_load_for_execution",
	"e5rt_program_function_release",
	"e5rt_precompiled_compute_op_create_options_create_with_program_function",
	"e5rt_precompiled_compute_op_create_options_set_operation_name",
	"e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers",
	"e5rt_precompiled_compute_op_create_options_release",
	"e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options",
	"e5rt_execution_stream_operation_release",

	// Bind.
	"e5rt_buffer_object_alloc",
	"e5rt_buffer_object_get_data_ptr",
	"e5rt_buffer_object_release",
	"e5rt_io_port_bind_buffer_object",
	"e5rt_io_port_release",
	"e5rt_execution_stream_operation_retain_input_port",
	"e5rt_execution_stream_operation_retain_inout_port",
	"e5rt_execution_stream_operation_retain_output_port",

	// Dispatch.
	"e5rt_execution_stream_create",
	"e5rt_execution_stream_operation_prepare_op_for_encode",
	"e5rt_execution_stream_encode_operation",
	"e5rt_execution_stream_execute_sync",
	"e5rt_execution_stream_submit_async",
	"e5rt_execution_stream_reset",
	"e5rt_execution_stream_release",
	"e5rt_execution_stream_set_quality_of_service",
	"e5rt_execution_stream_set_ane_execution_priority",

	// Events.
	"e5rt_async_event_create",
	"e5rt_async_event_release",
	"e5rt_async_event_signal",
	"e5rt_async_event_sync_wait",
	"e5rt_async_event_get_last_signaled_value",
	"e5rt_async_event_get_active_future_value",
	"e5rt_async_event_set_active_future_value",
	"e5rt_execution_stream_operation_bind_completion_event",
	"e5rt_execution_stream_operation_bind_dependent_events",
}

// Compute device bits for [Lib.CompilerOptionsSetComputeDeviceTypesMask]. More
// than one bit lets the compiler choose; ANEForge reports that 0x3 selects BNNS
// and that any mask including [ComputeDeviceANE] selects the Neural Engine,
// falling back to BNNS when the Neural Engine compile fails.
//
// The values and that selection behavior are ANEForge's empirical finding
// (aneforge/_lib/e5rt_api.h:46-50, docs/e5rt-dispatch-reference.md:238-240),
// not an observation made here.
const (
	ComputeDeviceCPU uint64 = 0x1 // BNNS
	ComputeDeviceGPU uint64 = 0x2 // MPSGraph
	ComputeDeviceANE uint64 = 0x4
)

// A Lib is the loaded Espresso framework with its e5rt_* entry points resolved.
// The zero value is not usable; obtain one from [Open]. A Lib is safe for
// concurrent use.
type Lib struct {
	handle uintptr
	syms   map[string]uintptr // the names in Symbols, resolved by Open and never written again

	mu    sync.RWMutex
	extra map[string]uintptr // names resolved on demand by Lookup

	errorStringOnce sync.Once
	errorString     func(int32) uintptr
	errorStringErr  error
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
		l := &Lib{
			handle: h,
			syms:   make(map[string]uintptr, len(Symbols)),
			extra:  make(map[string]uintptr),
		}
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

// Sym returns the address of an entry point named in [Symbols] and resolved by
// [Open]. It is a map lookup, not a live dlsym, so it reaches nothing outside
// that list; use [Lib.Lookup] for the rest of the framework. Use it for symbols
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

// Lookup resolves any exported symbol of the loaded framework by name, whether
// or not it appears in [Symbols], and caches the result. Espresso exports on the
// order of two hundred e5rt_* names and this package lists only the ones it
// documents, so a caller driving a part of the route this package does not cover
// would otherwise have to dlopen the framework a second time. The caller must
// supply the calling convention, which this package makes no claim about.
//
// Use [Lib.Sym] instead when you want the lookup restricted to [Symbols], as a
// probe reporting on that list does.
func (l *Lib) Lookup(name string) (uintptr, error) {
	if l == nil {
		return 0, fmt.Errorf("e5rt: library not open")
	}
	if sym, ok := l.syms[name]; ok {
		return sym, nil
	}
	l.mu.RLock()
	sym, ok := l.extra[name]
	l.mu.RUnlock()
	if ok {
		return sym, nil
	}
	sym, err := purego.Dlsym(l.handle, name)
	if err != nil || sym == 0 {
		return 0, fmt.Errorf("e5rt: resolve %s: %w", name, err)
	}
	l.mu.Lock()
	l.extra[name] = sym
	l.mu.Unlock()
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

// Status is a normalized E5RT status code. E5RT declares its status carrier as
// a signed 32-bit C int. A zero code means the entry point accepted the call,
// not necessarily that it had its apparent effect. The meaning of nonzero
// values is not documented and has not been recovered; [Status.Err] reports
// the raw code.
type Status int64

// Err reports a non-nil error for any nonzero status.
func (s Status) Err(op string) error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("e5rt: %s: status %d", op, int64(s))
}

// ErrorString returns Espresso's static description of status.
//
// The recovered ane_bridge declaration says e5rt_error_code_get_string takes a
// signed 32-bit status and returns const char *. Its tests read static descriptions
// for status values 0 through 6. This method reads the returned C string; it
// does not retain it.
func (l *Lib) ErrorString(status Status) (string, error) {
	if l == nil {
		return "", fmt.Errorf("e5rt: library not open")
	}
	l.errorStringOnce.Do(func() {
		address, err := l.Lookup("e5rt_error_code_get_string")
		if err != nil {
			l.errorStringErr = err
			return
		}
		purego.RegisterFunc(&l.errorString, address)
	})
	if l.errorStringErr != nil {
		return "", l.errorStringErr
	}
	ptr := l.errorString(int32(status))
	if ptr == 0 {
		return "", fmt.Errorf("e5rt: no description for status %d", status)
	}
	return staticCString(ptr), nil
}

func staticCString(ptr uintptr) string {
	const maxLength = 4 << 10
	data := unsafe.Slice((*byte)(pointerAt(ptr)), maxLength)
	for n, b := range data {
		if b == 0 {
			return string(data[:n])
		}
	}
	return string(data)
}

// call invokes a resolved entry point and returns its int64 status.
func (l *Lib) call(name string, args ...uintptr) (Status, error) {
	sym, err := l.Sym(name)
	if err != nil {
		return 0, err
	}
	r, _, _ := purego.SyscallN(sym, args...)
	return normalizeStatus(int64(r)), nil
}

func normalizeStatus(raw int64) Status {
	return Status(int32(raw))
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

// newRef allocates a heap cell holding v, for the e5rt_*_release entry points,
// which take the address of the handle rather than the handle. See newOut for
// why the cell must be on the heap.
//
//go:noinline
func newRef(v uintptr) *uintptr {
	p := new(uintptr)
	*p = v
	return p
}

// boolArg encodes a Go bool as the C int the setters take.
func boolArg(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// release invokes an e5rt_*_release entry point on handle. The callee is passed
// a temporary cell, so if it clears the handle the caller's copy still holds the
// old value; treat a released handle as dead.
func (l *Lib) release(name string, handle uintptr) error {
	ref := newRef(handle)
	err := l.callErr(name, uintptr(unsafe.Pointer(ref)))
	runtime.KeepAlive(ref)
	return err
}

// What the citations below do and do not establish. Most wrappers name an
// ANEForge call site or typedef and say nothing else. That is provenance: it
// records where the argument list came from, not that it is right. What makes
// most of them right is that the compile, load, bind and dispatch wrappers are
// driven end to end by the e5rtdispatch example, whose output matches a float64
// CPU reference and moves when an input moves, and the rest by
// TestUnverifiedCalls. An argument list that were wrong in position or width
// would fault in C or return a nonzero status well before that.
//
// So a bare citation on a wrapper on that route is corroborated by the route
// running. A citation is the only evidence when the comment says so, and several
// do: the two scheduling setters, the custom Neural Engine compiler options, and
// the event family's progress calls are each exercised without any observable
// effect, which is not the same as being understood.
//
// Argument order across this family follows one rule, taken from the call sites
// in ANEForge's ane_e5rt_dispatch.mm: an entry point that creates or allocates
// an object takes the out-parameter FIRST, and every other entry point takes the
// object it acts on first and any out-parameter LAST. So
// e5rt_program_library_retain_program_function(library, "main", &function)
// (ane_e5rt_dispatch.mm:370) sits eight lines above
// e5rt_precompiled_compute_op_create_options_create_with_program_function(&op_options, function)
// (:373). The file-level comment in ANEForge's e5rt_api.h:17-18 states the
// opposite ("out goes LAST"); it is contradicted by three of that file's own
// per-symbol comments (:64, :76, :84) and by every call site, and is not
// followed here.

// CompilerConfigOptionsCreate creates a compiler configuration and returns it.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:347): out-parameter first, no other
// argument. Called from Go on macOS 26.x as the first step of a compile that
// went on to produce correct results; see the e5rtdispatch example.
func (l *Lib) CompilerConfigOptionsCreate() (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_config_options_create", uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// CompilerConfigOptionsSetCacheBundleLocation sets the directory the compiler
// writes its compiled bundle into.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:349): (config, const char *path).
func (l *Lib) CompilerConfigOptionsSetCacheBundleLocation(config uintptr, dir string) error {
	path, p := cstring(dir)
	err := l.callErr("e5rt_e5_compiler_config_options_set_cache_bundle_location", config, p)
	runtime.KeepAlive(path)
	return err
}

// CompilerConfigOptionsRelease releases a compiler configuration. ANEForge
// releases it only after the operation exists (ane_e5rt_dispatch.mm:388).
func (l *Lib) CompilerConfigOptionsRelease(config uintptr) error {
	return l.release("e5rt_e5_compiler_config_options_release", config)
}

// CompilerCreateWithConfig creates a compiler from a configuration handle and
// returns it. Build the configuration with [Lib.CompilerConfigOptionsCreate].
//
// LOCAL OBSERVATION: passing config == 0 returns status 1 on macOS 26.x;
// creating a configuration first and passing it returns 0. A zero config is
// therefore not accepted, contrary to the source paper.
//
// The signature int64_t(void **compiler_out, void *config) is the paper's
// listing 6.1 and agrees with ANEForge (e5rt_api.h:40, call site
// ane_e5rt_dispatch.mm:351).
func (l *Lib) CompilerCreateWithConfig(config uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_create_with_config", uintptr(unsafe.Pointer(out)), config)
	runtime.KeepAlive(out)
	return *out, err
}

// CompilerRelease releases a compiler (ane_e5rt_dispatch.mm:387).
func (l *Lib) CompilerRelease(compiler uintptr) error {
	return l.release("e5rt_e5_compiler_release", compiler)
}

// CompilerOptionsCreate creates the per-compile options object
// [Lib.CompilerCompile] takes, and returns it.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:354): out-parameter first.
func (l *Lib) CompilerOptionsCreate() (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_options_create", uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// CompilerOptionsSetComputeDeviceTypesMask selects which backends the compiler
// may target. Pass a bitwise OR of [ComputeDeviceCPU], [ComputeDeviceGPU] and
// [ComputeDeviceANE].
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:356): (options, uint64_t mask).
func (l *Lib) CompilerOptionsSetComputeDeviceTypesMask(options uintptr, mask uint64) error {
	return l.callErr("e5rt_e5_compiler_options_set_compute_device_types_mask", options, uintptr(mask))
}

// CompilerOptionsGetComputeDeviceTypesMask reports the mask set on options.
//
// ANEForge TYPEDEF (e5rt_api.h:55): (options, uint64_t *out). ANEForge does not
// call it, so the typedef was all the placement rested on. It no longer is: the
// mask round-trips through [Lib.CompilerOptionsSetComputeDeviceTypesMask] for
// four distinct values, which a wrong out-parameter position would not do. See
// TestComputeDeviceMaskRoundTrips.
func (l *Lib) CompilerOptionsGetComputeDeviceTypesMask(options uintptr) (uint64, error) {
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_options_get_compute_device_types_mask", options, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return uint64(*out), err
}

// CompilerOptionsSetForceRecompilation makes the compiler ignore any cached
// bundle. ANEForge sets it on every compile (ane_e5rt_dispatch.mm:357).
func (l *Lib) CompilerOptionsSetForceRecompilation(options uintptr, force bool) error {
	return l.callErr("e5rt_e5_compiler_options_set_force_recompilation", options, boolArg(force))
}

// CompilerOptionsSetSegmenter selects how the compiler partitions the network.
// ANEForge passes "graph" (ane_e5rt_dispatch.mm:358); no other value is known.
func (l *Lib) CompilerOptionsSetSegmenter(options uintptr, segmenter string) error {
	s, p := cstring(segmenter)
	err := l.callErr("e5rt_e5_compiler_options_set_segmenter", options, p)
	runtime.KeepAlive(s)
	return err
}

// CompilerOptionsSetCustomANECompilerOptions passes a string through to the
// Neural Engine compiler. ANEForge uses it for cross-target compile checks
// (ane_e5rt_dispatch.mm:441) and its Python side gives the syntax as
// TargetArchitecture=h13 (_runtime.py:102).
//
// No effect has been observed here. The call returns zero for every string
// tried, and the compile that follows succeeds either way: with the documented
// spelling, with an architecture that does not exist, with an option name that
// does not exist, and with a string that is not an assignment at all, each
// against its own cache directory so no result is a previous compile's. Since
// the negative controls do not fail, nothing separates this call reaching the
// Neural Engine compiler from it being ignored, and a caller should not treat a
// successful compile as evidence the option took effect.
func (l *Lib) CompilerOptionsSetCustomANECompilerOptions(options uintptr, custom string) error {
	s, p := cstring(custom)
	err := l.callErr("e5rt_e5_compiler_options_set_custom_ane_compiler_options", options, p)
	runtime.KeepAlive(s)
	return err
}

// CompilerOptionsRelease releases a compiler options object
// (ane_e5rt_dispatch.mm:386).
func (l *Lib) CompilerOptionsRelease(options uintptr) error {
	return l.release("e5rt_e5_compiler_options_release", options)
}

// CompilerCompile compiles the network description at modelPath and returns the
// resulting program library. The paper describes modelPath as accepting the
// .espresso.net netplist representation alongside .mil; that has not been
// verified here. ANEForge passes a MIL path (ane_e5rt_dispatch.mm:363) and
// names the parameter mil_or_mlmodelc_path (e5rt_api.h:42).
//
// Build options with [Lib.CompilerOptionsCreate]. This is a method on an
// existing compiler, so the out-parameter comes last.
//
// The signature
// int64_t(compiler, const char *model_path, void *options, void **library_out)
// is the paper's and agrees with ANEForge (e5rt_api.h:42, call site :363).
func (l *Lib) CompilerCompile(compiler uintptr, modelPath string, options uintptr) (uintptr, error) {
	path, p := cstring(modelPath)
	out := newOut()
	err := l.callErr("e5rt_e5_compiler_compile", compiler, p, options, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(path)
	runtime.KeepAlive(out)
	return *out, err
}

// ProgramLibraryCreate opens an already-compiled bundle at bundlePath and
// returns the program library it holds. The bundle must have been compiled by
// this process: ANEForge reports (docs/e5rt-dispatch-reference.md:305-309) that
// the signed program lives only in aned's per-PID cache, so loading a bundle in
// a fresh process fails.
//
// ANEForge TYPEDEF plus an empirical note (e5rt_api.h:64-65): the out-parameter
// comes first, and the reversed order returns "Invalid E5 path specified. @
// GetE5PathFromCompositeBundle". ANEForge resolves the symbol but does not call
// it, so the note is its author's probing, not a call site.
//
// Confirmed here on macOS 26.x. Given the .bundle directory the compiler leaves
// under its cache location, this returns a library handle from which "main" can
// be retained, with no compiler involved. See TestUnverifiedCalls.
//
// The reuse is across processes too, which contradicts the report above. One
// process compiles and exits; a second, started afterwards, opens the bundle,
// creates the operation, encodes and dispatches, and the output matches the
// reference. Three controls make that a measurement rather than a coincidence.
// The compiler emitted only main_ane, so the program being reused is an engine
// program and not a BNNS fallback that would return the same correct answer from
// the CPU. The model.mil and its weights are deleted after the compile and
// before the second process runs, so no silent recompile can explain it. And the
// same sequence inside a single process is run as a control, because a failure
// there would mean the cross-process arm had measured nothing.
//
// Note what this does not say. It was measured on macOS 26.x with both processes
// sharing a code-signing identity and a parent, which is the case ANEForge says
// needs posix_spawn'd children. Whether a bundle survives an aned restart, a
// reboot, or an unrelated process is untested. See TestCompiledBundleAcrossProcesses.
func (l *Lib) ProgramLibraryCreate(bundlePath string) (uintptr, error) {
	path, p := cstring(bundlePath)
	out := newOut()
	err := l.callErr("e5rt_program_library_create", uintptr(unsafe.Pointer(out)), p)
	runtime.KeepAlive(path)
	runtime.KeepAlive(out)
	return *out, err
}

// ProgramLibraryRelease releases a program library (ane_e5rt_dispatch.mm:334).
func (l *Lib) ProgramLibraryRelease(library uintptr) error {
	return l.release("e5rt_program_library_release", library)
}

// ProgramLibraryRetainProgramFunction retains the callable function a compiled
// program exposes under fnName. ANEForge always asks for "main"
// (ane_e5rt_dispatch.mm:370).
//
// The signature int64_t(library, const char *fn_name, void **function_out) is
// the paper's and agrees with ANEForge (e5rt_api.h:69, call site :370). It is a
// method on an existing library, so the out-parameter comes last.
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
// The single-argument signature int64_t(function) is the paper's chapter 4 and
// 5 listings and agrees with ANEForge's typedef (e5rt_api.h:72). Neither the
// paper's chapter 6 listings nor ANEForge calls it, and ANEForge's documented
// sequence goes straight from retaining the function to creating the operation
// options.
//
// It is gone. On macOS 26.x it returns status 2 and prints:
//
//	ProgramFunction LoadForExecution() is no longer supported.
//	Switch to ExecutionStreamOperation.
//
// The route the rest of this package drives is that replacement, and it works
// without this call, so the answer to where it belonged in the sequence is that
// it no longer belongs anywhere. Kept as a wrapper only so the status and the
// message are discoverable. See TestUnverifiedCalls.
func (l *Lib) ProgramFunctionLoadForExecution(function uintptr) error {
	return l.callErr("e5rt_program_function_load_for_execution", function)
}

// ProgramFunctionRelease releases a retained program function
// (ane_e5rt_dispatch.mm:332).
func (l *Lib) ProgramFunctionRelease(function uintptr) error {
	return l.release("e5rt_program_function_release", function)
}

// PrecompiledComputeOpOptionsCreate creates the options object that describes an
// operation built from a loaded program function, and returns it.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:373): (&op_options, function),
// out-parameter FIRST. The source paper gives this call with the out-parameter
// in both positions in different listings; ANEForge's e5rt_api.h:76 adds that
// the reversed order fails with "Cannot provide program function as nullptr. @
// Create". That settles the order without a local observation.
func (l *Lib) PrecompiledComputeOpOptionsCreate(function uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_precompiled_compute_op_create_options_create_with_program_function",
		uintptr(unsafe.Pointer(out)), function)
	runtime.KeepAlive(out)
	return *out, err
}

// PrecompiledComputeOpOptionsSetOperationName names the operation. ANEForge
// passes the program function's name, "main" (ane_e5rt_dispatch.mm:375).
func (l *Lib) PrecompiledComputeOpOptionsSetOperationName(options uintptr, name string) error {
	s, p := cstring(name)
	err := l.callErr("e5rt_precompiled_compute_op_create_options_set_operation_name", options, p)
	runtime.KeepAlive(s)
	return err
}

// PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers asks the runtime to
// allocate the operation's internal buffers. ANEForge passes true
// (ane_e5rt_dispatch.mm:376); the effect of false is not documented anywhere
// consulted here.
func (l *Lib) PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(options uintptr, allocate bool) error {
	return l.callErr("e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers",
		options, boolArg(allocate))
}

// PrecompiledComputeOpOptionsRelease releases an operation options object
// (ane_e5rt_dispatch.mm:330).
func (l *Lib) PrecompiledComputeOpOptionsRelease(options uintptr) error {
	return l.release("e5rt_precompiled_compute_op_create_options_release", options)
}

// OperationCreatePrecompiled creates the executable operation from an options
// object built by [Lib.PrecompiledComputeOpOptionsCreate], and returns it.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:378): (&operation, op_options),
// out-parameter FIRST, matching the typedef at e5rt_api.h:85. As with
// [Lib.PrecompiledComputeOpOptionsCreate], the paper gives both orders.
func (l *Lib) OperationCreatePrecompiled(options uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options",
		uintptr(unsafe.Pointer(out)), options)
	runtime.KeepAlive(out)
	return *out, err
}

// OperationRelease releases an operation (ane_e5rt_dispatch.mm:328).
func (l *Lib) OperationRelease(op uintptr) error {
	return l.release("e5rt_execution_stream_operation_release", op)
}

// BufferObjectAlloc allocates a buffer object of nbytes and returns it.
//
// typ selects the backing store:
//
//	0  host memory; [Lib.BufferObjectGetDataPtr] yields a plain CPU address
//	1  Neural Engine mapped
//	2  IOSurface backed
//
// Anything else is rejected with "Invalid BufferType @ AllocMemory". ANEForge
// passes 0 everywhere (ane_e5rt_dispatch.mm:278); the meanings above are its
// empirical notes (e5rt_api.h:99-104,
// docs/e5rt-dispatch-reference.md:235-237), not observations made here, and the
// note for 1 carries the author's own question mark.
//
// The signature int64_t(void **buf_out, size_t nbytes, uint32_t type) is the
// paper's and agrees with ANEForge (e5rt_api.h:104, call site :278).
func (l *Lib) BufferObjectAlloc(nbytes uintptr, typ int) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_buffer_object_alloc", uintptr(unsafe.Pointer(out)), nbytes, uintptr(typ))
	runtime.KeepAlive(out)
	return *out, err
}

// BufferObjectGetDataPtr returns the host address of a buffer object's storage.
//
// LOCAL OBSERVATION: called on macOS 26.x, this returns a pointer backed by
// writable memory. The two-argument shape (buf, &ptr), which the paper gives
// only inside a comment, is confirmed by ANEForge (e5rt_api.h:106, call site
// ane_e5rt_dispatch.mm:286).
func (l *Lib) BufferObjectGetDataPtr(buf uintptr) (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_buffer_object_get_data_ptr", buf, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// BufferObjectRelease releases a buffer object (ane_e5rt_dispatch.mm:312).
func (l *Lib) BufferObjectRelease(buf uintptr) error {
	return l.release("e5rt_buffer_object_release", buf)
}

// OperationRetainInputPort retains the named input port of an operation.
//
// The signature int64_t(op, const char *port_name, void **port_out) is the
// paper's, consistent across every listing in which it appears, and agrees with
// ANEForge (e5rt_api.h:87, call site ane_e5rt_dispatch.mm:267). A retained
// port from an engine-borrowed operation is engine-co-owned; do not release it.
func (l *Lib) OperationRetainInputPort(op uintptr, portName string) (uintptr, error) {
	name, p := cstring(portName)
	out := newOut()
	err := l.callErr("e5rt_execution_stream_operation_retain_input_port", op, p, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(name)
	runtime.KeepAlive(out)
	return *out, err
}

// OperationRetainInoutPort retains the named inout port of an operation.
//
// An inout is one port backed by one caller-supplied object for both its read
// and its write. It is distinct from binding the same object independently to
// an input and output port. The three-argument convention matches the input
// and output port calls: operation, NUL-terminated name, and out-parameter.
// Its first use is kept in a subprocess probe because E5RT failures can throw
// C++ exceptions across purego.
func (l *Lib) OperationRetainInoutPort(op uintptr, portName string) (uintptr, error) {
	name, p := cstring(portName)
	out := newOut()
	err := l.callErr("e5rt_execution_stream_operation_retain_inout_port", op, p, uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(name)
	runtime.KeepAlive(out)
	return *out, err
}

// OperationRetainOutputPort retains the named output port of an operation.
//
// Same shape as [Lib.OperationRetainInputPort]. The paper never calls it inside
// a chapter 6 listing; ANEForge does (e5rt_api.h:88, call site
// ane_e5rt_dispatch.mm:269). A retained port from an engine-borrowed operation
// is engine-co-owned; do not release it.
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
// The signature int64_t(port, buffer_object) is the paper's and agrees with
// ANEForge (e5rt_api.h:96, call site ane_e5rt_dispatch.mm:294). The paper
// further asserts that binding the same buffer to an input and an output port
// aliases it so state stays resident on the engine across steps; ANEForge builds
// its resident-state training path on exactly that
// (docs/e5rt-dispatch-reference.md:156-191, call site
// ane_e5rt_dispatch.mm:696). Neither has been reproduced here.
func (l *Lib) IOPortBindBufferObject(port, buf uintptr) error {
	return l.callErr("e5rt_io_port_bind_buffer_object", port, buf)
}

// IOPortRelease releases a retained I/O port from this package's standalone
// direct route.
//
// [Compile] and [CompilePipeline] create their own operation and stream, so
// their Close methods release their ports. Do not call IOPortRelease for a port
// retained from an engine-borrowed operation: that flow co-owns the port and
// requires callers to reuse it without releasing it.
func (l *Lib) IOPortRelease(port uintptr) error {
	return l.release("e5rt_io_port_release", port)
}

// ExecutionStreamCreate creates an execution stream.
//
// The signature int64_t(void **stream_out) is the paper's and agrees with
// ANEForge (e5rt_api.h:111, call site ane_e5rt_dispatch.mm:592).
func (l *Lib) ExecutionStreamCreate() (uintptr, error) {
	out := newOut()
	err := l.callErr("e5rt_execution_stream_create", uintptr(unsafe.Pointer(out)))
	runtime.KeepAlive(out)
	return *out, err
}

// PrepareOpForEncode resets an operation to its ready-to-encode state.
//
// The signature int64_t(op) is the paper's and agrees with ANEForge
// (e5rt_api.h:89). The paper places the call inside the hot loop ahead of every
// encode; ANEForge reports it is legal only on an operation that has already
// been encoded once, and so calls it only when re-encoding a used stream
// (ane_e5rt_dispatch.mm:460-464, call site :483).
//
// On this package's self-created direct stream, it returns zero in every
// position tried and leaves nothing usable behind. Called on an operation that
// has never been encoded, it returns zero and the ordinary encode and dispatch
// still work, so ANEForge's rejection claim does not hold as a status. Called
// on an encoded operation it also returns zero, and afterwards re-encoding it
// on the same stream fails with status 2, resetting that stream fails with
// status 2, and releasing that stream terminates the process:
//
//	libc++abi: terminating due to uncaught exception of type E5RT::E5RTError:
//	Op has not been encoded and hence cannot be reset to "ReadyForEncode" state
//
// So ANEForge's error string is real, but it arrives as a C++ exception at
// release time rather than as a status from this call. An exception thrown
// through a purego call cannot be caught in Go, which is why this wrapper
// reports success for a call that has already made the stream unreleasable.
// Do not use this wrapper with a self-created direct stream; encode each
// operation once.
//
// A separate engine-borrowed route reports this call works after resetting the
// borrowed stream, then retaining and binding its engine-co-owned I/O ports.
// That route is not reproduced here and has different port ownership, so it is
// intentionally not exposed through [Program] or [Pipeline].
func (l *Lib) PrepareOpForEncode(op uintptr) error {
	return l.callErr("e5rt_execution_stream_operation_prepare_op_for_encode", op)
}

// EncodeOperation encodes a prepared operation into a stream.
//
// The signature int64_t(stream, op) is the paper's and agrees with ANEForge
// (e5rt_api.h:113, call site ane_e5rt_dispatch.mm:488). ANEForge encodes each
// operation once, in submission order.
func (l *Lib) EncodeOperation(stream, op uintptr) error {
	return l.callErr("e5rt_execution_stream_encode_operation", stream, op)
}

// ExecuteSync submits an encoded stream and blocks until it completes.
//
// The signature int64_t(stream) is the paper's and agrees with ANEForge
// (e5rt_api.h:114, call site ane_e5rt_dispatch.mm:664), which reports this as
// the production path and says it serializes every encoded operation in the
// stream. Both hold here on macOS 26.x: the call returns having filled the
// bound output buffer, and the result matches a CPU reference, so it had
// completed rather than merely been queued.
func (l *Lib) ExecuteSync(stream uintptr) error {
	return l.callErr("e5rt_execution_stream_execute_sync", stream)
}

// ExecutionStreamReset resets a stream for reuse after execution.
//
// The signature int64_t(stream) is the paper's and agrees with ANEForge
// (ane_e5rt_dispatch.mm:99, call site :479).
//
// ANEForge reports that this call rejects a stream that has not been executed
// yet (:464). That is NOT reproduced here: on macOS 26.x a freshly created
// stream is reset successfully, twice in a row, both returning zero. The
// neighbouring claim in the same comment, about [Lib.PrepareOpForEncode] on a
// self-created stream, is also not reproduced; see that wrapper.
//
// This does fail with status 2 on a stream whose operation has been passed to
// [Lib.PrepareOpForEncode] on a self-created stream, which is one reason not
// to call it in that route.
func (l *Lib) ExecutionStreamReset(stream uintptr) error {
	return l.callErr("e5rt_execution_stream_reset", stream)
}

// ExecutionStreamRelease releases an execution stream
// (ane_e5rt_dispatch.mm:506).
func (l *Lib) ExecutionStreamRelease(stream uintptr) error {
	return l.release("e5rt_execution_stream_release", stream)
}

// SubmitAsync submits an encoded stream without blocking and returns a function
// that releases the completion block.
//
// completion runs on a thread E5RT owns once the work finishes, and the Status
// it returns is handed back to the runtime; return 0 unless there is a reason
// not to. It must not block, and it runs after SubmitAsync has returned,
// so a caller that needs the result waits on something the closure signals.
//
// Call the returned release function only after completion has run. The callee
// retains the block, so releasing early does not free it, but nothing else keeps
// the Go closure reachable.
//
// The two-argument signature, the second an Objective-C block of type
// e5rt_error_code_t (^)(void), comes from ANEForge, which settles it from the
// disassembly and drives it with a real block (ane_e5rt_dispatch.mm:88-92,
// :891-897, :933-949). A null second argument crashes inside objc_retain,
// because the callee retains unconditionally, and a raw purego callback where a
// block is expected is a segfault. [objc.NewBlock] builds a real
// __NSMallocBlock__ with copy and dispose helpers, which is what makes this
// wrappable; an unconditional retain is safe against it, and the Go closure
// stays reachable until dispose.
//
// ANEForge notes that the similarly named e5rt_execution_stream_async_submit is
// an older entry point that answers "Use submit_async".
func (l *Lib) SubmitAsync(stream uintptr, completion func() Status) (release func(), err error) {
	if completion == nil {
		return nil, fmt.Errorf("e5rt: submit_async requires a completion function")
	}
	block := objc.NewBlock(func(objc.Block) int64 {
		return int64(completion())
	})
	if block == 0 {
		return nil, fmt.Errorf("e5rt: could not build the completion block")
	}
	if err := l.callErr("e5rt_execution_stream_submit_async", stream, uintptr(block)); err != nil {
		block.Release()
		return nil, err
	}
	return func() { block.Release() }, nil
}

// ExecutionStreamSetQualityOfService sets the stream's dispatch quality of
// service.
//
// The signature int64_t(stream, uint64_t) is ANEForge's (e5rt_api.h:116). It
// has no call site there, so the meaning of the value is outside evidence;
// what is observed here is only that the call returns zero. Passing a value
// changes nothing measurable in this package's tests.
func (l *Lib) ExecutionStreamSetQualityOfService(stream uintptr, qos uint64) error {
	return l.callErr("e5rt_execution_stream_set_quality_of_service", stream, uintptr(qos))
}

// ExecutionStreamSetANEExecutionPriority sets the stream's priority on the
// Neural Engine.
//
// The signature int64_t(stream, uint64_t) is ANEForge's (e5rt_api.h:117), with
// the same caveat as [Lib.ExecutionStreamSetQualityOfService]: no call site,
// and the value's meaning is unrecovered.
func (l *Lib) ExecutionStreamSetANEExecutionPriority(stream uintptr, priority uint64) error {
	return l.callErr("e5rt_execution_stream_set_ane_execution_priority", stream, uintptr(priority))
}

// AsyncEventCreate creates an async event with the given name, which must not
// be empty; the callee reports a NULL name as an error. The event starts at
// zero, and [Lib.AsyncEventLastSignaledValue] counts up from there.
//
// The name is a label, not a key. Two events created with the same name while
// both are live are two distinct events with distinct handles.
//
// ANEForge CALL SITE (ane_e5rt_dispatch.mm:584): out-parameter first, then the
// name, then a third argument it always passes as 0 and describes as an initial
// value. That description is not confirmed here and this wrapper does not expose
// it: on macOS 26.x every nonzero third argument tried (1, 2, 7, each with a
// fresh name, so a name collision is ruled out) is rejected with status 1, and 0
// succeeds. Since no value that would distinguish an initial count from a flags
// word is accepted, the parameter's meaning is unrecovered and only its one
// working value is passed.
func (l *Lib) AsyncEventCreate(name string) (uintptr, error) {
	if name == "" {
		return 0, fmt.Errorf("e5rt: async event requires a name")
	}
	out := newOut()
	buf, p := cstring(name)
	err := l.callErr("e5rt_async_event_create", uintptr(unsafe.Pointer(out)), p, 0)
	event := *out
	runtime.KeepAlive(out)
	runtime.KeepAlive(buf)
	return event, err
}

// AsyncEventRelease releases an async event (ane_e5rt_dispatch.mm:326). Like
// the other release entry points it takes the address of the handle.
func (l *Lib) AsyncEventRelease(event uintptr) error {
	return l.release("e5rt_async_event_release", event)
}

// AsyncEventSignal signals event with value.
//
// The signature is int64_t(event, uint64_t value). A child-process probe
// signals fresh events with 7 and 31 and reads those exact values back. The
// old one-argument wrapper returned status 2 and did not signal anything.
func (l *Lib) AsyncEventSignal(event uintptr, value uint64) error {
	return l.callErr("e5rt_async_event_signal", event, uintptr(value))
}

// AsyncEventSyncWait waits for event to reach value, for at most timeoutNS.
//
// The recovered ane_bridge declaration gives the three-argument signature
// int64_t(event, uint64_t value, uint64_t timeout_ns). Its own test verifies a
// wait for an already-signaled value returns promptly. This package has not
// driven an unmet wait or a timeout, so it does not establish blocking behavior.
func (l *Lib) AsyncEventSyncWait(event uintptr, value, timeoutNS uint64) error {
	return l.callErr("e5rt_async_event_sync_wait", event, uintptr(value), uintptr(timeoutNS))
}

// AsyncEventLastSignaledValue reports the event's last signaled value.
//
// The signature int64_t(event, uint64_t *out) is ANEForge's (e5rt_api.h:126,
// call site ane_e5rt_dispatch.mm:876): object first, out-parameter last, which
// is this package's rule for a method on an existing object.
//
// A child-process probe verifies that [Lib.AsyncEventSignal] moves this value.
// It separately checks whether an operation's completion event advances under
// [Lib.SubmitAsync].
//
// What does work is the shape of the graph rather than its progress:
// [Lib.OperationBindCompletionEvent] and [Lib.OperationBindDependentEvents]
// accept a chain across two operations, which then encodes and dispatches
// correctly, and the latter rejects a cycle with status 2. That rejection is the
// control proving those calls inspect their arguments rather than accepting
// anything.
func (l *Lib) AsyncEventLastSignaledValue(event uintptr) (uint64, error) {
	out := newOut()
	err := l.callErr("e5rt_async_event_get_last_signaled_value", event, uintptr(unsafe.Pointer(out)))
	v := uint64(*out)
	runtime.KeepAlive(out)
	return v, err
}

// AsyncEventActiveFutureValue reports the event's active future value.
//
// The signature is int64_t(event, uint64_t *out). Its shape agrees across the
// shim, a CoreML call site, and the recovered ane_bridge declaration. The
// value is useful as a control for [Lib.AsyncEventSetActiveFutureValue]. On
// macOS 26.x a fresh event reports zero and reports a value set through that
// method. It is not a completion signal.
func (l *Lib) AsyncEventActiveFutureValue(event uintptr) (uint64, error) {
	out := newOut()
	err := l.callErr("e5rt_async_event_get_active_future_value", event, uintptr(unsafe.Pointer(out)))
	v := uint64(*out)
	runtime.KeepAlive(out)
	return v, err
}

// AsyncEventSetActiveFutureValue sets the value the event is expected to reach.
//
// The signature int64_t(event, uint64_t) is ANEForge's (e5rt_api.h:127). On
// macOS 26.x, [Lib.AsyncEventActiveFutureValue] reads the supplied value back
// from a fresh event. This does not signal the event: its last-signaled value
// remains unchanged.
func (l *Lib) AsyncEventSetActiveFutureValue(event uintptr, value uint64) error {
	return l.callErr("e5rt_async_event_set_active_future_value", event, uintptr(value))
}

// OperationBindCompletionEvent binds an event that the operation signals when
// it finishes. Bind it before encoding the operation.
//
// The signature int64_t(operation, event) is ANEForge's (e5rt_api.h:90, call
// site ane_e5rt_dispatch.mm:586).
//
// Binding succeeds and the operation then encodes and dispatches normally. The
// completion-event behavior under [Lib.SubmitAsync] is separately probed; use
// [Lib.SubmitAsync]'s completion function as the blocking completion signal.
func (l *Lib) OperationBindCompletionEvent(op, event uintptr) error {
	return l.callErr("e5rt_execution_stream_operation_bind_completion_event", op, event)
}

// OperationBindDependentEvents makes the operation wait for each of the given
// events before it runs. Bind them before encoding the operation. Chaining two
// operations means binding a completion event to the first with
// [Lib.OperationBindCompletionEvent] and passing that same event here to the
// second.
//
// The signature int64_t(operation, void **events, uint64_t count) is ANEForge's
// (e5rt_api.h:91, call site ane_e5rt_dispatch.mm:865).
//
// A chain across two operations over the same compiled function is accepted
// here, and the stream then encodes both and dispatches correctly. An operation
// given its own completion event is rejected with status 2, which is the control
// showing the call inspects what it is passed.
//
// That the call accepts a chain is not evidence that it enforces one. ANEForge
// reports that [Lib.ExecuteSync] serializes a stream's operations in submission
// order regardless of what is bound here, so on that path a dependency cannot be
// distinguished from the ordering that would happen anyway; it says the ordering
// has been confirmed only on the asynchronous path. The completion-event probe
// has not observed an engine-generated signal, so it cannot separate the two.
func (l *Lib) OperationBindDependentEvents(op uintptr, events []uintptr) error {
	if len(events) == 0 {
		return fmt.Errorf("e5rt: bind_dependent_events requires at least one event")
	}
	buf := append([]uintptr(nil), events...)
	err := l.callErr("e5rt_execution_stream_operation_bind_dependent_events",
		op, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	runtime.KeepAlive(buf)
	return err
}

// Deliberately unwrapped:
//
//   - e5rt_e5_compiler_is_new_compile_required is named only in the paper's
//     phase table, with no call site and no argument list anywhere in the text,
//     and does not appear in ANEForge at all.
//
//   - e5rt_execution_stream_async_submit answers "Use submit_async" and is
//     superseded by [Lib.SubmitAsync].
//
// Espresso exports on the order of two hundred e5rt_* names and [Symbols] lists
// only the ones documented here; reach the rest with [Lib.Lookup] and supply
// your own convention.
