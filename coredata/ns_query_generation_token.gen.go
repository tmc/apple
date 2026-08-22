// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSQueryGenerationToken] class.
var (
	_NSQueryGenerationTokenClass     NSQueryGenerationTokenClass
	_NSQueryGenerationTokenClassOnce sync.Once
)

func getNSQueryGenerationTokenClass() NSQueryGenerationTokenClass {
	_NSQueryGenerationTokenClassOnce.Do(func() {
		_NSQueryGenerationTokenClass = NSQueryGenerationTokenClass{class: objc.GetClass("NSQueryGenerationToken")}
	})
	return _NSQueryGenerationTokenClass
}

// GetNSQueryGenerationTokenClass returns the class object for NSQueryGenerationToken.
func GetNSQueryGenerationTokenClass() NSQueryGenerationTokenClass {
	return getNSQueryGenerationTokenClass()
}

type NSQueryGenerationTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSQueryGenerationTokenClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSQueryGenerationTokenClass) Alloc() NSQueryGenerationToken {
	rv := objc.Send[NSQueryGenerationToken](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A token that indicates which generation of the persistent store is being
// accessed.
//
// # Overview
//
// When a managed object context is pinned to a specific generation of the app
// data, a query generation token will be associated with that context.
//
// See: https://developer.apple.com/documentation/CoreData/NSQueryGenerationToken
type NSQueryGenerationToken struct {
	objectivec.Object
}

// NSQueryGenerationTokenFromID constructs a [NSQueryGenerationToken] from an objc.ID.
//
// A token that indicates which generation of the persistent store is being
// accessed.
func NSQueryGenerationTokenFromID(id objc.ID) NSQueryGenerationToken {
	return NSQueryGenerationToken{objectivec.Object{ID: id}}
}

// NOTE: NSQueryGenerationToken adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSQueryGenerationToken] class.
//
// See: https://developer.apple.com/documentation/CoreData/NSQueryGenerationToken
type INSQueryGenerationToken interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (q NSQueryGenerationToken) Init() NSQueryGenerationToken {
	rv := objc.Send[NSQueryGenerationToken](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q NSQueryGenerationToken) Autorelease() NSQueryGenerationToken {
	rv := objc.Send[NSQueryGenerationToken](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSQueryGenerationToken creates a new NSQueryGenerationToken instance.
func NewNSQueryGenerationToken() NSQueryGenerationToken {
	class := getNSQueryGenerationTokenClass()
	rv := objc.Send[NSQueryGenerationToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (q NSQueryGenerationToken) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](q.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A token that informs a context to use the current generation.
//
// See: https://developer.apple.com/documentation/CoreData/NSQueryGenerationToken/current
func (_NSQueryGenerationTokenClass NSQueryGenerationTokenClass) CurrentQueryGenerationToken() NSQueryGenerationToken {
	rv := objc.Send[objc.ID](objc.ID(_NSQueryGenerationTokenClass.class), objc.Sel("currentQueryGenerationToken"))
	return NSQueryGenerationTokenFromID(objc.ID(rv))
}
