// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSIndexSpecifier] class.
var (
	_NSIndexSpecifierClass     NSIndexSpecifierClass
	_NSIndexSpecifierClassOnce sync.Once
)

func getNSIndexSpecifierClass() NSIndexSpecifierClass {
	_NSIndexSpecifierClassOnce.Do(func() {
		_NSIndexSpecifierClass = NSIndexSpecifierClass{class: objc.GetClass("NSIndexSpecifier")}
	})
	return _NSIndexSpecifierClass
}

// GetNSIndexSpecifierClass returns the class object for NSIndexSpecifier.
func GetNSIndexSpecifierClass() NSIndexSpecifierClass {
	return getNSIndexSpecifierClass()
}

type NSIndexSpecifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSIndexSpecifierClass) Alloc() NSIndexSpecifier {
	rv := objc.Send[NSIndexSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier representing an object in a collection (or container) with an
// index number.
//
// # Overview
// 
// The script terms `first` and `front` specify the object with index `0`,
// while `last` specifies the object with index of `count-1`. A negative index
// indicates a location by counting backward from the last object in the
// collection.
// 
// You don’t normally subclass [NSIndexSpecifier].
//
// # Creating Index Specifiers
//
//   - [NSIndexSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyIndex]: Initializes an allocated [NSIndexSpecifier](<doc://com.apple.foundation/documentation/Foundation/NSIndexSpecifier>) object with a class description, container specifier, collection key, and object index.
//
// # Accessing the Index
//
//   - [NSIndexSpecifier.Index]: Sets the value of the receiver’s `index` property.
//   - [NSIndexSpecifier.SetIndex]
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier
type NSIndexSpecifier struct {
	NSScriptObjectSpecifier
}

// NSIndexSpecifierFromID constructs a [NSIndexSpecifier] from an objc.ID.
//
// A specifier representing an object in a collection (or container) with an
// index number.
func NSIndexSpecifierFromID(id objc.ID) NSIndexSpecifier {
	return NSIndexSpecifier{NSScriptObjectSpecifier: NSScriptObjectSpecifierFromID(id)}
}
// NOTE: NSIndexSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSIndexSpecifier] class.
//
// # Creating Index Specifiers
//
//   - [INSIndexSpecifier.InitWithContainerClassDescriptionContainerSpecifierKeyIndex]: Initializes an allocated [NSIndexSpecifier](<doc://com.apple.foundation/documentation/Foundation/NSIndexSpecifier>) object with a class description, container specifier, collection key, and object index.
//
// # Accessing the Index
//
//   - [INSIndexSpecifier.Index]: Sets the value of the receiver’s `index` property.
//   - [INSIndexSpecifier.SetIndex]
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier
type INSIndexSpecifier interface {
	INSScriptObjectSpecifier

	// Topic: Creating Index Specifiers

	// Initializes an allocated [NSIndexSpecifier](<doc://com.apple.foundation/documentation/Foundation/NSIndexSpecifier>) object with a class description, container specifier, collection key, and object index.
	InitWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, index int) NSIndexSpecifier

	// Topic: Accessing the Index

	// Sets the value of the receiver’s `index` property.
	Index() int
	SetIndex(value int)
}

// Init initializes the instance.
func (i NSIndexSpecifier) Init() NSIndexSpecifier {
	rv := objc.Send[NSIndexSpecifier](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSIndexSpecifier) Autorelease() NSIndexSpecifier {
	rv := objc.Send[NSIndexSpecifier](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSIndexSpecifier creates a new NSIndexSpecifier instance.
func NewNSIndexSpecifier() NSIndexSpecifier {
	class := getNSIndexSpecifierClass()
	rv := objc.Send[NSIndexSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func NewIndexSpecifierWithCoder(inCoder INSCoder) NSIndexSpecifier {
	instance := getNSIndexSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSIndexSpecifierFromID(rv)
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
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSIndexSpecifier {
	instance := getNSIndexSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSIndexSpecifierFromID(rv)
}

// Initializes an allocated [NSIndexSpecifier] object with a class
// description, container specifier, collection key, and object index.
//
// classDesc: Description for the container of the collection.
//
// container: Container of the collection.
//
// property: Name of the collection.
//
// index: The object within the `key` collection the index specifier is to identify.
//
// # Return Value
// 
// Initialized [NSIndexSpecifier] object with its `index` property set to
// `objectIndex`.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and sets
// the `index` property of the index specifier to `objectIndex`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/init(containerClassDescription:containerSpecifier:key:index:)
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, index int) NSIndexSpecifier {
	instance := getNSIndexSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:index:"), classDesc, container, objc.String(property), index)
	return NSIndexSpecifierFromID(rv)
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
func NewIndexSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSIndexSpecifier {
	instance := getNSIndexSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSIndexSpecifierFromID(rv)
}

// Initializes an allocated [NSIndexSpecifier] object with a class
// description, container specifier, collection key, and object index.
//
// classDesc: Description for the container of the collection.
//
// container: Container of the collection.
//
// property: Name of the collection.
//
// index: The object within the `key` collection the index specifier is to identify.
//
// # Return Value
// 
// Initialized [NSIndexSpecifier] object with its `index` property set to
// `objectIndex`.
//
// # Discussion
// 
// Invokes the super class’s
// [InitWithContainerClassDescriptionContainerSpecifierKey] method and sets
// the `index` property of the index specifier to `objectIndex`.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/init(containerClassDescription:containerSpecifier:key:index:)
func (i NSIndexSpecifier) InitWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string, index int) NSIndexSpecifier {
	rv := objc.Send[NSIndexSpecifier](i.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:index:"), classDesc, container, objc.String(property), index)
	return rv
}

// Sets the value of the receiver’s `index` property.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/index
func (i NSIndexSpecifier) Index() int {
	rv := objc.Send[int](i.ID, objc.Sel("index"))
	return rv
}
func (i NSIndexSpecifier) SetIndex(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setIndex:"), value)
}

