// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// An interface for handling incoming messages.
//
// See: https://developer.apple.com/documentation/Foundation/PortDelegate
type NSPortDelegate interface {
	objectivec.IObject
}

// NSPortDelegateObject wraps an existing Objective-C object that conforms to the NSPortDelegate protocol.
type NSPortDelegateObject struct {
	objectivec.Object
}

func (o NSPortDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPortDelegateObjectFromID constructs a [NSPortDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPortDelegateObjectFromID(id objc.ID) NSPortDelegateObject {
	return NSPortDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Processes a given incoming message on the port.
//
// message: An incoming port message.
//
// # Discussion
//
// See [NSPort] for more information.
//
// The delegate should implement either [HandlePortMessage] or the
// [NSMachPortDelegate] protocol method [HandleMachMessage]. You must not
// implement both delegate methods.
//
// See: https://developer.apple.com/documentation/Foundation/PortDelegate/handle(_:)
func (o NSPortDelegateObject) HandlePortMessage(message INSPortMessage) {
	objc.Send[struct{}](o.ID, objc.Sel("handlePortMessage:"), message)
}

// NSPortDelegateConfig holds optional typed callbacks for [NSPortDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsportdelegate
type NSPortDelegateConfig struct {

	// Other Methods
	// HandlePortMessage — Processes a given incoming message on the port.
	HandlePortMessage func(message NSPortMessage)
}

// NewNSPortDelegate creates an Objective-C object implementing the [NSPortDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSPortDelegateObject] satisfies the [NSPortDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsportdelegate
func NewNSPortDelegate(config NSPortDelegateConfig) NSPortDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSPortDelegate_%d", n)

	var methods []objc.MethodDef

	if config.HandlePortMessage != nil {
		fn := config.HandlePortMessage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handlePortMessage:"),
			Fn: func(self objc.ID, _cmd objc.SEL, messageID objc.ID) {
				message := NSPortMessageFromID(messageID)
				fn(message)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSPortDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSPortDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSPortDelegateObjectFromID(instance)
}
