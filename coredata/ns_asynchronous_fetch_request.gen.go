// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSAsynchronousFetchRequest] class.
var (
	_NSAsynchronousFetchRequestClass     NSAsynchronousFetchRequestClass
	_NSAsynchronousFetchRequestClassOnce sync.Once
)

func getNSAsynchronousFetchRequestClass() NSAsynchronousFetchRequestClass {
	_NSAsynchronousFetchRequestClassOnce.Do(func() {
		_NSAsynchronousFetchRequestClass = NSAsynchronousFetchRequestClass{class: objc.GetClass("NSAsynchronousFetchRequest")}
	})
	return _NSAsynchronousFetchRequestClass
}

// GetNSAsynchronousFetchRequestClass returns the class object for NSAsynchronousFetchRequest.
func GetNSAsynchronousFetchRequestClass() NSAsynchronousFetchRequestClass {
	return getNSAsynchronousFetchRequestClass()
}

type NSAsynchronousFetchRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAsynchronousFetchRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAsynchronousFetchRequestClass) Alloc() NSAsynchronousFetchRequest {
	rv := objc.Send[NSAsynchronousFetchRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A fetch request that retrieves results asynchronously and supports progress
// notification.
//
// # Initializing a Request
//
//   - [NSAsynchronousFetchRequest.InitWithFetchRequestCompletionBlock]: Initializes a new asynchronous fetch request configured with the provided fetch request and completion block.
//
// # Preparing a Request
//
//   - [NSAsynchronousFetchRequest.EstimatedResultCount]: A configuration parameter that assists Core Data with scheduling the asynchronous fetch request.
//   - [NSAsynchronousFetchRequest.SetEstimatedResultCount]
//   - [NSAsynchronousFetchRequest.FetchRequest]: The underlying fetch request that is executed asynchronously.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest
type NSAsynchronousFetchRequest struct {
	NSPersistentStoreRequest
}

// NSAsynchronousFetchRequestFromID constructs a [NSAsynchronousFetchRequest] from an objc.ID.
//
// A fetch request that retrieves results asynchronously and supports progress
// notification.
func NSAsynchronousFetchRequestFromID(id objc.ID) NSAsynchronousFetchRequest {
	return NSAsynchronousFetchRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSAsynchronousFetchRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAsynchronousFetchRequest] class.
//
// # Initializing a Request
//
//   - [INSAsynchronousFetchRequest.InitWithFetchRequestCompletionBlock]: Initializes a new asynchronous fetch request configured with the provided fetch request and completion block.
//
// # Preparing a Request
//
//   - [INSAsynchronousFetchRequest.EstimatedResultCount]: A configuration parameter that assists Core Data with scheduling the asynchronous fetch request.
//   - [INSAsynchronousFetchRequest.SetEstimatedResultCount]
//   - [INSAsynchronousFetchRequest.FetchRequest]: The underlying fetch request that is executed asynchronously.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest
type INSAsynchronousFetchRequest interface {
	INSPersistentStoreRequest

	// Topic: Initializing a Request

	// Initializes a new asynchronous fetch request configured with the provided fetch request and completion block.
	InitWithFetchRequestCompletionBlock(request INSFetchRequest, blk idNSFetchRequestResultAsynchronousFetchResultHandler) NSAsynchronousFetchRequest

	// Topic: Preparing a Request

	// A configuration parameter that assists Core Data with scheduling the asynchronous fetch request.
	EstimatedResultCount() int
	SetEstimatedResultCount(value int)
	// The underlying fetch request that is executed asynchronously.
	FetchRequest() INSFetchRequest
}

// Init initializes the instance.
func (a NSAsynchronousFetchRequest) Init() NSAsynchronousFetchRequest {
	rv := objc.Send[NSAsynchronousFetchRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAsynchronousFetchRequest) Autorelease() NSAsynchronousFetchRequest {
	rv := objc.Send[NSAsynchronousFetchRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAsynchronousFetchRequest creates a new NSAsynchronousFetchRequest instance.
func NewNSAsynchronousFetchRequest() NSAsynchronousFetchRequest {
	class := getNSAsynchronousFetchRequestClass()
	rv := objc.Send[NSAsynchronousFetchRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new asynchronous fetch request configured with the provided
// fetch request and completion block.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest/init(fetchRequest:completionBlock:)
func (a NSAsynchronousFetchRequest) InitWithFetchRequestCompletionBlock(request INSFetchRequest, blk idNSFetchRequestResultAsynchronousFetchResultHandler) NSAsynchronousFetchRequest {
	_block1, _ := NewidNSFetchRequestResultAsynchronousFetchResultBlock(blk)
	rv := objc.Send[NSAsynchronousFetchRequest](a.ID, objc.Sel("initWithFetchRequest:completionBlock:"), request, _block1)
	return rv
}

// A configuration parameter that assists Core Data with scheduling the
// asynchronous fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest/estimatedResultCount
func (a NSAsynchronousFetchRequest) EstimatedResultCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("estimatedResultCount"))
	return rv
}
func (a NSAsynchronousFetchRequest) SetEstimatedResultCount(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setEstimatedResultCount:"), value)
}

// The underlying fetch request that is executed asynchronously.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest/fetchRequest
func (a NSAsynchronousFetchRequest) FetchRequest() INSFetchRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}

// InitWithFetchRequestCompletionBlockSync is a synchronous wrapper around [NSAsynchronousFetchRequest.InitWithFetchRequestCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSAsynchronousFetchRequest) InitWithFetchRequestCompletionBlockSync(ctx context.Context, request INSFetchRequest) (*NSAsynchronousFetchResult, error) {
	done := make(chan *NSAsynchronousFetchResult, 1)
	a.InitWithFetchRequestCompletionBlock(request, func(val *NSAsynchronousFetchResult) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
