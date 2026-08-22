// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CLBeaconRegion] class.
var (
	_CLBeaconRegionClass     CLBeaconRegionClass
	_CLBeaconRegionClassOnce sync.Once
)

func getCLBeaconRegionClass() CLBeaconRegionClass {
	_CLBeaconRegionClassOnce.Do(func() {
		_CLBeaconRegionClass = CLBeaconRegionClass{class: objc.GetClass("CLBeaconRegion")}
	})
	return _CLBeaconRegionClass
}

// GetCLBeaconRegionClass returns the class object for CLBeaconRegion.
func GetCLBeaconRegionClass() CLBeaconRegionClass {
	return getCLBeaconRegionClass()
}

type CLBeaconRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLBeaconRegionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLBeaconRegionClass) Alloc() CLBeaconRegion {
	rv := objc.Send[CLBeaconRegion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A region for detecting the presence of iBeacon devices.
//
// # Overview
//
// A [CLBeaconRegion] object defines a region that you use to detect Bluetooth
// beacons conforming to the iBeacon specification. In contrast to a
// [CLCircularRegion] that centers on a geographic location, a
// [CLBeaconRegion] focuses on an iBeacon with specific identifying
// characteristics, which you provide. When a matching device comes in range,
// Core Location notifies your app.
//
// You monitor beacon regions in two ways. To detect when a beacon is in
// range, use the [CLLocationManager.StartMonitoringForRegion] method of your
// location manager object. After detecting a beacon, call the
// [startRangingBeacons(in:)] method to determine the relative distance to
// that beacon.
//
// When detecting an iBeacon, you need to specify the [proximityUUID],
// [major], and [minor] values that you programmed into the beacon hardware.
// You use the values to identify your beacons uniquely, and you can specify a
// subset of values to detect multiple beacons. The [proximityUUID] property
// is typically the same for all of the beacons in your installation. Use the
// [major] and [minor] values to distinguish among different beacons in your
// installation.
//
// If you want to configure the current iOS device as a Bluetooth beacon,
// create a beacon region with the appropriate identifying information. You
// can then call the [peripheralData(withMeasuredPower:)] method of the region
// to get a dictionary that you can use to advertise the device with the Core
// Bluetooth framework. For more information about using that framework to
// advertise the device as a beacon, see [Turning an iOS device into an
// iBeacon device].
//
// For information about how to detect beacons, see [Determining the proximity
// to an iBeacon device].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion
//
// [Determining the proximity to an iBeacon device]: https://developer.apple.com/documentation/CoreLocation/determining-the-proximity-to-an-ibeacon-device
// [Turning an iOS device into an iBeacon device]: https://developer.apple.com/documentation/CoreLocation/turning-an-ios-device-into-an-ibeacon-device
// [major]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/major
// [minor]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/minor
// [peripheralData(withMeasuredPower:)]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/peripheralData(withMeasuredPower:)
// [proximityUUID]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/proximityUUID
// [startRangingBeacons(in:)]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startRangingBeacons(in:)
type CLBeaconRegion struct {
	CLRegion
}

// CLBeaconRegionFromID constructs a [CLBeaconRegion] from an objc.ID.
//
// A region for detecting the presence of iBeacon devices.
func CLBeaconRegionFromID(id objc.ID) CLBeaconRegion {
	return CLBeaconRegion{CLRegion: CLRegionFromID(id)}
}

// NOTE: CLBeaconRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLBeaconRegion] class.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion
type ICLBeaconRegion interface {
	ICLRegion
}

// Init initializes the instance.
func (b CLBeaconRegion) Init() CLBeaconRegion {
	rv := objc.Send[CLBeaconRegion](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CLBeaconRegion) Autorelease() CLBeaconRegion {
	rv := objc.Send[CLBeaconRegion](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLBeaconRegion creates a new CLBeaconRegion instance.
func NewCLBeaconRegion() CLBeaconRegion {
	class := getCLBeaconRegionClass()
	rv := objc.Send[CLBeaconRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(coder:)
func NewBeaconRegionWithCoder(coder foundation.INSCoder) CLBeaconRegion {
	instance := getCLBeaconRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLBeaconRegionFromID(rv)
}
