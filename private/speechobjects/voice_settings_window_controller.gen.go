// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VoiceSettingsWindowController] class.
var (
	_VoiceSettingsWindowControllerClass     VoiceSettingsWindowControllerClass
	_VoiceSettingsWindowControllerClassOnce sync.Once
)

func getVoiceSettingsWindowControllerClass() VoiceSettingsWindowControllerClass {
	_VoiceSettingsWindowControllerClassOnce.Do(func() {
		_VoiceSettingsWindowControllerClass = VoiceSettingsWindowControllerClass{class: objc.GetClass("VoiceSettingsWindowController")}
	})
	return _VoiceSettingsWindowControllerClass
}

// GetVoiceSettingsWindowControllerClass returns the class object for VoiceSettingsWindowController.
func GetVoiceSettingsWindowControllerClass() VoiceSettingsWindowControllerClass {
	return getVoiceSettingsWindowControllerClass()
}

type VoiceSettingsWindowControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceSettingsWindowControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceSettingsWindowControllerClass) Alloc() VoiceSettingsWindowController {
	rv := objc.Send[VoiceSettingsWindowController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VoiceSettingsWindowController.CancelVoiceSettings]
//   - [VoiceSettingsWindowController.CloseWindow]
//   - [VoiceSettingsWindowController.PercentOfNormalFromSlider]
//   - [VoiceSettingsWindowController.PlayStopVoiceSettings]
//   - [VoiceSettingsWindowController.RateCheckboxClicked]
//   - [VoiceSettingsWindowController.RateSliderChanged]
//   - [VoiceSettingsWindowController.SaveVoiceSettings]
//   - [VoiceSettingsWindowController.SetUpWindowWithVoiceSettingsModalDelegate]
//   - [VoiceSettingsWindowController.SetValueOfSliderUsingPercentOfNormal]
//   - [VoiceSettingsWindowController.SetValueOfSliderUsingWordsPerMinute]
//   - [VoiceSettingsWindowController.SpeechSynthesizerDidFinishSpeaking]
//   - [VoiceSettingsWindowController.VoicePopupMenuChanged]
//   - [VoiceSettingsWindowController.VoiceSettingsFromWindow]
//   - [VoiceSettingsWindowController.VolumeCheckboxClicked]
//   - [VoiceSettingsWindowController.VolumeSliderChanged]
//   - [VoiceSettingsWindowController.WordsPerMinuteFromSlider]
//   - [VoiceSettingsWindowController.DebugDescription]
//   - [VoiceSettingsWindowController.Description]
//   - [VoiceSettingsWindowController.Hash]
//   - [VoiceSettingsWindowController.Superclass]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController
type VoiceSettingsWindowController struct {
	appkit.NSWindowController
}

// VoiceSettingsWindowControllerFromID constructs a [VoiceSettingsWindowController] from an objc.ID.
func VoiceSettingsWindowControllerFromID(id objc.ID) VoiceSettingsWindowController {
	return VoiceSettingsWindowController{NSWindowController: appkit.NSWindowControllerFromID(id)}
}

// Ensure VoiceSettingsWindowController implements IVoiceSettingsWindowController.
var _ IVoiceSettingsWindowController = VoiceSettingsWindowController{}

// An interface definition for the [VoiceSettingsWindowController] class.
//
// # Methods
//
//   - [IVoiceSettingsWindowController.CancelVoiceSettings]
//   - [IVoiceSettingsWindowController.CloseWindow]
//   - [IVoiceSettingsWindowController.PercentOfNormalFromSlider]
//   - [IVoiceSettingsWindowController.PlayStopVoiceSettings]
//   - [IVoiceSettingsWindowController.RateCheckboxClicked]
//   - [IVoiceSettingsWindowController.RateSliderChanged]
//   - [IVoiceSettingsWindowController.SaveVoiceSettings]
//   - [IVoiceSettingsWindowController.SetUpWindowWithVoiceSettingsModalDelegate]
//   - [IVoiceSettingsWindowController.SetValueOfSliderUsingPercentOfNormal]
//   - [IVoiceSettingsWindowController.SetValueOfSliderUsingWordsPerMinute]
//   - [IVoiceSettingsWindowController.SpeechSynthesizerDidFinishSpeaking]
//   - [IVoiceSettingsWindowController.VoicePopupMenuChanged]
//   - [IVoiceSettingsWindowController.VoiceSettingsFromWindow]
//   - [IVoiceSettingsWindowController.VolumeCheckboxClicked]
//   - [IVoiceSettingsWindowController.VolumeSliderChanged]
//   - [IVoiceSettingsWindowController.WordsPerMinuteFromSlider]
//   - [IVoiceSettingsWindowController.DebugDescription]
//   - [IVoiceSettingsWindowController.Description]
//   - [IVoiceSettingsWindowController.Hash]
//   - [IVoiceSettingsWindowController.Superclass]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController
type IVoiceSettingsWindowController interface {
	appkit.INSWindowController

	// Topic: Methods

	CancelVoiceSettings(settings objectivec.IObject)
	CloseWindow()
	PercentOfNormalFromSlider(slider objectivec.IObject) float32
	PlayStopVoiceSettings(settings objectivec.IObject)
	RateCheckboxClicked(clicked objectivec.IObject)
	RateSliderChanged(changed objectivec.IObject)
	SaveVoiceSettings(settings objectivec.IObject)
	SetUpWindowWithVoiceSettingsModalDelegate(settings objectivec.IObject, delegate objectivec.IObject)
	SetValueOfSliderUsingPercentOfNormal(slider objectivec.IObject, normal float32)
	SetValueOfSliderUsingWordsPerMinute(slider objectivec.IObject, minute float32)
	SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool)
	VoicePopupMenuChanged(changed objectivec.IObject)
	VoiceSettingsFromWindow() objectivec.IObject
	VolumeCheckboxClicked(clicked objectivec.IObject)
	VolumeSliderChanged(changed objectivec.IObject)
	WordsPerMinuteFromSlider(slider objectivec.IObject) float32
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VoiceSettingsWindowController) Init() VoiceSettingsWindowController {
	rv := objc.Send[VoiceSettingsWindowController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceSettingsWindowController) Autorelease() VoiceSettingsWindowController {
	rv := objc.Send[VoiceSettingsWindowController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceSettingsWindowController creates a new VoiceSettingsWindowController instance.
func NewVoiceSettingsWindowController() VoiceSettingsWindowController {
	class := getVoiceSettingsWindowControllerClass()
	rv := objc.Send[VoiceSettingsWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/cancelVoiceSettings:
func (v VoiceSettingsWindowController) CancelVoiceSettings(settings objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("cancelVoiceSettings:"), settings)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/closeWindow
func (v VoiceSettingsWindowController) CloseWindow() {
	objc.Send[objc.ID](v.ID, objc.Sel("closeWindow"))
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/percentOfNormalFromSlider:
func (v VoiceSettingsWindowController) PercentOfNormalFromSlider(slider objectivec.IObject) float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("percentOfNormalFromSlider:"), slider)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/playStopVoiceSettings:
func (v VoiceSettingsWindowController) PlayStopVoiceSettings(settings objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("playStopVoiceSettings:"), settings)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/rateCheckboxClicked:
func (v VoiceSettingsWindowController) RateCheckboxClicked(clicked objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rateCheckboxClicked:"), clicked)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/rateSliderChanged:
func (v VoiceSettingsWindowController) RateSliderChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rateSliderChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/saveVoiceSettings:
func (v VoiceSettingsWindowController) SaveVoiceSettings(settings objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("saveVoiceSettings:"), settings)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/setUpWindowWithVoiceSettings:modalDelegate:
func (v VoiceSettingsWindowController) SetUpWindowWithVoiceSettingsModalDelegate(settings objectivec.IObject, delegate objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setUpWindowWithVoiceSettings:modalDelegate:"), settings, delegate)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/setValueOfSlider:usingPercentOfNormal:
func (v VoiceSettingsWindowController) SetValueOfSliderUsingPercentOfNormal(slider objectivec.IObject, normal float32) {
	objc.Send[objc.ID](v.ID, objc.Sel("setValueOfSlider:usingPercentOfNormal:"), slider, normal)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/setValueOfSlider:usingWordsPerMinute:
func (v VoiceSettingsWindowController) SetValueOfSliderUsingWordsPerMinute(slider objectivec.IObject, minute float32) {
	objc.Send[objc.ID](v.ID, objc.Sel("setValueOfSlider:usingWordsPerMinute:"), slider, minute)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/speechSynthesizer:didFinishSpeaking:
func (v VoiceSettingsWindowController) SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), synthesizer, speaking)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/voicePopupMenuChanged:
func (v VoiceSettingsWindowController) VoicePopupMenuChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("voicePopupMenuChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/voiceSettingsFromWindow
func (v VoiceSettingsWindowController) VoiceSettingsFromWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("voiceSettingsFromWindow"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/volumeCheckboxClicked:
func (v VoiceSettingsWindowController) VolumeCheckboxClicked(clicked objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("volumeCheckboxClicked:"), clicked)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/volumeSliderChanged:
func (v VoiceSettingsWindowController) VolumeSliderChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("volumeSliderChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/wordsPerMinuteFromSlider:
func (v VoiceSettingsWindowController) WordsPerMinuteFromSlider(slider objectivec.IObject) float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("wordsPerMinuteFromSlider:"), slider)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/debugDescription
func (v VoiceSettingsWindowController) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/description
func (v VoiceSettingsWindowController) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/hash
func (v VoiceSettingsWindowController) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceSettingsWindowController/superclass
func (v VoiceSettingsWindowController) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
