// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Returns the maximum dimensions for a `sourceFrame` for the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/maximumDimensions-1x3l0
type maximumDimensions interface {
	objectivec.IObject
}

// maximumDimensionsObject wraps an existing Objective-C object that conforms to the maximumDimensions protocol.
type maximumDimensionsObject struct {
	objectivec.Object
}

func (o maximumDimensionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// maximumDimensionsObjectFromID constructs a [maximumDimensionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func maximumDimensionsObjectFromID(id objc.ID) maximumDimensionsObject {
	return maximumDimensionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}
