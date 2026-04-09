// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LAAuthenticationRequirement] class.
var (
	_LAAuthenticationRequirementClass     LAAuthenticationRequirementClass
	_LAAuthenticationRequirementClassOnce sync.Once
)

func getLAAuthenticationRequirementClass() LAAuthenticationRequirementClass {
	_LAAuthenticationRequirementClassOnce.Do(func() {
		_LAAuthenticationRequirementClass = LAAuthenticationRequirementClass{class: objc.GetClass("LAAuthenticationRequirement")}
	})
	return _LAAuthenticationRequirementClass
}

// GetLAAuthenticationRequirementClass returns the class object for LAAuthenticationRequirement.
func GetLAAuthenticationRequirementClass() LAAuthenticationRequirementClass {
	return getLAAuthenticationRequirementClass()
}

type LAAuthenticationRequirementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAAuthenticationRequirementClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAAuthenticationRequirementClass) Alloc() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A set of requirements that protect a right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement
type LAAuthenticationRequirement struct {
	objectivec.Object
}

// LAAuthenticationRequirementFromID constructs a [LAAuthenticationRequirement] from an objc.ID.
//
// A set of requirements that protect a right.
func LAAuthenticationRequirementFromID(id objc.ID) LAAuthenticationRequirement {
	return LAAuthenticationRequirement{objectivec.Object{ID: id}}
}

// NOTE: LAAuthenticationRequirement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAAuthenticationRequirement] class.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement
type ILAAuthenticationRequirement interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a LAAuthenticationRequirement) Init() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a LAAuthenticationRequirement) Autorelease() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAAuthenticationRequirement creates a new LAAuthenticationRequirement instance.
func NewLAAuthenticationRequirement() LAAuthenticationRequirement {
	class := getLAAuthenticationRequirementClass()
	rv := objc.Send[LAAuthenticationRequirement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a requirement that requires biometric authentication or a fallback
// requirement that you specify.
//
// fallback: A requirement to use a fallback if biometric authentication fails or is
// unavailable, or if the user prefers not to use biometric authentication.
//
// # Return Value
//
// Returns a requirement that requires biometric authentication or a fallback
// requirement that you specify.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry(fallback:)
func (_LAAuthenticationRequirementClass LAAuthenticationRequirementClass) BiometryRequirementWithFallback(fallback ILABiometryFallbackRequirement) LAAuthenticationRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LAAuthenticationRequirementClass.class), objc.Sel("biometryRequirementWithFallback:"), fallback)
	return LAAuthenticationRequirementFromID(rv)
}

// The requirement that requires user authentication.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/default
func (_LAAuthenticationRequirementClass LAAuthenticationRequirementClass) DefaultRequirement() LAAuthenticationRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LAAuthenticationRequirementClass.class), objc.Sel("defaultRequirement"))
	return LAAuthenticationRequirementFromID(objc.ID(rv))
}

// The requirement that requires biometric authentication.
//
// # Discussion
//
// Authorizations with this requirement fail when biometrics aren’t
// available on the current device or there aren’t any enrolled biometrics
// on the current device.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry
func (_LAAuthenticationRequirementClass LAAuthenticationRequirementClass) BiometryRequirement() LAAuthenticationRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LAAuthenticationRequirementClass.class), objc.Sel("biometryRequirement"))
	return LAAuthenticationRequirementFromID(objc.ID(rv))
}

// The requirement that requires user authentication with the current set of
// biometrics.
//
// # Discussion
//
// Authorizations with this requirement fail when:
//
// - Biometrics aren’t available on the current device. - There aren’t any
// enrolled biometrics on the current device. - There’s a change in enrolled
// biometrics on the current device. For example, adding a new finger to Touch
// ID changes the set of enrolled biometrics.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometryCurrentSet
func (_LAAuthenticationRequirementClass LAAuthenticationRequirementClass) BiometryCurrentSetRequirement() LAAuthenticationRequirement {
	rv := objc.Send[objc.ID](objc.ID(_LAAuthenticationRequirementClass.class), objc.Sel("biometryCurrentSetRequirement"))
	return LAAuthenticationRequirementFromID(objc.ID(rv))
}
