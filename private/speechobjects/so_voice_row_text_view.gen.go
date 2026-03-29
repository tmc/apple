// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
)

// The class instance for the [SOVoiceRowTextView] class.
var (
	_SOVoiceRowTextViewClass     SOVoiceRowTextViewClass
	_SOVoiceRowTextViewClassOnce sync.Once
)

func getSOVoiceRowTextViewClass() SOVoiceRowTextViewClass {
	_SOVoiceRowTextViewClassOnce.Do(func() {
		_SOVoiceRowTextViewClass = SOVoiceRowTextViewClass{class: objc.GetClass("SOVoiceRowTextView")}
	})
	return _SOVoiceRowTextViewClass
}

// GetSOVoiceRowTextViewClass returns the class object for SOVoiceRowTextView.
func GetSOVoiceRowTextViewClass() SOVoiceRowTextViewClass {
	return getSOVoiceRowTextViewClass()
}

type SOVoiceRowTextViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOVoiceRowTextViewClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOVoiceRowTextViewClass) Alloc() SOVoiceRowTextView {
	rv := objc.Send[SOVoiceRowTextView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceRowTextView
type SOVoiceRowTextView struct {
	appkit.NSTextField
}

// SOVoiceRowTextViewFromID constructs a [SOVoiceRowTextView] from an objc.ID.
func SOVoiceRowTextViewFromID(id objc.ID) SOVoiceRowTextView {
	return SOVoiceRowTextView{NSTextField: appkit.NSTextFieldFromID(id)}
}
// Ensure SOVoiceRowTextView implements ISOVoiceRowTextView.
var _ ISOVoiceRowTextView = SOVoiceRowTextView{}

// An interface definition for the [SOVoiceRowTextView] class.
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceRowTextView
type ISOVoiceRowTextView interface {
	appkit.INSTextField
}

// Init initializes the instance.
func (s SOVoiceRowTextView) Init() SOVoiceRowTextView {
	rv := objc.Send[SOVoiceRowTextView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOVoiceRowTextView) Autorelease() SOVoiceRowTextView {
	rv := objc.Send[SOVoiceRowTextView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOVoiceRowTextView creates a new SOVoiceRowTextView instance.
func NewSOVoiceRowTextView() SOVoiceRowTextView {
	class := getSOVoiceRowTextViewClass()
	rv := objc.Send[SOVoiceRowTextView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

