// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A Boolean value that indicates whether the current configuration supports the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/isSupported
type isSupported interface {
	objectivec.IObject
}

// isSupportedObject wraps an existing Objective-C object that conforms to the isSupported protocol.
type isSupportedObject struct {
	objectivec.Object
}

func (o isSupportedObject) BaseObject() objectivec.Object {
	return o.Object
}

// isSupportedObjectFromID constructs a [isSupportedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func isSupportedObjectFromID(id objc.ID) isSupportedObject {
	return isSupportedObject{
		Object: objectivec.ObjectFromID(id),
	}
}
