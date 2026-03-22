// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that are implemented by the source object in a dragging session.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource
type NSDraggingSource interface {
	objectivec.IObject

	// Declares the types of operations the source allows to be performed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
	DraggingSessionSourceOperationMaskForDraggingContext(session INSDraggingSession, context NSDraggingContext) NSDragOperation
}

// NSDraggingSourceObject wraps an existing Objective-C object that conforms to the NSDraggingSource protocol.
type NSDraggingSourceObject struct {
	objectivec.Object
}
func (o NSDraggingSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDraggingSourceObjectFromID constructs a [NSDraggingSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDraggingSourceObjectFromID(id objc.ID) NSDraggingSourceObject {
	return NSDraggingSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Declares the types of operations the source allows to be performed.
//
// session: The dragging session.
//
// context: The dragging context. See [NSDraggingContext] for the supported values.
// //
// [NSDraggingContext]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
//
// # Return Value
// 
// The appropriate dragging operation as defined in
//
// # Discussion
// 
// In the future Apple may provide more specific “within” values in the
// future. To account for this, for unrecognized localities, return the
// operation mask for the most specific context that you are concerned with.
// The following code is an example of how to implement this functionality:
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
func (o NSDraggingSourceObject) DraggingSessionSourceOperationMaskForDraggingContext(session INSDraggingSession, context NSDraggingContext) NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("draggingSession:sourceOperationMaskForDraggingContext:"), session, context)
	return rv
	}
// Invoked when the drag will begin.
//
// session: The dragging session.
//
// screenPoint: The point where the drag will begin, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:willBeginAt:)
func (o NSDraggingSourceObject) DraggingSessionWillBeginAtPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("draggingSession:willBeginAtPoint:"), session, screenPoint)
	}
// Invoked when the drag moves on the screen.
//
// session: The dragging session.
//
// screenPoint: The point where the drag moved to, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:movedTo:)
func (o NSDraggingSourceObject) DraggingSessionMovedToPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("draggingSession:movedToPoint:"), session, screenPoint)
	}
// Invoked when the dragging session has completed.
//
// session: The dragging session.
//
// screenPoint: The point where the drag ended, in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for drag operation types.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:endedAt:operation:)
func (o NSDraggingSourceObject) DraggingSessionEndedAtPointOperation(session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[struct{}](o.ID, objc.Sel("draggingSession:endedAtPoint:operation:"), session, screenPoint, operation)
	}
// Returns whether the modifier keys will be ignored for this dragging
// session.
//
// session: The dragging session.
//
// # Return Value
// 
// [true] if the modifier keys will be ignored, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/ignoreModifierKeys(for:)
func (o NSDraggingSourceObject) IgnoreModifierKeysForDraggingSession(session INSDraggingSession) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("ignoreModifierKeysForDraggingSession:"), session)
	return rv
	}

