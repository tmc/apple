// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEProcedureData] class.
var (
	_ANEProcedureDataClass     ANEProcedureDataClass
	_ANEProcedureDataClassOnce sync.Once
)

func getANEProcedureDataClass() ANEProcedureDataClass {
	_ANEProcedureDataClassOnce.Do(func() {
		_ANEProcedureDataClass = ANEProcedureDataClass{class: objc.GetClass("_ANEProcedureData")}
	})
	return _ANEProcedureDataClass
}

// GetANEProcedureDataClass returns the class object for _ANEProcedureData.
func GetANEProcedureDataClass() ANEProcedureDataClass {
	return getANEProcedureDataClass()
}

type ANEProcedureDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEProcedureDataClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEProcedureDataClass) Alloc() ANEProcedureData {
	rv := objc.Send[ANEProcedureData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEProcedureData.EncodeWithCoder]
//   - [ANEProcedureData.ProcedureSymbol]
//   - [ANEProcedureData.WeightArray]
//   - [ANEProcedureData.InitWithCoder]
//   - [ANEProcedureData.InitWithProcedureWeightArray]
type ANEProcedureData struct {
	objectivec.Object
}

// ANEProcedureDataFromID constructs a [ANEProcedureData] from an objc.ID.
func ANEProcedureDataFromID(id objc.ID) ANEProcedureData {
	return ANEProcedureData{objectivec.Object{ID: id}}
}

// Ensure ANEProcedureData implements IANEProcedureData.
var _ IANEProcedureData = ANEProcedureData{}

// An interface definition for the [ANEProcedureData] class.
//
// # Methods
//
//   - [IANEProcedureData.EncodeWithCoder]
//   - [IANEProcedureData.ProcedureSymbol]
//   - [IANEProcedureData.WeightArray]
//   - [IANEProcedureData.InitWithCoder]
//   - [IANEProcedureData.InitWithProcedureWeightArray]
type IANEProcedureData interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	ProcedureSymbol() string
	WeightArray() foundation.INSArray
	InitWithCoder(coder foundation.INSCoder) ANEProcedureData
	InitWithProcedureWeightArray(procedure objectivec.IObject, array objectivec.IObject) ANEProcedureData
}

// Init initializes the instance.
func (a ANEProcedureData) Init() ANEProcedureData {
	rv := objc.Send[ANEProcedureData](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEProcedureData) Autorelease() ANEProcedureData {
	rv := objc.Send[ANEProcedureData](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEProcedureData creates a new ANEProcedureData instance.
func NewANEProcedureData() ANEProcedureData {
	class := getANEProcedureDataClass()
	rv := objc.Send[ANEProcedureData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEProcedureDataWithCoder(coder objectivec.IObject) ANEProcedureData {
	instance := getANEProcedureDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEProcedureDataFromID(rv)
}

func NewANEProcedureDataWithProcedureWeightArray(procedure objectivec.IObject, array objectivec.IObject) ANEProcedureData {
	instance := getANEProcedureDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcedure:weightArray:"), procedure, array)
	return ANEProcedureDataFromID(rv)
}

func (a ANEProcedureData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEProcedureData) InitWithCoder(coder foundation.INSCoder) ANEProcedureData {
	rv := objc.Send[ANEProcedureData](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEProcedureData) InitWithProcedureWeightArray(procedure objectivec.IObject, array objectivec.IObject) ANEProcedureData {
	rv := objc.Send[ANEProcedureData](a.ID, objc.Sel("initWithProcedure:weightArray:"), procedure, array)
	return rv
}

func (_ANEProcedureDataClass ANEProcedureDataClass) ProcedureDataWithSymbolWeightArray(symbol objectivec.IObject, array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEProcedureDataClass.class), objc.Sel("procedureDataWithSymbol:weightArray:"), symbol, array)
	return objectivec.Object{ID: rv}
}
func (_ANEProcedureDataClass ANEProcedureDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEProcedureDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a ANEProcedureData) ProcedureSymbol() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureSymbol"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEProcedureData) WeightArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
