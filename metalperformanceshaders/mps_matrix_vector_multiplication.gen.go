// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixVectorMultiplication] class.
var (
	_MPSMatrixVectorMultiplicationClass     MPSMatrixVectorMultiplicationClass
	_MPSMatrixVectorMultiplicationClassOnce sync.Once
)

func getMPSMatrixVectorMultiplicationClass() MPSMatrixVectorMultiplicationClass {
	_MPSMatrixVectorMultiplicationClassOnce.Do(func() {
		_MPSMatrixVectorMultiplicationClass = MPSMatrixVectorMultiplicationClass{class: objc.GetClass("MPSMatrixVectorMultiplication")}
	})
	return _MPSMatrixVectorMultiplicationClass
}

// GetMPSMatrixVectorMultiplicationClass returns the class object for MPSMatrixVectorMultiplication.
func GetMPSMatrixVectorMultiplicationClass() MPSMatrixVectorMultiplicationClass {
	return getMPSMatrixVectorMultiplicationClass()
}

type MPSMatrixVectorMultiplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixVectorMultiplicationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixVectorMultiplicationClass) Alloc() MPSMatrixVectorMultiplication {
	rv := objc.Send[MPSMatrixVectorMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A matrix-vector multiplication kernel
//
// # Initializers
//
//   - [MPSMatrixVectorMultiplication.InitWithDeviceTransposeRowsColumnsAlphaBeta]
//   - [MPSMatrixVectorMultiplication.InitWithDeviceRowsColumns]
//
// # Instance Methods
//
//   - [MPSMatrixVectorMultiplication.EncodeToCommandBufferInputMatrixInputVectorResultVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication
type MPSMatrixVectorMultiplication struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixVectorMultiplicationFromID constructs a [MPSMatrixVectorMultiplication] from an objc.ID.
//
// A matrix-vector multiplication kernel
func MPSMatrixVectorMultiplicationFromID(id objc.ID) MPSMatrixVectorMultiplication {
	return MPSMatrixVectorMultiplication{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixVectorMultiplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixVectorMultiplication] class.
//
// # Initializers
//
//   - [IMPSMatrixVectorMultiplication.InitWithDeviceTransposeRowsColumnsAlphaBeta]
//   - [IMPSMatrixVectorMultiplication.InitWithDeviceRowsColumns]
//
// # Instance Methods
//
//   - [IMPSMatrixVectorMultiplication.EncodeToCommandBufferInputMatrixInputVectorResultVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication
type IMPSMatrixVectorMultiplication interface {
	IMPSMatrixBinaryKernel

	// Topic: Initializers

	InitWithDeviceTransposeRowsColumnsAlphaBeta(device metal.MTLDevice, transpose bool, rows uint, columns uint, alpha float64, beta float64) MPSMatrixVectorMultiplication
	InitWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixVectorMultiplication

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, inputVector IMPSVector, resultVector IMPSVector)
}

// Init initializes the instance.
func (m MPSMatrixVectorMultiplication) Init() MPSMatrixVectorMultiplication {
	rv := objc.Send[MPSMatrixVectorMultiplication](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixVectorMultiplication) Autorelease() MPSMatrixVectorMultiplication {
	rv := objc.Send[MPSMatrixVectorMultiplication](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixVectorMultiplication creates a new MPSMatrixVectorMultiplication instance.
func NewMPSMatrixVectorMultiplication() MPSMatrixVectorMultiplication {
	class := getMPSMatrixVectorMultiplicationClass()
	rv := objc.Send[MPSMatrixVectorMultiplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixVectorMultiplicationWithCoder(aDecoder foundation.INSCoder) MPSMatrixVectorMultiplication {
	instance := getMPSMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixVectorMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixVectorMultiplicationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixVectorMultiplication {
	instance := getMPSMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixVectorMultiplicationFromID(rv)
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
func NewMatrixVectorMultiplicationWithDevice(device metal.MTLDevice) MPSMatrixVectorMultiplication {
	instance := getMPSMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixVectorMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication/init(device:rows:columns:)
func NewMatrixVectorMultiplicationWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixVectorMultiplication {
	instance := getMPSMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	return MPSMatrixVectorMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication/init(device:transpose:rows:columns:alpha:beta:)
func NewMatrixVectorMultiplicationWithDeviceTransposeRowsColumnsAlphaBeta(device metal.MTLDevice, transpose bool, rows uint, columns uint, alpha float64, beta float64) MPSMatrixVectorMultiplication {
	instance := getMPSMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:transpose:rows:columns:alpha:beta:"), device, transpose, rows, columns, alpha, beta)
	return MPSMatrixVectorMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication/init(device:transpose:rows:columns:alpha:beta:)
func (m MPSMatrixVectorMultiplication) InitWithDeviceTransposeRowsColumnsAlphaBeta(device metal.MTLDevice, transpose bool, rows uint, columns uint, alpha float64, beta float64) MPSMatrixVectorMultiplication {
	rv := objc.Send[MPSMatrixVectorMultiplication](m.ID, objc.Sel("initWithDevice:transpose:rows:columns:alpha:beta:"), device, transpose, rows, columns, alpha, beta)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication/init(device:rows:columns:)
func (m MPSMatrixVectorMultiplication) InitWithDeviceRowsColumns(device metal.MTLDevice, rows uint, columns uint) MPSMatrixVectorMultiplication {
	rv := objc.Send[MPSMatrixVectorMultiplication](m.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication/encode(commandBuffer:inputMatrix:inputVector:resultVector:)
func (m MPSMatrixVectorMultiplication) EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, inputVector IMPSVector, resultVector IMPSVector) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:inputVector:resultVector:"), commandBuffer, inputMatrix, inputVector, resultVector)
}
