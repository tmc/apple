// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSSystemExtensionRequest] class.
var (
	_OSSystemExtensionRequestClass     OSSystemExtensionRequestClass
	_OSSystemExtensionRequestClassOnce sync.Once
)

func getOSSystemExtensionRequestClass() OSSystemExtensionRequestClass {
	_OSSystemExtensionRequestClassOnce.Do(func() {
		_OSSystemExtensionRequestClass = OSSystemExtensionRequestClass{class: objc.GetClass("OSSystemExtensionRequest")}
	})
	return _OSSystemExtensionRequestClass
}

// GetOSSystemExtensionRequestClass returns the class object for OSSystemExtensionRequest.
func GetOSSystemExtensionRequestClass() OSSystemExtensionRequestClass {
	return getOSSystemExtensionRequestClass()
}

type OSSystemExtensionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSSystemExtensionRequestClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSSystemExtensionRequestClass) Alloc() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A request to activate or deactivate a system extension.
//
// # Working with a Delegate
//
//   - [OSSystemExtensionRequest.Delegate]: A delegate to receive updates about the progress of a request.
//   - [OSSystemExtensionRequest.SetDelegate]
//
// # Identifying the Target Extension
//
//   - [OSSystemExtensionRequest.Identifier]: The bundle identifier of the target extension.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest
type OSSystemExtensionRequest struct {
	objectivec.Object
}

// OSSystemExtensionRequestFromID constructs a [OSSystemExtensionRequest] from an objc.ID.
//
// A request to activate or deactivate a system extension.
func OSSystemExtensionRequestFromID(id objc.ID) OSSystemExtensionRequest {
	return OSSystemExtensionRequest{objectivec.Object{ID: id}}
}

// NOTE: OSSystemExtensionRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSSystemExtensionRequest] class.
//
// # Working with a Delegate
//
//   - [IOSSystemExtensionRequest.Delegate]: A delegate to receive updates about the progress of a request.
//   - [IOSSystemExtensionRequest.SetDelegate]
//
// # Identifying the Target Extension
//
//   - [IOSSystemExtensionRequest.Identifier]: The bundle identifier of the target extension.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest
type IOSSystemExtensionRequest interface {
	objectivec.IObject

	// Topic: Working with a Delegate

	// A delegate to receive updates about the progress of a request.
	Delegate() OSSystemExtensionRequestDelegate
	SetDelegate(value OSSystemExtensionRequestDelegate)

	// Topic: Identifying the Target Extension

	// The bundle identifier of the target extension.
	Identifier() string
}

// Init initializes the instance.
func (o OSSystemExtensionRequest) Init() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSSystemExtensionRequest) Autorelease() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionRequest creates a new OSSystemExtensionRequest instance.
func NewOSSystemExtensionRequest() OSSystemExtensionRequest {
	class := getOSSystemExtensionRequestClass()
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a request to activate a System Extension.
//
// identifier: The bundle identifier of the target extension.
//
// queue: The dispatch queue to use when calling delegate methods.
//
// # Return Value
//
// A new extension request.
//
// # Discussion
//
// Create and submit an activation request whenever you want to use a given
// extension. If the extension is inactive, the system may need to prompt the
// user for approval. The request succeeds only after the user gives their
// approval.
//
// If the extension is already active, the request succeeds in short order,
// without significant delay or user interaction. If you request activation of
// a new version of an already-active extension, the system prompts the user
// to resolve the conflict before proceeding.
//
// An activation request may succeed, but also indicate that the extension
// requires a restart to become active. This can occur when replacing an
// extension that required a restart to deactivate. The most recently
// activated extension becomes active when the user restarts their Mac.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/activationRequest(forExtensionWithIdentifier:queue:)
func (_OSSystemExtensionRequestClass OSSystemExtensionRequestClass) ActivationRequestForExtensionQueue(identifier string, queue dispatch.Queue) OSSystemExtensionRequest {
	rv := objc.Send[objc.ID](objc.ID(_OSSystemExtensionRequestClass.class), objc.Sel("activationRequestForExtension:queue:"), objc.String(identifier), uintptr(queue.Handle()))
	return OSSystemExtensionRequestFromID(rv)
}

// Creates a request to deactivate a System Extension.
//
// identifier: The bundle identifier of the extension to deactivate.
//
// queue: The dispatch queue to use when calling delegate methods.
//
// # Discussion
//
// The system discovers existing system extensions in the
// `Contents/Library/SystemExtensions` directory of the main app bundle.
//
// A deactivation request may require a restart before deactivating the
// extension. If the request succeeds but requires a restart to complete, the
// extension may still appear operational until the next restart.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/deactivationRequest(forExtensionWithIdentifier:queue:)
func (_OSSystemExtensionRequestClass OSSystemExtensionRequestClass) DeactivationRequestForExtensionQueue(identifier string, queue dispatch.Queue) OSSystemExtensionRequest {
	rv := objc.Send[objc.ID](objc.ID(_OSSystemExtensionRequestClass.class), objc.Sel("deactivationRequestForExtension:queue:"), objc.String(identifier), uintptr(queue.Handle()))
	return OSSystemExtensionRequestFromID(rv)
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/propertiesRequest(forExtensionWithIdentifier:queue:)
func (_OSSystemExtensionRequestClass OSSystemExtensionRequestClass) PropertiesRequestForExtensionQueue(identifier string, queue dispatch.Queue) OSSystemExtensionRequest {
	rv := objc.Send[objc.ID](objc.ID(_OSSystemExtensionRequestClass.class), objc.Sel("propertiesRequestForExtension:queue:"), objc.String(identifier), uintptr(queue.Handle()))
	return OSSystemExtensionRequestFromID(rv)
}

// A delegate to receive updates about the progress of a request.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/delegate
func (o OSSystemExtensionRequest) Delegate() OSSystemExtensionRequestDelegate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("delegate"))
	return OSSystemExtensionRequestDelegateObjectFromID(rv)
}
func (o OSSystemExtensionRequest) SetDelegate(value OSSystemExtensionRequestDelegate) {
	objc.Send[struct{}](o.ID, objc.Sel("setDelegate:"), value)
}

// The bundle identifier of the target extension.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/identifier
func (o OSSystemExtensionRequest) Identifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
