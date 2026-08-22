// Code generated from Apple documentation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/apple/foundation"
)

// CLBeaconMajorValue is the most significant value in a beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconMajorValue
type CLBeaconMajorValue = uint16

// CLBeaconMinorValue is the least significant value in a beacon.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBeaconMinorValue
type CLBeaconMinorValue = uint16

// CLGeocodeCompletionHandler is a block to be called when a geocoding request is complete.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLGeocodeCompletionHandler
type CLGeocodeCompletionHandler = func([]CLPlacemark, foundation.NSError)

// CLHeadingComponentValue is a type used to report magnetic differences reported by the onboard hardware.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeadingComponentValue
type CLHeadingComponentValue = float64

// CLLocationAccuracy is the accuracy of a geographical coordinate.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationAccuracy
type CLLocationAccuracy = float64

// CLLocationDegrees is a latitude or longitude value specified in degrees.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationDegrees
type CLLocationDegrees = float64

// CLLocationDirection is an azimuth that is measured in degrees relative to true north.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationDirection
type CLLocationDirection = float64

// CLLocationDirectionAccuracy is the accuracy of a compass heading.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationDirectionAccuracy
type CLLocationDirectionAccuracy = float64

// CLLocationDistance is a distance in meters from an existing location.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationDistance
type CLLocationDistance = float64

// CLLocationSpeed is the velocity (measured in meters per second) at which the device is moving.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSpeed
type CLLocationSpeed = float64

// CLLocationSpeedAccuracy is the accuracy of a speed.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSpeedAccuracy
type CLLocationSpeedAccuracy = float64
