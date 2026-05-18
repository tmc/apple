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
// # Discussion
//
// The property typically matches one of the common counter set names that
// [MTLCommonCounterSet] defines (see [Confirming which counters and counter
// sets a GPU supports]).
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/name
//
// [Confirming which counters and counter sets a GPU supports]: https://developer.apple.com/documentation/Metal/confirming-which-counters-and-counter-sets-a-gpu-supports
func (o MTLCounterSetObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// An array of the counter instances a GPU device supports.
//
// # Discussion
//
// Check whether a GPU device supports a specific counter by comparing its
// common name (see [MTLCommonCounter]) with each element in the property’s
// array.
//
// For more information, see [Confirming which counters and counter sets a GPU
// supports].
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSet/counters
//
// [Confirming which counters and counter sets a GPU supports]: https://developer.apple.com/documentation/Metal/confirming-which-counters-and-counter-sets-a-gpu-supports
func (o MTLCounterSetObject) Counters() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("counters"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}
