// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothObject] class.
var (
	_IOBluetoothObjectClass     IOBluetoothObjectClass
	_IOBluetoothObjectClassOnce sync.Once
)

func getIOBluetoothObjectClass() IOBluetoothObjectClass {
	_IOBluetoothObjectClassOnce.Do(func() {
		_IOBluetoothObjectClass = IOBluetoothObjectClass{class: objc.GetClass("IOBluetoothObject")}
	})
	return _IOBluetoothObjectClass
}

// GetIOBluetoothObjectClass returns the class object for IOBluetoothObject.
func GetIOBluetoothObjectClass() IOBluetoothObjectClass {
	return getIOBluetoothObjectClass()
}

type IOBluetoothObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothObjectClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothObjectClass) Alloc() IOBluetoothObject {
	rv := objc.Send[IOBluetoothObject](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject
type IOBluetoothObject struct {
	objectivec.Object
}

// IOBluetoothObjectFromID constructs a [IOBluetoothObject] from an objc.ID.
func IOBluetoothObjectFromID(id objc.ID) IOBluetoothObject {
	return IOBluetoothObject{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothObject] class.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject
type IIOBluetoothObject interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b IOBluetoothObject) Init() IOBluetoothObject {
	rv := objc.Send[IOBluetoothObject](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothObject) Autorelease() IOBluetoothObject {
	rv := objc.Send[IOBluetoothObject](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothObject creates a new IOBluetoothObject instance.
func NewIOBluetoothObject() IOBluetoothObject {
	class := getIOBluetoothObjectClass()
	rv := objc.Send[IOBluetoothObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}
