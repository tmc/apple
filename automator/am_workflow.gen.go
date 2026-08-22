// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AMWorkflow] class.
var (
	_AMWorkflowClass     AMWorkflowClass
	_AMWorkflowClassOnce sync.Once
)

func getAMWorkflowClass() AMWorkflowClass {
	_AMWorkflowClassOnce.Do(func() {
		_AMWorkflowClass = AMWorkflowClass{class: objc.GetClass("AMWorkflow")}
	})
	return _AMWorkflowClass
}

// GetAMWorkflowClass returns the class object for AMWorkflow.
func GetAMWorkflowClass() AMWorkflowClass {
	return getAMWorkflowClass()
}

type AMWorkflowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMWorkflowClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMWorkflowClass) Alloc() AMWorkflow {
	rv := objc.Send[AMWorkflow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that lets you use an Automator workflow in your app.
//
// # Overview
//
// A workflow consists of one or more actions, discrete tasks which together
// can perform complex automation tasks. Your app can use workflows to package
// its own features and to take advantage of features provided by other apps.
// You create actions with Xcode, while you create workflows with the
// Automator app.
//
// You can load and run a workflow with minimal overhead by using the class
// method [AMWorkflowClass.RunWorkflowAtURLWithInputError]. However, in
// situations where you need greater control, such as the ability to start and
// stop the workflow, you can use an instance of the [AMWorkflowController]
// class instead. In that case, you must create and initialize both the
// workflow and the workflow controller objects.
//
// In either case, the workflow runs in a separate process so that any actions
// it contains are executed in a separate memory space. That helps to insulate
// your app from crashes, memory leaks, or exceptions that might occur from
// running the actions in the workflow.
//
// You can display a workflow with an instance of [AMWorkflowView].
//
// # Creating a Workflow
//
//   - [AMWorkflow.InitWithContentsOfURLError]: Creates and initializes a workflow based on the contents of the specified file.
//
// # Saving Changes to a Workflow
//
//   - [AMWorkflow.WriteToURLError]: Writes the workflow to the specified file.
//
// # Getting Information About a Workflow
//
//   - [AMWorkflow.Actions]: An array of the workflow’s actions.
//   - [AMWorkflow.FileURL]: A URL that specifies the location of the workflow file.
//   - [AMWorkflow.ValueForVariableWithName]: Returns the value of the workflow variable with the specified name.
//
// # Working with the Workflow’s Input and Output
//
//   - [AMWorkflow.Input]: The input data that is passed to the first action in the workflow.
//   - [AMWorkflow.SetInput]
//   - [AMWorkflow.Output]: The output data that is provided by the last action in the workflow.
//
// # Manipulating the Workflow
//
//   - [AMWorkflow.SetValueForVariableWithName]: Sets the value of the workflow variable with the specified name.
//
// # Manipulating the Workflow’s Actions
//
//   - [AMWorkflow.AddAction]: Adds the specified action at the end of the receiving workflow.
//   - [AMWorkflow.InsertActionAtIndex]: Inserts the specified action at the specified position of the receiving workflow.
//   - [AMWorkflow.MoveActionAtIndexToIndex]: Moves the action from the specified start position to the specified end position in the receiving workflow.
//   - [AMWorkflow.RemoveAction]: Removes the specified action from the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow
type AMWorkflow struct {
	objectivec.Object
}

// AMWorkflowFromID constructs a [AMWorkflow] from an objc.ID.
//
// An object that lets you use an Automator workflow in your app.
func AMWorkflowFromID(id objc.ID) AMWorkflow {
	return AMWorkflow{objectivec.Object{ID: id}}
}

// NOTE: AMWorkflow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMWorkflow] class.
//
// # Creating a Workflow
//
//   - [IAMWorkflow.InitWithContentsOfURLError]: Creates and initializes a workflow based on the contents of the specified file.
//
// # Saving Changes to a Workflow
//
//   - [IAMWorkflow.WriteToURLError]: Writes the workflow to the specified file.
//
// # Getting Information About a Workflow
//
//   - [IAMWorkflow.Actions]: An array of the workflow’s actions.
//   - [IAMWorkflow.FileURL]: A URL that specifies the location of the workflow file.
//   - [IAMWorkflow.ValueForVariableWithName]: Returns the value of the workflow variable with the specified name.
//
// # Working with the Workflow’s Input and Output
//
//   - [IAMWorkflow.Input]: The input data that is passed to the first action in the workflow.
//   - [IAMWorkflow.SetInput]
//   - [IAMWorkflow.Output]: The output data that is provided by the last action in the workflow.
//
// # Manipulating the Workflow
//
//   - [IAMWorkflow.SetValueForVariableWithName]: Sets the value of the workflow variable with the specified name.
//
// # Manipulating the Workflow’s Actions
//
//   - [IAMWorkflow.AddAction]: Adds the specified action at the end of the receiving workflow.
//   - [IAMWorkflow.InsertActionAtIndex]: Inserts the specified action at the specified position of the receiving workflow.
//   - [IAMWorkflow.MoveActionAtIndexToIndex]: Moves the action from the specified start position to the specified end position in the receiving workflow.
//   - [IAMWorkflow.RemoveAction]: Removes the specified action from the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow
type IAMWorkflow interface {
	objectivec.IObject

	// Topic: Creating a Workflow

	// Creates and initializes a workflow based on the contents of the specified file.
	InitWithContentsOfURLError(fileURL foundation.NSURL) (AMWorkflow, error)

	// Topic: Saving Changes to a Workflow

	// Writes the workflow to the specified file.
	WriteToURLError(fileURL foundation.NSURL) (bool, error)

	// Topic: Getting Information About a Workflow

	// An array of the workflow’s actions.
	Actions() []AMAction
	// A URL that specifies the location of the workflow file.
	FileURL() foundation.NSURL
	// Returns the value of the workflow variable with the specified name.
	ValueForVariableWithName(variableName string) objectivec.IObject

	// Topic: Working with the Workflow’s Input and Output

	// The input data that is passed to the first action in the workflow.
	Input() objectivec.IObject
	SetInput(value objectivec.IObject)
	// The output data that is provided by the last action in the workflow.
	Output() objectivec.IObject

	// Topic: Manipulating the Workflow

	// Sets the value of the workflow variable with the specified name.
	SetValueForVariableWithName(value objectivec.IObject, variableName string) bool

	// Topic: Manipulating the Workflow’s Actions

	// Adds the specified action at the end of the receiving workflow.
	AddAction(action IAMAction)
	// Inserts the specified action at the specified position of the receiving workflow.
	InsertActionAtIndex(action IAMAction, index uint)
	// Moves the action from the specified start position to the specified end position in the receiving workflow.
	MoveActionAtIndexToIndex(startIndex uint, endIndex uint)
	// Removes the specified action from the workflow.
	RemoveAction(action IAMAction)
}

// Init initializes the instance.
func (a AMWorkflow) Init() AMWorkflow {
	rv := objc.Send[AMWorkflow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMWorkflow) Autorelease() AMWorkflow {
	rv := objc.Send[AMWorkflow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflow creates a new AMWorkflow instance.
func NewAMWorkflow() AMWorkflow {
	class := getAMWorkflowClass()
	rv := objc.Send[AMWorkflow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes a workflow based on the contents of the specified
// file.
//
// fileURL: URL that specifies the location of a workflow file.
//
// # Return Value
//
// The initialized workflow object. On error, returns nil.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/init(contentsOf:)
func NewAMWorkflowWithContentsOfURLError(fileURL foundation.NSURL) (AMWorkflow, error) {
	var errorPtr objc.ID
	instance := getAMWorkflowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMWorkflow{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AMWorkflow{}, objc.ErrInitFailed
	}
	return AMWorkflowFromID(rv), nil
}

// Creates and initializes a workflow based on the contents of the specified
// file.
//
// fileURL: URL that specifies the location of a workflow file.
//
// # Return Value
//
// The initialized workflow object. On error, returns nil.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/init(contentsOf:)
func (a AMWorkflow) InitWithContentsOfURLError(fileURL foundation.NSURL) (AMWorkflow, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMWorkflow{}, foundation.NSErrorFrom(errorPtr)
	}
	return AMWorkflowFromID(rv), nil

}

// Writes the workflow to the specified file.
//
// fileURL: URL that specifies the file location to write the workflow.
//
// # Discussion
//
// You might want to save the workflow, for example, because you have made
// changes to a variable it contains.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/write(to:)
func (a AMWorkflow) WriteToURLError(fileURL foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("writeToURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the value of the workflow variable with the specified name.
//
// variableName: The variable name.
//
// # Return Value
//
// The value for the variable. Returns `nil` if no variable is found with the
// specified name.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/valueForVariable(withName:)
func (a AMWorkflow) ValueForVariableWithName(variableName string) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("valueForVariableWithName:"), objc.String(variableName))
	return objectivec.Object{ID: rv}
}

// Sets the value of the workflow variable with the specified name.
//
// value: The value to set for the named variable.
//
// variableName: The name of a variable to set the value for.
//
// # Return Value
//
// true if `variableName` was found and its value is set; otherwise false.
//
// # Discussion
//
// This method does nothing if the variable specified by `variableName` is not
// found.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/setValue(_:forVariableWithName:)
func (a AMWorkflow) SetValueForVariableWithName(value objectivec.IObject, variableName string) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setValue:forVariableWithName:"), value, objc.String(variableName))
	return rv
}

// Adds the specified action at the end of the receiving workflow.
//
// action: The action to add.
//
// # Discussion
//
// The workflow retains the action but does not copy it.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/addAction(_:)
func (a AMWorkflow) AddAction(action IAMAction) {
	objc.Send[objc.ID](a.ID, objc.Sel("addAction:"), action)
}

// Inserts the specified action at the specified position of the receiving
// workflow.
//
// action: The action to insert.
//
// index: The position in the receiver at which to insert the action. If the position
// is invalid, this method does nothing.
//
// # Discussion
//
// The workflow retains the action but does not copy it.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/insertAction(_:at:)
func (a AMWorkflow) InsertActionAtIndex(action IAMAction, index uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("insertAction:atIndex:"), action, index)
}

// Moves the action from the specified start position to the specified end
// position in the receiving workflow.
//
// startIndex: The start position of the action to move.
//
// endIndex: The end position for the action that is moved.
//
// # Discussion
//
// If either index is invalid, this method does nothing.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/moveAction(at:to:)
func (a AMWorkflow) MoveActionAtIndexToIndex(startIndex uint, endIndex uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("moveActionAtIndex:toIndex:"), startIndex, endIndex)
}

// Removes the specified action from the workflow.
//
// action: The action to be removed.
//
// # Discussion
//
// The action receives an `AMAction closed` message before being released.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/removeAction(_:)
func (a AMWorkflow) RemoveAction(action IAMAction) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeAction:"), action)
}

// Loads and runs the specified workflow file.
//
// fileURL: A URL that specifies the location of a workflow file.
//
// input: The input for the first action in the workflow. Pass `nil` if the first
// action doesn’t need input.
//
// # Return Value
//
// This method returns `nil` on error, or if the action completes successfully
// without producing output. The error argument must be examined to determine
// which scenario occurred. Otherwise, this method returns the output of the
// last action in the workflow. Your application may need to convert the data
// to a desired type.
//
// # Discussion
//
// Use this method to run a workflow without the overhead of performing a
// separate allocation, setting up a workflow controller, and so on. In
// situations where you need greater control, such as the ability to start and
// stop the workflow, use an instance of the [AMWorkflowController] class
// instead.
//
// The workflow is run in a separate process so that any actions it contains
// are executed in a separate memory space. This helps to insulate the app
// from crashes, memory leaks, or exceptions that might occur from running the
// actions in the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/run(at:withInput:)
func (_AMWorkflowClass AMWorkflowClass) RunWorkflowAtURLWithInputError(fileURL foundation.NSURL, input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AMWorkflowClass.class), objc.Sel("runWorkflowAtURL:withInput:error:"), fileURL, input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// An array of the workflow’s actions.
//
// # Return Value
//
// An array of actions for the workflow file. Actions are instances of classes
// such as [AMBundleAction], [AMAppleScriptAction], and [AMShellScriptAction].
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/actions
//
// [AMAppleScriptAction]: https://developer.apple.com/documentation/Automator/AMAppleScriptAction
func (a AMWorkflow) Actions() []AMAction {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("actions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AMAction {
		return AMActionFromID(id)
	})
}

// A URL that specifies the location of the workflow file.
//
// # Return Value
//
// URL that specifies the location of the workflow file.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/fileURL
func (a AMWorkflow) FileURL() foundation.NSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The input data that is passed to the first action in the workflow.
//
// # Return Value
//
// The input for the first action in the workflow. Should be a data type the
// action can use, or a type that can be converted to one the action can use.
// Use “ to set the input data for the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/input
func (a AMWorkflow) Input() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("input"))
	return objectivec.Object{ID: rv}
}
func (a AMWorkflow) SetInput(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setInput:"), value)
}

// The output data that is provided by the last action in the workflow.
//
// # Return Value
//
// The output for the for last action in the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflow/output
func (a AMWorkflow) Output() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("output"))
	return objectivec.Object{ID: rv}
}
