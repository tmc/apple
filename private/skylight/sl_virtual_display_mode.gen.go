// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLVirtualDisplayMode] class.
var (
	_SLVirtualDisplayModeClass     SLVirtualDisplayModeClass
	_SLVirtualDisplayModeClassOnce sync.Once
)

func getSLVirtualDisplayModeClass() SLVirtualDisplayModeClass {
	_SLVirtualDisplayModeClassOnce.Do(func() {
		_SLVirtualDisplayModeClass = SLVirtualDisplayModeClass{class: objc.GetClass("SLVirtualDisplayMode")}
	})
	return _SLVirtualDisplayModeClass
}

// GetSLVirtualDisplayModeClass returns the class object for SLVirtualDisplayMode.
func GetSLVirtualDisplayModeClass() SLVirtualDisplayModeClass {
	return getSLVirtualDisplayModeClass()
}

type SLVirtualDisplayModeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLVirtualDisplayModeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLVirtualDisplayModeClass) Alloc() SLVirtualDisplayMode {
	rv := objc.SendIfResponds[SLVirtualDisplayMode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLVirtualDisplayMode.DictionaryRepresentation]
//   - [SLVirtualDisplayMode.Eotf]
//   - [SLVirtualDisplayMode.SetEotf]
//   - [SLVirtualDisplayMode.IsEqualToMode]
//   - [SLVirtualDisplayMode.Options]
//   - [SLVirtualDisplayMode.SetOptions]
//   - [SLVirtualDisplayMode.RefreshDeadline]
//   - [SLVirtualDisplayMode.SetRefreshDeadline]
//   - [SLVirtualDisplayMode.RefreshRate]
//   - [SLVirtualDisplayMode.SizeInPixels]
//   - [SLVirtualDisplayMode.SizeInPoints]
//   - [SLVirtualDisplayMode.InitWithSizeInPixelsSizeInPointsRefreshRateError]
type SLVirtualDisplayMode struct {
	objectivec.Object
}

// SLVirtualDisplayModeFromID constructs a [SLVirtualDisplayMode] from an objc.ID.
func SLVirtualDisplayModeFromID(id objc.ID) SLVirtualDisplayMode {
	return SLVirtualDisplayMode{objectivec.Object{ID: id}}
}

// Ensure SLVirtualDisplayMode implements ISLVirtualDisplayMode.
var _ ISLVirtualDisplayMode = SLVirtualDisplayMode{}

// An interface definition for the [SLVirtualDisplayMode] class.
//
// # Methods
//
//   - [ISLVirtualDisplayMode.DictionaryRepresentation]
//   - [ISLVirtualDisplayMode.Eotf]
//   - [ISLVirtualDisplayMode.SetEotf]
//   - [ISLVirtualDisplayMode.IsEqualToMode]
//   - [ISLVirtualDisplayMode.Options]
//   - [ISLVirtualDisplayMode.SetOptions]
//   - [ISLVirtualDisplayMode.RefreshDeadline]
//   - [ISLVirtualDisplayMode.SetRefreshDeadline]
//   - [ISLVirtualDisplayMode.RefreshRate]
//   - [ISLVirtualDisplayMode.SizeInPixels]
//   - [ISLVirtualDisplayMode.SizeInPoints]
//   - [ISLVirtualDisplayMode.InitWithSizeInPixelsSizeInPointsRefreshRateError]
type ISLVirtualDisplayMode interface {
	objectivec.IObject

	// Topic: Methods

	DictionaryRepresentation() objectivec.IObject
	Eotf() uint64
	SetEotf(value uint64)
	IsEqualToMode(mode objectivec.IObject) bool
	Options() uint64
	SetOptions(value uint64)
	RefreshDeadline() float64
	SetRefreshDeadline(value float64)
	RefreshRate() float32
	SizeInPixels() unsafe.Pointer
	SizeInPoints() unsafe.Pointer
	InitWithSizeInPixelsSizeInPointsRefreshRateError(pixels unsafe.Pointer, points unsafe.Pointer, rate float32) (SLVirtualDisplayMode, error)
}

// Init initializes the instance.
func (s SLVirtualDisplayMode) Init() SLVirtualDisplayMode {
	rv := objc.SendIfResponds[SLVirtualDisplayMode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplayMode) Autorelease() SLVirtualDisplayMode {
	rv := objc.SendIfResponds[SLVirtualDisplayMode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplayMode creates a new SLVirtualDisplayMode instance.
func NewSLVirtualDisplayMode() SLVirtualDisplayMode {
	class := getSLVirtualDisplayModeClass()
	rv := objc.SendIfResponds[SLVirtualDisplayMode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLVirtualDisplayModeWithSizeInPixelsSizeInPointsRefreshRateError(pixels unsafe.Pointer, points unsafe.Pointer, rate float32) (SLVirtualDisplayMode, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplayModeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSizeInPixels:sizeInPoints:refreshRate:error:"), pixels, points, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayMode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SLVirtualDisplayMode{}, objc.ErrInitFailed
	}
	return SLVirtualDisplayModeFromID(rv), nil
}

func (s SLVirtualDisplayMode) DictionaryRepresentation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (s SLVirtualDisplayMode) IsEqualToMode(mode objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isEqualToMode:"), mode)
	return rv
}
func (s SLVirtualDisplayMode) InitWithSizeInPixelsSizeInPointsRefreshRateError(pixels unsafe.Pointer, points unsafe.Pointer, rate float32) (SLVirtualDisplayMode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithSizeInPixels:sizeInPoints:refreshRate:error:"), pixels, points, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayMode{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayModeFromID(rv), nil

}

func (_SLVirtualDisplayModeClass SLVirtualDisplayModeClass) ModeWithBackendMode(mode objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLVirtualDisplayModeClass.class), objc.Sel("modeWithBackendMode:"), mode)
	return objectivec.Object{ID: rv}
}
func (_SLVirtualDisplayModeClass SLVirtualDisplayModeClass) ModeWithDictionaryRepresentation(representation objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLVirtualDisplayModeClass.class), objc.Sel("modeWithDictionaryRepresentation:"), representation)
	return objectivec.Object{ID: rv}
}

func (s SLVirtualDisplayMode) Eotf() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("eotf"))
	return rv
}
func (s SLVirtualDisplayMode) SetEotf(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setEotf:"), value)
}
func (s SLVirtualDisplayMode) Options() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("options"))
	return rv
}
func (s SLVirtualDisplayMode) SetOptions(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setOptions:"), value)
}
func (s SLVirtualDisplayMode) RefreshDeadline() float64 {
	rv := objc.SendIfResponds[float64](s.ID, objc.Sel("refreshDeadline"))
	return rv
}
func (s SLVirtualDisplayMode) SetRefreshDeadline(value float64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setRefreshDeadline:"), value)
}
func (s SLVirtualDisplayMode) RefreshRate() float32 {
	rv := objc.SendIfResponds[float32](s.ID, objc.Sel("refreshRate"))
	return rv
}
func (s SLVirtualDisplayMode) SizeInPixels() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("sizeInPixels"))
	return rv
}
func (s SLVirtualDisplayMode) SizeInPoints() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("sizeInPoints"))
	return rv
}
