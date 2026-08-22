// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that describes the configuration of a processor to use during a video processing session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration
type VTFrameProcessorConfiguration interface {
	objectivec.IObject

	// A dictionary of pixel buffer attributes that define what the source and reference frames passed to the processor must conform to.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/sourcePixelBufferAttributes
	SourcePixelBufferAttributes() foundation.INSDictionary

	// A dictionary of pixel buffer attributes that define which output frames passed to the processor must conform to.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/destinationPixelBufferAttributes
	DestinationPixelBufferAttributes() foundation.INSDictionary

	// Returns the number of “next” frames that this processor requires for processing.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
	NextFrameCount() int

	// Returns the number of “previous” frames that this processor requires for processing.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
	PreviousFrameCount() int

	// A list of supported pixel formats for the current configuration.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/frameSupportedPixelFormats
	FrameSupportedPixelFormats() []foundation.NSNumber
}

// VTFrameProcessorConfigurationObject wraps an existing Objective-C object that conforms to the VTFrameProcessorConfiguration protocol.
type VTFrameProcessorConfigurationObject struct {
	objectivec.Object
}

func (o VTFrameProcessorConfigurationObject) BaseObject() objectivec.Object {
	return o.Object
}

// VTFrameProcessorConfigurationObjectFromID constructs a [VTFrameProcessorConfigurationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VTFrameProcessorConfigurationObjectFromID(id objc.ID) VTFrameProcessorConfigurationObject {
	return VTFrameProcessorConfigurationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A dictionary of pixel buffer attributes that define what the source and
// reference frames passed to the processor must conform to.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/sourcePixelBufferAttributes
func (o VTFrameProcessorConfigurationObject) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(rv)
}

// A dictionary of pixel buffer attributes that define which output frames
// passed to the processor must conform to.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/destinationPixelBufferAttributes
func (o VTFrameProcessorConfigurationObject) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(rv)
}

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTFrameProcessorConfigurationObject) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTFrameProcessorConfigurationObject) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

// A list of supported pixel formats for the current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/frameSupportedPixelFormats
func (o VTFrameProcessorConfigurationObject) FrameSupportedPixelFormats() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("frameSupportedPixelFormats"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}
