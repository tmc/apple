// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ETDataSource protocol.
type ETDataSource interface {
	objectivec.IObject

	// DataPointAtIndex protocol.
	DataPointAtIndex(index int) objectivec.IObject

	// NumberOfDataPoints protocol.
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

func (o ETDataSourceObject) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (o ETDataSourceObject) NumberOfDataPoints() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
