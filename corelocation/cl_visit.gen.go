// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLVisit] class.
var (
	_CLVisitClass     CLVisitClass
	_CLVisitClassOnce sync.Once
)

func getCLVisitClass() CLVisitClass {
	_CLVisitClassOnce.Do(func() {
		_CLVisitClass = CLVisitClass{class: objc.GetClass("CLVisit")}
	})
	return _CLVisitClass
}

// GetCLVisitClass returns the class object for CLVisit.
func GetCLVisitClass() CLVisitClass {
	return getCLVisitClass()
}

type CLVisitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLVisitClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLVisitClass) Alloc() CLVisit {
	rv := objc.Send[CLVisit](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about the user’s location during a specific period of time.
//
// # Overview
//
// A [CLVisit] object encapsulates information about places that the user has
// been. Visit objects are created by the system and delivered by the
// [CLLocationManager] object to its delegate after you start the delivery of
// events. The visit includes the location where the visit occurred and
// information about the arrival and departure times as relevant. You do not
// create visit objects directly, nor should you subclass [CLVisit].
//
// Visit objects contain as much information about the visit as possible but
// may not always include both the arrival and departure times. For example,
// when the user arrives at a location, the system may send an event with only
// an arrival time. When the user departs a location, the event can contain
// both the arrival time (if your app was monitoring visits prior to the
// user’s arrival) and the departure time.
//
// # Getting the location
//
//   - [CLVisit.Coordinate]: The geographical coordinate information.
//   - [CLVisit.HorizontalAccuracy]: The horizontal accuracy (in meters) of the specified coordinate.
//
// # Getting the visit duration
//
//   - [CLVisit.ArrivalDate]: The approximate time at which the user arrived at the specified location.
//   - [CLVisit.DepartureDate]: The approximate time at which the user left the specified location.
//
// # Initializers
//
//   - [CLVisit.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit
type CLVisit struct {
	objectivec.Object
}

// CLVisitFromID constructs a [CLVisit] from an objc.ID.
//
// Information about the user’s location during a specific period of time.
func CLVisitFromID(id objc.ID) CLVisit {
	return CLVisit{objectivec.Object{ID: id}}
}

// NOTE: CLVisit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLVisit] class.
//
// # Getting the location
//
//   - [ICLVisit.Coordinate]: The geographical coordinate information.
//   - [ICLVisit.HorizontalAccuracy]: The horizontal accuracy (in meters) of the specified coordinate.
//
// # Getting the visit duration
//
//   - [ICLVisit.ArrivalDate]: The approximate time at which the user arrived at the specified location.
//   - [ICLVisit.DepartureDate]: The approximate time at which the user left the specified location.
//
// # Initializers
//
//   - [ICLVisit.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit
type ICLVisit interface {
	objectivec.IObject

	// Topic: Getting the location

	// The geographical coordinate information.
	Coordinate() CLLocationCoordinate2D
	// The horizontal accuracy (in meters) of the specified coordinate.
	HorizontalAccuracy() CLLocationAccuracy

	// Topic: Getting the visit duration

	// The approximate time at which the user arrived at the specified location.
	ArrivalDate() foundation.NSDate
	// The approximate time at which the user left the specified location.
	DepartureDate() foundation.NSDate

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLVisit

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v CLVisit) Init() CLVisit {
	rv := objc.Send[CLVisit](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v CLVisit) Autorelease() CLVisit {
	rv := objc.Send[CLVisit](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLVisit creates a new CLVisit instance.
func NewCLVisit() CLVisit {
	class := getCLVisitClass()
	rv := objc.Send[CLVisit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/init(coder:)
func NewVisitWithCoder(coder foundation.INSCoder) CLVisit {
	instance := getCLVisitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLVisitFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/init(coder:)
func (v CLVisit) InitWithCoder(coder foundation.INSCoder) CLVisit {
	rv := objc.Send[CLVisit](v.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (v CLVisit) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The geographical coordinate information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/coordinate
func (v CLVisit) Coordinate() CLLocationCoordinate2D {
	rv := objc.Send[CLLocationCoordinate2D](v.ID, objc.Sel("coordinate"))
	return CLLocationCoordinate2D(rv)
}

// The horizontal accuracy (in meters) of the specified coordinate.
//
// # Discussion
//
// The latitude and longitude specified by the [CLVisit.Coordinate] property
// identify the center of the circle, and this value indicates the radius of
// that circle.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/horizontalAccuracy
func (v CLVisit) HorizontalAccuracy() CLLocationAccuracy {
	rv := objc.Send[CLLocationAccuracy](v.ID, objc.Sel("horizontalAccuracy"))
	return CLLocationAccuracy(rv)
}

// The approximate time at which the user arrived at the specified location.
//
// # Discussion
//
// When the visit object does not include arrival information, this property
// is set to the date returned by the [distantPast] method of [NSDate].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/arrivalDate
//
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [distantPast]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
func (v CLVisit) ArrivalDate() foundation.NSDate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("arrivalDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The approximate time at which the user left the specified location.
//
// # Discussion
//
// When the visit object does not include departure information, this property
// is set to the date returned by the [distantFuture] method of [NSDate].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLVisit/departureDate
//
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [distantFuture]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
func (v CLVisit) DepartureDate() foundation.NSDate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("departureDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
