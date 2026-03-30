// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioPlayer] class.
var (
	_AVAudioPlayerClass     AVAudioPlayerClass
	_AVAudioPlayerClassOnce sync.Once
)

func getAVAudioPlayerClass() AVAudioPlayerClass {
	_AVAudioPlayerClassOnce.Do(func() {
		_AVAudioPlayerClass = AVAudioPlayerClass{class: objc.GetClass("AVAudioPlayer")}
	})
	return _AVAudioPlayerClass
}

// GetAVAudioPlayerClass returns the class object for AVAudioPlayer.
func GetAVAudioPlayerClass() AVAudioPlayerClass {
	return getAVAudioPlayerClass()
}

type AVAudioPlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPlayerClass) Alloc() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioPlayer.STSLabel]
//   - [AVAudioPlayer.DecodeError]
//   - [AVAudioPlayer.FinishedPlaying]
//   - [AVAudioPlayer.Impl]
//   - [AVAudioPlayer.MixToUplink]
//   - [AVAudioPlayer.PrivRemoveSessionListener]
//   - [AVAudioPlayer.SetMixToUplink]
//   - [AVAudioPlayer.SetSTSLabel]
//   - [AVAudioPlayer.SetUseInjectionDevice]
//   - [AVAudioPlayer.Url]
//   - [AVAudioPlayer.UseInjectionDevice]
//   - [AVAudioPlayer.InitBase]
//   - [AVAudioPlayer.MeteringEnabled]
//   - [AVAudioPlayer.SetMeteringEnabled]
//   - [AVAudioPlayer.Playing]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer
type AVAudioPlayer struct {
	objectivec.Object
}

// AVAudioPlayerFromID constructs a [AVAudioPlayer] from an objc.ID.
func AVAudioPlayerFromID(id objc.ID) AVAudioPlayer {
	return AVAudioPlayer{objectivec.Object{ID: id}}
}

// Ensure AVAudioPlayer implements IAVAudioPlayer.
var _ IAVAudioPlayer = AVAudioPlayer{}

// An interface definition for the [AVAudioPlayer] class.
//
// # Methods
//
//   - [IAVAudioPlayer.STSLabel]
//   - [IAVAudioPlayer.DecodeError]
//   - [IAVAudioPlayer.FinishedPlaying]
//   - [IAVAudioPlayer.Impl]
//   - [IAVAudioPlayer.MixToUplink]
//   - [IAVAudioPlayer.PrivRemoveSessionListener]
//   - [IAVAudioPlayer.SetMixToUplink]
//   - [IAVAudioPlayer.SetSTSLabel]
//   - [IAVAudioPlayer.SetUseInjectionDevice]
//   - [IAVAudioPlayer.Url]
//   - [IAVAudioPlayer.UseInjectionDevice]
//   - [IAVAudioPlayer.InitBase]
//   - [IAVAudioPlayer.MeteringEnabled]
//   - [IAVAudioPlayer.SetMeteringEnabled]
//   - [IAVAudioPlayer.Playing]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer
type IAVAudioPlayer interface {
	objectivec.IObject

	// Topic: Methods

	STSLabel() objectivec.IObject
	DecodeError(error_ objectivec.IObject)
	FinishedPlaying(playing objectivec.IObject)
	Impl() objectivec.IObject
	MixToUplink() bool
	PrivRemoveSessionListener()
	SetMixToUplink(uplink bool)
	SetSTSLabel(sTSLabel objectivec.IObject)
	SetUseInjectionDevice(device bool)
	Url() foundation.INSURL
	UseInjectionDevice() bool
	InitBase() AVAudioPlayer
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	Playing() bool
}

// Init initializes the instance.
func (a AVAudioPlayer) Init() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPlayer) Autorelease() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPlayer creates a new AVAudioPlayer instance.
func NewAVAudioPlayer() AVAudioPlayer {
	class := getAVAudioPlayerClass()
	rv := objc.Send[AVAudioPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/initBase
func NewAudioPlayerBase() AVAudioPlayer {
	instance := getAVAudioPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initBase"))
	return AVAudioPlayerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/STSLabel
func (a AVAudioPlayer) STSLabel() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("STSLabel"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/decodeError:
func (a AVAudioPlayer) DecodeError(error_ objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("decodeError:"), error_)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/finishedPlaying:
func (a AVAudioPlayer) FinishedPlaying(playing objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishedPlaying:"), playing)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/impl
func (a AVAudioPlayer) Impl() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("impl"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/mixToUplink
func (a AVAudioPlayer) MixToUplink() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("mixToUplink"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/privRemoveSessionListener
func (a AVAudioPlayer) PrivRemoveSessionListener() {
	objc.Send[objc.ID](a.ID, objc.Sel("privRemoveSessionListener"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setMixToUplink:
func (a AVAudioPlayer) SetMixToUplink(uplink bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setMixToUplink:"), uplink)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setSTSLabel:
func (a AVAudioPlayer) SetSTSLabel(sTSLabel objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSTSLabel:"), sTSLabel)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setUseInjectionDevice:
func (a AVAudioPlayer) SetUseInjectionDevice(device bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setUseInjectionDevice:"), device)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/useInjectionDevice
func (a AVAudioPlayer) UseInjectionDevice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("useInjectionDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/initBase
func (a AVAudioPlayer) InitBase() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](a.ID, objc.Sel("initBase"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/meteringEnabled
func (a AVAudioPlayer) MeteringEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("meteringEnabled"))
	return rv
}
func (a AVAudioPlayer) SetMeteringEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMeteringEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/playing
func (a AVAudioPlayer) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("playing"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/url
func (a AVAudioPlayer) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
