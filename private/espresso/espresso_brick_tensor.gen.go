// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoBrickTensor] class.
var (
	_EspressoBrickTensorClass     EspressoBrickTensorClass
	_EspressoBrickTensorClassOnce sync.Once
)

func getEspressoBrickTensorClass() EspressoBrickTensorClass {
	_EspressoBrickTensorClassOnce.Do(func() {
		_EspressoBrickTensorClass = EspressoBrickTensorClass{class: objc.GetClass("EspressoBrickTensor")}
	})
	return _EspressoBrickTensorClass
}

// GetEspressoBrickTensorClass returns the class object for EspressoBrickTensor.
func GetEspressoBrickTensorClass() EspressoBrickTensorClass {
	return getEspressoBrickTensorClass()
}

type EspressoBrickTensorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoBrickTensorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoBrickTensorClass) Alloc() EspressoBrickTensor {
	rv := objc.Send[EspressoBrickTensor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoBrickTensor.Shape]
//   - [EspressoBrickTensor.SetShape]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensor
type EspressoBrickTensor struct {
	objectivec.Object
}

// EspressoBrickTensorFromID constructs a [EspressoBrickTensor] from an objc.ID.
func EspressoBrickTensorFromID(id objc.ID) EspressoBrickTensor {
	return EspressoBrickTensor{objectivec.Object{ID: id}}
}

// Ensure EspressoBrickTensor implements IEspressoBrickTensor.
var _ IEspressoBrickTensor = EspressoBrickTensor{}

// An interface definition for the [EspressoBrickTensor] class.
//
// # Methods
//
//   - [IEspressoBrickTensor.Shape]
//   - [IEspressoBrickTensor.SetShape]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensor
type IEspressoBrickTensor interface {
	objectivec.IObject

	// Topic: Methods

	Shape() IEspressoBrickTensorShape
	SetShape(value IEspressoBrickTensorShape)
}

// Init initializes the instance.
func (e EspressoBrickTensor) Init() EspressoBrickTensor {
	rv := objc.Send[EspressoBrickTensor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickTensor) Autorelease() EspressoBrickTensor {
	rv := objc.Send[EspressoBrickTensor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickTensor creates a new EspressoBrickTensor instance.
func NewEspressoBrickTensor() EspressoBrickTensor {
	class := getEspressoBrickTensorClass()
	rv := objc.Send[EspressoBrickTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensor/shape
func (e EspressoBrickTensor) Shape() IEspressoBrickTensorShape {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shape"))
	return EspressoBrickTensorShapeFromID(objc.ID(rv))
}
func (e EspressoBrickTensor) SetShape(value IEspressoBrickTensorShape) {
	objc.Send[struct{}](e.ID, objc.Sel("setShape:"), value)
}
