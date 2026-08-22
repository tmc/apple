// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNYOLOLoss] class.
var (
	_MPSCNNYOLOLossClass     MPSCNNYOLOLossClass
	_MPSCNNYOLOLossClassOnce sync.Once
)

func getMPSCNNYOLOLossClass() MPSCNNYOLOLossClass {
	_MPSCNNYOLOLossClassOnce.Do(func() {
		_MPSCNNYOLOLossClass = MPSCNNYOLOLossClass{class: objc.GetClass("MPSCNNYOLOLoss")}
	})
	return _MPSCNNYOLOLossClass
}

// GetMPSCNNYOLOLossClass returns the class object for MPSCNNYOLOLoss.
func GetMPSCNNYOLOLossClass() MPSCNNYOLOLossClass {
	return getMPSCNNYOLOLossClass()
}

type MPSCNNYOLOLossClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNYOLOLossClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNYOLOLossClass) Alloc() MPSCNNYOLOLoss {
	rv := objc.Send[MPSCNNYOLOLoss](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that computes the YOLO loss and loss gradient between specified
// predictions and labels.
//
// # Initializers
//
//   - [MPSCNNYOLOLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNYOLOLoss.AnchorBoxes]
//   - [MPSCNNYOLOLoss.LossClasses]
//   - [MPSCNNYOLOLoss.LossConfidence]
//   - [MPSCNNYOLOLoss.LossWH]
//   - [MPSCNNYOLOLoss.LossXY]
//   - [MPSCNNYOLOLoss.MaxIOUForObjectAbsence]
//   - [MPSCNNYOLOLoss.MinIOUForObjectPresence]
//   - [MPSCNNYOLOLoss.NumberOfAnchorBoxes]
//   - [MPSCNNYOLOLoss.ReduceAcrossBatch]
//   - [MPSCNNYOLOLoss.ReductionType]
//   - [MPSCNNYOLOLoss.ScaleClass]
//   - [MPSCNNYOLOLoss.ScaleNoObject]
//   - [MPSCNNYOLOLoss.ScaleObject]
//   - [MPSCNNYOLOLoss.ScaleWH]
//   - [MPSCNNYOLOLoss.ScaleXY]
//
// # Instance Methods
//
//   - [MPSCNNYOLOLoss.EncodeToCommandBufferSourceImageLabels]
//   - [MPSCNNYOLOLoss.EncodeToCommandBufferSourceImageLabelsDestinationImage]
//   - [MPSCNNYOLOLoss.EncodeBatchToCommandBufferSourceImagesLabels]
//   - [MPSCNNYOLOLoss.EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss
type MPSCNNYOLOLoss struct {
	MPSCNNKernel
}

// MPSCNNYOLOLossFromID constructs a [MPSCNNYOLOLoss] from an objc.ID.
//
// A kernel that computes the YOLO loss and loss gradient between specified
// predictions and labels.
func MPSCNNYOLOLossFromID(id objc.ID) MPSCNNYOLOLoss {
	return MPSCNNYOLOLoss{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNYOLOLoss adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNYOLOLoss] class.
//
// # Initializers
//
//   - [IMPSCNNYOLOLoss.InitWithDeviceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNYOLOLoss.AnchorBoxes]
//   - [IMPSCNNYOLOLoss.LossClasses]
//   - [IMPSCNNYOLOLoss.LossConfidence]
//   - [IMPSCNNYOLOLoss.LossWH]
//   - [IMPSCNNYOLOLoss.LossXY]
//   - [IMPSCNNYOLOLoss.MaxIOUForObjectAbsence]
//   - [IMPSCNNYOLOLoss.MinIOUForObjectPresence]
//   - [IMPSCNNYOLOLoss.NumberOfAnchorBoxes]
//   - [IMPSCNNYOLOLoss.ReduceAcrossBatch]
//   - [IMPSCNNYOLOLoss.ReductionType]
//   - [IMPSCNNYOLOLoss.ScaleClass]
//   - [IMPSCNNYOLOLoss.ScaleNoObject]
//   - [IMPSCNNYOLOLoss.ScaleObject]
//   - [IMPSCNNYOLOLoss.ScaleWH]
//   - [IMPSCNNYOLOLoss.ScaleXY]
//
// # Instance Methods
//
//   - [IMPSCNNYOLOLoss.EncodeToCommandBufferSourceImageLabels]
//   - [IMPSCNNYOLOLoss.EncodeToCommandBufferSourceImageLabelsDestinationImage]
//   - [IMPSCNNYOLOLoss.EncodeBatchToCommandBufferSourceImagesLabels]
//   - [IMPSCNNYOLOLoss.EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss
type IMPSCNNYOLOLoss interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLoss

	// Topic: Instance Properties

	AnchorBoxes() foundation.NSData
	LossClasses() IMPSCNNLoss
	LossConfidence() IMPSCNNLoss
	LossWH() IMPSCNNLoss
	LossXY() IMPSCNNLoss
	MaxIOUForObjectAbsence() float32
	MinIOUForObjectPresence() float32
	NumberOfAnchorBoxes() uint
	ReduceAcrossBatch() bool
	ReductionType() MPSCNNReductionType
	ScaleClass() float32
	ScaleNoObject() float32
	ScaleObject() float32
	ScaleWH() float32
	ScaleXY() float32

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageLabels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels) IMPSImage
	EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels, destinationImage IMPSImage)
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch, destinationImage MPSImageBatch)
}

// Init initializes the instance.
func (c MPSCNNYOLOLoss) Init() MPSCNNYOLOLoss {
	rv := objc.Send[MPSCNNYOLOLoss](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNYOLOLoss) Autorelease() MPSCNNYOLOLoss {
	rv := objc.Send[MPSCNNYOLOLoss](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNYOLOLoss creates a new MPSCNNYOLOLoss instance.
func NewMPSCNNYOLOLoss() MPSCNNYOLOLoss {
	class := getMPSCNNYOLOLossClass()
	rv := objc.Send[MPSCNNYOLOLoss](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNYOLOLossWithCoder(aDecoder foundation.INSCoder) MPSCNNYOLOLoss {
	instance := getMPSCNNYOLOLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNYOLOLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/init(coder:device:)
func NewCNNYOLOLossWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNYOLOLoss {
	instance := getMPSCNNYOLOLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNYOLOLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNYOLOLossWithDevice(device metal.MTLDevice) MPSCNNYOLOLoss {
	instance := getMPSCNNYOLOLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNYOLOLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/init(device:lossDescriptor:)
func NewCNNYOLOLossWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLoss {
	instance := getMPSCNNYOLOLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return MPSCNNYOLOLossFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/init(device:lossDescriptor:)
func (c MPSCNNYOLOLoss) InitWithDeviceLossDescriptor(device metal.MTLDevice, lossDescriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLoss {
	rv := objc.Send[MPSCNNYOLOLoss](c.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/encode(commandBuffer:sourceImage:labels:)
func (c MPSCNNYOLOLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), commandBuffer, sourceImage, labels)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/encode(commandBuffer:sourceImage:labels:destinationImage:)
func (c MPSCNNYOLOLoss) EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, labels IMPSCNNLossLabels, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:destinationImage:"), commandBuffer, sourceImage, labels, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/encode(commandBuffer:sourceImages:labels:)
func (c MPSCNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), commandBuffer, sourceImage, labels)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/encode(commandBuffer:sourceImages:labels:destinationImages:)
func (c MPSCNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, labels MPSCNNLossLabelsBatch, destinationImage MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:destinationImages:"), commandBuffer, sourceImage, labels, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/anchorBoxes
func (c MPSCNNYOLOLoss) AnchorBoxes() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("anchorBoxes"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/lossClasses
func (c MPSCNNYOLOLoss) LossClasses() IMPSCNNLoss {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lossClasses"))
	return MPSCNNLossFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/lossConfidence
func (c MPSCNNYOLOLoss) LossConfidence() IMPSCNNLoss {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lossConfidence"))
	return MPSCNNLossFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/lossWH
func (c MPSCNNYOLOLoss) LossWH() IMPSCNNLoss {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lossWH"))
	return MPSCNNLossFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/lossXY
func (c MPSCNNYOLOLoss) LossXY() IMPSCNNLoss {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lossXY"))
	return MPSCNNLossFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/maxIOUForObjectAbsence
func (c MPSCNNYOLOLoss) MaxIOUForObjectAbsence() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/minIOUForObjectPresence
func (c MPSCNNYOLOLoss) MinIOUForObjectPresence() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("minIOUForObjectPresence"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/numberOfAnchorBoxes
func (c MPSCNNYOLOLoss) NumberOfAnchorBoxes() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfAnchorBoxes"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/reduceAcrossBatch
func (c MPSCNNYOLOLoss) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/reductionType
func (c MPSCNNYOLOLoss) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](c.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/scaleClass
func (c MPSCNNYOLOLoss) ScaleClass() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleClass"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/scaleNoObject
func (c MPSCNNYOLOLoss) ScaleNoObject() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleNoObject"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/scaleObject
func (c MPSCNNYOLOLoss) ScaleObject() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleObject"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/scaleWH
func (c MPSCNNYOLOLoss) ScaleWH() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleWH"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss/scaleXY
func (c MPSCNNYOLOLoss) ScaleXY() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleXY"))
	return rv
}
