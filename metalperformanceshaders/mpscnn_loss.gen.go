// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLoss] class.
var (
	_MPSCNNLossClass     MPSCNNLossClass
	_MPSCNNLossClassOnce sync.Once
)

func getMPSCNNLossClass() MPSCNNLossClass {
	_MPSCNNLossClassOnce.Do(func() {
		_MPSCNNLossClass = MPSCNNLossClass{class: objc.GetClass("MPSCNNLoss")}
	})
	return _MPSCNNLossClass
}

// GetMPSCNNLossClass returns the class object for MPSCNNLoss.
func GetMPSCNNLossClass() MPSCNNLossClass {
	return getMPSCNNLossClass()
}

type MPSCNNLossClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLossClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLossClass) Alloc() MPSCNNLoss {
	rv := objc.Send[MPSCNNLoss](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that computes the loss and loss gradient between specified
// predictions and labels.
//
// # Initializers
//
//   - [MPSCNNLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNLoss.Delta]
//   - [MPSCNNLoss.Epsilon]
//   - [MPSCNNLoss.LabelSmoothing]
//   - [MPSCNNLoss.LossType]
//   - [MPSCNNLoss.NumberOfClasses]
//   - [MPSCNNLoss.ReduceAcrossBatch]
//   - [MPSCNNLoss.ReductionType]
//   - [MPSCNNLoss.Weight]
//
// # Instance Methods
//
//   - [MPSCNNLoss.EncodeToCommandBufferSourceImageLabels]
//   - [MPSCNNLoss.EncodeToCommandBufferSourceImageLabelsDestinationImage]
//   - [MPSCNNLoss.EncodeBatchToCommandBufferSourceImagesLabels]
//   - [MPSCNNLoss.EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss
type MPSCNNLoss struct {
	MPSCNNKernel
}

// MPSCNNLossFromID constructs a [MPSCNNLoss] from an objc.ID.
//
// A kernel that computes the loss and loss gradient between specified
// predictions and labels.
func MPSCNNLossFromID(id objc.ID) MPSCNNLoss {
	return MPSCNNLoss{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNLoss adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLoss] class.
//
// # Initializers
//
//   - [IMPSCNNLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNLoss.Delta]
//   - [IMPSCNNLoss.Epsilon]
//   - [IMPSCNNLoss.LabelSmoothing]
//   - [IMPSCNNLoss.LossType]
//   - [IMPSCNNLoss.NumberOfClasses]
//   - [IMPSCNNLoss.ReduceAcrossBatch]
//   - [IMPSCNNLoss.ReductionType]
//   - [IMPSCNNLoss.Weight]
//
// # Instance Methods
//
//   - [IMPSCNNLoss.EncodeToCommandBufferSourceImageLabels]
//   - [IMPSCNNLoss.EncodeToCommandBufferSourceImageLabelsDestinationImage]
//   - [IMPSCNNLoss.EncodeBatchToCommandBufferSourceImagesLabels]
//   - [IMPSCNNLoss.EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss
type IMPSCNNLoss interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSCNNLoss

	// Topic: Instance Properties

	Delta() float32
	Epsilon() float32
	LabelSmoothing() float32
	LossType() MPSCNNLossType
	NumberOfClasses() uint
	ReduceAcrossBatch() bool
	ReductionType() MPSCNNReductionType
	Weight() float32

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageLabels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels) IMPSImage
	EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels, destinationImage IMPSImage)
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch, destinationImage MPSImageBatch)
}

// Init initializes the instance.
func (c MPSCNNLoss) Init() MPSCNNLoss {
	rv := objc.Send[MPSCNNLoss](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLoss) Autorelease() MPSCNNLoss {
	rv := objc.Send[MPSCNNLoss](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLoss creates a new MPSCNNLoss instance.
func NewMPSCNNLoss() MPSCNNLoss {
	class := getMPSCNNLossClass()
	rv := objc.Send[MPSCNNLoss](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNLossWithCoder(aDecoder foundation.INSCoder) MPSCNNLoss {
	instance := getMPSCNNLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/init(coder:device:)
func NewCNNLossWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNLoss {
	instance := getMPSCNNLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNLossWithDevice(device metal.MTLDevice) MPSCNNLoss {
	instance := getMPSCNNLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/init(device:lossDescriptor:)
func NewCNNLossWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSCNNLoss {
	instance := getMPSCNNLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return MPSCNNLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/init(device:lossDescriptor:)
func (c MPSCNNLoss) InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNLossDescriptor) MPSCNNLoss {
	rv := objc.Send[MPSCNNLoss](c.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/encode(commandBuffer:sourceImage:labels:)
func (c MPSCNNLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), commandBuffer, sourceImage, labels)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/encode(commandBuffer:sourceImage:labels:destinationImage:)
func (c MPSCNNLoss) EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:destinationImage:"), commandBuffer, sourceImage, labels, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/encode(commandBuffer:sourceImages:labels:)
func (c MPSCNNLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), commandBuffer, sourceImage, labels)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/encode(commandBuffer:sourceImages:labels:destinationImages:)
func (c MPSCNNLoss) EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch, destinationImage MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:destinationImages:"), commandBuffer, sourceImage, labels, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/delta
func (c MPSCNNLoss) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/epsilon
func (c MPSCNNLoss) Epsilon() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/labelSmoothing
func (c MPSCNNLoss) LabelSmoothing() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("labelSmoothing"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/lossType
func (c MPSCNNLoss) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](c.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/numberOfClasses
func (c MPSCNNLoss) NumberOfClasses() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfClasses"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/reduceAcrossBatch
func (c MPSCNNLoss) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/reductionType
func (c MPSCNNLoss) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](c.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss/weight
func (c MPSCNNLoss) Weight() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("weight"))
	return rv
}
