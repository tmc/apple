// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (ec EspressoInnerProductWeightsForMPSClass) Alloc() EspressoInnerProductWeightsForMPS {
	rv := objc.Send[EspressoInnerProductWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS
type EspressoInnerProductWeightsForMPS struct {
	objectivec.Object
}

// EspressoInnerProductWeightsForMPSFromID constructs a [EspressoInnerProductWeightsForMPS] from an objc.ID.
func EspressoInnerProductWeightsForMPSFromID(id objc.ID) EspressoInnerProductWeightsForMPS {
	return EspressoInnerProductWeightsForMPS{objectivec.Object{id}}
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
//
// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS
type IEspressoInnerProductWeightsForMPS interface {
	objectivec.IObject

	// Topic: Methods

	BiasTerms() unsafe.Pointer
	DataType() uint32
	Descriptor() objectivec.IObject
	Label() objectivec.IObject
	LookupTableForUInt8Kernel() unsafe.Pointer
	Purge()
	RangesForUInt8Kernel() []objectivec.IObject
	Ready() bool
	Weights() unsafe.Pointer
	InitWithParams(params objectivec.IObject) EspressoInnerProductWeightsForMPS
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e EspressoInnerProductWeightsForMPS) Init() EspressoInnerProductWeightsForMPS {
	rv := objc.Send[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoInnerProductWeightsForMPS) Autorelease() EspressoInnerProductWeightsForMPS {
	rv := objc.Send[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoInnerProductWeightsForMPS creates a new EspressoInnerProductWeightsForMPS instance.
func NewEspressoInnerProductWeightsForMPS() EspressoInnerProductWeightsForMPS {
	class := getEspressoInnerProductWeightsForMPSClass()
	rv := objc.Send[EspressoInnerProductWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/initWithParams:
func NewEspressoInnerProductWeightsForMPSWithParams(params objectivec.IObject) EspressoInnerProductWeightsForMPS {
	instance := getEspressoInnerProductWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoInnerProductWeightsForMPSFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/biasTerms
func (e EspressoInnerProductWeightsForMPS) BiasTerms() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("biasTerms"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/dataType
func (e EspressoInnerProductWeightsForMPS) DataType() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/descriptor
func (e EspressoInnerProductWeightsForMPS) Descriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/label
func (e EspressoInnerProductWeightsForMPS) Label() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/lookupTableForUInt8Kernel
func (e EspressoInnerProductWeightsForMPS) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/purge
func (e EspressoInnerProductWeightsForMPS) Purge() {
	objc.Send[objc.ID](e.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/rangesForUInt8Kernel
func (e EspressoInnerProductWeightsForMPS) RangesForUInt8Kernel() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("rangesForUInt8Kernel"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/ready
func (e EspressoInnerProductWeightsForMPS) Ready() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("ready"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/weights
func (e EspressoInnerProductWeightsForMPS) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("weights"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/initWithParams:
func (e EspressoInnerProductWeightsForMPS) InitWithParams(params objectivec.IObject) EspressoInnerProductWeightsForMPS {
	rv := objc.Send[EspressoInnerProductWeightsForMPS](e.ID, objc.Sel("initWithParams:"), params)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/debugDescription
func (e EspressoInnerProductWeightsForMPS) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/description
func (e EspressoInnerProductWeightsForMPS) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/hash
func (e EspressoInnerProductWeightsForMPS) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/superclass
func (e EspressoInnerProductWeightsForMPS) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

