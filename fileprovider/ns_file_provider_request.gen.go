// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderRequest] class.
var (
	_NSFileProviderRequestClass     NSFileProviderRequestClass
	_NSFileProviderRequestClassOnce sync.Once
)

func getNSFileProviderRequestClass() NSFileProviderRequestClass {
	_NSFileProviderRequestClassOnce.Do(func() {
		_NSFileProviderRequestClass = NSFileProviderRequestClass{class: objc.GetClass("NSFileProviderRequest")}
	})
	return _NSFileProviderRequestClass
}

// GetNSFileProviderRequestClass returns the class object for NSFileProviderRequest.
func GetNSFileProviderRequestClass() NSFileProviderRequestClass {
	return getNSFileProviderRequestClass()
}

type NSFileProviderRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderRequestClass) Alloc() NSFileProviderRequest {
	rv := objc.Send[NSFileProviderRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the application requesting data
// from the File Provider extension.
//
// # Accessing Application Information
//
//   - [NSFileProviderRequest.DomainVersion]: The version of the domain for the request.
//   - [NSFileProviderRequest.RequestingExecutable]: The URL of the requesting executable.
//   - [NSFileProviderRequest.IsFileViewerRequest]: A Boolean value that indicates whether the request came from Finder or related system file browsers.
//   - [NSFileProviderRequest.IsSystemRequest]: A Boolean value that indicates whether the request came from a system process.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest
type NSFileProviderRequest struct {
	objectivec.Object
}

// NSFileProviderRequestFromID constructs a [NSFileProviderRequest] from an objc.ID.
//
// An object that provides information about the application requesting data
// from the File Provider extension.
func NSFileProviderRequestFromID(id objc.ID) NSFileProviderRequest {
	return NSFileProviderRequest{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderRequest] class.
//
// # Accessing Application Information
//
//   - [INSFileProviderRequest.DomainVersion]: The version of the domain for the request.
//   - [INSFileProviderRequest.RequestingExecutable]: The URL of the requesting executable.
//   - [INSFileProviderRequest.IsFileViewerRequest]: A Boolean value that indicates whether the request came from Finder or related system file browsers.
//   - [INSFileProviderRequest.IsSystemRequest]: A Boolean value that indicates whether the request came from a system process.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest
type INSFileProviderRequest interface {
	objectivec.IObject

	// Topic: Accessing Application Information

	// The version of the domain for the request.
	DomainVersion() INSFileProviderDomainVersion
	// The URL of the requesting executable.
	RequestingExecutable() foundation.NSURL
	// A Boolean value that indicates whether the request came from Finder or related system file browsers.
	IsFileViewerRequest() bool
	// A Boolean value that indicates whether the request came from a system process.
	IsSystemRequest() bool
}

// Init initializes the instance.
func (f NSFileProviderRequest) Init() NSFileProviderRequest {
	rv := objc.Send[NSFileProviderRequest](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderRequest) Autorelease() NSFileProviderRequest {
	rv := objc.Send[NSFileProviderRequest](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderRequest creates a new NSFileProviderRequest instance.
func NewNSFileProviderRequest() NSFileProviderRequest {
	class := getNSFileProviderRequestClass()
	rv := objc.Send[NSFileProviderRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The version of the domain for the request.
//
// # Discussion
//
// If the file provider extension doesn’t implement the
// [NSFileProviderDomainState] protocol, this property is `nil`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/domainVersion
func (f NSFileProviderRequest) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(objc.ID(rv))
}

// The URL of the requesting executable.
//
// # Discussion
//
// This property is `nil` unless the device has a Mobile Device Management
// (MDM) profile, and the profile’s administrator installed the file
// provider’s app using the MDM profile.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/requestingExecutable
func (f NSFileProviderRequest) RequestingExecutable() foundation.NSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("requestingExecutable"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the request came from Finder or
// related system file browsers.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/isFileViewerRequest
func (f NSFileProviderRequest) IsFileViewerRequest() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isFileViewerRequest"))
	return rv
}

// A Boolean value that indicates whether the request came from a system
// process.
//
// # Discussion
//
// System requests occur, for example, when the system needs to update a file
// after receiving a push notification about a change from the remote storage.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/isSystemRequest
func (f NSFileProviderRequest) IsSystemRequest() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isSystemRequest"))
	return rv
}
