// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BSInvalidatable protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BSInvalidatable
type BSInvalidatable interface {
	objectivec.IObject

	// Invalidate protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/BSInvalidatable/invalidate
	Invalidate()
}

// BSInvalidatableObject wraps an existing Objective-C object that conforms to the BSInvalidatable protocol.
type BSInvalidatableObject struct {
	objectivec.Object
}

func (o BSInvalidatableObject) BaseObject() objectivec.Object {
	return o.Object
}

// BSInvalidatableObjectFromID constructs a [BSInvalidatableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BSInvalidatableObjectFromID(id objc.ID) BSInvalidatableObject {
	return BSInvalidatableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BSInvalidatable/invalidate
func (o BSInvalidatableObject) Invalidate() {
	objc.Send[struct{}](o.ID, objc.Sel("invalidate"))
}
