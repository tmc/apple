// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LADomainStateCompanion] class.
var (
	_LADomainStateCompanionClass     LADomainStateCompanionClass
	_LADomainStateCompanionClassOnce sync.Once
)

func getLADomainStateCompanionClass() LADomainStateCompanionClass {
	_LADomainStateCompanionClassOnce.Do(func() {
		_LADomainStateCompanionClass = LADomainStateCompanionClass{class: objc.GetClass("LADomainStateCompanion")}
	})
	return _LADomainStateCompanionClass
}

// GetLADomainStateCompanionClass returns the class object for LADomainStateCompanion.
func GetLADomainStateCompanionClass() LADomainStateCompanionClass {
	return getLADomainStateCompanionClass()
}

type LADomainStateCompanionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LADomainStateCompanionClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LADomainStateCompanionClass) Alloc() LADomainStateCompanion {
	rv := objc.Send[LADomainStateCompanion](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [LADomainStateCompanion.StateHash]: Contains combined state hash data for all available companion types. . Returns `nil` if no companion devices are paired.
//
// # Instance Methods
//
//   - [LADomainStateCompanion.StateHashForCompanionType]: Returns state hash data for the given companion type.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion
type LADomainStateCompanion struct {
	objectivec.Object
}

// LADomainStateCompanionFromID constructs a [LADomainStateCompanion] from an objc.ID.
func LADomainStateCompanionFromID(id objc.ID) LADomainStateCompanion {
	return LADomainStateCompanion{objectivec.Object{ID: id}}
}

// NOTE: LADomainStateCompanion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LADomainStateCompanion] class.
//
// # Instance Properties
//
//   - [ILADomainStateCompanion.StateHash]: Contains combined state hash data for all available companion types. . Returns `nil` if no companion devices are paired.
//
// # Instance Methods
//
//   - [ILADomainStateCompanion.StateHashForCompanionType]: Returns state hash data for the given companion type.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion
type ILADomainStateCompanion interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Contains combined state hash data for all available companion types. . Returns `nil` if no companion devices are paired.
	StateHash() foundation.INSData

	// Topic: Instance Methods

	// Returns state hash data for the given companion type.
	StateHashForCompanionType(companionType LACompanionType) foundation.INSData

	// Indicates types of companions paired with the device. The elements are NSNumber-wrapped instances of @c [LACompanionType].
	AvailableCompanionTypes() foundation.INSSet
}

// Init initializes the instance.
func (d LADomainStateCompanion) Init() LADomainStateCompanion {
	rv := objc.Send[LADomainStateCompanion](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d LADomainStateCompanion) Autorelease() LADomainStateCompanion {
	rv := objc.Send[LADomainStateCompanion](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewLADomainStateCompanion creates a new LADomainStateCompanion instance.
func NewLADomainStateCompanion() LADomainStateCompanion {
	class := getLADomainStateCompanionClass()
	rv := objc.Send[LADomainStateCompanion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns state hash data for the given companion type.
//
// companionType: The companion type for which state hash data should be returned.
//
// # Discussion
//
// If database of paired devices of the given type was modified state hash
// data will change. Nature of such database changes cannot be determined but
// comparing data of state hash after different policy evaluation will reveal
// the fact database was changed between calls.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash(for:)
func (d LADomainStateCompanion) StateHashForCompanionType(companionType LACompanionType) foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stateHashForCompanionType:"), companionType)
	return foundation.NSDataFromID(rv)
}

// Contains combined state hash data for all available companion types. .
// Returns `nil` if no companion devices are paired.
//
// # Discussion
//
// As long as database of paired companion devices doesn’t change,
// `stateHash` stays the same for the same set of `availableCompanions`.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash
func (d LADomainStateCompanion) StateHash() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stateHash"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// Indicates types of companions paired with the device. The elements are
// NSNumber-wrapped instances of @c [LACompanionType].
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/availableCompanionTypes-1ggnh
func (d LADomainStateCompanion) AvailableCompanionTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("availableCompanionTypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}
