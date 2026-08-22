// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayMultiaryKernel] class.
var (
	_MPSNDArrayMultiaryKernelClass     MPSNDArrayMultiaryKernelClass
	_MPSNDArrayMultiaryKernelClassOnce sync.Once
)

func getMPSNDArrayMultiaryKernelClass() MPSNDArrayMultiaryKernelClass {
	_MPSNDArrayMultiaryKernelClassOnce.Do(func() {
		_MPSNDArrayMultiaryKernelClass = MPSNDArrayMultiaryKernelClass{class: objc.GetClass("MPSNDArrayMultiaryKernel")}
	})
	return _MPSNDArrayMultiaryKernelClass
}

// GetMPSNDArrayMultiaryKernelClass returns the class object for MPSNDArrayMultiaryKernel.
func GetMPSNDArrayMultiaryKernelClass() MPSNDArrayMultiaryKernelClass {
	return getMPSNDArrayMultiaryKernelClass()
}

type MPSNDArrayMultiaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayMultiaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayMultiaryKernelClass) Alloc() MPSNDArrayMultiaryKernel {
	rv := objc.Send[MPSNDArrayMultiaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayMultiaryKernel.EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray]
//   - [MPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArrays]
//   - [MPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysDestinationArray]
//   - [MPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysResultStateDestinationArray]
//   - [MPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel
type MPSNDArrayMultiaryKernel struct {
	MPSNDArrayMultiaryBase
}

// MPSNDArrayMultiaryKernelFromID constructs a [MPSNDArrayMultiaryKernel] from an objc.ID.
func MPSNDArrayMultiaryKernelFromID(id objc.ID) MPSNDArrayMultiaryKernel {
	return MPSNDArrayMultiaryKernel{MPSNDArrayMultiaryBase: MPSNDArrayMultiaryBaseFromID(id)}
}

// NOTE: MPSNDArrayMultiaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayMultiaryKernel] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayMultiaryKernel.EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray]
//   - [IMPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArrays]
//   - [IMPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysDestinationArray]
//   - [IMPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysResultStateDestinationArray]
//   - [IMPSNDArrayMultiaryKernel.EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel
type IMPSNDArrayMultiaryKernel interface {
	IMPSNDArrayMultiaryBase

	// Topic: Instance Methods

	EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder metal.MTLComputeCommandEncoder, commandBuffer metal.MTLCommandBuffer, sourceArrays []MPSNDArray, destination IMPSNDArray)
	EncodeToCommandBufferSourceArrays(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray) IMPSNDArray
	EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, destination IMPSNDArray)
	EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, outGradientState IMPSState, destination IMPSNDArray)
	EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray
}

// Init initializes the instance.
func (n MPSNDArrayMultiaryKernel) Init() MPSNDArrayMultiaryKernel {
	rv := objc.Send[MPSNDArrayMultiaryKernel](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayMultiaryKernel) Autorelease() MPSNDArrayMultiaryKernel {
	rv := objc.Send[MPSNDArrayMultiaryKernel](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayMultiaryKernel creates a new MPSNDArrayMultiaryKernel instance.
func NewMPSNDArrayMultiaryKernel() MPSNDArrayMultiaryKernel {
	class := getMPSNDArrayMultiaryKernelClass()
	rv := objc.Send[MPSNDArrayMultiaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayMultiaryKernelWithCoder(aDecoder foundation.INSCoder) MPSNDArrayMultiaryKernel {
	instance := getMPSNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayMultiaryKernelWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayMultiaryKernel {
	instance := getMPSNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayMultiaryKernelFromID(rv)
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
func NewNDArrayMultiaryKernelWithDevice(device metal.MTLDevice) MPSNDArrayMultiaryKernel {
	instance := getMPSNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayMultiaryKernelWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMultiaryKernel {
	instance := getMPSNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:commandBuffer:sourceArrays:destinationArray:)
func (n MPSNDArrayMultiaryKernel) EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder metal.MTLComputeCommandEncoder, commandBuffer metal.MTLCommandBuffer, sourceArrays []MPSNDArray, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandEncoder:commandBuffer:sourceArrays:destinationArray:"), encoder, commandBuffer, objectivec.IObjectSliceToNSArray(sourceArrays), destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:)
func (n MPSNDArrayMultiaryKernel) EncodeToCommandBufferSourceArrays(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:"), cmdBuf, objectivec.IObjectSliceToNSArray(sourceArrays))
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:destinationArray:)
func (n MPSNDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), cmdBuf, objectivec.IObjectSliceToNSArray(sourceArrays), destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:resultState:destinationArray:)
func (n MPSNDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, outGradientState IMPSState, destination IMPSNDArray) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:destinationArray:"), cmdBuf, objectivec.IObjectSliceToNSArray(sourceArrays), outGradientState, destination)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/encode(to:sourceArrays:resultState:outputStateIsTemporary:)
func (n MPSNDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf metal.MTLCommandBuffer, sourceArrays []MPSNDArray, outGradientState IMPSState, outputStateIsTemporary bool) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), cmdBuf, objectivec.IObjectSliceToNSArray(sourceArrays), outGradientState, outputStateIsTemporary)
	return MPSNDArrayFromID(rv)
}
