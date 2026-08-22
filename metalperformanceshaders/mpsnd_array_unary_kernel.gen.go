// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayUnaryKernel] class.
var (
	_MPSNDArrayUnaryKernelClass     MPSNDArrayUnaryKernelClass
	_MPSNDArrayUnaryKernelClassOnce sync.Once
)

func getMPSNDArrayUnaryKernelClass() MPSNDArrayUnaryKernelClass {
	_MPSNDArrayUnaryKernelClassOnce.Do(func() {
		_MPSNDArrayUnaryKernelClass = MPSNDArrayUnaryKernelClass{class: objc.GetClass("MPSNDArrayUnaryKernel")}
	})
	return _MPSNDArrayUnaryKernelClass
}

// GetMPSNDArrayUnaryKernelClass returns the class object for MPSNDArrayUnaryKernel.
func GetMPSNDArrayUnaryKernelClass() MPSNDArrayUnaryKernelClass {
	return getMPSNDArrayUnaryKernelClass()
}

type MPSNDArrayUnaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayUnaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayUnaryKernelClass) Alloc() MPSNDArrayUnaryKernel {
	rv := objc.Send[MPSNDArrayUnaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArray]
//   - [MPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayDestinationArray]
//   - [MPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayResultStateDestinationArray]
//   - [MPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel
type MPSNDArrayUnaryKernel struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayUnaryKernelFromID constructs a [MPSNDArrayUnaryKernel] from an objc.ID.
func MPSNDArrayUnaryKernelFromID(id objc.ID) MPSNDArrayUnaryKernel {
	return MPSNDArrayUnaryKernel{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayUnaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayUnaryKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArray]
//   - [IMPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayDestinationArray]
//   - [IMPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayResultStateDestinationArray]
//   - [IMPSNDArrayUnaryKernel.EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel
type IMPSNDArrayUnaryKernel interface {
	IMPSNDArrayMultiaryKernel

	// Topic: Instance Methods

	EncodeToCommandBufferSourceArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray) IMPSNDArray
	EncodeToCommandBufferSourceArrayDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, destination IMPSNDArray)
	EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, outGradientState IMPSState, destination IMPSNDArray)
	EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray
}

// Init initializes the instance.
func (n MPSNDArrayUnaryKernel) Init() MPSNDArrayUnaryKernel {
	rv := objc.Send[MPSNDArrayUnaryKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayUnaryKernel) Autorelease() MPSNDArrayUnaryKernel {
	rv := objc.Send[MPSNDArrayUnaryKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayUnaryKernel creates a new MPSNDArrayUnaryKernel instance.
func NewMPSNDArrayUnaryKernel() MPSNDArrayUnaryKernel {
	class := getMPSNDArrayUnaryKernelClass()
	rv := objc.Send[MPSNDArrayUnaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayUnaryKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayUnaryKernel {
	instance := getMPSNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/init(coder:device:)
func NewNDArrayUnaryKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayUnaryKernel {
	instance := getMPSNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/init(device:)
func NewNDArrayUnaryKernelWithDevice(device metal.MTLDevice) MPSNDArrayUnaryKernel {
	instance := getMPSNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayUnaryKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayUnaryKernel {
	instance := getMPSNDArrayUnaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayUnaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:)
func (n MPSNDArrayUnaryKernel) EncodeToCommandBufferSourceArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:"), cmdBuf, sourceArray)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:destinationArray:)
func (n MPSNDArrayUnaryKernel) EncodeToCommandBufferSourceArrayDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:destinationArray:"), cmdBuf, sourceArray, destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:resultState:destinationArray:)
func (n MPSNDArrayUnaryKernel) EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, outGradientState IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:destinationArray:"), cmdBuf, sourceArray, outGradientState, destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/encode(to:sourceArray:resultState:outputStateIsTemporary:)
func (n MPSNDArrayUnaryKernel) EncodeToCommandBufferSourceArrayResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:outputStateIsTemporary:"), cmdBuf, sourceArray, outGradientState, outputStateIsTemporary)
	return MPSNDArrayFromID(rv)
}
