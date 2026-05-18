// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [FSGenericURLResource] class.
var (
	_FSGenericURLResourceClass     FSGenericURLResourceClass
	_FSGenericURLResourceClassOnce sync.Once
)

func getFSGenericURLResourceClass() FSGenericURLResourceClass {
	_FSGenericURLResourceClassOnce.Do(func() {
		_FSGenericURLResourceClass = FSGenericURLResourceClass{class: objc.GetClass("FSGenericURLResource")}
	})
	return _FSGenericURLResourceClass
}

// GetFSGenericURLResourceClass returns the class object for FSGenericURLResource.
func GetFSGenericURLResourceClass() FSGenericURLResourceClass {
	return getFSGenericURLResourceClass()
}

type FSGenericURLResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSGenericURLResourceClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSGenericURLResourceClass) Alloc() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A resource that represents an abstract URL.
//
// # Overview
//
// An [FSGenericURLResource] is a completely abstract resource. The only
// reference to its contents is a single URL, the contents of which are
// arbitrary. This URL might represent a PCI locator string like
// `/pci@f0000000/usb@5`, or some sort of network address for a remote file
// system. FSKit leaves interpretation of the URL and its contents entirely up
// to your implementation.
//
// Use the `Info.Plist()` key [FSSupportedSchemes] to provide an array of
// case-insensitive URL schemes that your implementation supports. The
// following example shows how a hypothetical [FSGenericURLResource]
// implementation declares support for the `rsh` and `ssh` URL schemes:
//
// # Creating a generic URL resource
//
//   - [FSGenericURLResource.InitWithURL]: Creates a generic URL resource with the given URL.
//
// # Accessing resource properties
//
//   - [FSGenericURLResource.Url]: The URL represented by the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSGenericURLResource
type FSGenericURLResource struct {
	FSResource
}

// FSGenericURLResourceFromID constructs a [FSGenericURLResource] from an objc.ID.
//
// A resource that represents an abstract URL.
func FSGenericURLResourceFromID(id objc.ID) FSGenericURLResource {
	return FSGenericURLResource{FSResource: FSResourceFromID(id)}
}

// NOTE: FSGenericURLResource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSGenericURLResource] class.
//
// # Creating a generic URL resource
//
//   - [IFSGenericURLResource.InitWithURL]: Creates a generic URL resource with the given URL.
//
// # Accessing resource properties
//
//   - [IFSGenericURLResource.Url]: The URL represented by the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSGenericURLResource
type IFSGenericURLResource interface {
	IFSResource

	// Topic: Creating a generic URL resource

	// Creates a generic URL resource with the given URL.
	InitWithURL(url foundation.NSURL) FSGenericURLResource

	// Topic: Accessing resource properties

	// The URL represented by the resource.
	Url() foundation.NSURL
}

// Init initializes the instance.
func (g FSGenericURLResource) Init() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g FSGenericURLResource) Autorelease() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSGenericURLResource creates a new FSGenericURLResource instance.
func NewFSGenericURLResource() FSGenericURLResource {
	class := getFSGenericURLResourceClass()
	rv := objc.Send[FSGenericURLResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a generic URL resource with the given URL.
//
// url: A URL that provides the content of the file system. The format of this URL
// is completely arbitrary. It’s up to your extension to access the contents
// represented by the URL and make them available as an [FSVolume] that FSKit
// can load.
//
// See: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/init(url:)-2cmhi
func NewGenericURLResourceWithURL(url foundation.NSURL) FSGenericURLResource {
	instance := getFSGenericURLResourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return FSGenericURLResourceFromID(rv)
}

// Creates a generic URL resource with the given URL.
//
// url: A URL that provides the content of the file system. The format of this URL
// is completely arbitrary. It’s up to your extension to access the contents
// represented by the URL and make them available as an [FSVolume] that FSKit
// can load.
//
// See: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/init(url:)-2cmhi
func (g FSGenericURLResource) InitWithURL(url foundation.NSURL) FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](g.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// The URL represented by the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/url
func (g FSGenericURLResource) Url() foundation.NSURL {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
