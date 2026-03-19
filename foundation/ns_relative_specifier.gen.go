// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSRelativeSpecifier] class.
var (
	_NSRelativeSpecifierClass     NSRelativeSpecifierClass
	_NSRelativeSpecifierClassOnce sync.Once
)

func getNSRelativeSpecifierClass() NSRelativeSpecifierClass {
	_NSRelativeSpecifierClassOnce.Do(func() {
		_NSRelativeSpecifierClass = NSRelativeSpecifierClass{class: objc.GetClass("NSRelativeSpecifier")}
	})
	return _NSRelativeSpecifierClass
}

// GetNSRelativeSpecifierClass returns the class object for NSRelativeSpecifier.
func GetNSRelativeSpecifierClass() NSRelativeSpecifierClass {
	return getNSRelativeSpecifierClass()
}

type NSRelativeSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRelativeSpecifierClass) Alloc() NSRelativeSpecifier {
	rv := objc.Send[NSRelativeSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier that indicates an object in a collection by its position
// relative to another object.
//
// # Overview
// 
// You don’t normally subclass [NSRelativeSpecifier].
//
// # Initializing a relative specifier
//
//   - [NSRelativeSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier]: Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and initializes the relative position and base specifier to `relPos` and `baseSpecifier`.
//
// # Accessing a relative specifier
//
//   - [NSRelativeSpecifier.BaseSpecifier]: Sets the specifier for the base object.
//   - [NSRelativeSpecifier.SetBaseSpecifier]
//   - [NSRelativeSpecifier.RelativePosition]: Sets the relative position encapsulated by the receiver.
//   - [NSRelativeSpecifier.SetRelativePosition]
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier
type NSRelativeSpecifier struct {
	NSScriptObjectSpecifier
}

// NSRelativeSpecifierFromID constructs a [NSRelativeSpecifier] from an objc.ID.
//
// A specifier that indicates an object in a collection by its position
// relative to another object.
func NSRelativeSpecifierFromID(id objc.ID) NSRelativeSpecifier {
	return NSRelativeSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSRelativeSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRelativeSpecifier] class.
//
// # Initializing a relative specifier
//
//   - [INSRelativeSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier]: Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and initializes the relative position and base specifier to `relPos` and `baseSpecifier`.
//
// # Accessing a relative specifier
//
//   - [INSRelativeSpecifier.BaseSpecifier]: Sets the specifier for the base object.
//   - [INSRelativeSpecifier.SetBaseSpecifier]
//   - [INSRelativeSpecifier.RelativePosition]: Sets the relative position encapsulated by the receiver.
//   - [INSRelativeSpecifier.SetRelativePosition]
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier
type INSRelativeSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Initializing a relative specifier

	// Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and initializes the relative position and base specifier to `relPos` and `baseSpecifier`.
	InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, relPos NSRelativePosition, baseSpecifier INSScriptObjectSpecifier) NSRelativeSpecifier

	// Topic: Accessing a relative specifier

	// Sets the specifier for the base object.
	BaseSpecifier() INSScriptObjectSpecifier
	SetBaseSpecifier(value INSScriptObjectSpecifier)
	// Sets the relative position encapsulated by the receiver.
	RelativePosition() NSRelativePosition
	SetRelativePosition(value NSRelativePosition)
}

// Init initializes the instance.
func (r NSRelativeSpecifier) Init() NSRelativeSpecifier {
	rv := objc.Send[NSRelativeSpecifier](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRelativeSpecifier) Autorelease() NSRelativeSpecifier {
	rv := objc.Send[NSRelativeSpecifier](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRelativeSpecifier creates a new NSRelativeSpecifier instance.
func NewNSRelativeSpecifier() NSRelativeSpecifier {
	class := getNSRelativeSpecifierClass()
	rv := objc.Send[NSRelativeSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/init(coder:)
func NewRelativeSpecifierWithCoder(inCoder INSCoder) NSRelativeSpecifier {
	instance := getNSRelativeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSRelativeSpecifierFromID(rv)
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
func NewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSRelativeSpecifier {
	instance := getNSRelativeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSRelativeSpecifierFromID(rv)
}

// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and
// initializes the relative position and base specifier to `relPos` and
// `baseSpecifier`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/init(containerClassDescription:containerSpecifier:key:relativePosition:baseSpecifier:)
func NewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, relPos NSRelativePosition, baseSpecifier INSScriptObjectSpecifier) NSRelativeSpecifier {
	instance := getNSRelativeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:relativePosition:baseSpecifier:"), classDesc, container, objc.String(property), relPos, baseSpecifier)
	return NSRelativeSpecifierFromID(rv)
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
func NewRelativeSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSRelativeSpecifier {
	instance := getNSRelativeSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSRelativeSpecifierFromID(rv)
}

// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and
// initializes the relative position and base specifier to `relPos` and
// `baseSpecifier`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/init(containerClassDescription:containerSpecifier:key:relativePosition:baseSpecifier:)
func (r NSRelativeSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, relPos NSRelativePosition, baseSpecifier INSScriptObjectSpecifier) NSRelativeSpecifier {
	rv := objc.Send[NSRelativeSpecifier](r.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:relativePosition:baseSpecifier:"), classDesc, container, objc.String(property), relPos, baseSpecifier)
	return rv
}

// Sets the specifier for the base object.
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/baseSpecifier
func (r NSRelativeSpecifier) BaseSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("baseSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (r NSRelativeSpecifier) SetBaseSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](r.ID, objc.Sel("setBaseSpecifier:"), value)
}
// Sets the relative position encapsulated by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/relativePosition-swift.property
func (r NSRelativeSpecifier) RelativePosition() NSRelativePosition {
	rv := objc.Send[NSRelativePosition](r.ID, objc.Sel("relativePosition"))
	return NSRelativePosition(rv)
}
func (r NSRelativeSpecifier) SetRelativePosition(value NSRelativePosition) {
	objc.Send[struct{}](r.ID, objc.Sel("setRelativePosition:"), value)
}

