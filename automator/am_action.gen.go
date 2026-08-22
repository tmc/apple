// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AMAction] class.
var (
	_AMActionClass     AMActionClass
	_AMActionClassOnce sync.Once
)

func getAMActionClass() AMActionClass {
	_AMActionClassOnce.Do(func() {
		_AMActionClass = AMActionClass{class: objc.GetClass("AMAction")}
	})
	return _AMActionClass
}

// GetAMActionClass returns the class object for AMAction.
func GetAMActionClass() AMActionClass {
	return getAMActionClass()
}

type AMActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMActionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMActionClass) Alloc() AMAction {
	rv := objc.Send[AMAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the interface and general characteristics of
// Automator actions.
//
// # Overview
//
// Automator is an Apple app that allows users to construct and execute
// workflows consisting of a sequence of discrete modules called actions. An
// action performs a specific task, such as copying a file or cropping an
// image, and passes its output to Automator to give to the next action in the
// workflow. Actions are currently implemented as loadable bundles owned by
// objects of the [AMBundleAction] class, a subclass of [AMAction].
//
// The critically important method declared by [AMAction] is
// [AMAction.RunWithInputError]. When Automator executes a workflow, it sends
// this message to each action object in the workflow (in workflow sequence),
// in most cases passing in the output of the previous action as input. The
// action object performs its task in this method and ends by returning an
// output object for the next action in the workflow.
//
// Subclassing [AMAction] is not recommended. For most situations requiring an
// enhancement to the Automator framework, it is sufficient to subclass
// [AMBundleAction].
//
// # Initializing and Encoding
//
//   - [AMAction.InitWithDefinitionFromArchive]: Initializes the action with the specified definition.
//   - [AMAction.InitWithContentsOfURLError]: Loads an Automator action from a file URL.
//   - [AMAction.WriteToDictionary]: Examines the parameters and other configuration information specified in the passed dictionary and adds its own information to it if appropriate.
//
// # Controlling the Action
//
//   - [AMAction.RunWithInputError]: Requests the action to perform its task using the specified input.
//   - [AMAction.RunAsynchronouslyWithInput]: Causes Automator to wait for notification that the action has completed execution, which allows the action to perform an asynchronous operation.
//   - [AMAction.FinishRunningWithError]: Causes the action to stop running and return an error, which, in turn, causes the workflow to stop.
//   - [AMAction.WillFinishRunning]: Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
//   - [AMAction.Stop]: Stops the action from running.
//   - [AMAction.Reset]: Resets the action to its initial state.
//
// # Initializing and Synchronizing the Action User Interface
//
//   - [AMAction.Activated]: Allows the action to synchronize its information with settings in another app.
//   - [AMAction.Opened]: Allows the action to initialize its user interface.
//
// # Updating Action Parameters
//
//   - [AMAction.ParametersUpdated]: Requests the action to update its user interface from its stored parameters, which have changed.
//   - [AMAction.UpdateParameters]: Requests the action to update its stored set of parameters from the settings in the action’s user interface.
//
// # Getting Action Information
//
//   - [AMAction.Name]: The name of the action.
//   - [AMAction.ProgressValue]: A float value between 0 and 1, which indicates how far along the action is while processing.
//   - [AMAction.SetProgressValue]
//   - [AMAction.IgnoresInput]: A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//   - [AMAction.Output]: The action’s output.
//   - [AMAction.SetOutput]
//   - [AMAction.SelectedInputType]: The type of input, in UTI format, of the input received by the action.
//   - [AMAction.SetSelectedInputType]
//   - [AMAction.SelectedOutputType]: The type of output, in UTI format, of the output to be produced by the action.
//   - [AMAction.SetSelectedOutputType]
//   - [AMAction.IsStopped]: A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// # Performing Cleanup Operations
//
//   - [AMAction.Closed]: Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
//
// See: https://developer.apple.com/documentation/Automator/AMAction
type AMAction struct {
	objectivec.Object
}

// AMActionFromID constructs a [AMAction] from an objc.ID.
//
// An abstract class that defines the interface and general characteristics of
// Automator actions.
func AMActionFromID(id objc.ID) AMAction {
	return AMAction{objectivec.Object{ID: id}}
}

// NOTE: AMAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMAction] class.
//
// # Initializing and Encoding
//
//   - [IAMAction.InitWithDefinitionFromArchive]: Initializes the action with the specified definition.
//   - [IAMAction.InitWithContentsOfURLError]: Loads an Automator action from a file URL.
//   - [IAMAction.WriteToDictionary]: Examines the parameters and other configuration information specified in the passed dictionary and adds its own information to it if appropriate.
//
// # Controlling the Action
//
//   - [IAMAction.RunWithInputError]: Requests the action to perform its task using the specified input.
//   - [IAMAction.RunAsynchronouslyWithInput]: Causes Automator to wait for notification that the action has completed execution, which allows the action to perform an asynchronous operation.
//   - [IAMAction.FinishRunningWithError]: Causes the action to stop running and return an error, which, in turn, causes the workflow to stop.
//   - [IAMAction.WillFinishRunning]: Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
//   - [IAMAction.Stop]: Stops the action from running.
//   - [IAMAction.Reset]: Resets the action to its initial state.
//
// # Initializing and Synchronizing the Action User Interface
//
//   - [IAMAction.Activated]: Allows the action to synchronize its information with settings in another app.
//   - [IAMAction.Opened]: Allows the action to initialize its user interface.
//
// # Updating Action Parameters
//
//   - [IAMAction.ParametersUpdated]: Requests the action to update its user interface from its stored parameters, which have changed.
//   - [IAMAction.UpdateParameters]: Requests the action to update its stored set of parameters from the settings in the action’s user interface.
//
// # Getting Action Information
//
//   - [IAMAction.Name]: The name of the action.
//   - [IAMAction.ProgressValue]: A float value between 0 and 1, which indicates how far along the action is while processing.
//   - [IAMAction.SetProgressValue]
//   - [IAMAction.IgnoresInput]: A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//   - [IAMAction.Output]: The action’s output.
//   - [IAMAction.SetOutput]
//   - [IAMAction.SelectedInputType]: The type of input, in UTI format, of the input received by the action.
//   - [IAMAction.SetSelectedInputType]
//   - [IAMAction.SelectedOutputType]: The type of output, in UTI format, of the output to be produced by the action.
//   - [IAMAction.SetSelectedOutputType]
//   - [IAMAction.IsStopped]: A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// # Performing Cleanup Operations
//
//   - [IAMAction.Closed]: Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
//
// See: https://developer.apple.com/documentation/Automator/AMAction
type IAMAction interface {
	objectivec.IObject

	// Topic: Initializing and Encoding

	// Initializes the action with the specified definition.
	InitWithDefinitionFromArchive(dict foundation.INSDictionary, archived bool) AMAction
	// Loads an Automator action from a file URL.
	InitWithContentsOfURLError(fileURL foundation.NSURL) (AMAction, error)
	// Examines the parameters and other configuration information specified in the passed dictionary and adds its own information to it if appropriate.
	WriteToDictionary(dictionary foundation.INSDictionary)

	// Topic: Controlling the Action

	// Requests the action to perform its task using the specified input.
	RunWithInputError(input objectivec.IObject) (objectivec.IObject, error)
	// Causes Automator to wait for notification that the action has completed execution, which allows the action to perform an asynchronous operation.
	RunAsynchronouslyWithInput(input objectivec.IObject)
	// Causes the action to stop running and return an error, which, in turn, causes the workflow to stop.
	FinishRunningWithError(error_ foundation.NSError)
	// Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
	WillFinishRunning()
	// Stops the action from running.
	Stop()
	// Resets the action to its initial state.
	Reset()

	// Topic: Initializing and Synchronizing the Action User Interface

	// Allows the action to synchronize its information with settings in another app.
	Activated()
	// Allows the action to initialize its user interface.
	Opened()

	// Topic: Updating Action Parameters

	// Requests the action to update its user interface from its stored parameters, which have changed.
	ParametersUpdated()
	// Requests the action to update its stored set of parameters from the settings in the action’s user interface.
	UpdateParameters()

	// Topic: Getting Action Information

	// The name of the action.
	Name() string
	// A float value between 0 and 1, which indicates how far along the action is while processing.
	ProgressValue() float64
	SetProgressValue(value float64)
	// A Boolean value that indicates whether the action acts upon its input or the input is ignored.
	IgnoresInput() bool
	// The action’s output.
	Output() objectivec.IObject
	SetOutput(value objectivec.IObject)
	// The type of input, in UTI format, of the input received by the action.
	SelectedInputType() string
	SetSelectedInputType(value string)
	// The type of output, in UTI format, of the output to be produced by the action.
	SelectedOutputType() string
	SetSelectedOutputType(value string)
	// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
	IsStopped() bool

	// Topic: Performing Cleanup Operations

	// Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
	Closed()
}

// Init initializes the instance.
func (a AMAction) Init() AMAction {
	rv := objc.Send[AMAction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMAction) Autorelease() AMAction {
	rv := objc.Send[AMAction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMAction creates a new AMAction instance.
func NewAMAction() AMAction {
	class := getAMActionClass()
	rv := objc.Send[AMAction](objc.ID(class.class), objc.Sel("new"))
	return rv
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
func NewAMActionWithContentsOfURLError(fileURL foundation.NSURL) (AMAction, error) {
	var errorPtr objc.ID
	instance := getAMActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMAction{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AMAction{}, objc.ErrInitFailed
	}
	return AMActionFromID(rv), nil
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
func NewAMActionWithDefinitionFromArchive(dict foundation.INSDictionary, archived bool) AMAction {
	instance := getAMActionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	return AMActionFromID(rv)
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
func (a AMAction) InitWithDefinitionFromArchive(dict foundation.INSDictionary, archived bool) AMAction {
	rv := objc.Send[AMAction](a.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	return rv
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
func (a AMAction) InitWithContentsOfURLError(fileURL foundation.NSURL) (AMAction, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AMAction{}, foundation.NSErrorFrom(errorPtr)
	}
	return AMActionFromID(rv), nil

}

// Examines the parameters and other configuration information specified in
// the passed dictionary and adds its own information to it if appropriate.
//
// dictionary: A dictionary that contains parameter and other configuration information
// about the action.
//
// # Discussion
//
// Automator sends this message to an action object prior to archiving it. In
// its implementation of this method, the action object should first invoke
// the superclass implementation.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/write(to:)
func (a AMAction) WriteToDictionary(dictionary foundation.INSDictionary) {
	objc.Send[objc.ID](a.ID, objc.Sel("writeToDictionary:"), dictionary)
}

// Requests the action to perform its task using the specified input.
//
// input: The input for the receiving action. Should contain one or more objects
// compatible with one of the types specified in the action’s
// [AMAction.SelectedInputType] property.
//
// # Return Value
//
// An [NSArray] object that contains one or more objects of a data type
// compatible with a type specified in the receiving action’s [AMProvides]
// property. If the action doesn’t modify the data passed in `input`, it
// should return it unchanged. If the action doesn’t have any data to
// provide, it should return an empty [NSArray] object.
//
// # Discussion
//
// This method is intended to be overridden.
//
// The input and output objects for actions are usually instances of
// [NSArray]. If the action encounters problems, it should return by
// indirection an [NSError] object that describes the error.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/run(withInput:)
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
func (a AMAction) RunWithInputError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("runWithInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Causes Automator to wait for notification that the action has completed
// execution, which allows the action to perform an asynchronous operation.
//
// input: The input for the action. Should contain one or more objects compatible
// with one of the types specified in the action’s
// [AMAction.SelectedInputType] property.
//
// # Discussion
//
// Override this method in actions that need to make asynchronous calls. After
// [AMAction.RunAsynchronouslyWithInput] is invoked, Automator doesn’t
// continue until the action invokes [AMAction.FinishRunningWithError]. In
// your override of this method, you can make an asynchronous call, wait to be
// notified of its completion, then invoke [AMAction.FinishRunningWithError]
// to signal to Automator that the action has completed.
//
// For actions that don’t need to make asynchronous calls, use
// [runWithInput:fromAction:error:] instead.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/runAsynchronously(withInput:)
//
// [runWithInput:fromAction:error:]: https://developer.apple.com/documentation/Automator/AMAction/runWithInput:fromAction:error:
func (a AMAction) RunAsynchronouslyWithInput(input objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("runAsynchronouslyWithInput:"), input)
}

// Causes the action to stop running and return an error, which, in turn,
// causes the workflow to stop.
//
// error: The error to be returned to Automator.
//
// # Discussion
//
// Call this method on any action that overrides
// [AMAction.RunAsynchronouslyWithInput] in order to make asynchronous calls.
// When [AMAction.FinishRunningWithError] is invoked, it immediately calls
// [AMAction.WillFinishRunning].
//
// See: https://developer.apple.com/documentation/Automator/AMAction/finishRunningWithError(_:)
func (a AMAction) FinishRunningWithError(error_ foundation.NSError) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishRunningWithError:"), error_)
}

// Provides an opportunity for an action to perform cleanup operations, such
// as closing windows and deallocating memory.
//
// # Discussion
//
// Overridde this method in actions that need to make asynchronous calls.
// Automator invokes this method when the action is about to complete its run
// phase.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/willFinishRunning()
func (a AMAction) WillFinishRunning() {
	objc.Send[objc.ID](a.ID, objc.Sel("willFinishRunning"))
}

// Stops the action from running.
//
// # Discussion
//
// The output acquired by the action during execution of the current workflow
// is still accessible to Automator.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/stop()
func (a AMAction) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}

// Resets the action to its initial state.
//
// # Discussion
//
// Resetting causes the action to release its output generated from the
// current execution of the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/reset()
func (a AMAction) Reset() {
	objc.Send[objc.ID](a.ID, objc.Sel("reset"))
}

// Allows the action to synchronize its information with settings in another
// app.
//
// # Discussion
//
// The system invokes this method when the window of the Automator workflow to
// which the action belongs becomes the main window.
//
// Be sure to invoke the superclass implementation of this method as the last
// thing in your implementation.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/activated()
func (a AMAction) Activated() {
	objc.Send[objc.ID](a.ID, objc.Sel("activated"))
}

// Allows the action to initialize its user interface.
//
// # Discussion
//
// The system invokes this method when the action is first added to a
// workflow.
//
// You should perform all initializations of an action’s user interface in
// this method and not in `awakeFromNib`. Be sure to invoke the superclass
// implementation of this method as the final step of your implementation.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/opened()
func (a AMAction) Opened() {
	objc.Send[objc.ID](a.ID, objc.Sel("opened"))
}

// Requests the action to update its user interface from its stored
// parameters, which have changed.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/parametersUpdated()
func (a AMAction) ParametersUpdated() {
	objc.Send[objc.ID](a.ID, objc.Sel("parametersUpdated"))
}

// Requests the action to update its stored set of parameters from the
// settings in the action’s user interface.
//
// # Discussion
//
// This message sends just before an action is saved, copied, or run.
// Preferably, an action’s settings should not solely reside in the controls
// of its view, but if they do, the action can fetch and save them in this
// method.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/updateParameters()
func (a AMAction) UpdateParameters() {
	objc.Send[objc.ID](a.ID, objc.Sel("updateParameters"))
}

// Invoked by Automator when the receiving action is removed from a workflow,
// allowing it to perform cleanup operations.
//
// # Discussion
//
// This method is intended to be overridden, so that your action can perform
// its specific cleanup operations.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/closed()
func (a AMAction) Closed() {
	objc.Send[objc.ID](a.ID, objc.Sel("closed"))
}

// The name of the action.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/name
func (a AMAction) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// A float value between 0 and 1, which indicates how far along the action is
// while processing.
//
// # Discussion
//
// Setting this value causes Automator’s action progress indicator
// (displayed as a workflow runs) to update in order to provide the user with
// an indication of progress.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/progressValue
func (a AMAction) ProgressValue() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("progressValue"))
	return rv
}
func (a AMAction) SetProgressValue(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgressValue:"), value)
}

// A Boolean value that indicates whether the action acts upon its input or
// the input is ignored.
//
// # Discussion
//
// true if the action acts upon its input, otherwise false.
//
// Many actions act upon their input, but an action may merely pass on its
// input or, rarely, ignore it.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/ignoresInput
func (a AMAction) IgnoresInput() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("ignoresInput"))
	return rv
}

// The action’s output.
//
// # Return Value
//
// The receiving action’s output, or `nil` if called before the action is
// run.
//
// # Discussion
//
// `nil` if called before the action is run.
//
// This method is used in conjunction with the [AMWorkflow] class, which
// allows access to the actions in a workflow. Within a workflow, for example,
// you might iteratively inspect the output of each action. Or, on completion
// of a workflow, you might examine the output of the last action, to
// determine the output of the workflow.
//
// This parameter can also be used when running an action asynchronously. Call
// `setOutput` to specify the output the action produces.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/output
func (a AMAction) Output() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("output"))
	return objectivec.Object{ID: rv}
}
func (a AMAction) SetOutput(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutput:"), value)
}

// The type of input, in UTI format, of the input received by the action.
//
// # Discussion
//
// Getting this value provides the type of input the action is configured to
// accept. For example, your action may have the ability to accept files and
// folders, or documents, depending on how it’s configured and what action
// precedes it in the workflow.
//
// The input types the action supports are specified in the action’s
// `Info.Plist()` file. By default, this property defaults to first input type
// in the `Info.Plist()` file.
//
// Through its interface, the action can could be configured to allow the user
// to specify the type of input the action should accept. For example, a
// Contacts action may allow the user to configure whether the action accepts
// people or groups. In cases like this, set this property value to explicitly
// indicate the input type the action accepts. Setting this value to
// accurately reflect the appropriate type of input helps Automator determine
// whether the input the action receives is compatible, or can be made
// compatible, with the action.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/selectedInputType
func (a AMAction) SelectedInputType() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("selectedInputType"))
	return foundation.NSStringFromID(rv).String()
}
func (a AMAction) SetSelectedInputType(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSelectedInputType:"), objc.String(value))
}

// The type of output, in UTI format, of the output to be produced by the
// action.
//
// # Discussion
//
// Getting this value provides the type of output the action is configured to
// produce. For example, your action may have the ability to output files and
// folders, or documents, depending on how it’s configured or what it
// encounters while processing.
//
// The output types the action supports are specified in the action’s
// `Info.Plist()` file. By default, this property defaults to the first output
// type in the `Info.Plist()` file.
//
// Set this value to explicitly specify the output type the action produces.
// Setting this value to accurately reflect the appropriate output helps
// Automator determine whether the output is compatible, or can be made
// compatible, with the next action in the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/selectedOutputType
func (a AMAction) SelectedOutputType() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("selectedOutputType"))
	return foundation.NSStringFromID(rv).String()
}
func (a AMAction) SetSelectedOutputType(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSelectedOutputType:"), objc.String(value))
}

// A Boolean value that indicates whether the user clicked the stop button on
// the parent workflow.
//
// # Discussion
//
// This value is true if the user clicked the stop button, or false if the
// workflow is still running. This property should be referenced during
// lengthy action processes, such as a loop, in order to determine whether to
// exit the operation and stop the action.
//
// See: https://developer.apple.com/documentation/Automator/AMAction/isStopped
func (a AMAction) IsStopped() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isStopped"))
	return rv
}
