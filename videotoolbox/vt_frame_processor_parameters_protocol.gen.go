// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The base protocol for input and output processing parameters for a frame processor implementation.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters
type VTFrameProcessorParameters interface {
	objectivec.IObject

	// A processor frame that contains the current source frame to use for all processing features.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/sourceFrame
	SourceFrame() IVTFrameProcessorFrame

	// Destination frame that contains the destination frame for processors which output a single processed frame.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrame-3im3l
	DestinationFrame() IVTFrameProcessorFrame

	// Array of destination frames for processors which may output more than one processed frame.
	//
	// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
	DestinationFrames() []VTFrameProcessorFrame
}

// VTFrameProcessorParametersObject wraps an existing Objective-C object that conforms to the VTFrameProcessorParameters protocol.
type VTFrameProcessorParametersObject struct {
	objectivec.Object
}

func (o VTFrameProcessorParametersObject) BaseObject() objectivec.Object {
	return o.Object
}

// VTFrameProcessorParametersObjectFromID constructs a [VTFrameProcessorParametersObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VTFrameProcessorParametersObjectFromID(id objc.ID) VTFrameProcessorParametersObject {
	return VTFrameProcessorParametersObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A processor frame that contains the current source frame to use for all
// processing features.
//
// # Discussion
//
// This property must not be [NULL].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/sourceFrame
func (o VTFrameProcessorParametersObject) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(rv)
}

// Destination frame that contains the destination frame for processors which
// output a single processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrame-3im3l
func (o VTFrameProcessorParametersObject) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(rv)
}

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTFrameProcessorParametersObject) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}
