// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCountCommand] class.
var (
	_NSCountCommandClass     NSCountCommandClass
	_NSCountCommandClassOnce sync.Once
)

func getNSCountCommandClass() NSCountCommandClass {
	_NSCountCommandClassOnce.Do(func() {
		_NSCountCommandClass = NSCountCommandClass{class: objc.GetClass("NSCountCommand")}
	})
	return _NSCountCommandClass
}

// GetNSCountCommandClass returns the class object for NSCountCommand.
func GetNSCountCommandClass() NSCountCommandClass {
	return getNSCountCommandClass()
}

type NSCountCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCountCommandClass) Alloc() NSCountCommand {
	rv := objc.Send[NSCountCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A command that counts the number of objects of a specified class in the
// specified object container.
//
// # Overview
// 
// An instance of [NSCountCommand] counts the number of objects of a specified
// class in the specified object container (such as the number of words in a
// paragraph or document) and returns the result.
// 
// [NSCountCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `count` command through key-value coding. Most
// applications don’t need to subclass [NSCountCommand] or call its methods.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountCommand
type NSCountCommand struct {
	NSScriptCommand
}

// NSCountCommandFromID constructs a [NSCountCommand] from an objc.ID.
//
// A command that counts the number of objects of a specified class in the
// specified object container.
func NSCountCommandFromID(id objc.ID) NSCountCommand {
	return NSCountCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSCountCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCountCommand] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountCommand
type INSCountCommand interface {
	INSScriptCommand
}





// Init initializes the instance.
func (c NSCountCommand) Init() NSCountCommand {
	rv := objc.Send[NSCountCommand](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCountCommand) Autorelease() NSCountCommand {
	rv := objc.Send[NSCountCommand](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCountCommand creates a new NSCountCommand instance.
func NewNSCountCommand() NSCountCommand {
	class := getNSCountCommandClass()
	rv := objc.Send[NSCountCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewCountCommandWithCoder(inCoder INSCoder) NSCountCommand {
	instance := getNSCountCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSCountCommandFromID(rv)
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
func NewCountCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSCountCommand {
	instance := getNSCountCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSCountCommandFromID(rv)
}









































