// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CLBeaconIdentityCondition] class.
var (
	_CLBeaconIdentityConditionClass     CLBeaconIdentityConditionClass
	_CLBeaconIdentityConditionClassOnce sync.Once
)

func getCLBeaconIdentityConditionClass() CLBeaconIdentityConditionClass {
	_CLBeaconIdentityConditionClassOnce.Do(func() {
		_CLBeaconIdentityConditionClass = CLBeaconIdentityConditionClass{class: objc.GetClass("CLBeaconIdentityCondition")}
	})
	return _CLBeaconIdentityConditionClass
}

// GetCLBeaconIdentityConditionClass returns the class object for CLBeaconIdentityCondition.
func GetCLBeaconIdentityConditionClass() CLBeaconIdentityConditionClass {
	return getCLBeaconIdentityConditionClass()
}

type CLBeaconIdentityConditionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLBeaconIdentityConditionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLBeaconIdentityConditionClass) Alloc() CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A condition that describes the identity characteristics of a beacon.
//
// # Overview
//
// Core Location defines a beacon identity by UUID, and major and minor
// values. You need to specify the UUID. If you only specify a UUID, the
// framework treats the major and minor values as wildcards and any beacons
// with the same UUID satisfy the condition. Similarly, if you specify only a
// UUID and a major value, the framework treats the minor value as a wildcard
// and any beacons with the same UUID and major value satisfy the condition.
//
// # Creating beacon identity conditions
//
//   - [CLBeaconIdentityCondition.InitWithUUID]: Creates a new beacon identity condition with the identifier you specify.
//   - [CLBeaconIdentityCondition.InitWithUUIDMajor]: Creates a new beacon identity condition with the identifier and major value you specify.
//   - [CLBeaconIdentityCondition.InitWithUUIDMajorMinor]: Creates a new beacon identity condition with the identifier, and major and minor values you specify.
//
// # Accessing the beacon’s properties
//
//   - [CLBeaconIdentityCondition.UUID]: A universally unique identifier that represent the beacon’s identifier.
//   - [CLBeaconIdentityCondition.Major]: The most significant value associated with the beacon.
//   - [CLBeaconIdentityCondition.Minor]: The least significant value associated with the beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition
type CLBeaconIdentityCondition struct {
	CLCondition
}

// CLBeaconIdentityConditionFromID constructs a [CLBeaconIdentityCondition] from an objc.ID.
//
// A condition that describes the identity characteristics of a beacon.
func CLBeaconIdentityConditionFromID(id objc.ID) CLBeaconIdentityCondition {
	return CLBeaconIdentityCondition{CLCondition: CLConditionFromID(id)}
}

// NOTE: CLBeaconIdentityCondition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLBeaconIdentityCondition] class.
//
// # Creating beacon identity conditions
//
//   - [ICLBeaconIdentityCondition.InitWithUUID]: Creates a new beacon identity condition with the identifier you specify.
//   - [ICLBeaconIdentityCondition.InitWithUUIDMajor]: Creates a new beacon identity condition with the identifier and major value you specify.
//   - [ICLBeaconIdentityCondition.InitWithUUIDMajorMinor]: Creates a new beacon identity condition with the identifier, and major and minor values you specify.
//
// # Accessing the beacon’s properties
//
//   - [ICLBeaconIdentityCondition.UUID]: A universally unique identifier that represent the beacon’s identifier.
//   - [ICLBeaconIdentityCondition.Major]: The most significant value associated with the beacon.
//   - [ICLBeaconIdentityCondition.Minor]: The least significant value associated with the beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition
type ICLBeaconIdentityCondition interface {
	ICLCondition

	// Topic: Creating beacon identity conditions

	// Creates a new beacon identity condition with the identifier you specify.
	InitWithUUID(uuid foundation.NSUUID) CLBeaconIdentityCondition
	// Creates a new beacon identity condition with the identifier and major value you specify.
	InitWithUUIDMajor(uuid foundation.NSUUID, major CLBeaconMajorValue) CLBeaconIdentityCondition
	// Creates a new beacon identity condition with the identifier, and major and minor values you specify.
	InitWithUUIDMajorMinor(uuid foundation.NSUUID, major CLBeaconMajorValue, minor CLBeaconMinorValue) CLBeaconIdentityCondition

	// Topic: Accessing the beacon’s properties

	// A universally unique identifier that represent the beacon’s identifier.
	UUID() foundation.NSUUID
	// The most significant value associated with the beacon.
	Major() foundation.NSNumber
	// The least significant value associated with the beacon.
	Minor() foundation.NSNumber
}

// Init initializes the instance.
func (b CLBeaconIdentityCondition) Init() CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CLBeaconIdentityCondition) Autorelease() CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLBeaconIdentityCondition creates a new CLBeaconIdentityCondition instance.
func NewCLBeaconIdentityCondition() CLBeaconIdentityCondition {
	class := getCLBeaconIdentityConditionClass()
	rv := objc.Send[CLBeaconIdentityCondition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new beacon identity condition with the identifier you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:
func NewBeaconIdentityConditionWithUUID(uuid foundation.NSUUID) CLBeaconIdentityCondition {
	instance := getCLBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:"), uuid)
	return CLBeaconIdentityConditionFromID(rv)
}

// Creates a new beacon identity condition with the identifier and major value
// you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// major: A [CLBeaconMajorValue] to use as the beacon’s major value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:
func NewBeaconIdentityConditionWithUUIDMajor(uuid foundation.NSUUID, major CLBeaconMajorValue) CLBeaconIdentityCondition {
	instance := getCLBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:major:"), uuid, major)
	return CLBeaconIdentityConditionFromID(rv)
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
func NewBeaconIdentityConditionWithUUIDMajorMinor(uuid foundation.NSUUID, major CLBeaconMajorValue, minor CLBeaconMinorValue) CLBeaconIdentityCondition {
	instance := getCLBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:major:minor:"), uuid, major, minor)
	return CLBeaconIdentityConditionFromID(rv)
}

// Creates a new beacon identity condition with the identifier you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:
func (b CLBeaconIdentityCondition) InitWithUUID(uuid foundation.NSUUID) CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](b.ID, objc.Sel("initWithUUID:"), uuid)
	return rv
}

// Creates a new beacon identity condition with the identifier and major value
// you specify.
//
// uuid: A [CLBeaconIdentityCondition.UUID] to use as the beacon’s identifier.
//
// major: A [CLBeaconMajorValue] to use as the beacon’s major value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:
func (b CLBeaconIdentityCondition) InitWithUUIDMajor(uuid foundation.NSUUID, major CLBeaconMajorValue) CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](b.ID, objc.Sel("initWithUUID:major:"), uuid, major)
	return rv
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
func (b CLBeaconIdentityCondition) InitWithUUIDMajorMinor(uuid foundation.NSUUID, major CLBeaconMajorValue, minor CLBeaconMinorValue) CLBeaconIdentityCondition {
	rv := objc.Send[CLBeaconIdentityCondition](b.ID, objc.Sel("initWithUUID:major:minor:"), uuid, major, minor)
	return rv
}

// A universally unique identifier that represent the beacon’s identifier.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/UUID
func (b CLBeaconIdentityCondition) UUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// The most significant value associated with the beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/major
func (b CLBeaconIdentityCondition) Major() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("major"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The least significant value associated with the beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/minor
func (b CLBeaconIdentityCondition) Minor() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("minor"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
