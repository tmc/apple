// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEModelInstanceParameters] class.
var (
	_ANEModelInstanceParametersClass     ANEModelInstanceParametersClass
	_ANEModelInstanceParametersClassOnce sync.Once
)

func getANEModelInstanceParametersClass() ANEModelInstanceParametersClass {
	_ANEModelInstanceParametersClassOnce.Do(func() {
		_ANEModelInstanceParametersClass = ANEModelInstanceParametersClass{class: objc.GetClass("_ANEModelInstanceParameters")}
	})
	return _ANEModelInstanceParametersClass
}

// GetANEModelInstanceParametersClass returns the class object for _ANEModelInstanceParameters.
func GetANEModelInstanceParametersClass() ANEModelInstanceParametersClass {
	return getANEModelInstanceParametersClass()
}

type ANEModelInstanceParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEModelInstanceParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEModelInstanceParametersClass) Alloc() ANEModelInstanceParameters {
	rv := objc.Send[ANEModelInstanceParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEModelInstanceParameters.EncodeWithCoder]
//   - [ANEModelInstanceParameters.InstanceName]
//   - [ANEModelInstanceParameters.ProcedureArray]
//   - [ANEModelInstanceParameters.InitWithCoder]
//   - [ANEModelInstanceParameters.InitWithProcedureDataProcedureArray]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters
type ANEModelInstanceParameters struct {
	objectivec.Object
}

// ANEModelInstanceParametersFromID constructs a [ANEModelInstanceParameters] from an objc.ID.
func ANEModelInstanceParametersFromID(id objc.ID) ANEModelInstanceParameters {
	return ANEModelInstanceParameters{objectivec.Object{ID: id}}
}
// Ensure ANEModelInstanceParameters implements IANEModelInstanceParameters.
var _ IANEModelInstanceParameters = ANEModelInstanceParameters{}

// An interface definition for the [ANEModelInstanceParameters] class.
//
// # Methods
//
//   - [IANEModelInstanceParameters.EncodeWithCoder]
//   - [IANEModelInstanceParameters.InstanceName]
//   - [IANEModelInstanceParameters.ProcedureArray]
//   - [IANEModelInstanceParameters.InitWithCoder]
//   - [IANEModelInstanceParameters.InitWithProcedureDataProcedureArray]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters
type IANEModelInstanceParameters interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	InstanceName() string
	ProcedureArray() foundation.INSArray
	InitWithCoder(coder foundation.INSCoder) ANEModelInstanceParameters
	InitWithProcedureDataProcedureArray(data objectivec.IObject, array objectivec.IObject) ANEModelInstanceParameters
}

// Init initializes the instance.
func (a ANEModelInstanceParameters) Init() ANEModelInstanceParameters {
	rv := objc.Send[ANEModelInstanceParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEModelInstanceParameters) Autorelease() ANEModelInstanceParameters {
	rv := objc.Send[ANEModelInstanceParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEModelInstanceParameters creates a new ANEModelInstanceParameters instance.
func NewANEModelInstanceParameters() ANEModelInstanceParameters {
	class := getANEModelInstanceParametersClass()
	rv := objc.Send[ANEModelInstanceParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/initWithCoder:
func NewANEModelInstanceParametersWithCoder(coder objectivec.IObject) ANEModelInstanceParameters {
	instance := getANEModelInstanceParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEModelInstanceParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/initWithProcedureData:procedureArray:
func NewANEModelInstanceParametersWithProcedureDataProcedureArray(data objectivec.IObject, array objectivec.IObject) ANEModelInstanceParameters {
	instance := getANEModelInstanceParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcedureData:procedureArray:"), data, array)
	return ANEModelInstanceParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/encodeWithCoder:
func (a ANEModelInstanceParameters) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/initWithCoder:
func (a ANEModelInstanceParameters) InitWithCoder(coder foundation.INSCoder) ANEModelInstanceParameters {
	rv := objc.Send[ANEModelInstanceParameters](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/initWithProcedureData:procedureArray:
func (a ANEModelInstanceParameters) InitWithProcedureDataProcedureArray(data objectivec.IObject, array objectivec.IObject) ANEModelInstanceParameters {
	rv := objc.Send[ANEModelInstanceParameters](a.ID, objc.Sel("initWithProcedureData:procedureArray:"), data, array)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/supportsSecureCoding
func (_ANEModelInstanceParametersClass ANEModelInstanceParametersClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEModelInstanceParametersClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/withProcedureData:procedureArray:
func (_ANEModelInstanceParametersClass ANEModelInstanceParametersClass) WithProcedureDataProcedureArray(data objectivec.IObject, array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelInstanceParametersClass.class), objc.Sel("withProcedureData:procedureArray:"), data, array)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/instanceName
func (a ANEModelInstanceParameters) InstanceName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("instanceName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelInstanceParameters/procedureArray
func (a ANEModelInstanceParameters) ProcedureArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

