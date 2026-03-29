// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacSerialNumber] class.
var (
	_VZMacSerialNumberClass     VZMacSerialNumberClass
	_VZMacSerialNumberClassOnce sync.Once
)

func getVZMacSerialNumberClass() VZMacSerialNumberClass {
	_VZMacSerialNumberClassOnce.Do(func() {
		_VZMacSerialNumberClass = VZMacSerialNumberClass{class: objc.GetClass("_VZMacSerialNumber")}
	})
	return _VZMacSerialNumberClass
}

// GetVZMacSerialNumberClass returns the class object for _VZMacSerialNumber.
func GetVZMacSerialNumberClass() VZMacSerialNumberClass {
	return getVZMacSerialNumberClass()
}

type VZMacSerialNumberClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacSerialNumberClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacSerialNumberClass) Alloc() VZMacSerialNumber {
	rv := objc.Send[VZMacSerialNumber](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacSerialNumber.String]
//   - [VZMacSerialNumber.InitWithString]
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSerialNumber
type VZMacSerialNumber struct {
	objectivec.Object
}

// VZMacSerialNumberFromID constructs a [VZMacSerialNumber] from an objc.ID.
func VZMacSerialNumberFromID(id objc.ID) VZMacSerialNumber {
	return VZMacSerialNumber{objectivec.Object{ID: id}}
}
// Ensure VZMacSerialNumber implements IVZMacSerialNumber.
var _ IVZMacSerialNumber = VZMacSerialNumber{}

// An interface definition for the [VZMacSerialNumber] class.
//
// # Methods
//
//   - [IVZMacSerialNumber.String]
//   - [IVZMacSerialNumber.InitWithString]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSerialNumber
type IVZMacSerialNumber interface {
	objectivec.IObject

	// Topic: Methods

	String() string
	InitWithString(string_ objectivec.IObject) VZMacSerialNumber
}

// Init initializes the instance.
func (v VZMacSerialNumber) Init() VZMacSerialNumber {
	rv := objc.Send[VZMacSerialNumber](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacSerialNumber) Autorelease() VZMacSerialNumber {
	rv := objc.Send[VZMacSerialNumber](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacSerialNumber creates a new VZMacSerialNumber instance.
func NewVZMacSerialNumber() VZMacSerialNumber {
	class := getVZMacSerialNumberClass()
	rv := objc.Send[VZMacSerialNumber](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSerialNumber/initWithString:
func NewVZMacSerialNumberWithString(string_ objectivec.IObject) VZMacSerialNumber {
	instance := getVZMacSerialNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), string_)
	return VZMacSerialNumberFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacSerialNumber/initWithString:
func (v VZMacSerialNumber) InitWithString(string_ objectivec.IObject) VZMacSerialNumber {
	rv := objc.Send[VZMacSerialNumber](v.ID, objc.Sel("initWithString:"), string_)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacSerialNumber/string
func (v VZMacSerialNumber) String() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

