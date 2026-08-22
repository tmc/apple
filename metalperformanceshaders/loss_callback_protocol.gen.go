// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSNNLossCallback protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossCallback
type MPSNNLossCallback interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSCopying
	foundation.NSSecureCoding

	// ScalarWeightForSourceImageDestinationImage protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossCallback/scalarWeight(forSourceImage:destinationImage:)
	ScalarWeightForSourceImageDestinationImage(sourceImage IMPSImage, destinationImage IMPSImage) float32
}

// MPSNNLossCallbackObject wraps an existing Objective-C object that conforms to the MPSNNLossCallback protocol.
type MPSNNLossCallbackObject struct {
	foundation.NSCodingObject
}

func (o MPSNNLossCallbackObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSNNLossCallbackObjectFromID constructs a [MPSNNLossCallbackObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSNNLossCallbackObjectFromID(id objc.ID) MPSNNLossCallbackObject {
	return MPSNNLossCallbackObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossCallback/scalarWeight(forSourceImage:destinationImage:)
func (o MPSNNLossCallbackObject) ScalarWeightForSourceImageDestinationImage(sourceImage IMPSImage, destinationImage IMPSImage) float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scalarWeightForSourceImage:destinationImage:"), sourceImage, destinationImage)
	return rv
}
