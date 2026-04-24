// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLVirtualDisplayCapabilities] class.
var (
	_SLVirtualDisplayCapabilitiesClass     SLVirtualDisplayCapabilitiesClass
	_SLVirtualDisplayCapabilitiesClassOnce sync.Once
)

func getSLVirtualDisplayCapabilitiesClass() SLVirtualDisplayCapabilitiesClass {
	_SLVirtualDisplayCapabilitiesClassOnce.Do(func() {
		_SLVirtualDisplayCapabilitiesClass = SLVirtualDisplayCapabilitiesClass{class: objc.GetClass("SLVirtualDisplayCapabilities")}
	})
	return _SLVirtualDisplayCapabilitiesClass
}

// GetSLVirtualDisplayCapabilitiesClass returns the class object for SLVirtualDisplayCapabilities.
func GetSLVirtualDisplayCapabilitiesClass() SLVirtualDisplayCapabilitiesClass {
	return getSLVirtualDisplayCapabilitiesClass()
}

type SLVirtualDisplayCapabilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLVirtualDisplayCapabilitiesClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLVirtualDisplayCapabilitiesClass) Alloc() SLVirtualDisplayCapabilities {
	rv := objc.Send[SLVirtualDisplayCapabilities](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLVirtualDisplayCapabilities.DictionaryRepresentation]
//   - [SLVirtualDisplayCapabilities.MaximumPixelsPerPoint]
//   - [SLVirtualDisplayCapabilities.MaximumSizeInPixels]
//   - [SLVirtualDisplayCapabilities.MinimumPixelsPerPoint]
//   - [SLVirtualDisplayCapabilities.MinimumRefreshRate]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities
type SLVirtualDisplayCapabilities struct {
	objectivec.Object
}

// SLVirtualDisplayCapabilitiesFromID constructs a [SLVirtualDisplayCapabilities] from an objc.ID.
func SLVirtualDisplayCapabilitiesFromID(id objc.ID) SLVirtualDisplayCapabilities {
	return SLVirtualDisplayCapabilities{objectivec.Object{ID: id}}
}

// Ensure SLVirtualDisplayCapabilities implements ISLVirtualDisplayCapabilities.
var _ ISLVirtualDisplayCapabilities = SLVirtualDisplayCapabilities{}

// An interface definition for the [SLVirtualDisplayCapabilities] class.
//
// # Methods
//
//   - [ISLVirtualDisplayCapabilities.DictionaryRepresentation]
//   - [ISLVirtualDisplayCapabilities.MaximumPixelsPerPoint]
//   - [ISLVirtualDisplayCapabilities.MaximumSizeInPixels]
//   - [ISLVirtualDisplayCapabilities.MinimumPixelsPerPoint]
//   - [ISLVirtualDisplayCapabilities.MinimumRefreshRate]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities
type ISLVirtualDisplayCapabilities interface {
	objectivec.IObject

	// Topic: Methods

	DictionaryRepresentation() objectivec.IObject
	MaximumPixelsPerPoint() objectivec.IObject
	MaximumSizeInPixels() objectivec.IObject
	MinimumPixelsPerPoint() objectivec.IObject
	MinimumRefreshRate() float32
}

// Init initializes the instance.
func (s SLVirtualDisplayCapabilities) Init() SLVirtualDisplayCapabilities {
	rv := objc.Send[SLVirtualDisplayCapabilities](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplayCapabilities) Autorelease() SLVirtualDisplayCapabilities {
	rv := objc.Send[SLVirtualDisplayCapabilities](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplayCapabilities creates a new SLVirtualDisplayCapabilities instance.
func NewSLVirtualDisplayCapabilities() SLVirtualDisplayCapabilities {
	class := getSLVirtualDisplayCapabilitiesClass()
	rv := objc.Send[SLVirtualDisplayCapabilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities/dictionaryRepresentation
func (s SLVirtualDisplayCapabilities) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities/maximumPixelsPerPoint
func (s SLVirtualDisplayCapabilities) MaximumPixelsPerPoint() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("maximumPixelsPerPoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities/maximumSizeInPixels
func (s SLVirtualDisplayCapabilities) MaximumSizeInPixels() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("maximumSizeInPixels"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities/minimumPixelsPerPoint
func (s SLVirtualDisplayCapabilities) MinimumPixelsPerPoint() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("minimumPixelsPerPoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayCapabilities/minimumRefreshRate
func (s SLVirtualDisplayCapabilities) MinimumRefreshRate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("minimumRefreshRate"))
	return rv
}
