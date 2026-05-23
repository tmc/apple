// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [JSVirtualMachine] class.
var (
	_JSVirtualMachineClass     JSVirtualMachineClass
	_JSVirtualMachineClassOnce sync.Once
)

func getJSVirtualMachineClass() JSVirtualMachineClass {
	_JSVirtualMachineClassOnce.Do(func() {
		_JSVirtualMachineClass = JSVirtualMachineClass{class: objc.GetClass("JSVirtualMachine")}
	})
	return _JSVirtualMachineClass
}

// GetJSVirtualMachineClass returns the class object for JSVirtualMachine.
func GetJSVirtualMachineClass() JSVirtualMachineClass {
	return getJSVirtualMachineClass()
}

type JSVirtualMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (jc JSVirtualMachineClass) Class() objc.Class {
	return jc.class
}

// Alloc allocates memory for a new instance of the class.
func (jc JSVirtualMachineClass) Alloc() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// A self-contained environment for JavaScript execution.
//
// # Overview
//
// You use this class for two main purposes: to support concurrent JavaScript
// execution, and to manage memory for objects that bridge between JavaScript
// and Objective-C or Swift.
//
// # Support Threading and Concurrent JavaScript Execution
//
// Each JavaScript context (a [JSContext] object) belongs to a virtual
// machine. Each virtual machine can encompass multiple contexts, allowing
// values ([JSValue] objects) to pass between contexts. However, each virtual
// machine is distinct—you can’t pass a value that you create in one
// virtual machine to a context in another virtual machine.
//
// The JavaScriptCore API is thread-safe—for example, you can create
// [JSValue] objects or evaluate scripts from any thread—however, all other
// threads attempting to use the same virtual machine must wait. To run
// JavaScript concurrently on multiple threads, use a separate
// [JSVirtualMachine] instance for each thread.
//
// # Manage Memory for Exported Objects
//
// When you export an Objective-C or Swift object to JavaScript, you must not
// to store JavaScript values in that object. This action creates a retain
// cycle—[JSValue] objects hold strong references to their enclosing
// JavaScript contexts, and [JSContext] objects hold strong references to the
// native objects you export to JavaScript. Instead, use the [JSManagedValue]
// class to conditionally retain a JavaScript value, and report the native
// ownership chain for that managed value to the JavaScriptCore virtual
// machine. Use the [JSVirtualMachine.AddManagedReferenceWithOwner] and
// [JSVirtualMachine.RemoveManagedReferenceWithOwner] methods to describe your
// native object graph to JavaScriptCore. After you remove the last managed
// reference for an object, the JavaScript garbage collector can safely
// destroy that object.
//
// # Managing Memory for Bridged Values
//
//   - [JSVirtualMachine.AddManagedReferenceWithOwner]: Notifies the JavaScriptCore virtual machine of an external object relationship.
//   - [JSVirtualMachine.RemoveManagedReferenceWithOwner]: Notifies the JavaScriptCore virtual machine that a previously registered object relationship no longer exists.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine
type JSVirtualMachine struct {
	objectivec.Object
}

// JSVirtualMachineFromID constructs a [JSVirtualMachine] from an objc.ID.
//
// A self-contained environment for JavaScript execution.
func JSVirtualMachineFromID(id objc.ID) JSVirtualMachine {
	return JSVirtualMachine{objectivec.Object{ID: id}}
}

// NOTE: JSVirtualMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [JSVirtualMachine] class.
//
// # Managing Memory for Bridged Values
//
//   - [IJSVirtualMachine.AddManagedReferenceWithOwner]: Notifies the JavaScriptCore virtual machine of an external object relationship.
//   - [IJSVirtualMachine.RemoveManagedReferenceWithOwner]: Notifies the JavaScriptCore virtual machine that a previously registered object relationship no longer exists.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine
type IJSVirtualMachine interface {
	objectivec.IObject

	// Topic: Managing Memory for Bridged Values

	// Notifies the JavaScriptCore virtual machine of an external object relationship.
	AddManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject)
	// Notifies the JavaScriptCore virtual machine that a previously registered object relationship no longer exists.
	RemoveManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject)
}

// Init initializes the instance.
func (j JSVirtualMachine) Init() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j JSVirtualMachine) Autorelease() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSVirtualMachine creates a new JSVirtualMachine instance.
func NewJSVirtualMachine() JSVirtualMachine {
	class := getJSVirtualMachineClass()
	rv := objc.Send[JSVirtualMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Notifies the JavaScriptCore virtual machine of an external object
// relationship.
//
// object: The object to be referenced by the JavaScript memory management graph.
//
// owner: The other object responsible for the lifetime of the reference.
//
// # Discussion
//
// Use this method to make the JavaScript runtime aware of arbitrary external
// Objective-C or Swift object graphs. The runtime can then use this
// information to retain any JavaScript values that are referenced from
// somewhere in said object graph.
//
// For correct behavior, clients must make their external object graphs
// reachable from within the JavaScript runtime. If an Objective-C or Swift
// object is reachable from within the JavaScript runtime, all managed
// references transitively reachable from it as recorded using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method are scanned by the
// garbage collector.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/addManagedReference(_:withOwner:)
func (j JSVirtualMachine) AddManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject) {
	objc.Send[objc.ID](j.ID, objc.Sel("addManagedReference:withOwner:"), object, owner)
}

// Notifies the JavaScriptCore virtual machine that a previously registered
// object relationship no longer exists.
//
// object: The object formerly referenced by the JavaScript memory management graph.
//
// owner: The other object responsible for the lifetime of the reference.
//
// # Discussion
//
// Use this method to deregister object relationships recorded using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method.
//
// The JavaScript garbage collector continues to scan any references that were
// reported to it until you use this method to remove those references.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/removeManagedReference(_:withOwner:)
func (j JSVirtualMachine) RemoveManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject) {
	objc.Send[objc.ID](j.ID, objc.Sel("removeManagedReference:withOwner:"), object, owner)
}
