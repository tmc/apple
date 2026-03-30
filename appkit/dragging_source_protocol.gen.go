// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
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
//
// # Return Value
//
// A dragging operation you specify.
//
// # Discussion
//
// To account for unexpected contexts, set a `default` case for the most
// specific context your app handles. The following code shows an example that
// handles different dragging contexts and includes a default case.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
//
// [NSDraggingContext]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
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
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:endedAt:operation:)
//
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
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
// true if the modifier keys will be ignored, false otherwise.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/ignoreModifierKeys(for:)
func (o NSDraggingSourceObject) IgnoreModifierKeysForDraggingSession(session INSDraggingSession) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("ignoreModifierKeysForDraggingSession:"), session)
	return rv
}
