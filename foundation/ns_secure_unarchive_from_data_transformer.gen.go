// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSecureUnarchiveFromDataTransformer] class.
var (
	_NSSecureUnarchiveFromDataTransformerClass     NSSecureUnarchiveFromDataTransformerClass
	_NSSecureUnarchiveFromDataTransformerClassOnce sync.Once
)

func getNSSecureUnarchiveFromDataTransformerClass() NSSecureUnarchiveFromDataTransformerClass {
	_NSSecureUnarchiveFromDataTransformerClassOnce.Do(func() {
		_NSSecureUnarchiveFromDataTransformerClass = NSSecureUnarchiveFromDataTransformerClass{class: objc.GetClass("NSSecureUnarchiveFromDataTransformer")}
	})
	return _NSSecureUnarchiveFromDataTransformerClass
}

// GetNSSecureUnarchiveFromDataTransformerClass returns the class object for NSSecureUnarchiveFromDataTransformer.
func GetNSSecureUnarchiveFromDataTransformerClass() NSSecureUnarchiveFromDataTransformerClass {
	return getNSSecureUnarchiveFromDataTransformerClass()
}

type NSSecureUnarchiveFromDataTransformerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSecureUnarchiveFromDataTransformerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSecureUnarchiveFromDataTransformerClass) Alloc() NSSecureUnarchiveFromDataTransformer {
	rv := objc.Send[NSSecureUnarchiveFromDataTransformer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A value transformer that converts data to and from classes that support
// secure coding.
//
// # Overview
//
// This class provides a default [NSValueTransformer] implementation for
// secure decoding. This class attempts to decode data into the classes listed
// within [NSSecureUnarchiveFromDataTransformer.AllowedTopLevelClasses], which includes [NSArray], [NSDictionary],
// [NSSet], [NSString], [NSNumber], [NSDate], [NSData], [NSURL], [NSUUID], and
// [NSNull].
//
// To archive or unarchive other classes that support [NSSecureCoding], create
// a subclass and override [NSSecureUnarchiveFromDataTransformer.AllowedTopLevelClasses] to list the classes to
// transform.
//
// To use [NSSecureUnarchiveFromDataTransformer] with [Core Data], use the
// name of this class, or the name of a subclass you implement, as the name of
// the transformer for an entity’s attribute within a Core Data Model. If
// you use your own transformer subclass, register it with your app before
// intializing your persistent container with Core Data.
//
// For an example of subclassing [NSSecureUnarchiveFromDataTransformer], see
// [Handling Different Data Types in Core Data], which has a
// [ColorToDataTransformer] class that transforms [UIColor] to [NSData] and
// the reverse, to support archiving instances of [UIColor].
//
// See: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer
//
// [Core Data]: https://developer.apple.com/documentation/CoreData
// [Handling Different Data Types in Core Data]: https://developer.apple.com/documentation/CoreData/handling-different-data-types-in-core-data
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
type NSSecureUnarchiveFromDataTransformer struct {
	NSValueTransformer
}

// NSSecureUnarchiveFromDataTransformerFromID constructs a [NSSecureUnarchiveFromDataTransformer] from an objc.ID.
//
// A value transformer that converts data to and from classes that support
// secure coding.
func NSSecureUnarchiveFromDataTransformerFromID(id objc.ID) NSSecureUnarchiveFromDataTransformer {
	return NSSecureUnarchiveFromDataTransformer{NSValueTransformer: NSValueTransformerFromID(id)}
}

// NOTE: NSSecureUnarchiveFromDataTransformer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSecureUnarchiveFromDataTransformer] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer
type INSSecureUnarchiveFromDataTransformer interface {
	INSValueTransformer
}

// Init initializes the instance.
func (s NSSecureUnarchiveFromDataTransformer) Init() NSSecureUnarchiveFromDataTransformer {
	rv := objc.Send[NSSecureUnarchiveFromDataTransformer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSecureUnarchiveFromDataTransformer) Autorelease() NSSecureUnarchiveFromDataTransformer {
	rv := objc.Send[NSSecureUnarchiveFromDataTransformer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSecureUnarchiveFromDataTransformer creates a new NSSecureUnarchiveFromDataTransformer instance.
func NewNSSecureUnarchiveFromDataTransformer() NSSecureUnarchiveFromDataTransformer {
	class := getNSSecureUnarchiveFromDataTransformerClass()
	rv := objc.Send[NSSecureUnarchiveFromDataTransformer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A list of allowed classes the top-level object in an archive must conform
// to, for encoding and decoding.
//
// # Discussion
//
// This property contains the value of [TransformedValueClass] if that value
// isn’t `nil`. Otherwise, it holds a list of the top level classes that it
// decodes, which includes [NSArray], [NSDictionary], [NSSet], [NSString],
// [NSNumber], [NSDate], [NSData], [NSURL], [NSUUID], and [NSNull].
//
// Override this property in subclasses to provide an expanded or different
// set of allowed transformation classes.
//
// See: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer/allowedTopLevelClasses
func (_NSSecureUnarchiveFromDataTransformerClass NSSecureUnarchiveFromDataTransformerClass) AllowedTopLevelClasses() []objc.Class {
	rv := objc.Send[[]objc.ID](objc.ID(_NSSecureUnarchiveFromDataTransformerClass.class), objc.Sel("allowedTopLevelClasses"))
	return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
		return objc.Class(id)
	})
}
