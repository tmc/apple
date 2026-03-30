// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoConvolutionWeightsForMPS] class.
var (
	_EspressoConvolutionWeightsForMPSClass     EspressoConvolutionWeightsForMPSClass
	_EspressoConvolutionWeightsForMPSClassOnce sync.Once
)

func getEspressoConvolutionWeightsForMPSClass() EspressoConvolutionWeightsForMPSClass {
	_EspressoConvolutionWeightsForMPSClassOnce.Do(func() {
		_EspressoConvolutionWeightsForMPSClass = EspressoConvolutionWeightsForMPSClass{class: objc.GetClass("EspressoConvolutionWeightsForMPS")}
	})
	return _EspressoConvolutionWeightsForMPSClass
}

// GetEspressoConvolutionWeightsForMPSClass returns the class object for EspressoConvolutionWeightsForMPS.
func GetEspressoConvolutionWeightsForMPSClass() EspressoConvolutionWeightsForMPSClass {
	return getEspressoConvolutionWeightsForMPSClass()
}

type EspressoConvolutionWeightsForMPSClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoConvolutionWeightsForMPSClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoConvolutionWeightsForMPSClass) Alloc() EspressoConvolutionWeightsForMPS {
	rv := objc.Send[EspressoConvolutionWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoConvolutionWeightsForMPS.BiasTerms]
//   - [EspressoConvolutionWeightsForMPS.DataType]
//   - [EspressoConvolutionWeightsForMPS.Descriptor]
//   - [EspressoConvolutionWeightsForMPS.Label]
//   - [EspressoConvolutionWeightsForMPS.LookupTableForUInt8Kernel]
//   - [EspressoConvolutionWeightsForMPS.Purge]
//   - [EspressoConvolutionWeightsForMPS.RangesForUInt8Kernel]
//   - [EspressoConvolutionWeightsForMPS.Ready]
//   - [EspressoConvolutionWeightsForMPS.Weights]
//   - [EspressoConvolutionWeightsForMPS.InitWithParams]
//   - [EspressoConvolutionWeightsForMPS.DebugDescription]
//   - [EspressoConvolutionWeightsForMPS.Description]
//   - [EspressoConvolutionWeightsForMPS.Hash]
//   - [EspressoConvolutionWeightsForMPS.Superclass]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS
type EspressoConvolutionWeightsForMPS struct {
	objectivec.Object
}

// EspressoConvolutionWeightsForMPSFromID constructs a [EspressoConvolutionWeightsForMPS] from an objc.ID.
func EspressoConvolutionWeightsForMPSFromID(id objc.ID) EspressoConvolutionWeightsForMPS {
	return EspressoConvolutionWeightsForMPS{objectivec.Object{ID: id}}
}

// Ensure EspressoConvolutionWeightsForMPS implements IEspressoConvolutionWeightsForMPS.
var _ IEspressoConvolutionWeightsForMPS = EspressoConvolutionWeightsForMPS{}

// An interface definition for the [EspressoConvolutionWeightsForMPS] class.
//
// # Methods
//
//   - [IEspressoConvolutionWeightsForMPS.BiasTerms]
//   - [IEspressoConvolutionWeightsForMPS.DataType]
//   - [IEspressoConvolutionWeightsForMPS.Descriptor]
//   - [IEspressoConvolutionWeightsForMPS.Label]
//   - [IEspressoConvolutionWeightsForMPS.LookupTableForUInt8Kernel]
//   - [IEspressoConvolutionWeightsForMPS.Purge]
//   - [IEspressoConvolutionWeightsForMPS.RangesForUInt8Kernel]
//   - [IEspressoConvolutionWeightsForMPS.Ready]
//   - [IEspressoConvolutionWeightsForMPS.Weights]
//   - [IEspressoConvolutionWeightsForMPS.InitWithParams]
//   - [IEspressoConvolutionWeightsForMPS.DebugDescription]
//   - [IEspressoConvolutionWeightsForMPS.Description]
//   - [IEspressoConvolutionWeightsForMPS.Hash]
//   - [IEspressoConvolutionWeightsForMPS.Superclass]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS
type IEspressoConvolutionWeightsForMPS interface {
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
	InitWithParams(params objectivec.IObject) EspressoConvolutionWeightsForMPS
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e EspressoConvolutionWeightsForMPS) Init() EspressoConvolutionWeightsForMPS {
	rv := objc.Send[EspressoConvolutionWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoConvolutionWeightsForMPS) Autorelease() EspressoConvolutionWeightsForMPS {
	rv := objc.Send[EspressoConvolutionWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoConvolutionWeightsForMPS creates a new EspressoConvolutionWeightsForMPS instance.
func NewEspressoConvolutionWeightsForMPS() EspressoConvolutionWeightsForMPS {
	class := getEspressoConvolutionWeightsForMPSClass()
	rv := objc.Send[EspressoConvolutionWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/initWithParams:
func NewEspressoConvolutionWeightsForMPSWithParams(params objectivec.IObject) EspressoConvolutionWeightsForMPS {
	instance := getEspressoConvolutionWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoConvolutionWeightsForMPSFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/biasTerms
func (e EspressoConvolutionWeightsForMPS) BiasTerms() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("biasTerms"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/dataType
func (e EspressoConvolutionWeightsForMPS) DataType() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/descriptor
func (e EspressoConvolutionWeightsForMPS) Descriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/label
func (e EspressoConvolutionWeightsForMPS) Label() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/lookupTableForUInt8Kernel
func (e EspressoConvolutionWeightsForMPS) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/purge
func (e EspressoConvolutionWeightsForMPS) Purge() {
	objc.Send[objc.ID](e.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/rangesForUInt8Kernel
func (e EspressoConvolutionWeightsForMPS) RangesForUInt8Kernel() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("rangesForUInt8Kernel"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/ready
func (e EspressoConvolutionWeightsForMPS) Ready() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("ready"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/weights
func (e EspressoConvolutionWeightsForMPS) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("weights"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/initWithParams:
func (e EspressoConvolutionWeightsForMPS) InitWithParams(params objectivec.IObject) EspressoConvolutionWeightsForMPS {
	rv := objc.Send[EspressoConvolutionWeightsForMPS](e.ID, objc.Sel("initWithParams:"), params)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/debugDescription
func (e EspressoConvolutionWeightsForMPS) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/description
func (e EspressoConvolutionWeightsForMPS) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/hash
func (e EspressoConvolutionWeightsForMPS) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/superclass
func (e EspressoConvolutionWeightsForMPS) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
