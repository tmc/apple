// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayBinaryKernel] class.
var (
	_MPSNDArrayBinaryKernelClass     MPSNDArrayBinaryKernelClass
	_MPSNDArrayBinaryKernelClassOnce sync.Once
)

func getMPSNDArrayBinaryKernelClass() MPSNDArrayBinaryKernelClass {
	_MPSNDArrayBinaryKernelClassOnce.Do(func() {
		_MPSNDArrayBinaryKernelClass = MPSNDArrayBinaryKernelClass{class: objc.GetClass("MPSNDArrayBinaryKernel")}
	})
	return _MPSNDArrayBinaryKernelClass
}

// GetMPSNDArrayBinaryKernelClass returns the class object for MPSNDArrayBinaryKernel.
func GetMPSNDArrayBinaryKernelClass() MPSNDArrayBinaryKernelClass {
	return getMPSNDArrayBinaryKernelClass()
}

type MPSNDArrayBinaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayBinaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayBinaryKernelClass) Alloc() MPSNDArrayBinaryKernel {
	rv := objc.Send[MPSNDArrayBinaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArray]
//   - [MPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray]
//   - [MPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray]
//   - [MPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel
type MPSNDArrayBinaryKernel struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayBinaryKernelFromID constructs a [MPSNDArrayBinaryKernel] from an objc.ID.
func MPSNDArrayBinaryKernelFromID(id objc.ID) MPSNDArrayBinaryKernel {
	return MPSNDArrayBinaryKernel{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayBinaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayBinaryKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArray]
//   - [IMPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray]
//   - [IMPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray]
//   - [IMPSNDArrayBinaryKernel.EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel
type IMPSNDArrayBinaryKernel interface {
	IMPSNDArrayMultiaryKernel

	// Topic: Instance Methods

	EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray) IMPSNDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, destination IMPSNDArray)
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, outGradientState IMPSState, destination IMPSNDArray)
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray
}

// Init initializes the instance.
func (n MPSNDArrayBinaryKernel) Init() MPSNDArrayBinaryKernel {
	rv := objc.Send[MPSNDArrayBinaryKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayBinaryKernel) Autorelease() MPSNDArrayBinaryKernel {
	rv := objc.Send[MPSNDArrayBinaryKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayBinaryKernel creates a new MPSNDArrayBinaryKernel instance.
func NewMPSNDArrayBinaryKernel() MPSNDArrayBinaryKernel {
	class := getMPSNDArrayBinaryKernelClass()
	rv := objc.Send[MPSNDArrayBinaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayBinaryKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayBinaryKernel {
	instance := getMPSNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(coder:device:)
func NewNDArrayBinaryKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayBinaryKernel {
	instance := getMPSNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(device:)
func NewNDArrayBinaryKernelWithDevice(device metal.MTLDevice) MPSNDArrayBinaryKernel {
	instance := getMPSNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayBinaryKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayBinaryKernel {
	instance := getMPSNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:)
func (n MPSNDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:"), cmdBuf, primarySourceArray, secondarySourceArray)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:destinationArray:)
func (n MPSNDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:resultState:destinationArray:)
func (n MPSNDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, outGradientState IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, outGradientState, destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:resultState:outputStateIsTemporary:)
func (n MPSNDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:outputStateIsTemporary:"), cmdBuf, primarySourceArray, secondarySourceArray, outGradientState, outputStateIsTemporary)
	return MPSNDArrayFromID(rv)
}
