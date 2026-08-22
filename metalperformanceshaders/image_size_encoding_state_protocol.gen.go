// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that contain information about an image size elsewhere in the graph.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState
type MPSImageSizeEncodingState interface {
	objectivec.IObject

	// sourceHeight protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceHeight
	SourceHeight() uint

	// sourceWidth protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceWidth
	SourceWidth() uint
}

// MPSImageSizeEncodingStateObject wraps an existing Objective-C object that conforms to the MPSImageSizeEncodingState protocol.
type MPSImageSizeEncodingStateObject struct {
	objectivec.Object
}

func (o MPSImageSizeEncodingStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSImageSizeEncodingStateObjectFromID constructs a [MPSImageSizeEncodingStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSImageSizeEncodingStateObjectFromID(id objc.ID) MPSImageSizeEncodingStateObject {
	return MPSImageSizeEncodingStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceHeight
func (o MPSImageSizeEncodingStateObject) SourceHeight() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("sourceHeight"))
	return uint(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceWidth
func (o MPSImageSizeEncodingStateObject) SourceWidth() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("sourceWidth"))
	return uint(rv)
}
