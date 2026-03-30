// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

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
