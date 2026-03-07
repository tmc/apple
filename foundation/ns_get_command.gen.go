// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSGetCommand] class.
var (
	_NSGetCommandClass     NSGetCommandClass
	_NSGetCommandClassOnce sync.Once
)

func getNSGetCommandClass() NSGetCommandClass {
	_NSGetCommandClassOnce.Do(func() {
		_NSGetCommandClass = NSGetCommandClass{class: objc.GetClass("NSGetCommand")}
	})
	return _NSGetCommandClass
}

// GetNSGetCommandClass returns the class object for NSGetCommand.
func GetNSGetCommandClass() NSGetCommandClass {
	return getNSGetCommandClass()
}

type NSGetCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGetCommandClass) Alloc() NSGetCommand {
	rv := objc.Send[NSGetCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A command that retrieves a value or object from a scriptable object.
//
// # Overview
// 
// An instance of [NSGetCommand] gets the specified value or object from the
// specified scriptable object: for example, the words from a paragraph or the
// name of a document.
// 
// When an instance of [NSGetCommand] is executed, it evaluates the specified
// receivers, gathers the specified data, if any, and packages it in a return
// Apple event.
// 
// [NSGetCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `get` command through key-value coding. Most
// applications don’t need to subclass [NSGetCommand] or call its methods.
// 
// For information on working with `get` commands, see [Getting and Setting
// Properties and Elements] in [Cocoa Scripting Guide].
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Getting and Setting Properties and Elements]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_get_set/SAppsGetSet.html#//apple_ref/doc/uid/TP40002164-CH18
//
// See: https://developer.apple.com/documentation/Foundation/NSGetCommand
type NSGetCommand struct {
	NSScriptCommand
}

// NSGetCommandFromID constructs a [NSGetCommand] from an objc.ID.
//
// A command that retrieves a value or object from a scriptable object.
func NSGetCommandFromID(id objc.ID) NSGetCommand {
	return NSGetCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSGetCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSGetCommand] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSGetCommand
type INSGetCommand interface {
	INSScriptCommand
}





// Init initializes the instance.
func (g NSGetCommand) Init() NSGetCommand {
	rv := objc.Send[NSGetCommand](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGetCommand) Autorelease() NSGetCommand {
	rv := objc.Send[NSGetCommand](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGetCommand creates a new NSGetCommand instance.
func NewNSGetCommand() NSGetCommand {
	class := getNSGetCommandClass()
	rv := objc.Send[NSGetCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewGetCommandWithCoder(inCoder INSCoder) NSGetCommand {
	instance := getNSGetCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSGetCommandFromID(rv)
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
func NewGetCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSGetCommand {
	instance := getNSGetCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSGetCommandFromID(rv)
}









































