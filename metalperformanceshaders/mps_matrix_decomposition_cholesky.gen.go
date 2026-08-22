// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixDecompositionCholesky] class.
var (
	_MPSMatrixDecompositionCholeskyClass     MPSMatrixDecompositionCholeskyClass
	_MPSMatrixDecompositionCholeskyClassOnce sync.Once
)

func getMPSMatrixDecompositionCholeskyClass() MPSMatrixDecompositionCholeskyClass {
	_MPSMatrixDecompositionCholeskyClassOnce.Do(func() {
		_MPSMatrixDecompositionCholeskyClass = MPSMatrixDecompositionCholeskyClass{class: objc.GetClass("MPSMatrixDecompositionCholesky")}
	})
	return _MPSMatrixDecompositionCholeskyClass
}

// GetMPSMatrixDecompositionCholeskyClass returns the class object for MPSMatrixDecompositionCholesky.
func GetMPSMatrixDecompositionCholeskyClass() MPSMatrixDecompositionCholeskyClass {
	return getMPSMatrixDecompositionCholeskyClass()
}

type MPSMatrixDecompositionCholeskyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixDecompositionCholeskyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixDecompositionCholeskyClass) Alloc() MPSMatrixDecompositionCholesky {
	rv := objc.Send[MPSMatrixDecompositionCholesky](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for computing the Cholesky factorization of a matrix.
//
// # Overview
//
// This kernel computes one of the following factorizations of a matrix A:
//
// - A = LLᵀ
// - A = UᵀU
//
// where:
//
// - A is a symmetric positive-definite matrix for which the factorization is
// to be computed - L is the lower triangular matrix - U is the upper
// triangular matrix
//
// # Initializers
//
//   - [MPSMatrixDecompositionCholesky.InitWithDeviceLowerOrder]
//
// # Instance Methods
//
//   - [MPSMatrixDecompositionCholesky.EncodeToCommandBufferSourceMatrixResultMatrixStatus]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky
type MPSMatrixDecompositionCholesky struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixDecompositionCholeskyFromID constructs a [MPSMatrixDecompositionCholesky] from an objc.ID.
//
// A kernel for computing the Cholesky factorization of a matrix.
func MPSMatrixDecompositionCholeskyFromID(id objc.ID) MPSMatrixDecompositionCholesky {
	return MPSMatrixDecompositionCholesky{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixDecompositionCholesky adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixDecompositionCholesky] class.
//
// # Initializers
//
//   - [IMPSMatrixDecompositionCholesky.InitWithDeviceLowerOrder]
//
// # Instance Methods
//
//   - [IMPSMatrixDecompositionCholesky.EncodeToCommandBufferSourceMatrixResultMatrixStatus]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky
type IMPSMatrixDecompositionCholesky interface {
	IMPSMatrixUnaryKernel

	// Topic: Initializers

	InitWithDeviceLowerOrder(device metal.MTLDevice, lower bool, order uint) MPSMatrixDecompositionCholesky

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, resultMatrix IMPSMatrix, status metal.MTLBuffer)
}

// Init initializes the instance.
func (m MPSMatrixDecompositionCholesky) Init() MPSMatrixDecompositionCholesky {
	rv := objc.Send[MPSMatrixDecompositionCholesky](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixDecompositionCholesky) Autorelease() MPSMatrixDecompositionCholesky {
	rv := objc.Send[MPSMatrixDecompositionCholesky](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixDecompositionCholesky creates a new MPSMatrixDecompositionCholesky instance.
func NewMPSMatrixDecompositionCholesky() MPSMatrixDecompositionCholesky {
	class := getMPSMatrixDecompositionCholeskyClass()
	rv := objc.Send[MPSMatrixDecompositionCholesky](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixDecompositionCholeskyWithCoder(aDecoder foundation.INSCoder) MPSMatrixDecompositionCholesky {
	instance := getMPSMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixDecompositionCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixDecompositionCholeskyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixDecompositionCholesky {
	instance := getMPSMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixDecompositionCholeskyFromID(rv)
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
func NewMatrixDecompositionCholeskyWithDevice(device metal.MTLDevice) MPSMatrixDecompositionCholesky {
	instance := getMPSMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixDecompositionCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky/init(device:lower:order:)
func NewMatrixDecompositionCholeskyWithDeviceLowerOrder(device metal.MTLDevice, lower bool, order uint) MPSMatrixDecompositionCholesky {
	instance := getMPSMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lower:order:"), device, lower, order)
	return MPSMatrixDecompositionCholeskyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky/init(device:lower:order:)
func (m MPSMatrixDecompositionCholesky) InitWithDeviceLowerOrder(device metal.MTLDevice, lower bool, order uint) MPSMatrixDecompositionCholesky {
	rv := objc.Send[MPSMatrixDecompositionCholesky](m.ID, objc.Sel("initWithDevice:lower:order:"), device, lower, order)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky/encode(commandBuffer:sourceMatrix:resultMatrix:status:)
func (m MPSMatrixDecompositionCholesky) EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, resultMatrix IMPSMatrix, status metal.MTLBuffer) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:status:"), commandBuffer, sourceMatrix, resultMatrix, status)
}
