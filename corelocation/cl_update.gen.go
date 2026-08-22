// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLUpdate] class.
var (
	_CLUpdateClass     CLUpdateClass
	_CLUpdateClassOnce sync.Once
)

func getCLUpdateClass() CLUpdateClass {
	_CLUpdateClassOnce.Do(func() {
		_CLUpdateClass = CLUpdateClass{class: objc.GetClass("CLUpdate")}
	})
	return _CLUpdateClass
}

// GetCLUpdateClass returns the class object for CLUpdate.
func GetCLUpdateClass() CLUpdateClass {
	return getCLUpdateClass()
}

type CLUpdateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLUpdateClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLUpdateClass) Alloc() CLUpdate {
	rv := objc.Send[CLUpdate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a location update.
//
// # Update properties
//
//   - [CLUpdate.IsStationary]: A Boolean value that indicates whether the device is stationary.
//   - [CLUpdate.Location]: A person’s location, if available.
//
// # Instance Properties
//
//   - [CLUpdate.AccuracyLimited]
//   - [CLUpdate.AuthorizationDenied]
//   - [CLUpdate.AuthorizationDeniedGlobally]
//   - [CLUpdate.AuthorizationRequestInProgress]
//   - [CLUpdate.AuthorizationRestricted]
//   - [CLUpdate.InsufficientlyInUse]
//   - [CLUpdate.LocationUnavailable]
//   - [CLUpdate.ServiceSessionRequired]
//   - [CLUpdate.Stationary]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate
type CLUpdate struct {
	objectivec.Object
}

// CLUpdateFromID constructs a [CLUpdate] from an objc.ID.
//
// An object that represents a location update.
func CLUpdateFromID(id objc.ID) CLUpdate {
	return CLUpdate{objectivec.Object{ID: id}}
}

// Ensure CLUpdate implements ICLUpdate.
var _ ICLUpdate = CLUpdate{}

// An interface definition for the [CLUpdate] class.
//
// # Update properties
//
//   - [ICLUpdate.IsStationary]: A Boolean value that indicates whether the device is stationary.
//   - [ICLUpdate.Location]: A person’s location, if available.
//
// # Instance Properties
//
//   - [ICLUpdate.AccuracyLimited]
//   - [ICLUpdate.AuthorizationDenied]
//   - [ICLUpdate.AuthorizationDeniedGlobally]
//   - [ICLUpdate.AuthorizationRequestInProgress]
//   - [ICLUpdate.AuthorizationRestricted]
//   - [ICLUpdate.InsufficientlyInUse]
//   - [ICLUpdate.LocationUnavailable]
//   - [ICLUpdate.ServiceSessionRequired]
//   - [ICLUpdate.Stationary]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate
type ICLUpdate interface {
	objectivec.IObject

	// Topic: Update properties

	// A Boolean value that indicates whether the device is stationary.
	IsStationary() bool
	// A person’s location, if available.
	Location() ICLLocation

	// Topic: Instance Properties

	AccuracyLimited() bool
	AuthorizationDenied() bool
	AuthorizationDeniedGlobally() bool
	AuthorizationRequestInProgress() bool
	AuthorizationRestricted() bool
	InsufficientlyInUse() bool
	LocationUnavailable() bool
	ServiceSessionRequired() bool
	Stationary() bool
}

// Init initializes the instance.
func (u CLUpdate) Init() CLUpdate {
	rv := objc.Send[CLUpdate](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u CLUpdate) Autorelease() CLUpdate {
	rv := objc.Send[CLUpdate](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLUpdate creates a new CLUpdate instance.
func NewCLUpdate() CLUpdate {
	class := getCLUpdateClass()
	rv := objc.Send[CLUpdate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the device is stationary.
//
// # Discussion
//
// Updates may stop flowing temporarily for several reasons including if the
// app is no longer authorized to receive location updates or if its location
// becomes unknown. If Core Location stops delivering updates because the
// device is stationary, then it sets `isStationary` to true; otherwise,
// it’s false.
//
// If `isStationary` is true, then the framework can suspend updates until the
// person starts moving, or their location becomes unknown.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/isStationary
func (u CLUpdate) IsStationary() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isStationary"))
	return rv
}

// A person’s location, if available.
//
// # Discussion
//
// If the location isn’t available, the value is `nil`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/location
func (u CLUpdate) Location() ICLLocation {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("location"))
	return CLLocationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/accuracyLimited
func (u CLUpdate) AccuracyLimited() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("accuracyLimited"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationDenied
func (u CLUpdate) AuthorizationDenied() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("authorizationDenied"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationDeniedGlobally
func (u CLUpdate) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationRequestInProgress
func (u CLUpdate) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/authorizationRestricted
func (u CLUpdate) AuthorizationRestricted() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("authorizationRestricted"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/insufficientlyInUse
func (u CLUpdate) InsufficientlyInUse() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("insufficientlyInUse"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/locationUnavailable
func (u CLUpdate) LocationUnavailable() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("locationUnavailable"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/serviceSessionRequired
func (u CLUpdate) ServiceSessionRequired() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("serviceSessionRequired"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLUpdate/stationary
func (u CLUpdate) Stationary() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("stationary"))
	return rv
}
