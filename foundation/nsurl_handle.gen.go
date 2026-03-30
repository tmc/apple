// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLHandle] class.
var (
	_NSURLHandleClass     NSURLHandleClass
	_NSURLHandleClassOnce sync.Once
)

func getNSURLHandleClass() NSURLHandleClass {
	_NSURLHandleClassOnce.Do(func() {
		_NSURLHandleClass = NSURLHandleClass{class: objc.GetClass("NSURLHandle")}
	})
	return _NSURLHandleClass
}

// GetNSURLHandleClass returns the class object for NSURLHandle.
func GetNSURLHandleClass() NSURLHandleClass {
	return getNSURLHandleClass()
}

type NSURLHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSURLHandleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLHandleClass) Alloc() NSURLHandle {
	rv := objc.Send[NSURLHandle](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that accesses and manages resource data indicated by a URL.
//
// # Overview
//
// A single [NSURLHandle] can service multiple equivalent [NSURL] objects, but
// only if these URLs map to the same resource.
//
// # Overview
//
// Cocoa provides private concrete subclasses to handle HTTP and file URL
// schemes. If you want to implement support for additional URL schemes, you
// would do so by creating a subclass of [NSURLHandle]. You can use [NSURL]
// and [NSURLHandle] to download from FTP sites without subclassing.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLHandle
type NSURLHandle struct {
	objectivec.Object
}

// NSURLHandleFromID constructs a [NSURLHandle] from an objc.ID.
//
// An object that accesses and manages resource data indicated by a URL.
func NSURLHandleFromID(id objc.ID) NSURLHandle {
	return NSURLHandle{objectivec.Object{ID: id}}
}

// NOTE: NSURLHandle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURLHandle] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLHandle
type INSURLHandle interface {
	objectivec.IObject
}

// Init initializes the instance.
func (u NSURLHandle) Init() NSURLHandle {
	rv := objc.Send[NSURLHandle](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLHandle) Autorelease() NSURLHandle {
	rv := objc.Send[NSURLHandle](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLHandle creates a new NSURLHandle instance.
func NewNSURLHandle() NSURLHandle {
	class := getNSURLHandleClass()
	rv := objc.Send[NSURLHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}
