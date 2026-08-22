// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLRegion] class.
var (
	_CLRegionClass     CLRegionClass
	_CLRegionClassOnce sync.Once
)

func getCLRegionClass() CLRegionClass {
	_CLRegionClassOnce.Do(func() {
		_CLRegionClass = CLRegionClass{class: objc.GetClass("CLRegion")}
	})
	return _CLRegionClass
}

// GetCLRegionClass returns the class object for CLRegion.
func GetCLRegionClass() CLRegionClass {
	return getCLRegionClass()
}

type CLRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLRegionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLRegionClass) Alloc() CLRegion {
	rv := objc.Send[CLRegion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A base class representing an area that can be monitored.
//
// # Overview
//
// This is an abstract base class. Instantiate one of the provided subclasses
// that define specific types of regions. After you create a region, register
// it with a [CLLocationManager] object with the
// [CLLocationManager.StartMonitoringForRegion] method. The location manager
// generates appropriate events whenever the user crosses the boundaries of
// the region.
//
// # Getting the region identifier
//
//   - [CLRegion.Identifier]: The identifier for the region object.
//
// # Specifying the notification conditions
//
//   - [CLRegion.NotifyOnEntry]: A Boolean indicating that notifications are generated upon entry into the region.
//   - [CLRegion.SetNotifyOnEntry]
//   - [CLRegion.NotifyOnExit]: A Boolean indicating that notifications are generated upon exit from the region.
//   - [CLRegion.SetNotifyOnExit]
//
// # Initializers
//
//   - [CLRegion.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLRegion
type CLRegion struct {
	objectivec.Object
}

// CLRegionFromID constructs a [CLRegion] from an objc.ID.
//
// A base class representing an area that can be monitored.
func CLRegionFromID(id objc.ID) CLRegion {
	return CLRegion{objectivec.Object{ID: id}}
}

// NOTE: CLRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLRegion] class.
//
// # Getting the region identifier
//
//   - [ICLRegion.Identifier]: The identifier for the region object.
//
// # Specifying the notification conditions
//
//   - [ICLRegion.NotifyOnEntry]: A Boolean indicating that notifications are generated upon entry into the region.
//   - [ICLRegion.SetNotifyOnEntry]
//   - [ICLRegion.NotifyOnExit]: A Boolean indicating that notifications are generated upon exit from the region.
//   - [ICLRegion.SetNotifyOnExit]
//
// # Initializers
//
//   - [ICLRegion.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLRegion
type ICLRegion interface {
	objectivec.IObject

	// Topic: Getting the region identifier

	// The identifier for the region object.
	Identifier() string

	// Topic: Specifying the notification conditions

	// A Boolean indicating that notifications are generated upon entry into the region.
	NotifyOnEntry() bool
	SetNotifyOnEntry(value bool)
	// A Boolean indicating that notifications are generated upon exit from the region.
	NotifyOnExit() bool
	SetNotifyOnExit(value bool)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLRegion

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r CLRegion) Init() CLRegion {
	rv := objc.Send[CLRegion](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CLRegion) Autorelease() CLRegion {
	rv := objc.Send[CLRegion](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLRegion creates a new CLRegion instance.
func NewCLRegion() CLRegion {
	class := getCLRegionClass()
	rv := objc.Send[CLRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(coder:)
func NewRegionWithCoder(coder foundation.INSCoder) CLRegion {
	instance := getCLRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLRegionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(coder:)
func (r CLRegion) InitWithCoder(coder foundation.INSCoder) CLRegion {
	rv := objc.Send[CLRegion](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (r CLRegion) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The identifier for the region object.
//
// # Discussion
//
// Use this value to identify this region inside your application.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/identifier
func (r CLRegion) Identifier() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean indicating that notifications are generated upon entry into the
// region.
//
// # Discussion
//
// When this property is true, a device crossing from outside the region to
// inside the region triggers the delivery of a notification. If the property
// is false, a notification is not generated. The default value of this
// property is true.
//
// If the app is not running when a boundary crossing occurs, the system
// launches the app into the background to handle it. Upon launch, your app
// must configure new location manager and delegate objects to receive the
// notification. The notification is sent to your delegate’s
// [LocationManagerDidEnterRegion] method.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnEntry
func (r CLRegion) NotifyOnEntry() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("notifyOnEntry"))
	return rv
}
func (r CLRegion) SetNotifyOnEntry(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setNotifyOnEntry:"), value)
}

// A Boolean indicating that notifications are generated upon exit from the
// region.
//
// # Discussion
//
// When this property is true, a device crossing from inside the region to
// outside the region triggers the delivery of a notification. If the property
// is false, a notification is not generated. The default value of this
// property is true.
//
// If the app is not running when a boundary crossing occurs, the system
// launches the app into the background to handle it. Upon launch, your app
// must configure new location manager and delegate objects to receive the
// notification. The notification is sent to your delegate’s
// [LocationManagerDidExitRegion] method.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnExit
func (r CLRegion) NotifyOnExit() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("notifyOnExit"))
	return rv
}
func (r CLRegion) SetNotifyOnExit(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setNotifyOnExit:"), value)
}
