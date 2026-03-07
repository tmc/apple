// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMiddleSpecifier] class.
var (
	_NSMiddleSpecifierClass     NSMiddleSpecifierClass
	_NSMiddleSpecifierClassOnce sync.Once
)

func getNSMiddleSpecifierClass() NSMiddleSpecifierClass {
	_NSMiddleSpecifierClassOnce.Do(func() {
		_NSMiddleSpecifierClass = NSMiddleSpecifierClass{class: objc.GetClass("NSMiddleSpecifier")}
	})
	return _NSMiddleSpecifierClass
}

// GetNSMiddleSpecifierClass returns the class object for NSMiddleSpecifier.
func GetNSMiddleSpecifierClass() NSMiddleSpecifierClass {
	return getNSMiddleSpecifierClass()
}

type NSMiddleSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMiddleSpecifierClass) Alloc() NSMiddleSpecifier {
	rv := objc.Send[NSMiddleSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A specifier indicating the middle object in a collection or, if not a
// one-to-many relationship, the sole object.
//
// # Overview
// 
// You don’t typically subclass [NSMiddleSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSMiddleSpecifier
type NSMiddleSpecifier struct {
	NSScriptObjectSpecifier
}

// NSMiddleSpecifierFromID constructs a [NSMiddleSpecifier] from an objc.ID.
//
// A specifier indicating the middle object in a collection or, if not a
// one-to-many relationship, the sole object.
func NSMiddleSpecifierFromID(id objc.ID) NSMiddleSpecifier {
	return NSMiddleSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSMiddleSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMiddleSpecifier] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSMiddleSpecifier
type INSMiddleSpecifier interface {
	INSScriptObjectSpecifier
}





// Init initializes the instance.
func (m NSMiddleSpecifier) Init() NSMiddleSpecifier {
	rv := objc.Send[NSMiddleSpecifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMiddleSpecifier) Autorelease() NSMiddleSpecifier {
	rv := objc.Send[NSMiddleSpecifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMiddleSpecifier creates a new NSMiddleSpecifier instance.
func NewNSMiddleSpecifier() NSMiddleSpecifier {
	class := getNSMiddleSpecifierClass()
	rv := objc.Send[NSMiddleSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func NewMiddleSpecifierWithCoder(inCoder INSCoder) NSMiddleSpecifier {
	instance := getNSMiddleSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSMiddleSpecifierFromID(rv)
}


// Returns an [NSScriptObjectSpecifier] object initialized with the given
// attributes.
//
// # Return Value
// 
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier`, key `key`, and the class description of the object specifier
// `classDescription`, derived from the value of the specifier’s key.
//
// # Discussion
// 
// You should never pass `nil` for the value of `classDescription`. The
// receiver’s child reference is set to `nil`.
// 
// This is the designated initializer for [NSScriptObjectSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)
func NewMiddleSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSMiddleSpecifier {
	instance := getNSMiddleSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSMiddleSpecifierFromID(rv)
}


// Returns an [NSScriptObjectSpecifier] object initialized with a given
// container specifier and key.
//
// # Return Value
// 
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier` and key `key`.
//
// # Discussion
// 
// The class description of the container is set automatically.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerSpecifier:key:)
func NewMiddleSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSMiddleSpecifier {
	instance := getNSMiddleSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSMiddleSpecifierFromID(rv)
}









































