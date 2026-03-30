// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ETDataSource protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ETDataSource
type ETDataSource interface {
	objectivec.IObject

	// NumberOfDataPoints protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ETDataSource/numberOfDataPoints
	NumberOfDataPoints() int
}

// ETDataSourceObject wraps an existing Objective-C object that conforms to the ETDataSource protocol.
type ETDataSourceObject struct {
	objectivec.Object
}

func (o ETDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// ETDataSourceObjectFromID constructs a [ETDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ETDataSourceObjectFromID(id objc.ID) ETDataSourceObject {
	return ETDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ETDataSource/dataPointAtIndex:
func (o ETDataSourceObject) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ETDataSource/numberOfDataPoints
func (o ETDataSourceObject) NumberOfDataPoints() int {
	rv := objc.Send[int](o.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
