// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileAccessIntent] class.
var (
	_NSFileAccessIntentClass     NSFileAccessIntentClass
	_NSFileAccessIntentClassOnce sync.Once
)

func getNSFileAccessIntentClass() NSFileAccessIntentClass {
	_NSFileAccessIntentClassOnce.Do(func() {
		_NSFileAccessIntentClass = NSFileAccessIntentClass{class: objc.GetClass("NSFileAccessIntent")}
	})
	return _NSFileAccessIntentClass
}

// GetNSFileAccessIntentClass returns the class object for NSFileAccessIntent.
func GetNSFileAccessIntentClass() NSFileAccessIntentClass {
	return getNSFileAccessIntentClass()
}

type NSFileAccessIntentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileAccessIntentClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileAccessIntentClass) Alloc() NSFileAccessIntent {
	rv := objc.Send[NSFileAccessIntent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The details of a coordinated-read or coordinated-write operation.
//
// # Overview
// 
// Use this class when performing asynchronous operations with a file
// coordinator using the coordinator’s
// [CoordinateAccessWithIntentsQueueByAccessor] method.
//
// # Accessing the Current URL
//
//   - [NSFileAccessIntent.URL]: The current URL for the item managed by the file access intent instance. (read-only)
//
// See: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent
type NSFileAccessIntent struct {
	objectivec.Object
}

// NSFileAccessIntentFromID constructs a [NSFileAccessIntent] from an objc.ID.
//
// The details of a coordinated-read or coordinated-write operation.
func NSFileAccessIntentFromID(id objc.ID) NSFileAccessIntent {
	return NSFileAccessIntent{objectivec.Object{ID: id}}
}
// NOTE: NSFileAccessIntent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileAccessIntent] class.
//
// # Accessing the Current URL
//
//   - [INSFileAccessIntent.URL]: The current URL for the item managed by the file access intent instance. (read-only)
//
// See: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent
type INSFileAccessIntent interface {
	objectivec.IObject

	// Topic: Accessing the Current URL

	// The current URL for the item managed by the file access intent instance. (read-only)
	URL() INSURL
}

// Init initializes the instance.
func (f NSFileAccessIntent) Init() NSFileAccessIntent {
	rv := objc.Send[NSFileAccessIntent](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileAccessIntent) Autorelease() NSFileAccessIntent {
	rv := objc.Send[NSFileAccessIntent](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileAccessIntent creates a new NSFileAccessIntent instance.
func NewNSFileAccessIntent() NSFileAccessIntent {
	class := getNSFileAccessIntentClass()
	rv := objc.Send[NSFileAccessIntent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a file access intent object for reading the given URL with the
// provided options.
//
// url: The URL of the document you intend to read from.
//
// options: The coordinated reading options. For a list of valid values, see
// [NSFileCoordinator.ReadingOptions] in the [NSFileCoordinator].
// //
// [NSFileCoordinator.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
//
// # Return Value
// 
// A newly instantiated and configured file access intent object.
//
// # Discussion
// 
// When calling a file coordinator’s
// [CoordinateAccessWithIntentsQueueByAccessor] method, you pass an array of
// file access intent objects. Each intent object represents a specific read
// or write operation on a single document or directory. Use `` to create an
// intent object suitable for reading.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/readingIntent(with:options:)
func (_NSFileAccessIntentClass NSFileAccessIntentClass) ReadingIntentWithURLOptions(url INSURL, options NSFileCoordinatorReadingOptions) NSFileAccessIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSFileAccessIntentClass.class), objc.Sel("readingIntentWithURL:options:"), url, options)
	return NSFileAccessIntentFromID(rv)
}
// Returns a file access intent object for writing to the given URL with the
// provided options.
//
// url: The URL of the document you intend to write to.
//
// options: The coordinated writing options. For a list of valid values, see
// [NSFileCoordinator.WritingOptions] in the [NSFileCoordinator].
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// # Return Value
// 
// A newly instantiated and configured file access intent object.
//
// # Discussion
// 
// When calling a file coordinator’s
// [CoordinateAccessWithIntentsQueueByAccessor] method, you pass an array of
// file access intent objects. Each intent object represents a specific read
// or write operation on a single document or directory. Use `` to create an
// intent object suitable for writing.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/writingIntent(with:options:)
func (_NSFileAccessIntentClass NSFileAccessIntentClass) WritingIntentWithURLOptions(url INSURL, options NSFileCoordinatorWritingOptions) NSFileAccessIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSFileAccessIntentClass.class), objc.Sel("writingIntentWithURL:options:"), url, options)
	return NSFileAccessIntentFromID(rv)
}

// The current URL for the item managed by the file access intent instance.
// (read-only)
//
// # Discussion
// 
// Always use the URL returned by this property inside the accessor block of a
// file coordinator’s [CoordinateAccessWithIntentsQueueByAccessor] method.
// This property’s value may be different from the original URL, because the
// item was either moved or renamed while the file coordinator waited for
// access.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/url
func (f NSFileAccessIntent) URL() INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}

