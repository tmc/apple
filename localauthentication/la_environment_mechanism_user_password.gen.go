// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [LAEnvironmentMechanismUserPassword] class.
var (
	_LAEnvironmentMechanismUserPasswordClass     LAEnvironmentMechanismUserPasswordClass
	_LAEnvironmentMechanismUserPasswordClassOnce sync.Once
)

func getLAEnvironmentMechanismUserPasswordClass() LAEnvironmentMechanismUserPasswordClass {
	_LAEnvironmentMechanismUserPasswordClassOnce.Do(func() {
		_LAEnvironmentMechanismUserPasswordClass = LAEnvironmentMechanismUserPasswordClass{class: objc.GetClass("LAEnvironmentMechanismUserPassword")}
	})
	return _LAEnvironmentMechanismUserPasswordClass
}

// GetLAEnvironmentMechanismUserPasswordClass returns the class object for LAEnvironmentMechanismUserPassword.
func GetLAEnvironmentMechanismUserPasswordClass() LAEnvironmentMechanismUserPasswordClass {
	return getLAEnvironmentMechanismUserPasswordClass()
}

type LAEnvironmentMechanismUserPasswordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentMechanismUserPasswordClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentMechanismUserPasswordClass) Alloc() LAEnvironmentMechanismUserPassword {
	rv := objc.Send[LAEnvironmentMechanismUserPassword](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironmentMechanismUserPassword.IsSet]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword
type LAEnvironmentMechanismUserPassword struct {
	LAEnvironmentMechanism
}

// LAEnvironmentMechanismUserPasswordFromID constructs a [LAEnvironmentMechanismUserPassword] from an objc.ID.
func LAEnvironmentMechanismUserPasswordFromID(id objc.ID) LAEnvironmentMechanismUserPassword {
	return LAEnvironmentMechanismUserPassword{LAEnvironmentMechanism: LAEnvironmentMechanismFromID(id)}
}

// NOTE: LAEnvironmentMechanismUserPassword adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironmentMechanismUserPassword] class.
//
// # Instance Properties
//
//   - [ILAEnvironmentMechanismUserPassword.IsSet]
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword
type ILAEnvironmentMechanismUserPassword interface {
	ILAEnvironmentMechanism

	// Topic: Instance Properties

	IsSet() bool
}

// Init initializes the instance.
func (e LAEnvironmentMechanismUserPassword) Init() LAEnvironmentMechanismUserPassword {
	rv := objc.Send[LAEnvironmentMechanismUserPassword](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironmentMechanismUserPassword) Autorelease() LAEnvironmentMechanismUserPassword {
	rv := objc.Send[LAEnvironmentMechanismUserPassword](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironmentMechanismUserPassword creates a new LAEnvironmentMechanismUserPassword instance.
func NewLAEnvironmentMechanismUserPassword() LAEnvironmentMechanismUserPassword {
	class := getLAEnvironmentMechanismUserPasswordClass()
	rv := objc.Send[LAEnvironmentMechanismUserPassword](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// Whether the local user password or passcode is set on this device.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword/isSet
func (e LAEnvironmentMechanismUserPassword) IsSet() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isSet"))
	return rv
}
