// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AMWorkflowController] class.
var (
	_AMWorkflowControllerClass     AMWorkflowControllerClass
	_AMWorkflowControllerClassOnce sync.Once
)

func getAMWorkflowControllerClass() AMWorkflowControllerClass {
	_AMWorkflowControllerClassOnce.Do(func() {
		_AMWorkflowControllerClass = AMWorkflowControllerClass{class: objc.GetClass("AMWorkflowController")}
	})
	return _AMWorkflowControllerClass
}

// GetAMWorkflowControllerClass returns the class object for AMWorkflowController.
func GetAMWorkflowControllerClass() AMWorkflowControllerClass {
	return getAMWorkflowControllerClass()
}

type AMWorkflowControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMWorkflowControllerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMWorkflowControllerClass) Alloc() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that lets you manage an Automator workflow in your app.
//
// # Overview
//
// A controller can run and stop a workflow and obtain information about its
// state. The controller’s delegate ([AMWorkflowControllerDelegate])
// receives messages as the workflow is executed and its actions are run.
//
// You can load and run a workflow with minimal overhead by using the
// [AMWorkflow] class method [AMWorkflowClass.RunWorkflowAtURLWithInputError].
// Use [AMWorkflowController] where you need greater control, such as the
// ability to start and stop the workflow. In that case, you must create and
// initialize both the workflow and the workflow controller objects.
//
// # Accessing the Workflow
//
//   - [AMWorkflowController.Workflow]: The controller’s workflow.
//   - [AMWorkflowController.SetWorkflow]
//
// # Accessing the Workflow View
//
//   - [AMWorkflowController.WorkflowView]: The controller’s workflow view.
//   - [AMWorkflowController.SetWorkflowView]
//
// # Accessing the Delegate
//
//   - [AMWorkflowController.Delegate]: The controller’s delegate.
//   - [AMWorkflowController.SetDelegate]
//
// # Controlling the Workflow
//
//   - [AMWorkflowController.Pause]: Pauses a workflow that’s running.
//   - [AMWorkflowController.Reset]: Stops a workflow, clears any action results, and resets the workflow back to an un-run state.
//   - [AMWorkflowController.Run]: Runs the associated workflow, after first clearing any results stored by its actions during any previous run.
//   - [AMWorkflowController.Step]: In a paused workflow, runs the next action in the workflow and then pauses again.
//   - [AMWorkflowController.Stop]: Stops the associated workflow.
//
// # Getting Workflow Information
//
//   - [AMWorkflowController.CanRun]: A Boolean value that indicates whether the controller’s workflow is able to run.
//   - [AMWorkflowController.IsPaused]: A Boolean value that indicates whether the controller’s workflow is currently paused.
//   - [AMWorkflowController.IsRunning]: A Boolean value that indicates whether the controller’s workflow is currently running.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController
type AMWorkflowController struct {
	appkit.NSController
}

// AMWorkflowControllerFromID constructs a [AMWorkflowController] from an objc.ID.
//
// An object that lets you manage an Automator workflow in your app.
func AMWorkflowControllerFromID(id objc.ID) AMWorkflowController {
	return AMWorkflowController{NSController: appkit.NSControllerFromID(id)}
}

// NOTE: AMWorkflowController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMWorkflowController] class.
//
// # Accessing the Workflow
//
//   - [IAMWorkflowController.Workflow]: The controller’s workflow.
//   - [IAMWorkflowController.SetWorkflow]
//
// # Accessing the Workflow View
//
//   - [IAMWorkflowController.WorkflowView]: The controller’s workflow view.
//   - [IAMWorkflowController.SetWorkflowView]
//
// # Accessing the Delegate
//
//   - [IAMWorkflowController.Delegate]: The controller’s delegate.
//   - [IAMWorkflowController.SetDelegate]
//
// # Controlling the Workflow
//
//   - [IAMWorkflowController.Pause]: Pauses a workflow that’s running.
//   - [IAMWorkflowController.Reset]: Stops a workflow, clears any action results, and resets the workflow back to an un-run state.
//   - [IAMWorkflowController.Run]: Runs the associated workflow, after first clearing any results stored by its actions during any previous run.
//   - [IAMWorkflowController.Step]: In a paused workflow, runs the next action in the workflow and then pauses again.
//   - [IAMWorkflowController.Stop]: Stops the associated workflow.
//
// # Getting Workflow Information
//
//   - [IAMWorkflowController.CanRun]: A Boolean value that indicates whether the controller’s workflow is able to run.
//   - [IAMWorkflowController.IsPaused]: A Boolean value that indicates whether the controller’s workflow is currently paused.
//   - [IAMWorkflowController.IsRunning]: A Boolean value that indicates whether the controller’s workflow is currently running.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController
type IAMWorkflowController interface {
	appkit.INSController

	// Topic: Accessing the Workflow

	// The controller’s workflow.
	Workflow() IAMWorkflow
	SetWorkflow(value IAMWorkflow)

	// Topic: Accessing the Workflow View

	// The controller’s workflow view.
	WorkflowView() IAMWorkflowView
	SetWorkflowView(value IAMWorkflowView)

	// Topic: Accessing the Delegate

	// The controller’s delegate.
	Delegate() AMWorkflowControllerDelegate
	SetDelegate(value AMWorkflowControllerDelegate)

	// Topic: Controlling the Workflow

	// Pauses a workflow that’s running.
	Pause(sender objectivec.IObject)
	// Stops a workflow, clears any action results, and resets the workflow back to an un-run state.
	Reset(sender objectivec.IObject)
	// Runs the associated workflow, after first clearing any results stored by its actions during any previous run.
	Run(sender objectivec.IObject)
	// In a paused workflow, runs the next action in the workflow and then pauses again.
	Step(sender objectivec.IObject)
	// Stops the associated workflow.
	Stop(sender objectivec.IObject)

	// Topic: Getting Workflow Information

	// A Boolean value that indicates whether the controller’s workflow is able to run.
	CanRun() bool
	// A Boolean value that indicates whether the controller’s workflow is currently paused.
	IsPaused() bool
	// A Boolean value that indicates whether the controller’s workflow is currently running.
	IsRunning() bool
}

// Init initializes the instance.
func (a AMWorkflowController) Init() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMWorkflowController) Autorelease() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflowController creates a new AMWorkflowController instance.
func NewAMWorkflowController() AMWorkflowController {
	class := getAMWorkflowControllerClass()
	rv := objc.Send[AMWorkflowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Pauses a workflow that’s running.
//
// sender: Object that initiated the pause action.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/pause(_:)
func (a AMWorkflowController) Pause(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("pause:"), sender)
}

// Stops a workflow, clears any action results, and resets the workflow back
// to an un-run state.
//
// sender: Object that initiated the reset action.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/reset(_:)
func (a AMWorkflowController) Reset(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("reset:"), sender)
}

// Runs the associated workflow, after first clearing any results stored by
// its actions during any previous run.
//
// sender: Object that initiated the run action.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/run(_:)
func (a AMWorkflowController) Run(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("run:"), sender)
}

// In a paused workflow, runs the next action in the workflow and then pauses
// again.
//
// sender: Object that initiated the step action.
//
// # Discussion
//
// Stepping allows a workflow to be executed one action at a time. This is
// useful for ensuring that the workflow is doing what it’s supposed to do,
// as the results of each individual action can be inspected before moving on
// to the next.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/step(_:)
func (a AMWorkflowController) Step(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("step:"), sender)
}

// Stops the associated workflow.
//
// sender: Object that initiated the stop action.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/stop(_:)
func (a AMWorkflowController) Stop(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("stop:"), sender)
}

// The controller’s workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflow
func (a AMWorkflowController) Workflow() IAMWorkflow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("workflow"))
	return AMWorkflowFromID(objc.ID(rv))
}
func (a AMWorkflowController) SetWorkflow(value IAMWorkflow) {
	objc.Send[struct{}](a.ID, objc.Sel("setWorkflow:"), value)
}

// The controller’s workflow view.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-swift.property
func (a AMWorkflowController) WorkflowView() IAMWorkflowView {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("workflowView"))
	return AMWorkflowViewFromID(objc.ID(rv))
}
func (a AMWorkflowController) SetWorkflowView(value IAMWorkflowView) {
	objc.Send[struct{}](a.ID, objc.Sel("setWorkflowView:"), value)
}

// The controller’s delegate.
//
// # Return Value
//
// The controller’s delegate.
//
// # Discussion
//
// This object receives updates on the progress and state of the workflow
// controller.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/delegate
func (a AMWorkflowController) Delegate() AMWorkflowControllerDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return AMWorkflowControllerDelegateObjectFromID(rv)
}
func (a AMWorkflowController) SetDelegate(value AMWorkflowControllerDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the controller’s workflow is able
// to run.
//
// # Return Value
//
// true if the controller’s workflow is able to run; false otherwise.
//
// # Discussion
//
// You might use this method to determine when to enable a “Run” button or
// other UI element you use to run the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/canRun
func (a AMWorkflowController) CanRun() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canRun"))
	return rv
}

// A Boolean value that indicates whether the controller’s workflow is
// currently paused.
//
// # Return Value
//
// true if the controller’s workflow is currently paused; false otherwise.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/isPaused
func (a AMWorkflowController) IsPaused() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPaused"))
	return rv
}

// A Boolean value that indicates whether the controller’s workflow is
// currently running.
//
// # Discussion
//
// true if the controller’s workflow is currently running; false otherwise.
// Use “ to determine whether the receiver’s workflow is currently running.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowController/isRunning
func (a AMWorkflowController) IsRunning() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRunning"))
	return rv
}
