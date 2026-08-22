// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CLBeaconIdentityConstraint] class.
var (
	_CLBeaconIdentityConstraintClass     CLBeaconIdentityConstraintClass
	_CLBeaconIdentityConstraintClassOnce sync.Once
)

func getCLBeaconIdentityConstraintClass() CLBeaconIdentityConstraintClass {
	_CLBeaconIdentityConstraintClassOnce.Do(func() {
		_CLBeaconIdentityConstraintClass = CLBeaconIdentityConstraintClass{class: objc.GetClass("CLBeaconIdentityConstraint")}
	})
	return _CLBeaconIdentityConstraintClass
}

// GetCLBeaconIdentityConstraintClass returns the class object for CLBeaconIdentityConstraint.
func GetCLBeaconIdentityConstraintClass() CLBeaconIdentityConstraintClass {
	return getCLBeaconIdentityConstraintClass()
}

type CLBeaconIdentityConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLBeaconIdentityConstraintClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLBeaconIdentityConstraintClass) Alloc() CLBeaconIdentityConstraint {
	rv := objc.Send[CLBeaconIdentityConstraint](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Identity characteristics that can match one or more beacons.
//
// # Overview
//
// A constraint specifies beacon identity characteristics. Use constraints to
// check for matching beacons by comparing the beacon’s identity
// characteristics ([CLBeacon.UUID], [CLBeacon.Major], and [CLBeacon.Minor])
// to those in the constraint.
//
// Constraints always specify a UUID value, but the major and minor values are
// optional. A beacon satisfies the constraint if all three identity
// characteristics of the beacon match the same characteristic of the
// constraint. Major and minor characteristics are wildcards if they have no
// value. A major or minor wildcard value matches any value in the beacon’s
// corresponding characteristic.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityConstraint
type CLBeaconIdentityConstraint struct {
	CLBeaconIdentityCondition
}

// CLBeaconIdentityConstraintFromID constructs a [CLBeaconIdentityConstraint] from an objc.ID.
//
// Identity characteristics that can match one or more beacons.
func CLBeaconIdentityConstraintFromID(id objc.ID) CLBeaconIdentityConstraint {
	return CLBeaconIdentityConstraint{CLBeaconIdentityCondition: CLBeaconIdentityConditionFromID(id)}
}

// NOTE: CLBeaconIdentityConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLBeaconIdentityConstraint] class.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityConstraint
type ICLBeaconIdentityConstraint interface {
	ICLBeaconIdentityCondition
}

// Init initializes the instance.
func (b CLBeaconIdentityConstraint) Init() CLBeaconIdentityConstraint {
	rv := objc.Send[CLBeaconIdentityConstraint](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CLBeaconIdentityConstraint) Autorelease() CLBeaconIdentityConstraint {
	rv := objc.Send[CLBeaconIdentityConstraint](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLBeaconIdentityConstraint creates a new CLBeaconIdentityConstraint instance.
func NewCLBeaconIdentityConstraint() CLBeaconIdentityConstraint {
	class := getCLBeaconIdentityConstraintClass()
	rv := objc.Send[CLBeaconIdentityConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new beacon identity condition with the identifier you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:
func NewBeaconIdentityConstraintWithUUID(uuid foundation.NSUUID) CLBeaconIdentityConstraint {
	instance := getCLBeaconIdentityConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:"), uuid)
	return CLBeaconIdentityConstraintFromID(rv)
}

// Creates a new beacon identity condition with the identifier and major value
// you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// major: A [CLBeaconMajorValue] to use as the beacon’s major value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:
func NewBeaconIdentityConstraintWithUUIDMajor(uuid foundation.NSUUID, major CLBeaconMajorValue) CLBeaconIdentityConstraint {
	instance := getCLBeaconIdentityConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:major:"), uuid, major)
	return CLBeaconIdentityConstraintFromID(rv)
}

// Creates a new beacon identity condition with the identifier, and major and
// minor values you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// major: A [CLBeaconMajorValue] to use as the beacon’s major value.
//
// minor: A [CLBeaconMinorValue] to use as the beacon’s minor value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:minor:
func NewBeaconIdentityConstraintWithUUIDMajorMinor(uuid foundation.NSUUID, major CLBeaconMajorValue, minor CLBeaconMinorValue) CLBeaconIdentityConstraint {
	instance := getCLBeaconIdentityConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:major:minor:"), uuid, major, minor)
	return CLBeaconIdentityConstraintFromID(rv)
}
