// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LAEnvironmentState] class.
var (
	_LAEnvironmentStateClass     LAEnvironmentStateClass
	_LAEnvironmentStateClassOnce sync.Once
)

func getLAEnvironmentStateClass() LAEnvironmentStateClass {
	_LAEnvironmentStateClassOnce.Do(func() {
		_LAEnvironmentStateClass = LAEnvironmentStateClass{class: objc.GetClass("LAEnvironmentState")}
	})
	return _LAEnvironmentStateClass
}

// GetLAEnvironmentStateClass returns the class object for LAEnvironmentState.
func GetLAEnvironmentStateClass() LAEnvironmentStateClass {
	return getLAEnvironmentStateClass()
}

type LAEnvironmentStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentStateClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentStateClass) Alloc() LAEnvironmentState {
	rv := objc.Send[LAEnvironmentState](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironmentState.AllMechanisms]
//   - [LAEnvironmentState.Biometry]
//   - [LAEnvironmentState.Companions]
//   - [LAEnvironmentState.UserPassword]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class
type LAEnvironmentState struct {
	objectivec.Object
}

// LAEnvironmentStateFromID constructs a [LAEnvironmentState] from an objc.ID.
func LAEnvironmentStateFromID(id objc.ID) LAEnvironmentState {
	return LAEnvironmentState{objectivec.Object{ID: id}}
}

// NOTE: LAEnvironmentState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironmentState] class.
//
// # Instance Properties
//
//   - [ILAEnvironmentState.AllMechanisms]
//   - [ILAEnvironmentState.Biometry]
//   - [ILAEnvironmentState.Companions]
//   - [ILAEnvironmentState.UserPassword]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class
type ILAEnvironmentState interface {
	objectivec.IObject

	// Topic: Instance Properties

	AllMechanisms() []LAEnvironmentMechanism
	Biometry() ILAEnvironmentMechanismBiometry
	Companions() []LAEnvironmentMechanismCompanion
	UserPassword() ILAEnvironmentMechanismUserPassword
}

// Init initializes the instance.
func (e LAEnvironmentState) Init() LAEnvironmentState {
	rv := objc.Send[LAEnvironmentState](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironmentState) Autorelease() LAEnvironmentState {
	rv := objc.Send[LAEnvironmentState](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironmentState creates a new LAEnvironmentState instance.
func NewLAEnvironmentState() LAEnvironmentState {
	class := getLAEnvironmentStateClass()
	rv := objc.Send[LAEnvironmentState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// Information about all authentication mechanisms.
//
// This property aggregates @c biometry, @c userPassword, @c companions and
// any future authentication mechanisms.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/allMechanisms
func (e LAEnvironmentState) AllMechanisms() []LAEnvironmentMechanism {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("allMechanisms"))
	return objc.ConvertSlice(rv, func(id objc.ID) LAEnvironmentMechanism {
		return LAEnvironmentMechanismFromID(id)
	})
}

// # Discussion
//
// Information about biometric authentication (Touch ID, Face ID or Optic ID).
//
// @c nil if biometry is not supported by this device.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/biometry
func (e LAEnvironmentState) Biometry() ILAEnvironmentMechanismBiometry {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("biometry"))
	return LAEnvironmentMechanismBiometryFromID(objc.ID(rv))
}

// # Discussion
//
// Companion authentication mechanisms.
//
// Companion mechanisms such as Apple Watch can appear and disappear as they
// get in and out of reach, but this property enumerates paired companions,
// even if they are not reachable at the moment. Check @c isUsable property to
// determine if a particular companion type is available for use. Note that
// items in this array represent paired companion types, not individual
// devices. Therefore, even if the user has paired multiple Apple Watch
// devices for companion authentication, the array will contain only one
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/companions
func (e LAEnvironmentState) Companions() []LAEnvironmentMechanismCompanion {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("companions"))
	return objc.ConvertSlice(rv, func(id objc.ID) LAEnvironmentMechanismCompanion {
		return LAEnvironmentMechanismCompanionFromID(id)
	})
}

// # Discussion
//
// Information about local user password (on macOS) or passcode (on embedded
// platforms).
//
// @c nil if user password or passcode is not supported by this device.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/userPassword
func (e LAEnvironmentState) UserPassword() ILAEnvironmentMechanismUserPassword {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userPassword"))
	return LAEnvironmentMechanismUserPasswordFromID(objc.ID(rv))
}
