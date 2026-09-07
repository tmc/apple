// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[AVAudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioFormat.Interleaved]
//   - [AVAudioFormat.Standard]
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
//   - [IAVAudioFormat.Interleaved]
//   - [IAVAudioFormat.Standard]
type IAVAudioFormat interface {
	objectivec.IObject

	// Topic: Methods

	Interleaved() bool
	Standard() bool
}

// Init initializes the instance.
func (a AVAudioFormat) Init() AVAudioFormat {
	rv := objc.SendIfResponds[AVAudioFormat](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioFormat) Autorelease() AVAudioFormat {
	rv := objc.SendIfResponds[AVAudioFormat](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioFormat creates a new AVAudioFormat instance.
func NewAVAudioFormat() AVAudioFormat {
	class := getAVAudioFormatClass()
	rv := objc.SendIfResponds[AVAudioFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_AVAudioFormatClass AVAudioFormatClass) FormatWithInvalidSampleRateAndChannelCount() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVAudioFormatClass.class), objc.Sel("formatWithInvalidSampleRateAndChannelCount"))
	return objectivec.Object{ID: rv}
}
func (_AVAudioFormatClass AVAudioFormatClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioFormatClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a AVAudioFormat) Interleaved() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("interleaved"))
	return rv
}
func (a AVAudioFormat) Standard() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("standard"))
	return rv
}
