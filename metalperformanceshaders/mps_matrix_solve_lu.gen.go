// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixSolveLU] class.
var (
	_MPSMatrixSolveLUClass     MPSMatrixSolveLUClass
	_MPSMatrixSolveLUClassOnce sync.Once
)

func getMPSMatrixSolveLUClass() MPSMatrixSolveLUClass {
	_MPSMatrixSolveLUClassOnce.Do(func() {
		_MPSMatrixSolveLUClass = MPSMatrixSolveLUClass{class: objc.GetClass("MPSMatrixSolveLU")}
	})
	return _MPSMatrixSolveLUClass
}

// GetMPSMatrixSolveLUClass returns the class object for MPSMatrixSolveLU.
func GetMPSMatrixSolveLUClass() MPSMatrixSolveLUClass {
	return getMPSMatrixSolveLUClass()
}

type MPSMatrixSolveLUClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSolveLUClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSolveLUClass) Alloc() MPSMatrixSolveLU {
	rv := objc.Send[MPSMatrixSolveLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the solution of a linear system of equations using
// an LU factorization.
//
// # Overview
//
// This kernel finds the solution matrix to the system op(A) * X = B, where:
//
// - op(A) is Aᵀ or A - X is the resulting matrix of solutions - B is the
// array of right hand sides for which the equations are to be solved
//
// # Initializers
//
//   - [MPSMatrixSolveLU.InitWithDeviceTransposeOrderNumberOfRightHandSides]
//
// # Instance Methods
//
//   - [MPSMatrixSolveLU.EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU
type MPSMatrixSolveLU struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixSolveLUFromID constructs a [MPSMatrixSolveLU] from an objc.ID.
//
// A kernel for computing the solution of a linear system of equations using
// an LU factorization.
func MPSMatrixSolveLUFromID(id objc.ID) MPSMatrixSolveLU {
	return MPSMatrixSolveLU{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixSolveLU adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSolveLU] class.
//
// # Initializers
//
//   - [IMPSMatrixSolveLU.InitWithDeviceTransposeOrderNumberOfRightHandSides]
//
// # Instance Methods
//
//   - [IMPSMatrixSolveLU.EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU
type IMPSMatrixSolveLU interface {
	IMPSMatrixBinaryKernel

	// Topic: Initializers

	InitWithDeviceTransposeOrderNumberOfRightHandSides(device metal.MTLDevice, transpose bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveLU

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, pivotIndices IMPSMatrix, solutionMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixSolveLU) Init() MPSMatrixSolveLU {
	rv := objc.Send[MPSMatrixSolveLU](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSolveLU) Autorelease() MPSMatrixSolveLU {
	rv := objc.Send[MPSMatrixSolveLU](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSolveLU creates a new MPSMatrixSolveLU instance.
func NewMPSMatrixSolveLU() MPSMatrixSolveLU {
	class := getMPSMatrixSolveLUClass()
	rv := objc.Send[MPSMatrixSolveLU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSolveLUWithCoder(aDecoder foundation.INSCoder) MPSMatrixSolveLU {
	instance := getMPSMatrixSolveLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSolveLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixSolveLUWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSolveLU {
	instance := getMPSMatrixSolveLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSolveLUFromID(rv)
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
func NewMatrixSolveLUWithDevice(device metal.MTLDevice) MPSMatrixSolveLU {
	instance := getMPSMatrixSolveLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSolveLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU/init(device:transpose:order:numberOfRightHandSides:)
func NewMatrixSolveLUWithDeviceTransposeOrderNumberOfRightHandSides(device metal.MTLDevice, transpose bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveLU {
	instance := getMPSMatrixSolveLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:transpose:order:numberOfRightHandSides:"), device, transpose, order, numberOfRightHandSides)
	return MPSMatrixSolveLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU/init(device:transpose:order:numberOfRightHandSides:)
func (m MPSMatrixSolveLU) InitWithDeviceTransposeOrderNumberOfRightHandSides(device metal.MTLDevice, transpose bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveLU {
	rv := objc.Send[MPSMatrixSolveLU](m.ID, objc.Sel("initWithDevice:transpose:order:numberOfRightHandSides:"), device, transpose, order, numberOfRightHandSides)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU/encode(commandBuffer:sourceMatrix:rightHandSideMatrix:pivotIndices:solutionMatrix:)
func (m MPSMatrixSolveLU) EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, pivotIndices IMPSMatrix, solutionMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:pivotIndices:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, pivotIndices, solutionMatrix)
}
