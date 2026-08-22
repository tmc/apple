// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVPlayerItemSampleBufferOutputAudioConfiguration] class.
var (
	_AVPlayerItemSampleBufferOutputAudioConfigurationClass     AVPlayerItemSampleBufferOutputAudioConfigurationClass
	_AVPlayerItemSampleBufferOutputAudioConfigurationClassOnce sync.Once
)

func getAVPlayerItemSampleBufferOutputAudioConfigurationClass() AVPlayerItemSampleBufferOutputAudioConfigurationClass {
	_AVPlayerItemSampleBufferOutputAudioConfigurationClassOnce.Do(func() {
		_AVPlayerItemSampleBufferOutputAudioConfigurationClass = AVPlayerItemSampleBufferOutputAudioConfigurationClass{class: objc.GetClass("AVPlayerItemSampleBufferOutputAudioConfiguration")}
	})
	return _AVPlayerItemSampleBufferOutputAudioConfigurationClass
}

// GetAVPlayerItemSampleBufferOutputAudioConfigurationClass returns the class object for AVPlayerItemSampleBufferOutputAudioConfiguration.
func GetAVPlayerItemSampleBufferOutputAudioConfigurationClass() AVPlayerItemSampleBufferOutputAudioConfigurationClass {
	return getAVPlayerItemSampleBufferOutputAudioConfigurationClass()
}

type AVPlayerItemSampleBufferOutputAudioConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemSampleBufferOutputAudioConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemSampleBufferOutputAudioConfigurationClass) Alloc() AVPlayerItemSampleBufferOutputAudioConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputAudioConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Audio-specific configuration options specified when creating an
// [AVPlayerItemSampleBufferOutput].
//
// # Configuring audio output
//
//   - [AVPlayerItemSampleBufferOutputAudioConfiguration.RequestedAudioFormat]: Indicates the audio format in which the client prefers to receive the output sample buffers.
//   - [AVPlayerItemSampleBufferOutputAudioConfiguration.SetRequestedAudioFormat]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputAudioConfiguration
type AVPlayerItemSampleBufferOutputAudioConfiguration struct {
	AVPlayerItemSampleBufferOutputConfiguration
}

// AVPlayerItemSampleBufferOutputAudioConfigurationFromID constructs a [AVPlayerItemSampleBufferOutputAudioConfiguration] from an objc.ID.
//
// Audio-specific configuration options specified when creating an
// [AVPlayerItemSampleBufferOutput].
func AVPlayerItemSampleBufferOutputAudioConfigurationFromID(id objc.ID) AVPlayerItemSampleBufferOutputAudioConfiguration {
	return AVPlayerItemSampleBufferOutputAudioConfiguration{AVPlayerItemSampleBufferOutputConfiguration: AVPlayerItemSampleBufferOutputConfigurationFromID(id)}
}

// NOTE: AVPlayerItemSampleBufferOutputAudioConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemSampleBufferOutputAudioConfiguration] class.
//
// # Configuring audio output
//
//   - [IAVPlayerItemSampleBufferOutputAudioConfiguration.RequestedAudioFormat]: Indicates the audio format in which the client prefers to receive the output sample buffers.
//   - [IAVPlayerItemSampleBufferOutputAudioConfiguration.SetRequestedAudioFormat]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputAudioConfiguration
type IAVPlayerItemSampleBufferOutputAudioConfiguration interface {
	IAVPlayerItemSampleBufferOutputConfiguration

	// Topic: Configuring audio output

	// Indicates the audio format in which the client prefers to receive the output sample buffers.
	RequestedAudioFormat() coremedia.CMFormatDescriptionRef
	SetRequestedAudioFormat(value coremedia.CMFormatDescriptionRef)
}

// Init initializes the instance.
func (p AVPlayerItemSampleBufferOutputAudioConfiguration) Init() AVPlayerItemSampleBufferOutputAudioConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputAudioConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemSampleBufferOutputAudioConfiguration) Autorelease() AVPlayerItemSampleBufferOutputAudioConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputAudioConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemSampleBufferOutputAudioConfiguration creates a new AVPlayerItemSampleBufferOutputAudioConfiguration instance.
func NewAVPlayerItemSampleBufferOutputAudioConfiguration() AVPlayerItemSampleBufferOutputAudioConfiguration {
	class := getAVPlayerItemSampleBufferOutputAudioConfigurationClass()
	rv := objc.Send[AVPlayerItemSampleBufferOutputAudioConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates the audio format in which the client prefers to receive the
// output sample buffers.
//
// # Discussion
//
// Must be a PCM format.
//
// The output `CMSampleBuffers'` [CMFormatDescription] may not exactly match
// this format description, but it will match the parts described in the
// [AudioStreamBasicDescription]. The output format may differ from the
// requestedAudioFormat in its LPCM numeric type, channel interleaving and
// sample size. If any of these differs from the format in which you wish to
// operate, you can set up conversions between the format of audio sample
// buffers provided by the AVPlayerItemSampleBufferOutput and your required
// processing format by using AudioConverter or AVAudioEngine.
//
// Specifying a PCM format is currently required. In the future it may be
// optional.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputAudioConfiguration/requestedAudioFormat
func (p AVPlayerItemSampleBufferOutputAudioConfiguration) RequestedAudioFormat() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](p.ID, objc.Sel("requestedAudioFormat"))
	return coremedia.CMFormatDescriptionRef(rv)
}
func (p AVPlayerItemSampleBufferOutputAudioConfiguration) SetRequestedAudioFormat(value coremedia.CMFormatDescriptionRef) {
	objc.Send[struct{}](p.ID, objc.Sel("setRequestedAudioFormat:"), value)
}
