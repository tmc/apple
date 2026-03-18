// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that support the mediation of a custom segue.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming
type NSSeguePerforming interface {
	objectivec.IObject
}

// NSSeguePerformingObject wraps an existing Objective-C object that conforms to the NSSeguePerforming protocol.
type NSSeguePerformingObject struct {
	objectivec.Object
}
func (o NSSeguePerformingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSeguePerformingObjectFromID constructs a [NSSeguePerformingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSeguePerformingObjectFromID(id objc.ID) NSSeguePerformingObject {
	return NSSeguePerformingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Performs the specified segue.
//
// identifier: The string that uniquely identifies the segue in the storyboard file.
// 
// In Interface Builder, you can provide an identifier string to a segue using
// the inspector. Pass this string to this parameter.
//
// sender: The object that you want to use to initiate the segue. This parameter makes
// the object available to your implementation during the segue.
//
// # Discussion
// 
// Apps typically do not need to trigger segues programmatically. If needed,
// you can call this method to trigger a segue for an action that cannot be
// expressed in a storyboard file, such as a transition between scenes in
// different storyboards.
// 
// Typically, a segue is triggered by a user action, such as clicking a
// button. In Interface Builder, configure an object, such as a control
// embedded in the view controller’s view hierarchy, to trigger the segue.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/performSegue(withIdentifier:sender:)

func (o NSSeguePerformingObject) PerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("performSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
	}

// Called when a segue is about to be performed.
//
// segue: The segue object containing information about the view controllers involved
// in the segue.
//
// sender: The object that initiated the segue. You might use this parameter to
// perform different actions based on which control (or other object)
// initiated the segue.
//
// # Discussion
// 
// The default implementation of this method does nothing; you can override it
// to pass relevant data to the new view controller or window controller,
// based on the context of the segue. The `segue` object describes the
// transition and includes references to both controllers involved in the
// segue.
// 
// Segues can be triggered from multiple sources, so use the information in
// the `segue` and `sender` parameters to disambiguate between different
// logical paths in your app. For example, if the segue originated from a
// table view, the `sender` parameter would identify the cell that the user
// clicked. You could use that information to set the data on the destination
// view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/prepare(for:sender:)

func (o NSSeguePerformingObject) PrepareForSegueSender(segue INSStoryboardSegue, sender objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("prepareForSegue:sender:"), segue, sender)
	}

// Called immediately prior to the performance of a storyboard segue.
//
// identifier: The string that identifies the segue to be performed.
// 
// Using the Interface Builder inspector, provide a unique identifier string
// for each segue in a storyboard. The system provides a segue’s identifier
// to this parameter when it calls this method. The identifier string is used
// to locate the segue inside the storyboard file that contains the view
// controller.
//
// sender: The object that initiated the segue. This object is made available for
// informational purposes during the segue.
//
// # Return Value
// 
// [true] to allow a segue to proceed or [false] to stop it from proceeding.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Override this method to return [false] for cases where you want to prevent
// the performance of a segue. By default, invocation of a segue results in
// the segue being performed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/shouldPerformSegue(withIdentifier:sender:)

func (o NSSeguePerformingObject) ShouldPerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("shouldPerformSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
	return rv
	}

