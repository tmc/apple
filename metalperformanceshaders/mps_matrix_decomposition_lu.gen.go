// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixDecompositionLU] class.
var (
	_MPSMatrixDecompositionLUClass     MPSMatrixDecompositionLUClass
	_MPSMatrixDecompositionLUClassOnce sync.Once
)

func getMPSMatrixDecompositionLUClass() MPSMatrixDecompositionLUClass {
	_MPSMatrixDecompositionLUClassOnce.Do(func() {
		_MPSMatrixDecompositionLUClass = MPSMatrixDecompositionLUClass{class: objc.GetClass("MPSMatrixDecompositionLU")}
	})
	return _MPSMatrixDecompositionLUClass
}

// GetMPSMatrixDecompositionLUClass returns the class object for MPSMatrixDecompositionLU.
func GetMPSMatrixDecompositionLUClass() MPSMatrixDecompositionLUClass {
	return getMPSMatrixDecompositionLUClass()
}

type MPSMatrixDecompositionLUClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixDecompositionLUClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixDecompositionLUClass) Alloc() MPSMatrixDecompositionLU {
	rv := objc.Send[MPSMatrixDecompositionLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the LU factorization of a matrix using partial
// pivoting with row interchanges.
//
// # Overview
//
// This kernel object computes an LU factorization, PA = LU, where:
//
// - A is a matrix for which the LU factorization is to be computed - L is a
// unit lower triangular matrix - U is an upper triangular matrix - P is a
// permutation matrix
//
// # Initializers
//
//   - [MPSMatrixDecompositionLU.InitWithDeviceRowsColumns]
//
// # Instance Methods
//
//   - [MPSMatrixDecompositionLU.EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU
type MPSMatrixDecompositionLU struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixDecompositionLUFromID constructs a [MPSMatrixDecompositionLU] from an objc.ID.
//
// A kernel for computing the LU factorization of a matrix using partial
// pivoting with row interchanges.
func MPSMatrixDecompositionLUFromID(id objc.ID) MPSMatrixDecompositionLU {
	return MPSMatrixDecompositionLU{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixDecompositionLU adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixDecompositionLU] class.
//
// # Initializers
//
//   - [IMPSMatrixDecompositionLU.InitWithDeviceRowsColumns]
//
// # Instance Methods
//
//   - [IMPSMatrixDecompositionLU.EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU
type IMPSMatrixDecompositionLU interface {
	IMPSMatrixUnaryKernel

	// Topic: Initializers

	InitWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixDecompositionLU

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, resultMatrix IMPSMatrix, pivotIndices IMPSMatrix, status metal.MTLBuffer)
}

// Init initializes the instance.
func (m MPSMatrixDecompositionLU) Init() MPSMatrixDecompositionLU {
	rv := objc.Send[MPSMatrixDecompositionLU](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixDecompositionLU) Autorelease() MPSMatrixDecompositionLU {
	rv := objc.Send[MPSMatrixDecompositionLU](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixDecompositionLU creates a new MPSMatrixDecompositionLU instance.
func NewMPSMatrixDecompositionLU() MPSMatrixDecompositionLU {
	class := getMPSMatrixDecompositionLUClass()
	rv := objc.Send[MPSMatrixDecompositionLU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixDecompositionLUWithCoder(aDecoder foundation.INSCoder) MPSMatrixDecompositionLU {
	instance := getMPSMatrixDecompositionLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixDecompositionLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixDecompositionLUWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixDecompositionLU {
	instance := getMPSMatrixDecompositionLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixDecompositionLUFromID(rv)
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
func NewMatrixDecompositionLUWithDevice(device metal.MTLDevice) MPSMatrixDecompositionLU {
	instance := getMPSMatrixDecompositionLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixDecompositionLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU/init(device:rows:columns:)
func NewMatrixDecompositionLUWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixDecompositionLU {
	instance := getMPSMatrixDecompositionLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	return MPSMatrixDecompositionLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU/init(device:rows:columns:)
func (m MPSMatrixDecompositionLU) InitWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixDecompositionLU {
	rv := objc.Send[MPSMatrixDecompositionLU](m.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU/encode(commandBuffer:sourceMatrix:resultMatrix:pivotIndices:info:)
func (m MPSMatrixDecompositionLU) EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, resultMatrix IMPSMatrix, pivotIndices IMPSMatrix, status metal.MTLBuffer) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:pivotIndices:status:"), commandBuffer, sourceMatrix, resultMatrix, pivotIndices, status)
}
