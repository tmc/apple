// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods for creating new proxy objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCProxyCreating
type NSXPCProxyCreating interface {
	objectivec.IObject
}

// NSXPCProxyCreatingObject wraps an existing Objective-C object that conforms to the NSXPCProxyCreating protocol.
type NSXPCProxyCreatingObject struct {
	objectivec.Object
}
func (o NSXPCProxyCreatingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSXPCProxyCreatingObjectFromID constructs a [NSXPCProxyCreatingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSXPCProxyCreatingObjectFromID(id objc.ID) NSXPCProxyCreatingObject {
	return NSXPCProxyCreatingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a proxy object with no error handling block.
//
// # Discussion
// 
// Messages sent to the proxy object are sent over the wire to the other side
// of the connection. All messages must be ‘void’ return type. Control may
// be returned to the caller before the message is sent. The resulting proxy
// object conforms to the [NSXPCProxyCreating] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCProxyCreating/remoteObjectProxy()
func (o NSXPCProxyCreatingObject) RemoteObjectProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("remoteObjectProxy"))
	return objectivec.Object{ID: rv}
	}
// Returns a proxy object that invokes the error handling block if an error
// occurs on the connection.
//
// handler: The error handling block that the proxy object should call when an error
// occurs while waiting for a reply.
//
// # Discussion
// 
// If the message sent to the proxy has a reply handler, then either the error
// handler or the reply handler is called exactly once.
// 
// The resulting proxy object conforms to the [NSXPCProxyCreating] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCProxyCreating/remoteObjectProxyWithErrorHandler(_:)
func (o NSXPCProxyCreatingObject) RemoteObjectProxyWithErrorHandler(handler ErrorHandler) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return objectivec.Object{ID: rv}
	}
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCProxyCreating/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (o NSXPCProxyCreatingObject) SynchronousRemoteObjectProxyWithErrorHandler(handler ErrorHandler) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return objectivec.Object{ID: rv}
	}

