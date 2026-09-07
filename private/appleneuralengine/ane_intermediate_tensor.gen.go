// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEIntermediateTensor] class.
var (
	_ANEIntermediateTensorClass     ANEIntermediateTensorClass
	_ANEIntermediateTensorClassOnce sync.Once
)

func getANEIntermediateTensorClass() ANEIntermediateTensorClass {
	_ANEIntermediateTensorClassOnce.Do(func() {
		_ANEIntermediateTensorClass = ANEIntermediateTensorClass{class: objc.GetClass("_ANEIntermediateTensor")}
	})
	return _ANEIntermediateTensorClass
}

// GetANEIntermediateTensorClass returns the class object for _ANEIntermediateTensor.
func GetANEIntermediateTensorClass() ANEIntermediateTensorClass {
	return getANEIntermediateTensorClass()
}

type ANEIntermediateTensorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEIntermediateTensorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEIntermediateTensorClass) Alloc() ANEIntermediateTensor {
	rv := objc.SendIfResponds[ANEIntermediateTensor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEIntermediateTensor.EncodeWithCoder]
//   - [ANEIntermediateTensor.TensorIndex]
//   - [ANEIntermediateTensor.TensorName]
//   - [ANEIntermediateTensor.InitWithCoder]
//   - [ANEIntermediateTensor.InitWithNameIndex]
type ANEIntermediateTensor struct {
	objectivec.Object
}

// ANEIntermediateTensorFromID constructs a [ANEIntermediateTensor] from an objc.ID.
func ANEIntermediateTensorFromID(id objc.ID) ANEIntermediateTensor {
	return ANEIntermediateTensor{objectivec.Object{ID: id}}
}

// Ensure ANEIntermediateTensor implements IANEIntermediateTensor.
var _ IANEIntermediateTensor = ANEIntermediateTensor{}

// An interface definition for the [ANEIntermediateTensor] class.
//
// # Methods
//
//   - [IANEIntermediateTensor.EncodeWithCoder]
//   - [IANEIntermediateTensor.TensorIndex]
//   - [IANEIntermediateTensor.TensorName]
//   - [IANEIntermediateTensor.InitWithCoder]
//   - [IANEIntermediateTensor.InitWithNameIndex]
type IANEIntermediateTensor interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	TensorIndex() uint64
	TensorName() string
	InitWithCoder(coder foundation.INSCoder) ANEIntermediateTensor
	InitWithNameIndex(name objectivec.IObject, index uint64) ANEIntermediateTensor
}

// Init initializes the instance.
func (a ANEIntermediateTensor) Init() ANEIntermediateTensor {
	rv := objc.SendIfResponds[ANEIntermediateTensor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEIntermediateTensor) Autorelease() ANEIntermediateTensor {
	rv := objc.SendIfResponds[ANEIntermediateTensor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEIntermediateTensor creates a new ANEIntermediateTensor instance.
func NewANEIntermediateTensor() ANEIntermediateTensor {
	class := getANEIntermediateTensorClass()
	rv := objc.SendIfResponds[ANEIntermediateTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEIntermediateTensorWithCoder(coder objectivec.IObject) ANEIntermediateTensor {
	instance := getANEIntermediateTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEIntermediateTensorFromID(rv)
}

func NewANEIntermediateTensorWithNameIndex(name objectivec.IObject, index uint64) ANEIntermediateTensor {
	instance := getANEIntermediateTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:index:"), name, index)
	return ANEIntermediateTensorFromID(rv)
}

func (a ANEIntermediateTensor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEIntermediateTensor) InitWithCoder(coder foundation.INSCoder) ANEIntermediateTensor {
	rv := objc.SendIfResponds[ANEIntermediateTensor](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEIntermediateTensor) InitWithNameIndex(name objectivec.IObject, index uint64) ANEIntermediateTensor {
	rv := objc.SendIfResponds[ANEIntermediateTensor](a.ID, objc.Sel("initWithName:index:"), name, index)
	return rv
}

func (_ANEIntermediateTensorClass ANEIntermediateTensorClass) DescriptorWithNameIndex(name objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEIntermediateTensorClass.class), objc.Sel("descriptorWithName:index:"), name, index)
	return objectivec.Object{ID: rv}
}
func (_ANEIntermediateTensorClass ANEIntermediateTensorClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEIntermediateTensorClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a ANEIntermediateTensor) TensorIndex() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("tensorIndex"))
	return rv
}
func (a ANEIntermediateTensor) TensorName() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("tensorName"))
	return foundation.NSStringFromID(rv).String()
}
