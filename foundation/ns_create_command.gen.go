// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCreateCommand] class.
var (
	_NSCreateCommandClass     NSCreateCommandClass
	_NSCreateCommandClassOnce sync.Once
)

func getNSCreateCommandClass() NSCreateCommandClass {
	_NSCreateCommandClassOnce.Do(func() {
		_NSCreateCommandClass = NSCreateCommandClass{class: objc.GetClass("NSCreateCommand")}
	})
	return _NSCreateCommandClass
}

// GetNSCreateCommandClass returns the class object for NSCreateCommand.
func GetNSCreateCommandClass() NSCreateCommandClass {
	return getNSCreateCommandClass()
}

type NSCreateCommandClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCreateCommandClass) Alloc() NSCreateCommand {
	rv := objc.Send[NSCreateCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that creates a scriptable object.
//
// # Overview
// 
// An instance of [NSCreateCommand] creates the specified scriptable object
// (such as a document), optionally supplying the new object with the
// specified attributes. This command corresponds to AppleScript’s `make`
// command.
// 
// [NSCreateCommand] is part of Cocoa’s built-in scripting support. Most
// applications don’t need to subclass [NSCreateCommand] or invoke its
// methods.
// 
// When an instance of [NSCreateCommand] is executed, it creates a new object
// using `[[theClassToBeCreated NULL] init]` (where `theClassToBeCreated` is
// the class of the object to be created), unless the command has a `with
// data` argument. In the latter case, the new object is created by invoking
// `[[NSScriptCoercionHandler sharedCoercionHandler] theDataAsAnObject
// theClassToBeCreated]`. Any properties specified by a `with properties`
// argument are then set in the new object using `-`.
// 
// If an [NSCreateCommand] object with no argument corresponding to the `at`
// parameter is executed (for example, `tell application "Mail" to make new
// mailbox with properties {"testFolder"}`), and the receiver of the command
// (not necessarily the application object) has a to-many relationship to
// objects of the class to be instantiated, and the class description for the
// receiving class returns [false] when sent an `` message, the
// [NSCreateCommand] object creates a new object and sends the receiver an
// [insertValue(_:at:inPropertyWithKey:)] message to place the new object in
// the container. This is part of Cocoa’s scripting support for inserting
// newly-created objects into containers without explicitly specifying a
// location.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [insertValue(_:at:inPropertyWithKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/insertValue(_:at:inPropertyWithKey:)
//
// # Getting information about a create command
//
//   - [NSCreateCommand.CreateClassDescription]: Returns the class description for the class that is to be created.
//   - [NSCreateCommand.ResolvedKeyDictionary]: Returns a dictionary that contains the properties that were specified in the `make` Apple event command that has been converted to this [NSCreateCommand] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateCommand
type NSCreateCommand struct {
	NSScriptCommand
}

// NSCreateCommandFromID constructs a [NSCreateCommand] from an objc.ID.
//
// A command that creates a scriptable object.
func NSCreateCommandFromID(id objc.ID) NSCreateCommand {
	return NSCreateCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}
// NOTE: NSCreateCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCreateCommand] class.
//
// # Getting information about a create command
//
//   - [INSCreateCommand.CreateClassDescription]: Returns the class description for the class that is to be created.
//   - [INSCreateCommand.ResolvedKeyDictionary]: Returns a dictionary that contains the properties that were specified in the `make` Apple event command that has been converted to this [NSCreateCommand] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateCommand
type INSCreateCommand interface {
	INSScriptCommand

	// Topic: Getting information about a create command

	// Returns the class description for the class that is to be created.
	CreateClassDescription() INSScriptClassDescription
	// Returns a dictionary that contains the properties that were specified in the `make` Apple event command that has been converted to this [NSCreateCommand] object.
	ResolvedKeyDictionary() INSDictionary
}

// Init initializes the instance.
func (c NSCreateCommand) Init() NSCreateCommand {
	rv := objc.Send[NSCreateCommand](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCreateCommand) Autorelease() NSCreateCommand {
	rv := objc.Send[NSCreateCommand](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCreateCommand creates a new NSCreateCommand instance.
func NewNSCreateCommand() NSCreateCommand {
	class := getNSCreateCommandClass()
	rv := objc.Send[NSCreateCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewCreateCommandWithCoder(inCoder INSCoder) NSCreateCommand {
	instance := getNSCreateCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSCreateCommandFromID(rv)
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
func NewCreateCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSCreateCommand {
	instance := getNSCreateCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSCreateCommandFromID(rv)
}

// Returns the class description for the class that is to be created.
//
// # Return Value
// 
// The class description for the class that is to be created.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateCommand/createClassDescription
func (c NSCreateCommand) CreateClassDescription() INSScriptClassDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("createClassDescription"))
	return NSScriptClassDescriptionFromID(objc.ID(rv))
}
// Returns a dictionary that contains the properties that were specified in
// the `make` Apple event command that has been converted to this
// [NSCreateCommand] object.
//
// # Return Value
// 
// A dictionary that contains the properties that were specified in the `make`
// Apple event script command that has been converted to this
// [NSCreateCommand] object.
// 
// # Discussion
// 
// The keys in the returned dictionary are the names of properties (attributes
// or relationships, in the script suite) that have been specified for the
// command, and the corresponding values in the dictionary are the values that
// those properties should take. The required and optional arguments for the
// `make` command are specified in the core suite definition,
// `NSCoreSuite.ScriptSuite()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateCommand/resolvedKeyDictionary
func (c NSCreateCommand) ResolvedKeyDictionary() INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resolvedKeyDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}

