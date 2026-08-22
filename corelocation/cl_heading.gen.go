// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLHeading] class.
var (
	_CLHeadingClass     CLHeadingClass
	_CLHeadingClassOnce sync.Once
)

func getCLHeadingClass() CLHeadingClass {
	_CLHeadingClassOnce.Do(func() {
		_CLHeadingClass = CLHeadingClass{class: objc.GetClass("CLHeading")}
	})
	return _CLHeadingClass
}

// GetCLHeadingClass returns the class object for CLHeading.
func GetCLHeadingClass() CLHeadingClass {
	return getCLHeadingClass()
}

type CLHeadingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLHeadingClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLHeadingClass) Alloc() CLHeading {
	rv := objc.Send[CLHeading](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The orientation of the user’s device, relative to true or magnetic north.
//
// # Overview
//
// A [CLHeading] object contains computed values for the device’s azimuth
// (orientation) relative to true or magnetic north. It also includes the raw
// data for the three-dimensional vector used to compute those values. A
// navigation app might use the information to rotate a map so that it
// reflects the direction that the user is facing.
//
// Typically, you don’t create instances of this class yourself, nor do you
// subclass it. Instead, you receive instances of this class through the
// delegate assigned to the [CLLocationManager] object whose
// [CLLocationManager.StartUpdatingHeading] method you called.
//
// # Getting the heading values
//
//   - [CLHeading.MagneticHeading]: The heading (measured in degrees) relative to magnetic north.
//   - [CLHeading.TrueHeading]: The heading (measured in degrees) relative to true north.
//   - [CLHeading.HeadingAccuracy]: The maximum deviation (measured in degrees) between the reported heading and the true geomagnetic heading.
//
// # Getting the raw heading data
//
//   - [CLHeading.X]: The geomagnetic data (measured in microteslas) for the x-axis.
//   - [CLHeading.Y]: The geomagnetic data (measured in microteslas) for the y-axis.
//   - [CLHeading.Z]: The geomagnetic data (measured in microteslas) for the z-axis.
//
// # Getting the event timestamp
//
//   - [CLHeading.Timestamp]: The time at which this heading was determined.
//
// # Initializers
//
//   - [CLHeading.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading
type CLHeading struct {
	objectivec.Object
}

// CLHeadingFromID constructs a [CLHeading] from an objc.ID.
//
// The orientation of the user’s device, relative to true or magnetic north.
func CLHeadingFromID(id objc.ID) CLHeading {
	return CLHeading{objectivec.Object{ID: id}}
}

// NOTE: CLHeading adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLHeading] class.
//
// # Getting the heading values
//
//   - [ICLHeading.MagneticHeading]: The heading (measured in degrees) relative to magnetic north.
//   - [ICLHeading.TrueHeading]: The heading (measured in degrees) relative to true north.
//   - [ICLHeading.HeadingAccuracy]: The maximum deviation (measured in degrees) between the reported heading and the true geomagnetic heading.
//
// # Getting the raw heading data
//
//   - [ICLHeading.X]: The geomagnetic data (measured in microteslas) for the x-axis.
//   - [ICLHeading.Y]: The geomagnetic data (measured in microteslas) for the y-axis.
//   - [ICLHeading.Z]: The geomagnetic data (measured in microteslas) for the z-axis.
//
// # Getting the event timestamp
//
//   - [ICLHeading.Timestamp]: The time at which this heading was determined.
//
// # Initializers
//
//   - [ICLHeading.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading
type ICLHeading interface {
	objectivec.IObject

	// Topic: Getting the heading values

	// The heading (measured in degrees) relative to magnetic north.
	MagneticHeading() CLLocationDirection
	// The heading (measured in degrees) relative to true north.
	TrueHeading() CLLocationDirection
	// The maximum deviation (measured in degrees) between the reported heading and the true geomagnetic heading.
	HeadingAccuracy() CLLocationDirection

	// Topic: Getting the raw heading data

	// The geomagnetic data (measured in microteslas) for the x-axis.
	X() CLHeadingComponentValue
	// The geomagnetic data (measured in microteslas) for the y-axis.
	Y() CLHeadingComponentValue
	// The geomagnetic data (measured in microteslas) for the z-axis.
	Z() CLHeadingComponentValue

	// Topic: Getting the event timestamp

	// The time at which this heading was determined.
	Timestamp() foundation.NSDate

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLHeading

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (h CLHeading) Init() CLHeading {
	rv := objc.Send[CLHeading](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h CLHeading) Autorelease() CLHeading {
	rv := objc.Send[CLHeading](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLHeading creates a new CLHeading instance.
func NewCLHeading() CLHeading {
	class := getCLHeadingClass()
	rv := objc.Send[CLHeading](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/init(coder:)
func NewHeadingWithCoder(coder foundation.INSCoder) CLHeading {
	instance := getCLHeadingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLHeadingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/init(coder:)
func (h CLHeading) InitWithCoder(coder foundation.INSCoder) CLHeading {
	rv := objc.Send[CLHeading](h.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (h CLHeading) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](h.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The heading (measured in degrees) relative to magnetic north.
//
// # Discussion
//
// The value in this property represents the heading relative to the magnetic
// North Pole, which is different from the geographic North Pole. The value
// `0` means the device is pointed toward magnetic north, `90` means it is
// pointed east, `180` means it is pointed south, and so on. The value in this
// property should always be valid.
//
// In iOS 3.x and earlier, the value in this property is always measured
// relative to the top of the device in a portrait orientation, regardless of
// the device’s actual physical or interface orientation. In iOS 4.0 and
// later, the value is measured relative to the heading orientation specified
// by the location manager. For more information, see the
// [CLLocationManager.HeadingOrientation] property in [CLLocationManager].
//
// If the [CLHeading.HeadingAccuracy] property contains a negative value, the
// value in this property should be considered unreliable.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/magneticHeading
func (h CLHeading) MagneticHeading() CLLocationDirection {
	rv := objc.Send[CLLocationDirection](h.ID, objc.Sel("magneticHeading"))
	return CLLocationDirection(rv)
}

// The heading (measured in degrees) relative to true north.
//
// # Discussion
//
// The value in this property represents the heading relative to the
// geographic North Pole. The value `0` means the device is pointed toward
// true north, `90` means it is pointed due east, `180` means it is pointed
// due south, and so on. A negative value indicates that the heading could not
// be determined.
//
// In iOS 3.x and earlier, the value in this property is always measured
// relative to the top of the device in a portrait orientation, regardless of
// the device’s actual physical or interface orientation. In iOS 4.0 and
// later, the value is measured relative to the heading orientation specified
// by the location manager. For more information, see the
// [CLLocationManager.HeadingOrientation] property in [CLLocationManager].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/trueHeading
func (h CLHeading) TrueHeading() CLLocationDirection {
	rv := objc.Send[CLLocationDirection](h.ID, objc.Sel("trueHeading"))
	return CLLocationDirection(rv)
}

// The maximum deviation (measured in degrees) between the reported heading
// and the true geomagnetic heading.
//
// # Discussion
//
// A positive value in this property represents the potential error between
// the value reported by the [CLHeading.MagneticHeading] property and the
// actual direction of magnetic north. Thus, the lower the value of this
// property, the more accurate the heading. A negative value means that the
// reported heading is invalid, which can occur when the device is
// uncalibrated or there is strong interference from local magnetic fields.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/headingAccuracy
func (h CLHeading) HeadingAccuracy() CLLocationDirection {
	rv := objc.Send[CLLocationDirection](h.ID, objc.Sel("headingAccuracy"))
	return CLLocationDirection(rv)
}

// The geomagnetic data (measured in microteslas) for the x-axis.
//
// # Discussion
//
// This value represents the x-axis deviation from the magnetic field lines
// being tracked by the device.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/x
func (h CLHeading) X() CLHeadingComponentValue {
	rv := objc.Send[CLHeadingComponentValue](h.ID, objc.Sel("x"))
	return CLHeadingComponentValue(rv)
}

// The geomagnetic data (measured in microteslas) for the y-axis.
//
// # Discussion
//
// This value represents the y-axis deviation from the magnetic field lines
// being tracked by the device.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/y
func (h CLHeading) Y() CLHeadingComponentValue {
	rv := objc.Send[CLHeadingComponentValue](h.ID, objc.Sel("y"))
	return CLHeadingComponentValue(rv)
}

// The geomagnetic data (measured in microteslas) for the z-axis.
//
// # Discussion
//
// This value represents the z-axis deviation from the magnetic field lines
// being tracked by the device.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/z
func (h CLHeading) Z() CLHeadingComponentValue {
	rv := objc.Send[CLHeadingComponentValue](h.ID, objc.Sel("z"))
	return CLHeadingComponentValue(rv)
}

// The time at which this heading was determined.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLHeading/timestamp
func (h CLHeading) Timestamp() foundation.NSDate {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("timestamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}
