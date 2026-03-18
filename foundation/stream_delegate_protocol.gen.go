// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// An interface that delegates of a stream instance use to handle events on the stream.
//
// See: https://developer.apple.com/documentation/Foundation/StreamDelegate
type NSStreamDelegate interface {
	objectivec.IObject
}

// NSStreamDelegateObject wraps an existing Objective-C object that conforms to the NSStreamDelegate protocol.
type NSStreamDelegateObject struct {
	objectivec.Object
}
func (o NSStreamDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSStreamDelegateObjectFromID constructs a [NSStreamDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSStreamDelegateObjectFromID(id objc.ID) NSStreamDelegateObject {
	return NSStreamDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The delegate receives this message when a given event has occurred on a
// given stream.
//
// aStream: The stream on which `streamEvent` occurred.
//
// eventCode: The stream event that occurred.
//
// # Discussion
// 
// The delegate receives this message only if `theStream` is scheduled on a
// run loop. The message is sent on the stream object’s thread. The delegate
// should examine `streamEvent` to determine the appropriate action it should
// take.
//
// See: https://developer.apple.com/documentation/Foundation/StreamDelegate/stream(_:handle:)

func (o NSStreamDelegateObject) StreamHandleEvent(aStream INSStream, eventCode NSStreamEvent) {
	
	objc.Send[struct{}](o.ID, objc.Sel("stream:handleEvent:"), aStream, eventCode)
	}

// NSStreamDelegateConfig holds optional typed callbacks for [NSStreamDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsstreamdelegate
type NSStreamDelegateConfig struct {

	// Other Methods
	// StreamHandleEvent — The delegate receives this message when a given event has occurred on a given stream.
	StreamHandleEvent func(aStream NSStream, eventCode NSStreamEvent)
}

// NewNSStreamDelegate creates an Objective-C object implementing the [NSStreamDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSStreamDelegateObject] satisfies the [NSStreamDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsstreamdelegate
func NewNSStreamDelegate(config NSStreamDelegateConfig) NSStreamDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSStreamDelegate_%d", n)

	var methods []objc.MethodDef

	if config.StreamHandleEvent != nil {
		fn := config.StreamHandleEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("stream:handleEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, aStreamID objc.ID, eventCode NSStreamEvent) {
				aStream := NSStreamFromID(aStreamID)
				fn(aStream, eventCode)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSStreamDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSStreamDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSStreamDelegateObjectFromID(instance)
}

