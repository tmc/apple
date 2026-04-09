// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [LAEnvironmentMechanismBiometry] class.
var (
	_LAEnvironmentMechanismBiometryClass     LAEnvironmentMechanismBiometryClass
	_LAEnvironmentMechanismBiometryClassOnce sync.Once
)

func getLAEnvironmentMechanismBiometryClass() LAEnvironmentMechanismBiometryClass {
	_LAEnvironmentMechanismBiometryClassOnce.Do(func() {
		_LAEnvironmentMechanismBiometryClass = LAEnvironmentMechanismBiometryClass{class: objc.GetClass("LAEnvironmentMechanismBiometry")}
	})
	return _LAEnvironmentMechanismBiometryClass
}

// GetLAEnvironmentMechanismBiometryClass returns the class object for LAEnvironmentMechanismBiometry.
func GetLAEnvironmentMechanismBiometryClass() LAEnvironmentMechanismBiometryClass {
	return getLAEnvironmentMechanismBiometryClass()
}

type LAEnvironmentMechanismBiometryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentMechanismBiometryClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentMechanismBiometryClass) Alloc() LAEnvironmentMechanismBiometry {
	rv := objc.Send[LAEnvironmentMechanismBiometry](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironmentMechanismBiometry.BiometryType]
//   - [LAEnvironmentMechanismBiometry.BuiltInSensorInaccessible]
//   - [LAEnvironmentMechanismBiometry.IsEnrolled]
//   - [LAEnvironmentMechanismBiometry.IsLockedOut]
//   - [LAEnvironmentMechanismBiometry.StateHash]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry
type LAEnvironmentMechanismBiometry struct {
	LAEnvironmentMechanism
}

// LAEnvironmentMechanismBiometryFromID constructs a [LAEnvironmentMechanismBiometry] from an objc.ID.
func LAEnvironmentMechanismBiometryFromID(id objc.ID) LAEnvironmentMechanismBiometry {
	return LAEnvironmentMechanismBiometry{LAEnvironmentMechanism: LAEnvironmentMechanismFromID(id)}
}

// NOTE: LAEnvironmentMechanismBiometry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironmentMechanismBiometry] class.
//
// # Instance Properties
//
//   - [ILAEnvironmentMechanismBiometry.BiometryType]
//   - [ILAEnvironmentMechanismBiometry.BuiltInSensorInaccessible]
//   - [ILAEnvironmentMechanismBiometry.IsEnrolled]
//   - [ILAEnvironmentMechanismBiometry.IsLockedOut]
//   - [ILAEnvironmentMechanismBiometry.StateHash]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry
type ILAEnvironmentMechanismBiometry interface {
	ILAEnvironmentMechanism

	// Topic: Instance Properties

	BiometryType() LABiometryType
	BuiltInSensorInaccessible() bool
	IsEnrolled() bool
	IsLockedOut() bool
	StateHash() foundation.INSData
}

// Init initializes the instance.
func (e LAEnvironmentMechanismBiometry) Init() LAEnvironmentMechanismBiometry {
	rv := objc.Send[LAEnvironmentMechanismBiometry](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironmentMechanismBiometry) Autorelease() LAEnvironmentMechanismBiometry {
	rv := objc.Send[LAEnvironmentMechanismBiometry](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironmentMechanismBiometry creates a new LAEnvironmentMechanismBiometry instance.
func NewLAEnvironmentMechanismBiometry() LAEnvironmentMechanismBiometry {
	class := getLAEnvironmentMechanismBiometryClass()
	rv := objc.Send[LAEnvironmentMechanismBiometry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// Type of biometry supported by the device.
//
// This property does not indicate whether biometry is available or not. It
// always reads the type of biometry supported by device hardware. You should
// check @c isUsable property to see if it is available for use.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/biometryType
func (e LAEnvironmentMechanismBiometry) BiometryType() LABiometryType {
	rv := objc.Send[LABiometryType](e.ID, objc.Sel("biometryType"))
	return LABiometryType(rv)
}

// # Discussion
//
// Whether the built in biometric sensor is inaccessible in the current
// configuration, preventing the use of biometry.
//
// Currently, the only example of this is a Clamshell Mode on macOS. The user
// will be not able to use Touch ID if the MacBook lid is closed while
// connected to external monitor and keyboard, unless the external keyboard
// has Touch ID.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/builtInSensorInaccessible
func (e LAEnvironmentMechanismBiometry) BuiltInSensorInaccessible() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("builtInSensorInaccessible"))
	return rv
}

// # Discussion
//
// Whether the user has enrolled this biometry.
//
// Even if biometry is enrolled, it does not necessarily mean that it can be
// used. You should check @c isUsable property to see if it is available for
// use.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isEnrolled
func (e LAEnvironmentMechanismBiometry) IsEnrolled() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isEnrolled"))
	return rv
}

// # Discussion
//
// Whether biometry is locked out.
//
// The system might lock the user out of biometry for various reasons. For
// example, with Face ID, the user is locked out after 5 failed match attempts
// in row. To recover from bio lockout, users need to enter their passcode
// (e.g. during device ulock).
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isLockedOut
func (e LAEnvironmentMechanismBiometry) IsLockedOut() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isLockedOut"))
	return rv
}

// # Discussion
//
// The application specific state of the biometric enrollment as returned by
// @c LAContext.domainState.biometry.stateHash
//
// This value represents the state of the enrollment and changes whenever the
// biometric enrollment is changed. It does not directly map to the enrolled
// templates, e.g. if a finger is added to Touch ID enrollement and then
// removed, the final state would be different. It also returns different
// values to different apps to prevent tracking of user identity.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/stateHash
func (e LAEnvironmentMechanismBiometry) StateHash() foundation.INSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("stateHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
