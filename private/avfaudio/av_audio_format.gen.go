// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioFormat] class.
var (
	_AVAudioFormatClass     AVAudioFormatClass
	_AVAudioFormatClassOnce sync.Once
)

func getAVAudioFormatClass() AVAudioFormatClass {
	_AVAudioFormatClassOnce.Do(func() {
		_AVAudioFormatClass = AVAudioFormatClass{class: objc.GetClass("AVAudioFormat")}
	})
	return _AVAudioFormatClass
}

// GetAVAudioFormatClass returns the class object for AVAudioFormat.
func GetAVAudioFormatClass() AVAudioFormatClass {
	return getAVAudioFormatClass()
}

type AVAudioFormatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioFormatClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioFormatClass) Alloc() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioFormat.InitWithCoder]
//   - [AVAudioFormat.Interleaved]
//   - [AVAudioFormat.Standard]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
type AVAudioFormat struct {
	objectivec.Object
}

// AVAudioFormatFromID constructs a [AVAudioFormat] from an objc.ID.
func AVAudioFormatFromID(id objc.ID) AVAudioFormat {
	return AVAudioFormat{objectivec.Object{ID: id}}
}
// Ensure AVAudioFormat implements IAVAudioFormat.
var _ IAVAudioFormat = AVAudioFormat{}

// An interface definition for the [AVAudioFormat] class.
//
// # Methods
//
//   - [IAVAudioFormat.InitWithCoder]
//   - [IAVAudioFormat.Interleaved]
//   - [IAVAudioFormat.Standard]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
type IAVAudioFormat interface {
	objectivec.IObject

	// Topic: Methods

	InitWithCoder(coder foundation.INSCoder) AVAudioFormat
	Interleaved() bool
	Standard() bool
}

// Init initializes the instance.
func (a AVAudioFormat) Init() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioFormat) Autorelease() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioFormat creates a new AVAudioFormat instance.
func NewAVAudioFormat() AVAudioFormat {
	class := getAVAudioFormatClass()
	rv := objc.Send[AVAudioFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/initWithCoder:
func NewAudioFormatWithCoder(coder objectivec.IObject) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVAudioFormatFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/initWithCoder:
func (a AVAudioFormat) InitWithCoder(coder foundation.INSCoder) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/formatWithInvalidSampleRateAndChannelCount
func (_AVAudioFormatClass AVAudioFormatClass) FormatWithInvalidSampleRateAndChannelCount() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioFormatClass.class), objc.Sel("formatWithInvalidSampleRateAndChannelCount"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/supportsSecureCoding
func (_AVAudioFormatClass AVAudioFormatClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioFormatClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/interleaved
func (a AVAudioFormat) Interleaved() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("interleaved"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/standard
func (a AVAudioFormat) Standard() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("standard"))
	return rv
}

