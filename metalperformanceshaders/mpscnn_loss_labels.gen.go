// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNLossLabels] class.
var (
	_MPSCNNLossLabelsClass     MPSCNNLossLabelsClass
	_MPSCNNLossLabelsClassOnce sync.Once
)

func getMPSCNNLossLabelsClass() MPSCNNLossLabelsClass {
	_MPSCNNLossLabelsClassOnce.Do(func() {
		_MPSCNNLossLabelsClass = MPSCNNLossLabelsClass{class: objc.GetClass("MPSCNNLossLabels")}
	})
	return _MPSCNNLossLabelsClass
}

// GetMPSCNNLossLabelsClass returns the class object for MPSCNNLossLabels.
func GetMPSCNNLossLabelsClass() MPSCNNLossLabelsClass {
	return getMPSCNNLossLabelsClass()
}

type MPSCNNLossLabelsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLossLabelsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLossLabelsClass) Alloc() MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that stores the per-element weight buffer used by loss and gradient
// loss kernels.
//
// # Initializers
//
//   - [MPSCNNLossLabels.InitWithDeviceLabelsDescriptor]
//   - [MPSCNNLossLabels.InitWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor]
//   - [MPSCNNLossLabels.InitWithDeviceLossImageSizeLabelsImageWeightsImage]
//
// # Instance Methods
//
//   - [MPSCNNLossLabels.LabelsImage]
//   - [MPSCNNLossLabels.LossImage]
//   - [MPSCNNLossLabels.WeightsImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels
type MPSCNNLossLabels struct {
	MPSState
}

// MPSCNNLossLabelsFromID constructs a [MPSCNNLossLabels] from an objc.ID.
//
// A class that stores the per-element weight buffer used by loss and gradient
// loss kernels.
func MPSCNNLossLabelsFromID(id objc.ID) MPSCNNLossLabels {
	return MPSCNNLossLabels{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSCNNLossLabels adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLossLabels] class.
//
// # Initializers
//
//   - [IMPSCNNLossLabels.InitWithDeviceLabelsDescriptor]
//   - [IMPSCNNLossLabels.InitWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor]
//   - [IMPSCNNLossLabels.InitWithDeviceLossImageSizeLabelsImageWeightsImage]
//
// # Instance Methods
//
//   - [IMPSCNNLossLabels.LabelsImage]
//   - [IMPSCNNLossLabels.LossImage]
//   - [IMPSCNNLossLabels.WeightsImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels
type IMPSCNNLossLabels interface {
	IMPSState

	// Topic: Initializers

	InitWithDeviceLabelsDescriptor(device metal.MTLDevice, labelsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels
	InitWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsDescriptor IMPSCNNLossDataDescriptor, weightsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels
	InitWithDeviceLossImageSizeLabelsImageWeightsImage(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsImage IMPSImage, weightsImage IMPSImage) MPSCNNLossLabels

	// Topic: Instance Methods

	LabelsImage() IMPSImage
	LossImage() IMPSImage
	WeightsImage() IMPSImage
}

// Init initializes the instance.
func (c MPSCNNLossLabels) Init() MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLossLabels) Autorelease() MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLossLabels creates a new MPSCNNLossLabels instance.
func NewMPSCNNLossLabels() MPSCNNLossLabels {
	class := getMPSCNNLossLabelsClass()
	rv := objc.Send[MPSCNNLossLabels](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNLossLabelsWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:labelsDescriptor:)
func NewCNNLossLabelsWithDeviceLabelsDescriptor(device metal.MTLDevice, labelsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:labelsDescriptor:"), device, labelsDescriptor)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:lossImageSize:labelsDescriptor:weightsDescriptor:)
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsDescriptor IMPSCNNLossDataDescriptor, weightsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsDescriptor:weightsDescriptor:"), device, lossImageSize, labelsDescriptor, weightsDescriptor)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:lossImageSize:labelsImage:weightsImage:)
func NewCNNLossLabelsWithDeviceLossImageSizeLabelsImageWeightsImage(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsImage IMPSImage, weightsImage IMPSImage) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossImageSize:labelsImage:weightsImage:"), device, lossImageSize, labelsImage, weightsImage)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNLossLabelsWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNLossLabelsWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNLossLabelsWithResource(resource metal.MTLResource) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNLossLabelsWithResources(resources []objectivec.IObject) MPSCNNLossLabels {
	instance := getMPSCNNLossLabelsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNLossLabelsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:labelsDescriptor:)
func (c MPSCNNLossLabels) InitWithDeviceLabelsDescriptor(device metal.MTLDevice, labelsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](c.ID, objc.Sel("initWithDevice:labelsDescriptor:"), device, labelsDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:lossImageSize:labelsDescriptor:weightsDescriptor:)
func (c MPSCNNLossLabels) InitWithDeviceLossImageSizeLabelsDescriptorWeightsDescriptor(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsDescriptor IMPSCNNLossDataDescriptor, weightsDescriptor IMPSCNNLossDataDescriptor) MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](c.ID, objc.Sel("initWithDevice:lossImageSize:labelsDescriptor:weightsDescriptor:"), device, lossImageSize, labelsDescriptor, weightsDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/init(device:lossImageSize:labelsImage:weightsImage:)
func (c MPSCNNLossLabels) InitWithDeviceLossImageSizeLabelsImageWeightsImage(device metal.MTLDevice, lossImageSize metal.MTLSize, labelsImage IMPSImage, weightsImage IMPSImage) MPSCNNLossLabels {
	rv := objc.Send[MPSCNNLossLabels](c.ID, objc.Sel("initWithDevice:lossImageSize:labelsImage:weightsImage:"), device, lossImageSize, labelsImage, weightsImage)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/labelsImage()
func (c MPSCNNLossLabels) LabelsImage() IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("labelsImage"))
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/lossImage()
func (c MPSCNNLossLabels) LossImage() IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lossImage"))
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossLabels/weightsImage()
func (c MPSCNNLossLabels) WeightsImage() IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("weightsImage"))
	return MPSImageFromID(rv)
}
