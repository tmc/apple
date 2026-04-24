// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLVirtualDisplaySettings] class.
var (
	_SLVirtualDisplaySettingsClass     SLVirtualDisplaySettingsClass
	_SLVirtualDisplaySettingsClassOnce sync.Once
)

func getSLVirtualDisplaySettingsClass() SLVirtualDisplaySettingsClass {
	_SLVirtualDisplaySettingsClassOnce.Do(func() {
		_SLVirtualDisplaySettingsClass = SLVirtualDisplaySettingsClass{class: objc.GetClass("SLVirtualDisplaySettings")}
	})
	return _SLVirtualDisplaySettingsClass
}

// GetSLVirtualDisplaySettingsClass returns the class object for SLVirtualDisplaySettings.
func GetSLVirtualDisplaySettingsClass() SLVirtualDisplaySettingsClass {
	return getSLVirtualDisplaySettingsClass()
}

type SLVirtualDisplaySettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLVirtualDisplaySettingsClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLVirtualDisplaySettingsClass) Alloc() SLVirtualDisplaySettings {
	rv := objc.Send[SLVirtualDisplaySettings](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLVirtualDisplaySettings.DictionaryRepresentation]
//   - [SLVirtualDisplaySettings.NativeMode]
//   - [SLVirtualDisplaySettings.OptionalModes]
//   - [SLVirtualDisplaySettings.PreferredMode]
//   - [SLVirtualDisplaySettings.Rotations]
//   - [SLVirtualDisplaySettings.InitWithNativeModePreferredModeOptionalModesRotationsError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings
type SLVirtualDisplaySettings struct {
	objectivec.Object
}

// SLVirtualDisplaySettingsFromID constructs a [SLVirtualDisplaySettings] from an objc.ID.
func SLVirtualDisplaySettingsFromID(id objc.ID) SLVirtualDisplaySettings {
	return SLVirtualDisplaySettings{objectivec.Object{ID: id}}
}

// Ensure SLVirtualDisplaySettings implements ISLVirtualDisplaySettings.
var _ ISLVirtualDisplaySettings = SLVirtualDisplaySettings{}

// An interface definition for the [SLVirtualDisplaySettings] class.
//
// # Methods
//
//   - [ISLVirtualDisplaySettings.DictionaryRepresentation]
//   - [ISLVirtualDisplaySettings.NativeMode]
//   - [ISLVirtualDisplaySettings.OptionalModes]
//   - [ISLVirtualDisplaySettings.PreferredMode]
//   - [ISLVirtualDisplaySettings.Rotations]
//   - [ISLVirtualDisplaySettings.InitWithNativeModePreferredModeOptionalModesRotationsError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings
type ISLVirtualDisplaySettings interface {
	objectivec.IObject

	// Topic: Methods

	DictionaryRepresentation() objectivec.IObject
	NativeMode() ISLVirtualDisplayMode
	OptionalModes() foundation.INSArray
	PreferredMode() ISLVirtualDisplayMode
	Rotations() uint64
	InitWithNativeModePreferredModeOptionalModesRotationsError(mode objectivec.IObject, mode2 objectivec.IObject, modes objectivec.IObject, rotations uint64) (SLVirtualDisplaySettings, error)
}

// Init initializes the instance.
func (s SLVirtualDisplaySettings) Init() SLVirtualDisplaySettings {
	rv := objc.Send[SLVirtualDisplaySettings](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplaySettings) Autorelease() SLVirtualDisplaySettings {
	rv := objc.Send[SLVirtualDisplaySettings](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplaySettings creates a new SLVirtualDisplaySettings instance.
func NewSLVirtualDisplaySettings() SLVirtualDisplaySettings {
	class := getSLVirtualDisplaySettingsClass()
	rv := objc.Send[SLVirtualDisplaySettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/initWithNativeMode:preferredMode:optionalModes:rotations:error:
func NewSLVirtualDisplaySettingsWithNativeModePreferredModeOptionalModesRotationsError(mode objectivec.IObject, mode2 objectivec.IObject, modes objectivec.IObject, rotations uint64) (SLVirtualDisplaySettings, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplaySettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNativeMode:preferredMode:optionalModes:rotations:error:"), mode, mode2, modes, rotations, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplaySettings{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplaySettingsFromID(rv), nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/dictionaryRepresentation
func (s SLVirtualDisplaySettings) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/initWithNativeMode:preferredMode:optionalModes:rotations:error:
func (s SLVirtualDisplaySettings) InitWithNativeModePreferredModeOptionalModesRotationsError(mode objectivec.IObject, mode2 objectivec.IObject, modes objectivec.IObject, rotations uint64) (SLVirtualDisplaySettings, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithNativeMode:preferredMode:optionalModes:rotations:error:"), mode, mode2, modes, rotations, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplaySettings{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplaySettingsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/settingsWithBackendSettings:
func (_SLVirtualDisplaySettingsClass SLVirtualDisplaySettingsClass) SettingsWithBackendSettings(settings objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplaySettingsClass.class), objc.Sel("settingsWithBackendSettings:"), settings)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/settingsWithDictionaryRepresentation:
func (_SLVirtualDisplaySettingsClass SLVirtualDisplaySettingsClass) SettingsWithDictionaryRepresentation(representation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplaySettingsClass.class), objc.Sel("settingsWithDictionaryRepresentation:"), representation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/nativeMode
func (s SLVirtualDisplaySettings) NativeMode() ISLVirtualDisplayMode {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nativeMode"))
	return SLVirtualDisplayModeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/optionalModes
func (s SLVirtualDisplaySettings) OptionalModes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("optionalModes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/preferredMode
func (s SLVirtualDisplaySettings) PreferredMode() ISLVirtualDisplayMode {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("preferredMode"))
	return SLVirtualDisplayModeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplaySettings/rotations
func (s SLVirtualDisplaySettings) Rotations() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("rotations"))
	return rv
}
