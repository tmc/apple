// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixMultiplication] class.
var (
	_MPSMatrixMultiplicationClass     MPSMatrixMultiplicationClass
	_MPSMatrixMultiplicationClassOnce sync.Once
)

func getMPSMatrixMultiplicationClass() MPSMatrixMultiplicationClass {
	_MPSMatrixMultiplicationClassOnce.Do(func() {
		_MPSMatrixMultiplicationClass = MPSMatrixMultiplicationClass{class: objc.GetClass("MPSMatrixMultiplication")}
	})
	return _MPSMatrixMultiplicationClass
}

// GetMPSMatrixMultiplicationClass returns the class object for MPSMatrixMultiplication.
func GetMPSMatrixMultiplicationClass() MPSMatrixMultiplicationClass {
	return getMPSMatrixMultiplicationClass()
}

type MPSMatrixMultiplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixMultiplicationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixMultiplicationClass) Alloc() MPSMatrixMultiplication {
	rv := objc.Send[MPSMatrixMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A matrix multiplication kernel.
//
// # Overview
//
// An [MPSMatrixMultiplication] object computes the following operation:
//
// C = alpha * op(A) * op(B) + beta * C
//
// Where A, B_,_ and C are matrices represented by [MPSMatrix] objects, and
// alpha and beta are scalar values of the same data type as the values of C.
// A and B may each have an optional transposition operation applied.
//
// Matrices A, B, and C are also referred to as the left input matrix, the
// right input matrix, and the result matrix respectively.
//
// # Methods
//
//   - [MPSMatrixMultiplication.InitWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta]: Initializes a matrix multiplication kernel.
//   - [MPSMatrixMultiplication.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix]: Encodes a matrix multiplication kernel to a command buffer.
//
// # Properties
//
//   - [MPSMatrixMultiplication.LeftMatrixOrigin]: The origin of the left input matrix.
//   - [MPSMatrixMultiplication.SetLeftMatrixOrigin]
//   - [MPSMatrixMultiplication.RightMatrixOrigin]: The origin of the right input matrix.
//   - [MPSMatrixMultiplication.SetRightMatrixOrigin]
//   - [MPSMatrixMultiplication.ResultMatrixOrigin]: The origin of the result matrix.
//   - [MPSMatrixMultiplication.SetResultMatrixOrigin]
//   - [MPSMatrixMultiplication.BatchSize]
//   - [MPSMatrixMultiplication.SetBatchSize]
//   - [MPSMatrixMultiplication.BatchStart]
//   - [MPSMatrixMultiplication.SetBatchStart]
//
// # Initializers
//
//   - [MPSMatrixMultiplication.InitWithDeviceResultRowsResultColumnsInteriorColumns]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication
type MPSMatrixMultiplication struct {
	MPSKernel
}

// MPSMatrixMultiplicationFromID constructs a [MPSMatrixMultiplication] from an objc.ID.
//
// A matrix multiplication kernel.
func MPSMatrixMultiplicationFromID(id objc.ID) MPSMatrixMultiplication {
	return MPSMatrixMultiplication{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixMultiplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixMultiplication] class.
//
// # Methods
//
//   - [IMPSMatrixMultiplication.InitWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta]: Initializes a matrix multiplication kernel.
//   - [IMPSMatrixMultiplication.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix]: Encodes a matrix multiplication kernel to a command buffer.
//
// # Properties
//
//   - [IMPSMatrixMultiplication.LeftMatrixOrigin]: The origin of the left input matrix.
//   - [IMPSMatrixMultiplication.SetLeftMatrixOrigin]
//   - [IMPSMatrixMultiplication.RightMatrixOrigin]: The origin of the right input matrix.
//   - [IMPSMatrixMultiplication.SetRightMatrixOrigin]
//   - [IMPSMatrixMultiplication.ResultMatrixOrigin]: The origin of the result matrix.
//   - [IMPSMatrixMultiplication.SetResultMatrixOrigin]
//   - [IMPSMatrixMultiplication.BatchSize]
//   - [IMPSMatrixMultiplication.SetBatchSize]
//   - [IMPSMatrixMultiplication.BatchStart]
//   - [IMPSMatrixMultiplication.SetBatchStart]
//
// # Initializers
//
//   - [IMPSMatrixMultiplication.InitWithDeviceResultRowsResultColumnsInteriorColumns]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication
type IMPSMatrixMultiplication interface {
	IMPSKernel

	// Topic: Methods

	// Initializes a matrix multiplication kernel.
	InitWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(device metal.MTLDevice, transposeLeft bool, transposeRight bool, resultRows uint, resultColumns uint, interiorColumns uint, alpha float64, beta float64) MPSMatrixMultiplication
	// Encodes a matrix multiplication kernel to a command buffer.
	EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, leftMatrix IMPSMatrix, rightMatrix IMPSMatrix, resultMatrix IMPSMatrix)

	// Topic: Properties

	// The origin of the left input matrix.
	LeftMatrixOrigin() metal.MTLOrigin
	SetLeftMatrixOrigin(value metal.MTLOrigin)
	// The origin of the right input matrix.
	RightMatrixOrigin() metal.MTLOrigin
	SetRightMatrixOrigin(value metal.MTLOrigin)
	// The origin of the result matrix.
	ResultMatrixOrigin() metal.MTLOrigin
	SetResultMatrixOrigin(value metal.MTLOrigin)
	BatchSize() uint
	SetBatchSize(value uint)
	BatchStart() uint
	SetBatchStart(value uint)

	// Topic: Initializers

	InitWithDeviceResultRowsResultColumnsInteriorColumns(device metal.MTLDevice, resultRows uint, resultColumns uint, interiorColumns uint) MPSMatrixMultiplication
}

// Init initializes the instance.
func (m MPSMatrixMultiplication) Init() MPSMatrixMultiplication {
	rv := objc.Send[MPSMatrixMultiplication](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixMultiplication) Autorelease() MPSMatrixMultiplication {
	rv := objc.Send[MPSMatrixMultiplication](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixMultiplication creates a new MPSMatrixMultiplication instance.
func NewMPSMatrixMultiplication() MPSMatrixMultiplication {
	class := getMPSMatrixMultiplicationClass()
	rv := objc.Send[MPSMatrixMultiplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixMultiplicationWithCoder(aDecoder foundation.INSCoder) MPSMatrixMultiplication {
	instance := getMPSMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixMultiplicationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixMultiplication {
	instance := getMPSMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixMultiplicationFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewMatrixMultiplicationWithDevice(device metal.MTLDevice) MPSMatrixMultiplication {
	instance := getMPSMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/init(device:resultRows:resultColumns:interiorColumns:)
func NewMatrixMultiplicationWithDeviceResultRowsResultColumnsInteriorColumns(device metal.MTLDevice, resultRows uint, resultColumns uint, interiorColumns uint) MPSMatrixMultiplication {
	instance := getMPSMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultRows:resultColumns:interiorColumns:"), device, resultRows, resultColumns, interiorColumns)
	return MPSMatrixMultiplicationFromID(rv)
}

// Initializes a matrix multiplication kernel.
//
// device: The device on which the matrix multiplication kernel will run.
//
// transposeLeft: A boolean value that indicates if the left input matrix should be used in
// its transposed form. If the value is true, then `op(A) = A**T`; otherwise,
// `op(A) = A`.
//
// transposeRight: A boolean value that indicates if the right input matrix should be used in
// its transposed form. If the value is true, then `op(B) = B**T`; otherwise,
// `op(B) = B`.
//
// resultRows: The number of rows in the result matrix ([M] in the BLAS GEMM description).
//
// resultColumns: The number of columns in the result matrix ([N] in the BLAS GEMM
// description).
//
// interiorColumns: The number of columns of the left input matrix after the appropriate
// transpose operation has been applied ([K] in the BLAS GEMM description).
//
// alpha: The scale factor to apply to the product, specified in `double` precision.
// This value will be converted to the appropriate precision in the
// implementation itself, subject to rounding and/or clamping as necessary.
//
// beta: The scale factor to apply to the initial values of [C], specified in
// `double` precision. This value will be converted to the appropriate
// precision in the implementation itself, subject to rounding and/or clamping
// as necessary.
//
// # Return Value
//
// A valid [MPSMatrixMultiplication] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/init(device:transposeLeft:transposeRight:resultRows:resultColumns:interiorColumns:alpha:beta:)
func NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(device metal.MTLDevice, transposeLeft bool, transposeRight bool, resultRows uint, resultColumns uint, interiorColumns uint, alpha float64, beta float64) MPSMatrixMultiplication {
	instance := getMPSMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:transposeLeft:transposeRight:resultRows:resultColumns:interiorColumns:alpha:beta:"), device, transposeLeft, transposeRight, resultRows, resultColumns, interiorColumns, alpha, beta)
	return MPSMatrixMultiplicationFromID(rv)
}

// Initializes a matrix multiplication kernel.
//
// device: The device on which the matrix multiplication kernel will run.
//
// transposeLeft: A boolean value that indicates if the left input matrix should be used in
// its transposed form. If the value is true, then `op(A) = A**T`; otherwise,
// `op(A) = A`.
//
// transposeRight: A boolean value that indicates if the right input matrix should be used in
// its transposed form. If the value is true, then `op(B) = B**T`; otherwise,
// `op(B) = B`.
//
// resultRows: The number of rows in the result matrix ([M] in the BLAS GEMM description).
//
// resultColumns: The number of columns in the result matrix ([N] in the BLAS GEMM
// description).
//
// interiorColumns: The number of columns of the left input matrix after the appropriate
// transpose operation has been applied ([K] in the BLAS GEMM description).
//
// alpha: The scale factor to apply to the product, specified in `double` precision.
// This value will be converted to the appropriate precision in the
// implementation itself, subject to rounding and/or clamping as necessary.
//
// beta: The scale factor to apply to the initial values of [C], specified in
// `double` precision. This value will be converted to the appropriate
// precision in the implementation itself, subject to rounding and/or clamping
// as necessary.
//
// # Return Value
//
// A valid [MPSMatrixMultiplication] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/init(device:transposeLeft:transposeRight:resultRows:resultColumns:interiorColumns:alpha:beta:)
func (m MPSMatrixMultiplication) InitWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(device metal.MTLDevice, transposeLeft bool, transposeRight bool, resultRows uint, resultColumns uint, interiorColumns uint, alpha float64, beta float64) MPSMatrixMultiplication {
	rv := objc.Send[MPSMatrixMultiplication](m.ID, objc.Sel("initWithDevice:transposeLeft:transposeRight:resultRows:resultColumns:interiorColumns:alpha:beta:"), device, transposeLeft, transposeRight, resultRows, resultColumns, interiorColumns, alpha, beta)
	return rv
}

// Encodes a matrix multiplication kernel to a command buffer.
//
// commandBuffer: The command buffer that will receive the encoded kernel.
//
// leftMatrix: The left input matrix.
//
// rightMatrix: The right input matrix.
//
// resultMatrix: The addend matrix which will also be overwritten by the operation result.
//
// # Discussion
//
// The following constraints apply to the sizes of the matrices depending on
// the transposition operations and the sizes requested at initialization
// time, as well as the origins at the time this method is called:
//
// - The left input matrix must be large enough to hold an array of size
// `resultRows x interiorColumns` elements, beginning at the value of the
// [MPSMatrixMultiplication.LeftMatrixOrigin] property. - The right input
// matrix must be large enough to hold an array of size `interiorColumns x
// resultColumns` elements, beginning at the value of the
// [MPSMatrixMultiplication.RightMatrixOrigin] property. - The result matrix
// must be large enough to hold an array of size `resultRows x resultColumns`
// elements, beginning at the value of the
// [MPSMatrixMultiplication.ResultMatrixOrigin] property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/encode(commandBuffer:leftMatrix:rightMatrix:resultMatrix:)
func (m MPSMatrixMultiplication) EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer metal.MTLCommandBuffer, leftMatrix IMPSMatrix, rightMatrix IMPSMatrix, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:leftMatrix:rightMatrix:resultMatrix:"), commandBuffer, leftMatrix, rightMatrix, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/init(device:resultRows:resultColumns:interiorColumns:)
func (m MPSMatrixMultiplication) InitWithDeviceResultRowsResultColumnsInteriorColumns(device metal.MTLDevice, resultRows uint, resultColumns uint, interiorColumns uint) MPSMatrixMultiplication {
	rv := objc.Send[MPSMatrixMultiplication](m.ID, objc.Sel("initWithDevice:resultRows:resultColumns:interiorColumns:"), device, resultRows, resultColumns, interiorColumns)
	return rv
}

// The origin of the left input matrix.
//
// # Discussion
//
// The origin, relative to `(0,0)`, at which to start reading values. If a
// different origin is desired, you must modify this property before encoding
// the matrix multiplication kernel. The default value is `(0,0,0)` (the `z`
// value must always be 0).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/leftMatrixOrigin
func (m MPSMatrixMultiplication) LeftMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("leftMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixMultiplication) SetLeftMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setLeftMatrixOrigin:"), value)
}

// The origin of the right input matrix.
//
// # Discussion
//
// The origin, relative to `(0,0)`, at which to start reading values. If a
// different origin is desired, you must modify this property before encoding
// the matrix multiplication kernel. The default value is `(0,0,0)` (the `z`
// value must always be 0).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/rightMatrixOrigin
func (m MPSMatrixMultiplication) RightMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("rightMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixMultiplication) SetRightMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setRightMatrixOrigin:"), value)
}

// The origin of the result matrix.
//
// # Discussion
//
// The origin, relative to `(0,0)`, at which to start reading values. If a
// different origin is desired, you must modify this property before encoding
// the matrix multiplication kernel. The default value is `(0,0,0)` (the `z`
// value must always be 0).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/resultMatrixOrigin
func (m MPSMatrixMultiplication) ResultMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("resultMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixMultiplication) SetResultMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setResultMatrixOrigin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/batchSize
func (m MPSMatrixMultiplication) BatchSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MPSMatrixMultiplication) SetBatchSize(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication/batchStart
func (m MPSMatrixMultiplication) BatchStart() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchStart"))
	return rv
}
func (m MPSMatrixMultiplication) SetBatchStart(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchStart:"), value)
}
