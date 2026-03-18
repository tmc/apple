// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSExistsCommand] class.
var (
	_NSExistsCommandClass     NSExistsCommandClass
	_NSExistsCommandClassOnce sync.Once
)

func getNSExistsCommandClass() NSExistsCommandClass {
	_NSExistsCommandClassOnce.Do(func() {
		_NSExistsCommandClass = NSExistsCommandClass{class: objc.GetClass("NSExistsCommand")}
	})
	return _NSExistsCommandClass
}

// GetNSExistsCommandClass returns the class object for NSExistsCommand.
func GetNSExistsCommandClass() NSExistsCommandClass {
	return getNSExistsCommandClass()
}

type NSExistsCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExistsCommandClass) Alloc() NSExistsCommand {
	rv := objc.Send[NSExistsCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that determines whether a scriptable object exists.
//
// # Overview
// 
// An instance of [NSExistsCommand] determines whether a specified scriptable
// object, such as a word, paragraph, or image, exists.
// 
// When an instance of [NSExistsCommand] is executed, it evaluates the
// receiver specifier for the command to determine if it specifies any
// objects.
// 
// [NSExistsCommand] is part of Cocoa’s built-in scripting support. Most
// applications don’t need to subclass [NSExistsCommand].
//
// See: https://developer.apple.com/documentation/Foundation/NSExistsCommand
type NSExistsCommand struct {
	NSScriptCommand
}

// NSExistsCommandFromID constructs a [NSExistsCommand] from an objc.ID.
//
// A command that determines whether a scriptable object exists.
func NSExistsCommandFromID(id objc.ID) NSExistsCommand {
	return NSExistsCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSExistsCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExistsCommand] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSExistsCommand
type INSExistsCommand interface {
	INSScriptCommand
}

// Init initializes the instance.
func (e NSExistsCommand) Init() NSExistsCommand {
	rv := objc.Send[NSExistsCommand](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExistsCommand) Autorelease() NSExistsCommand {
	rv := objc.Send[NSExistsCommand](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExistsCommand creates a new NSExistsCommand instance.
func NewNSExistsCommand() NSExistsCommand {
	class := getNSExistsCommandClass()
	rv := objc.Send[NSExistsCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewExistsCommandWithCoder(inCoder INSCoder) NSExistsCommand {
	instance := getNSExistsCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSExistsCommandFromID(rv)
}

// Returns an a script command object initialized from the passed command
// description.
//
// commandDef: A command description for the command to be created.
//
// # Return Value
// 
// A newly initialized instance of [NSScriptCommand] or a subclass.
//
// # Discussion
// 
// To make this command object usable, you must set its receiving objects and
// arguments (if any) after invoking this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(commandDescription:)
func NewExistsCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSExistsCommand {
	instance := getNSExistsCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSExistsCommandFromID(rv)
}

