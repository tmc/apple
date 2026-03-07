// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods you use to respond to a stack view detaching and reattaching views.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackViewDelegate
type NSStackViewDelegate interface {
	objectivec.IObject
}



// NSStackViewDelegateObject wraps an existing Objective-C object that conforms to the NSStackViewDelegate protocol.
type NSStackViewDelegateObject struct {
	objectivec.Object
}
func (o NSStackViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSStackViewDelegateObjectFromID constructs a [NSStackViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSStackViewDelegateObjectFromID(id objc.ID) NSStackViewDelegateObject {
	return NSStackViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Called when the stack view has automatically reattached one or more
// previously-detached views.
//
// stackView: The stack view that has reattached one or more detached views.
//
// views: An array of one or more views, managed by the stack view, that were
// reattached.
//
// # Discussion
// 
// To configure a custom class to respond to the automatic reattachment of
// views to a stack view’s view hierarchy, implement this method in the
// class. This method is not called when your code explicitly adds a view to a
// stack view’s [Views] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackViewDelegate/stackView(_:didReattach:)

func (o NSStackViewDelegateObject) StackViewDidReattachViews(stackView INSStackView, views []NSView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("stackView:didReattachViews:"), stackView, objectivec.IObjectSliceToNSArray(views))
	}

// Called when the stack view is about to automatically detach one or more of
// its views.
//
// stackView: The stack view that is about to detach one or more of its views.
//
// views: An array of one or more views, managed by the stack view, that are about to
// be automatically detached.
//
// # Discussion
// 
// To configure a custom class to respond to the automatic detachment of views
// from a stack view’s view hierarchy, implement this method in the class.
// This method is not called when your code explicitly removes a view from a
// stack view’s [Views] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackViewDelegate/stackView(_:willDetach:)

func (o NSStackViewDelegateObject) StackViewWillDetachViews(stackView INSStackView, views []NSView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("stackView:willDetachViews:"), stackView, objectivec.IObjectSliceToNSArray(views))
	}







