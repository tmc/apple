// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AMWorkflowView] class.
var (
	_AMWorkflowViewClass     AMWorkflowViewClass
	_AMWorkflowViewClassOnce sync.Once
)

func getAMWorkflowViewClass() AMWorkflowViewClass {
	_AMWorkflowViewClassOnce.Do(func() {
		_AMWorkflowViewClass = AMWorkflowViewClass{class: objc.GetClass("AMWorkflowView")}
	})
	return _AMWorkflowViewClass
}

// GetAMWorkflowViewClass returns the class object for AMWorkflowView.
func GetAMWorkflowViewClass() AMWorkflowViewClass {
	return getAMWorkflowViewClass()
}

type AMWorkflowViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AMWorkflowViewClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AMWorkflowViewClass) Alloc() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that lets you view and edit Automator workflows in your app.
//
// # Overview
//
// A workflow view displays an instance of [AMWorkflow].
//
// You can use Interface Builder to add an instance of [AMWorkflowView] to a
// window in your app. You can then add an [AMWorkflowView] object to the nib
// window and use the controller’s [AMWorkflowController.WorkflowView]
// outlet to connect it to the workflow view. The controller object also has
// [AMWorkflowController.Run] and [AMWorkflowController.Stop] actions that can
// be connected to buttons or other user interface elements.
//
// # Configuring the Workflow View
//
//   - [AMWorkflowView.IsEditable]: A Boolean value that indicates whether the workflow view is editable.
//   - [AMWorkflowView.SetEditable]
//   - [AMWorkflowView.WorkflowController]: The view’s workflow controller.
//   - [AMWorkflowView.SetWorkflowController]
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowView
type AMWorkflowView struct {
	appkit.NSView
}

// AMWorkflowViewFromID constructs a [AMWorkflowView] from an objc.ID.
//
// An object that lets you view and edit Automator workflows in your app.
func AMWorkflowViewFromID(id objc.ID) AMWorkflowView {
	return AMWorkflowView{NSView: appkit.NSViewFromID(id)}
}

// NOTE: AMWorkflowView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AMWorkflowView] class.
//
// # Configuring the Workflow View
//
//   - [IAMWorkflowView.IsEditable]: A Boolean value that indicates whether the workflow view is editable.
//   - [IAMWorkflowView.SetEditable]
//   - [IAMWorkflowView.WorkflowController]: The view’s workflow controller.
//   - [IAMWorkflowView.SetWorkflowController]
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowView
type IAMWorkflowView interface {
	appkit.INSView

	// Topic: Configuring the Workflow View

	// A Boolean value that indicates whether the workflow view is editable.
	IsEditable() bool
	SetEditable(value bool)
	// The view’s workflow controller.
	WorkflowController() IAMWorkflowController
	SetWorkflowController(value IAMWorkflowController)
}

// Init initializes the instance.
func (a AMWorkflowView) Init() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AMWorkflowView) Autorelease() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflowView creates a new AMWorkflowView instance.
func NewAMWorkflowView() AMWorkflowView {
	class := getAMWorkflowViewClass()
	rv := objc.Send[AMWorkflowView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the workflow view is editable.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowView/isEditable
func (a AMWorkflowView) IsEditable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEditable"))
	return rv
}
func (a AMWorkflowView) SetEditable(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setEditable:"), value)
}

// The view’s workflow controller.
//
// See: https://developer.apple.com/documentation/Automator/AMWorkflowView/workflowController
func (a AMWorkflowView) WorkflowController() IAMWorkflowController {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("workflowController"))
	return AMWorkflowControllerFromID(objc.ID(rv))
}
func (a AMWorkflowView) SetWorkflowController(value IAMWorkflowController) {
	objc.Send[struct{}](a.ID, objc.Sel("setWorkflowController:"), value)
}
