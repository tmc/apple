// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptCommandDescription] class.
var (
	_NSScriptCommandDescriptionClass     NSScriptCommandDescriptionClass
	_NSScriptCommandDescriptionClassOnce sync.Once
)

func getNSScriptCommandDescriptionClass() NSScriptCommandDescriptionClass {
	_NSScriptCommandDescriptionClassOnce.Do(func() {
		_NSScriptCommandDescriptionClass = NSScriptCommandDescriptionClass{class: objc.GetClass("NSScriptCommandDescription")}
	})
	return _NSScriptCommandDescriptionClass
}

// GetNSScriptCommandDescriptionClass returns the class object for NSScriptCommandDescription.
func GetNSScriptCommandDescriptionClass() NSScriptCommandDescriptionClass {
	return getNSScriptCommandDescriptionClass()
}

type NSScriptCommandDescriptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptCommandDescriptionClass) Alloc() NSScriptCommandDescription {
	rv := objc.Send[NSScriptCommandDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A script command that a macOS app supports.
//
// # Overview
// 
// A scriptable application provides scriptability information that describes
// the commands and objects scripters can use in scripts that target the
// application. An application’s scripting information is collected
// automatically by an instance of [NSScriptSuiteRegistry], which creates an
// [NSScriptCommandDescription] for each command it finds, caches these
// objects in memory, and installs a command handler for each command.
// 
// A script command instance stores the name, class, argument types, and
// return type of a command. For example, commands in AppleScript’s Core
// suite include `clone`, `count`, `create`, `delete`, `exists`, and `move`.
// 
// The public methods of [NSScriptCommandDescription] are used primarily by
// Cocoa’s built-in scripting support in responding to Apple events that
// target the application. Although you can subclass the
// [NSScriptCommandDescription] class, it is unlikely that you would need to
// do so, or to create instances of it.
//
// # Initializing a Script Command Description
//
//   - [NSScriptCommandDescription.InitWithSuiteNameCommandNameDictionary]: Initializes and returns a newly allocated instance of [NSScriptCommandDescription].
//
// # Getting Basic Information About the Command
//
//   - [NSScriptCommandDescription.AppleEventClassCode]: Returns the four-character code for the Apple event class of the receiver’s command.
//   - [NSScriptCommandDescription.AppleEventCode]: Returns the four-character code for the Apple event ID of the receiver’s command.
//   - [NSScriptCommandDescription.CommandClassName]: Returns the name of the class that will be instantiated to handle the command.
//   - [NSScriptCommandDescription.CommandName]: Returns the name of the command.
//   - [NSScriptCommandDescription.SuiteName]: Returns the name of the suite that contains the command described by the receiver.
//
// # Getting Command Argument Information
//
//   - [NSScriptCommandDescription.AppleEventCodeForArgumentWithName]: Returns the Apple event code for the specified command argument of the receiver.
//   - [NSScriptCommandDescription.ArgumentNames]: Returns the names (or keys) for all arguments of the receiver’s command.
//   - [NSScriptCommandDescription.IsOptionalArgumentWithName]: Returns a Boolean value that indicates whether the command argument identified by the specified argument key is an optional argument.
//   - [NSScriptCommandDescription.TypeForArgumentWithName]: Returns the type of the command argument identified by the specified key.
//
// # Getting Command Return-Type Information
//
//   - [NSScriptCommandDescription.AppleEventCodeForReturnType]: Returns the Apple event code that identifies the command’s return type.
//   - [NSScriptCommandDescription.ReturnType]: Returns the return type of the command.
//
// # Creating Commands
//
//   - [NSScriptCommandDescription.CreateCommandInstance]: Creates and returns an instance of the command object described by the receiver.
//   - [NSScriptCommandDescription.CreateCommandInstanceWithZone]: Creates and returns an instance of the command object described by the receiver in the specified memory zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription
type NSScriptCommandDescription struct {
	objectivec.Object
}

// NSScriptCommandDescriptionFromID constructs a [NSScriptCommandDescription] from an objc.ID.
//
// A script command that a macOS app supports.
func NSScriptCommandDescriptionFromID(id objc.ID) NSScriptCommandDescription {
	return NSScriptCommandDescription{objectivec.Object{ID: id}}
}
// NOTE: NSScriptCommandDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptCommandDescription] class.
//
// # Initializing a Script Command Description
//
//   - [INSScriptCommandDescription.InitWithSuiteNameCommandNameDictionary]: Initializes and returns a newly allocated instance of [NSScriptCommandDescription].
//
// # Getting Basic Information About the Command
//
//   - [INSScriptCommandDescription.AppleEventClassCode]: Returns the four-character code for the Apple event class of the receiver’s command.
//   - [INSScriptCommandDescription.AppleEventCode]: Returns the four-character code for the Apple event ID of the receiver’s command.
//   - [INSScriptCommandDescription.CommandClassName]: Returns the name of the class that will be instantiated to handle the command.
//   - [INSScriptCommandDescription.CommandName]: Returns the name of the command.
//   - [INSScriptCommandDescription.SuiteName]: Returns the name of the suite that contains the command described by the receiver.
//
// # Getting Command Argument Information
//
//   - [INSScriptCommandDescription.AppleEventCodeForArgumentWithName]: Returns the Apple event code for the specified command argument of the receiver.
//   - [INSScriptCommandDescription.ArgumentNames]: Returns the names (or keys) for all arguments of the receiver’s command.
//   - [INSScriptCommandDescription.IsOptionalArgumentWithName]: Returns a Boolean value that indicates whether the command argument identified by the specified argument key is an optional argument.
//   - [INSScriptCommandDescription.TypeForArgumentWithName]: Returns the type of the command argument identified by the specified key.
//
// # Getting Command Return-Type Information
//
//   - [INSScriptCommandDescription.AppleEventCodeForReturnType]: Returns the Apple event code that identifies the command’s return type.
//   - [INSScriptCommandDescription.ReturnType]: Returns the return type of the command.
//
// # Creating Commands
//
//   - [INSScriptCommandDescription.CreateCommandInstance]: Creates and returns an instance of the command object described by the receiver.
//   - [INSScriptCommandDescription.CreateCommandInstanceWithZone]: Creates and returns an instance of the command object described by the receiver in the specified memory zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription
type INSScriptCommandDescription interface {
	objectivec.IObject
	NSCoding

	// Topic: Initializing a Script Command Description

	// Initializes and returns a newly allocated instance of [NSScriptCommandDescription].
	InitWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration INSDictionary) NSScriptCommandDescription

	// Topic: Getting Basic Information About the Command

	// Returns the four-character code for the Apple event class of the receiver’s command.
	AppleEventClassCode() uint32
	// Returns the four-character code for the Apple event ID of the receiver’s command.
	AppleEventCode() uint32
	// Returns the name of the class that will be instantiated to handle the command.
	CommandClassName() string
	// Returns the name of the command.
	CommandName() string
	// Returns the name of the suite that contains the command described by the receiver.
	SuiteName() string

	// Topic: Getting Command Argument Information

	// Returns the Apple event code for the specified command argument of the receiver.
	AppleEventCodeForArgumentWithName(argumentName string) uint32
	// Returns the names (or keys) for all arguments of the receiver’s command.
	ArgumentNames() []string
	// Returns a Boolean value that indicates whether the command argument identified by the specified argument key is an optional argument.
	IsOptionalArgumentWithName(argumentName string) bool
	// Returns the type of the command argument identified by the specified key.
	TypeForArgumentWithName(argumentName string) string

	// Topic: Getting Command Return-Type Information

	// Returns the Apple event code that identifies the command’s return type.
	AppleEventCodeForReturnType() uint32
	// Returns the return type of the command.
	ReturnType() string

	// Topic: Creating Commands

	// Creates and returns an instance of the command object described by the receiver.
	CreateCommandInstance() INSScriptCommand
	// Creates and returns an instance of the command object described by the receiver in the specified memory zone.
	CreateCommandInstanceWithZone(zone NSZone) INSScriptCommand
}

// Init initializes the instance.
func (s NSScriptCommandDescription) Init() NSScriptCommandDescription {
	rv := objc.Send[NSScriptCommandDescription](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptCommandDescription) Autorelease() NSScriptCommandDescription {
	rv := objc.Send[NSScriptCommandDescription](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptCommandDescription creates a new NSScriptCommandDescription instance.
func NewNSScriptCommandDescription() NSScriptCommandDescription {
	class := getNSScriptCommandDescriptionClass()
	rv := objc.Send[NSScriptCommandDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(coder:)
func NewScriptCommandDescriptionWithCoder(inCoder INSCoder) NSScriptCommandDescription {
	instance := getNSScriptCommandDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSScriptCommandDescriptionFromID(rv)
}

// Initializes and returns a newly allocated instance of
// [NSScriptCommandDescription].
//
// suiteName: The name of the suite (in the application’s scriptability information)
// that the command belongs to. For example, `"AppName Suite"`.
//
// commandName: The name of the script command that this instance describes.
//
// commandDeclaration: A command declaration dictionary of the sort that is valid in script suite
// property list files. This dictionary provides information about the command
// such as its argument names and types and return type (if any).
//
// # Return Value
// 
// The initialized command description instance. Returns `nil` if the event
// constant or class name for the command description is missing; also returns
// `nil` if the return type or argument values are of the wrong type.
//
// # Discussion
// 
// This method registers `self` with the application’s global instance of
// [NSScriptSuiteRegistry] and also registers all command arguments with the
// registry.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(suiteName:commandName:dictionary:)
func NewScriptCommandDescriptionWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration INSDictionary) NSScriptCommandDescription {
	instance := getNSScriptCommandDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSuiteName:commandName:dictionary:"), objc.String(suiteName), objc.String(commandName), commandDeclaration)
	return NSScriptCommandDescriptionFromID(rv)
}

// Initializes and returns a newly allocated instance of
// [NSScriptCommandDescription].
//
// suiteName: The name of the suite (in the application’s scriptability information)
// that the command belongs to. For example, `"AppName Suite"`.
//
// commandName: The name of the script command that this instance describes.
//
// commandDeclaration: A command declaration dictionary of the sort that is valid in script suite
// property list files. This dictionary provides information about the command
// such as its argument names and types and return type (if any).
//
// # Return Value
// 
// The initialized command description instance. Returns `nil` if the event
// constant or class name for the command description is missing; also returns
// `nil` if the return type or argument values are of the wrong type.
//
// # Discussion
// 
// This method registers `self` with the application’s global instance of
// [NSScriptSuiteRegistry] and also registers all command arguments with the
// registry.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(suiteName:commandName:dictionary:)
func (s NSScriptCommandDescription) InitWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration INSDictionary) NSScriptCommandDescription {
	rv := objc.Send[NSScriptCommandDescription](s.ID, objc.Sel("initWithSuiteName:commandName:dictionary:"), objc.String(suiteName), objc.String(commandName), commandDeclaration)
	return rv
}

// Returns the Apple event code for the specified command argument of the
// receiver.
//
// argumentName: The argument name (used as a key) for which to obtain the corresponding
// Apple event code.
//
// # Return Value
// 
// The code for the specified argument.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCodeForArgument(withName:)
func (s NSScriptCommandDescription) AppleEventCodeForArgumentWithName(argumentName string) uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCodeForArgumentWithName:"), objc.String(argumentName))
	return rv
}

// Returns a Boolean value that indicates whether the command argument
// identified by the specified argument key is an optional argument.
//
// argumentName: Argument name (used as a key) that identifies the command argument to
// examine.
//
// # Return Value
// 
// [true] if the specified argument exists and is optional; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/isOptionalArgument(withName:)
func (s NSScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isOptionalArgumentWithName:"), objc.String(argumentName))
	return rv
}

// Returns the type of the command argument identified by the specified key.
//
// argumentName: Argument name (used as a key) that identifies the command argument to
// examine.
//
// # Return Value
// 
// The type of the specified command argument. Returns `nil` if there is no
// such argument.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/typeForArgument(withName:)
func (s NSScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("typeForArgumentWithName:"), objc.String(argumentName))
	return NSStringFromID(rv).String()
}

// Creates and returns an instance of the command object described by the
// receiver.
//
// # Return Value
// 
// The command object, instantiated from [NSScriptCommand] or a subclass.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/createCommandInstance()
func (s NSScriptCommandDescription) CreateCommandInstance() INSScriptCommand {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createCommandInstance"))
	return NSScriptCommandFromID(rv)
}

// Creates and returns an instance of the command object described by the
// receiver in the specified memory zone.
//
// zone: The memory zone from which to allocate the command.
//
// # Return Value
// 
// The command object, instantiated from [NSScriptCommand] or a subclass.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/createCommandInstance(with:)
func (s NSScriptCommandDescription) CreateCommandInstanceWithZone(zone NSZone) INSScriptCommand {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createCommandInstanceWithZone:"), zone)
	return NSScriptCommandFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(coder:)
func (s NSScriptCommandDescription) InitWithCoder(inCoder INSCoder) NSScriptCommandDescription {
	rv := objc.Send[NSScriptCommandDescription](s.ID, objc.Sel("initWithCoder:"), inCoder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSScriptCommandDescription) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the four-character code for the Apple event class of the
// receiver’s command.
//
// # Return Value
// 
// The Apple event code associated with the receiver’s command. This is the
// primary code used to identify the command in Apple events.
// 
// # Discussion
// 
// In an Apple event that specifies a script command, two four character
// codes—the event class and event ID—together identify the command. You
// use this method to obtain the event class. You use [AppleEventCode] to
// obtain the event ID.
// 
// For example, commands in AppleScript’s Core suite, such as `clone`,
// `count`, and `create`, have an event class code of `'core'`. This code and
// the event ID code returned by `appleEventCode` together specify the
// necessary information for identifying and dispatching an Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventClassCode
func (s NSScriptCommandDescription) AppleEventClassCode() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventClassCode"))
	return rv
}

// Returns the four-character code for the Apple event ID of the receiver’s
// command.
//
// # Return Value
// 
// The code for the event ID of the receiver’s command.
// 
// # Discussion
// 
// This value of the event ID returned by this method, together with the event
// class code returned by [AppleEventClassCode], specifies the necessary
// information for identifying and dispatching an Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCode
func (s NSScriptCommandDescription) AppleEventCode() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCode"))
	return rv
}

// Returns the name of the class that will be instantiated to handle the
// command.
//
// # Return Value
// 
// The Objective-C class name (for example, `"NSGetCommand"`). This is always
// [NSScriptCommand] or a subclass.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/commandClassName
func (s NSScriptCommandDescription) CommandClassName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandClassName"))
	return NSStringFromID(rv).String()
}

// Returns the name of the command.
//
// # Return Value
// 
// The command name as it appears in the application’s scriptability
// information; may be different from what is displayed to the scripter.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/commandName
func (s NSScriptCommandDescription) CommandName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandName"))
	return NSStringFromID(rv).String()
}

// Returns the name of the suite that contains the command described by the
// receiver.
//
// # Return Value
// 
// The receiver’s suite name. Within an application’s scriptability
// information, named suites contain related sets of information.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/suiteName
func (s NSScriptCommandDescription) SuiteName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("suiteName"))
	return NSStringFromID(rv).String()
}

// Returns the names (or keys) for all arguments of the receiver’s command.
//
// # Return Value
// 
// The array of argument names. If there are no arguments for the command,
// returns an empty array.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/argumentNames
func (s NSScriptCommandDescription) ArgumentNames() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("argumentNames"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the Apple event code that identifies the command’s return type.
//
// # Return Value
// 
// The event code for the command’s return type.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCodeForReturnType
func (s NSScriptCommandDescription) AppleEventCodeForReturnType() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCodeForReturnType"))
	return rv
}

// Returns the return type of the command.
//
// # Return Value
// 
// The receiver’s command return type; for example, `"NSNumber"` or
// `"NSDictionary"`).
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/returnType
func (s NSScriptCommandDescription) ReturnType() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("returnType"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCoding
			

