// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LADomainState] class.
var (
	_LADomainStateClass     LADomainStateClass
	_LADomainStateClassOnce sync.Once
)

func getLADomainStateClass() LADomainStateClass {
	_LADomainStateClassOnce.Do(func() {
		_LADomainStateClass = LADomainStateClass{class: objc.GetClass("LADomainState")}
	})
	return _LADomainStateClass
}

// GetLADomainStateClass returns the class object for LADomainState.
func GetLADomainStateClass() LADomainStateClass {
	return getLADomainStateClass()
}

type LADomainStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LADomainStateClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LADomainStateClass) Alloc() LADomainState {
	rv := objc.Send[LADomainState](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LADomainState.Biometry]: Contains biometric domain state.
//   - [LADomainState.Companion]: Contains companion domain state.
//   - [LADomainState.StateHash]: Contains combined state hash data for biometry and companion state hashes.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainState
type LADomainState struct {
	objectivec.Object
}

// LADomainStateFromID constructs a [LADomainState] from an objc.ID.
func LADomainStateFromID(id objc.ID) LADomainState {
	return LADomainState{objectivec.Object{ID: id}}
}

// NOTE: LADomainState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LADomainState] class.
//
// # Instance Properties
//
//   - [ILADomainState.Biometry]: Contains biometric domain state.
//   - [ILADomainState.Companion]: Contains companion domain state.
//   - [ILADomainState.StateHash]: Contains combined state hash data for biometry and companion state hashes.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainState
type ILADomainState interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Contains biometric domain state.
	Biometry() ILADomainStateBiometry
	// Contains companion domain state.
	Companion() ILADomainStateCompanion
	// Contains combined state hash data for biometry and companion state hashes.
	StateHash() foundation.INSData
}

// Init initializes the instance.
func (d LADomainState) Init() LADomainState {
	rv := objc.Send[LADomainState](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d LADomainState) Autorelease() LADomainState {
	rv := objc.Send[LADomainState](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewLADomainState creates a new LADomainState instance.
func NewLADomainState() LADomainState {
	class := getLADomainStateClass()
	rv := objc.Send[LADomainState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Contains biometric domain state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/biometry
func (d LADomainState) Biometry() ILADomainStateBiometry {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("biometry"))
	return LADomainStateBiometryFromID(objc.ID(rv))
}

// Contains companion domain state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/companion
func (d LADomainState) Companion() ILADomainStateCompanion {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("companion"))
	return LADomainStateCompanionFromID(objc.ID(rv))
}

// Contains combined state hash data for biometry and companion state hashes.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/stateHash
func (d LADomainState) StateHash() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stateHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}
