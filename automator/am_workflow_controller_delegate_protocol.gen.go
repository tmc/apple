// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods that a delegate of a workflow controller implements.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate
type AMWorkflowControllerDelegate interface {
	objectivec.IObject
}

// AMWorkflowControllerDelegateObject wraps an existing Objective-C object that conforms to the AMWorkflowControllerDelegate protocol.
type AMWorkflowControllerDelegateObject struct {
	objectivec.Object
}

func (o AMWorkflowControllerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AMWorkflowControllerDelegateObjectFromID constructs a [AMWorkflowControllerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AMWorkflowControllerDelegateObjectFromID(id objc.ID) AMWorkflowControllerDelegateObject {
	return AMWorkflowControllerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate when the specified action is about to run.
//
// controller: The controller object sending the message.
//
// action: The workflow action to run.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowController(_:willRun:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerWillRunAction(controller IAMWorkflowController, action IAMAction) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowController:willRunAction:"), controller, action)
}

// Notifies the delegate when the workflow controller object is about to run.
//
// controller: The workflow controller object to run.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowControllerWillRun(_:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerWillRun(controller IAMWorkflowController) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowControllerWillRun:"), controller)
}

// Notifies the delegate when the specified action finishes running.
//
// controller: The controller object sending the message.
//
// action: The workflow action that ran.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowController(_:didRun:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerDidRunAction(controller IAMWorkflowController, action IAMAction) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowController:didRunAction:"), controller, action)
}

// Notifies the delegate when the workflow controller object finishes running.
//
// controller: The workflow controller object that finished running.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowControllerDidRun(_:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerDidRun(controller IAMWorkflowController) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowControllerDidRun:"), controller)
}

// Tells the delegate that the workflow controller object is about to stop.
//
// controller: The workflow controller object to stop.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowControllerWillStop(_:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerWillStop(controller IAMWorkflowController) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowControllerWillStop:"), controller)
}

// Tells the delegate that the workflow controller object has stopped.
//
// controller: The workflow controller object that stopped.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowControllerDidStop(_:)
func (o AMWorkflowControllerDelegateObject) WorkflowControllerDidStop(controller IAMWorkflowController) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowControllerDidStop:"), controller)
}

// Notifies the delegate when the workflow encounters an error.
//
// controller: The controller object sending the message.
//
// error: If a workflow error occurs, upon return contains an instance of [NSError]
// that describes the problem.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowControllerDelegate/workflowController(_:didError:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (o AMWorkflowControllerDelegateObject) WorkflowControllerDidError(controller IAMWorkflowController, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("workflowController:didError:"), controller, error_)
}

// AMWorkflowControllerDelegateConfig holds optional typed callbacks for [AMWorkflowControllerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/automator/amworkflowcontrollerdelegate
type AMWorkflowControllerDelegateConfig struct {

	// Preparing to Run
	// WorkflowControllerWillRun — Notifies the delegate when the workflow controller object is about to run.
	WorkflowControllerWillRun func(controller AMWorkflowController)

	// Running
	// WorkflowControllerDidRun — Notifies the delegate when the workflow controller object finishes running.
	WorkflowControllerDidRun func(controller AMWorkflowController)

	// Stopping
	// WorkflowControllerWillStop — Tells the delegate that the workflow controller object is about to stop.
	WorkflowControllerWillStop func(controller AMWorkflowController)
	// WorkflowControllerDidStop — Tells the delegate that the workflow controller object has stopped.
	WorkflowControllerDidStop func(controller AMWorkflowController)

	// Handling Errors
	// WorkflowControllerDidError — Notifies the delegate when the workflow encounters an error.
	WorkflowControllerDidError func(controller AMWorkflowController, error_ foundation.NSError)

	// Other Methods
	// WorkflowControllerWillRunAction — Notifies the delegate when the specified action is about to run.
	WorkflowControllerWillRunAction func(controller AMWorkflowController, action AMAction)
	// WorkflowControllerDidRunAction — Notifies the delegate when the specified action finishes running.
	WorkflowControllerDidRunAction func(controller AMWorkflowController, action AMAction)
}

// NewAMWorkflowControllerDelegate creates an Objective-C object implementing the [AMWorkflowControllerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AMWorkflowControllerDelegateObject] satisfies the [AMWorkflowControllerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/automator/amworkflowcontrollerdelegate
func NewAMWorkflowControllerDelegate(config AMWorkflowControllerDelegateConfig) AMWorkflowControllerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAMWorkflowControllerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WorkflowControllerWillRunAction != nil {
		fn := config.WorkflowControllerWillRunAction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowController:willRunAction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID, actionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowController:willRunAction:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				action := AMActionFromID(actionID)
				fn(controller, action)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerWillRun != nil {
		fn := config.WorkflowControllerWillRun
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowControllerWillRun:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowControllerWillRun:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerDidRunAction != nil {
		fn := config.WorkflowControllerDidRunAction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowController:didRunAction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID, actionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowController:didRunAction:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				action := AMActionFromID(actionID)
				fn(controller, action)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerDidRun != nil {
		fn := config.WorkflowControllerDidRun
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowControllerDidRun:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowControllerDidRun:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerWillStop != nil {
		fn := config.WorkflowControllerWillStop
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowControllerWillStop:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowControllerWillStop:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerDidStop != nil {
		fn := config.WorkflowControllerDidStop
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowControllerDidStop:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowControllerDidStop:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	if config.WorkflowControllerDidError != nil {
		fn := config.WorkflowControllerDidError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("workflowController:didError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AMWorkflowControllerDelegate", "workflowController:didError:")
					}
				}()
				controller := AMWorkflowControllerFromID(controllerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(controller, error_)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AMWorkflowControllerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAMWorkflowControllerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AMWorkflowControllerDelegateObjectFromID(instance)
}
