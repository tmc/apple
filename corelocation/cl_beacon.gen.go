// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLBeacon] class.
var (
	_CLBeaconClass     CLBeaconClass
	_CLBeaconClassOnce sync.Once
)

func getCLBeaconClass() CLBeaconClass {
	_CLBeaconClassOnce.Do(func() {
		_CLBeaconClass = CLBeaconClass{class: objc.GetClass("CLBeacon")}
	})
	return _CLBeaconClass
}

// GetCLBeaconClass returns the class object for CLBeacon.
func GetCLBeaconClass() CLBeaconClass {
	return getCLBeaconClass()
}

type CLBeaconClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLBeaconClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLBeaconClass) Alloc() CLBeacon {
	rv := objc.Send[CLBeacon](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about an observed iBeacon device and its relative distance to a
// person’s device.
//
// # Overview
//
// The [CLBeacon] class represents a beacon that was observed during beacon
// ranging. You do not create instances of this class directly. The location
// manager ([CLLocationManager]) object reports observed beacons to its
// associated delegate object.
//
// The identity of a beacon is defined by its [CLBeacon.UUID],
// [CLBeacon.Major], and [CLBeacon.Minor] properties. These values are coded
// into the beacon itself. For a more thorough description of the meaning of
// those values, see [CLBeaconRegion].
//
// # Getting the beacon identity
//
//   - [CLBeacon.UUID]: The UUID that the observed beacon transmitted.
//   - [CLBeacon.Major]: The major value that the observed beacon transmitted.
//   - [CLBeacon.Minor]: The minor value that the observed beacon transmitted.
//
// # Determining the distance to the beacon
//
//   - [CLBeacon.Proximity]: The relative distance to the beacon.
//   - [CLBeacon.Accuracy]: The accuracy of the proximity value, measured in meters from the beacon.
//   - [CLBeacon.Rssi]: The received signal strength of the beacon, measured in decibels.
//
// # Getting the observation timestamp
//
//   - [CLBeacon.Timestamp]: A timestamp representing when the beacon was observed.
//
// # Initializers
//
//   - [CLBeacon.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon
type CLBeacon struct {
	objectivec.Object
}

// CLBeaconFromID constructs a [CLBeacon] from an objc.ID.
//
// Information about an observed iBeacon device and its relative distance to a
// person’s device.
func CLBeaconFromID(id objc.ID) CLBeacon {
	return CLBeacon{objectivec.Object{ID: id}}
}

// NOTE: CLBeacon adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLBeacon] class.
//
// # Getting the beacon identity
//
//   - [ICLBeacon.UUID]: The UUID that the observed beacon transmitted.
//   - [ICLBeacon.Major]: The major value that the observed beacon transmitted.
//   - [ICLBeacon.Minor]: The minor value that the observed beacon transmitted.
//
// # Determining the distance to the beacon
//
//   - [ICLBeacon.Proximity]: The relative distance to the beacon.
//   - [ICLBeacon.Accuracy]: The accuracy of the proximity value, measured in meters from the beacon.
//   - [ICLBeacon.Rssi]: The received signal strength of the beacon, measured in decibels.
//
// # Getting the observation timestamp
//
//   - [ICLBeacon.Timestamp]: A timestamp representing when the beacon was observed.
//
// # Initializers
//
//   - [ICLBeacon.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon
type ICLBeacon interface {
	objectivec.IObject

	// Topic: Getting the beacon identity

	// The UUID that the observed beacon transmitted.
	UUID() foundation.NSUUID
	// The major value that the observed beacon transmitted.
	Major() foundation.NSNumber
	// The minor value that the observed beacon transmitted.
	Minor() foundation.NSNumber

	// Topic: Determining the distance to the beacon

	// The relative distance to the beacon.
	Proximity() CLProximity
	// The accuracy of the proximity value, measured in meters from the beacon.
	Accuracy() CLLocationAccuracy
	// The received signal strength of the beacon, measured in decibels.
	Rssi() int

	// Topic: Getting the observation timestamp

	// A timestamp representing when the beacon was observed.
	Timestamp() foundation.NSDate

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLBeacon

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b CLBeacon) Init() CLBeacon {
	rv := objc.Send[CLBeacon](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CLBeacon) Autorelease() CLBeacon {
	rv := objc.Send[CLBeacon](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLBeacon creates a new CLBeacon instance.
func NewCLBeacon() CLBeacon {
	class := getCLBeaconClass()
	rv := objc.Send[CLBeacon](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/init(coder:)
func NewBeaconWithCoder(coder foundation.INSCoder) CLBeacon {
	instance := getCLBeaconClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLBeaconFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/init(coder:)
func (b CLBeacon) InitWithCoder(coder foundation.INSCoder) CLBeacon {
	rv := objc.Send[CLBeacon](b.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (b CLBeacon) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The UUID that the observed beacon transmitted.
//
// # Discussion
//
// The UUID is the most significant beacon identity characteristic. Multiple
// beacon can transmit the same UUID.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/uuid
func (b CLBeacon) UUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// The major value that the observed beacon transmitted.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/major
func (b CLBeacon) Major() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("major"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The minor value that the observed beacon transmitted.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/minor
func (b CLBeacon) Minor() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("minor"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The relative distance to the beacon.
//
// # Discussion
//
// The value in this property gives a general sense of the relative distance
// to the beacon. Use it to quickly identify beacons that are nearer to the
// user rather than farther away.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/proximity
func (b CLBeacon) Proximity() CLProximity {
	rv := objc.Send[CLProximity](b.ID, objc.Sel("proximity"))
	return CLProximity(rv)
}

// The accuracy of the proximity value, measured in meters from the beacon.
//
// # Discussion
//
// A beacon with a smaller value for accuracy is typically nearer than a
// beacon with a larger accuracy value.
//
// Use this property to differentiate between beacons with the same proximity
// value. Do not use it to identify a precise location for the beacon.
// Accuracy values may fluctuate due to RF interference.
//
// A negative value in this property signifies that the actual accuracy could
// not be determined.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/accuracy
func (b CLBeacon) Accuracy() CLLocationAccuracy {
	rv := objc.Send[CLLocationAccuracy](b.ID, objc.Sel("accuracy"))
	return CLLocationAccuracy(rv)
}

// The received signal strength of the beacon, measured in decibels.
//
// # Discussion
//
// This value is the average signal strength of the samples received since
// Core Location last reported the range of the beacon to your app.
//
// Use this value for calibrating beacon transmission power.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/rssi
func (b CLBeacon) Rssi() int {
	rv := objc.Send[int](b.ID, objc.Sel("rssi"))
	return rv
}

// A timestamp representing when the beacon was observed.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeacon/timestamp
func (b CLBeacon) Timestamp() foundation.NSDate {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("timestamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}
