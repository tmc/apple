// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AMShellScriptAction] class.
var (
	_AMShellScriptActionClass     AMShellScriptActionClass
	_AMShellScriptActionClassOnce sync.Once
)

func getAMShellScriptActionClass() AMShellScriptActionClass {
	_AMShellScriptActionClassOnce.Do(func() {
		_AMShellScriptActionClass = AMShellScriptActionClass{class: objc.GetClass("AMShellScriptAction")}
	})
	return _AMShellScriptActionClass
}

// GetAMShellScriptActionClass returns the class object for AMShellScriptAction.
func GetAMShellScriptActionClass() AMShellScriptActionClass {
	return getAMShellScriptActionClass()
}

type AMShellScriptActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMShellScriptActionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMShellScriptActionClass) Alloc() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents Automator actions whose runtime behavior is
// driven by a shell script or by a Perl or Python script.
//
// # Overview
//
// When you create a Shell Script Automator Action project in Xcode, the
// project template supplies an [AMShellScriptAction] instance as the
// Principal Class of the action bundle. This ready-made instance provides a
// default implementation of the [AMAction] [AMAction.RunWithInputError]
// method that uses the logic defined in the script. You can substitute your
// own subclass of [AMShellScriptAction] for Principal Class if you need to.
//
// # Handling the I/O Separator Character
//
//   - [AMShellScriptAction.InputFieldSeparator]: A string to use as the delimiter between items in the string passed to the action through standard input.
//   - [AMShellScriptAction.OutputFieldSeparator]: A string to use as a delimiter in the string output by the action.
//   - [AMShellScriptAction.RemapLineEndings]: A Boolean value that indicates whether you want automatic remapping of carriage return (`\r`) to newline (`\n`) characters in the input string.
//
// See: https://developer.apple.com/documentation/Automator/AMShellScriptAction
type AMShellScriptAction struct {
	AMBundleAction
}

// AMShellScriptActionFromID constructs a [AMShellScriptAction] from an objc.ID.
//
// An object that represents Automator actions whose runtime behavior is
// driven by a shell script or by a Perl or Python script.
func AMShellScriptActionFromID(id objc.ID) AMShellScriptAction {
	return AMShellScriptAction{AMBundleAction: AMBundleActionFromID(id)}
}

// NOTE: AMShellScriptAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMShellScriptAction] class.
//
// # Handling the I/O Separator Character
//
//   - [IAMShellScriptAction.InputFieldSeparator]: A string to use as the delimiter between items in the string passed to the action through standard input.
//   - [IAMShellScriptAction.OutputFieldSeparator]: A string to use as a delimiter in the string output by the action.
//   - [IAMShellScriptAction.RemapLineEndings]: A Boolean value that indicates whether you want automatic remapping of carriage return (`\r`) to newline (`\n`) characters in the input string.
//
// See: https://developer.apple.com/documentation/Automator/AMShellScriptAction
type IAMShellScriptAction interface {
	IAMBundleAction

	// Topic: Handling the I/O Separator Character

	// A string to use as the delimiter between items in the string passed to the action through standard input.
	InputFieldSeparator() string
	// A string to use as a delimiter in the string output by the action.
	OutputFieldSeparator() string
	// A Boolean value that indicates whether you want automatic remapping of carriage return (`\r`) to newline (`\n`) characters in the input string.
	RemapLineEndings() bool
}

// Init initializes the instance.
func (a AMShellScriptAction) Init() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMShellScriptAction) Autorelease() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMShellScriptAction creates a new AMShellScriptAction instance.
func NewAMShellScriptAction() AMShellScriptAction {
	class := getAMShellScriptActionClass()
	rv := objc.Send[AMShellScriptAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Automator/AMBundleAction/init(coder:)
func NewAMShellScriptActionWithCoder(coder foundation.INSCoder) AMShellScriptAction {
	instance := getAMShellScriptActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AMShellScriptActionFromID(rv)
}

// Loads an Automator action from a file URL.
//
// fileURL: URL that specifies the location of an action file.
//
// # Return Value
//
// The initialized action.
//
// # Discussion
//
// This method is typically invoked by app that use the [AMWorkflow] class to
// embed Automator workflows. It is used to allow creation of actions for a
// workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/init(contentsOf:)
func NewAMShellScriptActionWithContentsOfURLError(fileURL foundation.NSURL) (AMShellScriptAction, error) {
	var errorPtr objc.ID
	instance := getAMShellScriptActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMShellScriptAction{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AMShellScriptAction{}, objc.ErrInitFailed
	}
	return AMShellScriptActionFromID(rv), nil
}

// Initializes the action with the specified definition.
//
// dict: A dictionary that describes the action, including any custom definition
// properties.
//
// archived: If the action is being unarchived, true; otherwise, false.
//
// # Return Value
//
// The initialized action.
//
// # Discussion
//
// This is the primary initializer for all Automator classes. The Automator
// app sends this message to instances of [AMAction] both when it loads
// actions bundles and when it unarchives them.
//
// The [AMAction] object being instantiated should perform whatever
// initializations are necessary after invoking `super`’s implementation of
// this method. It can then examine the values in `dict`, particularly if the
// action had been archived with custom definition properties.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/init(definition:fromArchive:)
func NewAMShellScriptActionWithDefinitionFromArchive(dict foundation.INSDictionary, archived bool) AMShellScriptAction {
	instance := getAMShellScriptActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	return AMShellScriptActionFromID(rv)
}

// A string to use as the delimiter between items in the string passed to the
// action through standard input.
//
// # Discussion
//
// The Automator framework converts the output from the previous action (which
// is usually in the form of a list or array) into a single string in which
// the array elements are concatenated by the input field separator. By
// default, this separator is the newline character (`\n`). You can override
// this method to, for example, return a null character (`\0`) to provide
// null-terminated strings for `xargs -0`.
//
// See: https://developer.apple.com/documentation/Automator/AMShellScriptAction/inputFieldSeparator
func (a AMShellScriptAction) InputFieldSeparator() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFieldSeparator"))
	return foundation.NSStringFromID(rv).String()
}

// A string to use as a delimiter in the string output by the action.
//
// # Discussion
//
// Upon completion, the Automator framework converts an output string provided
// by the action into an array (or list), to be passed to the next action in
// the workflow for further processing. The elements in this array are derived
// from fields delimited by the output field separator. The default value is
// the separator character returned by
// [AMShellScriptAction.InputFieldSeparator]. Override this method if you want
// a different delimiter for output.
//
// See: https://developer.apple.com/documentation/Automator/AMShellScriptAction/outputFieldSeparator
func (a AMShellScriptAction) OutputFieldSeparator() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFieldSeparator"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether you want automatic remapping of
// carriage return (`\r`) to newline (`\n`) characters in the input string.
//
// # Discussion
//
// The default is false. Override to return true if you want the remapping to
// occur.
//
// See: https://developer.apple.com/documentation/Automator/AMShellScriptAction/remapLineEndings
func (a AMShellScriptAction) RemapLineEndings() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("remapLineEndings"))
	return rv
}
