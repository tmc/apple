// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMoveCommand] class.
var (
	_NSMoveCommandClass     NSMoveCommandClass
	_NSMoveCommandClassOnce sync.Once
)

func getNSMoveCommandClass() NSMoveCommandClass {
	_NSMoveCommandClassOnce.Do(func() {
		_NSMoveCommandClass = NSMoveCommandClass{class: objc.GetClass("NSMoveCommand")}
	})
	return _NSMoveCommandClass
}

// GetNSMoveCommandClass returns the class object for NSMoveCommand.
func GetNSMoveCommandClass() NSMoveCommandClass {
	return getNSMoveCommandClass()
}

type NSMoveCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMoveCommandClass) Alloc() NSMoveCommand {
	rv := objc.Send[NSMoveCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that moves one or more scriptable objects.
//
// # Overview
// 
// An instance of [NSMoveCommand] moves the specified scriptable object or
// objects; for example, it may move words to a new location in a document or
// a file to a new directory.
// 
// [NSMoveCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `move` AppleScript command through key-value
// coding. Most applications don’t need to subclass [NSMoveCommand] or
// invoke its methods. However, for circumstances where you might choose to
// subclass this command, see “Modifying a Standard Command” in [Script
// Commands] in [Cocoa Scripting Guide].
// 
// When an instance of [NSMoveCommand] is executed, it does not make copies of
// moved objects. It removes objects from the source container or containers,
// then inserts them into the destination container.
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Script Commands]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_script_cmds/SAppsScriptCmds.html#//apple_ref/doc/uid/20001242
//
// # Working with specifiers
//
//   - [NSMoveCommand.KeySpecifier]: Returns a specifier for the object or objects to be moved.
//
// See: https://developer.apple.com/documentation/Foundation/NSMoveCommand
type NSMoveCommand struct {
	NSScriptCommand
}

// NSMoveCommandFromID constructs a [NSMoveCommand] from an objc.ID.
//
// A command that moves one or more scriptable objects.
func NSMoveCommandFromID(id objc.ID) NSMoveCommand {
	return NSMoveCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSMoveCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMoveCommand] class.
//
// # Working with specifiers
//
//   - [INSMoveCommand.KeySpecifier]: Returns a specifier for the object or objects to be moved.
//
// See: https://developer.apple.com/documentation/Foundation/NSMoveCommand
type INSMoveCommand interface {
	INSScriptCommand

	// Topic: Working with specifiers

	// Returns a specifier for the object or objects to be moved.
	KeySpecifier() INSScriptObjectSpecifier
}

// Init initializes the instance.
func (m NSMoveCommand) Init() NSMoveCommand {
	rv := objc.Send[NSMoveCommand](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMoveCommand) Autorelease() NSMoveCommand {
	rv := objc.Send[NSMoveCommand](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMoveCommand creates a new NSMoveCommand instance.
func NewNSMoveCommand() NSMoveCommand {
	class := getNSMoveCommandClass()
	rv := objc.Send[NSMoveCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewMoveCommandWithCoder(inCoder INSCoder) NSMoveCommand {
	instance := getNSMoveCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSMoveCommandFromID(rv)
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
func NewMoveCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSMoveCommand {
	instance := getNSMoveCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSMoveCommandFromID(rv)
}

// Returns a specifier for the object or objects to be moved.
//
// # Return Value
// 
// A specifier for the object or objects to be moved.
// 
// # Discussion
// 
// Note that this specifier may be different than the specifier set by
// [SetReceiversSpecifier], which sets the container specifier. For example,
// for a command such as `move the third circle to the location of the first
// circle`, the receiver might identify a document (which has a list of
// graphics), while the key specifier identifies the particular graphic to be
// moved.
//
// See: https://developer.apple.com/documentation/Foundation/NSMoveCommand/keySpecifier
func (m NSMoveCommand) KeySpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keySpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}

