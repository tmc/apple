// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines methods that a batch normalization state uses to initialize scale factors, bias terms, and batch statistics.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource
type MPSCNNBatchNormalizationDataSource interface {
	objectivec.IObject
	foundation.NSCopying

	// Beta protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/beta()
	Beta() unsafe.Pointer

	// Gamma protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/gamma()
	Gamma() unsafe.Pointer

	// Label protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/label()
	Label() string

	// Load protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/load()
	Load() bool

	// Mean protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/mean()
	Mean() unsafe.Pointer

	// NumberOfFeatureChannels protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/numberOfFeatureChannels()
	NumberOfFeatureChannels() uint

	// Purge protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/purge()
	Purge()

	// Variance protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/variance()
	Variance() unsafe.Pointer
}

// MPSCNNBatchNormalizationDataSourceObject wraps an existing Objective-C object that conforms to the MPSCNNBatchNormalizationDataSource protocol.
type MPSCNNBatchNormalizationDataSourceObject struct {
	foundation.NSCopyingObject
}

func (o MPSCNNBatchNormalizationDataSourceObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MPSCNNBatchNormalizationDataSourceObjectFromID constructs a [MPSCNNBatchNormalizationDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSCNNBatchNormalizationDataSourceObjectFromID(id objc.ID) MPSCNNBatchNormalizationDataSourceObject {
	return MPSCNNBatchNormalizationDataSourceObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/beta()
func (o MPSCNNBatchNormalizationDataSourceObject) Beta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("beta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/gamma()
func (o MPSCNNBatchNormalizationDataSourceObject) Gamma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("gamma"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/label()
func (o MPSCNNBatchNormalizationDataSourceObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/load()
func (o MPSCNNBatchNormalizationDataSourceObject) Load() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("load"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/mean()
func (o MPSCNNBatchNormalizationDataSourceObject) Mean() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("mean"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/numberOfFeatureChannels()
func (o MPSCNNBatchNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("numberOfFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/purge()
func (o MPSCNNBatchNormalizationDataSourceObject) Purge() {
	objc.Send[struct{}](o.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/variance()
func (o MPSCNNBatchNormalizationDataSourceObject) Variance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("variance"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/copy(with:device:)
func (o MPSCNNBatchNormalizationDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/encode(with:)
func (o MPSCNNBatchNormalizationDataSourceObject) EncodeWithCoder(aCoder foundation.INSCoder) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), aCoder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/epsilon()
func (o MPSCNNBatchNormalizationDataSourceObject) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/updateGammaAndBeta(with:)
func (o MPSCNNBatchNormalizationDataSourceObject) UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState IMPSCNNBatchNormalizationState) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateGammaAndBetaWithBatchNormalizationState:"), batchNormalizationState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/updateGammaAndBeta(with:batchNormalizationState:)
func (o MPSCNNBatchNormalizationDataSourceObject) UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState) IMPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateGammaAndBetaWithCommandBuffer:batchNormalizationState:"), commandBuffer, batchNormalizationState)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/updateMeanAndVariance(with:)
func (o MPSCNNBatchNormalizationDataSourceObject) UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState IMPSCNNBatchNormalizationState) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateMeanAndVarianceWithBatchNormalizationState:"), batchNormalizationState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource/updateMeanAndVariance(with:batchNormalizationState:)
func (o MPSCNNBatchNormalizationDataSourceObject) UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState) IMPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateMeanAndVarianceWithCommandBuffer:batchNormalizationState:"), commandBuffer, batchNormalizationState)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}
