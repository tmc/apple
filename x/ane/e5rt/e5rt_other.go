//go:build !darwin

// This file is the non-darwin stub. Its doc comments are deliberately terse and
// carry no behavioral claims: pkg.go.dev renders linux/amd64 by default, so this
// is the published documentation, and every statement about what an e5rt_* entry
// point does is an UNVERIFIED reading of Bryngelson, arXiv 2606.22283 ch. 6
// rather than something observed. The darwin build of this package carries those
// caveats per function; read e5rt.go for what is and is not established.

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
	"e5rt_e5_compiler_create_with_config",
	"e5rt_e5_compiler_compile",
	"e5rt_e5_compiler_is_new_compile_required",
	"e5rt_program_library_create",
	"e5rt_program_library_retain_program_function",
	"e5rt_program_function_load_for_execution",
	"e5rt_precompiled_compute_op_create_options_create_with_program_function",
	"e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options",
	"e5rt_buffer_object_alloc",
	"e5rt_buffer_object_get_data_ptr",
	"e5rt_io_port_bind_buffer_object",
	"e5rt_execution_stream_operation_retain_input_port",
	"e5rt_execution_stream_operation_retain_output_port",
	"e5rt_execution_stream_create",
	"e5rt_execution_stream_operation_prepare_op_for_encode",
	"e5rt_execution_stream_encode_operation",
	"e5rt_execution_stream_execute_sync",
	"e5rt_execution_stream_submit_async",
	"e5rt_execution_stream_reset",
}

// A Lib is the loaded Espresso framework with its e5rt_* entry points resolved.
type Lib struct{}

// Open loads the Espresso framework and resolves every name in [Symbols].
func Open() (*Lib, error) { return nil, ErrUnsupported }

// Sym returns the address of a resolved e5rt_* entry point.
func (l *Lib) Sym(string) (uintptr, error) { return 0, ErrUnsupported }

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

// CompilerCreateWithConfig creates a compiler from an options handle.
func (l *Lib) CompilerCreateWithConfig(uintptr) (uintptr, error) { return 0, ErrUnsupported }

// CompilerCompile compiles the network description at modelPath.
func (l *Lib) CompilerCompile(uintptr, string, uintptr) (uintptr, error) { return 0, ErrUnsupported }

// ProgramLibraryRetainProgramFunction retains a callable program function.
func (l *Lib) ProgramLibraryRetainProgramFunction(uintptr, string) (uintptr, error) {
	return 0, ErrUnsupported
}

// ProgramFunctionLoadForExecution prepares a program function for execution.
func (l *Lib) ProgramFunctionLoadForExecution(uintptr) error { return ErrUnsupported }

// BufferObjectAlloc allocates a buffer object.
func (l *Lib) BufferObjectAlloc(uintptr, int) (uintptr, error) { return 0, ErrUnsupported }

// BufferObjectGetDataPtr returns the host address of a buffer object's storage.
func (l *Lib) BufferObjectGetDataPtr(uintptr) (uintptr, error) { return 0, ErrUnsupported }

// OperationRetainInputPort retains the named input port of an operation.
func (l *Lib) OperationRetainInputPort(uintptr, string) (uintptr, error) { return 0, ErrUnsupported }

// OperationRetainOutputPort retains the named output port of an operation.
func (l *Lib) OperationRetainOutputPort(uintptr, string) (uintptr, error) { return 0, ErrUnsupported }

// IOPortBindBufferObject binds a buffer object to a retained I/O port.
func (l *Lib) IOPortBindBufferObject(uintptr, uintptr) error { return ErrUnsupported }

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
