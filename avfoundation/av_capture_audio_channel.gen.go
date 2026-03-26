// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureAudioChannel] class.
var (
	_AVCaptureAudioChannelClass     AVCaptureAudioChannelClass
	_AVCaptureAudioChannelClassOnce sync.Once
)

func getAVCaptureAudioChannelClass() AVCaptureAudioChannelClass {
	_AVCaptureAudioChannelClassOnce.Do(func() {
		_AVCaptureAudioChannelClass = AVCaptureAudioChannelClass{class: objc.GetClass("AVCaptureAudioChannel")}
	})
	return _AVCaptureAudioChannelClass
}

// GetAVCaptureAudioChannelClass returns the class object for AVCaptureAudioChannel.
func GetAVCaptureAudioChannelClass() AVCaptureAudioChannelClass {
	return getAVCaptureAudioChannelClass()
}

type AVCaptureAudioChannelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureAudioChannelClass) Alloc() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that monitors average and peak power levels for an audio channel
// in a capture connection.
//
// # Overview
// 
// You don’t create instances of this class directly. Instead, an
// [AVCaptureConnection] object that connects an audio input to an audio
// output provides an array of [AVCaptureAudioChannel] objects, one for each
// channel of audio available. You can poll for audio levels by iterating
// through these audio channel objects.
//
// # Configuring a channel
//
//   - [AVCaptureAudioChannel.Enabled]: A Boolean value that indicates whether the channel is in an enabled state.
//   - [AVCaptureAudioChannel.SetEnabled]
//   - [AVCaptureAudioChannel.Volume]: The current volume (gain) of the channel.
//   - [AVCaptureAudioChannel.SetVolume]
//
// # Accessing power levels
//
//   - [AVCaptureAudioChannel.AveragePowerLevel]: The instantaneous average power level in decibels.
//   - [AVCaptureAudioChannel.PeakHoldLevel]: The peak hold power level in decibels.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel
type AVCaptureAudioChannel struct {
	objectivec.Object
}

// AVCaptureAudioChannelFromID constructs a [AVCaptureAudioChannel] from an objc.ID.
//
// An object that monitors average and peak power levels for an audio channel
// in a capture connection.
func AVCaptureAudioChannelFromID(id objc.ID) AVCaptureAudioChannel {
	return AVCaptureAudioChannel{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureAudioChannel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureAudioChannel] class.
//
// # Configuring a channel
//
//   - [IAVCaptureAudioChannel.Enabled]: A Boolean value that indicates whether the channel is in an enabled state.
//   - [IAVCaptureAudioChannel.SetEnabled]
//   - [IAVCaptureAudioChannel.Volume]: The current volume (gain) of the channel.
//   - [IAVCaptureAudioChannel.SetVolume]
//
// # Accessing power levels
//
//   - [IAVCaptureAudioChannel.AveragePowerLevel]: The instantaneous average power level in decibels.
//   - [IAVCaptureAudioChannel.PeakHoldLevel]: The peak hold power level in decibels.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel
type IAVCaptureAudioChannel interface {
	objectivec.IObject

	// Topic: Configuring a channel

	// A Boolean value that indicates whether the channel is in an enabled state.
	Enabled() bool
	SetEnabled(value bool)
	// The current volume (gain) of the channel.
	Volume() float32
	SetVolume(value float32)

	// Topic: Accessing power levels

	// The instantaneous average power level in decibels.
	AveragePowerLevel() float32
	// The peak hold power level in decibels.
	PeakHoldLevel() float32

	// The connections between inputs and outputs that a capture session contains.
	Connections() IAVCaptureConnection
	SetConnections(value IAVCaptureConnection)
}

// Init initializes the instance.
func (c AVCaptureAudioChannel) Init() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureAudioChannel) Autorelease() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAudioChannel creates a new AVCaptureAudioChannel instance.
func NewAVCaptureAudioChannel() AVCaptureAudioChannel {
	class := getAVCaptureAudioChannelClass()
	rv := objc.Send[AVCaptureAudioChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the channel is in an enabled state.
//
// # Discussion
// 
// By default, a connection enables all audio channels that it exposes. You
// can set this value to [false] to stop the flow of data for a particular
// channel.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/isEnabled
func (c AVCaptureAudioChannel) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c AVCaptureAudioChannel) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}
// The current volume (gain) of the channel.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/volume
func (c AVCaptureAudioChannel) Volume() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("volume"))
	return rv
}
func (c AVCaptureAudioChannel) SetVolume(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setVolume:"), value)
}
// The instantaneous average power level in decibels.
//
// # Discussion
// 
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/averagePowerLevel
func (c AVCaptureAudioChannel) AveragePowerLevel() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("averagePowerLevel"))
	return rv
}
// The peak hold power level in decibels.
//
// # Discussion
// 
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/peakHoldLevel
func (c AVCaptureAudioChannel) PeakHoldLevel() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("peakHoldLevel"))
	return rv
}
// The connections between inputs and outputs that a capture session contains.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturesession/connections
func (c AVCaptureAudioChannel) Connections() IAVCaptureConnection {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("connections"))
	return AVCaptureConnectionFromID(objc.ID(rv))
}
func (c AVCaptureAudioChannel) SetConnections(value IAVCaptureConnection) {
	objc.Send[struct{}](c.ID, objc.Sel("setConnections:"), value)
}

