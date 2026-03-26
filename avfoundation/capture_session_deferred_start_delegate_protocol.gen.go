// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that defines the interface to respond to events about a capture session’s deferred start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate
type AVCaptureSessionDeferredStartDelegate interface {
	objectivec.IObject

	// This method gets called by the session when deferred start has finished running.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate/sessionDidRunDeferredStart(_:)
	SessionDidRunDeferredStart(session IAVCaptureSession)

	// This method gets called by the session when deferred start is about to run.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate/sessionWillRunDeferredStart(_:)
	SessionWillRunDeferredStart(session IAVCaptureSession)
}

// AVCaptureSessionDeferredStartDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureSessionDeferredStartDelegate protocol.
type AVCaptureSessionDeferredStartDelegateObject struct {
	objectivec.Object
}
func (o AVCaptureSessionDeferredStartDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureSessionDeferredStartDelegateObjectFromID constructs a [AVCaptureSessionDeferredStartDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureSessionDeferredStartDelegateObjectFromID(id objc.ID) AVCaptureSessionDeferredStartDelegateObject {
	return AVCaptureSessionDeferredStartDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// This method gets called by the session when deferred start has finished
// running.
//
// session: The [AVCaptureSession] instance that runs the deferred start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate/sessionDidRunDeferredStart(_:)
func (o AVCaptureSessionDeferredStartDelegateObject) SessionDidRunDeferredStart(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionDidRunDeferredStart:"), session)
	}
// This method gets called by the session when deferred start is about to run.
//
// session: The [AVCaptureSession] instance that runs the deferred start.
//
// # Discussion
// 
// Delegates receive this message when the session has finished the deferred
// start. This message will be sent regardless of whether the session’s
// [AutomaticallyRunsDeferredStart] property is set. See
// [SetDeferredStartDelegateDeferredStartDelegateCallbackQueue] documentation
// for more information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate/sessionWillRunDeferredStart(_:)
func (o AVCaptureSessionDeferredStartDelegateObject) SessionWillRunDeferredStart(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionWillRunDeferredStart:"), session)
	}

// AVCaptureSessionDeferredStartDelegateConfig holds optional typed callbacks for [AVCaptureSessionDeferredStartDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturesessiondeferredstartdelegate
type AVCaptureSessionDeferredStartDelegateConfig struct {

	// Responding to deferred start events
	// SessionDidRunDeferredStart — This method gets called by the session when deferred start has finished running.
	SessionDidRunDeferredStart func(session AVCaptureSession)
	// SessionWillRunDeferredStart — This method gets called by the session when deferred start is about to run.
	SessionWillRunDeferredStart func(session AVCaptureSession)
}

// NewAVCaptureSessionDeferredStartDelegate creates an Objective-C object implementing the [AVCaptureSessionDeferredStartDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureSessionDeferredStartDelegateObject] satisfies the [AVCaptureSessionDeferredStartDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturesessiondeferredstartdelegate
func NewAVCaptureSessionDeferredStartDelegate(config AVCaptureSessionDeferredStartDelegateConfig) AVCaptureSessionDeferredStartDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureSessionDeferredStartDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SessionDidRunDeferredStart != nil {
		fn := config.SessionDidRunDeferredStart
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionDidRunDeferredStart:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.SessionWillRunDeferredStart != nil {
		fn := config.SessionWillRunDeferredStart
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionWillRunDeferredStart:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureSessionDeferredStartDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureSessionDeferredStartDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureSessionDeferredStartDelegateObjectFromID(instance)
}

