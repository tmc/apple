// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// The class instance for the [EspressoInnerProductWeightsForMPS] class.
var (
	_EspressoInnerProductWeightsForMPSClass     EspressoInnerProductWeightsForMPSClass
	_EspressoInnerProductWeightsForMPSClassOnce sync.Once
)

func getEspressoInnerProductWeightsForMPSClass() EspressoInnerProductWeightsForMPSClass {
	_EspressoInnerProductWeightsForMPSClassOnce.Do(func() {
		_EspressoInnerProductWeightsForMPSClass = EspressoInnerProductWeightsForMPSClass{class: objc.GetClass("EspressoInnerProductWeightsForMPS")}
	})
	return _EspressoInnerProductWeightsForMPSClass
}

// GetEspressoInnerProductWeightsForMPSClass returns the class object for EspressoInnerProductWeightsForMPS.
func GetEspressoInnerProductWeightsForMPSClass() EspressoInnerProductWeightsForMPSClass {
	return getEspressoInnerProductWeightsForMPSClass()
}

type EspressoInnerProductWeightsForMPSClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoInnerProductWeightsForMPSClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoInnerProductWeightsForMPSClass) Alloc() EspressoInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoInnerProductWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoInnerProductWeightsForMPS.BiasTerms]
//   - [EspressoInnerProductWeightsForMPS.DataType]
//   - [EspressoInnerProductWeightsForMPS.Descriptor]
//   - [EspressoInnerProductWeightsForMPS.Label]
//   - [EspressoInnerProductWeightsForMPS.LookupTableForUInt8Kernel]
//   - [EspressoInnerProductWeightsForMPS.Purge]
//   - [EspressoInnerProductWeightsForMPS.RangesForUInt8Kernel]
//   - [EspressoInnerProductWeightsForMPS.Ready]
//   - [EspressoInnerProductWeightsForMPS.Weights]
//   - [EspressoInnerProductWeightsForMPS.InitWithParams]
//   - [EspressoInnerProductWeightsForMPS.DebugDescription]
//   - [EspressoInnerProductWeightsForMPS.Description]
//   - [EspressoInnerProductWeightsForMPS.Hash]
//   - [EspressoInnerProductWeightsForMPS.Superclass]
type EspressoInnerProductWeightsForMPS struct {
	objectivec.Object
}

// EspressoInnerProductWeightsForMPSFromID constructs a [EspressoInnerProductWeightsForMPS] from an objc.ID.
func EspressoInnerProductWeightsForMPSFromID(id objc.ID) EspressoInnerProductWeightsForMPS {
	return EspressoInnerProductWeightsForMPS{objectivec.Object{ID: id}}
}

// Ensure EspressoInnerProductWeightsForMPS implements IEspressoInnerProductWeightsForMPS.
var _ IEspressoInnerProductWeightsForMPS = EspressoInnerProductWeightsForMPS{}

// An interface definition for the [EspressoInnerProductWeightsForMPS] class.
//
// # Methods
//
//   - [IEspressoInnerProductWeightsForMPS.BiasTerms]
//   - [IEspressoInnerProductWeightsForMPS.DataType]
//   - [IEspressoInnerProductWeightsForMPS.Descriptor]
//   - [IEspressoInnerProductWeightsForMPS.Label]
//   - [IEspressoInnerProductWeightsForMPS.LookupTableForUInt8Kernel]
//   - [IEspressoInnerProductWeightsForMPS.Purge]
//   - [IEspressoInnerProductWeightsForMPS.RangesForUInt8Kernel]
//   - [IEspressoInnerProductWeightsForMPS.Ready]
//   - [IEspressoInnerProductWeightsForMPS.Weights]
//   - [IEspressoInnerProductWeightsForMPS.InitWithParams]
//   - [IEspressoInnerProductWeightsForMPS.DebugDescription]
//   - [IEspressoInnerProductWeightsForMPS.Description]
//   - [IEspressoInnerProductWeightsForMPS.Hash]
//   - [IEspressoInnerProductWeightsForMPS.Superclass]
type IEspressoInnerProductWeightsForMPS interface {
	objectivec.IObject

	// Topic: Methods

	BiasTerms() unsafe.Pointer
	DataType() uint32
	Descriptor() objectivec.IObject
	Label() objectivec.IObject
	LookupTableForUInt8Kernel() unsafe.Pointer
	Purge()
	RangesForUInt8Kernel() unsafe.Pointer
	Ready() bool
	Weights() unsafe.Pointer
	InitWithParams(params appleneuralengine.InnerProductUniforms) EspressoInnerProductWeightsForMPS
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (e EspressoInnerProductWeightsForMPS) Init() EspressoInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoInnerProductWeightsForMPS) Autorelease() EspressoInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoInnerProductWeightsForMPS creates a new EspressoInnerProductWeightsForMPS instance.
func NewEspressoInnerProductWeightsForMPS() EspressoInnerProductWeightsForMPS {
	class := getEspressoInnerProductWeightsForMPSClass()
	rv := objc.SendIfResponds[EspressoInnerProductWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoInnerProductWeightsForMPSWithParams(params appleneuralengine.InnerProductUniforms) EspressoInnerProductWeightsForMPS {
	instance := getEspressoInnerProductWeightsForMPSClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoInnerProductWeightsForMPSFromID(rv)
}

func (e EspressoInnerProductWeightsForMPS) BiasTerms() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("biasTerms"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) DataType() uint32 {
	rv := objc.SendIfResponds[uint32](e.ID, objc.Sel("dataType"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) Descriptor() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}
func (e EspressoInnerProductWeightsForMPS) Label() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}
func (e EspressoInnerProductWeightsForMPS) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) Purge() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("purge"))
}
func (e EspressoInnerProductWeightsForMPS) RangesForUInt8Kernel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("rangesForUInt8Kernel"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) Ready() bool {
	rv := objc.SendIfResponds[bool](e.ID, objc.Sel("ready"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) Weights() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("weights"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) InitWithParams(params appleneuralengine.InnerProductUniforms) EspressoInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("initWithParams:"), params)
	return rv
}

func (e EspressoInnerProductWeightsForMPS) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoInnerProductWeightsForMPS) Description() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoInnerProductWeightsForMPS) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("hash"))
	return rv
}
func (e EspressoInnerProductWeightsForMPS) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](e.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
