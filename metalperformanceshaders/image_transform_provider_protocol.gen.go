// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A general interface for objects that provide image resampling.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTransformProvider
type MPSImageTransformProvider interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSSecureCoding

	// TransformForSourceImageHandle protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTransformProvider/transform(forSourceImage:handle:)
	TransformForSourceImageHandle(image IMPSImage, handle MPSHandle) MPSScaleTransform
}

// MPSImageTransformProviderObject wraps an existing Objective-C object that conforms to the MPSImageTransformProvider protocol.
type MPSImageTransformProviderObject struct {
	foundation.NSCodingObject
}

func (o MPSImageTransformProviderObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSImageTransformProviderObjectFromID constructs a [MPSImageTransformProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSImageTransformProviderObjectFromID(id objc.ID) MPSImageTransformProviderObject {
	return MPSImageTransformProviderObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTransformProvider/transform(forSourceImage:handle:)
func (o MPSImageTransformProviderObject) TransformForSourceImageHandle(image IMPSImage, handle MPSHandle) MPSScaleTransform {
	rv := objc.Send[MPSScaleTransform](o.ID, objc.Sel("transformForSourceImage:handle:"), image, handle)
	return rv
}
