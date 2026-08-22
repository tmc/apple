// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CLCircularGeographicCondition] class.
var (
	_CLCircularGeographicConditionClass     CLCircularGeographicConditionClass
	_CLCircularGeographicConditionClassOnce sync.Once
)

func getCLCircularGeographicConditionClass() CLCircularGeographicConditionClass {
	_CLCircularGeographicConditionClassOnce.Do(func() {
		_CLCircularGeographicConditionClass = CLCircularGeographicConditionClass{class: objc.GetClass("CLCircularGeographicCondition")}
	})
	return _CLCircularGeographicConditionClass
}

// GetCLCircularGeographicConditionClass returns the class object for CLCircularGeographicCondition.
func GetCLCircularGeographicConditionClass() CLCircularGeographicConditionClass {
	return getCLCircularGeographicConditionClass()
}

type CLCircularGeographicConditionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLCircularGeographicConditionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLCircularGeographicConditionClass) Alloc() CLCircularGeographicCondition {
	rv := objc.Send[CLCircularGeographicCondition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A circular geographic condition that a center point and radius define.
//
// # Overview
//
// Use [CLCircularGeographicCondition] to monitor events that occur in a
// circular geographic condition that you describe.
//
// # Creating a circular geographic condition
//
//   - [CLCircularGeographicCondition.InitWithCenterRadius]: Creates a new circular geographic condition with the center point and radius you provide.
//
// # Instance properties
//
//   - [CLCircularGeographicCondition.Center]: The center of the circular geographic condition.
//   - [CLCircularGeographicCondition.Radius]: The radius of the circular geographic condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition
type CLCircularGeographicCondition struct {
	CLCondition
}

// CLCircularGeographicConditionFromID constructs a [CLCircularGeographicCondition] from an objc.ID.
//
// A circular geographic condition that a center point and radius define.
func CLCircularGeographicConditionFromID(id objc.ID) CLCircularGeographicCondition {
	return CLCircularGeographicCondition{CLCondition: CLConditionFromID(id)}
}

// NOTE: CLCircularGeographicCondition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLCircularGeographicCondition] class.
//
// # Creating a circular geographic condition
//
//   - [ICLCircularGeographicCondition.InitWithCenterRadius]: Creates a new circular geographic condition with the center point and radius you provide.
//
// # Instance properties
//
//   - [ICLCircularGeographicCondition.Center]: The center of the circular geographic condition.
//   - [ICLCircularGeographicCondition.Radius]: The radius of the circular geographic condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition
type ICLCircularGeographicCondition interface {
	ICLCondition

	// Topic: Creating a circular geographic condition

	// Creates a new circular geographic condition with the center point and radius you provide.
	InitWithCenterRadius(center CLLocationCoordinate2D, radius CLLocationDistance) CLCircularGeographicCondition

	// Topic: Instance properties

	// The center of the circular geographic condition.
	Center() CLLocationCoordinate2D
	// The radius of the circular geographic condition.
	Radius() CLLocationDistance
}

// Init initializes the instance.
func (c CLCircularGeographicCondition) Init() CLCircularGeographicCondition {
	rv := objc.Send[CLCircularGeographicCondition](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CLCircularGeographicCondition) Autorelease() CLCircularGeographicCondition {
	rv := objc.Send[CLCircularGeographicCondition](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLCircularGeographicCondition creates a new CLCircularGeographicCondition instance.
func NewCLCircularGeographicCondition() CLCircularGeographicCondition {
	class := getCLCircularGeographicConditionClass()
	rv := objc.Send[CLCircularGeographicCondition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new circular geographic condition with the center point and
// radius you provide.
//
// center: The center of the circular geographic condition.
//
// radius: The radius of the circular geographic condition.
//
// # Return Value
//
// Returns an instance of [CLCircularGeographicCondition] with the specified
// center coordinate and radius.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/initWithCenter:radius:
func NewCircularGeographicConditionWithCenterRadius(center CLLocationCoordinate2D, radius CLLocationDistance) CLCircularGeographicCondition {
	instance := getCLCircularGeographicConditionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	return CLCircularGeographicConditionFromID(rv)
}

// Creates a new circular geographic condition with the center point and
// radius you provide.
//
// center: The center of the circular geographic condition.
//
// radius: The radius of the circular geographic condition.
//
// # Return Value
//
// Returns an instance of [CLCircularGeographicCondition] with the specified
// center coordinate and radius.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/initWithCenter:radius:
func (c CLCircularGeographicCondition) InitWithCenterRadius(center CLLocationCoordinate2D, radius CLLocationDistance) CLCircularGeographicCondition {
	rv := objc.Send[CLCircularGeographicCondition](c.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	return rv
}

// The center of the circular geographic condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/center
func (c CLCircularGeographicCondition) Center() CLLocationCoordinate2D {
	rv := objc.Send[CLLocationCoordinate2D](c.ID, objc.Sel("center"))
	return CLLocationCoordinate2D(rv)
}

// The radius of the circular geographic condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/radius
func (c CLCircularGeographicCondition) Radius() CLLocationDistance {
	rv := objc.Send[CLLocationDistance](c.ID, objc.Sel("radius"))
	return CLLocationDistance(rv)
}
