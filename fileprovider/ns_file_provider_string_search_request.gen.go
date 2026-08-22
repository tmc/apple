// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderStringSearchRequest] class.
var (
	_NSFileProviderStringSearchRequestClass     NSFileProviderStringSearchRequestClass
	_NSFileProviderStringSearchRequestClassOnce sync.Once
)

func getNSFileProviderStringSearchRequestClass() NSFileProviderStringSearchRequestClass {
	_NSFileProviderStringSearchRequestClassOnce.Do(func() {
		_NSFileProviderStringSearchRequestClass = NSFileProviderStringSearchRequestClass{class: objc.GetClass("NSFileProviderStringSearchRequest")}
	})
	return _NSFileProviderStringSearchRequestClass
}

// GetNSFileProviderStringSearchRequestClass returns the class object for NSFileProviderStringSearchRequest.
func GetNSFileProviderStringSearchRequestClass() NSFileProviderStringSearchRequestClass {
	return getNSFileProviderStringSearchRequestClass()
}

type NSFileProviderStringSearchRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderStringSearchRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderStringSearchRequestClass) Alloc() NSFileProviderStringSearchRequest {
	rv := objc.Send[NSFileProviderStringSearchRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that contains details of a string-based search request.
//
// # Working with request properties
//
//   - [NSFileProviderStringSearchRequest.Query]: A plaintext string, representing the query a person entered into the system search UI.
//
// # Instance Properties
//
//   - [NSFileProviderStringSearchRequest.DesiredNumberOfResults]: How many results the system is requesting. This is a hint to the extension, to help avoid unnecessary work. The extension may return more results than this.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest
type NSFileProviderStringSearchRequest struct {
	objectivec.Object
}

// NSFileProviderStringSearchRequestFromID constructs a [NSFileProviderStringSearchRequest] from an objc.ID.
//
// A type that contains details of a string-based search request.
func NSFileProviderStringSearchRequestFromID(id objc.ID) NSFileProviderStringSearchRequest {
	return NSFileProviderStringSearchRequest{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderStringSearchRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderStringSearchRequest] class.
//
// # Working with request properties
//
//   - [INSFileProviderStringSearchRequest.Query]: A plaintext string, representing the query a person entered into the system search UI.
//
// # Instance Properties
//
//   - [INSFileProviderStringSearchRequest.DesiredNumberOfResults]: How many results the system is requesting. This is a hint to the extension, to help avoid unnecessary work. The extension may return more results than this.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest
type INSFileProviderStringSearchRequest interface {
	objectivec.IObject

	// Topic: Working with request properties

	// A plaintext string, representing the query a person entered into the system search UI.
	Query() string

	// Topic: Instance Properties

	// How many results the system is requesting. This is a hint to the extension, to help avoid unnecessary work. The extension may return more results than this.
	DesiredNumberOfResults() int
}

// Init initializes the instance.
func (f NSFileProviderStringSearchRequest) Init() NSFileProviderStringSearchRequest {
	rv := objc.Send[NSFileProviderStringSearchRequest](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderStringSearchRequest) Autorelease() NSFileProviderStringSearchRequest {
	rv := objc.Send[NSFileProviderStringSearchRequest](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderStringSearchRequest creates a new NSFileProviderStringSearchRequest instance.
func NewNSFileProviderStringSearchRequest() NSFileProviderStringSearchRequest {
	class := getNSFileProviderStringSearchRequestClass()
	rv := objc.Send[NSFileProviderStringSearchRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A plaintext string, representing the query a person entered into the system
// search UI.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest/query
func (f NSFileProviderStringSearchRequest) Query() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("query"))
	return foundation.NSStringFromID(rv).String()
}

// How many results the system is requesting. This is a hint to the extension,
// to help avoid unnecessary work. The extension may return more results than
// this.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest/desiredNumberOfResults
func (f NSFileProviderStringSearchRequest) DesiredNumberOfResults() int {
	rv := objc.Send[int](f.ID, objc.Sel("desiredNumberOfResults"))
	return rv
}
