// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoBrickTensorMetal] class.
var (
	_EspressoBrickTensorMetalClass     EspressoBrickTensorMetalClass
	_EspressoBrickTensorMetalClassOnce sync.Once
)

func getEspressoBrickTensorMetalClass() EspressoBrickTensorMetalClass {
	_EspressoBrickTensorMetalClassOnce.Do(func() {
		_EspressoBrickTensorMetalClass = EspressoBrickTensorMetalClass{class: objc.GetClass("EspressoBrickTensorMetal")}
	})
	return _EspressoBrickTensorMetalClass
}

// GetEspressoBrickTensorMetalClass returns the class object for EspressoBrickTensorMetal.
func GetEspressoBrickTensorMetalClass() EspressoBrickTensorMetalClass {
	return getEspressoBrickTensorMetalClass()
}

type EspressoBrickTensorMetalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoBrickTensorMetalClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoBrickTensorMetalClass) Alloc() EspressoBrickTensorMetal {
	rv := objc.Send[EspressoBrickTensorMetal](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoBrickTensorMetal.Texture]
//   - [EspressoBrickTensorMetal.SetTexture]
type EspressoBrickTensorMetal struct {
	EspressoBrickTensor
}

// EspressoBrickTensorMetalFromID constructs a [EspressoBrickTensorMetal] from an objc.ID.
func EspressoBrickTensorMetalFromID(id objc.ID) EspressoBrickTensorMetal {
	return EspressoBrickTensorMetal{EspressoBrickTensor: EspressoBrickTensorFromID(id)}
}

// Ensure EspressoBrickTensorMetal implements IEspressoBrickTensorMetal.
var _ IEspressoBrickTensorMetal = EspressoBrickTensorMetal{}

// An interface definition for the [EspressoBrickTensorMetal] class.
//
// # Methods
//
//   - [IEspressoBrickTensorMetal.Texture]
//   - [IEspressoBrickTensorMetal.SetTexture]
type IEspressoBrickTensorMetal interface {
	IEspressoBrickTensor

	// Topic: Methods

	Texture() unsafe.Pointer
	SetTexture(value kernel.Pointer)
}

// Init initializes the instance.
func (e EspressoBrickTensorMetal) Init() EspressoBrickTensorMetal {
	rv := objc.Send[EspressoBrickTensorMetal](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickTensorMetal) Autorelease() EspressoBrickTensorMetal {
	rv := objc.Send[EspressoBrickTensorMetal](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickTensorMetal creates a new EspressoBrickTensorMetal instance.
func NewEspressoBrickTensorMetal() EspressoBrickTensorMetal {
	class := getEspressoBrickTensorMetalClass()
	rv := objc.Send[EspressoBrickTensorMetal](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e EspressoBrickTensorMetal) Texture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("texture"))
	return rv
}
func (e EspressoBrickTensorMetal) SetTexture(value kernel.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setTexture:"), value)
}
