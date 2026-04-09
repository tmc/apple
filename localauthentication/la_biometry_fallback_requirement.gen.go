// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LABiometryFallbackRequirement] class.
var (
	_LABiometryFallbackRequirementClass     LABiometryFallbackRequirementClass
	_LABiometryFallbackRequirementClassOnce sync.Once
)

func getLABiometryFallbackRequirementClass() LABiometryFallbackRequirementClass {
	_LABiometryFallbackRequirementClassOnce.Do(func() {
		_LABiometryFallbackRequirementClass = LABiometryFallbackRequirementClass{class: objc.GetClass("LABiometryFallbackRequirement")}
	})
	return _LABiometryFallbackRequirementClass
}

// GetLABiometryFallbackRequirementClass returns the class object for LABiometryFallbackRequirement.
func GetLABiometryFallbackRequirementClass() LABiometryFallbackRequirementClass {
	return getLABiometryFallbackRequirementClass()
}

type LABiometryFallbackRequirementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LABiometryFallbackRequirementClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LABiometryFallbackRequirementClass) Alloc() LABiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A set of requirements to fall back on if biometrics aren’t present.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement
type LABiometryFallbackRequirement struct {
	objectivec.Object
}

// LABiometryFallbackRequirementFromID constructs a [LABiometryFallbackRequirement] from an objc.ID.
//
// A set of requirements to fall back on if biometrics aren’t present.
func LABiometryFallbackRequirementFromID(id objc.ID) LABiometryFallbackRequirement {
	return LABiometryFallbackRequirement{objectivec.Object{ID: id}}
}

// NOTE: LABiometryFallbackRequirement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LABiometryFallbackRequirement] class.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement
type ILABiometryFallbackRequirement interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b LABiometryFallbackRequirement) Init() LABiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b LABiometryFallbackRequirement) Autorelease() LABiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewLABiometryFallbackRequirement creates a new LABiometryFallbackRequirement instance.
func NewLABiometryFallbackRequirement() LABiometryFallbackRequirement {
	class := getLABiometryFallbackRequirementClass()
	rv := objc.Send[LABiometryFallbackRequirement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The default biometric fallback requirement.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/default
func (_LABiometryFallbackRequirementClass LABiometryFallbackRequirementClass) DefaultRequirement() LABiometryFallbackRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LABiometryFallbackRequirementClass.class), objc.Sel("defaultRequirement"))
	return LABiometryFallbackRequirementFromID(objc.ID(rv))
}

// The fallback requirement that requires entering the device passcode.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/devicePasscode
func (_LABiometryFallbackRequirementClass LABiometryFallbackRequirementClass) DevicePasscodeRequirement() LABiometryFallbackRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LABiometryFallbackRequirementClass.class), objc.Sel("devicePasscodeRequirement"))
	return LABiometryFallbackRequirementFromID(objc.ID(rv))
}
