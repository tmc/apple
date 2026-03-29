// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptCommand] class.
var (
	_NSScriptCommandClass     NSScriptCommandClass
	_NSScriptCommandClassOnce sync.Once
)

func getNSScriptCommandClass() NSScriptCommandClass {
	_NSScriptCommandClassOnce.Do(func() {
		_NSScriptCommandClass = NSScriptCommandClass{class: objc.GetClass("NSScriptCommand")}
	})
	return _NSScriptCommandClass
}

// GetNSScriptCommandClass returns the class object for NSScriptCommand.
func GetNSScriptCommandClass() NSScriptCommandClass {
	return getNSScriptCommandClass()
}

type NSScriptCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScriptCommandClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptCommandClass) Alloc() NSScriptCommand {
	rv := objc.Send[NSScriptCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A self-contained scripting statement.
//
// # Overview
// 
// An instance of [NSScriptCommand] represents a scripting statement, such as
// `set word 5 of the front document to word 1 of the second document`, and
// contains the information needed to perform the operation specified by the
// statement.
// 
// When an Apple event reaches a Cocoa application, Cocoa’s built-in
// scripting support transforms it into a script command (that is, an instance
// of [NSScriptCommand] or one of the subclasses provided by Cocoa scripting
// or by your application) and executes the command in the context of the
// application. Executing a command means either invoking the selector
// associated with the command on the object or objects designated to receive
// the command, or having the command perform its default implementation
// method ([NSScriptCommand.PerformDefaultImplementation]).
// 
// Your application most likely calls methods of [NSScriptCommand] to extract
// the command arguments. You do this either in the
// `performDefaultImplementation` method of a command subclass you have
// created, or in an object method designated as the selector to handle a
// particular command.
// 
// As part of Cocoa’s standard scripting implementation, [NSScriptCommand]
// and its subclasses can handle the default command set for AppleScript’s
// Standard suite for most applications without any subclassing. The Standard
// suite includes commands such as `copy`, `count`, `create`, `delete`,
// `exists`, and `move`, as well as common object classes such as
// `application`, `document`, and `window`.
// 
// For more information on working with script commands, see [Script Commands]
// in [Cocoa Scripting Guide].
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Script Commands]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_script_cmds/SAppsScriptCmds.html#//apple_ref/doc/uid/20001242
//
// # Initializing a script command
//
//   - [NSScriptCommand.InitWithCommandDescription]: Returns an a script command object initialized from the passed command description.
//
// # Getting the Apple event
//
//   - [NSScriptCommand.AppleEvent]: If the receiver was constructed by Cocoa scripting’s built-in Apple event handling, returns the Apple event descriptor from which it was constructed.
//
// # Executing the command
//
//   - [NSScriptCommand.ExecuteCommand]: Executes the command if it is valid and returns the result, if any.
//   - [NSScriptCommand.PerformDefaultImplementation]: Overridden by subclasses to provide a default implementation for the command represented by the receiver.
//
// # Accessing receivers
//
//   - [NSScriptCommand.EvaluatedReceivers]: Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
//   - [NSScriptCommand.ReceiversSpecifier]: Sets the object specifier to `receiversSpec` that, when evaluated, indicates the receiver or receivers of the command.
//   - [NSScriptCommand.SetReceiversSpecifier]
//
// # Accessing arguments
//
//   - [NSScriptCommand.Arguments]: Sets the arguments of the command to `args`.
//   - [NSScriptCommand.SetArguments]
//   - [NSScriptCommand.EvaluatedArguments]: Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names.
//
// # Accessing the direct parameter
//
//   - [NSScriptCommand.DirectParameter]: Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives.
//   - [NSScriptCommand.SetDirectParameter]
//
// # Getting command information
//
//   - [NSScriptCommand.CommandDescription]: Returns the command description for the command.
//   - [NSScriptCommand.WellFormed]: Returns a Boolean value indicating whether the receiver is well formed according to its command description.
//
// # Handling script execution errors
//
//   - [NSScriptCommand.ScriptErrorExpectedTypeDescriptor]: Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//   - [NSScriptCommand.SetScriptErrorExpectedTypeDescriptor]
//   - [NSScriptCommand.ScriptErrorNumber]: Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender.
//   - [NSScriptCommand.SetScriptErrorNumber]
//   - [NSScriptCommand.ScriptErrorOffendingObjectDescriptor]: Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//   - [NSScriptCommand.SetScriptErrorOffendingObjectDescriptor]
//   - [NSScriptCommand.ScriptErrorString]: Sets a script error string that is associated with execution of the command.
//   - [NSScriptCommand.SetScriptErrorString]
//
// # Suspending and resuming commands
//
//   - [NSScriptCommand.SuspendExecution]: Suspends the execution of the receiver.
//   - [NSScriptCommand.ResumeExecutionWithResult]: If a successful, unmatched, invocation of [suspendExecution()](<doc://com.apple.foundation/documentation/Foundation/NSScriptCommand/suspendExecution()>) has been made, resume the execution of the command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand
type NSScriptCommand struct {
	objectivec.Object
}

// NSScriptCommandFromID constructs a [NSScriptCommand] from an objc.ID.
//
// A self-contained scripting statement.
func NSScriptCommandFromID(id objc.ID) NSScriptCommand {
	return NSScriptCommand{objectivec.Object{ID: id}}
}
// NOTE: NSScriptCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptCommand] class.
//
// # Initializing a script command
//
//   - [INSScriptCommand.InitWithCommandDescription]: Returns an a script command object initialized from the passed command description.
//
// # Getting the Apple event
//
//   - [INSScriptCommand.AppleEvent]: If the receiver was constructed by Cocoa scripting’s built-in Apple event handling, returns the Apple event descriptor from which it was constructed.
//
// # Executing the command
//
//   - [INSScriptCommand.ExecuteCommand]: Executes the command if it is valid and returns the result, if any.
//   - [INSScriptCommand.PerformDefaultImplementation]: Overridden by subclasses to provide a default implementation for the command represented by the receiver.
//
// # Accessing receivers
//
//   - [INSScriptCommand.EvaluatedReceivers]: Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
//   - [INSScriptCommand.ReceiversSpecifier]: Sets the object specifier to `receiversSpec` that, when evaluated, indicates the receiver or receivers of the command.
//   - [INSScriptCommand.SetReceiversSpecifier]
//
// # Accessing arguments
//
//   - [INSScriptCommand.Arguments]: Sets the arguments of the command to `args`.
//   - [INSScriptCommand.SetArguments]
//   - [INSScriptCommand.EvaluatedArguments]: Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names.
//
// # Accessing the direct parameter
//
//   - [INSScriptCommand.DirectParameter]: Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives.
//   - [INSScriptCommand.SetDirectParameter]
//
// # Getting command information
//
//   - [INSScriptCommand.CommandDescription]: Returns the command description for the command.
//   - [INSScriptCommand.WellFormed]: Returns a Boolean value indicating whether the receiver is well formed according to its command description.
//
// # Handling script execution errors
//
//   - [INSScriptCommand.ScriptErrorExpectedTypeDescriptor]: Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//   - [INSScriptCommand.SetScriptErrorExpectedTypeDescriptor]
//   - [INSScriptCommand.ScriptErrorNumber]: Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender.
//   - [INSScriptCommand.SetScriptErrorNumber]
//   - [INSScriptCommand.ScriptErrorOffendingObjectDescriptor]: Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//   - [INSScriptCommand.SetScriptErrorOffendingObjectDescriptor]
//   - [INSScriptCommand.ScriptErrorString]: Sets a script error string that is associated with execution of the command.
//   - [INSScriptCommand.SetScriptErrorString]
//
// # Suspending and resuming commands
//
//   - [INSScriptCommand.SuspendExecution]: Suspends the execution of the receiver.
//   - [INSScriptCommand.ResumeExecutionWithResult]: If a successful, unmatched, invocation of [suspendExecution()](<doc://com.apple.foundation/documentation/Foundation/NSScriptCommand/suspendExecution()>) has been made, resume the execution of the command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand
type INSScriptCommand interface {
	objectivec.IObject
	NSCoding

	// Topic: Initializing a script command

	// Returns an a script command object initialized from the passed command description.
	InitWithCommandDescription(commandDef INSScriptCommandDescription) NSScriptCommand

	// Topic: Getting the Apple event

	// If the receiver was constructed by Cocoa scripting’s built-in Apple event handling, returns the Apple event descriptor from which it was constructed.
	AppleEvent() INSAppleEventDescriptor

	// Topic: Executing the command

	// Executes the command if it is valid and returns the result, if any.
	ExecuteCommand() objectivec.IObject
	// Overridden by subclasses to provide a default implementation for the command represented by the receiver.
	PerformDefaultImplementation() objectivec.IObject

	// Topic: Accessing receivers

	// Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
	EvaluatedReceivers() objectivec.IObject
	// Sets the object specifier to `receiversSpec` that, when evaluated, indicates the receiver or receivers of the command.
	ReceiversSpecifier() INSScriptObjectSpecifier
	SetReceiversSpecifier(value INSScriptObjectSpecifier)

	// Topic: Accessing arguments

	// Sets the arguments of the command to `args`.
	Arguments() INSDictionary
	SetArguments(value INSDictionary)
	// Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names.
	EvaluatedArguments() INSDictionary

	// Topic: Accessing the direct parameter

	// Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives.
	DirectParameter() objectivec.IObject
	SetDirectParameter(value objectivec.IObject)

	// Topic: Getting command information

	// Returns the command description for the command.
	CommandDescription() INSScriptCommandDescription
	// Returns a Boolean value indicating whether the receiver is well formed according to its command description.
	WellFormed() bool

	// Topic: Handling script execution errors

	// Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
	ScriptErrorExpectedTypeDescriptor() INSAppleEventDescriptor
	SetScriptErrorExpectedTypeDescriptor(value INSAppleEventDescriptor)
	// Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender.
	ScriptErrorNumber() int
	SetScriptErrorNumber(value int)
	// Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
	ScriptErrorOffendingObjectDescriptor() INSAppleEventDescriptor
	SetScriptErrorOffendingObjectDescriptor(value INSAppleEventDescriptor)
	// Sets a script error string that is associated with execution of the command.
	ScriptErrorString() string
	SetScriptErrorString(value string)

	// Topic: Suspending and resuming commands

	// Suspends the execution of the receiver.
	SuspendExecution()
	// If a successful, unmatched, invocation of [suspendExecution()](<doc://com.apple.foundation/documentation/Foundation/NSScriptCommand/suspendExecution()>) has been made, resume the execution of the command.
	ResumeExecutionWithResult(result objectivec.IObject)
}

// Init initializes the instance.
func (s NSScriptCommand) Init() NSScriptCommand {
	rv := objc.Send[NSScriptCommand](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptCommand) Autorelease() NSScriptCommand {
	rv := objc.Send[NSScriptCommand](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptCommand creates a new NSScriptCommand instance.
func NewNSScriptCommand() NSScriptCommand {
	class := getNSScriptCommandClass()
	rv := objc.Send[NSScriptCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewScriptCommandWithCoder(inCoder INSCoder) NSScriptCommand {
	instance := getNSScriptCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSScriptCommandFromID(rv)
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
func NewScriptCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSScriptCommand {
	instance := getNSScriptCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSScriptCommandFromID(rv)
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
func (s NSScriptCommand) InitWithCommandDescription(commandDef INSScriptCommandDescription) NSScriptCommand {
	rv := objc.Send[NSScriptCommand](s.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return rv
}
// Executes the command if it is valid and returns the result, if any.
//
// # Discussion
// 
// Before this method executes the command (through [NSInvocation]
// mechanisms), it evaluates all object specifiers involved in the command,
// validates that the receivers can actually handle the command, and verifies
// that the types of any arguments that were initially object specifiers are
// valid.
// 
// You shouldn’t have to override this method. If the command’s receivers
// want to handle the command themselves, this method invokes their defined
// handler. Otherwise, it invokes [PerformDefaultImplementation].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/execute()
func (s NSScriptCommand) ExecuteCommand() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("executeCommand"))
	return objectivec.Object{ID: rv}
}
// Overridden by subclasses to provide a default implementation for the
// command represented by the receiver.
//
// # Discussion
// 
// Do not invoke this method directly. [ExecuteCommand] invokes this method
// when the command being executed is not supported by the class of the
// objects receiving the command. The default implementation returns `nil`.
// 
// You need to create a subclass of [NSScriptCommand] only if you need to
// provide a default implementation of a command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/performDefaultImplementation()
func (s NSScriptCommand) PerformDefaultImplementation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("performDefaultImplementation"))
	return objectivec.Object{ID: rv}
}
// Suspends the execution of the receiver.
//
// # Discussion
// 
// Suspends the execution of the receiver only if the receiver is being
// executed in the current thread by Cocoa scripting’s built-in Apple event
// handling (that is, the receiver would be returned by `[NSScriptCommand
// currentCommand]`)—otherwise, does nothing. A matching invocation of
// [ResumeExecutionWithResult] must be made.
// 
// Another command can execute while a command is suspended.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/suspendExecution()
func (s NSScriptCommand) SuspendExecution() {
	objc.Send[objc.ID](s.ID, objc.Sel("suspendExecution"))
}
// If a successful, unmatched, invocation of [SuspendExecution] has been made,
// resume the execution of the command.
//
// # Discussion
// 
// Resumes the execution of the command if a successful, unmatched, invocation
// of [SuspendExecution] has been made—otherwise, does nothing. The value
// for `result` is dependent on the segment of command execution that was
// suspended:
// 
// - If [SuspendExecution] was invoked from within a command handler of one of
// the command’s receivers, `result` is considered to be the return value of
// the handler. Unless the command has received a [ScriptErrorNumber] message
// with a nonzero error number, execution of the command will continue and the
// command handlers of other receivers will be invoked. - If
// [SuspendExecution] was invoked from within an override of
// [PerformDefaultImplementation] the result is treated as if it were the
// return value of the invocation of [PerformDefaultImplementation].
// 
// [ResumeExecutionWithResult] may be invoked in any thread, not just the one
// in which the corresponding invocation of [SuspendExecution] occurred.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/resumeExecution(withResult:)
func (s NSScriptCommand) ResumeExecutionWithResult(result objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("resumeExecutionWithResult:"), result)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func (s NSScriptCommand) InitWithCoder(inCoder INSCoder) NSScriptCommand {
	rv := objc.Send[NSScriptCommand](s.ID, objc.Sel("initWithCoder:"), inCoder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSScriptCommand) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// If a command is being executed in the current thread by Cocoa scripting’s
// built-in Apple event handling, return the command.
//
// # Discussion
// 
// A command is being executed in the current thread by Cocoa scripting’s
// built-in Apple event handling if an instance of [NSScriptCommand] is
// handling an [ExecuteCommand] message at this instant as the result of the
// dispatch of an Apple event. Returns `nil` otherwise. [ScriptErrorNumber]
// and [ScriptErrorString] messages sent to the returned command object will
// affect the reply event sent to the sender of the event from which the
// command was constructed, if the sender has requested a reply.
// 
// A suspended command is not considered the current command. If a command is
// suspended and no other command is being executed in the current thread,
// [CurrentCommand] returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/current()
func (_NSScriptCommandClass NSScriptCommandClass) CurrentCommand() NSScriptCommand {
	rv := objc.Send[objc.ID](objc.ID(_NSScriptCommandClass.class), objc.Sel("currentCommand"))
	return NSScriptCommandFromID(rv)
}

// If the receiver was constructed by Cocoa scripting’s built-in Apple event
// handling, returns the Apple event descriptor from which it was constructed.
//
// # Discussion
// 
// The effects of mutating or retaining this descriptor are undefined,
// although it may be copied.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/appleEvent
func (s NSScriptCommand) AppleEvent() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("appleEvent"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}
// Returns the object or objects to which the command is to be sent (called
// both the “receivers” or “targets” of script commands).
//
// # Discussion
// 
// It evaluates receivers, which are always object specifiers, to a proper
// object. If the command does not specify a receiver, or if the receiver
// doesn’t accept the command, it returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/evaluatedReceivers
func (s NSScriptCommand) EvaluatedReceivers() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("evaluatedReceivers"))
	return objectivec.Object{ID: rv}
}
// Sets the object specifier to `receiversSpec` that, when evaluated,
// indicates the receiver or receivers of the command.
//
// # Discussion
// 
// If you create a subclass of [NSScriptCommand], you don’t necessarily need
// to override this method, though some of Cocoa’s subclasses do. An
// override should perform the same function as the superclass method, with a
// critical difference: it causes the container specifier part of the
// passed-in object specifier to become the receiver specifier of the command,
// and the key part of the passed-in object specifier to become the key
// specifier. In an override, for example, if `receiversRef` is a specifier
// for `the third rectangle of the first document`, the receiver specifier is
// `the first document` while the key specifier is `the third rectangle`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/receiversSpecifier
func (s NSScriptCommand) ReceiversSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("receiversSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (s NSScriptCommand) SetReceiversSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](s.ID, objc.Sel("setReceiversSpecifier:"), value)
}
// Sets the arguments of the command to `args`.
//
// # Discussion
// 
// Each argument in the dictionary is identified by the same name key used for
// the argument in the command’s class declaration in the script suite file.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/arguments
func (s NSScriptCommand) Arguments() INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("arguments"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (s NSScriptCommand) SetArguments(value INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setArguments:"), value)
}
// Returns a dictionary containing the arguments of the command, evaluated
// from object specifiers to objects if necessary. The keys in the dictionary
// are the argument names.
//
// # Discussion
// 
// Arguments initially can be either a normal object or an object specifier
// such as `word 5` (represented as an instance of an
// [NSScriptObjectSpecifier] subclass). If arguments are object specifiers,
// the receiver evaluates them before using the referenced objects. Returns
// `nil` if the command is not well formed. Also returns `nil` if an object
// specifier does not evaluate to an object or if there is no type defined for
// the argument in the command description.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/evaluatedArguments
func (s NSScriptCommand) EvaluatedArguments() INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("evaluatedArguments"))
	return NSDictionaryFromID(objc.ID(rv))
}
// Sets the object that corresponds to the direct parameter of the Apple event
// from which the receiver derives.
//
// # Discussion
// 
// You don’t normally override this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/directParameter
func (s NSScriptCommand) DirectParameter() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("directParameter"))
	return objectivec.Object{ID: rv}
}
func (s NSScriptCommand) SetDirectParameter(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDirectParameter:"), value)
}
// Returns the command description for the command.
//
// # Discussion
// 
// Once a command is created, its command description is immutable.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/commandDescription
func (s NSScriptCommand) CommandDescription() INSScriptCommandDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandDescription"))
	return NSScriptCommandDescriptionFromID(objc.ID(rv))
}
// Returns a Boolean value indicating whether the receiver is well formed
// according to its command description.
//
// # Discussion
// 
// The method ensures that there is a description of the command and that the
// number of arguments and the types of non-specifier arguments conform to the
// command description.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/isWellFormed
func (s NSScriptCommand) WellFormed() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isWellFormed"))
	return rv
}
// Sets a descriptor for the expected type that will be put in the reply Apple
// event if the sender requested a reply, execution of the receiver completes,
// and an error number was set.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/scriptErrorExpectedTypeDescriptor
func (s NSScriptCommand) ScriptErrorExpectedTypeDescriptor() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scriptErrorExpectedTypeDescriptor"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}
func (s NSScriptCommand) SetScriptErrorExpectedTypeDescriptor(value INSAppleEventDescriptor) {
	objc.Send[struct{}](s.ID, objc.Sel("setScriptErrorExpectedTypeDescriptor:"), value)
}
// Sets a script error number that is associated with the execution of the
// command and is returned in the reply Apple event, if a reply was requested
// by the sender.
//
// # Discussion
// 
// If you override [PerformDefaultImplementation] and an error occurs, you
// should call this method to supply an appropriate error number. In fact, any
// script handler should call this method when an error occurs. The error
// number you supply is returned in the reply Apple event.
// 
// Invoking `` causes an error message to be displayed. To associate a
// specific error message with the error number, you invoke
// [ScriptErrorString]. This make sense, for example, when you set an error
// number that is specific to your application, or when you can supply a
// specific and useful error message to the user.
// 
// If `` is invoked on an [NSScriptCommand] with multiple receivers, the
// command will stop sending command handling messages to more receivers.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/scriptErrorNumber
func (s NSScriptCommand) ScriptErrorNumber() int {
	rv := objc.Send[int](s.ID, objc.Sel("scriptErrorNumber"))
	return rv
}
func (s NSScriptCommand) SetScriptErrorNumber(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setScriptErrorNumber:"), value)
}
// Sets a descriptor for an object that will be put in the reply Apple event
// if the sender requested a reply, execution of the receiver completes, and
// an error number was set.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/scriptErrorOffendingObjectDescriptor
func (s NSScriptCommand) ScriptErrorOffendingObjectDescriptor() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scriptErrorOffendingObjectDescriptor"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}
func (s NSScriptCommand) SetScriptErrorOffendingObjectDescriptor(value INSAppleEventDescriptor) {
	objc.Send[struct{}](s.ID, objc.Sel("setScriptErrorOffendingObjectDescriptor:"), value)
}
// Sets a script error string that is associated with execution of the
// command.
//
// # Discussion
// 
// If you override [PerformDefaultImplementation] and an error occurs, you
// should call this method to supply a string that provides a useful
// explanation. In fact, any script handler should call this method when an
// error occurs.
// 
// Calling this method alone does not cause an error message to be
// displayed—you must also call [ScriptErrorNumber] to supply an error
// number.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/scriptErrorString
func (s NSScriptCommand) ScriptErrorString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scriptErrorString"))
	return NSStringFromID(rv).String()
}
func (s NSScriptCommand) SetScriptErrorString(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setScriptErrorString:"), objc.String(value))
}

			// Protocol methods for NSCoding
			

