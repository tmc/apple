// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLFeatureFlags](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
	UserDefaults() foundation.UserDefaults
}

// Init initializes the instance.
func (m MLFeatureFlags) Init() MLFeatureFlags {
	rv := objc.SendIfResponds[MLFeatureFlags](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLFeatureFlags) Autorelease() MLFeatureFlags {
	rv := objc.SendIfResponds[MLFeatureFlags](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureFlags creates a new MLFeatureFlags instance.
func NewMLFeatureFlags() MLFeatureFlags {
	class := getMLFeatureFlagsClass()
	rv := objc.SendIfResponds[MLFeatureFlags](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLFeatureFlags) AddFeatureWithControlNameDefaultValue(feature objectivec.IObject, name objectivec.IObject, value bool) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("addFeature:withControlName:defaultValue:"), feature, name, value)
}
func (m MLFeatureFlags) ControlKeyForFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("controlKeyForFeature:"), feature)
	return objectivec.Object{ID: rv}
}
func (m MLFeatureFlags) DefineFeatures() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defineFeatures"))
}
func (m MLFeatureFlags) IsFeatureEnabled(enabled objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isFeatureEnabled:"), enabled)
	return rv
}
func (m MLFeatureFlags) IsMPSGraphEnabled() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isMPSGraphEnabled"))
	return rv
}
func (m MLFeatureFlags) IsMPSGraphFP16Enabled() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isMPSGraphFP16Enabled"))
	return rv
}
func (m MLFeatureFlags) RemoveOverrideForFeature(feature objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("removeOverrideForFeature:"), feature)
	return rv
}
func (m MLFeatureFlags) SetOverrideForFeature(override bool, feature objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("setOverride:forFeature:"), override, feature)
	return rv
}

func (_MLFeatureFlagsClass MLFeatureFlagsClass) SharedFeatureFlags() MLFeatureFlags {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLFeatureFlagsClass.class), objc.Sel("sharedFeatureFlags"))
	return MLFeatureFlagsFromID(rv)
}

func (m MLFeatureFlags) Flags() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("flags"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLFeatureFlags) OverrideOriginalValues() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("overrideOriginalValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLFeatureFlags) UserDefaults() foundation.UserDefaults {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("userDefaults"))
	return foundation.UserDefaultsFromID(objc.ID(rv))
}
