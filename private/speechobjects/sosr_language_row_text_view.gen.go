// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
)

// The class instance for the [SOSRLanguageRowTextView] class.
var (
	_SOSRLanguageRowTextViewClass     SOSRLanguageRowTextViewClass
	_SOSRLanguageRowTextViewClassOnce sync.Once
)

func getSOSRLanguageRowTextViewClass() SOSRLanguageRowTextViewClass {
	_SOSRLanguageRowTextViewClassOnce.Do(func() {
		_SOSRLanguageRowTextViewClass = SOSRLanguageRowTextViewClass{class: objc.GetClass("SOSRLanguageRowTextView")}
	})
	return _SOSRLanguageRowTextViewClass
}

// GetSOSRLanguageRowTextViewClass returns the class object for SOSRLanguageRowTextView.
func GetSOSRLanguageRowTextViewClass() SOSRLanguageRowTextViewClass {
	return getSOSRLanguageRowTextViewClass()
}

type SOSRLanguageRowTextViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguageRowTextViewClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguageRowTextViewClass) Alloc() SOSRLanguageRowTextView {
	rv := objc.Send[SOSRLanguageRowTextView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRowTextView
type SOSRLanguageRowTextView struct {
	appkit.NSTextField
}

// SOSRLanguageRowTextViewFromID constructs a [SOSRLanguageRowTextView] from an objc.ID.
func SOSRLanguageRowTextViewFromID(id objc.ID) SOSRLanguageRowTextView {
	return SOSRLanguageRowTextView{NSTextField: appkit.NSTextFieldFromID(id)}
}
// Ensure SOSRLanguageRowTextView implements ISOSRLanguageRowTextView.
var _ ISOSRLanguageRowTextView = SOSRLanguageRowTextView{}

// An interface definition for the [SOSRLanguageRowTextView] class.
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRowTextView
type ISOSRLanguageRowTextView interface {
	appkit.INSTextField
}

// Init initializes the instance.
func (s SOSRLanguageRowTextView) Init() SOSRLanguageRowTextView {
	rv := objc.Send[SOSRLanguageRowTextView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguageRowTextView) Autorelease() SOSRLanguageRowTextView {
	rv := objc.Send[SOSRLanguageRowTextView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguageRowTextView creates a new SOSRLanguageRowTextView instance.
func NewSOSRLanguageRowTextView() SOSRLanguageRowTextView {
	class := getSOSRLanguageRowTextViewClass()
	rv := objc.Send[SOSRLanguageRowTextView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

