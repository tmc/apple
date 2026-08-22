// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixSolveCholesky] class.
var (
	_MPSMatrixSolveCholeskyClass     MPSMatrixSolveCholeskyClass
	_MPSMatrixSolveCholeskyClassOnce sync.Once
)

func getMPSMatrixSolveCholeskyClass() MPSMatrixSolveCholeskyClass {
	_MPSMatrixSolveCholeskyClassOnce.Do(func() {
		_MPSMatrixSolveCholeskyClass = MPSMatrixSolveCholeskyClass{class: objc.GetClass("MPSMatrixSolveCholesky")}
	})
	return _MPSMatrixSolveCholeskyClass
}

// GetMPSMatrixSolveCholeskyClass returns the class object for MPSMatrixSolveCholesky.
func GetMPSMatrixSolveCholeskyClass() MPSMatrixSolveCholeskyClass {
	return getMPSMatrixSolveCholeskyClass()
}

type MPSMatrixSolveCholeskyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSolveCholeskyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSolveCholeskyClass) Alloc() MPSMatrixSolveCholesky {
	rv := objc.Send[MPSMatrixSolveCholesky](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the solution of a linear system of equations using a
// Cholesky factorization.
//
// # Overview
//
// This kernel finds the solution matrix to the system AX=B, where:
//
// - A is a symmetric positive-definite matrix - X is the resulting matrix of
// solutions - B is the array of right-hand-sides for which the equations are
// to be solved
//
// # Initializers
//
//   - [MPSMatrixSolveCholesky.InitWithDeviceUpperOrderNumberOfRightHandSides]
//
// # Instance Methods
//
//   - [MPSMatrixSolveCholesky.EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky
type MPSMatrixSolveCholesky struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixSolveCholeskyFromID constructs a [MPSMatrixSolveCholesky] from an objc.ID.
//
// A kernel for computing the solution of a linear system of equations using a
// Cholesky factorization.
func MPSMatrixSolveCholeskyFromID(id objc.ID) MPSMatrixSolveCholesky {
	return MPSMatrixSolveCholesky{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixSolveCholesky adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSolveCholesky] class.
//
// # Initializers
//
//   - [IMPSMatrixSolveCholesky.InitWithDeviceUpperOrderNumberOfRightHandSides]
//
// # Instance Methods
//
//   - [IMPSMatrixSolveCholesky.EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky
type IMPSMatrixSolveCholesky interface {
	IMPSMatrixBinaryKernel

	// Topic: Initializers

	InitWithDeviceUpperOrderNumberOfRightHandSides(device metal.MTLDevice, upper bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveCholesky

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, solutionMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixSolveCholesky) Init() MPSMatrixSolveCholesky {
	rv := objc.Send[MPSMatrixSolveCholesky](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSolveCholesky) Autorelease() MPSMatrixSolveCholesky {
	rv := objc.Send[MPSMatrixSolveCholesky](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSolveCholesky creates a new MPSMatrixSolveCholesky instance.
func NewMPSMatrixSolveCholesky() MPSMatrixSolveCholesky {
	class := getMPSMatrixSolveCholeskyClass()
	rv := objc.Send[MPSMatrixSolveCholesky](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSolveCholeskyWithCoder(aDecoder foundation.INSCoder) MPSMatrixSolveCholesky {
	instance := getMPSMatrixSolveCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSolveCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixSolveCholeskyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSolveCholesky {
	instance := getMPSMatrixSolveCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSolveCholeskyFromID(rv)
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
func NewMatrixSolveCholeskyWithDevice(device metal.MTLDevice) MPSMatrixSolveCholesky {
	instance := getMPSMatrixSolveCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSolveCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky/init(device:upper:order:numberOfRightHandSides:)
func NewMatrixSolveCholeskyWithDeviceUpperOrderNumberOfRightHandSides(device metal.MTLDevice, upper bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveCholesky {
	instance := getMPSMatrixSolveCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:upper:order:numberOfRightHandSides:"), device, upper, order, numberOfRightHandSides)
	return MPSMatrixSolveCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky/init(device:upper:order:numberOfRightHandSides:)
func (m MPSMatrixSolveCholesky) InitWithDeviceUpperOrderNumberOfRightHandSides(device metal.MTLDevice, upper bool, order uint, numberOfRightHandSides uint) MPSMatrixSolveCholesky {
	rv := objc.Send[MPSMatrixSolveCholesky](m.ID, objc.Sel("initWithDevice:upper:order:numberOfRightHandSides:"), device, upper, order, numberOfRightHandSides)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky/encode(commandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:)
func (m MPSMatrixSolveCholesky) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, solutionMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, solutionMatrix)
}
