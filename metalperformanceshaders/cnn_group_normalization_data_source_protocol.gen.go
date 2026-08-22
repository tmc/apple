// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSCNNGroupNormalizationDataSource protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource
type MPSCNNGroupNormalizationDataSource interface {
	objectivec.IObject
	foundation.NSCopying

	// Beta protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/beta()
	Beta() unsafe.Pointer

	// Gamma protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/gamma()
	Gamma() unsafe.Pointer

	// Label protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/label()
	Label() string

	// numberOfFeatureChannels protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/numberOfFeatureChannels
	NumberOfFeatureChannels() uint

	// numberOfGroups protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/numberOfGroups
	NumberOfGroups() uint
	SetNumberOfGroups(value uint)
}

// MPSCNNGroupNormalizationDataSourceObject wraps an existing Objective-C object that conforms to the MPSCNNGroupNormalizationDataSource protocol.
type MPSCNNGroupNormalizationDataSourceObject struct {
	foundation.NSCopyingObject
}

func (o MPSCNNGroupNormalizationDataSourceObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MPSCNNGroupNormalizationDataSourceObjectFromID constructs a [MPSCNNGroupNormalizationDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSCNNGroupNormalizationDataSourceObjectFromID(id objc.ID) MPSCNNGroupNormalizationDataSourceObject {
	return MPSCNNGroupNormalizationDataSourceObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/beta()
func (o MPSCNNGroupNormalizationDataSourceObject) Beta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("beta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/gamma()
func (o MPSCNNGroupNormalizationDataSourceObject) Gamma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("gamma"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/label()
func (o MPSCNNGroupNormalizationDataSourceObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/copy(with:device:)
func (o MPSCNNGroupNormalizationDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/encode(with:)
func (o MPSCNNGroupNormalizationDataSourceObject) EncodeWithCoder(aCoder foundation.INSCoder) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), aCoder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/epsilon()
func (o MPSCNNGroupNormalizationDataSourceObject) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/updateGammaAndBeta(with:groupNormalizationStateBatch:)
func (o MPSCNNGroupNormalizationDataSourceObject) UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer metal.MTLCommandBuffer, groupNormalizationStateBatch MPSCNNGroupNormalizationGradientStateBatch) IMPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateGammaAndBetaWithCommandBuffer:groupNormalizationStateBatch:"), commandBuffer, groupNormalizationStateBatch)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/updateGammaAndBeta(withGroupNormalizationStateBatch:)
func (o MPSCNNGroupNormalizationDataSourceObject) UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch MPSCNNGroupNormalizationGradientStateBatch) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateGammaAndBetaWithGroupNormalizationStateBatch:"), groupNormalizationStateBatch)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/numberOfFeatureChannels
func (o MPSCNNGroupNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("numberOfFeatureChannels"))
	return uint(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource/numberOfGroups
func (o MPSCNNGroupNormalizationDataSourceObject) NumberOfGroups() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("numberOfGroups"))
	return uint(rv)
}

func (o MPSCNNGroupNormalizationDataSourceObject) SetNumberOfGroups(value uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setNumberOfGroups:"), value)
}
