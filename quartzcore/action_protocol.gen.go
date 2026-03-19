// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An interface that allows instances to respond to actions triggered by a Core Animation layer change.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAction
type CAAction interface {
	objectivec.IObject
}

// CAActionObject wraps an existing Objective-C object that conforms to the CAAction protocol.
type CAActionObject struct {
	objectivec.Object
}
func (o CAActionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CAActionObjectFromID constructs a [CAActionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CAActionObjectFromID(id objc.ID) CAActionObject {
	return CAActionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called to trigger the action specified by the identifier.
//
// event: The identifier of the action. The identifier may be a key or key path
// relative to `anObject`, an arbitrary external action, or one of the action
// identifiers defined in [CALayer].
//
// anObject: The layer on which the action should occur.
//
// dict: A dictionary containing parameters associated with this event. May be
// `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAction/run(forKey:object:arguments:)
func (o CAActionObject) RunActionForKeyObjectArguments(event string, anObject objectivec.IObject, dict foundation.INSDictionary) {
	
	objc.Send[struct{}](o.ID, objc.Sel("runActionForKey:object:arguments:"), objc.String(event), anObject, dict)
	}

