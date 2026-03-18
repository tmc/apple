// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCloseCommand] class.
var (
	_NSCloseCommandClass     NSCloseCommandClass
	_NSCloseCommandClassOnce sync.Once
)

func getNSCloseCommandClass() NSCloseCommandClass {
	_NSCloseCommandClassOnce.Do(func() {
		_NSCloseCommandClass = NSCloseCommandClass{class: objc.GetClass("NSCloseCommand")}
	})
	return _NSCloseCommandClass
}

// GetNSCloseCommandClass returns the class object for NSCloseCommand.
func GetNSCloseCommandClass() NSCloseCommandClass {
	return getNSCloseCommandClass()
}

type NSCloseCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCloseCommandClass) Alloc() NSCloseCommand {
	rv := objc.Send[NSCloseCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that closes one or more scriptable objects.
//
// # Overview
// 
// An instance of [NSCloseCommand] closes the specified scriptable object or
// objects—typically a document or window (and its associated document, if
// any). The command may optionally specify a location to save in and how to
// handle modified documents (by automatically saving changes, not saving
// them, or asking the user).
// 
// [NSCloseCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `close` command through key-value coding. Most
// applications don’t need to subclass [NSCloseCommand] or call its methods.
//
// # Accessing save options
//
//   - [NSCloseCommand.SaveOptions]: Returns a constant indicating how to deal with closing any modified documents.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloseCommand
type NSCloseCommand struct {
	NSScriptCommand
}

// NSCloseCommandFromID constructs a [NSCloseCommand] from an objc.ID.
//
// A command that closes one or more scriptable objects.
func NSCloseCommandFromID(id objc.ID) NSCloseCommand {
	return NSCloseCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSCloseCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCloseCommand] class.
//
// # Accessing save options
//
//   - [INSCloseCommand.SaveOptions]: Returns a constant indicating how to deal with closing any modified documents.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloseCommand
type INSCloseCommand interface {
	INSScriptCommand

	// Topic: Accessing save options

	// Returns a constant indicating how to deal with closing any modified documents.
	SaveOptions() NSSaveOptions
}

// Init initializes the instance.
func (c NSCloseCommand) Init() NSCloseCommand {
	rv := objc.Send[NSCloseCommand](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCloseCommand) Autorelease() NSCloseCommand {
	rv := objc.Send[NSCloseCommand](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCloseCommand creates a new NSCloseCommand instance.
func NewNSCloseCommand() NSCloseCommand {
	class := getNSCloseCommandClass()
	rv := objc.Send[NSCloseCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewCloseCommandWithCoder(inCoder INSCoder) NSCloseCommand {
	instance := getNSCloseCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSCloseCommandFromID(rv)
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
func NewCloseCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSCloseCommand {
	instance := getNSCloseCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSCloseCommandFromID(rv)
}

// Returns a constant indicating how to deal with closing any modified
// documents.
//
// # Return Value
// 
// A constant indicating how to deal with closing any modified documents. The
// default value returned is [NSSaveOptionsAsk]. See Constants for a list of
// possible return values.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloseCommand/saveOptions
func (c NSCloseCommand) SaveOptions() NSSaveOptions {
	rv := objc.Send[NSSaveOptions](c.ID, objc.Sel("saveOptions"))
	return NSSaveOptions(rv)
}

