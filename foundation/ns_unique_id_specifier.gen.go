// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUniqueIDSpecifier] class.
var (
	_NSUniqueIDSpecifierClass     NSUniqueIDSpecifierClass
	_NSUniqueIDSpecifierClassOnce sync.Once
)

func getNSUniqueIDSpecifierClass() NSUniqueIDSpecifierClass {
	_NSUniqueIDSpecifierClassOnce.Do(func() {
		_NSUniqueIDSpecifierClass = NSUniqueIDSpecifierClass{class: objc.GetClass("NSUniqueIDSpecifier")}
	})
	return _NSUniqueIDSpecifierClass
}

// GetNSUniqueIDSpecifierClass returns the class object for NSUniqueIDSpecifier.
func GetNSUniqueIDSpecifierClass() NSUniqueIDSpecifierClass {
	return getNSUniqueIDSpecifierClass()
}

type NSUniqueIDSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUniqueIDSpecifierClass) Alloc() NSUniqueIDSpecifier {
	rv := objc.Send[NSUniqueIDSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A specifier for an object in a collection (or container) by unique ID.
//
// # Overview
// 
// This specifier works only for objects that have an ID property. The unique
// ID object passed to an instance of [NSUniqueIDSpecifier] must be either an
// [NSNumber] object or an [NSString] object. The exact type should match the
// scripting dictionary declaration of the ID attribute for the relevant
// scripting class.
// 
// You can expect that the ID property will be for any object that supports
// it. Therefore a scripter can obtain the unique ID for an object and refer
// to the object by the ID, but cannot set the unique ID.
// 
// You don’t normally subclass [NSUniqueIDSpecifier].
// 
// The evaluation of [NSUniqueIDSpecifier] objects follows these steps until
// the specified object is found:
// 
// - If the container implements a method whose selector matches the relevant
// `valueInWithUniqueID:` pattern established by scripting key-value coding,
// the method is invoked. This method can potentially be very fast, and it may
// be relatively easy to implement. - As is the case when evaluating any
// script object specifier, the container of the specified object is given a
// chance to evaluate the object specifier. If the container class implements
// the [indicesOfObjects(byEvaluatingObjectSpecifier:)] method, the method is
// invoked. This method can potentially be very fast, but it is relatively
// difficult to implement. - An [NSWhoseSpecifier] object that specifies the
// first object whose relevant `'ID '` attribute matches the ID is synthesized
// and evaluated. The [NSWhoseSpecifier] object must search through all of the
// keyed elements in the container, looking for a match. The search is
// potentially very slow.
//
// [indicesOfObjects(byEvaluatingObjectSpecifier:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/indicesOfObjects(byEvaluatingObjectSpecifier:)
//
// # Initializing a unique ID specifier
//
//   - [NSUniqueIDSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID]: Returns an [NSUniqueIDSpecifier] object, initialized with the given arguments.
//
// # Accessing unique ID information
//
//   - [NSUniqueIDSpecifier.UniqueID]: Returns the ID encapsulated by the receiver.
//   - [NSUniqueIDSpecifier.SetUniqueID]
//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier
type NSUniqueIDSpecifier struct {
	NSScriptObjectSpecifier
}

// NSUniqueIDSpecifierFromID constructs a [NSUniqueIDSpecifier] from an objc.ID.
//
// A specifier for an object in a collection (or container) by unique ID.
func NSUniqueIDSpecifierFromID(id objc.ID) NSUniqueIDSpecifier {
	return NSUniqueIDSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSUniqueIDSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSUniqueIDSpecifier] class.
//
// # Initializing a unique ID specifier
//
//   - [INSUniqueIDSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID]: Returns an [NSUniqueIDSpecifier] object, initialized with the given arguments.
//
// # Accessing unique ID information
//
//   - [INSUniqueIDSpecifier.UniqueID]: Returns the ID encapsulated by the receiver.
//   - [INSUniqueIDSpecifier.SetUniqueID]
//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier
type INSUniqueIDSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Initializing a unique ID specifier

	// Returns an [NSUniqueIDSpecifier] object, initialized with the given arguments.
	InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, uniqueID objectivec.IObject) NSUniqueIDSpecifier

	// Topic: Accessing unique ID information

	// Returns the ID encapsulated by the receiver.
	UniqueID() objectivec.IObject
	SetUniqueID(value objectivec.IObject)
}





// Init initializes the instance.
func (u NSUniqueIDSpecifier) Init() NSUniqueIDSpecifier {
	rv := objc.Send[NSUniqueIDSpecifier](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUniqueIDSpecifier) Autorelease() NSUniqueIDSpecifier {
	rv := objc.Send[NSUniqueIDSpecifier](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUniqueIDSpecifier creates a new NSUniqueIDSpecifier instance.
func NewNSUniqueIDSpecifier() NSUniqueIDSpecifier {
	class := getNSUniqueIDSpecifierClass()
	rv := objc.Send[NSUniqueIDSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier/init(coder:)
func NewUniqueIDSpecifierWithCoder(inCoder INSCoder) NSUniqueIDSpecifier {
	instance := getNSUniqueIDSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSUniqueIDSpecifierFromID(rv)
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
func NewUniqueIDSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSUniqueIDSpecifier {
	instance := getNSUniqueIDSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSUniqueIDSpecifierFromID(rv)
}


// Returns an [NSUniqueIDSpecifier] object, initialized with the given
// arguments.
//
// classDesc: The class description for the new object.
//
// container: The container for the new object.
//
// property: The property for the new object.
//
// uniqueID: The unique ID for the new object.
// 
// `uniqueID` must be an instance of [NSNumber] or [NSString]. The type should
// match the declared type of the attribute of the specified scriptable class
// whose four-character code is `'ID '`.
//
// # Return Value
// 
// An [NSUniqueIDSpecifier] object, initialized with the given arguments.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and sets
// the ID to `uniqueID`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier/init(containerClassDescription:containerSpecifier:key:uniqueID:)
func NewUniqueIDSpecifierWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, uniqueID objectivec.IObject) NSUniqueIDSpecifier {
	instance := getNSUniqueIDSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:uniqueID:"), classDesc, container, objc.String(property), uniqueID)
	return NSUniqueIDSpecifierFromID(rv)
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
func NewUniqueIDSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSUniqueIDSpecifier {
	instance := getNSUniqueIDSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSUniqueIDSpecifierFromID(rv)
}







// Returns an [NSUniqueIDSpecifier] object, initialized with the given
// arguments.
//
// classDesc: The class description for the new object.
//
// container: The container for the new object.
//
// property: The property for the new object.
//
// uniqueID: The unique ID for the new object.
// 
// `uniqueID` must be an instance of [NSNumber] or [NSString]. The type should
// match the declared type of the attribute of the specified scriptable class
// whose four-character code is `'ID '`.
//
// # Return Value
// 
// An [NSUniqueIDSpecifier] object, initialized with the given arguments.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and sets
// the ID to `uniqueID`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier/init(containerClassDescription:containerSpecifier:key:uniqueID:)
func (u NSUniqueIDSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyUniqueID(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, uniqueID objectivec.IObject) NSUniqueIDSpecifier {
	rv := objc.Send[NSUniqueIDSpecifier](u.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:uniqueID:"), classDesc, container, objc.String(property), uniqueID)
	return rv
}












// Returns the ID encapsulated by the receiver.
//
// # Return Value
// 
// The ID encapsulated by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier/uniqueID
func (u NSUniqueIDSpecifier) UniqueID() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uniqueID"))
	return objectivec.Object{ID: rv}
}
func (u NSUniqueIDSpecifier) SetUniqueID(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setUniqueID:"), value)
}


























