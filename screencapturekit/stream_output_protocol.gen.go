// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A delegate protocol your app implements to receive capture stream output events.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutput
type SCStreamOutput interface {
	objectivec.IObject
}



// SCStreamOutputObject wraps an existing Objective-C object that conforms to the SCStreamOutput protocol.
type SCStreamOutputObject struct {
	objectivec.Object
}
func (o SCStreamOutputObject) BaseObject() objectivec.Object {
	return o.Object
}



// SCStreamOutputObjectFromID constructs a [SCStreamOutputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SCStreamOutputObjectFromID(id objc.ID) SCStreamOutputObject {
	return SCStreamOutputObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that a capture stream produced a frame.
//
// stream: The frame capture stream that produced this frame.
//
// sampleBuffer: The sample buffer containing capture data.
//
// type: The type of capture contained in the sample buffer.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutput/stream(_:didOutputSampleBuffer:of:)

func (o SCStreamOutputObject) StreamDidOutputSampleBufferOfType(stream ISCStream, sampleBuffer objectivec.IObject, type_ SCStreamOutputType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("stream:didOutputSampleBuffer:ofType:"), stream, sampleBuffer, type_)
	}








