// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAsynchronousFetchResult] class.
var (
	_NSAsynchronousFetchResultClass     NSAsynchronousFetchResultClass
	_NSAsynchronousFetchResultClassOnce sync.Once
)

func getNSAsynchronousFetchResultClass() NSAsynchronousFetchResultClass {
	_NSAsynchronousFetchResultClassOnce.Do(func() {
		_NSAsynchronousFetchResultClass = NSAsynchronousFetchResultClass{class: objc.GetClass("NSAsynchronousFetchResult")}
	})
	return _NSAsynchronousFetchResultClass
}

// GetNSAsynchronousFetchResultClass returns the class object for NSAsynchronousFetchResult.
func GetNSAsynchronousFetchResultClass() NSAsynchronousFetchResultClass {
	return getNSAsynchronousFetchResultClass()
}

type NSAsynchronousFetchResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAsynchronousFetchResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAsynchronousFetchResultClass) Alloc() NSAsynchronousFetchResult {
	rv := objc.Send[NSAsynchronousFetchResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A fetch result object that encompasses the response from an executed
// asynchronous fetch request.
//
// # Getting Information About a Result
//
//   - [NSAsynchronousFetchResult.FetchRequest]: The underlying fetch request that was executed.
//   - [NSAsynchronousFetchResult.FinalResult]: The results that were received from the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchResult
type NSAsynchronousFetchResult struct {
	NSPersistentStoreAsynchronousResult
}

// NSAsynchronousFetchResultFromID constructs a [NSAsynchronousFetchResult] from an objc.ID.
//
// A fetch result object that encompasses the response from an executed
// asynchronous fetch request.
func NSAsynchronousFetchResultFromID(id objc.ID) NSAsynchronousFetchResult {
	return NSAsynchronousFetchResult{NSPersistentStoreAsynchronousResult: NSPersistentStoreAsynchronousResultFromID(id)}
}

// NOTE: NSAsynchronousFetchResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAsynchronousFetchResult] class.
//
// # Getting Information About a Result
//
//   - [INSAsynchronousFetchResult.FetchRequest]: The underlying fetch request that was executed.
//   - [INSAsynchronousFetchResult.FinalResult]: The results that were received from the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchResult
type INSAsynchronousFetchResult interface {
	INSPersistentStoreAsynchronousResult

	// Topic: Getting Information About a Result

	// The underlying fetch request that was executed.
	FetchRequest() INSAsynchronousFetchRequest
	// The results that were received from the fetch request.
	FinalResult() []objectivec.IObject
}

// Init initializes the instance.
func (a NSAsynchronousFetchResult) Init() NSAsynchronousFetchResult {
	rv := objc.Send[NSAsynchronousFetchResult](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAsynchronousFetchResult) Autorelease() NSAsynchronousFetchResult {
	rv := objc.Send[NSAsynchronousFetchResult](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAsynchronousFetchResult creates a new NSAsynchronousFetchResult instance.
func NewNSAsynchronousFetchResult() NSAsynchronousFetchResult {
	class := getNSAsynchronousFetchResultClass()
	rv := objc.Send[NSAsynchronousFetchResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The underlying fetch request that was executed.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchResult/fetchRequest
func (a NSAsynchronousFetchResult) FetchRequest() INSAsynchronousFetchRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fetchRequest"))
	return NSAsynchronousFetchRequestFromID(objc.ID(rv))
}

// The results that were received from the fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchResult/finalResult
func (a NSAsynchronousFetchResult) FinalResult() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("finalResult"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
