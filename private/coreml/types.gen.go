// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objectivec"
)

// C struct types

// BayesianProbitRegression
type BayesianProbitRegression struct {
}

// CGImage
type CGImage struct {
}

// FeatureValues
type FeatureValues struct {
}

// IRProgram
type IRProgram struct {
}

// KDBoundingBox
type KDBoundingBox struct {
	Field1 [3]uint64
	Field2 uint64
}

// KDInterval
type KDInterval struct {
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

// MMappedFile
type MMappedFile struct {
}

// Model
type Model struct {
}

// MultiArrayBufferLayout
type MultiArrayBufferLayout struct {
}

// OpaqueVTPixelTransferSession
type OpaqueVTPixelTransferSession struct {
}

// Path
type Path struct {
	Field1 [3]uint64
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

// Long
type Long struct {
	__data_    *byte
	__size_    uint64
	__cap_     objectivec.Object
	__is_long_ objectivec.Object
}

// Rep
type Rep struct {
	__s int16
	__l int
}

// SharedWeakCount
type SharedWeakCount struct {
}

// Short
type Short struct {
	__data_    [23]int8
	__size_    objectivec.Object
	__is_long_ objectivec.Object
}

// E5rtAsyncEvent
type E5rtAsyncEvent struct {
}

// E5rt_async_event is a type alias for E5rtAsyncEvent for use in objc.Send[T] calls.
type E5rt_async_event = E5rtAsyncEvent

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
	Numer uint32
	Denom uint32
}

// Mach_timebase_info is a type alias for MachTimebaseInfo for use in objc.Send[T] calls.
type Mach_timebase_info = MachTimebaseInfo

// Net
type Net struct {
}

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque [56]int8
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

// OSUnfairLockS
type OSUnfairLockS struct {
	_os_unfair_lock_opaque uint32
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// SvmModel
type SvmModel struct {
	Field1  SvmParameter
	Field2  int32
	Field3  int32
	Field4  *SvmNodeRef
	Field5  *float64
	Field6  *float64
	Field7  *float64
	Field8  *float64
	Field9  *int32
	Field10 *int32
	Field11 *int32
	Field12 int32
}

// Svm_model is a type alias for SvmModel for use in objc.Send[T] calls.
type Svm_model = SvmModel

// SvmNode
type SvmNode struct {
	Field1 int32
	Field2 float64
}

// Svm_node is a type alias for SvmNode for use in objc.Send[T] calls.
type Svm_node = SvmNode

// SvmParameter
type SvmParameter struct {
	Field1  int32
	Field2  int32
	Field3  int32
	Field4  float64
	Field5  float64
	Field6  float64
	Field7  float64
	Field8  float64
	Field9  int32
	Field10 *int32
	Field11 *float64
	Field12 float64
	Field13 float64
	Field14 int32
	Field15 int32
}

// Svm_parameter is a type alias for SvmParameter for use in objc.Send[T] calls.
type Svm_parameter = SvmParameter

// Vimage2espressoParam
type Vimage2espressoParam struct {
	Field1  float32
	Field2  int32
	Field3  int32
	Field4  int32
	Field5  float32
	Field6  float32
	Field7  float32
	Field8  float32
	Field9  int32
	Field10 uint32
	Field11 uint32
	Field12 uint32
	Field13 int32
	Field14 int32
	Field15 int32
}

// Vimage2espresso_param is a type alias for Vimage2espressoParam for use in objc.Send[T] calls.
type Vimage2espresso_param = Vimage2espressoParam
