// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
)

// The class instance for the [SOSRLanguageTableCellView] class.
var (
	_SOSRLanguageTableCellViewClass     SOSRLanguageTableCellViewClass
	_SOSRLanguageTableCellViewClassOnce sync.Once
)

func getSOSRLanguageTableCellViewClass() SOSRLanguageTableCellViewClass {
	_SOSRLanguageTableCellViewClassOnce.Do(func() {
		_SOSRLanguageTableCellViewClass = SOSRLanguageTableCellViewClass{class: objc.GetClass("SOSRLanguageTableCellView")}
	})
	return _SOSRLanguageTableCellViewClass
}

// GetSOSRLanguageTableCellViewClass returns the class object for SOSRLanguageTableCellView.
func GetSOSRLanguageTableCellViewClass() SOSRLanguageTableCellViewClass {
	return getSOSRLanguageTableCellViewClass()
}

type SOSRLanguageTableCellViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguageTableCellViewClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguageTableCellViewClass) Alloc() SOSRLanguageTableCellView {
	rv := objc.Send[SOSRLanguageTableCellView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOSRLanguageTableCellView.ActiveCheckbox]
//   - [SOSRLanguageTableCellView.DownloadMessageTextField]
//   - [SOSRLanguageTableCellView.DownloadVariantPopUpButton]
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageTableCellView
type SOSRLanguageTableCellView struct {
	appkit.NSTableCellView
}

// SOSRLanguageTableCellViewFromID constructs a [SOSRLanguageTableCellView] from an objc.ID.
func SOSRLanguageTableCellViewFromID(id objc.ID) SOSRLanguageTableCellView {
	return SOSRLanguageTableCellView{NSTableCellView: appkit.NSTableCellViewFromID(id)}
}
// Ensure SOSRLanguageTableCellView implements ISOSRLanguageTableCellView.
var _ ISOSRLanguageTableCellView = SOSRLanguageTableCellView{}

// An interface definition for the [SOSRLanguageTableCellView] class.
//
// # Methods
//
//   - [ISOSRLanguageTableCellView.ActiveCheckbox]
//   - [ISOSRLanguageTableCellView.DownloadMessageTextField]
//   - [ISOSRLanguageTableCellView.DownloadVariantPopUpButton]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageTableCellView
type ISOSRLanguageTableCellView interface {
	appkit.INSTableCellView

	// Topic: Methods

	ActiveCheckbox() ISOSRLanguageRowCheckboxButton
	DownloadMessageTextField() appkit.NSTextField
	DownloadVariantPopUpButton() appkit.NSPopUpButton
}

// Init initializes the instance.
func (s SOSRLanguageTableCellView) Init() SOSRLanguageTableCellView {
	rv := objc.Send[SOSRLanguageTableCellView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguageTableCellView) Autorelease() SOSRLanguageTableCellView {
	rv := objc.Send[SOSRLanguageTableCellView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguageTableCellView creates a new SOSRLanguageTableCellView instance.
func NewSOSRLanguageTableCellView() SOSRLanguageTableCellView {
	class := getSOSRLanguageTableCellViewClass()
	rv := objc.Send[SOSRLanguageTableCellView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageTableCellView/activeCheckbox
func (s SOSRLanguageTableCellView) ActiveCheckbox() ISOSRLanguageRowCheckboxButton {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("activeCheckbox"))
	return SOSRLanguageRowCheckboxButtonFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageTableCellView/downloadMessageTextField
func (s SOSRLanguageTableCellView) DownloadMessageTextField() appkit.NSTextField {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadMessageTextField"))
	return appkit.NSTextFieldFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageTableCellView/downloadVariantPopUpButton
func (s SOSRLanguageTableCellView) DownloadVariantPopUpButton() appkit.NSPopUpButton {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadVariantPopUpButton"))
	return appkit.NSPopUpButtonFromID(objc.ID(rv))
}

