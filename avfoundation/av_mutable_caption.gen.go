// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableCaption] class.
var (
	_AVMutableCaptionClass     AVMutableCaptionClass
	_AVMutableCaptionClassOnce sync.Once
)

func getAVMutableCaptionClass() AVMutableCaptionClass {
	_AVMutableCaptionClassOnce.Do(func() {
		_AVMutableCaptionClass = AVMutableCaptionClass{class: objc.GetClass("AVMutableCaption")}
	})
	return _AVMutableCaptionClass
}

// GetAVMutableCaptionClass returns the class object for AVMutableCaption.
func GetAVMutableCaptionClass() AVMutableCaptionClass {
	return getAVMutableCaptionClass()
}

type AVMutableCaptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableCaptionClass) Alloc() AVMutableCaption {
	rv := objc.Send[AVMutableCaption](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable caption subclass that you use to create new captions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption
type AVMutableCaption struct {
	AVCaption
}

// AVMutableCaptionFromID constructs a [AVMutableCaption] from an objc.ID.
//
// A mutable caption subclass that you use to create new captions.
func AVMutableCaptionFromID(id objc.ID) AVMutableCaption {
	return AVMutableCaption{AVCaption: AVCaptionFromID(id)}
}
// NOTE: AVMutableCaption adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableCaption] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption
type IAVMutableCaption interface {
	IAVCaption

	// Removes a background color from a range of text.
	RemoveBackgroundColorInRange(range_ foundation.NSRange)
	// Removes a decoration from a range of text.
	RemoveDecorationInRange(range_ foundation.NSRange)
	// Removes a font style from a range of text.
	RemoveFontStyleInRange(range_ foundation.NSRange)
	// Removes a font weight from a range of text.
	RemoveFontWeightInRange(range_ foundation.NSRange)
	// Removes ruby text from a range.
	RemoveRubyInRange(range_ foundation.NSRange)
	// Removes the text color for a range of text.
	RemoveTextColorInRange(range_ foundation.NSRange)
	// Removes text combine from a range of text.
	RemoveTextCombineInRange(range_ foundation.NSRange)
	// Sets the background color for a range of text.
	SetBackgroundColorInRange(color coregraphics.CGColorRef, range_ foundation.NSRange)
	// Sets a decoration for a range of text.
	SetDecorationInRange(decoration AVCaptionDecoration, range_ foundation.NSRange)
	// Sets the font style for a range of text.
	SetFontStyleInRange(fontStyle AVCaptionFontStyle, range_ foundation.NSRange)
	// Sets the font weight for a range of text.
	SetFontWeightInRange(fontWeight AVCaptionFontWeight, range_ foundation.NSRange)
	// Sets ruby text for a range.
	SetRubyInRange(ruby IAVCaptionRuby, range_ foundation.NSRange)
	// Sets the text color for a range of text.
	SetTextColorInRange(color coregraphics.CGColorRef, range_ foundation.NSRange)
	// Sets text combine for a range.
	SetTextCombineInRange(textCombine AVCaptionTextCombine, range_ foundation.NSRange)
}

// Init initializes the instance.
func (m AVMutableCaption) Init() AVMutableCaption {
	rv := objc.Send[AVMutableCaption](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableCaption) Autorelease() AVMutableCaption {
	rv := objc.Send[AVMutableCaption](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableCaption creates a new AVMutableCaption instance.
func NewAVMutableCaption() AVMutableCaption {
	class := getAVMutableCaptionClass()
	rv := objc.Send[AVMutableCaption](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a caption that contains text and a time range.
//
// text: The text the caption displays.
//
// timeRange: The range of time when the caption is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/init(_:timeRange:)
// timeRange is a [coremedia.CMTimeRange].
func NewMutableCaptionWithTextTimeRange(text string, timeRange objectivec.IObject) AVMutableCaption {
	instance := getAVMutableCaptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithText:timeRange:"), objc.String(text), timeRange)
	return AVMutableCaptionFromID(rv)
}

// Removes a background color from a range of text.
//
// range: The range from which the system removes the background color.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeBackgroundColorInRange:
func (m AVMutableCaption) RemoveBackgroundColorInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeBackgroundColorInRange:"), range_)
}
// Removes a decoration from a range of text.
//
// range: The range from which the system removes the decoration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeDecorationInRange:
func (m AVMutableCaption) RemoveDecorationInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeDecorationInRange:"), range_)
}
// Removes a font style from a range of text.
//
// range: The range from which to remove the style.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeFontStyleInRange:
func (m AVMutableCaption) RemoveFontStyleInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeFontStyleInRange:"), range_)
}
// Removes a font weight from a range of text.
//
// range: The range from which to remove the style.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeFontWeightInRange:
func (m AVMutableCaption) RemoveFontWeightInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeFontWeightInRange:"), range_)
}
// Removes ruby text from a range.
//
// range: The range from which the system removes the ruby text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeRubyInRange:
func (m AVMutableCaption) RemoveRubyInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeRubyInRange:"), range_)
}
// Removes the text color for a range of text.
//
// range: The range from which the system removes the text color.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeTextColorInRange:
func (m AVMutableCaption) RemoveTextColorInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTextColorInRange:"), range_)
}
// Removes text combine from a range of text.
//
// range: The range from which the system removes text combine.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/removeTextCombineInRange:
func (m AVMutableCaption) RemoveTextCombineInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTextCombineInRange:"), range_)
}
// Sets the background color for a range of text.
//
// color: The background color.
//
// range: The range to which this background color applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setBackgroundColor:inRange:
func (m AVMutableCaption) SetBackgroundColorInRange(color coregraphics.CGColorRef, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setBackgroundColor:inRange:"), color, range_)
}
// Sets a decoration for a range of text.
//
// decoration: The decoration.
//
// range: The range to which this decoration applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setDecoration:inRange:
func (m AVMutableCaption) SetDecorationInRange(decoration AVCaptionDecoration, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setDecoration:inRange:"), decoration, range_)
}
// Sets the font style for a range of text.
//
// fontStyle: The font style.
//
// range: The range to which this style applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setFontStyle:inRange:
func (m AVMutableCaption) SetFontStyleInRange(fontStyle AVCaptionFontStyle, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setFontStyle:inRange:"), fontStyle, range_)
}
// Sets the font weight for a range of text.
//
// fontWeight: The font weight.
//
// range: The range to which this font weight applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setFontWeight:inRange:
func (m AVMutableCaption) SetFontWeightInRange(fontWeight AVCaptionFontWeight, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setFontWeight:inRange:"), fontWeight, range_)
}
// Sets ruby text for a range.
//
// ruby: The ruby text.
//
// range: The range to which the ruby text applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setRuby:inRange:
func (m AVMutableCaption) SetRubyInRange(ruby IAVCaptionRuby, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setRuby:inRange:"), ruby, range_)
}
// Sets the text color for a range of text.
//
// color: The text color.
//
// range: The range to which this text color applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setTextColor:inRange:
func (m AVMutableCaption) SetTextColorInRange(color coregraphics.CGColorRef, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setTextColor:inRange:"), color, range_)
}
// Sets text combine for a range.
//
// textCombine: The text combine.
//
// range: The range to which the text combine applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaption/setTextCombine:inRange:
func (m AVMutableCaption) SetTextCombineInRange(textCombine AVCaptionTextCombine, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setTextCombine:inRange:"), textCombine, range_)
}

