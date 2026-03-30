// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataSourceBlobF4] class.
var (
	_ETDataSourceBlobF4Class     ETDataSourceBlobF4Class
	_ETDataSourceBlobF4ClassOnce sync.Once
)

func getETDataSourceBlobF4Class() ETDataSourceBlobF4Class {
	_ETDataSourceBlobF4ClassOnce.Do(func() {
		_ETDataSourceBlobF4Class = ETDataSourceBlobF4Class{class: objc.GetClass("ETDataSourceBlobF4")}
	})
	return _ETDataSourceBlobF4Class
}

// GetETDataSourceBlobF4Class returns the class object for ETDataSourceBlobF4.
func GetETDataSourceBlobF4Class() ETDataSourceBlobF4Class {
	return getETDataSourceBlobF4Class()
}

type ETDataSourceBlobF4Class struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataSourceBlobF4Class) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataSourceBlobF4Class) Alloc() ETDataSourceBlobF4 {
	rv := objc.Send[ETDataSourceBlobF4](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETDataSourceBlobF4.AddBlobForKey]
//   - [ETDataSourceBlobF4.DataPointAtIndex]
//   - [ETDataSourceBlobF4.NumberOfDataPoints]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceBlobF4
type ETDataSourceBlobF4 struct {
	objectivec.Object
}

// ETDataSourceBlobF4FromID constructs a [ETDataSourceBlobF4] from an objc.ID.
func ETDataSourceBlobF4FromID(id objc.ID) ETDataSourceBlobF4 {
	return ETDataSourceBlobF4{objectivec.Object{ID: id}}
}

// Ensure ETDataSourceBlobF4 implements IETDataSourceBlobF4.
var _ IETDataSourceBlobF4 = ETDataSourceBlobF4{}

// An interface definition for the [ETDataSourceBlobF4] class.
//
// # Methods
//
//   - [IETDataSourceBlobF4.AddBlobForKey]
//   - [IETDataSourceBlobF4.DataPointAtIndex]
//   - [IETDataSourceBlobF4.NumberOfDataPoints]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceBlobF4
type IETDataSourceBlobF4 interface {
	objectivec.IObject

	// Topic: Methods

	AddBlobForKey(blob objectivec.IObject, key objectivec.IObject)
	DataPointAtIndex(index int) objectivec.IObject
	NumberOfDataPoints() int
}

// Init initializes the instance.
func (e ETDataSourceBlobF4) Init() ETDataSourceBlobF4 {
	rv := objc.Send[ETDataSourceBlobF4](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataSourceBlobF4) Autorelease() ETDataSourceBlobF4 {
	rv := objc.Send[ETDataSourceBlobF4](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataSourceBlobF4 creates a new ETDataSourceBlobF4 instance.
func NewETDataSourceBlobF4() ETDataSourceBlobF4 {
	class := getETDataSourceBlobF4Class()
	rv := objc.Send[ETDataSourceBlobF4](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataSourceBlobF4/addBlob:forKey:
func (e ETDataSourceBlobF4) AddBlobForKey(blob objectivec.IObject, key objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("addBlob:forKey:"), blob, key)
}

// See: https://developer.apple.com/documentation/Espresso/ETDataSourceBlobF4/dataPointAtIndex:
func (e ETDataSourceBlobF4) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETDataSourceBlobF4/numberOfDataPoints
func (e ETDataSourceBlobF4) NumberOfDataPoints() int {
	rv := objc.Send[int](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
