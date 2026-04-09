// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [LAEnvironmentMechanismCompanion] class.
var (
	_LAEnvironmentMechanismCompanionClass     LAEnvironmentMechanismCompanionClass
	_LAEnvironmentMechanismCompanionClassOnce sync.Once
)

func getLAEnvironmentMechanismCompanionClass() LAEnvironmentMechanismCompanionClass {
	_LAEnvironmentMechanismCompanionClassOnce.Do(func() {
		_LAEnvironmentMechanismCompanionClass = LAEnvironmentMechanismCompanionClass{class: objc.GetClass("LAEnvironmentMechanismCompanion")}
	})
	return _LAEnvironmentMechanismCompanionClass
}

// GetLAEnvironmentMechanismCompanionClass returns the class object for LAEnvironmentMechanismCompanion.
func GetLAEnvironmentMechanismCompanionClass() LAEnvironmentMechanismCompanionClass {
	return getLAEnvironmentMechanismCompanionClass()
}

type LAEnvironmentMechanismCompanionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAEnvironmentMechanismCompanionClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAEnvironmentMechanismCompanionClass) Alloc() LAEnvironmentMechanismCompanion {
	rv := objc.Send[LAEnvironmentMechanismCompanion](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LAEnvironmentMechanismCompanion.StateHash]
//   - [LAEnvironmentMechanismCompanion.Type]: Type of the companion.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion
type LAEnvironmentMechanismCompanion struct {
	LAEnvironmentMechanism
}

// LAEnvironmentMechanismCompanionFromID constructs a [LAEnvironmentMechanismCompanion] from an objc.ID.
func LAEnvironmentMechanismCompanionFromID(id objc.ID) LAEnvironmentMechanismCompanion {
	return LAEnvironmentMechanismCompanion{LAEnvironmentMechanism: LAEnvironmentMechanismFromID(id)}
}

// NOTE: LAEnvironmentMechanismCompanion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAEnvironmentMechanismCompanion] class.
//
// # Instance Properties
//
//   - [ILAEnvironmentMechanismCompanion.StateHash]
//   - [ILAEnvironmentMechanismCompanion.Type]: Type of the companion.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion
type ILAEnvironmentMechanismCompanion interface {
	ILAEnvironmentMechanism

	// Topic: Instance Properties

	StateHash() foundation.INSData
	// Type of the companion.
	Type() LACompanionType
}

// Init initializes the instance.
func (e LAEnvironmentMechanismCompanion) Init() LAEnvironmentMechanismCompanion {
	rv := objc.Send[LAEnvironmentMechanismCompanion](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e LAEnvironmentMechanismCompanion) Autorelease() LAEnvironmentMechanismCompanion {
	rv := objc.Send[LAEnvironmentMechanismCompanion](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAEnvironmentMechanismCompanion creates a new LAEnvironmentMechanismCompanion instance.
func NewLAEnvironmentMechanismCompanion() LAEnvironmentMechanismCompanion {
	class := getLAEnvironmentMechanismCompanionClass()
	rv := objc.Send[LAEnvironmentMechanismCompanion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// Hash of the current companion pairing as returned by @c
// LAContext.domainState.companion.stateHash(for:)
//
// If no companion are paired for this companion type, @c stateHash property
// is @c nil. If at least one companion is paired for this companion type, @c
// stateHash is not @c nil and it changes whenever the set of paired
// companions of this type is changed.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/stateHash
func (e LAEnvironmentMechanismCompanion) StateHash() foundation.INSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("stateHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// Type of the companion.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/type
func (e LAEnvironmentMechanismCompanion) Type() LACompanionType {
	rv := objc.Send[LACompanionType](e.ID, objc.Sel("type"))
	return LACompanionType(rv)
}
