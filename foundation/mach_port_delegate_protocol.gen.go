// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// An interface for handling incoming Mach messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPortDelegate
type NSMachPortDelegate interface {
	objectivec.IObject
	NSPortDelegate
}

// NSMachPortDelegateObject wraps an existing Objective-C object that conforms to the NSMachPortDelegate protocol.
type NSMachPortDelegateObject struct {
	objectivec.Object
}

func (o NSMachPortDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSMachPortDelegateObjectFromID constructs a [NSMachPortDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMachPortDelegateObjectFromID(id objc.ID) NSMachPortDelegateObject {
	return NSMachPortDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Process an incoming Mach message.
//
// msg: A pointer to a Mach message, cast as a pointer to void.
//
// # Discussion
//
// The delegate should interpret this data as a pointer to a Mach message
// beginning with a msg_header_t structure and should handle the message
// appropriately.
//
// The delegate should implement either “ or the [NSPortDelegate] protocol
// method handlePortMessage:.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPortDelegate/handleMachMessage(_:)
func (o NSMachPortDelegateObject) HandleMachMessage(msg unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("handleMachMessage:"), msg)
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
func (o NSMachPortDelegateObject) HandlePortMessage(message INSPortMessage) {
	objc.Send[struct{}](o.ID, objc.Sel("handlePortMessage:"), message)
}

// NSMachPortDelegateConfig holds optional typed callbacks for [NSMachPortDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsmachportdelegate
type NSMachPortDelegateConfig struct {

	// Handling Mach messages
	// HandleMachMessage — Process an incoming Mach message.
	HandleMachMessage func(msg kernel.Pointer)
}

// NewNSMachPortDelegate creates an Objective-C object implementing the [NSMachPortDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSMachPortDelegateObject] satisfies the [NSMachPortDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsmachportdelegate
func NewNSMachPortDelegate(config NSMachPortDelegateConfig) NSMachPortDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSMachPortDelegate_%d", n)

	var methods []objc.MethodDef

	if config.HandleMachMessage != nil {
		fn := config.HandleMachMessage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handleMachMessage:"),
			Fn: func(self objc.ID, _cmd objc.SEL, msg kernel.Pointer) {
				fn(msg)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSMachPortDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSMachPortDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSMachPortDelegateObjectFromID(instance)
}
