// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AMWorkspace] class.
var (
	_AMWorkspaceClass     AMWorkspaceClass
	_AMWorkspaceClassOnce sync.Once
)

func getAMWorkspaceClass() AMWorkspaceClass {
	_AMWorkspaceClassOnce.Do(func() {
		_AMWorkspaceClass = AMWorkspaceClass{class: objc.GetClass("AMWorkspace")}
	})
	return _AMWorkspaceClass
}

// GetAMWorkspaceClass returns the class object for AMWorkspace.
func GetAMWorkspaceClass() AMWorkspaceClass {
	return getAMWorkspaceClass()
}

type AMWorkspaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMWorkspaceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMWorkspaceClass) Alloc() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A workspace for running an Automator workflow.
//
// # Overview
//
// The [AMWorkspace] class provides access to the shared workspace in the
// Automator framework, where you can run workflows without a workflow
// controller. Use [AMWorkspaceClass.SharedWorkspace] to access the shared
// workspace and [AMWorkspace.RunWorkflowAtPathWithInputError] to run your
// workflow in it.
//
// # Running Workflows
//
//   - [AMWorkspace.RunWorkflowAtPathWithInputError]: Loads and runs the specified workflow file.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkspace
type AMWorkspace struct {
	objectivec.Object
}

// AMWorkspaceFromID constructs a [AMWorkspace] from an objc.ID.
//
// A workspace for running an Automator workflow.
func AMWorkspaceFromID(id objc.ID) AMWorkspace {
	return AMWorkspace{objectivec.Object{ID: id}}
}

// NOTE: AMWorkspace adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMWorkspace] class.
//
// # Running Workflows
//
//   - [IAMWorkspace.RunWorkflowAtPathWithInputError]: Loads and runs the specified workflow file.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkspace
type IAMWorkspace interface {
	objectivec.IObject

	// Topic: Running Workflows

	// Loads and runs the specified workflow file.
	RunWorkflowAtPathWithInputError(path string, input objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (a AMWorkspace) Init() AMWorkspace {
	rv := objc.Send[AMWorkspace](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMWorkspace) Autorelease() AMWorkspace {
	rv := objc.Send[AMWorkspace](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkspace creates a new AMWorkspace instance.
func NewAMWorkspace() AMWorkspace {
	class := getAMWorkspaceClass()
	rv := objc.Send[AMWorkspace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads and runs the specified workflow file.
//
// path: A path that specifies the location of the workflow file.
//
// input: The input for the first action in the workflow. Pass `nil` if the first
// action doesn’t need input.
//
// # Return Value
//
// `nil` if an error occurs or if the action completes successfully without
// producing output; otherwise, the output of the last action in the workflow.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkspace/runWorkflow(atPath:withInput:)
func (a AMWorkspace) RunWorkflowAtPathWithInputError(path string, input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("runWorkflowAtPath:withInput:error:"), objc.String(path), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// The shared workspace object.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkspace/shared
func (_AMWorkspaceClass AMWorkspaceClass) SharedWorkspace() AMWorkspace {
	rv := objc.Send[objc.ID](objc.ID(_AMWorkspaceClass.class), objc.Sel("sharedWorkspace"))
	return AMWorkspaceFromID(objc.ID(rv))
}
