// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayBinaryPrimaryGradientKernel] class.
var (
	_MPSNDArrayBinaryPrimaryGradientKernelClass     MPSNDArrayBinaryPrimaryGradientKernelClass
	_MPSNDArrayBinaryPrimaryGradientKernelClassOnce sync.Once
)

func getMPSNDArrayBinaryPrimaryGradientKernelClass() MPSNDArrayBinaryPrimaryGradientKernelClass {
	_MPSNDArrayBinaryPrimaryGradientKernelClassOnce.Do(func() {
		_MPSNDArrayBinaryPrimaryGradientKernelClass = MPSNDArrayBinaryPrimaryGradientKernelClass{class: objc.GetClass("MPSNDArrayBinaryPrimaryGradientKernel")}
	})
	return _MPSNDArrayBinaryPrimaryGradientKernelClass
}

// GetMPSNDArrayBinaryPrimaryGradientKernelClass returns the class object for MPSNDArrayBinaryPrimaryGradientKernel.
func GetMPSNDArrayBinaryPrimaryGradientKernelClass() MPSNDArrayBinaryPrimaryGradientKernelClass {
	return getMPSNDArrayBinaryPrimaryGradientKernelClass()
}

type MPSNDArrayBinaryPrimaryGradientKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayBinaryPrimaryGradientKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayBinaryPrimaryGradientKernelClass) Alloc() MPSNDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinaryPrimaryGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayBinaryPrimaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState]
//   - [MPSNDArrayBinaryPrimaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel
type MPSNDArrayBinaryPrimaryGradientKernel struct {
	MPSNDArrayMultiaryGradientKernel
}

// MPSNDArrayBinaryPrimaryGradientKernelFromID constructs a [MPSNDArrayBinaryPrimaryGradientKernel] from an objc.ID.
func MPSNDArrayBinaryPrimaryGradientKernelFromID(id objc.ID) MPSNDArrayBinaryPrimaryGradientKernel {
	return MPSNDArrayBinaryPrimaryGradientKernel{MPSNDArrayMultiaryGradientKernel: MPSNDArrayMultiaryGradientKernelFromID(id)}
}

// NOTE: MPSNDArrayBinaryPrimaryGradientKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayBinaryPrimaryGradientKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayBinaryPrimaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState]
//   - [IMPSNDArrayBinaryPrimaryGradientKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel
type IMPSNDArrayBinaryPrimaryGradientKernel interface {
	IMPSNDArrayMultiaryGradientKernel

	// Topic: Instance Methods

	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray)
}

// Init initializes the instance.
func (n MPSNDArrayBinaryPrimaryGradientKernel) Init() MPSNDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinaryPrimaryGradientKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayBinaryPrimaryGradientKernel) Autorelease() MPSNDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[MPSNDArrayBinaryPrimaryGradientKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayBinaryPrimaryGradientKernel creates a new MPSNDArrayBinaryPrimaryGradientKernel instance.
func NewMPSNDArrayBinaryPrimaryGradientKernel() MPSNDArrayBinaryPrimaryGradientKernel {
	class := getMPSNDArrayBinaryPrimaryGradientKernelClass()
	rv := objc.Send[MPSNDArrayBinaryPrimaryGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayBinaryPrimaryGradientKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayBinaryPrimaryGradientKernel {
	instance := getMPSNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayBinaryPrimaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/init(coder:device:)
func NewNDArrayBinaryPrimaryGradientKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayBinaryPrimaryGradientKernel {
	instance := getMPSNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayBinaryPrimaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/init(device:)
func NewNDArrayBinaryPrimaryGradientKernelWithDevice(device metal.MTLDevice) MPSNDArrayBinaryPrimaryGradientKernel {
	instance := getMPSNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayBinaryPrimaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayBinaryPrimaryGradientKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayBinaryPrimaryGradientKernel {
	instance := getMPSNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayBinaryPrimaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayBinaryPrimaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayBinaryPrimaryGradientKernel {
	instance := getMPSNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayBinaryPrimaryGradientKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:)
func (n MPSNDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/encode(to:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:)
func (n MPSNDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, gradient IMPSNDArray, state IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state, destination)
}
