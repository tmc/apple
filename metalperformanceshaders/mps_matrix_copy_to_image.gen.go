// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixCopyToImage] class.
var (
	_MPSMatrixCopyToImageClass     MPSMatrixCopyToImageClass
	_MPSMatrixCopyToImageClassOnce sync.Once
)

func getMPSMatrixCopyToImageClass() MPSMatrixCopyToImageClass {
	_MPSMatrixCopyToImageClassOnce.Do(func() {
		_MPSMatrixCopyToImageClass = MPSMatrixCopyToImageClass{class: objc.GetClass("MPSMatrixCopyToImage")}
	})
	return _MPSMatrixCopyToImageClass
}

// GetMPSMatrixCopyToImageClass returns the class object for MPSMatrixCopyToImage.
func GetMPSMatrixCopyToImageClass() MPSMatrixCopyToImageClass {
	return getMPSMatrixCopyToImageClass()
}

type MPSMatrixCopyToImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixCopyToImageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixCopyToImageClass) Alloc() MPSMatrixCopyToImage {
	rv := objc.Send[MPSMatrixCopyToImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that copies matrix data to a Metal Performance Shaders image.
//
// # Initializers
//
//   - [MPSMatrixCopyToImage.InitWithDeviceDataLayout]
//
// # Instance Properties
//
//   - [MPSMatrixCopyToImage.DataLayout]
//   - [MPSMatrixCopyToImage.SourceMatrixBatchIndex]
//   - [MPSMatrixCopyToImage.SetSourceMatrixBatchIndex]
//   - [MPSMatrixCopyToImage.SourceMatrixOrigin]
//   - [MPSMatrixCopyToImage.SetSourceMatrixOrigin]
//
// # Instance Methods
//
//   - [MPSMatrixCopyToImage.EncodeToCommandBufferSourceMatrixDestinationImage]
//   - [MPSMatrixCopyToImage.EncodeBatchToCommandBufferSourceMatrixDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage
type MPSMatrixCopyToImage struct {
	MPSKernel
}

// MPSMatrixCopyToImageFromID constructs a [MPSMatrixCopyToImage] from an objc.ID.
//
// A kernel that copies matrix data to a Metal Performance Shaders image.
func MPSMatrixCopyToImageFromID(id objc.ID) MPSMatrixCopyToImage {
	return MPSMatrixCopyToImage{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixCopyToImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixCopyToImage] class.
//
// # Initializers
//
//   - [IMPSMatrixCopyToImage.InitWithDeviceDataLayout]
//
// # Instance Properties
//
//   - [IMPSMatrixCopyToImage.DataLayout]
//   - [IMPSMatrixCopyToImage.SourceMatrixBatchIndex]
//   - [IMPSMatrixCopyToImage.SetSourceMatrixBatchIndex]
//   - [IMPSMatrixCopyToImage.SourceMatrixOrigin]
//   - [IMPSMatrixCopyToImage.SetSourceMatrixOrigin]
//
// # Instance Methods
//
//   - [IMPSMatrixCopyToImage.EncodeToCommandBufferSourceMatrixDestinationImage]
//   - [IMPSMatrixCopyToImage.EncodeBatchToCommandBufferSourceMatrixDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage
type IMPSMatrixCopyToImage interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSMatrixCopyToImage

	// Topic: Instance Properties

	DataLayout() MPSDataLayout
	SourceMatrixBatchIndex() uint
	SetSourceMatrixBatchIndex(value uint)
	SourceMatrixOrigin() metal.MTLOrigin
	SetSourceMatrixOrigin(value metal.MTLOrigin)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, destinationImage IMPSImage)
	EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, destinationImages MPSImageBatch)
}

// Init initializes the instance.
func (m MPSMatrixCopyToImage) Init() MPSMatrixCopyToImage {
	rv := objc.Send[MPSMatrixCopyToImage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixCopyToImage) Autorelease() MPSMatrixCopyToImage {
	rv := objc.Send[MPSMatrixCopyToImage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixCopyToImage creates a new MPSMatrixCopyToImage instance.
func NewMPSMatrixCopyToImage() MPSMatrixCopyToImage {
	class := getMPSMatrixCopyToImageClass()
	rv := objc.Send[MPSMatrixCopyToImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixCopyToImageWithCoder(aDecoder foundation.INSCoder) MPSMatrixCopyToImage {
	instance := getMPSMatrixCopyToImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixCopyToImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/init(coder:device:)
func NewMatrixCopyToImageWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixCopyToImage {
	instance := getMPSMatrixCopyToImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixCopyToImageFromID(rv)
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
func NewMatrixCopyToImageWithDevice(device metal.MTLDevice) MPSMatrixCopyToImage {
	instance := getMPSMatrixCopyToImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixCopyToImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/init(device:dataLayout:)
func NewMatrixCopyToImageWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSMatrixCopyToImage {
	instance := getMPSMatrixCopyToImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	return MPSMatrixCopyToImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/init(device:dataLayout:)
func (m MPSMatrixCopyToImage) InitWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSMatrixCopyToImage {
	rv := objc.Send[MPSMatrixCopyToImage](m.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/encode(commandBuffer:sourceMatrix:destinationImage:)
func (m MPSMatrixCopyToImage) EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, destinationImage IMPSImage) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:destinationImage:"), commandBuffer, sourceMatrix, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/encodeBatch(commandBuffer:sourceMatrix:destinationImages:)
func (m MPSMatrixCopyToImage) EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceMatrix IMPSMatrix, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeBatchToCommandBuffer:sourceMatrix:destinationImages:"), commandBuffer, sourceMatrix, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/dataLayout
func (m MPSMatrixCopyToImage) DataLayout() MPSDataLayout {
	rv := objc.Send[MPSDataLayout](m.ID, objc.Sel("dataLayout"))
	return MPSDataLayout(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/sourceMatrixBatchIndex
func (m MPSMatrixCopyToImage) SourceMatrixBatchIndex() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceMatrixBatchIndex"))
	return rv
}
func (m MPSMatrixCopyToImage) SetSourceMatrixBatchIndex(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceMatrixBatchIndex:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage/sourceMatrixOrigin
func (m MPSMatrixCopyToImage) SourceMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](m.ID, objc.Sel("sourceMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (m MPSMatrixCopyToImage) SetSourceMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceMatrixOrigin:"), value)
}
