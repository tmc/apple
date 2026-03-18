// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSRandomSpecifier] class.
var (
	_NSRandomSpecifierClass     NSRandomSpecifierClass
	_NSRandomSpecifierClassOnce sync.Once
)

func getNSRandomSpecifierClass() NSRandomSpecifierClass {
	_NSRandomSpecifierClassOnce.Do(func() {
		_NSRandomSpecifierClass = NSRandomSpecifierClass{class: objc.GetClass("NSRandomSpecifier")}
	})
	return _NSRandomSpecifierClass
}

// GetNSRandomSpecifierClass returns the class object for NSRandomSpecifier.
func GetNSRandomSpecifierClass() NSRandomSpecifierClass {
	return getNSRandomSpecifierClass()
}

type NSRandomSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRandomSpecifierClass) Alloc() NSRandomSpecifier {
	rv := objc.Send[NSRandomSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier for an arbitrary object in a collection or, if not a
// one-to-many relationship, the sole object.
//
// See: https://developer.apple.com/documentation/Foundation/NSRandomSpecifier
type NSRandomSpecifier struct {
	NSScriptObjectSpecifier
}

// NSRandomSpecifierFromID constructs a [NSRandomSpecifier] from an objc.ID.
//
// A specifier for an arbitrary object in a collection or, if not a
// one-to-many relationship, the sole object.
func NSRandomSpecifierFromID(id objc.ID) NSRandomSpecifier {
	return NSRandomSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSRandomSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRandomSpecifier] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSRandomSpecifier
type INSRandomSpecifier interface {
	INSScriptObjectSpecifier
}

// Init initializes the instance.
func (r NSRandomSpecifier) Init() NSRandomSpecifier {
	rv := objc.Send[NSRandomSpecifier](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRandomSpecifier) Autorelease() NSRandomSpecifier {
	rv := objc.Send[NSRandomSpecifier](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRandomSpecifier creates a new NSRandomSpecifier instance.
func NewNSRandomSpecifier() NSRandomSpecifier {
	class := getNSRandomSpecifierClass()
	rv := objc.Send[NSRandomSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func NewRandomSpecifierWithCoder(inCoder INSCoder) NSRandomSpecifier {
	instance := getNSRandomSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSRandomSpecifierFromID(rv)
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
func NewRandomSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSRandomSpecifier {
	instance := getNSRandomSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSRandomSpecifierFromID(rv)
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
func NewRandomSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSRandomSpecifier {
	instance := getNSRandomSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSRandomSpecifierFromID(rv)
}

