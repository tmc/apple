// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines methods that an instance normalization uses to initialize scale factors and bias terms.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource
type MPSCNNInstanceNormalizationDataSource interface {
	objectivec.IObject
	foundation.NSCopying

	// Beta protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/beta()
	Beta() unsafe.Pointer

	// Gamma protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/gamma()
	Gamma() unsafe.Pointer

	// Label protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/label()
	Label() string

	// numberOfFeatureChannels protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/numberOfFeatureChannels
	NumberOfFeatureChannels() uint
}

// MPSCNNInstanceNormalizationDataSourceObject wraps an existing Objective-C object that conforms to the MPSCNNInstanceNormalizationDataSource protocol.
type MPSCNNInstanceNormalizationDataSourceObject struct {
	foundation.NSCopyingObject
}

func (o MPSCNNInstanceNormalizationDataSourceObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MPSCNNInstanceNormalizationDataSourceObjectFromID constructs a [MPSCNNInstanceNormalizationDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSCNNInstanceNormalizationDataSourceObjectFromID(id objc.ID) MPSCNNInstanceNormalizationDataSourceObject {
	return MPSCNNInstanceNormalizationDataSourceObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/beta()
func (o MPSCNNInstanceNormalizationDataSourceObject) Beta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("beta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/gamma()
func (o MPSCNNInstanceNormalizationDataSourceObject) Gamma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("gamma"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/label()
func (o MPSCNNInstanceNormalizationDataSourceObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/copy(with:device:)
func (o MPSCNNInstanceNormalizationDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/encode(with:)
func (o MPSCNNInstanceNormalizationDataSourceObject) EncodeWithCoder(aCoder foundation.INSCoder) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), aCoder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/epsilon()
func (o MPSCNNInstanceNormalizationDataSourceObject) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/load()
func (o MPSCNNInstanceNormalizationDataSourceObject) Load() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("load"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/purge()
func (o MPSCNNInstanceNormalizationDataSourceObject) Purge() {
	objc.Send[struct{}](o.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/updateGammaAndBeta(with:instanceNormalizationStateBatch:)
func (o MPSCNNInstanceNormalizationDataSourceObject) UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer metal.MTLCommandBuffer, instanceNormalizationStateBatch MPSCNNInstanceNormalizationGradientStateBatch) IMPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateGammaAndBetaWithCommandBuffer:instanceNormalizationStateBatch:"), commandBuffer, instanceNormalizationStateBatch)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/updateGammaAndBeta(withInstanceNormalizationStateBatch:)
func (o MPSCNNInstanceNormalizationDataSourceObject) UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch MPSCNNInstanceNormalizationGradientStateBatch) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateGammaAndBetaWithInstanceNormalizationStateBatch:"), instanceNormalizationStateBatch)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource/numberOfFeatureChannels
func (o MPSCNNInstanceNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("numberOfFeatureChannels"))
	return uint(rv)
}
