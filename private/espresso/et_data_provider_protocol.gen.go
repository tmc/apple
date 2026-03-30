// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ETDataProvider protocol.
//
// See: https://developer.apple.com/documentation/Espresso/ETDataProvider
type ETDataProvider interface {
	objectivec.IObject

	// NumberOfDataPoints protocol.
	//
	// See: https://developer.apple.com/documentation/Espresso/ETDataProvider/numberOfDataPoints
	NumberOfDataPoints() uint64

	// PrepareForEpoch protocol.
	//
	// See: https://developer.apple.com/documentation/Espresso/ETDataProvider/prepareForEpoch
	PrepareForEpoch()
}

// ETDataProviderObject wraps an existing Objective-C object that conforms to the ETDataProvider protocol.
type ETDataProviderObject struct {
	objectivec.Object
}

func (o ETDataProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// ETDataProviderObjectFromID constructs a [ETDataProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ETDataProviderObjectFromID(id objc.ID) ETDataProviderObject {
	return ETDataProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Espresso/ETDataProvider/dataPointAtIndex:error:
func (o ETDataProviderObject) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("dataPointAtIndex:error:"), index)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/Espresso/ETDataProvider/numberOfDataPoints
func (o ETDataProviderObject) NumberOfDataPoints() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numberOfDataPoints"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataProvider/prepareForEpoch
func (o ETDataProviderObject) PrepareForEpoch() {
	objc.Send[struct{}](o.ID, objc.Sel("prepareForEpoch"))
}
