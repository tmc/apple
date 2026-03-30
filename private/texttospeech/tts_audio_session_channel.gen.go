// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAudioSessionChannel] class.
var (
	_TTSAudioSessionChannelClass     TTSAudioSessionChannelClass
	_TTSAudioSessionChannelClassOnce sync.Once
)

func getTTSAudioSessionChannelClass() TTSAudioSessionChannelClass {
	_TTSAudioSessionChannelClassOnce.Do(func() {
		_TTSAudioSessionChannelClass = TTSAudioSessionChannelClass{class: objc.GetClass("TTSAudioSessionChannel")}
	})
	return _TTSAudioSessionChannelClass
}

// GetTTSAudioSessionChannelClass returns the class object for TTSAudioSessionChannel.
func GetTTSAudioSessionChannelClass() TTSAudioSessionChannelClass {
	return getTTSAudioSessionChannelClass()
}

type TTSAudioSessionChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAudioSessionChannelClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAudioSessionChannelClass) Alloc() TTSAudioSessionChannel {
	rv := objc.Send[TTSAudioSessionChannel](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAudioSessionChannel.Channel]
//   - [TTSAudioSessionChannel.SetChannel]
//   - [TTSAudioSessionChannel.ChannelLabel]
//   - [TTSAudioSessionChannel.SetChannelLabel]
//   - [TTSAudioSessionChannel.ChannelName]
//   - [TTSAudioSessionChannel.SetChannelName]
//   - [TTSAudioSessionChannel.ChannelNumber]
//   - [TTSAudioSessionChannel.SetChannelNumber]
//   - [TTSAudioSessionChannel.OwningPortUID]
//   - [TTSAudioSessionChannel.SetOwningPortUID]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel
type TTSAudioSessionChannel struct {
	objectivec.Object
}

// TTSAudioSessionChannelFromID constructs a [TTSAudioSessionChannel] from an objc.ID.
func TTSAudioSessionChannelFromID(id objc.ID) TTSAudioSessionChannel {
	return TTSAudioSessionChannel{objectivec.Object{ID: id}}
}

// Ensure TTSAudioSessionChannel implements ITTSAudioSessionChannel.
var _ ITTSAudioSessionChannel = TTSAudioSessionChannel{}

// An interface definition for the [TTSAudioSessionChannel] class.
//
// # Methods
//
//   - [ITTSAudioSessionChannel.Channel]
//   - [ITTSAudioSessionChannel.SetChannel]
//   - [ITTSAudioSessionChannel.ChannelLabel]
//   - [ITTSAudioSessionChannel.SetChannelLabel]
//   - [ITTSAudioSessionChannel.ChannelName]
//   - [ITTSAudioSessionChannel.SetChannelName]
//   - [ITTSAudioSessionChannel.ChannelNumber]
//   - [ITTSAudioSessionChannel.SetChannelNumber]
//   - [ITTSAudioSessionChannel.OwningPortUID]
//   - [ITTSAudioSessionChannel.SetOwningPortUID]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel
type ITTSAudioSessionChannel interface {
	objectivec.IObject

	// Topic: Methods

	Channel() unsafe.Pointer
	SetChannel(value unsafe.Pointer)
	ChannelLabel() uint32
	SetChannelLabel(value uint32)
	ChannelName() string
	SetChannelName(value string)
	ChannelNumber() uint64
	SetChannelNumber(value uint64)
	OwningPortUID() string
	SetOwningPortUID(value string)
}

// Init initializes the instance.
func (t TTSAudioSessionChannel) Init() TTSAudioSessionChannel {
	rv := objc.Send[TTSAudioSessionChannel](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAudioSessionChannel) Autorelease() TTSAudioSessionChannel {
	rv := objc.Send[TTSAudioSessionChannel](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAudioSessionChannel creates a new TTSAudioSessionChannel instance.
func NewTTSAudioSessionChannel() TTSAudioSessionChannel {
	class := getTTSAudioSessionChannelClass()
	rv := objc.Send[TTSAudioSessionChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/channelWithChannel:
func (_TTSAudioSessionChannelClass TTSAudioSessionChannelClass) ChannelWithChannel(channel objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSAudioSessionChannelClass.class), objc.Sel("channelWithChannel:"), channel)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/convertChannels:
func (_TTSAudioSessionChannelClass TTSAudioSessionChannelClass) ConvertChannels(channels objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSAudioSessionChannelClass.class), objc.Sel("convertChannels:"), channels)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/channel
func (t TTSAudioSessionChannel) Channel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("channel"))
	return rv
}
func (t TTSAudioSessionChannel) SetChannel(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannel:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/channelLabel
func (t TTSAudioSessionChannel) ChannelLabel() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("channelLabel"))
	return rv
}
func (t TTSAudioSessionChannel) SetChannelLabel(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannelLabel:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/channelName
func (t TTSAudioSessionChannel) ChannelName() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("channelName"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAudioSessionChannel) SetChannelName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannelName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/channelNumber
func (t TTSAudioSessionChannel) ChannelNumber() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("channelNumber"))
	return rv
}
func (t TTSAudioSessionChannel) SetChannelNumber(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannelNumber:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioSessionChannel/owningPortUID
func (t TTSAudioSessionChannel) OwningPortUID() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("owningPortUID"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAudioSessionChannel) SetOwningPortUID(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setOwningPortUID:"), objc.String(value))
}
