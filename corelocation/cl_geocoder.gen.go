// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLGeocoder] class.
var (
	_CLGeocoderClass     CLGeocoderClass
	_CLGeocoderClassOnce sync.Once
)

func getCLGeocoderClass() CLGeocoderClass {
	_CLGeocoderClassOnce.Do(func() {
		_CLGeocoderClass = CLGeocoderClass{class: objc.GetClass("CLGeocoder")}
	})
	return _CLGeocoderClass
}

// GetCLGeocoderClass returns the class object for CLGeocoder.
func GetCLGeocoderClass() CLGeocoderClass {
	return getCLGeocoderClass()
}

type CLGeocoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLGeocoderClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLGeocoderClass) Alloc() CLGeocoder {
	rv := objc.Send[CLGeocoder](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An interface for converting between geographic coordinates and place names.
//
// # Overview
//
// The [CLGeocoder] class provides services for converting between a
// coordinate (specified as a latitude and longitude) and the user-friendly
// representation of that coordinate. A user-friendly representation of the
// coordinate typically consists of the street, city, state, and country or
// region information corresponding to the given location, but it may also
// contain a relevant point of interest, landmarks, or other identifying
// information. A geocoder object is a single-shot object that works with a
// network-based service to look up placemark information for its specified
// coordinate value.
//
// To use a geocoder object, you create it and call one of its forward- or
// reverse-geocoding methods to begin the request. Reverse-geocoding requests
// take a latitude and longitude value and find a user-readable address.
// Forward-geocoding requests take a user-readable address and find the
// corresponding latitude and longitude value. Forward-geocoding requests may
// also return additional information about the specified location, such as a
// point of interest or building at that location. For both types of request,
// the results are returned using a [CLPlacemark] object. In the case of
// forward-geocoding requests, multiple placemark objects may be returned if
// the provided information yielded multiple possible locations.
//
// To make smart decisions about what types of information to return, the
// geocoder server uses all the information provided to it when processing the
// request. For example, if the user is moving quickly along a highway, it
// might return the name of the overall region, and not the name of a small
// park that the user is passing through.
//
// # Tips for Using a Geocoder Object
//
// Apps must be conscious of how they use geocoding. Geocoding requests are
// rate-limited for each app, so making too many requests in a short period of
// time may cause some of the requests to fail. (When the maximum rate is
// exceeded, the geocoder returns an error object with the [KCLErrorNetwork]
// error to the associated completion handler.) Here are some rules of thumb
// for using this class effectively:
//
// - Send at most one geocoding request for any one user action. - If the user
// performs multiple actions that involve geocoding the same location, reuse
// the results from the initial geocoding request instead of starting
// individual requests for each action. - When you want to update the user’s
// current location automatically (such as when the user is moving), issue new
// geocoding requests only when the user has moved a significant distance and
// after a reasonable amount of time has passed. For example, in a typical
// situation, you should not send more than one geocoding request per minute.
// - Do not start a geocoding request at a time when the user will not see the
// results immediately. For example, do not start a request if your
// application is inactive or in the background.
//
// The computer or device must have access to the network in order for the
// geocoder object to return detailed placemark information. Although, the
// geocoder stores enough information locally to report the localized country
// or region name and ISO country code for many locations. If this information
// isn’t available for a specific location, the geocoder may still report an
// error to your completion block.
//
// You can use geocoder objects either in conjunction with, or independent of,
// the classes of the [MapKit] framework.
//
// # Managing geocoding requests
//
//   - [CLGeocoder.IsGeocoding]: A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLGeocoder
//
// [MapKit]: https://developer.apple.com/documentation/MapKit
type CLGeocoder struct {
	objectivec.Object
}

// CLGeocoderFromID constructs a [CLGeocoder] from an objc.ID.
//
// An interface for converting between geographic coordinates and place names.
func CLGeocoderFromID(id objc.ID) CLGeocoder {
	return CLGeocoder{objectivec.Object{ID: id}}
}

// NOTE: CLGeocoder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLGeocoder] class.
//
// # Managing geocoding requests
//
//   - [ICLGeocoder.IsGeocoding]: A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLGeocoder
type ICLGeocoder interface {
	objectivec.IObject

	// Topic: Managing geocoding requests

	// A Boolean value indicating whether the receiver is in the middle of geocoding its value.
	IsGeocoding() bool
}

// Init initializes the instance.
func (g CLGeocoder) Init() CLGeocoder {
	rv := objc.Send[CLGeocoder](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g CLGeocoder) Autorelease() CLGeocoder {
	rv := objc.Send[CLGeocoder](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLGeocoder creates a new CLGeocoder instance.
func NewCLGeocoder() CLGeocoder {
	class := getCLGeocoderClass()
	rv := objc.Send[CLGeocoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value indicating whether the receiver is in the middle of
// geocoding its value.
//
// # Discussion
//
// This property contains the value true if the process is ongoing or false if
// the process is done or has not yet been initiated.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/isGeocoding
func (g CLGeocoder) IsGeocoding() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isGeocoding"))
	return rv
}
