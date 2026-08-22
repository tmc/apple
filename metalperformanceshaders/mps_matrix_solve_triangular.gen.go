// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixSolveTriangular] class.
var (
	_MPSMatrixSolveTriangularClass     MPSMatrixSolveTriangularClass
	_MPSMatrixSolveTriangularClassOnce sync.Once
)

func getMPSMatrixSolveTriangularClass() MPSMatrixSolveTriangularClass {
	_MPSMatrixSolveTriangularClassOnce.Do(func() {
		_MPSMatrixSolveTriangularClass = MPSMatrixSolveTriangularClass{class: objc.GetClass("MPSMatrixSolveTriangular")}
	})
	return _MPSMatrixSolveTriangularClass
}

// GetMPSMatrixSolveTriangularClass returns the class object for MPSMatrixSolveTriangular.
func GetMPSMatrixSolveTriangularClass() MPSMatrixSolveTriangularClass {
	return getMPSMatrixSolveTriangularClass()
}

type MPSMatrixSolveTriangularClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixSolveTriangularClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixSolveTriangularClass) Alloc() MPSMatrixSolveTriangular {
	rv := objc.Send[MPSMatrixSolveTriangular](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the solution of a linear system of equations using a
// triangular coefficient matrix.
//
// # Overview
//
// This kernel finds the solution matrix to the system op(A) * X = alpha * B
// or X * op(A) = alpha * B, where:
//
// - A is either an upper or lower triangular matrix - op(A) is either Aᵀ or
// A - X is the resulting matrix of solutions - B is the array of right hand
// sides for which the equations are to be solved
//
// # Initializers
//
//   - [MPSMatrixSolveTriangular.InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha]
//
// # Instance Methods
//
//   - [MPSMatrixSolveTriangular.EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular
type MPSMatrixSolveTriangular struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixSolveTriangularFromID constructs a [MPSMatrixSolveTriangular] from an objc.ID.
//
// A kernel for computing the solution of a linear system of equations using a
// triangular coefficient matrix.
func MPSMatrixSolveTriangularFromID(id objc.ID) MPSMatrixSolveTriangular {
	return MPSMatrixSolveTriangular{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixSolveTriangular adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixSolveTriangular] class.
//
// # Initializers
//
//   - [IMPSMatrixSolveTriangular.InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha]
//
// # Instance Methods
//
//   - [IMPSMatrixSolveTriangular.EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular
type IMPSMatrixSolveTriangular interface {
	IMPSMatrixBinaryKernel

	// Topic: Initializers

	InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device metal.MTLDevice, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MPSMatrixSolveTriangular

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, solutionMatrix IMPSMatrix)
}

// Init initializes the instance.
func (m MPSMatrixSolveTriangular) Init() MPSMatrixSolveTriangular {
	rv := objc.Send[MPSMatrixSolveTriangular](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixSolveTriangular) Autorelease() MPSMatrixSolveTriangular {
	rv := objc.Send[MPSMatrixSolveTriangular](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixSolveTriangular creates a new MPSMatrixSolveTriangular instance.
func NewMPSMatrixSolveTriangular() MPSMatrixSolveTriangular {
	class := getMPSMatrixSolveTriangularClass()
	rv := objc.Send[MPSMatrixSolveTriangular](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixSolveTriangularWithCoder(aDecoder foundation.INSCoder) MPSMatrixSolveTriangular {
	instance := getMPSMatrixSolveTriangularClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixSolveTriangularFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixSolveTriangularWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixSolveTriangular {
	instance := getMPSMatrixSolveTriangularClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixSolveTriangularFromID(rv)
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
func NewMatrixSolveTriangularWithDevice(device metal.MTLDevice) MPSMatrixSolveTriangular {
	instance := getMPSMatrixSolveTriangularClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixSolveTriangularFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular/init(device:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:)
func NewMatrixSolveTriangularWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device metal.MTLDevice, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MPSMatrixSolveTriangular {
	instance := getMPSMatrixSolveTriangularClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:"), device, right, upper, transpose, unit, order, numberOfRightHandSides, alpha)
	return MPSMatrixSolveTriangularFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular/init(device:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:)
func (m MPSMatrixSolveTriangular) InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device metal.MTLDevice, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MPSMatrixSolveTriangular {
	rv := objc.Send[MPSMatrixSolveTriangular](m.ID, objc.Sel("initWithDevice:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:"), device, right, upper, transpose, unit, order, numberOfRightHandSides, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular/encode(commandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:)
func (m MPSMatrixSolveTriangular) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, rightHandSideMatrix IMPSMatrix, solutionMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, solutionMatrix)
}
