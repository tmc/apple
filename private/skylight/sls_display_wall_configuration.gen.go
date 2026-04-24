// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayWallConfiguration] class.
var (
	_SLSDisplayWallConfigurationClass     SLSDisplayWallConfigurationClass
	_SLSDisplayWallConfigurationClassOnce sync.Once
)

func getSLSDisplayWallConfigurationClass() SLSDisplayWallConfigurationClass {
	_SLSDisplayWallConfigurationClassOnce.Do(func() {
		_SLSDisplayWallConfigurationClass = SLSDisplayWallConfigurationClass{class: objc.GetClass("SLSDisplayWallConfiguration")}
	})
	return _SLSDisplayWallConfigurationClass
}

// GetSLSDisplayWallConfigurationClass returns the class object for SLSDisplayWallConfiguration.
func GetSLSDisplayWallConfigurationClass() SLSDisplayWallConfigurationClass {
	return getSLSDisplayWallConfigurationClass()
}

type SLSDisplayWallConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayWallConfigurationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayWallConfigurationClass) Alloc() SLSDisplayWallConfiguration {
	rv := objc.Send[SLSDisplayWallConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayWallConfiguration.DisplayIDs]
//   - [SLSDisplayWallConfiguration.SetDisplayIDs]
//   - [SLSDisplayWallConfiguration.GridHeight]
//   - [SLSDisplayWallConfiguration.SetGridHeight]
//   - [SLSDisplayWallConfiguration.GridWidth]
//   - [SLSDisplayWallConfiguration.SetGridWidth]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallConfiguration
type SLSDisplayWallConfiguration struct {
	objectivec.Object
}

// SLSDisplayWallConfigurationFromID constructs a [SLSDisplayWallConfiguration] from an objc.ID.
func SLSDisplayWallConfigurationFromID(id objc.ID) SLSDisplayWallConfiguration {
	return SLSDisplayWallConfiguration{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayWallConfiguration implements ISLSDisplayWallConfiguration.
var _ ISLSDisplayWallConfiguration = SLSDisplayWallConfiguration{}

// An interface definition for the [SLSDisplayWallConfiguration] class.
//
// # Methods
//
//   - [ISLSDisplayWallConfiguration.DisplayIDs]
//   - [ISLSDisplayWallConfiguration.SetDisplayIDs]
//   - [ISLSDisplayWallConfiguration.GridHeight]
//   - [ISLSDisplayWallConfiguration.SetGridHeight]
//   - [ISLSDisplayWallConfiguration.GridWidth]
//   - [ISLSDisplayWallConfiguration.SetGridWidth]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallConfiguration
type ISLSDisplayWallConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	DisplayIDs() foundation.INSArray
	SetDisplayIDs(value foundation.INSArray)
	GridHeight() foundation.NSNumber
	SetGridHeight(value foundation.NSNumber)
	GridWidth() foundation.NSNumber
	SetGridWidth(value foundation.NSNumber)
}

// Init initializes the instance.
func (s SLSDisplayWallConfiguration) Init() SLSDisplayWallConfiguration {
	rv := objc.Send[SLSDisplayWallConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayWallConfiguration) Autorelease() SLSDisplayWallConfiguration {
	rv := objc.Send[SLSDisplayWallConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayWallConfiguration creates a new SLSDisplayWallConfiguration instance.
func NewSLSDisplayWallConfiguration() SLSDisplayWallConfiguration {
	class := getSLSDisplayWallConfigurationClass()
	rv := objc.Send[SLSDisplayWallConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallConfiguration/displayIDs
func (s SLSDisplayWallConfiguration) DisplayIDs() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIDs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLSDisplayWallConfiguration) SetDisplayIDs(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayIDs:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallConfiguration/gridHeight
func (s SLSDisplayWallConfiguration) GridHeight() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("gridHeight"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLSDisplayWallConfiguration) SetGridHeight(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setGridHeight:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallConfiguration/gridWidth
func (s SLSDisplayWallConfiguration) GridWidth() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("gridWidth"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLSDisplayWallConfiguration) SetGridWidth(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setGridWidth:"), value)
}
