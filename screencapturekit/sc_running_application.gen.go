// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCRunningApplication] class.
var (
	_SCRunningApplicationClass     SCRunningApplicationClass
	_SCRunningApplicationClassOnce sync.Once
)

func getSCRunningApplicationClass() SCRunningApplicationClass {
	_SCRunningApplicationClassOnce.Do(func() {
		_SCRunningApplicationClass = SCRunningApplicationClass{class: objc.GetClass("SCRunningApplication")}
	})
	return _SCRunningApplicationClass
}

// GetSCRunningApplicationClass returns the class object for SCRunningApplication.
func GetSCRunningApplicationClass() SCRunningApplicationClass {
	return getSCRunningApplicationClass()
}

type SCRunningApplicationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCRunningApplicationClass) Alloc() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An instance that represents an app running on a device.
//
// # Overview
// 
// Retrieve the available apps from an instance of [SCShareableContent].
// Select one or more apps to capture and use them to create an instance of
// [SCContentFilter]. Apply the filter to an instance of [SCStream] to limit
// its output to content matching your criteria.
//
// # Inspecting an app
//
//   - [SCRunningApplication.ProcessID]: The system process identifier of the app.
//   - [SCRunningApplication.BundleIdentifier]: The unique bundle identifier of the app.
//   - [SCRunningApplication.ApplicationName]: The display name of the app.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication
type SCRunningApplication struct {
	objectivec.Object
}

// SCRunningApplicationFromID constructs a [SCRunningApplication] from an objc.ID.
//
// An instance that represents an app running on a device.
func SCRunningApplicationFromID(id objc.ID) SCRunningApplication {
	return SCRunningApplication{objectivec.Object{id}}
}
// NOTE: SCRunningApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCRunningApplication] class.
//
// # Inspecting an app
//
//   - [ISCRunningApplication.ProcessID]: The system process identifier of the app.
//   - [ISCRunningApplication.BundleIdentifier]: The unique bundle identifier of the app.
//   - [ISCRunningApplication.ApplicationName]: The display name of the app.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication
type ISCRunningApplication interface {
	objectivec.IObject

	// Topic: Inspecting an app

	// The system process identifier of the app.
	ProcessID() int32
	// The unique bundle identifier of the app.
	BundleIdentifier() string
	// The display name of the app.
	ApplicationName() string
}





// Init initializes the instance.
func (r SCRunningApplication) Init() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SCRunningApplication) Autorelease() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCRunningApplication creates a new SCRunningApplication instance.
func NewSCRunningApplication() SCRunningApplication {
	class := getSCRunningApplicationClass()
	rv := objc.Send[SCRunningApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The system process identifier of the app.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/processID
func (r SCRunningApplication) ProcessID() int32 {
	rv := objc.Send[int32](r.ID, objc.Sel("processID"))
	return rv
}



// The unique bundle identifier of the app.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/bundleIdentifier
func (r SCRunningApplication) BundleIdentifier() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}



// The display name of the app.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/applicationName
func (r SCRunningApplication) ApplicationName() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("applicationName"))
	return foundation.NSStringFromID(rv).String()
}























