// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayMultiaryGradientKernel] class.
var (
	_MPSNDArrayMultiaryGradientKernelClass     MPSNDArrayMultiaryGradientKernelClass
	_MPSNDArrayMultiaryGradientKernelClassOnce sync.Once
)

func getMPSNDArrayMultiaryGradientKernelClass() MPSNDArrayMultiaryGradientKernelClass {
	_MPSNDArrayMultiaryGradientKernelClassOnce.Do(func() {
		_MPSNDArrayMultiaryGradientKernelClass = MPSNDArrayMultiaryGradientKernelClass{class: objc.GetClass("MPSNDArrayMultiaryGradientKernel")}
	})
	return _MPSNDArrayMultiaryGradientKernelClass
}

// GetMPSNDArrayMultiaryGradientKernelClass returns the class object for MPSNDArrayMultiaryGradientKernel.
func GetMPSNDArrayMultiaryGradientKernelClass() MPSNDArrayMultiaryGradientKernelClass {
	return getMPSNDArrayMultiaryGradientKernelClass()
}

type MPSNDArrayMultiaryGradientKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayMultiaryGradientKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayMultiaryGradientKernelClass) Alloc() MPSNDArrayMultiaryGradientKernel {
	rv := objc.Send[MPSNDArrayMultiaryGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayMultiaryGradientKernel.InitWithDeviceSourceCountSourceGradientIndex]
//
// # Instance Methods
//
//   - [MPSNDArrayMultiaryGradientKernel.EncodeToCommandBufferSourceArraysSourceGradientGradientState]
//   - [MPSNDArrayMultiaryGradientKernel.EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel
type MPSNDArrayMultiaryGradientKernel struct {
	MPSNDArrayMultiaryBase
}

// MPSNDArrayMultiaryGradientKernelFromID constructs a [MPSNDArrayMultiaryGradientKernel] from an objc.ID.
func MPSNDArrayMultiaryGradientKernelFromID(id objc.ID) MPSNDArrayMultiaryGradientKernel {
	return MPSNDArrayMultiaryGradientKernel{MPSNDArrayMultiaryBase: MPSNDArrayMultiaryBaseFromID(id)}
}

// NOTE: MPSNDArrayMultiaryGradientKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayMultiaryGradientKernel] class.
//
// # Initializers
//
//   - [IMPSNDArrayMultiaryGradientKernel.InitWithDeviceSourceCountSourceGradientIndex]
//
// # Instance Methods
//
//   - [IMPSNDArrayMultiaryGradientKernel.EncodeToCommandBufferSourceArraysSourceGradientGradientState]
//   - [IMPSNDArrayMultiaryGradientKernel.EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel
type IMPSNDArrayMultiaryGradientKernel interface {
	IMPSNDArrayMultiaryBase

	// Topic: Initializers

	InitWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayMultiaryGradientKernel

	// Topic: Instance Methods

	EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, sources []MPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray
	EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sources []MPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray)
}

// Init initializes the instance.
func (n MPSNDArrayMultiaryGradientKernel) Init() MPSNDArrayMultiaryGradientKernel {
	rv := objc.Send[MPSNDArrayMultiaryGradientKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayMultiaryGradientKernel) Autorelease() MPSNDArrayMultiaryGradientKernel {
	rv := objc.Send[MPSNDArrayMultiaryGradientKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayMultiaryGradientKernel creates a new MPSNDArrayMultiaryGradientKernel instance.
func NewMPSNDArrayMultiaryGradientKernel() MPSNDArrayMultiaryGradientKernel {
	class := getMPSNDArrayMultiaryGradientKernelClass()
	rv := objc.Send[MPSNDArrayMultiaryGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayMultiaryGradientKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayMultiaryGradientKernel {
	instance := getMPSNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayMultiaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(coder:device:)
func NewNDArrayMultiaryGradientKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayMultiaryGradientKernel {
	instance := getMPSNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayMultiaryGradientKernelFromID(rv)
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
func NewNDArrayMultiaryGradientKernelWithDevice(device metal.MTLDevice) MPSNDArrayMultiaryGradientKernel {
	instance := getMPSNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayMultiaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMultiaryGradientKernel {
	instance := getMPSNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayMultiaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayMultiaryGradientKernel {
	instance := getMPSNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayMultiaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func (n MPSNDArrayMultiaryGradientKernel) InitWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayMultiaryGradientKernel {
	rv := objc.Send[MPSNDArrayMultiaryGradientKernel](n.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/encode(to:sourceArrays:sourceGradient:gradientState:)
func (n MPSNDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, sources []MPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:"), cmdBuf, objectivec.IObjectSliceToNSArray(sources), gradient, state)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/encode(to:sourceArrays:sourceGradient:gradientState:destinationArray:)
func (n MPSNDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sources []MPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:destinationArray:"), cmdBuf, objectivec.IObjectSliceToNSArray(sources), gradient, state, destination)
}
