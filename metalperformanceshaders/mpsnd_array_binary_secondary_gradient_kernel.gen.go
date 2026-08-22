// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayBinarySecondaryGradientKernel] class.
var (
	_MPSNDArrayBinarySecondaryGradientKernelClass     MPSNDArrayBinarySecondaryGradientKernelClass
	_MPSNDArrayBinarySecondaryGradientKernelClassOnce sync.Once
)

func getMPSNDArrayBinarySecondaryGradientKernelClass() MPSNDArrayBinarySecondaryGradientKernelClass {
	_MPSNDArrayBinarySecondaryGradientKernelClassOnce.Do(func() {
		_MPSNDArrayBinarySecondaryGradientKernelClass = MPSNDArrayBinarySecondaryGradientKernelClass{class: objc.GetClass("MPSNDArrayBinarySecondaryGradientKernel")}
	})
	return _MPSNDArrayBinarySecondaryGradientKernelClass
}

// GetMPSNDArrayBinarySecondaryGradientKernelClass returns the class object for MPSNDArrayBinarySecondaryGradientKernel.
func GetMPSNDArrayBinarySecondaryGradientKernelClass() MPSNDArrayBinarySecondaryGradientKernelClass {
	return getMPSNDArrayBinarySecondaryGradientKernelClass()
}

type MPSNDArrayBinarySecondaryGradientKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayBinarySecondaryGradientKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayBinarySecondaryGradientKernelClass) Alloc() MPSNDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinarySecondaryGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayBinarySecondaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState]
//   - [MPSNDArrayBinarySecondaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel
type MPSNDArrayBinarySecondaryGradientKernel struct {
	MPSNDArrayMultiaryGradientKernel
}

// MPSNDArrayBinarySecondaryGradientKernelFromID constructs a [MPSNDArrayBinarySecondaryGradientKernel] from an objc.ID.
func MPSNDArrayBinarySecondaryGradientKernelFromID(id objc.ID) MPSNDArrayBinarySecondaryGradientKernel {
	return MPSNDArrayBinarySecondaryGradientKernel{MPSNDArrayMultiaryGradientKernel: MPSNDArrayMultiaryGradientKernelFromID(id)}
}

// NOTE: MPSNDArrayBinarySecondaryGradientKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayBinarySecondaryGradientKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayBinarySecondaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState]
//   - [IMPSNDArrayBinarySecondaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel
type IMPSNDArrayBinarySecondaryGradientKernel interface {
	IMPSNDArrayMultiaryGradientKernel

	// Topic: Instance Methods

	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray)
}

// Init initializes the instance.
func (n MPSNDArrayBinarySecondaryGradientKernel) Init() MPSNDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinarySecondaryGradientKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayBinarySecondaryGradientKernel) Autorelease() MPSNDArrayBinarySecondaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinarySecondaryGradientKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayBinarySecondaryGradientKernel creates a new MPSNDArrayBinarySecondaryGradientKernel instance.
func NewMPSNDArrayBinarySecondaryGradientKernel() MPSNDArrayBinarySecondaryGradientKernel {
	class := getMPSNDArrayBinarySecondaryGradientKernelClass()
	rv := objc.Send[MPSNDArrayBinarySecondaryGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayBinarySecondaryGradientKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayBinarySecondaryGradientKernel {
	instance := getMPSNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayBinarySecondaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel/init(coder:device:)
func NewNDArrayBinarySecondaryGradientKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayBinarySecondaryGradientKernel {
	instance := getMPSNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayBinarySecondaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel/init(device:)
func NewNDArrayBinarySecondaryGradientKernelWithDevice(device metal.MTLDevice) MPSNDArrayBinarySecondaryGradientKernel {
	instance := getMPSNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayBinarySecondaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayBinarySecondaryGradientKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayBinarySecondaryGradientKernel {
	instance := getMPSNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayBinarySecondaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayBinarySecondaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayBinarySecondaryGradientKernel {
	instance := getMPSNDArrayBinarySecondaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayBinarySecondaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:)
func (n MPSNDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinarySecondaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:)
func (n MPSNDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state, destination)
}
