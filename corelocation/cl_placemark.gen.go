// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLPlacemark] class.
var (
	_CLPlacemarkClass     CLPlacemarkClass
	_CLPlacemarkClassOnce sync.Once
)

func getCLPlacemarkClass() CLPlacemarkClass {
	_CLPlacemarkClassOnce.Do(func() {
		_CLPlacemarkClass = CLPlacemarkClass{class: objc.GetClass("CLPlacemark")}
	})
	return _CLPlacemarkClass
}

// GetCLPlacemarkClass returns the class object for CLPlacemark.
func GetCLPlacemarkClass() CLPlacemarkClass {
	return getCLPlacemarkClass()
}

type CLPlacemarkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLPlacemarkClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLPlacemarkClass) Alloc() CLPlacemark {
	rv := objc.Send[CLPlacemark](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A user-friendly description of a geographic coordinate, often containing
// the name of the place, its address, and other relevant information.
//
// # Overview
//
// A [CLPlacemark] object stores placemark data for a given latitude and
// longitude. Placemark data includes information such as the country or
// region, state, city, and street address associated with the specified
// coordinate. It can also include points of interest and geographically
// related data.
//
// When you reverse geocode a geographic coordinate using a [CLGeocoder]
// object, you receive a [CLPlacemark] object containing the descriptive
// information for that location. You can also create [CLPlacemark] object and
// fill it with address information yourself, which you might do when you want
// to determine the geographic coordinate associated with the location.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLPlacemark
type CLPlacemark struct {
	objectivec.Object
}

// CLPlacemarkFromID constructs a [CLPlacemark] from an objc.ID.
//
// A user-friendly description of a geographic coordinate, often containing
// the name of the place, its address, and other relevant information.
func CLPlacemarkFromID(id objc.ID) CLPlacemark {
	return CLPlacemark{objectivec.Object{ID: id}}
}

// NOTE: CLPlacemark adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLPlacemark] class.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLPlacemark
type ICLPlacemark interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p CLPlacemark) Init() CLPlacemark {
	rv := objc.Send[CLPlacemark](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p CLPlacemark) Autorelease() CLPlacemark {
	rv := objc.Send[CLPlacemark](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLPlacemark creates a new CLPlacemark instance.
func NewCLPlacemark() CLPlacemark {
	class := getCLPlacemarkClass()
	rv := objc.Send[CLPlacemark](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(coder:)
func NewPlacemarkWithCoder(coder foundation.INSCoder) CLPlacemark {
	instance := getCLPlacemarkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLPlacemarkFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(location:name:postalAddress:)
func NewPlacemarkWithLocationNamePostalAddress(location ICLLocation, name string, postalAddress unsafe.Pointer) CLPlacemark {
	rv := objc.Send[objc.ID](objc.ID(getCLPlacemarkClass().class), objc.Sel("placemarkWithLocation:name:postalAddress:"), location, objc.String(name), postalAddress)
	return CLPlacemarkFromID(rv)
}

func (p CLPlacemark) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}
