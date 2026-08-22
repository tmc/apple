// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Returns the minimum dimensions for a `sourceFrame` for the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/minimumDimensions-9n9r0
type minimumDimensions interface {
	objectivec.IObject
}

// minimumDimensionsObject wraps an existing Objective-C object that conforms to the minimumDimensions protocol.
type minimumDimensionsObject struct {
	objectivec.Object
}

func (o minimumDimensionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// minimumDimensionsObjectFromID constructs a [minimumDimensionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func minimumDimensionsObjectFromID(id objc.ID) minimumDimensionsObject {
	return minimumDimensionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}
