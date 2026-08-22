// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLLocation] class.
var (
	_CLLocationClass     CLLocationClass
	_CLLocationClassOnce sync.Once
)

func getCLLocationClass() CLLocationClass {
	_CLLocationClassOnce.Do(func() {
		_CLLocationClass = CLLocationClass{class: objc.GetClass("CLLocation")}
	})
	return _CLLocationClass
}

// GetCLLocationClass returns the class object for CLLocation.
func GetCLLocationClass() CLLocationClass {
	return getCLLocationClass()
}

type CLLocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLLocationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLLocationClass) Alloc() CLLocation {
	rv := objc.Send[CLLocation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The latitude, longitude, and course information reported by the system.
//
// # Overview
//
// A [CLLocation] object contains the geographical location and altitude of a
// device, along with values indicating the accuracy of those measurements and
// when they were collected. In iOS, a location object also contains course
// information — that is, the speed and heading in which the device was
// moving.
//
// Typically, you don’t create location objects yourself. After you request
// location updates from your [CLLocationManager] object, the system uses
// onboard sensors to gather location data and report that data to your app.
// Some services also return previously collected location data, which you can
// use as context to improve your services. You can always retrieve the most
// recently collected location from the [CLLocationManager.Location] property
// of your [CLLocationManager] object. You may create location objects
// yourself when you want to cache custom location data or calculate the
// distance between two geographical coordinates.
//
// Use [CLLocation] objects as-is, and don’t subclass them.
//
// # Creating a location object
//
//   - [CLLocation.InitWithLatitudeLongitude]: Creates a location object with the specified latitude and longitude.
//   - [CLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp]: Creates a location object with the specified coordinate and altitude information.
//   - [CLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseSpeedTimestamp]: Creates a location object with the specified coordinate, altitude, and course information.
//   - [CLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestamp]: Creates a location object with the specified coordinate, altitude, course, and accuracy information.
//   - [CLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestampSourceInfo]
//
// # Getting the location attributes
//
//   - [CLLocation.Coordinate]: The geographical coordinate information.
//   - [CLLocation.Altitude]: The altitude above mean sea level associated with a location, measured in meters.
//   - [CLLocation.EllipsoidalAltitude]: The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
//   - [CLLocation.Floor]: The logical floor of the building in which the user is located.
//   - [CLLocation.Timestamp]: The time at which this location was determined.
//   - [CLLocation.SourceInformation]: Information about the source that provides the location.
//
// # Getting the location accuracy
//
//   - [CLLocation.HorizontalAccuracy]: The radius of uncertainty for the location, measured in meters.
//   - [CLLocation.VerticalAccuracy]: The validity of the altitude values, and their estimated uncertainty, measured in meters.
//
// # Measuring the distance between coordinates
//
//   - [CLLocation.DistanceFromLocation]: Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// # Getting speed and course information
//
//   - [CLLocation.Speed]: The instantaneous speed of the device, measured in meters per second.
//   - [CLLocation.SpeedAccuracy]: The accuracy of the speed value, measured in meters per second.
//   - [CLLocation.Course]: The direction in which the device is traveling, measured in degrees and relative to due north.
//   - [CLLocation.CourseAccuracy]: The accuracy of the course value, measured in degrees.
//
// # Initializers
//
//   - [CLLocation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation
type CLLocation struct {
	objectivec.Object
}

// CLLocationFromID constructs a [CLLocation] from an objc.ID.
//
// The latitude, longitude, and course information reported by the system.
func CLLocationFromID(id objc.ID) CLLocation {
	return CLLocation{objectivec.Object{ID: id}}
}

// NOTE: CLLocation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLLocation] class.
//
// # Creating a location object
//
//   - [ICLLocation.InitWithLatitudeLongitude]: Creates a location object with the specified latitude and longitude.
//   - [ICLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp]: Creates a location object with the specified coordinate and altitude information.
//   - [ICLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseSpeedTimestamp]: Creates a location object with the specified coordinate, altitude, and course information.
//   - [ICLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestamp]: Creates a location object with the specified coordinate, altitude, course, and accuracy information.
//   - [ICLLocation.InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestampSourceInfo]
//
// # Getting the location attributes
//
//   - [ICLLocation.Coordinate]: The geographical coordinate information.
//   - [ICLLocation.Altitude]: The altitude above mean sea level associated with a location, measured in meters.
//   - [ICLLocation.EllipsoidalAltitude]: The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
//   - [ICLLocation.Floor]: The logical floor of the building in which the user is located.
//   - [ICLLocation.Timestamp]: The time at which this location was determined.
//   - [ICLLocation.SourceInformation]: Information about the source that provides the location.
//
// # Getting the location accuracy
//
//   - [ICLLocation.HorizontalAccuracy]: The radius of uncertainty for the location, measured in meters.
//   - [ICLLocation.VerticalAccuracy]: The validity of the altitude values, and their estimated uncertainty, measured in meters.
//
// # Measuring the distance between coordinates
//
//   - [ICLLocation.DistanceFromLocation]: Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// # Getting speed and course information
//
//   - [ICLLocation.Speed]: The instantaneous speed of the device, measured in meters per second.
//   - [ICLLocation.SpeedAccuracy]: The accuracy of the speed value, measured in meters per second.
//   - [ICLLocation.Course]: The direction in which the device is traveling, measured in degrees and relative to due north.
//   - [ICLLocation.CourseAccuracy]: The accuracy of the course value, measured in degrees.
//
// # Initializers
//
//   - [ICLLocation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation
type ICLLocation interface {
	objectivec.IObject

	// Topic: Creating a location object

	// Creates a location object with the specified latitude and longitude.
	InitWithLatitudeLongitude(latitude CLLocationDegrees, longitude CLLocationDegrees) CLLocation
	// Creates a location object with the specified coordinate and altitude information.
	InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, timestamp foundation.NSDate) CLLocation
	// Creates a location object with the specified coordinate, altitude, and course information.
	InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseSpeedTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, speed CLLocationSpeed, timestamp foundation.NSDate) CLLocation
	// Creates a location object with the specified coordinate, altitude, course, and accuracy information.
	InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate) CLLocation
	InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestampSourceInfo(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate, sourceInfo ICLLocationSourceInformation) CLLocation

	// Topic: Getting the location attributes

	// The geographical coordinate information.
	Coordinate() CLLocationCoordinate2D
	// The altitude above mean sea level associated with a location, measured in meters.
	Altitude() CLLocationDistance
	// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
	EllipsoidalAltitude() CLLocationDistance
	// The logical floor of the building in which the user is located.
	Floor() ICLFloor
	// The time at which this location was determined.
	Timestamp() foundation.NSDate
	// Information about the source that provides the location.
	SourceInformation() ICLLocationSourceInformation

	// Topic: Getting the location accuracy

	// The radius of uncertainty for the location, measured in meters.
	HorizontalAccuracy() CLLocationAccuracy
	// The validity of the altitude values, and their estimated uncertainty, measured in meters.
	VerticalAccuracy() CLLocationAccuracy

	// Topic: Measuring the distance between coordinates

	// Returns the distance (measured in meters) from the current object’s location to the specified location.
	DistanceFromLocation(location ICLLocation) CLLocationDistance

	// Topic: Getting speed and course information

	// The instantaneous speed of the device, measured in meters per second.
	Speed() CLLocationSpeed
	// The accuracy of the speed value, measured in meters per second.
	SpeedAccuracy() CLLocationSpeedAccuracy
	// The direction in which the device is traveling, measured in degrees and relative to due north.
	Course() CLLocationDirection
	// The accuracy of the course value, measured in degrees.
	CourseAccuracy() CLLocationDirectionAccuracy

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLLocation

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (l CLLocation) Init() CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l CLLocation) Autorelease() CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLLocation creates a new CLLocation instance.
func NewCLLocation() CLLocation {
	class := getCLLocationClass()
	rv := objc.Send[CLLocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coder:)
func NewLocationWithCoder(coder foundation.INSCoder) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLLocationFromID(rv)
}

// Creates a location object with the specified coordinate, altitude, course,
// and accuracy information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// course: The direction of travel for the location, measured in degrees relative to
// due north and continuing clockwise around the compass.
//
// courseAccuracy: The accuracy of the course value, measured in degrees. Specify a negative
// number to indicate that the course is invalid.
//
// speed: The current speed associated with this location, measured in meters per
// second.
//
// speedAccuracy: The accuracy of the speed value, measured in meters per second. Specify a
// negative number to indicate that the speed is invalid.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:)
func NewLocationWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, course, courseAccuracy, speed, speedAccuracy, timestamp)
	return CLLocationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:sourceInfo:)
func NewLocationWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestampSourceInfo(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate, sourceInfo ICLLocationSourceInformation) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:sourceInfo:"), coordinate, altitude, hAccuracy, vAccuracy, course, courseAccuracy, speed, speedAccuracy, timestamp, sourceInfo)
	return CLLocationFromID(rv)
}

// Creates a location object with the specified coordinate, altitude, and
// course information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// course: The direction of travel for the location, measured in degrees relative to
// due north and continuing clockwise around the compass.
//
// speed: The current speed associated with this location, measured in meters per
// second.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate,
// altitude, and course information.
//
// # Discussion
//
// Use this method to create location objects that aren’t necessarily based
// on the user’s current location.Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:speed:timestamp:)
func NewLocationWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseSpeedTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, speed CLLocationSpeed, timestamp foundation.NSDate) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:speed:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, course, speed, timestamp)
	return CLLocationFromID(rv)
}

// Creates a location object with the specified coordinate and altitude
// information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate
// and altitude information.
//
// # Discussion
//
// Use this method to create location objects that are not necessarily based
// on the user’s current location.Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// This method records the values you provide, and it initializes other
// properties to appropriate default values. Specifically, this method sets
// the speed and course values to `-1`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:timestamp:)
func NewLocationWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, timestamp foundation.NSDate) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, timestamp)
	return CLLocationFromID(rv)
}

// Creates a location object with the specified latitude and longitude.
//
// latitude: The latitude of the geographical coordinate.
//
// longitude: The longitude of the geographical coordinate.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate.
//
// # Discussion
//
// Use this method to create location objects that are not necessarily based
// on the user’s current location. Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// This method records the latitude and longitude values you provide, and it
// initializes other properties to appropriate default values. Specifically,
// this method sets the [CLLocation.Altitude] and
// [CLLocation.HorizontalAccuracy] properties to 0, sets the
// [CLLocation.VerticalAccuracy] property to `-1` to indicate that the
// altitude is invalid, sets the [CLLocation.Speed] and [CLLocation.Course]
// values to `-1`, and sets the [CLLocation.Timestamp] property to the time at
// which the returned object was created.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(latitude:longitude:)
func NewLocationWithLatitudeLongitude(latitude CLLocationDegrees, longitude CLLocationDegrees) CLLocation {
	instance := getCLLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLatitude:longitude:"), latitude, longitude)
	return CLLocationFromID(rv)
}

// Creates a location object with the specified latitude and longitude.
//
// latitude: The latitude of the geographical coordinate.
//
// longitude: The longitude of the geographical coordinate.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate.
//
// # Discussion
//
// Use this method to create location objects that are not necessarily based
// on the user’s current location. Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// This method records the latitude and longitude values you provide, and it
// initializes other properties to appropriate default values. Specifically,
// this method sets the [CLLocation.Altitude] and
// [CLLocation.HorizontalAccuracy] properties to 0, sets the
// [CLLocation.VerticalAccuracy] property to `-1` to indicate that the
// altitude is invalid, sets the [CLLocation.Speed] and [CLLocation.Course]
// values to `-1`, and sets the [CLLocation.Timestamp] property to the time at
// which the returned object was created.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(latitude:longitude:)
func (l CLLocation) InitWithLatitudeLongitude(latitude CLLocationDegrees, longitude CLLocationDegrees) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithLatitude:longitude:"), latitude, longitude)
	return rv
}

// Creates a location object with the specified coordinate and altitude
// information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate
// and altitude information.
//
// # Discussion
//
// Use this method to create location objects that are not necessarily based
// on the user’s current location.Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// This method records the values you provide, and it initializes other
// properties to appropriate default values. Specifically, this method sets
// the speed and course values to `-1`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:timestamp:)
func (l CLLocation) InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, timestamp foundation.NSDate) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, timestamp)
	return rv
}

// Creates a location object with the specified coordinate, altitude, and
// course information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// course: The direction of travel for the location, measured in degrees relative to
// due north and continuing clockwise around the compass.
//
// speed: The current speed associated with this location, measured in meters per
// second.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// # Return Value
//
// A location object initialized with the specified geographical coordinate,
// altitude, and course information.
//
// # Discussion
//
// Use this method to create location objects that aren’t necessarily based
// on the user’s current location.Typically, you acquire location objects
// from your [CLLocationManager] object, which returns the user’s actual
// location. However, you might use this method when you want to represent any
// location on a map. For example, you might create an object to represent the
// user’s intended destination.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:speed:timestamp:)
func (l CLLocation) InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseSpeedTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, speed CLLocationSpeed, timestamp foundation.NSDate) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:speed:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, course, speed, timestamp)
	return rv
}

// Creates a location object with the specified coordinate, altitude, course,
// and accuracy information.
//
// coordinate: A coordinate structure containing the latitude and longitude values.
//
// altitude: The altitude value for the location.
//
// hAccuracy: The radius of uncertainty for the geographical coordinate, measured in
// meters. Specify a negative number to indicate that the geographical
// coordinate is invalid.
//
// vAccuracy: The accuracy of the altitude value, measured in meters. Specify a negative
// number to indicate that the altitude is invalid.
//
// course: The direction of travel for the location, measured in degrees relative to
// due north and continuing clockwise around the compass.
//
// courseAccuracy: The accuracy of the course value, measured in degrees. Specify a negative
// number to indicate that the course is invalid.
//
// speed: The current speed associated with this location, measured in meters per
// second.
//
// speedAccuracy: The accuracy of the speed value, measured in meters per second. Specify a
// negative number to indicate that the speed is invalid.
//
// timestamp: The time to associate with the location object. Typically, you specify the
// current time.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:)
func (l CLLocation) InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestamp(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, course, courseAccuracy, speed, speedAccuracy, timestamp)
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:sourceInfo:)
func (l CLLocation) InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyCourseCourseAccuracySpeedSpeedAccuracyTimestampSourceInfo(coordinate CLLocationCoordinate2D, altitude CLLocationDistance, hAccuracy CLLocationAccuracy, vAccuracy CLLocationAccuracy, course CLLocationDirection, courseAccuracy CLLocationDirectionAccuracy, speed CLLocationSpeed, speedAccuracy CLLocationSpeedAccuracy, timestamp foundation.NSDate, sourceInfo ICLLocationSourceInformation) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:course:courseAccuracy:speed:speedAccuracy:timestamp:sourceInfo:"), coordinate, altitude, hAccuracy, vAccuracy, course, courseAccuracy, speed, speedAccuracy, timestamp, sourceInfo)
	return rv
}

// Returns the distance (measured in meters) from the current object’s
// location to the specified location.
//
// location: The destination location.
//
// # Return Value
//
// The distance (in meters) between the two locations.
//
// # Discussion
//
// This method measures the distance between the location in the current
// object and the value in the `location` parameter. The distance is
// calculated by tracing a line between the two points that follows the
// curvature of the Earth, and measuring the length of the resulting arc. The
// arc is a smooth curve that doesn’t take into account altitude changes
// between the two locations.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/distance(from:)
func (l CLLocation) DistanceFromLocation(location ICLLocation) CLLocationDistance {
	rv := objc.Send[CLLocationDistance](l.ID, objc.Sel("distanceFromLocation:"), location)
	return CLLocationDistance(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/init(coder:)
func (l CLLocation) InitWithCoder(coder foundation.INSCoder) CLLocation {
	rv := objc.Send[CLLocation](l.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (l CLLocation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The geographical coordinate information.
//
// # Discussion
//
// When running in the simulator, Core Location uses the values provided to it
// by the simulator. You must run your application on an iOS-based device to
// get the actual location of that device.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/coordinate
func (l CLLocation) Coordinate() CLLocationCoordinate2D {
	rv := objc.Send[CLLocationCoordinate2D](l.ID, objc.Sel("coordinate"))
	return CLLocationCoordinate2D(rv)
}

// The altitude above mean sea level associated with a location, measured in
// meters.
//
// # Discussion
//
// The [CLLocation.Altitude] property represents an orthometric height, which
// is the height above the approximate mean sea level. Positive values
// indicate altitudes above mean sea level. Negative values indicate altitudes
// below mean sea level.
//
// When [CLLocation.VerticalAccuracy] contains `0` or a negative number, the
// value of [CLLocation.Altitude] is invalid. The value of
// [CLLocation.Altitude] is valid when [CLLocation.VerticalAccuracy] contains
// a postive number.
//
// In most cases, Core Location approximates mean sea level using the Earth
// Gravitational Model 2008 (EGM 2008) geoid associated with the World
// Geodetic System 1984 (WGS84) standard. In some rare cases, Core Location
// approximates mean sea level using the DMA 10x10 geoid grid. The discrepancy
// between these two geoids is typically less than 5 meters.
//
// See [CLLocation.EllipsoidalAltitude] if your application uses an altitude
// with respect to the WGS84 reference frame.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/altitude
func (l CLLocation) Altitude() CLLocationDistance {
	rv := objc.Send[CLLocationDistance](l.ID, objc.Sel("altitude"))
	return CLLocationDistance(rv)
}

// The altitude as a height above the World Geodetic System 1984 (WGS84)
// ellipsoid, measured in meters.
//
// # Discussion
//
// The [CLLocation.EllipsoidalAltitude] property represents the altitude above
// the WGS84 ellipsoid associated with a location. Use the
// [CLLocation.EllipsoidalAltitude] property when your geodetic application
// needs an altitude with respect to a standard reference frame. Use
// [CLLocation.Altitude] if your application needs an altitude with respect to
// the approximate mean sea level.
//
// The [CLLocation.EllipsoidalAltitude] value is valid if
// [CLLocation.VerticalAccuracy] is greater than `0`, and invalid if
// [CLLocation.VerticalAccuracy] is `0` or below. If
// [CLLocation.VerticalAccuracy] is `0` or below,
// [CLLocation.EllipsoidalAltitude] is invalid and contains the value `0.0`.
//
// Valid values for [CLLocation.EllipsoidalAltitude] can be positive or
// negative. Positive values indicate altitudes above the ellipsoid. Negative
// values indicate altitudes below the ellipsoid.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/ellipsoidalAltitude
func (l CLLocation) EllipsoidalAltitude() CLLocationDistance {
	rv := objc.Send[CLLocationDistance](l.ID, objc.Sel("ellipsoidalAltitude"))
	return CLLocationDistance(rv)
}

// The logical floor of the building in which the user is located.
//
// # Discussion
//
// If floor information is not available for the current location, the value
// of this property is `nil`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/floor
func (l CLLocation) Floor() ICLFloor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("floor"))
	return CLFloorFromID(objc.ID(rv))
}

// The time at which this location was determined.
//
// # Discussion
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/timestamp
func (l CLLocation) Timestamp() foundation.NSDate {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("timestamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// Information about the source that provides the location.
//
// # Discussion
//
// This property enables developers to make better-informed decisions as to
// whether to treat certain locations differently, or reject potentially
// simulated locations that they generate during testing. An app may choose to
// check this property and reject locations if, for example, the
// [CLLocationSourceInformation.IsSimulatedBySoftware] property is `true` when
// the developer isn’t debugging or testing the app.
//
// Use the [CLLocation.SourceInformation] property when knowing the true
// location of the device (within a tolerance for estimation error and
// horizontal/vertical accuracy) is critical.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/sourceInformation
func (l CLLocation) SourceInformation() ICLLocationSourceInformation {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("sourceInformation"))
	return CLLocationSourceInformationFromID(objc.ID(rv))
}

// The radius of uncertainty for the location, measured in meters.
//
// # Discussion
//
// The location’s latitude and longitude identify the center of the circle,
// and this value indicates the radius of that circle. A negative value
// indicates that the latitude and longitude are invalid.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/horizontalAccuracy
func (l CLLocation) HorizontalAccuracy() CLLocationAccuracy {
	rv := objc.Send[CLLocationAccuracy](l.ID, objc.Sel("horizontalAccuracy"))
	return CLLocationAccuracy(rv)
}

// The validity of the altitude values, and their estimated uncertainty,
// measured in meters.
//
// # Discussion
//
// A positive [CLLocation.VerticalAccuracy] value represents the estimated
// uncertainty associated with [CLLocation.Altitude] and
// [CLLocation.EllipsoidalAltitude]. This value is available whenever altitude
// values are available.
//
// If [CLLocation.VerticalAccuracy] is `0` or a negative number,
// [CLLocation.Altitude] and [CLLocation.EllipsoidalAltitude] values are
// invalid. If [CLLocation.VerticalAccuracy] is a postive number,
// [CLLocation.Altitude] and [CLLocation.EllipsoidalAltitude] values are
// valid.
//
// A positive [CLLocation.VerticalAccuracy] value represents an uncertainty
// that’s approximately 68 percent, or one standard deviation, above and
// below the altitude values.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/verticalAccuracy
func (l CLLocation) VerticalAccuracy() CLLocationAccuracy {
	rv := objc.Send[CLLocationAccuracy](l.ID, objc.Sel("verticalAccuracy"))
	return CLLocationAccuracy(rv)
}

// The instantaneous speed of the device, measured in meters per second.
//
// # Discussion
//
// This value reflects the instantaneous speed of the device as it moves in
// the direction of its current heading. A negative value indicates an invalid
// speed. Because the actual speed can change many times between the delivery
// of location events, use this property for informational purposes only.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/speed
func (l CLLocation) Speed() CLLocationSpeed {
	rv := objc.Send[CLLocationSpeed](l.ID, objc.Sel("speed"))
	return CLLocationSpeed(rv)
}

// The accuracy of the speed value, measured in meters per second.
//
// # Discussion
//
// When this property contains `0` or a positive number, the value in the
// speed property is plus or minus the specified number of meters per second.
// When this property contains a negative number, the value in the speed
// property is invalid.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/speedAccuracy
func (l CLLocation) SpeedAccuracy() CLLocationSpeedAccuracy {
	rv := objc.Send[CLLocationSpeedAccuracy](l.ID, objc.Sel("speedAccuracy"))
	return CLLocationSpeedAccuracy(rv)
}

// The direction in which the device is traveling, measured in degrees and
// relative to due north.
//
// # Discussion
//
// Course values are measured in degrees starting at due north and continue
// clockwise around the compass. Thus, north is 0 degrees, east is 90 degrees,
// south is 180 degrees, and so on. Course values may not be available on all
// devices. A negative value indicates that the course information is invalid.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/course
func (l CLLocation) Course() CLLocationDirection {
	rv := objc.Send[CLLocationDirection](l.ID, objc.Sel("course"))
	return CLLocationDirection(rv)
}

// The accuracy of the course value, measured in degrees.
//
// # Discussion
//
// When this property contains `0` or a positive number, the value in the
// course property is plus or minus the specified number degrees, modulo 360.
// When this property contains a negative number, the value in the course
// property is invalid.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocation/courseAccuracy
func (l CLLocation) CourseAccuracy() CLLocationDirectionAccuracy {
	rv := objc.Send[CLLocationDirectionAccuracy](l.ID, objc.Sel("courseAccuracy"))
	return CLLocationDirectionAccuracy(rv)
}
