// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayUnaryGradientKernel] class.
var (
	_MPSNDArrayUnaryGradientKernelClass     MPSNDArrayUnaryGradientKernelClass
	_MPSNDArrayUnaryGradientKernelClassOnce sync.Once
)

func getMPSNDArrayUnaryGradientKernelClass() MPSNDArrayUnaryGradientKernelClass {
	_MPSNDArrayUnaryGradientKernelClassOnce.Do(func() {
		_MPSNDArrayUnaryGradientKernelClass = MPSNDArrayUnaryGradientKernelClass{class: objc.GetClass("MPSNDArrayUnaryGradientKernel")}
	})
	return _MPSNDArrayUnaryGradientKernelClass
}

// GetMPSNDArrayUnaryGradientKernelClass returns the class object for MPSNDArrayUnaryGradientKernel.
func GetMPSNDArrayUnaryGradientKernelClass() MPSNDArrayUnaryGradientKernelClass {
	return getMPSNDArrayUnaryGradientKernelClass()
}

type MPSNDArrayUnaryGradientKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayUnaryGradientKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayUnaryGradientKernelClass) Alloc() MPSNDArrayUnaryGradientKernel {
	rv := objc.Send[MPSNDArrayUnaryGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayUnaryGradientKernel.EncodeToCommandBufferSourceArraySourceGradientGradientState]
//   - [MPSNDArrayUnaryGradientKernel.EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel
type MPSNDArrayUnaryGradientKernel struct {
	MPSNDArrayMultiaryGradientKernel
}

// MPSNDArrayUnaryGradientKernelFromID constructs a [MPSNDArrayUnaryGradientKernel] from an objc.ID.
func MPSNDArrayUnaryGradientKernelFromID(id objc.ID) MPSNDArrayUnaryGradientKernel {
	return MPSNDArrayUnaryGradientKernel{MPSNDArrayMultiaryGradientKernel: MPSNDArrayMultiaryGradientKernelFromID(id)}
}

// NOTE: MPSNDArrayUnaryGradientKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayUnaryGradientKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayUnaryGradientKernel.EncodeToCommandBufferSourceArraySourceGradientGradientState]
//   - [IMPSNDArrayUnaryGradientKernel.EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel
type IMPSNDArrayUnaryGradientKernel interface {
	IMPSNDArrayMultiaryGradientKernel

	// Topic: Instance Methods

	EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray
	EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray)
}

// Init initializes the instance.
func (n MPSNDArrayUnaryGradientKernel) Init() MPSNDArrayUnaryGradientKernel {
	rv := objc.Send[MPSNDArrayUnaryGradientKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayUnaryGradientKernel) Autorelease() MPSNDArrayUnaryGradientKernel {
	rv := objc.Send[MPSNDArrayUnaryGradientKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayUnaryGradientKernel creates a new MPSNDArrayUnaryGradientKernel instance.
func NewMPSNDArrayUnaryGradientKernel() MPSNDArrayUnaryGradientKernel {
	class := getMPSNDArrayUnaryGradientKernelClass()
	rv := objc.Send[MPSNDArrayUnaryGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayUnaryGradientKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayUnaryGradientKernel {
	instance := getMPSNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayUnaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/init(coder:device:)
func NewNDArrayUnaryGradientKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayUnaryGradientKernel {
	instance := getMPSNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayUnaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/init(device:)
func NewNDArrayUnaryGradientKernelWithDevice(device metal.MTLDevice) MPSNDArrayUnaryGradientKernel {
	instance := getMPSNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayUnaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayUnaryGradientKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayUnaryGradientKernel {
	instance := getMPSNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayUnaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayUnaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayUnaryGradientKernel {
	instance := getMPSNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayUnaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/encode(to:sourceArray:sourceGradient:gradientState:)
func (n MPSNDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:"), cmdBuf, sourceArray, gradient, state)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/encode(to:sourceArray:sourceGradient:gradientState:destinationArray:)
func (n MPSNDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, sourceArray, gradient, state, destination)
}
