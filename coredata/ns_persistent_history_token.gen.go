// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentHistoryToken] class.
var (
	_NSPersistentHistoryTokenClass     NSPersistentHistoryTokenClass
	_NSPersistentHistoryTokenClassOnce sync.Once
)

func getNSPersistentHistoryTokenClass() NSPersistentHistoryTokenClass {
	_NSPersistentHistoryTokenClassOnce.Do(func() {
		_NSPersistentHistoryTokenClass = NSPersistentHistoryTokenClass{class: objc.GetClass("NSPersistentHistoryToken")}
	})
	return _NSPersistentHistoryTokenClass
}

// GetNSPersistentHistoryTokenClass returns the class object for NSPersistentHistoryToken.
func GetNSPersistentHistoryTokenClass() NSPersistentHistoryTokenClass {
	return getNSPersistentHistoryTokenClass()
}

type NSPersistentHistoryTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentHistoryTokenClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentHistoryTokenClass) Alloc() NSPersistentHistoryToken {
	rv := objc.Send[NSPersistentHistoryToken](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bookmark for keeping track the most recent history that you’ve
// processed.
//
// # Overview
//
// You can save a token to disk and fetch history when your app loads based on
// that token. See [Keep track of the most recent history] in [Consuming
// relevant store changes].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryToken
//
// [Consuming relevant store changes]: https://developer.apple.com/documentation/CoreData/consuming-relevant-store-changes
// [Keep track of the most recent history]: https://developer.apple.com/documentation/CoreData/consuming-relevant-store-changes#Keep-track-of-the-most-recent-history
type NSPersistentHistoryToken struct {
	objectivec.Object
}

// NSPersistentHistoryTokenFromID constructs a [NSPersistentHistoryToken] from an objc.ID.
//
// A bookmark for keeping track the most recent history that you’ve
// processed.
func NSPersistentHistoryTokenFromID(id objc.ID) NSPersistentHistoryToken {
	return NSPersistentHistoryToken{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentHistoryToken adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentHistoryToken] class.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryToken
type INSPersistentHistoryToken interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPersistentHistoryToken) Init() NSPersistentHistoryToken {
	rv := objc.Send[NSPersistentHistoryToken](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentHistoryToken) Autorelease() NSPersistentHistoryToken {
	rv := objc.Send[NSPersistentHistoryToken](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentHistoryToken creates a new NSPersistentHistoryToken instance.
func NewNSPersistentHistoryToken() NSPersistentHistoryToken {
	class := getNSPersistentHistoryTokenClass()
	rv := objc.Send[NSPersistentHistoryToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p NSPersistentHistoryToken) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}
