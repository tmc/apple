// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the interface for a video composition instruction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol
type AVVideoCompositionInstructionProtocol interface {
	objectivec.IObject

	// An identifier of a source track to pass through without compositing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/passthroughTrackID
	PassthroughTrackID() int32

	// The identifiers of the video tracks the instruction requires to compose frames.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/requiredSourceTrackIDs
	RequiredSourceTrackIDs() []foundation.NSValue

	// The identifiers of the sample data tracks the instruction requires to compose frames.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/requiredSourceSampleDataTrackIDs
	RequiredSourceSampleDataTrackIDs() []foundation.NSNumber

	// A Boolean value that indicates whether the composition contains tweening.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/containsTweening
	ContainsTweening() bool

	// A Boolean value that indicates whether the composition enables post-processing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/enablePostProcessing
	EnablePostProcessing() bool

	// The time range during which the instruction is effective.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/timeRange
	TimeRange() coremedia.CMTimeRange
}

// AVVideoCompositionInstructionProtocolObject wraps an existing Objective-C object that conforms to the AVVideoCompositionInstructionProtocol protocol.
type AVVideoCompositionInstructionProtocolObject struct {
	objectivec.Object
}
func (o AVVideoCompositionInstructionProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVVideoCompositionInstructionProtocolObjectFromID constructs a [AVVideoCompositionInstructionProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVVideoCompositionInstructionProtocolObjectFromID(id objc.ID) AVVideoCompositionInstructionProtocolObject {
	return AVVideoCompositionInstructionProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An identifier of a source track to pass through without compositing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/passthroughTrackID
func (o AVVideoCompositionInstructionProtocolObject) PassthroughTrackID() int32 {
	rv := objc.Send[int32](o.ID, objc.Sel("passthroughTrackID"))
	return rv
	}
// The identifiers of the video tracks the instruction requires to compose
// frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/requiredSourceTrackIDs
func (o AVVideoCompositionInstructionProtocolObject) RequiredSourceTrackIDs() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("requiredSourceTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}
// The identifiers of the sample data tracks the instruction requires to
// compose frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/requiredSourceSampleDataTrackIDs
func (o AVVideoCompositionInstructionProtocolObject) RequiredSourceSampleDataTrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
	}
// A Boolean value that indicates whether the composition contains tweening.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/containsTweening
func (o AVVideoCompositionInstructionProtocolObject) ContainsTweening() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("containsTweening"))
	return rv
	}
// A Boolean value that indicates whether the composition enables
// post-processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/enablePostProcessing
func (o AVVideoCompositionInstructionProtocolObject) EnablePostProcessing() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enablePostProcessing"))
	return rv
	}
// The time range during which the instruction is effective.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/timeRange
func (o AVVideoCompositionInstructionProtocolObject) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](o.ID, objc.Sel("timeRange"))
	return rv
	}

