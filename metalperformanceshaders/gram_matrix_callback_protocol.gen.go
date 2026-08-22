// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSNNGramMatrixCallback protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCallback
type MPSNNGramMatrixCallback interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSCopying
	foundation.NSSecureCoding

	// AlphaForSourceImageDestinationImage protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCallback/alpha(forSourceImage:destinationImage:)
	AlphaForSourceImageDestinationImage(sourceImage IMPSImage, destinationImage IMPSImage) float32
}

// MPSNNGramMatrixCallbackObject wraps an existing Objective-C object that conforms to the MPSNNGramMatrixCallback protocol.
type MPSNNGramMatrixCallbackObject struct {
	foundation.NSCodingObject
}

func (o MPSNNGramMatrixCallbackObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSNNGramMatrixCallbackObjectFromID constructs a [MPSNNGramMatrixCallbackObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSNNGramMatrixCallbackObjectFromID(id objc.ID) MPSNNGramMatrixCallbackObject {
	return MPSNNGramMatrixCallbackObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCallback/alpha(forSourceImage:destinationImage:)
func (o MPSNNGramMatrixCallbackObject) AlphaForSourceImageDestinationImage(sourceImage IMPSImage, destinationImage IMPSImage) float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("alphaForSourceImage:destinationImage:"), sourceImage, destinationImage)
	return rv
}
