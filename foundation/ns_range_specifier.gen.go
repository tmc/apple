// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSRangeSpecifier] class.
var (
	_NSRangeSpecifierClass     NSRangeSpecifierClass
	_NSRangeSpecifierClassOnce sync.Once
)

func getNSRangeSpecifierClass() NSRangeSpecifierClass {
	_NSRangeSpecifierClassOnce.Do(func() {
		_NSRangeSpecifierClass = NSRangeSpecifierClass{class: objc.GetClass("NSRangeSpecifier")}
	})
	return _NSRangeSpecifierClass
}

// GetNSRangeSpecifierClass returns the class object for NSRangeSpecifier.
func GetNSRangeSpecifierClass() NSRangeSpecifierClass {
	return getNSRangeSpecifierClass()
}

type NSRangeSpecifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSRangeSpecifierClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRangeSpecifierClass) Alloc() NSRangeSpecifier {
	rv := objc.Send[NSRangeSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier for a range of objects in a container.
//
// # Overview
// 
// An [NSRangeSpecifier] object specifies a range (that is, an uninterrupted
// series) of objects in a container through two delimiting objects. The range
// is represented by two object specifiers, a start specifier and an end
// specifier, which can be of any specifier type (such as [NSIndexSpecifier]
// or [NSWhoseSpecifier] object). These specifiers are evaluated in the
// context of the same container object as the range specifier itself.
// 
// You don’t normally subclass [NSRangeSpecifier].
//
// # Initializing a range specifier
//
//   - [NSRangeSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier]: Returns a range specifier initialized with the given properties.
//
// # Accessing a range specifier
//
//   - [NSRangeSpecifier.EndSpecifier]: Sets the object specifier representing the last object of the range to a given object.
//   - [NSRangeSpecifier.SetEndSpecifier]
//   - [NSRangeSpecifier.StartSpecifier]: Returns the object specifier representing the first object of the range.
//   - [NSRangeSpecifier.SetStartSpecifier]
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier
type NSRangeSpecifier struct {
	NSScriptObjectSpecifier
}

// NSRangeSpecifierFromID constructs a [NSRangeSpecifier] from an objc.ID.
//
// A specifier for a range of objects in a container.
func NSRangeSpecifierFromID(id objc.ID) NSRangeSpecifier {
	return NSRangeSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSRangeSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRangeSpecifier] class.
//
// # Initializing a range specifier
//
//   - [INSRangeSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier]: Returns a range specifier initialized with the given properties.
//
// # Accessing a range specifier
//
//   - [INSRangeSpecifier.EndSpecifier]: Sets the object specifier representing the last object of the range to a given object.
//   - [INSRangeSpecifier.SetEndSpecifier]
//   - [INSRangeSpecifier.StartSpecifier]: Returns the object specifier representing the first object of the range.
//   - [INSRangeSpecifier.SetStartSpecifier]
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier
type INSRangeSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Initializing a range specifier

	// Returns a range specifier initialized with the given properties.
	InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, startSpec INSScriptObjectSpecifier, endSpec INSScriptObjectSpecifier) NSRangeSpecifier

	// Topic: Accessing a range specifier

	// Sets the object specifier representing the last object of the range to a given object.
	EndSpecifier() INSScriptObjectSpecifier
	SetEndSpecifier(value INSScriptObjectSpecifier)
	// Returns the object specifier representing the first object of the range.
	StartSpecifier() INSScriptObjectSpecifier
	SetStartSpecifier(value INSScriptObjectSpecifier)
}

// Init initializes the instance.
func (r NSRangeSpecifier) Init() NSRangeSpecifier {
	rv := objc.Send[NSRangeSpecifier](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRangeSpecifier) Autorelease() NSRangeSpecifier {
	rv := objc.Send[NSRangeSpecifier](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRangeSpecifier creates a new NSRangeSpecifier instance.
func NewNSRangeSpecifier() NSRangeSpecifier {
	class := getNSRangeSpecifierClass()
	rv := objc.Send[NSRangeSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(coder:)
func NewRangeSpecifierWithCoder(inCoder INSCoder) NSRangeSpecifier {
	instance := getNSRangeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSRangeSpecifierFromID(rv)
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
func NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSRangeSpecifier {
	instance := getNSRangeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSRangeSpecifierFromID(rv)
}

// Returns a range specifier initialized with the given properties.
//
// classDesc: The class description.
//
// container: The container.
//
// property: The property.
//
// startSpec: The object specifier representing the first object of the range.
//
// endSpec: The object specifier representing the last object of the range.
//
// # Return Value
// 
// A range specifier initialized with the given properties.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and
// initializes the instance with the object specifiers representing the
// starting element, `startSpec`, and the ending element, `endSpec`, of a
// range of elements in the container.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(containerClassDescription:containerSpecifier:key:start:end:)
func NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, startSpec INSScriptObjectSpecifier, endSpec INSScriptObjectSpecifier) NSRangeSpecifier {
	instance := getNSRangeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:startSpecifier:endSpecifier:"), classDesc, container, objc.String(property), startSpec, endSpec)
	return NSRangeSpecifierFromID(rv)
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
func NewRangeSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSRangeSpecifier {
	instance := getNSRangeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSRangeSpecifierFromID(rv)
}

// Returns a range specifier initialized with the given properties.
//
// classDesc: The class description.
//
// container: The container.
//
// property: The property.
//
// startSpec: The object specifier representing the first object of the range.
//
// endSpec: The object specifier representing the last object of the range.
//
// # Return Value
// 
// A range specifier initialized with the given properties.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and
// initializes the instance with the object specifiers representing the
// starting element, `startSpec`, and the ending element, `endSpec`, of a
// range of elements in the container.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(containerClassDescription:containerSpecifier:key:start:end:)
func (r NSRangeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, startSpec INSScriptObjectSpecifier, endSpec INSScriptObjectSpecifier) NSRangeSpecifier {
	rv := objc.Send[NSRangeSpecifier](r.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:startSpecifier:endSpecifier:"), classDesc, container, objc.String(property), startSpec, endSpec)
	return rv
}

// Sets the object specifier representing the last object of the range to a
// given object.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/endSpecifier
func (r NSRangeSpecifier) EndSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("endSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (r NSRangeSpecifier) SetEndSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](r.ID, objc.Sel("setEndSpecifier:"), value)
}
// Returns the object specifier representing the first object of the range.
//
// # Return Value
// 
// The object specifier representing the first object of the range.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/startSpecifier
func (r NSRangeSpecifier) StartSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("startSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (r NSRangeSpecifier) SetStartSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](r.ID, objc.Sel("setStartSpecifier:"), value)
}

