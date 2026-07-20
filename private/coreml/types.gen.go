// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
)

// C struct types

// CGImage
type CGImage struct {
}

// KDBoundingBox
type KDBoundingBox struct {
	Field1 unsafe.Pointer
	Field2 uint64
}

// MLModelDescriptionSpecification
type MLModelDescriptionSpecification struct {
}

// MLModelInputArchiver
type MLModelInputArchiver struct {
}

// MLModelSpecification
type MLModelSpecification struct {
}

// OpaqueVTPixelTransferSession
type OpaqueVTPixelTransferSession struct {
}

// Path
type Path struct {
	Field1 unsafe.Pointer
}

// Prediction
type Prediction struct {
	Field1 float64
	Field2 float64
	Field3 float64
	Field4 float64
	Field5 float64
	Field6 bool
}

// CVBuffer
type CVBuffer struct {
}

// CVPixelBufferPool
type CVPixelBufferPool struct {
}

// E5rtExecutionStream
type E5rtExecutionStream struct {
}

// E5rt_execution_stream is a type alias for E5rtExecutionStream for use in objc.Send[T] calls.
type E5rt_execution_stream = E5rtExecutionStream

// E5rtExecutionStreamOperation
type E5rtExecutionStreamOperation struct {
}

// E5rt_execution_stream_operation is a type alias for E5rtExecutionStreamOperation for use in objc.Send[T] calls.
type E5rt_execution_stream_operation = E5rtExecutionStreamOperation

// E5rtIOPort
type E5rtIOPort struct {
}

// E5rt_io_port is a type alias for E5rtIOPort for use in objc.Send[T] calls.
type E5rt_io_port = E5rtIOPort

// E5rtProgramLibrary
type E5rtProgramLibrary struct {
}

// E5rt_program_library is a type alias for E5rtProgramLibrary for use in objc.Send[T] calls.
type E5rt_program_library = E5rtProgramLibrary

// MachTimebaseInfo
type MachTimebaseInfo struct {
	Numer uint
	Denom uint
}

// Mach_timebase_info is a type alias for MachTimebaseInfo for use in objc.Send[T] calls.
type Mach_timebase_info = MachTimebaseInfo

// Mutex
type Mutex struct {
	__m_ unsafe.Pointer
}

// OSUnfairLockS
type OSUnfairLockS struct {
	_os_unfair_lock_opaque uint
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// SvmModel
type SvmModel struct {
	Field1  unsafe.Pointer
	Field2  int
	Field3  int
	Field4  *SvmNodeRef
	Field5  []float64
	Field6  []float64
	Field7  []float64
	Field8  []float64
	Field9  []int
	Field10 []int
	Field11 []int
	Field12 int
}

// Svm_model is a type alias for SvmModel for use in objc.Send[T] calls.
type Svm_model = SvmModel

// SvmNode
type SvmNode struct {
	Field1 int
	Field2 float64
}

// Svm_node is a type alias for SvmNode for use in objc.Send[T] calls.
type Svm_node = SvmNode

// Vimage2espressoParam
type Vimage2espressoParam struct {
	Field1  float32
	Field2  int
	Field3  int
	Field4  int
	Field5  float32
	Field6  float32
	Field7  float32
	Field8  float32
	Field9  int
	Field10 uint
	Field11 uint
	Field12 uint
	Field13 int
	Field14 int
	Field15 int
}

// Vimage2espresso_param is a type alias for Vimage2espressoParam for use in objc.Send[T] calls.
type Vimage2espresso_param = Vimage2espressoParam
