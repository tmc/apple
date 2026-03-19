// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSNull] class.
var (
	_NSNullClass     NSNullClass
	_NSNullClassOnce sync.Once
)

func getNSNullClass() NSNullClass {
	_NSNullClassOnce.Do(func() {
		_NSNullClass = NSNullClass{class: objc.GetClass("NSNull")}
	})
	return _NSNullClass
}

// GetNSNullClass returns the class object for NSNull.
func GetNSNullClass() NSNullClass {
	return getNSNullClass()
}

type NSNullClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNullClass) Alloc() NSNull {
	rv := objc.Send[NSNull](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A singleton object used to represent null values in collection objects that
// don’t allow `nil` values.
//
// # Overview
// 
// [NSNull] is “toll-free bridged” with its Core Foundation counterpart,
// [CFNull]. See [Toll-Free Bridging] for more information on toll-free
// bridging.
//
// [CFNull]: https://developer.apple.com/documentation/CoreFoundation/CFNull
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// See: https://developer.apple.com/documentation/Foundation/NSNull
type NSNull struct {
	objectivec.Object
}

// NSNullFromID constructs a [NSNull] from an objc.ID.
//
// A singleton object used to represent null values in collection objects that
// don’t allow `nil` values.
func NSNullFromID(id objc.ID) NSNull {
	return NSNull{objectivec.Object{ID: id}}
}
// NOTE: NSNull adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSNull] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSNull
type INSNull interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// A value indicating that a requested item couldn’t be found or doesn’t exist.
	NSNotFound() int
	SetNSNotFound(value int)
}

// Init initializes the instance.
func (n NSNull) Init() NSNull {
	rv := objc.Send[NSNull](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNull) Autorelease() NSNull {
	rv := objc.Send[NSNull](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNull creates a new NSNull instance.
func NewNSNull() NSNull {
	class := getNSNullClass()
	rv := objc.Send[NSNull](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewNullWithCoder(coder INSCoder) NSNull {
	instance := getNSNullClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSNullFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (n NSNull) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (n NSNull) InitWithCoder(coder INSCoder) NSNull {
	rv := objc.Send[NSNull](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns the singleton instance of [NSNull].
//
// # Return Value
// 
// The singleton instance of [NSNull]. Equal to [kCFNull].
//
// [kCFNull]: https://developer.apple.com/documentation/CoreFoundation/kCFNull
//
// See: https://developer.apple.com/documentation/Foundation/NSNull/null
func (_NSNullClass NSNullClass) Null() NSNull {
	rv := objc.Send[objc.ID](objc.ID(_NSNullClass.class), objc.Sel("null"))
	return NSNullFromID(rv)
}

// A value indicating that a requested item couldn’t be found or doesn’t
// exist.
//
// See: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (n NSNull) NSNotFound() int {
	rv := objc.Send[int](n.ID, objc.Sel("NSNotFound"))
	return rv
}
func (n NSNull) SetNSNotFound(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setNSNotFound:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

