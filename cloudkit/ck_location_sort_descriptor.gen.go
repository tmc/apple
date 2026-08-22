// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/corelocation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKLocationSortDescriptor] class.
var (
	_CKLocationSortDescriptorClass     CKLocationSortDescriptorClass
	_CKLocationSortDescriptorClassOnce sync.Once
)

func getCKLocationSortDescriptorClass() CKLocationSortDescriptorClass {
	_CKLocationSortDescriptorClassOnce.Do(func() {
		_CKLocationSortDescriptorClass = CKLocationSortDescriptorClass{class: objc.GetClass("CKLocationSortDescriptor")}
	})
	return _CKLocationSortDescriptorClass
}

// GetCKLocationSortDescriptorClass returns the class object for CKLocationSortDescriptor.
func GetCKLocationSortDescriptorClass() CKLocationSortDescriptorClass {
	return getCKLocationSortDescriptorClass()
}

type CKLocationSortDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKLocationSortDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKLocationSortDescriptorClass) Alloc() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object for sorting records that contain location data.
//
// # Overview
//
// You can add a location sort descriptor to your queries when searching for
// records. At creation time, you must provide the sort descriptor with a key
// that has a [CLLocation] object as its value. The sort descriptor uses the
// value of that key to perform the sort.
//
// CloudKit computes distance by drawing a direct line between the two
// locations that follows the curvature of the Earth. Distances don’t
// account for altitude changes between the two locations.
//
// # Creating a Location Sort Descriptor
//
//   - [CKLocationSortDescriptor.InitWithKeyRelativeLocation]: Creates a location sort descriptor using the specified key and relative location.
//
// # Accessing the Location Value
//
//   - [CKLocationSortDescriptor.RelativeLocation]: The reference location for sorting records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor
//
// [CLLocation]: https://developer.apple.com/documentation/CoreLocation/CLLocation
type CKLocationSortDescriptor struct {
	foundation.NSSortDescriptor
}

// CKLocationSortDescriptorFromID constructs a [CKLocationSortDescriptor] from an objc.ID.
//
// An object for sorting records that contain location data.
func CKLocationSortDescriptorFromID(id objc.ID) CKLocationSortDescriptor {
	return CKLocationSortDescriptor{NSSortDescriptor: foundation.NSSortDescriptorFromID(id)}
}

// NOTE: CKLocationSortDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKLocationSortDescriptor] class.
//
// # Creating a Location Sort Descriptor
//
//   - [ICKLocationSortDescriptor.InitWithKeyRelativeLocation]: Creates a location sort descriptor using the specified key and relative location.
//
// # Accessing the Location Value
//
//   - [ICKLocationSortDescriptor.RelativeLocation]: The reference location for sorting records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor
type ICKLocationSortDescriptor interface {
	foundation.INSSortDescriptor

	// Topic: Creating a Location Sort Descriptor

	// Creates a location sort descriptor using the specified key and relative location.
	InitWithKeyRelativeLocation(key string, relativeLocation corelocation.CLLocation) CKLocationSortDescriptor

	// Topic: Accessing the Location Value

	// The reference location for sorting records.
	RelativeLocation() corelocation.CLLocation
}

// Init initializes the instance.
func (c CKLocationSortDescriptor) Init() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKLocationSortDescriptor) Autorelease() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKLocationSortDescriptor creates a new CKLocationSortDescriptor instance.
func NewCKLocationSortDescriptor() CKLocationSortDescriptor {
	class := getCKLocationSortDescriptorClass()
	rv := objc.Send[CKLocationSortDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a location sort descriptor from a serialized instance.
//
// aDecoder: The coder to use when deserializing the location sort descriptor.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(coder:)
func NewCKLocationSortDescriptorWithCoder(aDecoder foundation.INSCoder) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKLocationSortDescriptorFromID(rv)
}

// Creates a location sort descriptor using the specified key and relative
// location.
//
// key: The name of the key with a [CLLocation] object as its value. The key must
// belong to the records you’re sorting. The sort descriptor uses this key
// to retrieve the corresponding value from the record.
//
// relativeLocation: The reference location when sorting. CloudKit sorts records according to
// their distance from this location.
//
// # Discussion
//
// During sorting, the sort descriptor computes the distance between the value
// in the `relativeLocation` parameter and the location value in the specified
// key of each record. It then sorts the records in ascending order using the
// distance between the two points. You can’t change the sort order.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(key:relativeLocation:)
//
// [CLLocation]: https://developer.apple.com/documentation/CoreLocation/CLLocation
func NewCKLocationSortDescriptorWithKeyRelativeLocation(key string, relativeLocation corelocation.CLLocation) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:relativeLocation:"), objc.String(key), relativeLocation)
	return CKLocationSortDescriptorFromID(rv)
}

// Creates a location sort descriptor using the specified key and relative
// location.
//
// key: The name of the key with a [CLLocation] object as its value. The key must
// belong to the records you’re sorting. The sort descriptor uses this key
// to retrieve the corresponding value from the record.
//
// relativeLocation: The reference location when sorting. CloudKit sorts records according to
// their distance from this location.
//
// # Discussion
//
// During sorting, the sort descriptor computes the distance between the value
// in the `relativeLocation` parameter and the location value in the specified
// key of each record. It then sorts the records in ascending order using the
// distance between the two points. You can’t change the sort order.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(key:relativeLocation:)
//
// [CLLocation]: https://developer.apple.com/documentation/CoreLocation/CLLocation
func (c CKLocationSortDescriptor) InitWithKeyRelativeLocation(key string, relativeLocation corelocation.CLLocation) CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](c.ID, objc.Sel("initWithKey:relativeLocation:"), objc.String(key), relativeLocation)
	return rv
}

// The reference location for sorting records.
//
// See: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/relativeLocation
func (c CKLocationSortDescriptor) RelativeLocation() corelocation.CLLocation {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("relativeLocation"))
	return corelocation.CLLocationFromID(objc.ID(rv))
}
