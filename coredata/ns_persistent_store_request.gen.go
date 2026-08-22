// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentStoreRequest] class.
var (
	_NSPersistentStoreRequestClass     NSPersistentStoreRequestClass
	_NSPersistentStoreRequestClassOnce sync.Once
)

func getNSPersistentStoreRequestClass() NSPersistentStoreRequestClass {
	_NSPersistentStoreRequestClassOnce.Do(func() {
		_NSPersistentStoreRequestClass = NSPersistentStoreRequestClass{class: objc.GetClass("NSPersistentStoreRequest")}
	})
	return _NSPersistentStoreRequestClass
}

// GetNSPersistentStoreRequestClass returns the class object for NSPersistentStoreRequest.
func GetNSPersistentStoreRequestClass() NSPersistentStoreRequestClass {
	return getNSPersistentStoreRequestClass()
}

type NSPersistentStoreRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreRequestClass) Alloc() NSPersistentStoreRequest {
	rv := objc.Send[NSPersistentStoreRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Criteria used to retrieve data from or save data to a persistent store.
//
// # Configuring a Request
//
//   - [NSPersistentStoreRequest.AffectedStores]: The stores the request should be sent to.
//   - [NSPersistentStoreRequest.SetAffectedStores]
//   - [NSPersistentStoreRequest.RequestType]: The type of the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest
type NSPersistentStoreRequest struct {
	objectivec.Object
}

// NSPersistentStoreRequestFromID constructs a [NSPersistentStoreRequest] from an objc.ID.
//
// Criteria used to retrieve data from or save data to a persistent store.
func NSPersistentStoreRequestFromID(id objc.ID) NSPersistentStoreRequest {
	return NSPersistentStoreRequest{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentStoreRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStoreRequest] class.
//
// # Configuring a Request
//
//   - [INSPersistentStoreRequest.AffectedStores]: The stores the request should be sent to.
//   - [INSPersistentStoreRequest.SetAffectedStores]
//   - [INSPersistentStoreRequest.RequestType]: The type of the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest
type INSPersistentStoreRequest interface {
	objectivec.IObject

	// Topic: Configuring a Request

	// The stores the request should be sent to.
	AffectedStores() []NSPersistentStore
	SetAffectedStores(value []NSPersistentStore)
	// The type of the fetch request.
	RequestType() NSPersistentStoreRequestType
}

// Init initializes the instance.
func (p NSPersistentStoreRequest) Init() NSPersistentStoreRequest {
	rv := objc.Send[NSPersistentStoreRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStoreRequest) Autorelease() NSPersistentStoreRequest {
	rv := objc.Send[NSPersistentStoreRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStoreRequest creates a new NSPersistentStoreRequest instance.
func NewNSPersistentStoreRequest() NSPersistentStoreRequest {
	class := getNSPersistentStoreRequestClass()
	rv := objc.Send[NSPersistentStoreRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The stores the request should be sent to.
//
// # Discussion
//
// The array contains instances of [NSPersistentStore].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest/affectedStores
func (p NSPersistentStoreRequest) AffectedStores() []NSPersistentStore {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("affectedStores"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPersistentStore {
		return NSPersistentStoreFromID(id)
	})
}
func (p NSPersistentStoreRequest) SetAffectedStores(value []NSPersistentStore) {
	objc.Send[struct{}](p.ID, objc.Sel("setAffectedStores:"), objectivec.IObjectSliceToNSArray(value))
}

// The type of the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest/requestType
func (p NSPersistentStoreRequest) RequestType() NSPersistentStoreRequestType {
	rv := objc.Send[NSPersistentStoreRequestType](p.ID, objc.Sel("requestType"))
	return NSPersistentStoreRequestType(rv)
}
