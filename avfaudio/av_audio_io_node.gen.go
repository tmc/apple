// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioIONode] class.
var (
	_AVAudioIONodeClass     AVAudioIONodeClass
	_AVAudioIONodeClassOnce sync.Once
)

func getAVAudioIONodeClass() AVAudioIONodeClass {
	_AVAudioIONodeClassOnce.Do(func() {
		_AVAudioIONodeClass = AVAudioIONodeClass{class: objc.GetClass("AVAudioIONode")}
	})
	return _AVAudioIONodeClass
}

// GetAVAudioIONodeClass returns the class object for AVAudioIONode.
func GetAVAudioIONodeClass() AVAudioIONodeClass {
	return getAVAudioIONodeClass()
}

type AVAudioIONodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioIONodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioIONodeClass) Alloc() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that performs audio input or output in the engine.
//
// # Overview
//
// When rendering to and from an audio device in macOS, [AVAudioInputNode] and
// [AVAudioOutputNode] communicate with the system’s default input and
// output devices. In iOS, they communicate with the devices appropriate to
// the app’s [AVAudioSession] category, configurations, and user actions,
// such as connecting or disconnecting external devices.
//
// In the manual rendering mode, [AVAudioInputNode] and [AVAudioOutputNode]
// perform the input and output in the engine in response to the client’s
// request.
//
// # Getting the Audio Unit
//
//   - [AVAudioIONode.AudioUnit]: The node’s underlying audio unit, if any.
//
// # Getting the I/O Latency
//
//   - [AVAudioIONode.PresentationLatency]: The presentation or hardware latency, applicable when rendering to or from an audio device.
//
// # Getting and Setting the Voice Processing State
//
//   - [AVAudioIONode.SetVoiceProcessingEnabledError]: Enables or disables voice processing on the I/O node.
//   - [AVAudioIONode.VoiceProcessingEnabled]: A Boolean value that indicates whether voice processing is in an enabled state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
type AVAudioIONode struct {
	AVAudioNode
}

// AVAudioIONodeFromID constructs a [AVAudioIONode] from an objc.ID.
//
// An object that performs audio input or output in the engine.
func AVAudioIONodeFromID(id objc.ID) AVAudioIONode {
	return AVAudioIONode{AVAudioNode: AVAudioNodeFromID(id)}
}

// NOTE: AVAudioIONode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioIONode] class.
//
// # Getting the Audio Unit
//
//   - [IAVAudioIONode.AudioUnit]: The node’s underlying audio unit, if any.
//
// # Getting the I/O Latency
//
//   - [IAVAudioIONode.PresentationLatency]: The presentation or hardware latency, applicable when rendering to or from an audio device.
//
// # Getting and Setting the Voice Processing State
//
//   - [IAVAudioIONode.SetVoiceProcessingEnabledError]: Enables or disables voice processing on the I/O node.
//   - [IAVAudioIONode.VoiceProcessingEnabled]: A Boolean value that indicates whether voice processing is in an enabled state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
type IAVAudioIONode interface {
	IAVAudioNode

	// Topic: Getting the Audio Unit

	// The node’s underlying audio unit, if any.
	AudioUnit() IAVAudioUnit

	// Topic: Getting the I/O Latency

	// The presentation or hardware latency, applicable when rendering to or from an audio device.
	PresentationLatency() float64

	// Topic: Getting and Setting the Voice Processing State

	// Enables or disables voice processing on the I/O node.
	SetVoiceProcessingEnabledError(enabled bool) (bool, error)
	// A Boolean value that indicates whether voice processing is in an enabled state.
	VoiceProcessingEnabled() bool
}

// Init initializes the instance.
func (a AVAudioIONode) Init() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioIONode) Autorelease() AVAudioIONode {
	rv := objc.Send[AVAudioIONode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioIONode creates a new AVAudioIONode instance.
func NewAVAudioIONode() AVAudioIONode {
	class := getAVAudioIONodeClass()
	rv := objc.Send[AVAudioIONode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Enables or disables voice processing on the I/O node.
//
// enabled: The Boolean value that indicates whether to enable voice processing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/setVoiceProcessingEnabled(_:)
func (a AVAudioIONode) SetVoiceProcessingEnabledError(enabled bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setVoiceProcessingEnabled:error:"), enabled, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setVoiceProcessingEnabled:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The node’s underlying audio unit, if any.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/audioUnit
func (a AVAudioIONode) AudioUnit() IAVAudioUnit {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioUnit"))
	return AVAudioUnitFromID(objc.ID(rv))
}

// The presentation or hardware latency, applicable when rendering to or from
// an audio device.
//
// # Discussion
//
// This corresponds to `kAudioDevicePropertyLatency` and
// `kAudioStreamPropertyLatency`. For more information, see
// `AudioHardwareBase.H()` in `CoreAudio.Framework`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/presentationLatency
func (a AVAudioIONode) PresentationLatency() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("presentationLatency"))
	return rv
}

// A Boolean value that indicates whether voice processing is in an enabled
// state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/isVoiceProcessingEnabled
func (a AVAudioIONode) VoiceProcessingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isVoiceProcessingEnabled"))
	return rv
}
