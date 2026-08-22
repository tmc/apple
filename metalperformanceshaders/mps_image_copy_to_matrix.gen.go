// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageCopyToMatrix] class.
var (
	_MPSImageCopyToMatrixClass     MPSImageCopyToMatrixClass
	_MPSImageCopyToMatrixClassOnce sync.Once
)

func getMPSImageCopyToMatrixClass() MPSImageCopyToMatrixClass {
	_MPSImageCopyToMatrixClassOnce.Do(func() {
		_MPSImageCopyToMatrixClass = MPSImageCopyToMatrixClass{class: objc.GetClass("MPSImageCopyToMatrix")}
	})
	return _MPSImageCopyToMatrixClass
}

// GetMPSImageCopyToMatrixClass returns the class object for MPSImageCopyToMatrix.
func GetMPSImageCopyToMatrixClass() MPSImageCopyToMatrixClass {
	return getMPSImageCopyToMatrixClass()
}

type MPSImageCopyToMatrixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageCopyToMatrixClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageCopyToMatrixClass) Alloc() MPSImageCopyToMatrix {
	rv := objc.Send[MPSImageCopyToMatrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that copies image data to a matrix.
//
// # Overview
//
// This kernel copies image data to a [MPSMatrix] object. The image data is
// stored in a row of a matrix. The [MPSImageCopyToMatrix.DataLayout]
// specifies the order in which the feature channels in the image get stored
// in the matrix. If the [MPSImage] stores a batch of images, the images are
// copied into multiple rows, one row per image.
//
// The number of elements in a row in the matrix must be greater than the
// image width multiplied its height multiplied by the number of
// [MPSImage.FeatureChannels] in the image.
//
// # Initializers
//
//   - [MPSImageCopyToMatrix.InitWithDeviceDataLayout]
//
// # Instance Properties
//
//   - [MPSImageCopyToMatrix.DataLayout]
//   - [MPSImageCopyToMatrix.DestinationMatrixBatchIndex]
//   - [MPSImageCopyToMatrix.SetDestinationMatrixBatchIndex]
//   - [MPSImageCopyToMatrix.DestinationMatrixOrigin]
//   - [MPSImageCopyToMatrix.SetDestinationMatrixOrigin]
//
// # Instance Methods
//
//   - [MPSImageCopyToMatrix.EncodeToCommandBufferSourceImageDestinationMatrix]
//   - [MPSImageCopyToMatrix.EncodeBatchToCommandBufferSourceImagesDestinationMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix
type MPSImageCopyToMatrix struct {
	MPSKernel
}

// MPSImageCopyToMatrixFromID constructs a [MPSImageCopyToMatrix] from an objc.ID.
//
// A class that copies image data to a matrix.
func MPSImageCopyToMatrixFromID(id objc.ID) MPSImageCopyToMatrix {
	return MPSImageCopyToMatrix{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageCopyToMatrix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageCopyToMatrix] class.
//
// # Initializers
//
//   - [IMPSImageCopyToMatrix.InitWithDeviceDataLayout]
//
// # Instance Properties
//
//   - [IMPSImageCopyToMatrix.DataLayout]
//   - [IMPSImageCopyToMatrix.DestinationMatrixBatchIndex]
//   - [IMPSImageCopyToMatrix.SetDestinationMatrixBatchIndex]
//   - [IMPSImageCopyToMatrix.DestinationMatrixOrigin]
//   - [IMPSImageCopyToMatrix.SetDestinationMatrixOrigin]
//
// # Instance Methods
//
//   - [IMPSImageCopyToMatrix.EncodeToCommandBufferSourceImageDestinationMatrix]
//   - [IMPSImageCopyToMatrix.EncodeBatchToCommandBufferSourceImagesDestinationMatrix]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix
type IMPSImageCopyToMatrix interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSImageCopyToMatrix

	// Topic: Instance Properties

	DataLayout() MPSDataLayout
	DestinationMatrixBatchIndex() uint
	SetDestinationMatrixBatchIndex(value uint)
	DestinationMatrixOrigin() metal.MTLOrigin
	SetDestinationMatrixOrigin(value metal.MTLOrigin)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationMatrix IMPSMatrix)
	EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationMatrix IMPSMatrix)
}

// Init initializes the instance.
func (i MPSImageCopyToMatrix) Init() MPSImageCopyToMatrix {
	rv := objc.Send[MPSImageCopyToMatrix](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageCopyToMatrix) Autorelease() MPSImageCopyToMatrix {
	rv := objc.Send[MPSImageCopyToMatrix](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageCopyToMatrix creates a new MPSImageCopyToMatrix instance.
func NewMPSImageCopyToMatrix() MPSImageCopyToMatrix {
	class := getMPSImageCopyToMatrixClass()
	rv := objc.Send[MPSImageCopyToMatrix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageCopyToMatrixWithCoder(aDecoder foundation.INSCoder) MPSImageCopyToMatrix {
	instance := getMPSImageCopyToMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageCopyToMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/init(coder:device:)
func NewImageCopyToMatrixWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageCopyToMatrix {
	instance := getMPSImageCopyToMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageCopyToMatrixFromID(rv)
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
func NewImageCopyToMatrixWithDevice(device metal.MTLDevice) MPSImageCopyToMatrix {
	instance := getMPSImageCopyToMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageCopyToMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/init(device:dataLayout:)
func NewImageCopyToMatrixWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSImageCopyToMatrix {
	instance := getMPSImageCopyToMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	return MPSImageCopyToMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/init(device:dataLayout:)
func (i MPSImageCopyToMatrix) InitWithDeviceDataLayout(device metal.MTLDevice, dataLayout MPSDataLayout) MPSImageCopyToMatrix {
	rv := objc.Send[MPSImageCopyToMatrix](i.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/encode(commandBuffer:sourceImage:destinationMatrix:)
func (i MPSImageCopyToMatrix) EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationMatrix IMPSMatrix) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationMatrix:"), commandBuffer, sourceImage, destinationMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/encodeBatch(commandBuffer:sourceImages:destinationMatrix:)
func (i MPSImageCopyToMatrix) EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationMatrix IMPSMatrix) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationMatrix:"), commandBuffer, sourceImages, destinationMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/dataLayout
func (i MPSImageCopyToMatrix) DataLayout() MPSDataLayout {
	rv := objc.Send[MPSDataLayout](i.ID, objc.Sel("dataLayout"))
	return MPSDataLayout(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/destinationMatrixBatchIndex
func (i MPSImageCopyToMatrix) DestinationMatrixBatchIndex() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("destinationMatrixBatchIndex"))
	return rv
}
func (i MPSImageCopyToMatrix) SetDestinationMatrixBatchIndex(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setDestinationMatrixBatchIndex:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix/destinationMatrixOrigin
func (i MPSImageCopyToMatrix) DestinationMatrixOrigin() metal.MTLOrigin {
	rv := objc.Send[metal.MTLOrigin](i.ID, objc.Sel("destinationMatrixOrigin"))
	return metal.MTLOrigin(rv)
}
func (i MPSImageCopyToMatrix) SetDestinationMatrixOrigin(value metal.MTLOrigin) {
	objc.Send[struct{}](i.ID, objc.Sel("setDestinationMatrixOrigin:"), value)
}
