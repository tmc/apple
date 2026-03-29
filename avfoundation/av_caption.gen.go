// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaption] class.
var (
	_AVCaptionClass     AVCaptionClass
	_AVCaptionClassOnce sync.Once
)

func getAVCaptionClass() AVCaptionClass {
	_AVCaptionClassOnce.Do(func() {
		_AVCaptionClass = AVCaptionClass{class: objc.GetClass("AVCaption")}
	})
	return _AVCaptionClass
}

// GetAVCaptionClass returns the class object for AVCaption.
func GetAVCaptionClass() AVCaptionClass {
	return getAVCaptionClass()
}

type AVCaptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionClass) Alloc() AVCaption {
	rv := objc.Send[AVCaption](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents text to present over a time range.
//
// # Overview
// 
// A caption contains a cue, which is a single sentence or paragraph of text
// for a time range in the video timeline. Within the active range, the
// caption may animate (for example, Karaoke lyrics) by rolling-up, changing
// visibility, or using other dynamic styling.
//
// # Creating a caption
//
//   - [AVCaption.InitWithTextTimeRange]: Creates a caption that contains text and a time range.
//
// # Accessing text and timing
//
//   - [AVCaption.Text]: The caption text.
//   - [AVCaption.TimeRange]: The time range over which the system presents the caption.
//
// # Accessing the region
//
//   - [AVCaption.Region]: The region in which the caption exists.
//
// # Accessing alignment
//
//   - [AVCaption.TextAlignment]: The alignment for the caption text.
//
// # Accessing animation
//
//   - [AVCaption.Animation]: The animation that the system applies to this caption.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption
type AVCaption struct {
	objectivec.Object
}

// AVCaptionFromID constructs a [AVCaption] from an objc.ID.
//
// An object that represents text to present over a time range.
func AVCaptionFromID(id objc.ID) AVCaption {
	return AVCaption{objectivec.Object{ID: id}}
}
// NOTE: AVCaption adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaption] class.
//
// # Creating a caption
//
//   - [IAVCaption.InitWithTextTimeRange]: Creates a caption that contains text and a time range.
//
// # Accessing text and timing
//
//   - [IAVCaption.Text]: The caption text.
//   - [IAVCaption.TimeRange]: The time range over which the system presents the caption.
//
// # Accessing the region
//
//   - [IAVCaption.Region]: The region in which the caption exists.
//
// # Accessing alignment
//
//   - [IAVCaption.TextAlignment]: The alignment for the caption text.
//
// # Accessing animation
//
//   - [IAVCaption.Animation]: The animation that the system applies to this caption.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption
type IAVCaption interface {
	objectivec.IObject

	// Topic: Creating a caption

	// Creates a caption that contains text and a time range.
	InitWithTextTimeRange(text string, timeRange coremedia.CMTimeRange) AVCaption

	// Topic: Accessing text and timing

	// The caption text.
	Text() string
	// The time range over which the system presents the caption.
	TimeRange() coremedia.CMTimeRange

	// Topic: Accessing the region

	// The region in which the caption exists.
	Region() IAVCaptionRegion

	// Topic: Accessing alignment

	// The alignment for the caption text.
	TextAlignment() AVCaptionTextAlignment

	// Topic: Accessing animation

	// The animation that the system applies to this caption.
	Animation() AVCaptionAnimation

	// Returns the background color at the index position.
	BackgroundColorAtIndexRange(index int, outRange *foundation.NSRange) coregraphics.CGColorRef
	// Returns the text decoration at the index position.
	DecorationAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionDecoration
	// Returns the font style and range at the index position.
	FontStyleAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionFontStyle
	// Returns the font weight and range at the index position.
	FontWeightAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionFontWeight
	// Returns the ruby text at the index position.
	RubyAtIndexRange(index int, outRange *foundation.NSRange) IAVCaptionRuby
	// Returns the text color at the index position.
	TextColorAtIndexRange(index int, outRange *foundation.NSRange) coregraphics.CGColorRef
	// Returns the text combine at the index position.
	TextCombineAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionTextCombine
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c AVCaption) Init() AVCaption {
	rv := objc.Send[AVCaption](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaption) Autorelease() AVCaption {
	rv := objc.Send[AVCaption](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaption creates a new AVCaption instance.
func NewAVCaption() AVCaption {
	class := getAVCaptionClass()
	rv := objc.Send[AVCaption](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a caption that contains text and a time range.
//
// text: The text the caption displays.
//
// timeRange: The range of time when the caption is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/init(_:timeRange:)
func NewCaptionWithTextTimeRange(text string, timeRange coremedia.CMTimeRange) AVCaption {
	instance := getAVCaptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithText:timeRange:"), objc.String(text), timeRange)
	return AVCaptionFromID(rv)
}

// Creates a caption that contains text and a time range.
//
// text: The text the caption displays.
//
// timeRange: The range of time when the caption is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/init(_:timeRange:)
func (c AVCaption) InitWithTextTimeRange(text string, timeRange coremedia.CMTimeRange) AVCaption {
	rv := objc.Send[AVCaption](c.ID, objc.Sel("initWithText:timeRange:"), objc.String(text), timeRange)
	return rv
}
// Returns the background color at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned background color
// applies.
//
// # Return Value
// 
// The background color.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/backgroundColorAtIndex:range:
func (c AVCaption) BackgroundColorAtIndexRange(index int, outRange *foundation.NSRange) coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](c.ID, objc.Sel("backgroundColorAtIndex:range:"), index, outRange)
	return coregraphics.CGColorRef(rv)
}
// Returns the text decoration at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer to store the range to which the returned decoration applies.
//
// # Return Value
// 
// The text decoration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/decorationAtIndex:range:
func (c AVCaption) DecorationAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionDecoration {
	rv := objc.Send[AVCaptionDecoration](c.ID, objc.Sel("decorationAtIndex:range:"), index, outRange)
	return AVCaptionDecoration(rv)
}
// Returns the font style and range at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned style applies.
//
// # Return Value
// 
// The font style.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/fontStyleAtIndex:range:
func (c AVCaption) FontStyleAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionFontStyle {
	rv := objc.Send[AVCaptionFontStyle](c.ID, objc.Sel("fontStyleAtIndex:range:"), index, outRange)
	return AVCaptionFontStyle(rv)
}
// Returns the font weight and range at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned weight applies.
//
// # Return Value
// 
// The font weight.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/fontWeightAtIndex:range:
func (c AVCaption) FontWeightAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionFontWeight {
	rv := objc.Send[AVCaptionFontWeight](c.ID, objc.Sel("fontWeightAtIndex:range:"), index, outRange)
	return AVCaptionFontWeight(rv)
}
// Returns the ruby text at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned ruby text applies.
//
// # Return Value
// 
// The ruby text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/rubyAtIndex:range:
func (c AVCaption) RubyAtIndexRange(index int, outRange *foundation.NSRange) IAVCaptionRuby {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rubyAtIndex:range:"), index, outRange)
	return AVCaptionRubyFromID(rv)
}
// Returns the text color at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned text color applies.
//
// # Return Value
// 
// The text color.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/textColorAtIndex:range:
func (c AVCaption) TextColorAtIndexRange(index int, outRange *foundation.NSRange) coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](c.ID, objc.Sel("textColorAtIndex:range:"), index, outRange)
	return coregraphics.CGColorRef(rv)
}
// Returns the text combine at the index position.
//
// index: A character position in the caption text.
//
// outRange: A pointer that stores the range to which the returned text combine applies.
//
// # Return Value
// 
// The text combine.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/textCombineAtIndex:range:
func (c AVCaption) TextCombineAtIndexRange(index int, outRange *foundation.NSRange) AVCaptionTextCombine {
	rv := objc.Send[AVCaptionTextCombine](c.ID, objc.Sel("textCombineAtIndex:range:"), index, outRange)
	return AVCaptionTextCombine(rv)
}
func (c AVCaption) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The caption text.
//
// # Discussion
// 
// Apple iTT format text supports the same Unicode code points that an XML
// document supports. The system converts XML special characters like ‘&’
// to corresponding character-reference syntax when it writes the caption to a
// file. The text may contain the line-breaking character sequences LF, CR, or
// CF+LF.
// 
// CEA608 closed captions support the following Unicode characters. They
// don’t support the line-breaking character sequences LF, CR, or CF+LF.
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/text
func (c AVCaption) Text() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
// The time range over which the system presents the caption.
//
// # Discussion
// 
// Apple iTT format only permits captions to have overlapping time ranges if
// they’re associated with different regions.
// 
// CEA608 closed caption time ranges can’t start with zero, because the
// decoder needs transmission time. Align time ranges with the video frame
// rate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/timeRange
func (c AVCaption) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](c.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}
// The region in which the caption exists.
//
// # Discussion
// 
// This property is `nil` when the underlying caption format doesn’t support
// or use regions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/region
func (c AVCaption) Region() IAVCaptionRegion {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("region"))
	return AVCaptionRegionFromID(objc.ID(rv))
}
// The alignment for the caption text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/textAlignment-swift.property
func (c AVCaption) TextAlignment() AVCaptionTextAlignment {
	rv := objc.Send[AVCaptionTextAlignment](c.ID, objc.Sel("textAlignment"))
	return AVCaptionTextAlignment(rv)
}
// The animation that the system applies to this caption.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/animation-swift.property
func (c AVCaption) Animation() AVCaptionAnimation {
	rv := objc.Send[AVCaptionAnimation](c.ID, objc.Sel("animation"))
	return AVCaptionAnimation(rv)
}

