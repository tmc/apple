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
	rv := objc.Send[SLVirtualDisplayMode](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode
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
	SizeInPixels() objectivec.IObject
	SizeInPoints() objectivec.IObject
	InitWithSizeInPixelsSizeInPointsRefreshRateError(pixels objectivec.IObject, points objectivec.IObject, rate float32) (SLVirtualDisplayMode, error)
}

// Init initializes the instance.
func (s SLVirtualDisplayMode) Init() SLVirtualDisplayMode {
	rv := objc.Send[SLVirtualDisplayMode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplayMode) Autorelease() SLVirtualDisplayMode {
	rv := objc.Send[SLVirtualDisplayMode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplayMode creates a new SLVirtualDisplayMode instance.
func NewSLVirtualDisplayMode() SLVirtualDisplayMode {
	class := getSLVirtualDisplayModeClass()
	rv := objc.Send[SLVirtualDisplayMode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/initWithSizeInPixels:sizeInPoints:refreshRate:error:
func NewSLVirtualDisplayModeWithSizeInPixelsSizeInPointsRefreshRateError(pixels objectivec.IObject, points objectivec.IObject, rate float32) (SLVirtualDisplayMode, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplayModeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSizeInPixels:sizeInPoints:refreshRate:error:"), pixels, points, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayMode{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayModeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/dictionaryRepresentation
func (s SLVirtualDisplayMode) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/isEqualToMode:
func (s SLVirtualDisplayMode) IsEqualToMode(mode objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEqualToMode:"), mode)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/initWithSizeInPixels:sizeInPoints:refreshRate:error:
func (s SLVirtualDisplayMode) InitWithSizeInPixelsSizeInPointsRefreshRateError(pixels objectivec.IObject, points objectivec.IObject, rate float32) (SLVirtualDisplayMode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithSizeInPixels:sizeInPoints:refreshRate:error:"), pixels, points, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayMode{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayModeFromID(rv), nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/modeWithBackendMode:
func (_SLVirtualDisplayModeClass SLVirtualDisplayModeClass) ModeWithBackendMode(mode objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayModeClass.class), objc.Sel("modeWithBackendMode:"), mode)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/modeWithDictionaryRepresentation:
func (_SLVirtualDisplayModeClass SLVirtualDisplayModeClass) ModeWithDictionaryRepresentation(representation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayModeClass.class), objc.Sel("modeWithDictionaryRepresentation:"), representation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/eotf
func (s SLVirtualDisplayMode) Eotf() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("eotf"))
	return rv
}
func (s SLVirtualDisplayMode) SetEotf(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setEotf:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/options
func (s SLVirtualDisplayMode) Options() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("options"))
	return rv
}
func (s SLVirtualDisplayMode) SetOptions(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setOptions:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/refreshDeadline
func (s SLVirtualDisplayMode) RefreshDeadline() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("refreshDeadline"))
	return rv
}
func (s SLVirtualDisplayMode) SetRefreshDeadline(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setRefreshDeadline:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/refreshRate
func (s SLVirtualDisplayMode) RefreshRate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("refreshRate"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/sizeInPixels
func (s SLVirtualDisplayMode) SizeInPixels() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sizeInPixels"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayMode/sizeInPoints
func (s SLVirtualDisplayMode) SizeInPoints() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sizeInPoints"))
	return objectivec.Object{ID: rv}
}
