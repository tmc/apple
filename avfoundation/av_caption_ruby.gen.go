// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionRuby] class.
var (
	_AVCaptionRubyClass     AVCaptionRubyClass
	_AVCaptionRubyClassOnce sync.Once
)

func getAVCaptionRubyClass() AVCaptionRubyClass {
	_AVCaptionRubyClassOnce.Do(func() {
		_AVCaptionRubyClass = AVCaptionRubyClass{class: objc.GetClass("AVCaptionRuby")}
	})
	return _AVCaptionRubyClass
}

// GetAVCaptionRubyClass returns the class object for AVCaptionRuby.
func GetAVCaptionRubyClass() AVCaptionRubyClass {
	return getAVCaptionRubyClass()
}

type AVCaptionRubyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionRubyClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionRubyClass) Alloc() AVCaptionRuby {
	rv := objc.Send[AVCaptionRuby](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that presents ruby characters.
//
// # Overview
// 
// Ruby characters are small annotations, typically used in Japanese content,
// that render alongside the base text.
//
// # Creating Ruby text
//
//   - [AVCaptionRuby.InitWithText]: Creates ruby text.
//   - [AVCaptionRuby.InitWithTextPositionAlignment]: Creates ruby text with position and alignment.
//
// # Accessing text properties
//
//   - [AVCaptionRuby.Text]: The ruby text.
//   - [AVCaptionRuby.Position]: The ruby text position.
//   - [AVCaptionRuby.Alignment]: The ruby text alignment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby
type AVCaptionRuby struct {
	objectivec.Object
}

// AVCaptionRubyFromID constructs a [AVCaptionRuby] from an objc.ID.
//
// An object that presents ruby characters.
func AVCaptionRubyFromID(id objc.ID) AVCaptionRuby {
	return AVCaptionRuby{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionRuby adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionRuby] class.
//
// # Creating Ruby text
//
//   - [IAVCaptionRuby.InitWithText]: Creates ruby text.
//   - [IAVCaptionRuby.InitWithTextPositionAlignment]: Creates ruby text with position and alignment.
//
// # Accessing text properties
//
//   - [IAVCaptionRuby.Text]: The ruby text.
//   - [IAVCaptionRuby.Position]: The ruby text position.
//   - [IAVCaptionRuby.Alignment]: The ruby text alignment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby
type IAVCaptionRuby interface {
	objectivec.IObject

	// Topic: Creating Ruby text

	// Creates ruby text.
	InitWithText(text string) AVCaptionRuby
	// Creates ruby text with position and alignment.
	InitWithTextPositionAlignment(text string, position AVCaptionRubyPosition, alignment AVCaptionRubyAlignment) AVCaptionRuby

	// Topic: Accessing text properties

	// The ruby text.
	Text() string
	// The ruby text position.
	Position() AVCaptionRubyPosition
	// The ruby text alignment.
	Alignment() AVCaptionRubyAlignment

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c AVCaptionRuby) Init() AVCaptionRuby {
	rv := objc.Send[AVCaptionRuby](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionRuby) Autorelease() AVCaptionRuby {
	rv := objc.Send[AVCaptionRuby](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionRuby creates a new AVCaptionRuby instance.
func NewAVCaptionRuby() AVCaptionRuby {
	class := getAVCaptionRubyClass()
	rv := objc.Send[AVCaptionRuby](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates ruby text.
//
// text: The ruby text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:)
func NewCaptionRubyWithText(text string) AVCaptionRuby {
	instance := getAVCaptionRubyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithText:"), objc.String(text))
	return AVCaptionRubyFromID(rv)
}

// Creates ruby text with position and alignment.
//
// text: The ruby text.
//
// position: The ruby text position.
//
// alignment: The ruby text alignment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:position:alignment:)
func NewCaptionRubyWithTextPositionAlignment(text string, position AVCaptionRubyPosition, alignment AVCaptionRubyAlignment) AVCaptionRuby {
	instance := getAVCaptionRubyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithText:position:alignment:"), objc.String(text), position, alignment)
	return AVCaptionRubyFromID(rv)
}

// Creates ruby text.
//
// text: The ruby text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:)
func (c AVCaptionRuby) InitWithText(text string) AVCaptionRuby {
	rv := objc.Send[AVCaptionRuby](c.ID, objc.Sel("initWithText:"), objc.String(text))
	return rv
}
// Creates ruby text with position and alignment.
//
// text: The ruby text.
//
// position: The ruby text position.
//
// alignment: The ruby text alignment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:position:alignment:)
func (c AVCaptionRuby) InitWithTextPositionAlignment(text string, position AVCaptionRubyPosition, alignment AVCaptionRubyAlignment) AVCaptionRuby {
	rv := objc.Send[AVCaptionRuby](c.ID, objc.Sel("initWithText:position:alignment:"), objc.String(text), position, alignment)
	return rv
}
func (c AVCaptionRuby) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The ruby text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/text
func (c AVCaptionRuby) Text() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
// The ruby text position.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/position
func (c AVCaptionRuby) Position() AVCaptionRubyPosition {
	rv := objc.Send[AVCaptionRubyPosition](c.ID, objc.Sel("position"))
	return AVCaptionRubyPosition(rv)
}
// The ruby text alignment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/alignment
func (c AVCaptionRuby) Alignment() AVCaptionRubyAlignment {
	rv := objc.Send[AVCaptionRubyAlignment](c.ID, objc.Sel("alignment"))
	return AVCaptionRubyAlignment(rv)
}

