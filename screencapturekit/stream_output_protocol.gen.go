// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"fmt"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

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
func (o SCStreamOutputObject) StreamDidOutputSampleBufferOfType(stream ISCStream, sampleBuffer coremedia.CMSampleBufferRef, type_ SCStreamOutputType) {
	objc.Send[struct{}](o.ID, objc.Sel("stream:didOutputSampleBuffer:ofType:"), stream, sampleBuffer, type_)
}

// SCStreamOutputConfig holds optional typed callbacks for [SCStreamOutput] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/scstreamoutput
type SCStreamOutputConfig struct {

	// Other Methods
	// StreamDidOutputSampleBufferOfType — Tells the delegate that a capture stream produced a frame.
	StreamDidOutputSampleBufferOfType func(stream SCStream, sampleBuffer coremedia.CMSampleBufferRef, type_ SCStreamOutputType)
}

// NewSCStreamOutput creates an Objective-C object implementing the [SCStreamOutput] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SCStreamOutputObject] satisfies the [SCStreamOutput] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/scstreamoutput
func NewSCStreamOutput(config SCStreamOutputConfig) SCStreamOutputObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSCStreamOutput_%d", n)

	var methods []objc.MethodDef

	if config.StreamDidOutputSampleBufferOfType != nil {
		fn := config.StreamDidOutputSampleBufferOfType
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID, sampleBuffer coremedia.CMSampleBufferRef, type_ SCStreamOutputType) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SCStreamOutput", "stream:didOutputSampleBuffer:ofType:")
					}
				}()
				stream := SCStreamFromID(streamID)
				fn(stream, sampleBuffer, type_)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SCStreamOutput")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSCStreamOutput: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SCStreamOutputObjectFromID(instance)
}
