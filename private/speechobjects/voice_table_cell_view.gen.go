// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
)

// The class instance for the [VoiceTableCellView] class.
var (
	_VoiceTableCellViewClass     VoiceTableCellViewClass
	_VoiceTableCellViewClassOnce sync.Once
)

func getVoiceTableCellViewClass() VoiceTableCellViewClass {
	_VoiceTableCellViewClassOnce.Do(func() {
		_VoiceTableCellViewClass = VoiceTableCellViewClass{class: objc.GetClass("VoiceTableCellView")}
	})
	return _VoiceTableCellViewClass
}

// GetVoiceTableCellViewClass returns the class object for VoiceTableCellView.
func GetVoiceTableCellViewClass() VoiceTableCellViewClass {
	return getVoiceTableCellViewClass()
}

type VoiceTableCellViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceTableCellViewClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceTableCellViewClass) Alloc() VoiceTableCellView {
	rv := objc.Send[VoiceTableCellView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VoiceTableCellView.ActiveCheckbox]
//   - [VoiceTableCellView.DownloadCheckbox]
//   - [VoiceTableCellView.DownloadMessageTextField]
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableCellView
type VoiceTableCellView struct {
	appkit.NSTableCellView
}

// VoiceTableCellViewFromID constructs a [VoiceTableCellView] from an objc.ID.
func VoiceTableCellViewFromID(id objc.ID) VoiceTableCellView {
	return VoiceTableCellView{NSTableCellView: appkit.NSTableCellViewFromID(id)}
}
// Ensure VoiceTableCellView implements IVoiceTableCellView.
var _ IVoiceTableCellView = VoiceTableCellView{}

// An interface definition for the [VoiceTableCellView] class.
//
// # Methods
//
//   - [IVoiceTableCellView.ActiveCheckbox]
//   - [IVoiceTableCellView.DownloadCheckbox]
//   - [IVoiceTableCellView.DownloadMessageTextField]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableCellView
type IVoiceTableCellView interface {
	appkit.INSTableCellView

	// Topic: Methods

	ActiveCheckbox() ISOVoiceRowCheckboxButton
	DownloadCheckbox() ISOVoiceRowCheckboxButton
	DownloadMessageTextField() appkit.NSTextField
}

// Init initializes the instance.
func (v VoiceTableCellView) Init() VoiceTableCellView {
	rv := objc.Send[VoiceTableCellView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceTableCellView) Autorelease() VoiceTableCellView {
	rv := objc.Send[VoiceTableCellView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceTableCellView creates a new VoiceTableCellView instance.
func NewVoiceTableCellView() VoiceTableCellView {
	class := getVoiceTableCellViewClass()
	rv := objc.Send[VoiceTableCellView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableCellView/activeCheckbox
func (v VoiceTableCellView) ActiveCheckbox() ISOVoiceRowCheckboxButton {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("activeCheckbox"))
	return SOVoiceRowCheckboxButtonFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableCellView/downloadCheckbox
func (v VoiceTableCellView) DownloadCheckbox() ISOVoiceRowCheckboxButton {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("downloadCheckbox"))
	return SOVoiceRowCheckboxButtonFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableCellView/downloadMessageTextField
func (v VoiceTableCellView) DownloadMessageTextField() appkit.NSTextField {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("downloadMessageTextField"))
	return appkit.NSTextFieldFromID(objc.ID(rv))
}

