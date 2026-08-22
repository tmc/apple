//go:build !darwin

// This file is the non-darwin stub. Its doc comments are deliberately terse and
// carry no behavioral claims, because pkg.go.dev renders linux/amd64 by default
// and so this is the published documentation. When these comments were written
// nothing here had been called and every description was a reading of the source
// paper; most of them are now backed by direct observation on darwin instead.
// Either way the evidence lives next to the darwin implementation, which states
// per function what is established, what rests on ANEForge, and what has been
// contradicted. Read e5rt.go before relying on any of it.
//
// Every exported method and every name in [Symbols] must appear in both builds.
// Nothing about a separate stub implementation enforces that, so
// TestBuildsAgree parses the two files and compares them.

package e5rt

import (
	"errors"
	"fmt"
)

// FrameworkPath is the Espresso framework the e5rt_* symbols are exported from.
const FrameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"

// ErrUnsupported is returned on platforms that do not have Espresso.
var ErrUnsupported = errors.New("e5rt: unsupported platform")

// Symbols lists the e5rt_* entry points this package resolves on darwin.
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
	"e5rt_execution_stream_operation_retain_inout_port",
	"e5rt_execution_stream_operation_retain_output_port",
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

// Compute device bits for [Lib.CompilerOptionsSetComputeDeviceTypesMask].
const (
	ComputeDeviceCPU uint64 = 0x1 // BNNS
	ComputeDeviceGPU uint64 = 0x2 // MPSGraph
	ComputeDeviceANE uint64 = 0x4
)

// A Lib is the loaded Espresso framework with its e5rt_* entry points resolved.
type Lib struct{}

// Open loads the Espresso framework and resolves every name in [Symbols].
func Open() (*Lib, error) { return nil, ErrUnsupported }

// Sym returns the address of an entry point named in [Symbols].
func (l *Lib) Sym(string) (uintptr, error) { return 0, ErrUnsupported }

// Lookup resolves any exported symbol of the framework by name.
func (l *Lib) Lookup(string) (uintptr, error) { return 0, ErrUnsupported }

// Resolved reports the names from [Symbols] that resolved.
func (l *Lib) Resolved() []string { return nil }

// Status is the status code returned by an e5rt_* entry point; zero is
// reported as success. See the darwin build for what is unverified.
type Status int64

// Err reports a non-nil error for any nonzero status.
func (s Status) Err(op string) error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("e5rt: %s: status %d", op, int64(s))
}

// CompilerConfigOptionsCreate creates a compiler configuration.
func (l *Lib) CompilerConfigOptionsCreate() (uintptr, error) { return 0, ErrUnsupported }

// CompilerConfigOptionsSetCacheBundleLocation sets the compiled-bundle directory.
func (l *Lib) CompilerConfigOptionsSetCacheBundleLocation(uintptr, string) error {
	return ErrUnsupported
}

// CompilerConfigOptionsRelease releases a compiler configuration.
func (l *Lib) CompilerConfigOptionsRelease(uintptr) error { return ErrUnsupported }

// CompilerCreateWithConfig creates a compiler from a configuration handle.
func (l *Lib) CompilerCreateWithConfig(uintptr) (uintptr, error) { return 0, ErrUnsupported }

// CompilerRelease releases a compiler.
func (l *Lib) CompilerRelease(uintptr) error { return ErrUnsupported }

// CompilerOptionsCreate creates a per-compile options object.
func (l *Lib) CompilerOptionsCreate() (uintptr, error) { return 0, ErrUnsupported }

// CompilerOptionsSetComputeDeviceTypesMask selects the backends the compiler may target.
func (l *Lib) CompilerOptionsSetComputeDeviceTypesMask(uintptr, uint64) error {
	return ErrUnsupported
}

// CompilerOptionsGetComputeDeviceTypesMask reports the mask set on an options object.
func (l *Lib) CompilerOptionsGetComputeDeviceTypesMask(uintptr) (uint64, error) {
	return 0, ErrUnsupported
}

// CompilerOptionsSetForceRecompilation makes the compiler ignore a cached bundle.
func (l *Lib) CompilerOptionsSetForceRecompilation(uintptr, bool) error { return ErrUnsupported }

// CompilerOptionsSetSegmenter selects how the compiler partitions the network.
func (l *Lib) CompilerOptionsSetSegmenter(uintptr, string) error { return ErrUnsupported }

// CompilerOptionsSetCustomANECompilerOptions passes a string to the Neural Engine compiler.
func (l *Lib) CompilerOptionsSetCustomANECompilerOptions(uintptr, string) error {
	return ErrUnsupported
}

// CompilerOptionsRelease releases a compiler options object.
func (l *Lib) CompilerOptionsRelease(uintptr) error { return ErrUnsupported }

// CompilerCompile compiles the network description at modelPath.
func (l *Lib) CompilerCompile(uintptr, string, uintptr) (uintptr, error) { return 0, ErrUnsupported }

// ProgramLibraryCreate opens an already-compiled bundle.
func (l *Lib) ProgramLibraryCreate(string) (uintptr, error) { return 0, ErrUnsupported }

// ProgramLibraryRelease releases a program library.
func (l *Lib) ProgramLibraryRelease(uintptr) error { return ErrUnsupported }

// ProgramLibraryRetainProgramFunction retains a callable program function.
func (l *Lib) ProgramLibraryRetainProgramFunction(uintptr, string) (uintptr, error) {
	return 0, ErrUnsupported
}

// ProgramFunctionLoadForExecution prepares a program function for execution.
func (l *Lib) ProgramFunctionLoadForExecution(uintptr) error { return ErrUnsupported }

// ProgramFunctionRelease releases a retained program function.
func (l *Lib) ProgramFunctionRelease(uintptr) error { return ErrUnsupported }

// PrecompiledComputeOpOptionsCreate creates operation options from a program function.
func (l *Lib) PrecompiledComputeOpOptionsCreate(uintptr) (uintptr, error) {
	return 0, ErrUnsupported
}

// PrecompiledComputeOpOptionsSetOperationName names the operation.
func (l *Lib) PrecompiledComputeOpOptionsSetOperationName(uintptr, string) error {
	return ErrUnsupported
}

// PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers asks the runtime to
// allocate the operation's internal buffers.
func (l *Lib) PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(uintptr, bool) error {
	return ErrUnsupported
}

// PrecompiledComputeOpOptionsRelease releases an operation options object.
func (l *Lib) PrecompiledComputeOpOptionsRelease(uintptr) error { return ErrUnsupported }

// OperationCreatePrecompiled creates the executable operation from its options.
func (l *Lib) OperationCreatePrecompiled(uintptr) (uintptr, error) { return 0, ErrUnsupported }

// OperationRelease releases an operation.
func (l *Lib) OperationRelease(uintptr) error { return ErrUnsupported }

// BufferObjectAlloc allocates a buffer object.
func (l *Lib) BufferObjectAlloc(uintptr, int) (uintptr, error) { return 0, ErrUnsupported }

// BufferObjectGetDataPtr returns the host address of a buffer object's storage.
func (l *Lib) BufferObjectGetDataPtr(uintptr) (uintptr, error) { return 0, ErrUnsupported }

// BufferObjectRelease releases a buffer object.
func (l *Lib) BufferObjectRelease(uintptr) error { return ErrUnsupported }

// OperationRetainInputPort retains the named input port of an operation.
func (l *Lib) OperationRetainInputPort(uintptr, string) (uintptr, error) { return 0, ErrUnsupported }

// OperationRetainInoutPort retains the named inout port of an operation.
func (l *Lib) OperationRetainInoutPort(uintptr, string) (uintptr, error) { return 0, ErrUnsupported }

// OperationRetainOutputPort retains the named output port of an operation.
func (l *Lib) OperationRetainOutputPort(uintptr, string) (uintptr, error) { return 0, ErrUnsupported }

// IOPortBindBufferObject binds a buffer object to a retained I/O port.
func (l *Lib) IOPortBindBufferObject(uintptr, uintptr) error { return ErrUnsupported }

// IOPortRelease releases a retained I/O port.
func (l *Lib) IOPortRelease(uintptr) error { return ErrUnsupported }

// ExecutionStreamCreate creates an execution stream.
func (l *Lib) ExecutionStreamCreate() (uintptr, error) { return 0, ErrUnsupported }

// PrepareOpForEncode prepares an operation before each encode.
func (l *Lib) PrepareOpForEncode(uintptr) error { return ErrUnsupported }

// EncodeOperation encodes a prepared operation into a stream.
func (l *Lib) EncodeOperation(uintptr, uintptr) error { return ErrUnsupported }

// ExecuteSync submits an encoded stream and blocks until it completes.
func (l *Lib) ExecuteSync(uintptr) error { return ErrUnsupported }

// ExecutionStreamReset resets a stream for reuse after execution.
func (l *Lib) ExecutionStreamReset(uintptr) error { return ErrUnsupported }

// ExecutionStreamRelease releases an execution stream.
func (l *Lib) ExecutionStreamRelease(uintptr) error { return ErrUnsupported }

// SubmitAsync submits an encoded stream without blocking and returns a function
// that releases the completion block.
func (l *Lib) SubmitAsync(uintptr, func() Status) (func(), error) { return nil, ErrUnsupported }

// ExecutionStreamSetQualityOfService sets a stream's dispatch quality of service.
func (l *Lib) ExecutionStreamSetQualityOfService(uintptr, uint64) error { return ErrUnsupported }

// ExecutionStreamSetANEExecutionPriority sets a stream's Neural Engine priority.
func (l *Lib) ExecutionStreamSetANEExecutionPriority(uintptr, uint64) error {
	return ErrUnsupported
}

// AsyncEventCreate creates a named async event.
func (l *Lib) AsyncEventCreate(string) (uintptr, error) { return 0, ErrUnsupported }

// AsyncEventRelease releases an async event.
func (l *Lib) AsyncEventRelease(uintptr) error { return ErrUnsupported }

// AsyncEventSignal signals an async event from the host.
func (l *Lib) AsyncEventSignal(uintptr) error { return ErrUnsupported }

// AsyncEventSyncWait waits for an async event to be signaled.
func (l *Lib) AsyncEventSyncWait(uintptr) error { return ErrUnsupported }

// AsyncEventLastSignaledValue reports how many times an event has been signaled.
func (l *Lib) AsyncEventLastSignaledValue(uintptr) (uint64, error) { return 0, ErrUnsupported }

// AsyncEventSetActiveFutureValue sets the value an event is expected to reach.
func (l *Lib) AsyncEventSetActiveFutureValue(uintptr, uint64) error { return ErrUnsupported }

// OperationBindCompletionEvent binds the event an operation signals when it finishes.
func (l *Lib) OperationBindCompletionEvent(uintptr, uintptr) error { return ErrUnsupported }

// OperationBindDependentEvents makes an operation wait for the given events.
func (l *Lib) OperationBindDependentEvents(uintptr, []uintptr) error { return ErrUnsupported }
