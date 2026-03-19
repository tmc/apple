// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The interface through which a user activity instance notifies its delegate of updates.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivityDelegate
type NSUserActivityDelegate interface {
	objectivec.IObject
}

// NSUserActivityDelegateObject wraps an existing Objective-C object that conforms to the NSUserActivityDelegate protocol.
type NSUserActivityDelegateObject struct {
	objectivec.Object
}
func (o NSUserActivityDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserActivityDelegateObjectFromID constructs a [NSUserActivityDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserActivityDelegateObjectFromID(id objc.ID) NSUserActivityDelegateObject {
	return NSUserActivityDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the user activity delegate that an input and output streams are
// available to open.
//
// userActivity: The user activity that is continuing on another device. This user
// activity’s [SupportsContinuationStreams] property must be [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// inputStream: The stream from which the originating app can read data written from the
// continuing app.
//
// outputStream: The stream to which the originating app writes data to be read by the
// continuing app.
//
// # Discussion
// 
// If [SupportsContinuationStreams] is [true], the continuing app can request
// streams back to the originating app. This delegate callback is received
// with the streams from the continuing side. The streams are provided in an
// unopened state, and the delegate should open them immediately to start
// communicating with the continuing side.
// 
// Continuation streams are an optional feature of Handoff, and most user
// activities do not need them for successful continuation. When streams are
// needed, a simple request from the continuing app accompanied by a response
// from the originating app is enough for most continuation events.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivityDelegate/userActivity(_:didReceive:outputStream:)
func (o NSUserActivityDelegateObject) UserActivityDidReceiveInputStreamOutputStream(userActivity INSUserActivity, inputStream INSInputStream, outputStream INSOutputStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("userActivity:didReceiveInputStream:outputStream:"), userActivity, inputStream, outputStream)
	}
// Notifies the delegate that the user activity was continued on another
// device.
//
// userActivity: The user activity that was continued.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivityDelegate/userActivityWasContinued(_:)
func (o NSUserActivityDelegateObject) UserActivityWasContinued(userActivity INSUserActivity) {
	
	objc.Send[struct{}](o.ID, objc.Sel("userActivityWasContinued:"), userActivity)
	}
// Notifies the delegate that the user activity will be saved to be continued
// or persisted.
//
// userActivity: The user activity to update.
//
// # Discussion
// 
// The delegate overrides this method to update the activity with current
// state.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivityDelegate/userActivityWillSave(_:)
func (o NSUserActivityDelegateObject) UserActivityWillSave(userActivity INSUserActivity) {
	
	objc.Send[struct{}](o.ID, objc.Sel("userActivityWillSave:"), userActivity)
	}

// NSUserActivityDelegateConfig holds optional typed callbacks for [NSUserActivityDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate
type NSUserActivityDelegateConfig struct {

	// Managing activity continuation
	// UserActivityWasContinued — Notifies the delegate that the user activity was continued on another device.
	UserActivityWasContinued func(userActivity NSUserActivity)
	// UserActivityWillSave — Notifies the delegate that the user activity will be saved to be continued or persisted.
	UserActivityWillSave func(userActivity NSUserActivity)

	// Other Methods
	// UserActivityDidReceiveInputStreamOutputStream — Notifies the user activity delegate that an input and output streams are available to open.
	UserActivityDidReceiveInputStreamOutputStream func(userActivity NSUserActivity, inputStream NSInputStream, outputStream NSOutputStream)
}

// NewNSUserActivityDelegate creates an Objective-C object implementing the [NSUserActivityDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSUserActivityDelegateObject] satisfies the [NSUserActivityDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate
func NewNSUserActivityDelegate(config NSUserActivityDelegateConfig) NSUserActivityDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSUserActivityDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UserActivityDidReceiveInputStreamOutputStream != nil {
		fn := config.UserActivityDidReceiveInputStreamOutputStream
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("userActivity:didReceiveInputStream:outputStream:"),
			Fn: func(self objc.ID, _cmd objc.SEL, userActivityID objc.ID, inputStreamID objc.ID, outputStreamID objc.ID) {
				userActivity := NSUserActivityFromID(userActivityID)
				inputStream := NSInputStreamFromID(inputStreamID)
				outputStream := NSOutputStreamFromID(outputStreamID)
				fn(userActivity, inputStream, outputStream)
			},
		})
	}

	if config.UserActivityWasContinued != nil {
		fn := config.UserActivityWasContinued
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("userActivityWasContinued:"),
			Fn: func(self objc.ID, _cmd objc.SEL, userActivityID objc.ID) {
				userActivity := NSUserActivityFromID(userActivityID)
				fn(userActivity)
			},
		})
	}

	if config.UserActivityWillSave != nil {
		fn := config.UserActivityWillSave
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("userActivityWillSave:"),
			Fn: func(self objc.ID, _cmd objc.SEL, userActivityID objc.ID) {
				userActivity := NSUserActivityFromID(userActivityID)
				fn(userActivity)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSUserActivityDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSUserActivityDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSUserActivityDelegateObjectFromID(instance)
}

