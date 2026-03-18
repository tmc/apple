// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSNameSpecifier] class.
var (
	_NSNameSpecifierClass     NSNameSpecifierClass
	_NSNameSpecifierClassOnce sync.Once
)

func getNSNameSpecifierClass() NSNameSpecifierClass {
	_NSNameSpecifierClassOnce.Do(func() {
		_NSNameSpecifierClass = NSNameSpecifierClass{class: objc.GetClass("NSNameSpecifier")}
	})
	return _NSNameSpecifierClass
}

// GetNSNameSpecifierClass returns the class object for NSNameSpecifier.
func GetNSNameSpecifierClass() NSNameSpecifierClass {
	return getNSNameSpecifierClass()
}

type NSNameSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNameSpecifierClass) Alloc() NSNameSpecifier {
	rv := objc.Send[NSNameSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier for an object in a collection (or container) by name.
//
// # Overview
// 
// As an example, the following script specifies both an application and a
// window by name. In this script, the named window’s implicitly specified
// container is the Finder application’s list of open windows.
// 
// This specifier works only for objects that have a name property. You
// don’t normally subclass [NSNameSpecifier].
// 
// The evaluation of an instance of [NSNameSpecifier] follows these steps
// until the specified object is found:
// 
// - If the container implements a method whose selector matches the relevant
// `valueInWithName:` pattern established by scripting key-value coding, the
// method is invoked. This method can potentially be very fast, and it may be
// relatively easy to implement. - As is the case when evaluating any script
// object specifier, the container of the specified object is given a chance
// to evaluate the object specifier. If the container class implements the
// `indicesOfObjectsByEvaluatingObjectSpecifier` method, the method is
// invoked. This method can potentially be very fast, but it is relatively
// difficult to implement. - An instance of [NSWhoseSpecifier] that specifies
// the first object whose relevant `'pnam'` attribute matches the name is
// synthesized and evaluated. The instance of [NSWhoseSpecifier] must search
// through all of the keyed elements in the container, looking for a match.
// The search is potentially very slow.
//
// # Initializing a name specifier
//
//   - [NSNameSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyName]: Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and then sets the name instance variable to `name`.
//
// # Accessing a name specifier
//
//   - [NSNameSpecifier.Name]: Sets the name encapsulated with the receiver for the specified object in the container.
//   - [NSNameSpecifier.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier
type NSNameSpecifier struct {
	NSScriptObjectSpecifier
}

// NSNameSpecifierFromID constructs a [NSNameSpecifier] from an objc.ID.
//
// A specifier for an object in a collection (or container) by name.
func NSNameSpecifierFromID(id objc.ID) NSNameSpecifier {
	return NSNameSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSNameSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSNameSpecifier] class.
//
// # Initializing a name specifier
//
//   - [INSNameSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyName]: Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and then sets the name instance variable to `name`.
//
// # Accessing a name specifier
//
//   - [INSNameSpecifier.Name]: Sets the name encapsulated with the receiver for the specified object in the container.
//   - [INSNameSpecifier.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier
type INSNameSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Initializing a name specifier

	// Invokes the super class’s [init(containerClassDescription:containerSpecifier:key:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)>) method and then sets the name instance variable to `name`.
	InitWithContainerClassDescriptionContainerSpecifierKeyName(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, name string) NSNameSpecifier

	// Topic: Accessing a name specifier

	// Sets the name encapsulated with the receiver for the specified object in the container.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (n NSNameSpecifier) Init() NSNameSpecifier {
	rv := objc.Send[NSNameSpecifier](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNameSpecifier) Autorelease() NSNameSpecifier {
	rv := objc.Send[NSNameSpecifier](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNameSpecifier creates a new NSNameSpecifier instance.
func NewNSNameSpecifier() NSNameSpecifier {
	class := getNSNameSpecifierClass()
	rv := objc.Send[NSNameSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/init(coder:)
func NewNameSpecifierWithCoder(inCoder INSCoder) NSNameSpecifier {
	instance := getNSNameSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSNameSpecifierFromID(rv)
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
func NewNameSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSNameSpecifier {
	instance := getNSNameSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSNameSpecifierFromID(rv)
}

// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and then
// sets the name instance variable to `name`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/init(containerClassDescription:containerSpecifier:key:name:)
func NewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, name string) NSNameSpecifier {
	instance := getNSNameSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:name:"), classDesc, container, objc.String(property), objc.String(name))
	return NSNameSpecifierFromID(rv)
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
func NewNameSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSNameSpecifier {
	instance := getNSNameSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSNameSpecifierFromID(rv)
}

// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and then
// sets the name instance variable to `name`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/init(containerClassDescription:containerSpecifier:key:name:)
func (n NSNameSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyName(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, name string) NSNameSpecifier {
	rv := objc.Send[NSNameSpecifier](n.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:name:"), classDesc, container, objc.String(property), objc.String(name))
	return rv
}

// Sets the name encapsulated with the receiver for the specified object in
// the container.
//
// See: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/name
func (n NSNameSpecifier) Name() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (n NSNameSpecifier) SetName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setName:"), objc.String(value))
}

