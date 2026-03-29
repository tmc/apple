// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSWhoseSpecifier] class.
var (
	_NSWhoseSpecifierClass     NSWhoseSpecifierClass
	_NSWhoseSpecifierClassOnce sync.Once
)

func getNSWhoseSpecifierClass() NSWhoseSpecifierClass {
	_NSWhoseSpecifierClassOnce.Do(func() {
		_NSWhoseSpecifierClass = NSWhoseSpecifierClass{class: objc.GetClass("NSWhoseSpecifier")}
	})
	return _NSWhoseSpecifierClass
}

// GetNSWhoseSpecifierClass returns the class object for NSWhoseSpecifier.
func GetNSWhoseSpecifierClass() NSWhoseSpecifierClass {
	return getNSWhoseSpecifierClass()
}

type NSWhoseSpecifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWhoseSpecifierClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWhoseSpecifierClass) Alloc() NSWhoseSpecifier {
	rv := objc.Send[NSWhoseSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier that indicates every object in a collection matching a
// condition.
//
// # Overview
// 
// [NSWhoseSpecifier] specifies every object in a collection (or every element
// in a container) that matches the condition defined by a single Boolean
// expression or multiple Boolean expressions connected by logical operators.
// [NSWhoseSpecifier] is unique among object specifiers in that its top-level
// container is typically not the application object but an evaluated object
// specifier involved in the tested-for condition. An [NSWhoseSpecifier]
// object encapsulates a “test” object for defining this condition. A test
// object is instantiated from a subclass of the abstract [NSScriptWhoseTest]
// class, whose one declared method is [IsTrue]. See “Boolean Expressions
// and Logical Operations” in [NSScriptObjectSpecifier] and the descriptions
// in NSComparisonMethods and NSScriptingComparisonMethods for more
// information.
// 
// The set of elements specified by an [NSWhoseSpecifier] object can be a
// subset of those that pass the [NSWhoseSpecifier] object’s test. This
// subset is specified by the various sub-element properties of the
// [NSWhoseSpecifier] object . Consider as an example the specifier
// `paragraphs where color of third word is blue`. This would be represented
// by an [NSWhoseSpecifier] object that uses a test specifier and another
// object specifier to identify a subset of the objects with the specified
// property. That is, the specifier’s property is `paragraphs`; the test
// specifier is an index specifier with property `words` and `index 3`; and
// the qualifier is a key value qualifier for key `color` and value `[NSColor
// blueColor]`. The test object specifier (`word at index 3`) is evaluated for
// each object (paragraph) using that object as the container; the resulting
// objects (if any) are tested with the qualifier (`color blue`).
// 
// [NSWhoseSpecifier] is part of Cocoa’s built-in script handling. You
// don’t normally subclass it.
//
// # Initializing a whose specifier
//
//   - [NSWhoseSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyTest]: Returns an [NSWhoseSpecifier] object initialized with the given attributes.
//
// # Accessing information about a whose specifier
//
//   - [NSWhoseSpecifier.EndSubelementIdentifier]: Sets the end sub-element identifier for the specifier to the value of a given sub-element.
//   - [NSWhoseSpecifier.SetEndSubelementIdentifier]
//   - [NSWhoseSpecifier.EndSubelementIndex]: Sets the index position of the last sub-element within the range of objects being tested that pass the specifier’s test.
//   - [NSWhoseSpecifier.SetEndSubelementIndex]
//   - [NSWhoseSpecifier.StartSubelementIdentifier]: Returns the start sub-element identifier for the receiver.
//   - [NSWhoseSpecifier.SetStartSubelementIdentifier]
//   - [NSWhoseSpecifier.StartSubelementIndex]: Returns the index position of the first sub-element within the range of objects being tested that pass the receiver’s test.
//   - [NSWhoseSpecifier.SetStartSubelementIndex]
//   - [NSWhoseSpecifier.Test]: Returns the test object encapsulated by the receiver.
//   - [NSWhoseSpecifier.SetTest]
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier
type NSWhoseSpecifier struct {
	NSScriptObjectSpecifier
}

// NSWhoseSpecifierFromID constructs a [NSWhoseSpecifier] from an objc.ID.
//
// A specifier that indicates every object in a collection matching a
// condition.
func NSWhoseSpecifierFromID(id objc.ID) NSWhoseSpecifier {
	return NSWhoseSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSWhoseSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWhoseSpecifier] class.
//
// # Initializing a whose specifier
//
//   - [INSWhoseSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyTest]: Returns an [NSWhoseSpecifier] object initialized with the given attributes.
//
// # Accessing information about a whose specifier
//
//   - [INSWhoseSpecifier.EndSubelementIdentifier]: Sets the end sub-element identifier for the specifier to the value of a given sub-element.
//   - [INSWhoseSpecifier.SetEndSubelementIdentifier]
//   - [INSWhoseSpecifier.EndSubelementIndex]: Sets the index position of the last sub-element within the range of objects being tested that pass the specifier’s test.
//   - [INSWhoseSpecifier.SetEndSubelementIndex]
//   - [INSWhoseSpecifier.StartSubelementIdentifier]: Returns the start sub-element identifier for the receiver.
//   - [INSWhoseSpecifier.SetStartSubelementIdentifier]
//   - [INSWhoseSpecifier.StartSubelementIndex]: Returns the index position of the first sub-element within the range of objects being tested that pass the receiver’s test.
//   - [INSWhoseSpecifier.SetStartSubelementIndex]
//   - [INSWhoseSpecifier.Test]: Returns the test object encapsulated by the receiver.
//   - [INSWhoseSpecifier.SetTest]
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier
type INSWhoseSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Initializing a whose specifier

	// Returns an [NSWhoseSpecifier] object initialized with the given attributes.
	InitWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, test INSScriptWhoseTest) NSWhoseSpecifier

	// Topic: Accessing information about a whose specifier

	// Sets the end sub-element identifier for the specifier to the value of a given sub-element.
	EndSubelementIdentifier() NSWhoseSubelementIdentifier
	SetEndSubelementIdentifier(value NSWhoseSubelementIdentifier)
	// Sets the index position of the last sub-element within the range of objects being tested that pass the specifier’s test.
	EndSubelementIndex() int
	SetEndSubelementIndex(value int)
	// Returns the start sub-element identifier for the receiver.
	StartSubelementIdentifier() NSWhoseSubelementIdentifier
	SetStartSubelementIdentifier(value NSWhoseSubelementIdentifier)
	// Returns the index position of the first sub-element within the range of objects being tested that pass the receiver’s test.
	StartSubelementIndex() int
	SetStartSubelementIndex(value int)
	// Returns the test object encapsulated by the receiver.
	Test() INSScriptWhoseTest
	SetTest(value INSScriptWhoseTest)
}

// Init initializes the instance.
func (w NSWhoseSpecifier) Init() NSWhoseSpecifier {
	rv := objc.Send[NSWhoseSpecifier](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWhoseSpecifier) Autorelease() NSWhoseSpecifier {
	rv := objc.Send[NSWhoseSpecifier](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWhoseSpecifier creates a new NSWhoseSpecifier instance.
func NewNSWhoseSpecifier() NSWhoseSpecifier {
	class := getNSWhoseSpecifierClass()
	rv := objc.Send[NSWhoseSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/init(coder:)
func NewWhoseSpecifierWithCoder(inCoder INSCoder) NSWhoseSpecifier {
	instance := getNSWhoseSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSWhoseSpecifierFromID(rv)
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
func NewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSWhoseSpecifier {
	instance := getNSWhoseSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSWhoseSpecifierFromID(rv)
}

// Returns an [NSWhoseSpecifier] object initialized with the given attributes.
//
// classDesc: Class description for the receiver’s container object.
//
// container: An object specifier for the receiver’s container object.
//
// property: The key for the property for which to test.
//
// test: The test condition.
//
// # Return Value
// 
// An [NSWhoseSpecifier] object initialized with the given attributes.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] and sets the whose
// test condition to `test`.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/init(containerClassDescription:containerSpecifier:key:test:)
func NewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, test INSScriptWhoseTest) NSWhoseSpecifier {
	instance := getNSWhoseSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:test:"), classDesc, container, objc.String(property), test)
	return NSWhoseSpecifierFromID(rv)
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
func NewWhoseSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSWhoseSpecifier {
	instance := getNSWhoseSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSWhoseSpecifierFromID(rv)
}

// Returns an [NSWhoseSpecifier] object initialized with the given attributes.
//
// classDesc: Class description for the receiver’s container object.
//
// container: An object specifier for the receiver’s container object.
//
// property: The key for the property for which to test.
//
// test: The test condition.
//
// # Return Value
// 
// An [NSWhoseSpecifier] object initialized with the given attributes.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] and sets the whose
// test condition to `test`.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/init(containerClassDescription:containerSpecifier:key:test:)
func (w NSWhoseSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyTest(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, test INSScriptWhoseTest) NSWhoseSpecifier {
	rv := objc.Send[NSWhoseSpecifier](w.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:test:"), classDesc, container, objc.String(property), test)
	return rv
}

// Sets the end sub-element identifier for the specifier to the value of a
// given sub-element.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/endSubelementIdentifier
func (w NSWhoseSpecifier) EndSubelementIdentifier() NSWhoseSubelementIdentifier {
	rv := objc.Send[NSWhoseSubelementIdentifier](w.ID, objc.Sel("endSubelementIdentifier"))
	return NSWhoseSubelementIdentifier(rv)
}
func (w NSWhoseSpecifier) SetEndSubelementIdentifier(value NSWhoseSubelementIdentifier) {
	objc.Send[struct{}](w.ID, objc.Sel("setEndSubelementIdentifier:"), value)
}
// Sets the index position of the last sub-element within the range of objects
// being tested that pass the specifier’s test.
//
// # Discussion
// 
// Used only if the end sub-element identifier is [NSIndexSubelement].
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/endSubelementIndex
func (w NSWhoseSpecifier) EndSubelementIndex() int {
	rv := objc.Send[int](w.ID, objc.Sel("endSubelementIndex"))
	return rv
}
func (w NSWhoseSpecifier) SetEndSubelementIndex(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setEndSubelementIndex:"), value)
}
// Returns the start sub-element identifier for the receiver.
//
// # Return Value
// 
// The start sub-element identifier for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/startSubelementIdentifier
func (w NSWhoseSpecifier) StartSubelementIdentifier() NSWhoseSubelementIdentifier {
	rv := objc.Send[NSWhoseSubelementIdentifier](w.ID, objc.Sel("startSubelementIdentifier"))
	return NSWhoseSubelementIdentifier(rv)
}
func (w NSWhoseSpecifier) SetStartSubelementIdentifier(value NSWhoseSubelementIdentifier) {
	objc.Send[struct{}](w.ID, objc.Sel("setStartSubelementIdentifier:"), value)
}
// Returns the index position of the first sub-element within the range of
// objects being tested that pass the receiver’s test.
//
// # Return Value
// 
// The index position of the first sub-element within the range of objects
// being tested that pass the receiver’s test.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/startSubelementIndex
func (w NSWhoseSpecifier) StartSubelementIndex() int {
	rv := objc.Send[int](w.ID, objc.Sel("startSubelementIndex"))
	return rv
}
func (w NSWhoseSpecifier) SetStartSubelementIndex(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setStartSubelementIndex:"), value)
}
// Returns the test object encapsulated by the receiver.
//
// # Return Value
// 
// The test object encapsulated by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/test
func (w NSWhoseSpecifier) Test() INSScriptWhoseTest {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("test"))
	return NSScriptWhoseTestFromID(objc.ID(rv))
}
func (w NSWhoseSpecifier) SetTest(value INSScriptWhoseTest) {
	objc.Send[struct{}](w.ID, objc.Sel("setTest:"), value)
}

