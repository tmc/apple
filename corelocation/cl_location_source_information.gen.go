// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLLocationSourceInformation] class.
var (
	_CLLocationSourceInformationClass     CLLocationSourceInformationClass
	_CLLocationSourceInformationClassOnce sync.Once
)

func getCLLocationSourceInformationClass() CLLocationSourceInformationClass {
	_CLLocationSourceInformationClassOnce.Do(func() {
		_CLLocationSourceInformationClass = CLLocationSourceInformationClass{class: objc.GetClass("CLLocationSourceInformation")}
	})
	return _CLLocationSourceInformationClass
}

// GetCLLocationSourceInformationClass returns the class object for CLLocationSourceInformation.
func GetCLLocationSourceInformationClass() CLLocationSourceInformationClass {
	return getCLLocationSourceInformationClass()
}

type CLLocationSourceInformationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLLocationSourceInformationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLLocationSourceInformationClass) Alloc() CLLocationSourceInformation {
	rv := objc.Send[CLLocationSourceInformation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about the source that provides a location.
//
// # Overview
//
// [CLLocationSourceInformation] contains information about the source that
// provides a [CLLocation] instance, such as instances that
// [LocationManagerDidUpdateLocations] delivers. For example, an app may
// choose to check the source information and reject locations if the
// [CLLocationSourceInformation.IsSimulatedBySoftware] property is `true` when
// the developer isn’t debugging or testing the app.
//
// # Creating a location source information object
//
//   - [CLLocationSourceInformation.InitWithSoftwareSimulationStateAndExternalAccessoryState]: Creates an instance of location source information.
//
// # Identifying the source of location data
//
//   - [CLLocationSourceInformation.IsProducedByAccessory]: A Boolean value that indicates whether the system receives the location from an external accessory.
//   - [CLLocationSourceInformation.IsSimulatedBySoftware]: A Boolean value that indicates whether the system generates the location using on-device software simulation.
//
// # Initializers
//
//   - [CLLocationSourceInformation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation
type CLLocationSourceInformation struct {
	objectivec.Object
}

// CLLocationSourceInformationFromID constructs a [CLLocationSourceInformation] from an objc.ID.
//
// Information about the source that provides a location.
func CLLocationSourceInformationFromID(id objc.ID) CLLocationSourceInformation {
	return CLLocationSourceInformation{objectivec.Object{ID: id}}
}

// NOTE: CLLocationSourceInformation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLLocationSourceInformation] class.
//
// # Creating a location source information object
//
//   - [ICLLocationSourceInformation.InitWithSoftwareSimulationStateAndExternalAccessoryState]: Creates an instance of location source information.
//
// # Identifying the source of location data
//
//   - [ICLLocationSourceInformation.IsProducedByAccessory]: A Boolean value that indicates whether the system receives the location from an external accessory.
//   - [ICLLocationSourceInformation.IsSimulatedBySoftware]: A Boolean value that indicates whether the system generates the location using on-device software simulation.
//
// # Initializers
//
//   - [ICLLocationSourceInformation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation
type ICLLocationSourceInformation interface {
	objectivec.IObject

	// Topic: Creating a location source information object

	// Creates an instance of location source information.
	InitWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware bool, isAccessory bool) CLLocationSourceInformation

	// Topic: Identifying the source of location data

	// A Boolean value that indicates whether the system receives the location from an external accessory.
	IsProducedByAccessory() bool
	// A Boolean value that indicates whether the system generates the location using on-device software simulation.
	IsSimulatedBySoftware() bool

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLLocationSourceInformation

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (l CLLocationSourceInformation) Init() CLLocationSourceInformation {
	rv := objc.Send[CLLocationSourceInformation](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l CLLocationSourceInformation) Autorelease() CLLocationSourceInformation {
	rv := objc.Send[CLLocationSourceInformation](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLLocationSourceInformation creates a new CLLocationSourceInformation instance.
func NewCLLocationSourceInformation() CLLocationSourceInformation {
	class := getCLLocationSourceInformationClass()
	rv := objc.Send[CLLocationSourceInformation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/init(coder:)
func NewLocationSourceInformationWithCoder(coder foundation.INSCoder) CLLocationSourceInformation {
	instance := getCLLocationSourceInformationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLLocationSourceInformationFromID(rv)
}

// Creates an instance of location source information.
//
// isSoftware: A Boolean value that indicates software is generating or simulating the
// location information.
//
// isAccessory: A Boolean value that indicates an external device is providing the location
// information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/init(softwareSimulationState:andExternalAccessoryState:)
func NewLocationSourceInformationWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware bool, isAccessory bool) CLLocationSourceInformation {
	instance := getCLLocationSourceInformationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSoftwareSimulationState:andExternalAccessoryState:"), isSoftware, isAccessory)
	return CLLocationSourceInformationFromID(rv)
}

// Creates an instance of location source information.
//
// isSoftware: A Boolean value that indicates software is generating or simulating the
// location information.
//
// isAccessory: A Boolean value that indicates an external device is providing the location
// information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/init(softwareSimulationState:andExternalAccessoryState:)
func (l CLLocationSourceInformation) InitWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware bool, isAccessory bool) CLLocationSourceInformation {
	rv := objc.Send[CLLocationSourceInformation](l.ID, objc.Sel("initWithSoftwareSimulationState:andExternalAccessoryState:"), isSoftware, isAccessory)
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/init(coder:)
func (l CLLocationSourceInformation) InitWithCoder(coder foundation.INSCoder) CLLocationSourceInformation {
	rv := objc.Send[CLLocationSourceInformation](l.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (l CLLocationSourceInformation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether the system receives the location
// from an external accessory.
//
// # Discussion
//
// Core Location sets [CLLocationSourceInformation.IsProducedByAccessory] to
// `true` if the system retrieved the location from an external accessory
// attached to the device, such as a Made for iPhone GPS dongle or CarPlay.
// Otherwise, the default value is `false`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isProducedByAccessory
func (l CLLocationSourceInformation) IsProducedByAccessory() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isProducedByAccessory"))
	return rv
}

// A Boolean value that indicates whether the system generates the location
// using on-device software simulation.
//
// # Discussion
//
// Core Location sets [CLLocationSourceInformation.IsSimulatedBySoftware] to
// `true` if the system generated the location using on-device software
// simulation. You can simulate locations by loading GPX files using the Xcode
// debugger. The default value is `false`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isSimulatedBySoftware
func (l CLLocationSourceInformation) IsSimulatedBySoftware() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isSimulatedBySoftware"))
	return rv
}
