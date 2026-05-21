// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderBinaryRange] class.
var (
	_GTAGX2ShaderBinaryRangeClass     GTAGX2ShaderBinaryRangeClass
	_GTAGX2ShaderBinaryRangeClassOnce sync.Once
)

func getGTAGX2ShaderBinaryRangeClass() GTAGX2ShaderBinaryRangeClass {
	_GTAGX2ShaderBinaryRangeClassOnce.Do(func() {
		_GTAGX2ShaderBinaryRangeClass = GTAGX2ShaderBinaryRangeClass{class: objc.GetClass("GTAGX2ShaderBinaryRange")}
	})
	return _GTAGX2ShaderBinaryRangeClass
}

// GetGTAGX2ShaderBinaryRangeClass returns the class object for GTAGX2ShaderBinaryRange.
func GetGTAGX2ShaderBinaryRangeClass() GTAGX2ShaderBinaryRangeClass {
	return getGTAGX2ShaderBinaryRangeClass()
}

type GTAGX2ShaderBinaryRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderBinaryRangeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderBinaryRangeClass) Alloc() GTAGX2ShaderBinaryRange {
	rv := objc.Send[GTAGX2ShaderBinaryRange](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderBinaryRange.AddrEnd]
//   - [GTAGX2ShaderBinaryRange.SetAddrEnd]
//   - [GTAGX2ShaderBinaryRange.AddrStart]
//   - [GTAGX2ShaderBinaryRange.SetAddrStart]
//   - [GTAGX2ShaderBinaryRange.Binary]
//   - [GTAGX2ShaderBinaryRange.CallStack]
//   - [GTAGX2ShaderBinaryRange.CostForAddress]
//   - [GTAGX2ShaderBinaryRange.EncodeWithCoder]
//   - [GTAGX2ShaderBinaryRange.Location]
//   - [GTAGX2ShaderBinaryRange.NumSamples]
//   - [GTAGX2ShaderBinaryRange.SetNumSamples]
//   - [GTAGX2ShaderBinaryRange.SetParent]
//   - [GTAGX2ShaderBinaryRange.TotalCost]
//   - [GTAGX2ShaderBinaryRange.InitWithCoder]
//   - [GTAGX2ShaderBinaryRange.InitWithLocationAddrStartAddrEndParent]
//   - [GTAGX2ShaderBinaryRange.DebugDescription]
//   - [GTAGX2ShaderBinaryRange.Description]
//   - [GTAGX2ShaderBinaryRange.Hash]
//   - [GTAGX2ShaderBinaryRange.Superclass]
type GTAGX2ShaderBinaryRange struct {
	objectivec.Object
}

// GTAGX2ShaderBinaryRangeFromID constructs a [GTAGX2ShaderBinaryRange] from an objc.ID.
func GTAGX2ShaderBinaryRangeFromID(id objc.ID) GTAGX2ShaderBinaryRange {
	return GTAGX2ShaderBinaryRange{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderBinaryRange implements IGTAGX2ShaderBinaryRange.
var _ IGTAGX2ShaderBinaryRange = GTAGX2ShaderBinaryRange{}

// An interface definition for the [GTAGX2ShaderBinaryRange] class.
//
// # Methods
//
//   - [IGTAGX2ShaderBinaryRange.AddrEnd]
//   - [IGTAGX2ShaderBinaryRange.SetAddrEnd]
//   - [IGTAGX2ShaderBinaryRange.AddrStart]
//   - [IGTAGX2ShaderBinaryRange.SetAddrStart]
//   - [IGTAGX2ShaderBinaryRange.Binary]
//   - [IGTAGX2ShaderBinaryRange.CallStack]
//   - [IGTAGX2ShaderBinaryRange.CostForAddress]
//   - [IGTAGX2ShaderBinaryRange.EncodeWithCoder]
//   - [IGTAGX2ShaderBinaryRange.Location]
//   - [IGTAGX2ShaderBinaryRange.NumSamples]
//   - [IGTAGX2ShaderBinaryRange.SetNumSamples]
//   - [IGTAGX2ShaderBinaryRange.SetParent]
//   - [IGTAGX2ShaderBinaryRange.TotalCost]
//   - [IGTAGX2ShaderBinaryRange.InitWithCoder]
//   - [IGTAGX2ShaderBinaryRange.InitWithLocationAddrStartAddrEndParent]
//   - [IGTAGX2ShaderBinaryRange.DebugDescription]
//   - [IGTAGX2ShaderBinaryRange.Description]
//   - [IGTAGX2ShaderBinaryRange.Hash]
//   - [IGTAGX2ShaderBinaryRange.Superclass]
type IGTAGX2ShaderBinaryRange interface {
	objectivec.IObject

	// Topic: Methods

	AddrEnd() uint32
	SetAddrEnd(value uint32)
	AddrStart() uint32
	SetAddrStart(value uint32)
	Binary() unsafe.Pointer
	CallStack() foundation.INSArray
	CostForAddress(address uint32) float64
	EncodeWithCoder(coder foundation.INSCoder)
	Location() unsafe.Pointer
	NumSamples() uint64
	SetNumSamples(value uint64)
	SetParent(parent objectivec.IObject)
	TotalCost() float64
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinaryRange
	InitWithLocationAddrStartAddrEndParent(location objectivec.IObject, start uint32, end uint32, parent objectivec.IObject) GTAGX2ShaderBinaryRange
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderBinaryRange) Init() GTAGX2ShaderBinaryRange {
	rv := objc.Send[GTAGX2ShaderBinaryRange](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderBinaryRange) Autorelease() GTAGX2ShaderBinaryRange {
	rv := objc.Send[GTAGX2ShaderBinaryRange](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderBinaryRange creates a new GTAGX2ShaderBinaryRange instance.
func NewGTAGX2ShaderBinaryRange() GTAGX2ShaderBinaryRange {
	class := getGTAGX2ShaderBinaryRangeClass()
	rv := objc.Send[GTAGX2ShaderBinaryRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTAGX2ShaderBinaryRangeWithCoder(coder objectivec.IObject) GTAGX2ShaderBinaryRange {
	instance := getGTAGX2ShaderBinaryRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderBinaryRangeFromID(rv)
}

func NewGTAGX2ShaderBinaryRangeWithLocationAddrStartAddrEndParent(location objectivec.IObject, start uint32, end uint32, parent objectivec.IObject) GTAGX2ShaderBinaryRange {
	instance := getGTAGX2ShaderBinaryRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:addrStart:addrEnd:parent:"), location, start, end, parent)
	return GTAGX2ShaderBinaryRangeFromID(rv)
}

func (g GTAGX2ShaderBinaryRange) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("costForAddress:"), address)
	return rv
}
func (g GTAGX2ShaderBinaryRange) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTAGX2ShaderBinaryRange) SetParent(parent objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("setParent:"), parent)
}
func (g GTAGX2ShaderBinaryRange) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinaryRange {
	rv := objc.Send[GTAGX2ShaderBinaryRange](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GTAGX2ShaderBinaryRange) InitWithLocationAddrStartAddrEndParent(location objectivec.IObject, start uint32, end uint32, parent objectivec.IObject) GTAGX2ShaderBinaryRange {
	rv := objc.Send[GTAGX2ShaderBinaryRange](g.ID, objc.Sel("initWithLocation:addrStart:addrEnd:parent:"), location, start, end, parent)
	return rv
}

func (_GTAGX2ShaderBinaryRangeClass GTAGX2ShaderBinaryRangeClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderBinaryRangeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTAGX2ShaderBinaryRange) AddrEnd() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("addrEnd"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) SetAddrEnd(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAddrEnd:"), value)
}
func (g GTAGX2ShaderBinaryRange) AddrStart() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("addrStart"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) SetAddrStart(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAddrStart:"), value)
}
func (g GTAGX2ShaderBinaryRange) Binary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("binary"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) CallStack() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("callStack"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTAGX2ShaderBinaryRange) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryRange) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryRange) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("location"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) NumSamples() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numSamples"))
	return rv
}
func (g GTAGX2ShaderBinaryRange) SetNumSamples(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setNumSamples:"), value)
}
func (g GTAGX2ShaderBinaryRange) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTAGX2ShaderBinaryRange) TotalCost() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCost"))
	return rv
}
