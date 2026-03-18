// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A delegate protocol your app implements to respond to stream events.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate
type SCStreamDelegate interface {
	objectivec.IObject
}

// SCStreamDelegateObject wraps an existing Objective-C object that conforms to the SCStreamDelegate protocol.
type SCStreamDelegateObject struct {
	objectivec.Object
}
func (o SCStreamDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SCStreamDelegateObjectFromID constructs a [SCStreamDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SCStreamDelegateObjectFromID(id objc.ID) SCStreamDelegateObject {
	return SCStreamDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that Presenter Overlay started.
//
// stream: The stream using Presenter Overlay.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate/outputVideoEffectDidStart(for:)

func (o SCStreamDelegateObject) OutputVideoEffectDidStartForStream(stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outputVideoEffectDidStartForStream:"), stream)
	}

// Tells the delegate that Presenter Overlay stopped.
//
// stream: The stream that was using Presenter Overlay.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate/outputVideoEffectDidStop(for:)

func (o SCStreamDelegateObject) OutputVideoEffectDidStopForStream(stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outputVideoEffectDidStopForStream:"), stream)
	}

// Tells the delegate that the stream stopped with an error.
//
// stream: The stream that stopped.
//
// error: The error that caused the stream stoppage.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate/stream(_:didStopWithError:)

func (o SCStreamDelegateObject) StreamDidStopWithError(stream ISCStream, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("stream:didStopWithError:"), stream, error_)
	}

//
// stream: The SCStream object
//
// # Discussion
// 
// streamDidBecomeActive:
// 
// notifies the delegate the first time any window that was being shared in
// the stream is re-opened after all the windows being shared are closed. When
// all the windows being shared are closed, the client will receive
// streamDidBecomeInactive:.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate/streamDidBecomeActive(_:)

func (o SCStreamDelegateObject) StreamDidBecomeActive(stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("streamDidBecomeActive:"), stream)
	}

//
// stream: The SCStream object
//
// # Discussion
// 
// streamDidBecomeInactive:
// 
// notifies the delegate that all the windows that are currently being shared
// are exited. This callback occurs for all content filter types.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamDelegate/streamDidBecomeInactive(_:)

func (o SCStreamDelegateObject) StreamDidBecomeInactive(stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("streamDidBecomeInactive:"), stream)
	}

// SCStreamDelegateConfig holds optional typed callbacks for [SCStreamDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/scstreamdelegate
type SCStreamDelegateConfig struct {

	// Responding to stream stoppage
	// StreamDidStopWithError — Tells the delegate that the stream stopped with an error.
	StreamDidStopWithError func(stream SCStream, error_ foundation.NSError)

	// Instance Methods
	StreamDidBecomeActive func(stream SCStream)
	StreamDidBecomeInactive func(stream SCStream)

	// Other Methods
	// OutputVideoEffectDidStartForStream — Tells the delegate that Presenter Overlay started.
	OutputVideoEffectDidStartForStream func(stream SCStream)
	// OutputVideoEffectDidStopForStream — Tells the delegate that Presenter Overlay stopped.
	OutputVideoEffectDidStopForStream func(stream SCStream)
}

// NewSCStreamDelegate creates an Objective-C object implementing the [SCStreamDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SCStreamDelegateObject] satisfies the [SCStreamDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/scstreamdelegate
func NewSCStreamDelegate(config SCStreamDelegateConfig) SCStreamDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSCStreamDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OutputVideoEffectDidStartForStream != nil {
		fn := config.OutputVideoEffectDidStartForStream
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputVideoEffectDidStartForStream:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID) {
				stream := SCStreamFromID(streamID)
				fn(stream)
			},
		})
	}

	if config.OutputVideoEffectDidStopForStream != nil {
		fn := config.OutputVideoEffectDidStopForStream
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputVideoEffectDidStopForStream:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID) {
				stream := SCStreamFromID(streamID)
				fn(stream)
			},
		})
	}

	if config.StreamDidStopWithError != nil {
		fn := config.StreamDidStopWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("stream:didStopWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID, error_ID objc.ID) {
				stream := SCStreamFromID(streamID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(stream, error_)
			},
		})
	}

	if config.StreamDidBecomeActive != nil {
		fn := config.StreamDidBecomeActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("streamDidBecomeActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID) {
				stream := SCStreamFromID(streamID)
				fn(stream)
			},
		})
	}

	if config.StreamDidBecomeInactive != nil {
		fn := config.StreamDidBecomeInactive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("streamDidBecomeInactive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, streamID objc.ID) {
				stream := SCStreamFromID(streamID)
				fn(stream)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SCStreamDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSCStreamDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SCStreamDelegateObjectFromID(instance)
}

