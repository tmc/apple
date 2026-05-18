// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An individual counter a GPU device lists within one of its counter sets.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounter
type MTLCounter interface {
	objectivec.IObject

	// The name of a GPU’s counter instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounter/name
	Name() string
}

// MTLCounterObject wraps an existing Objective-C object that conforms to the MTLCounter protocol.
type MTLCounterObject struct {
	objectivec.Object
}

func (o MTLCounterObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCounterObjectFromID constructs a [MTLCounterObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCounterObjectFromID(id objc.ID) MTLCounterObject {
	return MTLCounterObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The name of a GPU’s counter instance.
//
// # Discussion
//
// The property typically matches one of the common counter names that
// [MTLCommonCounter] defines (see [Confirming which counters and counter sets
// a GPU supports]).
//
// See: https://developer.apple.com/documentation/Metal/MTLCounter/name
//
// [Confirming which counters and counter sets a GPU supports]: https://developer.apple.com/documentation/Metal/confirming-which-counters-and-counter-sets-a-gpu-supports
func (o MTLCounterObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
