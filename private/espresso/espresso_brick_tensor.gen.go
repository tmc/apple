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
	rv := objc.SendIfResponds[EspressoBrickTensor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoBrickTensor.Shape]
//   - [EspressoBrickTensor.SetShape]
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
type IEspressoBrickTensor interface {
	objectivec.IObject

	// Topic: Methods

	Shape() IEspressoBrickTensorShape
	SetShape(value IEspressoBrickTensorShape)
}

// Init initializes the instance.
func (e EspressoBrickTensor) Init() EspressoBrickTensor {
	rv := objc.SendIfResponds[EspressoBrickTensor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickTensor) Autorelease() EspressoBrickTensor {
	rv := objc.SendIfResponds[EspressoBrickTensor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickTensor creates a new EspressoBrickTensor instance.
func NewEspressoBrickTensor() EspressoBrickTensor {
	class := getEspressoBrickTensorClass()
	rv := objc.SendIfResponds[EspressoBrickTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e EspressoBrickTensor) Shape() IEspressoBrickTensorShape {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("shape"))
	return EspressoBrickTensorShapeFromID(objc.ID(rv))
}
func (e EspressoBrickTensor) SetShape(value IEspressoBrickTensorShape) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setShape:"), value)
}
