// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPropertySpecifier] class.
var (
	_NSPropertySpecifierClass     NSPropertySpecifierClass
	_NSPropertySpecifierClassOnce sync.Once
)

func getNSPropertySpecifierClass() NSPropertySpecifierClass {
	_NSPropertySpecifierClassOnce.Do(func() {
		_NSPropertySpecifierClass = NSPropertySpecifierClass{class: objc.GetClass("NSPropertySpecifier")}
	})
	return _NSPropertySpecifierClass
}

// GetNSPropertySpecifierClass returns the class object for NSPropertySpecifier.
func GetNSPropertySpecifierClass() NSPropertySpecifierClass {
	return getNSPropertySpecifierClass()
}

type NSPropertySpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPropertySpecifierClass) Alloc() NSPropertySpecifier {
	rv := objc.Send[NSPropertySpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier for a simple attribute value, a one-to-one relationship, or all
// elements of a to-many relationship.
//
// # Overview
// 
// You don’t typically subclass [NSPropertySpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSPropertySpecifier
type NSPropertySpecifier struct {
	NSScriptObjectSpecifier
}

// NSPropertySpecifierFromID constructs a [NSPropertySpecifier] from an objc.ID.
//
// A specifier for a simple attribute value, a one-to-one relationship, or all
// elements of a to-many relationship.
func NSPropertySpecifierFromID(id objc.ID) NSPropertySpecifier {
	return NSPropertySpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSPropertySpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPropertySpecifier] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPropertySpecifier
type INSPropertySpecifier interface {
	INSScriptObjectSpecifier
}

// Init initializes the instance.
func (p NSPropertySpecifier) Init() NSPropertySpecifier {
	rv := objc.Send[NSPropertySpecifier](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPropertySpecifier) Autorelease() NSPropertySpecifier {
	rv := objc.Send[NSPropertySpecifier](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPropertySpecifier creates a new NSPropertySpecifier instance.
func NewNSPropertySpecifier() NSPropertySpecifier {
	class := getNSPropertySpecifierClass()
	rv := objc.Send[NSPropertySpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func NewPropertySpecifierWithCoder(inCoder INSCoder) NSPropertySpecifier {
	instance := getNSPropertySpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSPropertySpecifierFromID(rv)
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
func NewPropertySpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSPropertySpecifier {
	instance := getNSPropertySpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSPropertySpecifierFromID(rv)
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
func NewPropertySpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSPropertySpecifier {
	instance := getNSPropertySpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSPropertySpecifierFromID(rv)
}

