// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPersistentHistoryChangeRequest] class.
var (
	_NSPersistentHistoryChangeRequestClass     NSPersistentHistoryChangeRequestClass
	_NSPersistentHistoryChangeRequestClassOnce sync.Once
)

func getNSPersistentHistoryChangeRequestClass() NSPersistentHistoryChangeRequestClass {
	_NSPersistentHistoryChangeRequestClassOnce.Do(func() {
		_NSPersistentHistoryChangeRequestClass = NSPersistentHistoryChangeRequestClass{class: objc.GetClass("NSPersistentHistoryChangeRequest")}
	})
	return _NSPersistentHistoryChangeRequestClass
}

// GetNSPersistentHistoryChangeRequestClass returns the class object for NSPersistentHistoryChangeRequest.
func GetNSPersistentHistoryChangeRequestClass() NSPersistentHistoryChangeRequestClass {
	return getNSPersistentHistoryChangeRequestClass()
}

type NSPersistentHistoryChangeRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentHistoryChangeRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentHistoryChangeRequestClass) Alloc() NSPersistentHistoryChangeRequest {
	rv := objc.Send[NSPersistentHistoryChangeRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A request to fetch or purge persistent history.
//
// # Configuring the Request
//
//   - [NSPersistentHistoryChangeRequest.FetchRequest]: The specified fetch request, when retrieving history.
//   - [NSPersistentHistoryChangeRequest.SetFetchRequest]
//   - [NSPersistentHistoryChangeRequest.ResultType]: The type of result that this request returns.
//   - [NSPersistentHistoryChangeRequest.SetResultType]
//
// # Getting the Token
//
//   - [NSPersistentHistoryChangeRequest.Token]: The specified token, when retrieving history defined by a token.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest
type NSPersistentHistoryChangeRequest struct {
	NSPersistentStoreRequest
}

// NSPersistentHistoryChangeRequestFromID constructs a [NSPersistentHistoryChangeRequest] from an objc.ID.
//
// A request to fetch or purge persistent history.
func NSPersistentHistoryChangeRequestFromID(id objc.ID) NSPersistentHistoryChangeRequest {
	return NSPersistentHistoryChangeRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSPersistentHistoryChangeRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentHistoryChangeRequest] class.
//
// # Configuring the Request
//
//   - [INSPersistentHistoryChangeRequest.FetchRequest]: The specified fetch request, when retrieving history.
//   - [INSPersistentHistoryChangeRequest.SetFetchRequest]
//   - [INSPersistentHistoryChangeRequest.ResultType]: The type of result that this request returns.
//   - [INSPersistentHistoryChangeRequest.SetResultType]
//
// # Getting the Token
//
//   - [INSPersistentHistoryChangeRequest.Token]: The specified token, when retrieving history defined by a token.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest
type INSPersistentHistoryChangeRequest interface {
	INSPersistentStoreRequest

	// Topic: Configuring the Request

	// The specified fetch request, when retrieving history.
	FetchRequest() INSFetchRequest
	SetFetchRequest(value INSFetchRequest)
	// The type of result that this request returns.
	ResultType() NSPersistentHistoryResultType
	SetResultType(value NSPersistentHistoryResultType)

	// Topic: Getting the Token

	// The specified token, when retrieving history defined by a token.
	Token() INSPersistentHistoryToken
}

// Init initializes the instance.
func (p NSPersistentHistoryChangeRequest) Init() NSPersistentHistoryChangeRequest {
	rv := objc.Send[NSPersistentHistoryChangeRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentHistoryChangeRequest) Autorelease() NSPersistentHistoryChangeRequest {
	rv := objc.Send[NSPersistentHistoryChangeRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentHistoryChangeRequest creates a new NSPersistentHistoryChangeRequest instance.
func NewNSPersistentHistoryChangeRequest() NSPersistentHistoryChangeRequest {
	class := getNSPersistentHistoryChangeRequestClass()
	rv := objc.Send[NSPersistentHistoryChangeRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves history since a given date.
//
// date: The date used to define the start of the fetch history.
//
// # Return Value
//
// A persistent history fetch request ([NSPersistentHistoryChangeRequest])
// with an initial date boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-qi5b
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) FetchHistoryAfterDate(date foundation.NSDate) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("fetchHistoryAfterDate:"), date)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Retrieves the request history after a given token.
//
// token: The bookmark that defines the start of the request history.
//
// # Return Value
//
// A persistent history fetch request ([NSPersistentHistoryChangeRequest])
// with an initial token bookmark boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-3rmfm
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) FetchHistoryAfterToken(token INSPersistentHistoryToken) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("fetchHistoryAfterToken:"), token)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Retrieves history since a given transaction.
//
// transaction: The transaction that marks the beginning of the history request.
//
// # Return Value
//
// A persistent history fetch request ([NSPersistentHistoryChangeRequest])
// with an initial transaction boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-9cuj5
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) FetchHistoryAfterTransaction(transaction INSPersistentHistoryTransaction) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("fetchHistoryAfterTransaction:"), transaction)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Retrieves history based on a fetch request.
//
// fetchRequest: The fetch request that defines the history bounds.
//
// # Return Value
//
// A persistent history fetch request ([NSPersistentHistoryChangeRequest])
// built using an existing fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(withFetch:)
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) FetchHistoryWithFetchRequest(fetchRequest INSFetchRequest) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("fetchHistoryWithFetchRequest:"), fetchRequest)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Purges history older than a given date.
//
// date: The date used to define the end of the delete history request.
//
// # Return Value
//
// A delete history change request ([NSPersistentHistoryChangeRequest]) using
// an end date boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/deleteHistory(before:)-7t2th
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) DeleteHistoryBeforeDate(date foundation.NSDate) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("deleteHistoryBeforeDate:"), date)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Purges history older than that defined by a given token.
//
// token: The bookmark that marks the end of the delete history request.
//
// # Return Value
//
// A delete history change request ([NSPersistentHistoryChangeRequest]) using
// an end token bookmark boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/deleteHistory(before:)-5kghb
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) DeleteHistoryBeforeToken(token INSPersistentHistoryToken) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("deleteHistoryBeforeToken:"), token)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// Purges history older than a given transaction.
//
// transaction: The transaction that marks the end of the delete history request.
//
// # Return Value
//
// A delete history change request ([NSPersistentHistoryChangeRequest]) using
// an end transaction boundary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/deleteHistory(before:)-9l06p
func (_NSPersistentHistoryChangeRequestClass NSPersistentHistoryChangeRequestClass) DeleteHistoryBeforeTransaction(transaction INSPersistentHistoryTransaction) NSPersistentHistoryChangeRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryChangeRequestClass.class), objc.Sel("deleteHistoryBeforeTransaction:"), transaction)
	return NSPersistentHistoryChangeRequestFromID(rv)
}

// The specified fetch request, when retrieving history.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchRequest
func (p NSPersistentHistoryChangeRequest) FetchRequest() INSFetchRequest {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}
func (p NSPersistentHistoryChangeRequest) SetFetchRequest(value INSFetchRequest) {
	objc.Send[struct{}](p.ID, objc.Sel("setFetchRequest:"), value)
}

// The type of result that this request returns.
//
// # Discussion
//
// This value defaults to
// [NSPersistentHistoryResultType.transactionsAndChanges].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/resultType
//
// [NSPersistentHistoryResultType.transactionsAndChanges]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResultType/transactionsAndChanges
func (p NSPersistentHistoryChangeRequest) ResultType() NSPersistentHistoryResultType {
	rv := objc.Send[NSPersistentHistoryResultType](p.ID, objc.Sel("resultType"))
	return NSPersistentHistoryResultType(rv)
}
func (p NSPersistentHistoryChangeRequest) SetResultType(value NSPersistentHistoryResultType) {
	objc.Send[struct{}](p.ID, objc.Sel("setResultType:"), value)
}

// The specified token, when retrieving history defined by a token.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/token
func (p NSPersistentHistoryChangeRequest) Token() INSPersistentHistoryToken {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("token"))
	return NSPersistentHistoryTokenFromID(objc.ID(rv))
}
