// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LAEnvironmentMechanism] class.
var (
	_LAEnvironmentMechanismClass     LAEnvironmentMechanismClass
	_LAEnvironmentMechanismClassOnce sync.Once
)

func getLAEnvironmentMechanismClass() LAEnvironmentMechanismClass {
	_LAEnvironmentMechanismClassOnce.Do(func() {
		_LAEnvironmentMechanismClass = LAEnvironmentMechanismClass{class: objc.GetClass("LAEnvironmentMechanism")}
	})
	return _LAEnvironmentMechanismClass
}

// GetLAEnvironmentMechanismClass returns the class object for LAEnvironmentMechanism.
func GetLAEnvironmentMechanismClass() LAEnvironmentMechanismClass {
	return getLAEnvironmentMechanismClass()
}

type LAEnvironmentMechanismClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentMechanismClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentMechanismClass) Alloc() LAEnvironmentMechanism {
	rv := objc.Send[LAEnvironmentMechanism](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironmentMechanism.IconSystemName]: Name of the SF Symbol representing this authentication mechanism.
//   - [LAEnvironmentMechanism.IsUsable]
//   - [LAEnvironmentMechanism.LocalizedName]: The localized name of the authentication mechanism, e.g. “Touch ID”, “Face ID” etc.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism
type LAEnvironmentMechanism struct {
	objectivec.Object
}

// LAEnvironmentMechanismFromID constructs a [LAEnvironmentMechanism] from an objc.ID.
func LAEnvironmentMechanismFromID(id objc.ID) LAEnvironmentMechanism {
	return LAEnvironmentMechanism{objectivec.Object{ID: id}}
}

// NOTE: LAEnvironmentMechanism adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironmentMechanism] class.
//
// # Instance Properties
//
//   - [ILAEnvironmentMechanism.IconSystemName]: Name of the SF Symbol representing this authentication mechanism.
//   - [ILAEnvironmentMechanism.IsUsable]
//   - [ILAEnvironmentMechanism.LocalizedName]: The localized name of the authentication mechanism, e.g. “Touch ID”, “Face ID” etc.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism
type ILAEnvironmentMechanism interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Name of the SF Symbol representing this authentication mechanism.
	IconSystemName() string
	IsUsable() bool
	// The localized name of the authentication mechanism, e.g. “Touch ID”, “Face ID” etc.
	LocalizedName() string
}

// Init initializes the instance.
func (e LAEnvironmentMechanism) Init() LAEnvironmentMechanism {
	rv := objc.Send[LAEnvironmentMechanism](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironmentMechanism) Autorelease() LAEnvironmentMechanism {
	rv := objc.Send[LAEnvironmentMechanism](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironmentMechanism creates a new LAEnvironmentMechanism instance.
func NewLAEnvironmentMechanism() LAEnvironmentMechanism {
	class := getLAEnvironmentMechanismClass()
	rv := objc.Send[LAEnvironmentMechanism](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Name of the SF Symbol representing this authentication mechanism.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/iconSystemName
func (e LAEnvironmentMechanism) IconSystemName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("iconSystemName"))
	return foundation.NSStringFromID(rv).String()
}

// # Discussion
//
// Whether the mechanism is available for use, i.e. whether the relevant
// preflight call of @c canEvaluatePolicy would succeed.
//
// properties of the subclass to determine why mechanism can’t be used.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/isUsable
func (e LAEnvironmentMechanism) IsUsable() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isUsable"))
	return rv
}

// The localized name of the authentication mechanism, e.g. “Touch ID”,
// “Face ID” etc.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/localizedName
func (e LAEnvironmentMechanism) LocalizedName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
