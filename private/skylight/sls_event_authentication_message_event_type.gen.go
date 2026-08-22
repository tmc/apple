// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSEventAuthenticationMessageEventType] class.
var (
	_SLSEventAuthenticationMessageEventTypeClass     SLSEventAuthenticationMessageEventTypeClass
	_SLSEventAuthenticationMessageEventTypeClassOnce sync.Once
)

func getSLSEventAuthenticationMessageEventTypeClass() SLSEventAuthenticationMessageEventTypeClass {
	_SLSEventAuthenticationMessageEventTypeClassOnce.Do(func() {
		_SLSEventAuthenticationMessageEventTypeClass = SLSEventAuthenticationMessageEventTypeClass{class: objc.GetClass("SLSEventAuthenticationMessageEventType")}
	})
	return _SLSEventAuthenticationMessageEventTypeClass
}

// GetSLSEventAuthenticationMessageEventTypeClass returns the class object for SLSEventAuthenticationMessageEventType.
func GetSLSEventAuthenticationMessageEventTypeClass() SLSEventAuthenticationMessageEventTypeClass {
	return getSLSEventAuthenticationMessageEventTypeClass()
}

type SLSEventAuthenticationMessageEventTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSEventAuthenticationMessageEventTypeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSEventAuthenticationMessageEventTypeClass) Alloc() SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSEventAuthenticationMessageEventType.AddToSigningContext]
//   - [SLSEventAuthenticationMessageEventType.CgSubType]
//   - [SLSEventAuthenticationMessageEventType.CgType]
//   - [SLSEventAuthenticationMessageEventType.EncodeWithCoder]
//   - [SLSEventAuthenticationMessageEventType.HidType]
//   - [SLSEventAuthenticationMessageEventType.IsCGType]
//   - [SLSEventAuthenticationMessageEventType.InitWithCoder]
//   - [SLSEventAuthenticationMessageEventType.InitWithEventRecord]
//   - [SLSEventAuthenticationMessageEventType.InitWithHIDTypeCgTypeCgSubType]
type SLSEventAuthenticationMessageEventType struct {
	objectivec.Object
}

// SLSEventAuthenticationMessageEventTypeFromID constructs a [SLSEventAuthenticationMessageEventType] from an objc.ID.
func SLSEventAuthenticationMessageEventTypeFromID(id objc.ID) SLSEventAuthenticationMessageEventType {
	return SLSEventAuthenticationMessageEventType{objectivec.Object{ID: id}}
}

// Ensure SLSEventAuthenticationMessageEventType implements ISLSEventAuthenticationMessageEventType.
var _ ISLSEventAuthenticationMessageEventType = SLSEventAuthenticationMessageEventType{}

// An interface definition for the [SLSEventAuthenticationMessageEventType] class.
//
// # Methods
//
//   - [ISLSEventAuthenticationMessageEventType.AddToSigningContext]
//   - [ISLSEventAuthenticationMessageEventType.CgSubType]
//   - [ISLSEventAuthenticationMessageEventType.CgType]
//   - [ISLSEventAuthenticationMessageEventType.EncodeWithCoder]
//   - [ISLSEventAuthenticationMessageEventType.HidType]
//   - [ISLSEventAuthenticationMessageEventType.IsCGType]
//   - [ISLSEventAuthenticationMessageEventType.InitWithCoder]
//   - [ISLSEventAuthenticationMessageEventType.InitWithEventRecord]
//   - [ISLSEventAuthenticationMessageEventType.InitWithHIDTypeCgTypeCgSubType]
type ISLSEventAuthenticationMessageEventType interface {
	objectivec.IObject

	// Topic: Methods

	AddToSigningContext(context objectivec.IObject)
	CgSubType() uint64
	CgType() uint32
	EncodeWithCoder(coder foundation.INSCoder)
	HidType() uint32
	IsCGType() bool
	InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessageEventType
	InitWithEventRecord(record *SLSEventRecord) SLSEventAuthenticationMessageEventType
	InitWithHIDTypeCgTypeCgSubType(hIDType uint32, type_ uint32, type_2 uint64) SLSEventAuthenticationMessageEventType
}

// Init initializes the instance.
func (s SLSEventAuthenticationMessageEventType) Init() SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSEventAuthenticationMessageEventType) Autorelease() SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSEventAuthenticationMessageEventType creates a new SLSEventAuthenticationMessageEventType instance.
func NewSLSEventAuthenticationMessageEventType() SLSEventAuthenticationMessageEventType {
	class := getSLSEventAuthenticationMessageEventTypeClass()
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSEventAuthenticationMessageEventTypeWithCoder(coder objectivec.IObject) SLSEventAuthenticationMessageEventType {
	instance := getSLSEventAuthenticationMessageEventTypeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSEventAuthenticationMessageEventTypeFromID(rv)
}

func NewSLSEventAuthenticationMessageEventTypeWithEventRecord(record *SLSEventRecord) SLSEventAuthenticationMessageEventType {
	instance := getSLSEventAuthenticationMessageEventTypeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithEventRecord:"), unsafe.Pointer(record))
	return SLSEventAuthenticationMessageEventTypeFromID(rv)
}

func NewSLSEventAuthenticationMessageEventTypeWithHIDTypeCgTypeCgSubType(hIDType uint32, type_ uint32, type_2 uint64) SLSEventAuthenticationMessageEventType {
	instance := getSLSEventAuthenticationMessageEventTypeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHIDType:cgType:cgSubType:"), hIDType, type_, type_2)
	return SLSEventAuthenticationMessageEventTypeFromID(rv)
}

func (s SLSEventAuthenticationMessageEventType) AddToSigningContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("addToSigningContext:"), context)
}
func (s SLSEventAuthenticationMessageEventType) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (s SLSEventAuthenticationMessageEventType) InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s SLSEventAuthenticationMessageEventType) InitWithEventRecord(record *SLSEventRecord) SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](s.ID, objc.Sel("initWithEventRecord:"), unsafe.Pointer(record))
	return rv
}
func (s SLSEventAuthenticationMessageEventType) InitWithHIDTypeCgTypeCgSubType(hIDType uint32, type_ uint32, type_2 uint64) SLSEventAuthenticationMessageEventType {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageEventType](s.ID, objc.Sel("initWithHIDType:cgType:cgSubType:"), hIDType, type_, type_2)
	return rv
}

func (_SLSEventAuthenticationMessageEventTypeClass SLSEventAuthenticationMessageEventTypeClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_SLSEventAuthenticationMessageEventTypeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (s SLSEventAuthenticationMessageEventType) CgSubType() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("cgSubType"))
	return rv
}
func (s SLSEventAuthenticationMessageEventType) CgType() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("cgType"))
	return rv
}
func (s SLSEventAuthenticationMessageEventType) HidType() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("hidType"))
	return rv
}
func (s SLSEventAuthenticationMessageEventType) IsCGType() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isCGType"))
	return rv
}
