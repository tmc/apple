// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAccessibilityCustomAction] class.
var (
	_NSAccessibilityCustomActionClass     NSAccessibilityCustomActionClass
	_NSAccessibilityCustomActionClassOnce sync.Once
)

func getNSAccessibilityCustomActionClass() NSAccessibilityCustomActionClass {
	_NSAccessibilityCustomActionClassOnce.Do(func() {
		_NSAccessibilityCustomActionClass = NSAccessibilityCustomActionClass{class: objc.GetClass("NSAccessibilityCustomAction")}
	})
	return _NSAccessibilityCustomActionClass
}

// GetNSAccessibilityCustomActionClass returns the class object for NSAccessibilityCustomAction.
func GetNSAccessibilityCustomActionClass() NSAccessibilityCustomActionClass {
	return getNSAccessibilityCustomActionClass()
}

type NSAccessibilityCustomActionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAccessibilityCustomActionClass) Alloc() NSAccessibilityCustomAction {
	rv := objc.Send[NSAccessibilityCustomAction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A custom action to perform on an accessible object.
//
// # Overview
// 
// Apps that support custom actions can create instances of this class,
// specifying the user-readable name of the action, and either a handler
// closure or the object and selector to use when performing the action.
// Assistive apps display custom actions in response to specific user cues.
// For example, VoiceOver lets users access actions quickly using the Actions
// rotor.
// 
// After creating an instance of this class, add it to the
// [accessibilityCustomActions] property of an appropriate accessible object.
//
// [accessibilityCustomActions]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
//
// # Creating a Custom Action
//
//   - [NSAccessibilityCustomAction.InitWithNameHandler]: Creates a custom action object with the specified name and handler.
//   - [NSAccessibilityCustomAction.InitWithNameTargetSelector]: Creates a custom action object with the specified name, target, and selector.
//
// # Getting the Action Name
//
//   - [NSAccessibilityCustomAction.Name]: A localized name that describes the action.
//   - [NSAccessibilityCustomAction.SetName]
//
// # Getting the Action
//
//   - [NSAccessibilityCustomAction.Handler]: The closure that handles the execution of the action.
//   - [NSAccessibilityCustomAction.SetHandler]
//   - [NSAccessibilityCustomAction.Target]: The object that performs the action through a selector.
//   - [NSAccessibilityCustomAction.SetTarget]
//   - [NSAccessibilityCustomAction.Selector]: The method to call on the target to perform the action.
//   - [NSAccessibilityCustomAction.SetSelector]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction
type NSAccessibilityCustomAction struct {
	objectivec.Object
}

// NSAccessibilityCustomActionFromID constructs a [NSAccessibilityCustomAction] from an objc.ID.
//
// A custom action to perform on an accessible object.
func NSAccessibilityCustomActionFromID(id objc.ID) NSAccessibilityCustomAction {
	return NSAccessibilityCustomAction{objectivec.Object{ID: id}}
}
// NOTE: NSAccessibilityCustomAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAccessibilityCustomAction] class.
//
// # Creating a Custom Action
//
//   - [INSAccessibilityCustomAction.InitWithNameHandler]: Creates a custom action object with the specified name and handler.
//   - [INSAccessibilityCustomAction.InitWithNameTargetSelector]: Creates a custom action object with the specified name, target, and selector.
//
// # Getting the Action Name
//
//   - [INSAccessibilityCustomAction.Name]: A localized name that describes the action.
//   - [INSAccessibilityCustomAction.SetName]
//
// # Getting the Action
//
//   - [INSAccessibilityCustomAction.Handler]: The closure that handles the execution of the action.
//   - [INSAccessibilityCustomAction.SetHandler]
//   - [INSAccessibilityCustomAction.Target]: The object that performs the action through a selector.
//   - [INSAccessibilityCustomAction.SetTarget]
//   - [INSAccessibilityCustomAction.Selector]: The method to call on the target to perform the action.
//   - [INSAccessibilityCustomAction.SetSelector]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction
type INSAccessibilityCustomAction interface {
	objectivec.IObject

	// Topic: Creating a Custom Action

	// Creates a custom action object with the specified name and handler.
	InitWithNameHandler(name string, handler ErrorHandler) NSAccessibilityCustomAction
	// Creates a custom action object with the specified name, target, and selector.
	InitWithNameTargetSelector(name string, target objectivec.Object, selector objc.SEL) NSAccessibilityCustomAction

	// Topic: Getting the Action Name

	// A localized name that describes the action.
	Name() string
	SetName(value string)

	// Topic: Getting the Action

	// The closure that handles the execution of the action.
	Handler() VoidHandler
	SetHandler(value VoidHandler)
	// The object that performs the action through a selector.
	Target() objectivec.Object
	SetTarget(value objectivec.Object)
	// The method to call on the target to perform the action.
	Selector() objc.SEL
	SetSelector(value objc.SEL)
}

// Init initializes the instance.
func (a NSAccessibilityCustomAction) Init() NSAccessibilityCustomAction {
	rv := objc.Send[NSAccessibilityCustomAction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAccessibilityCustomAction) Autorelease() NSAccessibilityCustomAction {
	rv := objc.Send[NSAccessibilityCustomAction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAccessibilityCustomAction creates a new NSAccessibilityCustomAction instance.
func NewNSAccessibilityCustomAction() NSAccessibilityCustomAction {
	class := getNSAccessibilityCustomActionClass()
	rv := objc.Send[NSAccessibilityCustomAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a custom action object with the specified name and handler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:handler:)
func NewAccessibilityCustomActionWithNameHandler(name string, handler bool) NSAccessibilityCustomAction {
	instance := getNSAccessibilityCustomActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:handler:"), objc.String(name), handler)
	return NSAccessibilityCustomActionFromID(rv)
}

// Creates a custom action object with the specified name, target, and
// selector.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:target:selector:)
func NewAccessibilityCustomActionWithNameTargetSelector(name string, target objectivec.Object, selector objc.SEL) NSAccessibilityCustomAction {
	instance := getNSAccessibilityCustomActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:target:selector:"), objc.String(name), target, selector)
	return NSAccessibilityCustomActionFromID(rv)
}

// Creates a custom action object with the specified name and handler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:handler:)
func (a NSAccessibilityCustomAction) InitWithNameHandler(name string, handler ErrorHandler) NSAccessibilityCustomAction {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithName:handler:"), objc.String(name), _block1)
	return NSAccessibilityCustomActionFromID(rv)
}

// Creates a custom action object with the specified name, target, and
// selector.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:target:selector:)
func (a NSAccessibilityCustomAction) InitWithNameTargetSelector(name string, target objectivec.Object, selector objc.SEL) NSAccessibilityCustomAction {
	rv := objc.Send[NSAccessibilityCustomAction](a.ID, objc.Sel("initWithName:target:selector:"), objc.String(name), target, selector)
	return rv
}

// A localized name that describes the action.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/name
func (a NSAccessibilityCustomAction) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAccessibilityCustomAction) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}

// The closure that handles the execution of the action.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/handler
func (a NSAccessibilityCustomAction) Handler() VoidHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("handler"))
	_ = rv
	return nil
}
func (a NSAccessibilityCustomAction) SetHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setHandler:"), block)
}

// The object that performs the action through a selector.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/target
func (a NSAccessibilityCustomAction) Target() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("target"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a NSAccessibilityCustomAction) SetTarget(value objectivec.Object) {
	objc.Send[struct{}](a.ID, objc.Sel("setTarget:"), value)
}

// The method to call on the target to perform the action.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/selector
func (a NSAccessibilityCustomAction) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](a.ID, objc.Sel("selector"))
	return rv
}
func (a NSAccessibilityCustomAction) SetSelector(value objc.SEL) {
	objc.Send[struct{}](a.ID, objc.Sel("setSelector:"), value)
}

