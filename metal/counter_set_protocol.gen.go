// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A collection of individual counters a GPU device supports for a counter set.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSet
type MTLCounterSet interface {
	objectivec.IObject

	// The name of the GPU’s counter set instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/name
	Name() string

	// An array of the counter instances a GPU device supports.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/counters
	Counters() []objectivec.IObject
}

// MTLCounterSetObject wraps an existing Objective-C object that conforms to the MTLCounterSet protocol.
type MTLCounterSetObject struct {
	objectivec.Object
}

func (o MTLCounterSetObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCounterSetObjectFromID constructs a [MTLCounterSetObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCounterSetObjectFromID(id objc.ID) MTLCounterSetObject {
	return MTLCounterSetObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The name of the GPU’s counter set instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/name
func (o MTLCounterSetObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// An array of the counter instances a GPU device supports.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/counters
func (o MTLCounterSetObject) Counters() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("counters"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
