// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that defines the interface to respond to capture control activation and presentation events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate
type AVCaptureSessionControlsDelegate interface {
	objectivec.IObject

	// Tells the delegate when a capture session’s controls become active and available for interaction.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsDidBecomeActive(_:)
	SessionControlsDidBecomeActive(session IAVCaptureSession)

	// Tells the delegate when a capture session’s controls are about to enter a fullscreen appearance.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsWillEnterFullscreenAppearance(_:)
	SessionControlsWillEnterFullscreenAppearance(session IAVCaptureSession)

	// Tells the delegate when a capture session’s controls are about to exit a fullscreen appearance.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsWillExitFullscreenAppearance(_:)
	SessionControlsWillExitFullscreenAppearance(session IAVCaptureSession)

	// Tells the delegate when a capture session’s controls become inactive and unavailable for interaction.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsDidBecomeInactive(_:)
	SessionControlsDidBecomeInactive(session IAVCaptureSession)
}

// AVCaptureSessionControlsDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureSessionControlsDelegate protocol.
type AVCaptureSessionControlsDelegateObject struct {
	objectivec.Object
}
func (o AVCaptureSessionControlsDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureSessionControlsDelegateObjectFromID constructs a [AVCaptureSessionControlsDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureSessionControlsDelegateObjectFromID(id objc.ID) AVCaptureSessionControlsDelegateObject {
	return AVCaptureSessionControlsDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when a capture session’s controls become active and
// available for interaction.
//
// session: The capture session with active controls.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsDidBecomeActive(_:)
func (o AVCaptureSessionControlsDelegateObject) SessionControlsDidBecomeActive(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionControlsDidBecomeActive:"), session)
	}
// Tells the delegate when a capture session’s controls are about to enter a
// fullscreen appearance.
//
// session: The capture session with controls that are entering a fullscreen
// appearance.
//
// # Discussion
// 
// When controls enter a fullscreen appearance, your app should hide portions
// of its user interface, including duplicative or unnecessary elements. Few
// onscreen elements should be visible so people can focus on their control
// interactions while viewing the camera preview unobstructed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsWillEnterFullscreenAppearance(_:)
func (o AVCaptureSessionControlsDelegateObject) SessionControlsWillEnterFullscreenAppearance(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionControlsWillEnterFullscreenAppearance:"), session)
	}
// Tells the delegate when a capture session’s controls are about to exit a
// fullscreen appearance.
//
// session: The capture session with controls that are exiting a fullscreen appearance.
//
// # Discussion
// 
// When your app receives this callback, it should resume showing portions of
// the interface it hid when controls entered a fullscreen appearance.
// 
// The system calls this method before [SessionControlsDidBecomeInactive].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsWillExitFullscreenAppearance(_:)
func (o AVCaptureSessionControlsDelegateObject) SessionControlsWillExitFullscreenAppearance(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionControlsWillExitFullscreenAppearance:"), session)
	}
// Tells the delegate when a capture session’s controls become inactive and
// unavailable for interaction.
//
// session: The capture session with inactive controls.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionControlsDelegate/sessionControlsDidBecomeInactive(_:)
func (o AVCaptureSessionControlsDelegateObject) SessionControlsDidBecomeInactive(session IAVCaptureSession) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionControlsDidBecomeInactive:"), session)
	}

// AVCaptureSessionControlsDelegateConfig holds optional typed callbacks for [AVCaptureSessionControlsDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturesessioncontrolsdelegate
type AVCaptureSessionControlsDelegateConfig struct {

	// Responding to control events
	// SessionControlsDidBecomeActive — Tells the delegate when a capture session’s controls become active and available for interaction.
	SessionControlsDidBecomeActive func(session AVCaptureSession)
	// SessionControlsWillEnterFullscreenAppearance — Tells the delegate when a capture session’s controls are about to enter a fullscreen appearance.
	SessionControlsWillEnterFullscreenAppearance func(session AVCaptureSession)
	// SessionControlsWillExitFullscreenAppearance — Tells the delegate when a capture session’s controls are about to exit a fullscreen appearance.
	SessionControlsWillExitFullscreenAppearance func(session AVCaptureSession)
	// SessionControlsDidBecomeInactive — Tells the delegate when a capture session’s controls become inactive and unavailable for interaction.
	SessionControlsDidBecomeInactive func(session AVCaptureSession)
}

// NewAVCaptureSessionControlsDelegate creates an Objective-C object implementing the [AVCaptureSessionControlsDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureSessionControlsDelegateObject] satisfies the [AVCaptureSessionControlsDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturesessioncontrolsdelegate
func NewAVCaptureSessionControlsDelegate(config AVCaptureSessionControlsDelegateConfig) AVCaptureSessionControlsDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureSessionControlsDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SessionControlsDidBecomeActive != nil {
		fn := config.SessionControlsDidBecomeActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionControlsDidBecomeActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.SessionControlsWillEnterFullscreenAppearance != nil {
		fn := config.SessionControlsWillEnterFullscreenAppearance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionControlsWillEnterFullscreenAppearance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.SessionControlsWillExitFullscreenAppearance != nil {
		fn := config.SessionControlsWillExitFullscreenAppearance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionControlsWillExitFullscreenAppearance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.SessionControlsDidBecomeInactive != nil {
		fn := config.SessionControlsDidBecomeInactive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sessionControlsDidBecomeInactive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVCaptureSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureSessionControlsDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureSessionControlsDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureSessionControlsDelegateObjectFromID(instance)
}

