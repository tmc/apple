// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderService] class.
var (
	_NSFileProviderServiceClass     NSFileProviderServiceClass
	_NSFileProviderServiceClassOnce sync.Once
)

func getNSFileProviderServiceClass() NSFileProviderServiceClass {
	_NSFileProviderServiceClassOnce.Do(func() {
		_NSFileProviderServiceClass = NSFileProviderServiceClass{class: objc.GetClass("NSFileProviderService")}
	})
	return _NSFileProviderServiceClass
}

// GetNSFileProviderServiceClass returns the class object for NSFileProviderService.
func GetNSFileProviderServiceClass() NSFileProviderServiceClass {
	return getNSFileProviderServiceClass()
}

type NSFileProviderServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderServiceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderServiceClass) Alloc() NSFileProviderService {
	rv := objc.Send[NSFileProviderService](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A service that provides a custom communication channel between your app and
// a File Provider extension.
//
// # Overview
// 
// To communicate, both your app and the File Provider extension must
// implement their part of the service:
// 
// - Your app requests the proxy object, and calls its methods. - The File
// Provider extension declares the supported services and vends a proxy object
// that implements the protocol for each service.
// 
// The app and File Provider extension communicate using an XPC service. This
// service performs actions only on items managed by the File Provider
// extension. For more information, see [Creating XPC Services].
// 
// # Defining the Service’s Protocol
// 
// Services let you define custom actions that are not provided by Apple’s
// APIs. Both the app and the File Provider extension must agree upon the
// service’s name and protocol. Communicate the name and protocol through an
// outside source (for example, posting a header file that defines both the
// name and protocol, or publishing a library that includes them both).
// 
// The service can be defined by either the app or the File Provider
// extension:
// 
// - Apps can define a service for features they would like to use. File
// providers can then choose to support those features by implementing the
// service. - File Provider extensions can provide a service for the features
// they support. Apps can then choose to use the specified service.
// 
// When defining a service’s protocol, the parameters for each method must
// adhere to the following rules:
// 
// - The parameter’s class must conform to [NSSecureCoding]. - The
// parameter’s class must be defined in both the app and the File Provider
// extension (for example, standard system types or classes defined in a
// library imported by both sides). - If a collection parameter contains types
// other than property list types (see [Property List Types and Objects]),
// declare the valid types using the [NSXPCInterface] class’s
// [ClassesForSelectorArgumentIndexOfReply] method.
//
// [Creating XPC Services]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/Chapters/CreatingXPCServices.html#//apple_ref/doc/uid/10000172i-SW6
// [Property List Types and Objects]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/PropertyList.html#//apple_ref/doc/uid/TP40008195-CH44-SW2
//
// # Accessing the Service
//
//   - [NSFileProviderService.Name]: The File Provider service’s name.
//   - [NSFileProviderService.GetFileProviderConnectionWithCompletionHandler]: Asynchronously returns the service’s connection object.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileProviderService
type NSFileProviderService struct {
	objectivec.Object
}

// NSFileProviderServiceFromID constructs a [NSFileProviderService] from an objc.ID.
//
// A service that provides a custom communication channel between your app and
// a File Provider extension.
func NSFileProviderServiceFromID(id objc.ID) NSFileProviderService {
	return NSFileProviderService{objectivec.Object{ID: id}}
}
// NOTE: NSFileProviderService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderService] class.
//
// # Accessing the Service
//
//   - [INSFileProviderService.Name]: The File Provider service’s name.
//   - [INSFileProviderService.GetFileProviderConnectionWithCompletionHandler]: Asynchronously returns the service’s connection object.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileProviderService
type INSFileProviderService interface {
	objectivec.IObject

	// Topic: Accessing the Service

	// The File Provider service’s name.
	Name() NSFileProviderServiceName
	// Asynchronously returns the service’s connection object.
	GetFileProviderConnectionWithCompletionHandler(completionHandler XPCConnectionErrorHandler)
}

// Init initializes the instance.
func (f NSFileProviderService) Init() NSFileProviderService {
	rv := objc.Send[NSFileProviderService](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderService) Autorelease() NSFileProviderService {
	rv := objc.Send[NSFileProviderService](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderService creates a new NSFileProviderService instance.
func NewNSFileProviderService() NSFileProviderService {
	class := getNSFileProviderServiceClass()
	rv := objc.Send[NSFileProviderService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Asynchronously returns the service’s connection object.
//
// completionHandler: A block that is called on an anonymous background queue. The system passes
// this block the following parameters:
// 
// `connection`: An [NSXPCConnection] object for the service, or `nil` if an
// error occurs. `error`: If an error occurs, this property contains an object
// that describes the error; otherwise, it is set to `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileProviderService/getFileProviderConnection(completionHandler:)
func (f NSFileProviderService) GetFileProviderConnectionWithCompletionHandler(completionHandler XPCConnectionErrorHandler) {
_block0, _ := NewXPCConnectionErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("getFileProviderConnectionWithCompletionHandler:"), _block0)
}

// The File Provider service’s name.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileProviderService/name
func (f NSFileProviderService) Name() NSFileProviderServiceName {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return NSFileProviderServiceName(NSStringFromID(rv).String())
}

// GetFileProviderConnection is a synchronous wrapper around [NSFileProviderService.GetFileProviderConnectionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderService) GetFileProviderConnection(ctx context.Context) (*NSXPCConnection, error) {
	type result struct {
		val *NSXPCConnection
		err error
	}
	done := make(chan result, 1)
	f.GetFileProviderConnectionWithCompletionHandler(func(val *NSXPCConnection, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

