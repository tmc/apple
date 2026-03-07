// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSDeleteCommand] class.
var (
	_NSDeleteCommandClass     NSDeleteCommandClass
	_NSDeleteCommandClassOnce sync.Once
)

func getNSDeleteCommandClass() NSDeleteCommandClass {
	_NSDeleteCommandClassOnce.Do(func() {
		_NSDeleteCommandClass = NSDeleteCommandClass{class: objc.GetClass("NSDeleteCommand")}
	})
	return _NSDeleteCommandClass
}

// GetNSDeleteCommandClass returns the class object for NSDeleteCommand.
func GetNSDeleteCommandClass() NSDeleteCommandClass {
	return getNSDeleteCommandClass()
}

type NSDeleteCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDeleteCommandClass) Alloc() NSDeleteCommand {
	rv := objc.Send[NSDeleteCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A command that deletes a scriptable object.
//
// # Overview
// 
// An instance of [NSDeleteCommand] deletes the specified scriptable object or
// objects (such as words, paragraphs, and so on).
// 
// Suppose, for example, a user executes a script that sends the command
// `delete the third rectangle in the first document` to the Sketch sample
// application (located in `/Developer/Examples/AppKit`). Cocoa creates an
// [NSDeleteCommand] object to perform the operation. When the command is
// executed, it uses the key-value coding mechanism (by invoking ``) to remove
// the specified object or objects from their container. See the description
// for [removeValue(at:fromPropertyWithKey:)] for related information.
// 
// [NSDeleteCommand] is part of Cocoa’s built-in scripting support. Most
// applications don’t need to subclass [NSDeleteCommand] or call its
// methods.
//
// [removeValue(at:fromPropertyWithKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/removeValue(at:fromPropertyWithKey:)
//
// # Working with specifiers
//
//   - [NSDeleteCommand.KeySpecifier]: Returns a specifier for the object or objects to be deleted.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeleteCommand
type NSDeleteCommand struct {
	NSScriptCommand
}

// NSDeleteCommandFromID constructs a [NSDeleteCommand] from an objc.ID.
//
// A command that deletes a scriptable object.
func NSDeleteCommandFromID(id objc.ID) NSDeleteCommand {
	return NSDeleteCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSDeleteCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDeleteCommand] class.
//
// # Working with specifiers
//
//   - [INSDeleteCommand.KeySpecifier]: Returns a specifier for the object or objects to be deleted.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeleteCommand
type INSDeleteCommand interface {
	INSScriptCommand

	// Topic: Working with specifiers

	// Returns a specifier for the object or objects to be deleted.
	KeySpecifier() INSScriptObjectSpecifier
}





// Init initializes the instance.
func (d NSDeleteCommand) Init() NSDeleteCommand {
	rv := objc.Send[NSDeleteCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDeleteCommand) Autorelease() NSDeleteCommand {
	rv := objc.Send[NSDeleteCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDeleteCommand creates a new NSDeleteCommand instance.
func NewNSDeleteCommand() NSDeleteCommand {
	class := getNSDeleteCommandClass()
	rv := objc.Send[NSDeleteCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewDeleteCommandWithCoder(inCoder INSCoder) NSDeleteCommand {
	instance := getNSDeleteCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSDeleteCommandFromID(rv)
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
func NewDeleteCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSDeleteCommand {
	instance := getNSDeleteCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSDeleteCommandFromID(rv)
}


















// Returns a specifier for the object or objects to be deleted.
//
// # Return Value
// 
// A specifier for the object or objects to be deleted.
// 
// # Discussion
// 
// Note that this may be different than the specifier or specifiers set by
// [SetReceiversSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSDeleteCommand/keySpecifier
func (d NSDeleteCommand) KeySpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("keySpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}


























