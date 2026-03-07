// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSQuitCommand] class.
var (
	_NSQuitCommandClass     NSQuitCommandClass
	_NSQuitCommandClassOnce sync.Once
)

func getNSQuitCommandClass() NSQuitCommandClass {
	_NSQuitCommandClassOnce.Do(func() {
		_NSQuitCommandClass = NSQuitCommandClass{class: objc.GetClass("NSQuitCommand")}
	})
	return _NSQuitCommandClass
}

// GetNSQuitCommandClass returns the class object for NSQuitCommand.
func GetNSQuitCommandClass() NSQuitCommandClass {
	return getNSQuitCommandClass()
}

type NSQuitCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSQuitCommandClass) Alloc() NSQuitCommand {
	rv := objc.Send[NSQuitCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A command that quits the specified app.
//
// # Overview
// 
// The quit command may optionally specify how to handle modified documents
// (automatically save changes, don’t save them, or ask the user). For
// details, see the description for the `quit` command in “Apple Events Sent
// By the Mac OS” in [How Cocoa Applications Handle Apple Events] in [Cocoa
// Scripting Guide].
// 
// [NSQuitCommand] is part of Cocoa’s built-in scripting support. Most
// applications don’t need to subclass [NSQuitCommand] or call its methods.
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [How Cocoa Applications Handle Apple Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_handle_AEs/SAppsHandleAEs.html#//apple_ref/doc/uid/20001239
//
// # Accessing options
//
//   - [NSQuitCommand.SaveOptions]: Returns a constant indicating how to deal with closing any modified documents.
//
// See: https://developer.apple.com/documentation/Foundation/NSQuitCommand
type NSQuitCommand struct {
	NSScriptCommand
}

// NSQuitCommandFromID constructs a [NSQuitCommand] from an objc.ID.
//
// A command that quits the specified app.
func NSQuitCommandFromID(id objc.ID) NSQuitCommand {
	return NSQuitCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSQuitCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSQuitCommand] class.
//
// # Accessing options
//
//   - [INSQuitCommand.SaveOptions]: Returns a constant indicating how to deal with closing any modified documents.
//
// See: https://developer.apple.com/documentation/Foundation/NSQuitCommand
type INSQuitCommand interface {
	INSScriptCommand

	// Topic: Accessing options

	// Returns a constant indicating how to deal with closing any modified documents.
	SaveOptions() NSSaveOptions
}





// Init initializes the instance.
func (q NSQuitCommand) Init() NSQuitCommand {
	rv := objc.Send[NSQuitCommand](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q NSQuitCommand) Autorelease() NSQuitCommand {
	rv := objc.Send[NSQuitCommand](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSQuitCommand creates a new NSQuitCommand instance.
func NewNSQuitCommand() NSQuitCommand {
	class := getNSQuitCommandClass()
	rv := objc.Send[NSQuitCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewQuitCommandWithCoder(inCoder INSCoder) NSQuitCommand {
	instance := getNSQuitCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSQuitCommandFromID(rv)
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
func NewQuitCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSQuitCommand {
	instance := getNSQuitCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSQuitCommandFromID(rv)
}


















// Returns a constant indicating how to deal with closing any modified
// documents.
//
// # Return Value
// 
// A constant indicating how to deal with closing any modified documents. The
// default value returned is [NSSaveOptionsAsk]. See “Constants” in
// [NSCloseCommand] for a list of possible return values.
//
// See: https://developer.apple.com/documentation/Foundation/NSQuitCommand/saveOptions
func (q NSQuitCommand) SaveOptions() NSSaveOptions {
	rv := objc.Send[NSSaveOptions](q.ID, objc.Sel("saveOptions"))
	return NSSaveOptions(rv)
}


























