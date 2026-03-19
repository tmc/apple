// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface an app extension uses to respond to a request from a host app.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionRequestHandling
type NSExtensionRequestHandling interface {
	objectivec.IObject

	// Tells the extension to prepare for a host app’s request.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionRequestHandling/beginRequest(with:)
	BeginRequestWithExtensionContext(context INSExtensionContext)
}

// NSExtensionRequestHandlingObject wraps an existing Objective-C object that conforms to the NSExtensionRequestHandling protocol.
type NSExtensionRequestHandlingObject struct {
	objectivec.Object
}
func (o NSExtensionRequestHandlingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSExtensionRequestHandlingObjectFromID constructs a [NSExtensionRequestHandlingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSExtensionRequestHandlingObjectFromID(id objc.ID) NSExtensionRequestHandlingObject {
	return NSExtensionRequestHandlingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the extension to prepare for a host app’s request.
//
// context: An [NSExtensionContext] object that represents the context in which the
// host app makes the request. Typically, the context contains data that the
// extension can work on.
//
// # Discussion
// 
// An extension prepares for a host app’s request by getting the context
// passed in this method and requesting related data items, if appropriate.
// This method is received after the extension is initialized, but before the
// principal object is asked to do anything with the context. For example, if
// the principal object is a view controller, it receives this message before
// [loadView()] is called. After an extension receives this message, the
// [extensionContext] property of the view controller returns a non`nil`
// value.
// 
// If your subclass conforms to this protocol and overrides ``, the subclass
// is expected to call `[super ]`.
//
// [extensionContext]: https://developer.apple.com/documentation/UIKit/UIViewController/extensionContext
// [loadView()]: https://developer.apple.com/documentation/UIKit/UIViewController/loadView()
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionRequestHandling/beginRequest(with:)
func (o NSExtensionRequestHandlingObject) BeginRequestWithExtensionContext(context INSExtensionContext) {
	
	objc.Send[struct{}](o.ID, objc.Sel("beginRequestWithExtensionContext:"), context)
	}

