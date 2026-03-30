// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCloneCommand] class.
var (
	_NSCloneCommandClass     NSCloneCommandClass
	_NSCloneCommandClassOnce sync.Once
)

func getNSCloneCommandClass() NSCloneCommandClass {
	_NSCloneCommandClassOnce.Do(func() {
		_NSCloneCommandClass = NSCloneCommandClass{class: objc.GetClass("NSCloneCommand")}
	})
	return _NSCloneCommandClass
}

// GetNSCloneCommandClass returns the class object for NSCloneCommand.
func GetNSCloneCommandClass() NSCloneCommandClass {
	return getNSCloneCommandClass()
}

type NSCloneCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCloneCommandClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCloneCommandClass) Alloc() NSCloneCommand {
	rv := objc.Send[NSCloneCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that clones one or more scriptable objects.
//
// # Overview
//
// An instance of [NSCloneCommand] clones the specified scriptable object or
// objects (such as words, paragraphs, images, and so on) and inserts them in
// the specified location, or the default location if no location is
// specified. The cloned scriptable objects typically correspond to objects in
// the application, but aren’t required to. This command corresponds to
// AppleScript’s `duplicate` command.
//
// [NSCloneCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `duplicate` command through key-value coding.
// Most applications don’t need to subclass [NSCloneCommand] or invoke its
// methods.
//
// When an instance of [NSCloneCommand] is executed, it clones the specified
// objects by sending them [copyWithZone:] messages.
//
// # Working with specifiers
//
//   - [NSCloneCommand.KeySpecifier]: Returns a specifier for the object or objects to be cloned.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloneCommand
//
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
type NSCloneCommand struct {
	NSScriptCommand
}

// NSCloneCommandFromID constructs a [NSCloneCommand] from an objc.ID.
//
// A command that clones one or more scriptable objects.
func NSCloneCommandFromID(id objc.ID) NSCloneCommand {
	return NSCloneCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}

// NOTE: NSCloneCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCloneCommand] class.
//
// # Working with specifiers
//
//   - [INSCloneCommand.KeySpecifier]: Returns a specifier for the object or objects to be cloned.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloneCommand
type INSCloneCommand interface {
	INSScriptCommand

	// Topic: Working with specifiers

	// Returns a specifier for the object or objects to be cloned.
	KeySpecifier() INSScriptObjectSpecifier
}

// Init initializes the instance.
func (c NSCloneCommand) Init() NSCloneCommand {
	rv := objc.Send[NSCloneCommand](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCloneCommand) Autorelease() NSCloneCommand {
	rv := objc.Send[NSCloneCommand](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCloneCommand creates a new NSCloneCommand instance.
func NewNSCloneCommand() NSCloneCommand {
	class := getNSCloneCommandClass()
	rv := objc.Send[NSCloneCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewCloneCommandWithCoder(inCoder INSCoder) NSCloneCommand {
	instance := getNSCloneCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSCloneCommandFromID(rv)
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
func NewCloneCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSCloneCommand {
	instance := getNSCloneCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSCloneCommandFromID(rv)
}

// Returns a specifier for the object or objects to be cloned.
//
// # Return Value
//
// A specifier for the object or objects to be cloned.
//
// # Discussion
//
// For example, the specifier may indicate that a document’s third rectangle
// should be cloned. The returned specifier is valid only in the context of
// the [NSCloneCommand] object; for example, if you send the specifier a
// [ContainerSpecifier] message, the result is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCloneCommand/keySpecifier
func (c NSCloneCommand) KeySpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keySpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
