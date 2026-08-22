// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AMBundleAction] class.
var (
	_AMBundleActionClass     AMBundleActionClass
	_AMBundleActionClassOnce sync.Once
)

func getAMBundleActionClass() AMBundleActionClass {
	_AMBundleActionClassOnce.Do(func() {
		_AMBundleActionClass = AMBundleActionClass{class: objc.GetClass("AMBundleAction")}
	})
	return _AMBundleActionClass
}

// GetAMBundleActionClass returns the class object for AMBundleAction.
func GetAMBundleActionClass() AMBundleActionClass {
	return getAMBundleActionClass()
}

type AMBundleActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMBundleActionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMBundleActionClass) Alloc() AMBundleAction {
	rv := objc.Send[AMBundleAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an Automator action that’s a loadable bundle.
//
// # Overview
//
// Automator loads action bundles from standard locations in the file system:
// `/System/Library/Automator`, `/Library/Automator`, and
// `~/Library/Automator`.
//
// [AMBundleAction] objects have several important properties:
//
// - The [AMBundleAction.Bundle] object associated with the action’s physical bundle - The
// action’s view, which holds its user interface - A parameters dictionary
// that reflects the settings in the user interface
//
// When you create a Cocoa Automator Action project in Xcode, the project
// template includes a custom subclass of [AMBundleAction]. This custom class
// uses the name of the project.
//
// You must provide an implementation of [AMAction.RunWithInputError], which
// is declared by the superclass [AMAction]. If you add any instance
// variables, you must override the
// [AMShellScriptAction.InitWithDefinitionFromArchive] method and the
// [AMAction.WriteToDictionary] method of [AMAction] to work with them.
//
// # Initializing the Action
//
//   - [AMBundleAction.AwakeFromBundle]: Allows the action object to perform setup tasks requiring the presence of all bundle objects.
//
// # Managing Action Properties
//
//   - [AMBundleAction.Bundle]: The action’s bundle object.
//   - [AMBundleAction.HasView]: A Boolean value that indicates whether the action has a view associated with it.
//   - [AMBundleAction.View]: The action’s view object.
//   - [AMBundleAction.Parameters]: The action’s parameters.
//   - [AMBundleAction.SetParameters]
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction
//
// [AMBundleAction.Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
type AMBundleAction struct {
	AMAction
}

// AMBundleActionFromID constructs a [AMBundleAction] from an objc.ID.
//
// An object that represents an Automator action that’s a loadable bundle.
func AMBundleActionFromID(id objc.ID) AMBundleAction {
	return AMBundleAction{AMAction: AMActionFromID(id)}
}

// NOTE: AMBundleAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMBundleAction] class.
//
// # Initializing the Action
//
//   - [IAMBundleAction.AwakeFromBundle]: Allows the action object to perform setup tasks requiring the presence of all bundle objects.
//
// # Managing Action Properties
//
//   - [IAMBundleAction.Bundle]: The action’s bundle object.
//   - [IAMBundleAction.HasView]: A Boolean value that indicates whether the action has a view associated with it.
//   - [IAMBundleAction.View]: The action’s view object.
//   - [IAMBundleAction.Parameters]: The action’s parameters.
//   - [IAMBundleAction.SetParameters]
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction
type IAMBundleAction interface {
	IAMAction

	// Topic: Initializing the Action

	// Allows the action object to perform setup tasks requiring the presence of all bundle objects.
	AwakeFromBundle()

	// Topic: Managing Action Properties

	// The action’s bundle object.
	Bundle() foundation.Bundle
	// A Boolean value that indicates whether the action has a view associated with it.
	HasView() bool
	// The action’s view object.
	View() appkit.NSView
	// The action’s parameters.
	Parameters() foundation.INSDictionary
	SetParameters(value foundation.INSDictionary)
}

// Init initializes the instance.
func (a AMBundleAction) Init() AMBundleAction {
	rv := objc.Send[AMBundleAction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMBundleAction) Autorelease() AMBundleAction {
	rv := objc.Send[AMBundleAction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMBundleAction creates a new AMBundleAction instance.
func NewAMBundleAction() AMBundleAction {
	class := getAMBundleActionClass()
	rv := objc.Send[AMBundleAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Automator/AMBundleAction/init(coder:)
func NewAMBundleActionWithCoder(coder foundation.INSCoder) AMBundleAction {
	instance := getAMBundleActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AMBundleActionFromID(rv)
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
func NewAMBundleActionWithContentsOfURLError(fileURL foundation.NSURL) (AMBundleAction, error) {
	var errorPtr objc.ID
	instance := getAMBundleActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMBundleAction{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AMBundleAction{}, objc.ErrInitFailed
	}
	return AMBundleActionFromID(rv), nil
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
func NewAMBundleActionWithDefinitionFromArchive(dict foundation.INSDictionary, archived bool) AMBundleAction {
	instance := getAMBundleActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	return AMBundleActionFromID(rv)
}

// Allows the action object to perform setup tasks requiring the presence of
// all bundle objects.
//
// # Discussion
//
// The system sends this message to the action object when all objects in its
// bundle have been unarchived. Use this method to perform setup tasks such as
// adding the action object as an observer of notifications, dynamically
// establishing bindings, and dynamically setting targets and actions.
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction/awakeFromBundle()
func (a AMBundleAction) AwakeFromBundle() {
	objc.Send[objc.ID](a.ID, objc.Sel("awakeFromBundle"))
}

// The action’s bundle object.
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction/bundle
func (a AMBundleAction) Bundle() foundation.Bundle {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bundle"))
	return foundation.BundleFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the action has a view associated
// with it.
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction/hasView
func (a AMBundleAction) HasView() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasView"))
	return rv
}

// The action’s view object.
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction/view
func (a AMBundleAction) View() appkit.NSView {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("view"))
	return appkit.NSViewFromID(objc.ID(rv))
}

// The action’s parameters.
//
// # Discussion
//
// The parameters of an action reflect the choices made and values entered in
// the action’s user interface. Keys in the parameters dictionary identify
// specific user-interface objects. If an action uses the Cocoa bindings
// mechanism, the parameters of an [AMBundleAction] object are automatically
// set.
//
// See: https://developer.apple.com/documentation/Automator/AMBundleAction/parameters
func (a AMBundleAction) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AMBundleAction) SetParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setParameters:"), value)
}
