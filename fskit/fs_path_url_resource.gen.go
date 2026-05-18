// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [FSPathURLResource] class.
var (
	_FSPathURLResourceClass     FSPathURLResourceClass
	_FSPathURLResourceClassOnce sync.Once
)

func getFSPathURLResourceClass() FSPathURLResourceClass {
	_FSPathURLResourceClassOnce.Do(func() {
		_FSPathURLResourceClass = FSPathURLResourceClass{class: objc.GetClass("FSPathURLResource")}
	})
	return _FSPathURLResourceClass
}

// GetFSPathURLResourceClass returns the class object for FSPathURLResource.
func GetFSPathURLResourceClass() FSPathURLResourceClass {
	return getFSPathURLResourceClass()
}

type FSPathURLResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSPathURLResourceClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSPathURLResourceClass) Alloc() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A resource that represents a path in the system file space.
//
// # Overview
//
// The URL passed to [FSPathURLResource] may be a security-scoped URL. If the
// URL is a security-scoped URL, FSKit transports it intact from a client
// application to your extension.
//
// # Creating a path URL resource
//
//   - [FSPathURLResource.InitWithURLWritable]: Creates a path URL resource.
//
// # Accessing resource properties
//
//   - [FSPathURLResource.Url]: The URL represented by the resource.
//   - [FSPathURLResource.IsWritable]: A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource
type FSPathURLResource struct {
	FSResource
}

// FSPathURLResourceFromID constructs a [FSPathURLResource] from an objc.ID.
//
// A resource that represents a path in the system file space.
func FSPathURLResourceFromID(id objc.ID) FSPathURLResource {
	return FSPathURLResource{FSResource: FSResourceFromID(id)}
}

// NOTE: FSPathURLResource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSPathURLResource] class.
//
// # Creating a path URL resource
//
//   - [IFSPathURLResource.InitWithURLWritable]: Creates a path URL resource.
//
// # Accessing resource properties
//
//   - [IFSPathURLResource.Url]: The URL represented by the resource.
//   - [IFSPathURLResource.IsWritable]: A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource
type IFSPathURLResource interface {
	IFSResource

	// Topic: Creating a path URL resource

	// Creates a path URL resource.
	InitWithURLWritable(URL foundation.NSURL, writable bool) FSPathURLResource

	// Topic: Accessing resource properties

	// The URL represented by the resource.
	Url() foundation.NSURL
	// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
	IsWritable() bool
}

// Init initializes the instance.
func (p FSPathURLResource) Init() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p FSPathURLResource) Autorelease() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSPathURLResource creates a new FSPathURLResource instance.
func NewFSPathURLResource() FSPathURLResource {
	class := getFSPathURLResourceClass()
	rv := objc.Send[FSPathURLResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a path URL resource.
//
// URL: A URL in the system file space that represents the contents of a file
// system. This parameter uses the “ scheme.
//
// writable: A Boolean value that indicates whether the file system supports writing to
// the contents of the URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource/init(url:writable:)-2l10q
func NewPathURLResourceWithURLWritable(URL foundation.NSURL, writable bool) FSPathURLResource {
	instance := getFSPathURLResourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:writable:"), URL, writable)
	return FSPathURLResourceFromID(rv)
}

// Creates a path URL resource.
//
// URL: A URL in the system file space that represents the contents of a file
// system. This parameter uses the “ scheme.
//
// writable: A Boolean value that indicates whether the file system supports writing to
// the contents of the URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource/init(url:writable:)-2l10q
func (p FSPathURLResource) InitWithURLWritable(URL foundation.NSURL, writable bool) FSPathURLResource {
	rv := objc.Send[FSPathURLResource](p.ID, objc.Sel("initWithURL:writable:"), URL, writable)
	return rv
}

// The URL represented by the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource/url
func (p FSPathURLResource) Url() foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the file system supports writing to
// the contents of the path URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSPathURLResource/isWritable
func (p FSPathURLResource) IsWritable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isWritable"))
	return rv
}
