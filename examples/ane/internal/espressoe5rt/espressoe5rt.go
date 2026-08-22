//go:build darwin

// Package espressoe5rt adapts Espresso's generated e5rt bindings for examples.
//
// It is deliberately kept under examples: the generated signatures are useful
// for experiments, but this package does not promote them to a public API.
package espressoe5rt

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/private/espresso"
)

// FrameworkPath is Espresso's private framework path.
const FrameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"

const (
	ComputeDeviceCPU uint64 = 0x1
	ComputeDeviceGPU uint64 = 0x2
	ComputeDeviceANE uint64 = 0x4
)

// Symbols is the direct-dispatch symbol set reported by the examples.
//
// Resolving a name establishes only that Espresso exports it. The typed methods
// in this package use signatures with the manifest's strongest grade, but that
// grade comes from the handwritten e5rt adapter's own call sites. It records
// that route being exercised; it is not independent confirmation. The
// remaining names are retained for the probe's explicitly raw experiments.
var Symbols = []string{
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
	"e5rt_buffer_object_alloc",
	"e5rt_buffer_object_get_data_ptr",
	"e5rt_buffer_object_release",
	"e5rt_io_port_bind_buffer_object",
	"e5rt_io_port_release",
	"e5rt_execution_stream_operation_retain_input_port",
	"e5rt_execution_stream_operation_retain_output_port",
	"e5rt_execution_stream_operation_retain_inout_port",
	"e5rt_execution_stream_create",
	"e5rt_execution_stream_operation_prepare_op_for_encode",
	"e5rt_execution_stream_encode_operation",
	"e5rt_execution_stream_execute_sync",
	"e5rt_execution_stream_submit_async",
	"e5rt_execution_stream_reset",
	"e5rt_execution_stream_release",
	"e5rt_execution_stream_set_quality_of_service",
	"e5rt_execution_stream_set_ane_execution_priority",
	"e5rt_async_event_create",
	"e5rt_async_event_release",
	"e5rt_async_event_signal",
	"e5rt_async_event_sync_wait",
	"e5rt_async_event_get_last_signaled_value",
	"e5rt_async_event_set_active_future_value",
	"e5rt_execution_stream_operation_bind_completion_event",
	"e5rt_execution_stream_operation_bind_dependent_events",
}

// Lib is the generated Espresso binding, checked for the symbols the examples
// use. Its zero value is not usable.
type Lib struct{ resolved map[string]bool }

// Open verifies that Espresso exports every name in Symbols. It does not verify
// a calling convention or promote a raw symbol to a supported binding.
func Open() (*Lib, error) {
	l := &Lib{resolved: make(map[string]bool, len(Symbols))}
	for _, name := range Symbols {
		if _, err := espresso.SymbolAddress(name); err != nil {
			return l, fmt.Errorf("resolve %s: %w", name, err)
		}
		l.resolved[name] = true
	}
	return l, nil
}

// Resolved returns the names in Symbols that Open resolved, in Symbols order.
func (l *Lib) Resolved() []string {
	var out []string
	for _, name := range Symbols {
		if l != nil && l.resolved[name] {
			out = append(out, name)
		}
	}
	return out
}

// Sym resolves a generated Espresso symbol by name. It is intended only for
// the probe's deliberately isolated raw-call experiments.
func (l *Lib) Sym(name string) (uintptr, error) {
	if l == nil || !l.resolved[name] {
		return 0, fmt.Errorf("espresso e5rt: symbol %s unresolved", name)
	}
	return espresso.SymbolAddress(name)
}

func check(name string, status int64, err error) error {
	if err != nil {
		return err
	}
	if status != 0 {
		return fmt.Errorf("espresso e5rt: %s: status %d", name, status)
	}
	return nil
}

// cstring returns a C string whose backing storage remains live through a call.
func cstring(s string) ([]byte, *byte) {
	b := append([]byte(s), 0)
	return b, (*byte)(unsafe.Pointer(&b[0]))
}

func boolInt32(v bool) int32 {
	if v {
		return 1
	}
	return 0
}

func release(name string, f func(*uintptr) (int64, error), handle uintptr) error {
	p := new(uintptr)
	*p = handle
	status, callErr := f(p)
	err := check(name, status, callErr)
	runtime.KeepAlive(p)
	return err
}

func (l *Lib) CompilerConfigOptionsCreate() (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtE5CompilerConfigOptionsCreate(out)
	err := check("config_options_create", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) CompilerConfigOptionsSetCacheBundleLocation(config uintptr, dir string) error {
	b, p := cstring(dir)
	status, callErr := espresso.E5rtE5CompilerConfigOptionsSetCacheBundleLocation(config, p)
	err := check("config_options_set_cache_bundle_location", status, callErr)
	runtime.KeepAlive(b)
	return err
}

func (l *Lib) CompilerConfigOptionsRelease(config uintptr) error {
	return release("config_options_release", espresso.E5rtE5CompilerConfigOptionsRelease, config)
}

func (l *Lib) CompilerCreateWithConfig(config uintptr) (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtE5CompilerCreateWithConfig(out, config)
	err := check("compiler_create_with_config", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) CompilerRelease(compiler uintptr) error {
	return release("compiler_release", espresso.E5rtE5CompilerRelease, compiler)
}

func (l *Lib) CompilerOptionsCreate() (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtE5CompilerOptionsCreate(out)
	err := check("compiler_options_create", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) CompilerOptionsSetComputeDeviceTypesMask(options uintptr, mask uint64) error {
	status, err := espresso.E5rtE5CompilerOptionsSetComputeDeviceTypesMask(options, mask)
	return check("compiler_options_set_compute_device_types_mask", status, err)
}

func (l *Lib) CompilerOptionsSetForceRecompilation(options uintptr, force bool) error {
	status, err := espresso.E5rtE5CompilerOptionsSetForceRecompilation(options, boolInt32(force))
	return check("compiler_options_set_force_recompilation", status, err)
}

func (l *Lib) CompilerOptionsSetSegmenter(options uintptr, segmenter string) error {
	b, p := cstring(segmenter)
	status, callErr := espresso.E5rtE5CompilerOptionsSetSegmenter(options, p)
	err := check("compiler_options_set_segmenter", status, callErr)
	runtime.KeepAlive(b)
	return err
}

func (l *Lib) CompilerOptionsRelease(options uintptr) error {
	return release("compiler_options_release", espresso.E5rtE5CompilerOptionsRelease, options)
}

func (l *Lib) CompilerCompile(compiler uintptr, modelPath string, options uintptr) (uintptr, error) {
	b, p := cstring(modelPath)
	out := new(uintptr)
	status, callErr := espresso.E5rtE5CompilerCompile(compiler, p, options, out)
	err := check("compiler_compile", status, callErr)
	runtime.KeepAlive(b)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) ProgramLibraryRelease(library uintptr) error {
	return release("program_library_release", espresso.E5rtProgramLibraryRelease, library)
}

func (l *Lib) ProgramFunctionLoadForExecution(function uintptr) error {
	status, err := espresso.E5rtProgramFunctionLoadForExecution(function)
	return check("program_function_load_for_execution", status, err)
}

func (l *Lib) ProgramFunctionRelease(function uintptr) error {
	return release("program_function_release", espresso.E5rtProgramFunctionRelease, function)
}

func (l *Lib) ProgramLibraryRetainProgramFunction(library uintptr, name string) (uintptr, error) {
	b, p := cstring(name)
	out := new(uintptr)
	status, callErr := espresso.E5rtProgramLibraryRetainProgramFunction(library, p, out)
	err := check("program_library_retain_program_function", status, callErr)
	runtime.KeepAlive(b)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) PrecompiledComputeOpOptionsCreate(function uintptr) (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction(out, function)
	err := check("precompiled_compute_op_create_options_create_with_program_function", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) PrecompiledComputeOpOptionsSetOperationName(options uintptr, name string) error {
	b, p := cstring(name)
	status, callErr := espresso.E5rtPrecompiledComputeOpCreateOptionsSetOperationName(options, p)
	err := check("precompiled_compute_op_create_options_set_operation_name", status, callErr)
	runtime.KeepAlive(b)
	return err
}

func (l *Lib) PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(options uintptr, allocate bool) error {
	status, err := espresso.E5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(options, boolInt32(allocate))
	return check("precompiled_compute_op_create_options_set_allocate_intermediate_buffers", status, err)
}

func (l *Lib) PrecompiledComputeOpOptionsRelease(options uintptr) error {
	return release("precompiled_compute_op_create_options_release", espresso.E5rtPrecompiledComputeOpCreateOptionsRelease, options)
}

func (l *Lib) OperationCreatePrecompiled(options uintptr) (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out, options)
	err := check("operation_create_precompiled", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) OperationRelease(operation uintptr) error {
	return release("operation_release", espresso.E5rtExecutionStreamOperationRelease, operation)
}

func (l *Lib) BufferObjectAlloc(nbytes uintptr, kind int) (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtBufferObjectAlloc(out, uint64(nbytes), uint32(kind))
	err := check("buffer_object_alloc", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) BufferObjectGetDataPtr(buffer uintptr) (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtBufferObjectGetDataPtr(buffer, out)
	err := check("buffer_object_get_data_ptr", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) BufferObjectRelease(buffer uintptr) error {
	return release("buffer_object_release", espresso.E5rtBufferObjectRelease, buffer)
}

func (l *Lib) retainPort(name string, f func(uintptr, *byte, *uintptr) (int64, error), operation uintptr, portName string) (uintptr, error) {
	b, p := cstring(portName)
	out := new(uintptr)
	status, callErr := f(operation, p, out)
	err := check(name, status, callErr)
	runtime.KeepAlive(b)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) OperationRetainInputPort(operation uintptr, portName string) (uintptr, error) {
	return l.retainPort("operation_retain_input_port", espresso.E5rtExecutionStreamOperationRetainInputPort, operation, portName)
}

func (l *Lib) OperationRetainOutputPort(operation uintptr, portName string) (uintptr, error) {
	return l.retainPort("operation_retain_output_port", espresso.E5rtExecutionStreamOperationRetainOutputPort, operation, portName)
}

func (l *Lib) IOPortBindBufferObject(port, buffer uintptr) error {
	status, err := espresso.E5rtIOPortBindBufferObject(port, buffer)
	return check("io_port_bind_buffer_object", status, err)
}

func (l *Lib) IOPortRelease(port uintptr) error {
	return release("io_port_release", espresso.E5rtIOPortRelease, port)
}

func (l *Lib) ExecutionStreamCreate() (uintptr, error) {
	out := new(uintptr)
	status, callErr := espresso.E5rtExecutionStreamCreate(out)
	err := check("execution_stream_create", status, callErr)
	runtime.KeepAlive(out)
	return *out, err
}

func (l *Lib) EncodeOperation(stream, operation uintptr) error {
	status, err := espresso.E5rtExecutionStreamEncodeOperation(stream, operation)
	return check("execution_stream_encode_operation", status, err)
}

func (l *Lib) ExecuteSync(stream uintptr) error {
	status, err := espresso.E5rtExecutionStreamExecuteSync(stream)
	return check("execution_stream_execute_sync", status, err)
}

func (l *Lib) ExecutionStreamReset(stream uintptr) error {
	status, err := espresso.E5rtExecutionStreamReset(stream)
	return check("execution_stream_reset", status, err)
}

func (l *Lib) ExecutionStreamRelease(stream uintptr) error {
	return release("execution_stream_release", espresso.E5rtExecutionStreamRelease, stream)
}
