// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol that provides a description of how kernels should pad images.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding
type MPSNNPadding interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSSecureCoding

	// PaddingMethod protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/paddingMethod()
	PaddingMethod() MPSNNPaddingMethod
}

// MPSNNPaddingObject wraps an existing Objective-C object that conforms to the MPSNNPadding protocol.
type MPSNNPaddingObject struct {
	foundation.NSCodingObject
}

func (o MPSNNPaddingObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSNNPaddingObjectFromID constructs a [MPSNNPaddingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSNNPaddingObjectFromID(id objc.ID) MPSNNPaddingObject {
	return MPSNNPaddingObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/paddingMethod()
func (o MPSNNPaddingObject) PaddingMethod() MPSNNPaddingMethod {
	rv := objc.Send[MPSNNPaddingMethod](o.ID, objc.Sel("paddingMethod"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/destinationImageDescriptor(forSourceImages:sourceStates:for:suggestedDescriptor:)
func (o MPSNNPaddingObject) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []MPSImage, sourceStates []MPSState, kernel IMPSKernel, inDescriptor IMPSImageDescriptor) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), kernel, inDescriptor)
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/label()
func (o MPSNNPaddingObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadding/inverse()
func (o MPSNNPaddingObject) Inverse() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inverse"))
	return objectivec.Object{ID: rv}
}
