// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSSystemExtensionManager] class.
var (
	_OSSystemExtensionManagerClass     OSSystemExtensionManagerClass
	_OSSystemExtensionManagerClassOnce sync.Once
)

func getOSSystemExtensionManagerClass() OSSystemExtensionManagerClass {
	_OSSystemExtensionManagerClassOnce.Do(func() {
		_OSSystemExtensionManagerClass = OSSystemExtensionManagerClass{class: objc.GetClass("OSSystemExtensionManager")}
	})
	return _OSSystemExtensionManagerClass
}

// GetOSSystemExtensionManagerClass returns the class object for OSSystemExtensionManager.
func GetOSSystemExtensionManagerClass() OSSystemExtensionManagerClass {
	return getOSSystemExtensionManagerClass()
}

type OSSystemExtensionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSSystemExtensionManagerClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSSystemExtensionManagerClass) Alloc() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A type that facilitates activation and deactivation of system extensions.
//
// # Overview
//
// Create an instance of [OSSystemExtensionRequest] with the class methods on
// that type, and submit it to the shared instance of the extension manager
// with [OSSystemExtensionManager.SubmitRequest]. Set the [OSSystemExtensionManager.Delegate] on the request to receive the
// result of the activation or deactivation. The delegate also receives
// notifications if the user needs to authorize the extension or if a version
// conflict occurs.
//
// # Submitting Requests
//
//   - [OSSystemExtensionManager.SubmitRequest]: Submits a system extension request to the manager.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager
type OSSystemExtensionManager struct {
	objectivec.Object
}

// OSSystemExtensionManagerFromID constructs a [OSSystemExtensionManager] from an objc.ID.
//
// A type that facilitates activation and deactivation of system extensions.
func OSSystemExtensionManagerFromID(id objc.ID) OSSystemExtensionManager {
	return OSSystemExtensionManager{objectivec.Object{ID: id}}
}

// NOTE: OSSystemExtensionManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSSystemExtensionManager] class.
//
// # Submitting Requests
//
//   - [IOSSystemExtensionManager.SubmitRequest]: Submits a system extension request to the manager.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager
type IOSSystemExtensionManager interface {
	objectivec.IObject

	// Topic: Submitting Requests

	// Submits a system extension request to the manager.
	SubmitRequest(request IOSSystemExtensionRequest)

	// A delegate to receive updates about the progress of a request.
	Delegate() OSSystemExtensionRequestDelegate
	SetDelegate(value OSSystemExtensionRequestDelegate)
}

// Init initializes the instance.
func (o OSSystemExtensionManager) Init() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSSystemExtensionManager) Autorelease() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionManager creates a new OSSystemExtensionManager instance.
func NewOSSystemExtensionManager() OSSystemExtensionManager {
	class := getOSSystemExtensionManagerClass()
	rv := objc.Send[OSSystemExtensionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Submits a system extension request to the manager.
//
// request: The request to process.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/submitRequest(_:)
func (o OSSystemExtensionManager) SubmitRequest(request IOSSystemExtensionRequest) {
	objc.Send[objc.ID](o.ID, objc.Sel("submitRequest:"), request)
}

// A delegate to receive updates about the progress of a request.
//
// See: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequest/delegate
func (o OSSystemExtensionManager) Delegate() OSSystemExtensionRequestDelegate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("delegate"))
	return OSSystemExtensionRequestDelegateObjectFromID(rv)
}
func (o OSSystemExtensionManager) SetDelegate(value OSSystemExtensionRequestDelegate) {
	objc.Send[struct{}](o.ID, objc.Sel("setDelegate:"), value)
}

// The shared instance of the extension manager.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/shared
func (_OSSystemExtensionManagerClass OSSystemExtensionManagerClass) SharedManager() OSSystemExtensionManager {
	rv := objc.Send[objc.ID](objc.ID(_OSSystemExtensionManagerClass.class), objc.Sel("sharedManager"))
	return OSSystemExtensionManagerFromID(objc.ID(rv))
}
