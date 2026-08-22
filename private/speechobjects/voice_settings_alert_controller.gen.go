// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VoiceSettingsAlertController] class.
var (
	_VoiceSettingsAlertControllerClass     VoiceSettingsAlertControllerClass
	_VoiceSettingsAlertControllerClassOnce sync.Once
)

func getVoiceSettingsAlertControllerClass() VoiceSettingsAlertControllerClass {
	_VoiceSettingsAlertControllerClassOnce.Do(func() {
		_VoiceSettingsAlertControllerClass = VoiceSettingsAlertControllerClass{class: objc.GetClass("VoiceSettingsAlertController")}
	})
	return _VoiceSettingsAlertControllerClass
}

// GetVoiceSettingsAlertControllerClass returns the class object for VoiceSettingsAlertController.
func GetVoiceSettingsAlertControllerClass() VoiceSettingsAlertControllerClass {
	return getVoiceSettingsAlertControllerClass()
}

type VoiceSettingsAlertControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceSettingsAlertControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceSettingsAlertControllerClass) Alloc() VoiceSettingsAlertController {
	rv := objc.SendIfResponds[VoiceSettingsAlertController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VoiceSettingsAlertController.ShowWindowWithVoiceSettingsModalDelegateWindowTitle]
type VoiceSettingsAlertController struct {
	VoiceSettingsWindowController
}

// VoiceSettingsAlertControllerFromID constructs a [VoiceSettingsAlertController] from an objc.ID.
func VoiceSettingsAlertControllerFromID(id objc.ID) VoiceSettingsAlertController {
	return VoiceSettingsAlertController{VoiceSettingsWindowController: VoiceSettingsWindowControllerFromID(id)}
}

// Ensure VoiceSettingsAlertController implements IVoiceSettingsAlertController.
var _ IVoiceSettingsAlertController = VoiceSettingsAlertController{}

// An interface definition for the [VoiceSettingsAlertController] class.
//
// # Methods
//
//   - [IVoiceSettingsAlertController.ShowWindowWithVoiceSettingsModalDelegateWindowTitle]
type IVoiceSettingsAlertController interface {
	IVoiceSettingsWindowController

	// Topic: Methods

	ShowWindowWithVoiceSettingsModalDelegateWindowTitle(settings objectivec.IObject, delegate objectivec.IObject, title objectivec.IObject)
}

// Init initializes the instance.
func (v VoiceSettingsAlertController) Init() VoiceSettingsAlertController {
	rv := objc.SendIfResponds[VoiceSettingsAlertController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceSettingsAlertController) Autorelease() VoiceSettingsAlertController {
	rv := objc.SendIfResponds[VoiceSettingsAlertController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceSettingsAlertController creates a new VoiceSettingsAlertController instance.
func NewVoiceSettingsAlertController() VoiceSettingsAlertController {
	class := getVoiceSettingsAlertControllerClass()
	rv := objc.SendIfResponds[VoiceSettingsAlertController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VoiceSettingsAlertController) ShowWindowWithVoiceSettingsModalDelegateWindowTitle(settings objectivec.IObject, delegate objectivec.IObject, title objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("showWindowWithVoiceSettings:modalDelegate:windowTitle:"), settings, delegate, title)
}

func (_VoiceSettingsAlertControllerClass VoiceSettingsAlertControllerClass) DefaultVoiceSettingsAlertController() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VoiceSettingsAlertControllerClass.class), objc.Sel("defaultVoiceSettingsAlertController"))
	return objectivec.Object{ID: rv}
}
