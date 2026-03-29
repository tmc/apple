// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/avfaudio"
)

// The class instance for the [TTSAudioFormat] class.
var (
	_TTSAudioFormatClass     TTSAudioFormatClass
	_TTSAudioFormatClassOnce sync.Once
)

func getTTSAudioFormatClass() TTSAudioFormatClass {
	_TTSAudioFormatClassOnce.Do(func() {
		_TTSAudioFormatClass = TTSAudioFormatClass{class: objc.GetClass("TTSAudioFormat")}
	})
	return _TTSAudioFormatClass
}

// GetTTSAudioFormatClass returns the class object for TTSAudioFormat.
func GetTTSAudioFormatClass() TTSAudioFormatClass {
	return getTTSAudioFormatClass()
}

type TTSAudioFormatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAudioFormatClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAudioFormatClass) Alloc() TTSAudioFormat {
	rv := objc.Send[TTSAudioFormat](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSAudioFormat.AvFormat]
//   - [TTSAudioFormat.ChannelCount]
//   - [TTSAudioFormat.ChannelLayoutTag]
//   - [TTSAudioFormat.SetChannelLayoutTag]
//   - [TTSAudioFormat.SampleRate]
//   - [TTSAudioFormat.StreamDescription]
//   - [TTSAudioFormat.SetStreamDescription]
//   - [TTSAudioFormat.InitWithStreamDescription]
//   - [TTSAudioFormat.InitWithStreamDescriptionChannelLayoutTag]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat
type TTSAudioFormat struct {
	objectivec.Object
}

// TTSAudioFormatFromID constructs a [TTSAudioFormat] from an objc.ID.
func TTSAudioFormatFromID(id objc.ID) TTSAudioFormat {
	return TTSAudioFormat{objectivec.Object{ID: id}}
}
// Ensure TTSAudioFormat implements ITTSAudioFormat.
var _ ITTSAudioFormat = TTSAudioFormat{}

// An interface definition for the [TTSAudioFormat] class.
//
// # Methods
//
//   - [ITTSAudioFormat.AvFormat]
//   - [ITTSAudioFormat.ChannelCount]
//   - [ITTSAudioFormat.ChannelLayoutTag]
//   - [ITTSAudioFormat.SetChannelLayoutTag]
//   - [ITTSAudioFormat.SampleRate]
//   - [ITTSAudioFormat.StreamDescription]
//   - [ITTSAudioFormat.SetStreamDescription]
//   - [ITTSAudioFormat.InitWithStreamDescription]
//   - [ITTSAudioFormat.InitWithStreamDescriptionChannelLayoutTag]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat
type ITTSAudioFormat interface {
	objectivec.IObject

	// Topic: Methods

	AvFormat() avfaudio.AVAudioFormat
	ChannelCount() uint32
	ChannelLayoutTag() uint32
	SetChannelLayoutTag(value uint32)
	SampleRate() float64
	StreamDescription() objectivec.IObject
	SetStreamDescription(value objectivec.IObject)
	InitWithStreamDescription(description objectivec.IObject) TTSAudioFormat
	InitWithStreamDescriptionChannelLayoutTag(description objectivec.IObject, tag uint32) TTSAudioFormat
}

// Init initializes the instance.
func (t TTSAudioFormat) Init() TTSAudioFormat {
	rv := objc.Send[TTSAudioFormat](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAudioFormat) Autorelease() TTSAudioFormat {
	rv := objc.Send[TTSAudioFormat](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAudioFormat creates a new TTSAudioFormat instance.
func NewTTSAudioFormat() TTSAudioFormat {
	class := getTTSAudioFormatClass()
	rv := objc.Send[TTSAudioFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/initWithStreamDescription:
func NewTTSAudioFormatWithStreamDescription(description objectivec.IObject) TTSAudioFormat {
	instance := getTTSAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamDescription:"), description)
	return TTSAudioFormatFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/initWithStreamDescription:channelLayoutTag:
func NewTTSAudioFormatWithStreamDescriptionChannelLayoutTag(description objectivec.IObject, tag uint32) TTSAudioFormat {
	instance := getTTSAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamDescription:channelLayoutTag:"), description, tag)
	return TTSAudioFormatFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/initWithStreamDescription:
func (t TTSAudioFormat) InitWithStreamDescription(description objectivec.IObject) TTSAudioFormat {
	rv := objc.Send[TTSAudioFormat](t.ID, objc.Sel("initWithStreamDescription:"), description)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/initWithStreamDescription:channelLayoutTag:
func (t TTSAudioFormat) InitWithStreamDescriptionChannelLayoutTag(description objectivec.IObject, tag uint32) TTSAudioFormat {
	rv := objc.Send[TTSAudioFormat](t.ID, objc.Sel("initWithStreamDescription:channelLayoutTag:"), description, tag)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/avFormat
func (t TTSAudioFormat) AvFormat() avfaudio.AVAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("avFormat"))
	return avfaudio.AVAudioFormatFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/channelCount
func (t TTSAudioFormat) ChannelCount() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("channelCount"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/channelLayoutTag
func (t TTSAudioFormat) ChannelLayoutTag() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("channelLayoutTag"))
	return rv
}
func (t TTSAudioFormat) SetChannelLayoutTag(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannelLayoutTag:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/sampleRate
func (t TTSAudioFormat) SampleRate() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("sampleRate"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioFormat/streamDescription
func (t TTSAudioFormat) StreamDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("streamDescription"))
	return objectivec.Object{ID: rv}
}
func (t TTSAudioFormat) SetStreamDescription(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setStreamDescription:"), value)
}

