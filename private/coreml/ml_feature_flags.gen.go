// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureFlags] class.
var (
	_MLFeatureFlagsClass     MLFeatureFlagsClass
	_MLFeatureFlagsClassOnce sync.Once
)

func getMLFeatureFlagsClass() MLFeatureFlagsClass {
	_MLFeatureFlagsClassOnce.Do(func() {
		_MLFeatureFlagsClass = MLFeatureFlagsClass{class: objc.GetClass("MLFeatureFlags")}
	})
	return _MLFeatureFlagsClass
}

// GetMLFeatureFlagsClass returns the class object for MLFeatureFlags.
func GetMLFeatureFlagsClass() MLFeatureFlagsClass {
	return getMLFeatureFlagsClass()
}

type MLFeatureFlagsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureFlagsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureFlagsClass) Alloc() MLFeatureFlags {
	rv := objc.Send[MLFeatureFlags](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLFeatureFlags.AddFeatureWithControlNameDefaultValue]
//   - [MLFeatureFlags.ControlKeyForFeature]
//   - [MLFeatureFlags.DefineFeatures]
//   - [MLFeatureFlags.Flags]
//   - [MLFeatureFlags.IsFeatureEnabled]
//   - [MLFeatureFlags.IsMPSGraphEnabled]
//   - [MLFeatureFlags.IsMPSGraphFP16Enabled]
//   - [MLFeatureFlags.OverrideOriginalValues]
//   - [MLFeatureFlags.RemoveOverrideForFeature]
//   - [MLFeatureFlags.SetOverrideForFeature]
//   - [MLFeatureFlags.UserDefaults]
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags
type MLFeatureFlags struct {
	objectivec.Object
}

// MLFeatureFlagsFromID constructs a [MLFeatureFlags] from an objc.ID.
func MLFeatureFlagsFromID(id objc.ID) MLFeatureFlags {
	return MLFeatureFlags{objectivec.Object{ID: id}}
}
// Ensure MLFeatureFlags implements IMLFeatureFlags.
var _ IMLFeatureFlags = MLFeatureFlags{}

// An interface definition for the [MLFeatureFlags] class.
//
// # Methods
//
//   - [IMLFeatureFlags.AddFeatureWithControlNameDefaultValue]
//   - [IMLFeatureFlags.ControlKeyForFeature]
//   - [IMLFeatureFlags.DefineFeatures]
//   - [IMLFeatureFlags.Flags]
//   - [IMLFeatureFlags.IsFeatureEnabled]
//   - [IMLFeatureFlags.IsMPSGraphEnabled]
//   - [IMLFeatureFlags.IsMPSGraphFP16Enabled]
//   - [IMLFeatureFlags.OverrideOriginalValues]
//   - [IMLFeatureFlags.RemoveOverrideForFeature]
//   - [IMLFeatureFlags.SetOverrideForFeature]
//   - [IMLFeatureFlags.UserDefaults]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags
type IMLFeatureFlags interface {
	objectivec.IObject

	// Topic: Methods

	AddFeatureWithControlNameDefaultValue(feature objectivec.IObject, name objectivec.IObject, value bool)
	ControlKeyForFeature(feature objectivec.IObject) objectivec.IObject
	DefineFeatures()
	Flags() foundation.INSDictionary
	IsFeatureEnabled(enabled objectivec.IObject) bool
	IsMPSGraphEnabled() bool
	IsMPSGraphFP16Enabled() bool
	OverrideOriginalValues() foundation.INSDictionary
	RemoveOverrideForFeature(feature objectivec.IObject) bool
	SetOverrideForFeature(override bool, feature objectivec.IObject) bool
	UserDefaults() *foundation.NSUserDefaults
}

// Init initializes the instance.
func (f MLFeatureFlags) Init() MLFeatureFlags {
	rv := objc.Send[MLFeatureFlags](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureFlags) Autorelease() MLFeatureFlags {
	rv := objc.Send[MLFeatureFlags](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureFlags creates a new MLFeatureFlags instance.
func NewMLFeatureFlags() MLFeatureFlags {
	class := getMLFeatureFlagsClass()
	rv := objc.Send[MLFeatureFlags](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/addFeature:withControlName:defaultValue:
func (f MLFeatureFlags) AddFeatureWithControlNameDefaultValue(feature objectivec.IObject, name objectivec.IObject, value bool) {
	objc.Send[objc.ID](f.ID, objc.Sel("addFeature:withControlName:defaultValue:"), feature, name, value)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/controlKeyForFeature:
func (f MLFeatureFlags) ControlKeyForFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("controlKeyForFeature:"), feature)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/defineFeatures
func (f MLFeatureFlags) DefineFeatures() {
	objc.Send[objc.ID](f.ID, objc.Sel("defineFeatures"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/isFeatureEnabled:
func (f MLFeatureFlags) IsFeatureEnabled(enabled objectivec.IObject) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isFeatureEnabled:"), enabled)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/isMPSGraphEnabled
func (f MLFeatureFlags) IsMPSGraphEnabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isMPSGraphEnabled"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/isMPSGraphFP16Enabled
func (f MLFeatureFlags) IsMPSGraphFP16Enabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isMPSGraphFP16Enabled"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/removeOverrideForFeature:
func (f MLFeatureFlags) RemoveOverrideForFeature(feature objectivec.IObject) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("removeOverrideForFeature:"), feature)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/setOverride:forFeature:
func (f MLFeatureFlags) SetOverrideForFeature(override bool, feature objectivec.IObject) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("setOverride:forFeature:"), override, feature)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/sharedFeatureFlags
func (_MLFeatureFlagsClass MLFeatureFlagsClass) SharedFeatureFlags() MLFeatureFlags {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureFlagsClass.class), objc.Sel("sharedFeatureFlags"))
	return MLFeatureFlagsFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/flags
func (f MLFeatureFlags) Flags() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("flags"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/overrideOriginalValues
func (f MLFeatureFlags) OverrideOriginalValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("overrideOriginalValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureFlags/userDefaults
func (f MLFeatureFlags) UserDefaults() *foundation.NSUserDefaults {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("userDefaults"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSUserDefaultsFromID(objc.ID(rv))
	return &val
}

