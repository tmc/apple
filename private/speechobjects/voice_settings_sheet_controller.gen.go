// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VoiceSettingsSheetController] class.
var (
	_VoiceSettingsSheetControllerClass     VoiceSettingsSheetControllerClass
	_VoiceSettingsSheetControllerClassOnce sync.Once
)

func getVoiceSettingsSheetControllerClass() VoiceSettingsSheetControllerClass {
	_VoiceSettingsSheetControllerClassOnce.Do(func() {
		_VoiceSettingsSheetControllerClass = VoiceSettingsSheetControllerClass{class: objc.GetClass("VoiceSettingsSheetController")}
	})
	return _VoiceSettingsSheetControllerClass
}

// GetVoiceSettingsSheetControllerClass returns the class object for VoiceSettingsSheetController.
func GetVoiceSettingsSheetControllerClass() VoiceSettingsSheetControllerClass {
	return getVoiceSettingsSheetControllerClass()
}

type VoiceSettingsSheetControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceSettingsSheetControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceSettingsSheetControllerClass) Alloc() VoiceSettingsSheetController {
	rv := objc.Send[VoiceSettingsSheetController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VoiceSettingsSheetController.ShowSheetWithVoiceSettingsModalDelegateModalForWindow]
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsSheetController
type VoiceSettingsSheetController struct {
	VoiceSettingsWindowController
}

// VoiceSettingsSheetControllerFromID constructs a [VoiceSettingsSheetController] from an objc.ID.
func VoiceSettingsSheetControllerFromID(id objc.ID) VoiceSettingsSheetController {
	return VoiceSettingsSheetController{VoiceSettingsWindowController: VoiceSettingsWindowControllerFromID(id)}
}
// Ensure VoiceSettingsSheetController implements IVoiceSettingsSheetController.
var _ IVoiceSettingsSheetController = VoiceSettingsSheetController{}

// An interface definition for the [VoiceSettingsSheetController] class.
//
// # Methods
//
//   - [IVoiceSettingsSheetController.ShowSheetWithVoiceSettingsModalDelegateModalForWindow]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsSheetController
type IVoiceSettingsSheetController interface {
	IVoiceSettingsWindowController

	// Topic: Methods

	ShowSheetWithVoiceSettingsModalDelegateModalForWindow(settings objectivec.IObject, delegate objectivec.IObject, window objectivec.IObject)
}

// Init initializes the instance.
func (v VoiceSettingsSheetController) Init() VoiceSettingsSheetController {
	rv := objc.Send[VoiceSettingsSheetController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceSettingsSheetController) Autorelease() VoiceSettingsSheetController {
	rv := objc.Send[VoiceSettingsSheetController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceSettingsSheetController creates a new VoiceSettingsSheetController instance.
func NewVoiceSettingsSheetController() VoiceSettingsSheetController {
	class := getVoiceSettingsSheetControllerClass()
	rv := objc.Send[VoiceSettingsSheetController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsSheetController/showSheetWithVoiceSettings:modalDelegate:modalForWindow:
func (v VoiceSettingsSheetController) ShowSheetWithVoiceSettingsModalDelegateModalForWindow(settings objectivec.IObject, delegate objectivec.IObject, window objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("showSheetWithVoiceSettings:modalDelegate:modalForWindow:"), settings, delegate, window)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsSheetController/defaultVoiceSettingsSheetController
func (_VoiceSettingsSheetControllerClass VoiceSettingsSheetControllerClass) DefaultVoiceSettingsSheetController() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VoiceSettingsSheetControllerClass.class), objc.Sel("defaultVoiceSettingsSheetController"))
	return objectivec.Object{ID: rv}
}

