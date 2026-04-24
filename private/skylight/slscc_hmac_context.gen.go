// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSCCHmacContext] class.
var (
	_SLSCCHmacContextClass     SLSCCHmacContextClass
	_SLSCCHmacContextClassOnce sync.Once
)

func getSLSCCHmacContextClass() SLSCCHmacContextClass {
	_SLSCCHmacContextClassOnce.Do(func() {
		_SLSCCHmacContextClass = SLSCCHmacContextClass{class: objc.GetClass("SLSCCHmacContext")}
	})
	return _SLSCCHmacContextClass
}

// GetSLSCCHmacContextClass returns the class object for SLSCCHmacContext.
func GetSLSCCHmacContextClass() SLSCCHmacContextClass {
	return getSLSCCHmacContextClass()
}

type SLSCCHmacContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSCCHmacContextClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSCCHmacContextClass) Alloc() SLSCCHmacContext {
	rv := objc.Send[SLSCCHmacContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSCCHmacContext.FinalizedData]
//   - [SLSCCHmacContext.UpdateSigningContextWithBytesLength]
//   - [SLSCCHmacContext.UpdateSigningContextWithData]
//   - [SLSCCHmacContext.UpdateSigningContextWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext
type SLSCCHmacContext struct {
	objectivec.Object
}

// SLSCCHmacContextFromID constructs a [SLSCCHmacContext] from an objc.ID.
func SLSCCHmacContextFromID(id objc.ID) SLSCCHmacContext {
	return SLSCCHmacContext{objectivec.Object{ID: id}}
}

// Ensure SLSCCHmacContext implements ISLSCCHmacContext.
var _ ISLSCCHmacContext = SLSCCHmacContext{}

// An interface definition for the [SLSCCHmacContext] class.
//
// # Methods
//
//   - [ISLSCCHmacContext.FinalizedData]
//   - [ISLSCCHmacContext.UpdateSigningContextWithBytesLength]
//   - [ISLSCCHmacContext.UpdateSigningContextWithData]
//   - [ISLSCCHmacContext.UpdateSigningContextWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext
type ISLSCCHmacContext interface {
	objectivec.IObject

	// Topic: Methods

	FinalizedData() foundation.INSData
	UpdateSigningContextWithBytesLength(bytes []byte)
	UpdateSigningContextWithData(data objectivec.IObject)
	UpdateSigningContextWithObject(object objectivec.IObject)
}

// Init initializes the instance.
func (s SLSCCHmacContext) Init() SLSCCHmacContext {
	rv := objc.Send[SLSCCHmacContext](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSCCHmacContext) Autorelease() SLSCCHmacContext {
	rv := objc.Send[SLSCCHmacContext](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSCCHmacContext creates a new SLSCCHmacContext instance.
func NewSLSCCHmacContext() SLSCCHmacContext {
	class := getSLSCCHmacContextClass()
	rv := objc.Send[SLSCCHmacContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext/updateSigningContextWithBytes:length:
func (s SLSCCHmacContext) UpdateSigningContextWithBytesLength(bytes []byte) {
	objc.Send[objc.ID](s.ID, objc.Sel("updateSigningContextWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext/updateSigningContextWithData:
func (s SLSCCHmacContext) UpdateSigningContextWithData(data objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("updateSigningContextWithData:"), data)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext/updateSigningContextWithObject:
func (s SLSCCHmacContext) UpdateSigningContextWithObject(object objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("updateSigningContextWithObject:"), object)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext/contextWithImplementingDigester:
func (_SLSCCHmacContextClass SLSCCHmacContextClass) ContextWithImplementingDigester(digester objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSCCHmacContextClass.class), objc.Sel("contextWithImplementingDigester:"), digester)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSCCHmacContext/finalizedData
func (s SLSCCHmacContext) FinalizedData() foundation.INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("finalizedData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
