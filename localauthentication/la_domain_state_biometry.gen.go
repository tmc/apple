// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LADomainStateBiometry] class.
var (
	_LADomainStateBiometryClass     LADomainStateBiometryClass
	_LADomainStateBiometryClassOnce sync.Once
)

func getLADomainStateBiometryClass() LADomainStateBiometryClass {
	_LADomainStateBiometryClassOnce.Do(func() {
		_LADomainStateBiometryClass = LADomainStateBiometryClass{class: objc.GetClass("LADomainStateBiometry")}
	})
	return _LADomainStateBiometryClass
}

// GetLADomainStateBiometryClass returns the class object for LADomainStateBiometry.
func GetLADomainStateBiometryClass() LADomainStateBiometryClass {
	return getLADomainStateBiometryClass()
}

type LADomainStateBiometryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LADomainStateBiometryClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LADomainStateBiometryClass) Alloc() LADomainStateBiometry {
	rv := objc.Send[LADomainStateBiometry](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LADomainStateBiometry.BiometryType]: Indicates biometry type available on the device.
//   - [LADomainStateBiometry.StateHash]: Contains state hash data for the available biometry type. Returns `nil` if no biometry entities are enrolled.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry
type LADomainStateBiometry struct {
	objectivec.Object
}

// LADomainStateBiometryFromID constructs a [LADomainStateBiometry] from an objc.ID.
func LADomainStateBiometryFromID(id objc.ID) LADomainStateBiometry {
	return LADomainStateBiometry{objectivec.Object{ID: id}}
}

// NOTE: LADomainStateBiometry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LADomainStateBiometry] class.
//
// # Instance Properties
//
//   - [ILADomainStateBiometry.BiometryType]: Indicates biometry type available on the device.
//   - [ILADomainStateBiometry.StateHash]: Contains state hash data for the available biometry type. Returns `nil` if no biometry entities are enrolled.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry
type ILADomainStateBiometry interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Indicates biometry type available on the device.
	BiometryType() LABiometryType
	// Contains state hash data for the available biometry type. Returns `nil` if no biometry entities are enrolled.
	StateHash() foundation.INSData
}

// Init initializes the instance.
func (d LADomainStateBiometry) Init() LADomainStateBiometry {
	rv := objc.Send[LADomainStateBiometry](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d LADomainStateBiometry) Autorelease() LADomainStateBiometry {
	rv := objc.Send[LADomainStateBiometry](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewLADomainStateBiometry creates a new LADomainStateBiometry instance.
func NewLADomainStateBiometry() LADomainStateBiometry {
	class := getLADomainStateBiometryClass()
	rv := objc.Send[LADomainStateBiometry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates biometry type available on the device.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/biometryType
func (d LADomainStateBiometry) BiometryType() LABiometryType {
	rv := objc.Send[LABiometryType](d.ID, objc.Sel("biometryType"))
	return LABiometryType(rv)
}

// Contains state hash data for the available biometry type. Returns `nil` if
// no biometry entities are enrolled.
//
// # Discussion
//
// If biometric database was modified (fingers, faces were removed or added),
// `stateHash` data will change. Nature of such database changes cannot be
// determined but comparing data of `stateHash` after different evaluatePolicy
// calls will reveal the fact database was changed between the calls.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/stateHash
func (d LADomainStateBiometry) StateHash() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stateHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
