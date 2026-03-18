// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptExecutionContext] class.
var (
	_NSScriptExecutionContextClass     NSScriptExecutionContextClass
	_NSScriptExecutionContextClassOnce sync.Once
)

func getNSScriptExecutionContextClass() NSScriptExecutionContextClass {
	_NSScriptExecutionContextClassOnce.Do(func() {
		_NSScriptExecutionContextClass = NSScriptExecutionContextClass{class: objc.GetClass("NSScriptExecutionContext")}
	})
	return _NSScriptExecutionContextClass
}

// GetNSScriptExecutionContextClass returns the class object for NSScriptExecutionContext.
func GetNSScriptExecutionContextClass() NSScriptExecutionContextClass {
	return getNSScriptExecutionContextClass()
}

type NSScriptExecutionContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptExecutionContextClass) Alloc() NSScriptExecutionContext {
	rv := objc.Send[NSScriptExecutionContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The context in which the current script command is executed.
//
// # Overview
// 
// An [NSScriptExecutionContext] object is a shared instance (there is only
// one instance of the class) that represents the context in which the current
// script command is executed. [NSScriptExecutionContext] tracks global state
// relating to the command being executed, especially the top-level container
// object (that is, the container implied by a specifier object that specifies
// no container) used in an evaluation of an [NSScriptObjectSpecifier] object.
// 
// In most cases, the top-level container for a complete series of nested
// object specifiers is automatically set to the application object ([NSApp]),
// and you can get this object with the [NSScriptExecutionContext.TopLevelObject] method. But you can
// also set this top-level container to something else (using
// [NSScriptExecutionContext.TopLevelObject]) if the situation warrants it.
// 
// It is unlikely that you will need to subclass [NSScriptExecutionContext].
//
// # Getting and setting the container object
//
//   - [NSScriptExecutionContext.TopLevelObject]: Sets the top-level object for an object-specifier evaluation.
//   - [NSScriptExecutionContext.SetTopLevelObject]
//   - [NSScriptExecutionContext.ObjectBeingTested]: Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//   - [NSScriptExecutionContext.SetObjectBeingTested]
//   - [NSScriptExecutionContext.RangeContainerObject]: Sets the top-level container object for a range-specifier evaluation to a give object.
//   - [NSScriptExecutionContext.SetRangeContainerObject]
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext
type NSScriptExecutionContext struct {
	objectivec.Object
}

// NSScriptExecutionContextFromID constructs a [NSScriptExecutionContext] from an objc.ID.
//
// The context in which the current script command is executed.
func NSScriptExecutionContextFromID(id objc.ID) NSScriptExecutionContext {
	return NSScriptExecutionContext{objectivec.Object{ID: id}}
}
// NOTE: NSScriptExecutionContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptExecutionContext] class.
//
// # Getting and setting the container object
//
//   - [INSScriptExecutionContext.TopLevelObject]: Sets the top-level object for an object-specifier evaluation.
//   - [INSScriptExecutionContext.SetTopLevelObject]
//   - [INSScriptExecutionContext.ObjectBeingTested]: Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//   - [INSScriptExecutionContext.SetObjectBeingTested]
//   - [INSScriptExecutionContext.RangeContainerObject]: Sets the top-level container object for a range-specifier evaluation to a give object.
//   - [INSScriptExecutionContext.SetRangeContainerObject]
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext
type INSScriptExecutionContext interface {
	objectivec.IObject

	// Topic: Getting and setting the container object

	// Sets the top-level object for an object-specifier evaluation.
	TopLevelObject() objectivec.IObject
	SetTopLevelObject(value objectivec.IObject)
	// Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
	ObjectBeingTested() objectivec.IObject
	SetObjectBeingTested(value objectivec.IObject)
	// Sets the top-level container object for a range-specifier evaluation to a give object.
	RangeContainerObject() objectivec.IObject
	SetRangeContainerObject(value objectivec.IObject)
}

// Init initializes the instance.
func (s NSScriptExecutionContext) Init() NSScriptExecutionContext {
	rv := objc.Send[NSScriptExecutionContext](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptExecutionContext) Autorelease() NSScriptExecutionContext {
	rv := objc.Send[NSScriptExecutionContext](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptExecutionContext creates a new NSScriptExecutionContext instance.
func NewNSScriptExecutionContext() NSScriptExecutionContext {
	class := getNSScriptExecutionContextClass()
	rv := objc.Send[NSScriptExecutionContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the shared [NSScriptExecutionContext] instance.
//
// # Return Value
// 
// The shared [NSScriptExecutionContext] instance, creating it first if it
// doesn’t exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/shared()
func (_NSScriptExecutionContextClass NSScriptExecutionContextClass) SharedScriptExecutionContext() NSScriptExecutionContext {
	rv := objc.Send[objc.ID](objc.ID(_NSScriptExecutionContextClass.class), objc.Sel("sharedScriptExecutionContext"))
	return NSScriptExecutionContextFromID(rv)
}

// Sets the top-level object for an object-specifier evaluation.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/topLevelObject
func (s NSScriptExecutionContext) TopLevelObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("topLevelObject"))
	return objectivec.Object{ID: rv}
}
func (s NSScriptExecutionContext) SetTopLevelObject(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setTopLevelObject:"), value)
}

// Sets the top-level container object currently being tested in a “whose”
// qualifier to a given object.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/objectBeingTested
func (s NSScriptExecutionContext) ObjectBeingTested() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectBeingTested"))
	return objectivec.Object{ID: rv}
}
func (s NSScriptExecutionContext) SetObjectBeingTested(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setObjectBeingTested:"), value)
}

// Sets the top-level container object for a range-specifier evaluation to a
// give object.
//
// # Discussion
// 
// Instances of [NSRangeSpecifier] contain object specifiers representing the
// first or last element in a range of elements, and these specifiers are
// evaluated in the context of `container`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/rangeContainerObject
func (s NSScriptExecutionContext) RangeContainerObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("rangeContainerObject"))
	return objectivec.Object{ID: rv}
}
func (s NSScriptExecutionContext) SetRangeContainerObject(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setRangeContainerObject:"), value)
}

