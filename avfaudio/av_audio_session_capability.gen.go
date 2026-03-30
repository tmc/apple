// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioSessionCapability] class.
var (
	_AVAudioSessionCapabilityClass     AVAudioSessionCapabilityClass
	_AVAudioSessionCapabilityClassOnce sync.Once
)

func getAVAudioSessionCapabilityClass() AVAudioSessionCapabilityClass {
	_AVAudioSessionCapabilityClassOnce.Do(func() {
		_AVAudioSessionCapabilityClass = AVAudioSessionCapabilityClass{class: objc.GetClass("AVAudioSessionCapability")}
	})
	return _AVAudioSessionCapabilityClass
}

// GetAVAudioSessionCapabilityClass returns the class object for AVAudioSessionCapability.
func GetAVAudioSessionCapabilityClass() AVAudioSessionCapabilityClass {
	return getAVAudioSessionCapabilityClass()
}

type AVAudioSessionCapabilityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSessionCapabilityClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSessionCapabilityClass) Alloc() AVAudioSessionCapability {
	rv := objc.Send[AVAudioSessionCapability](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Describes whether a specific capability is supported and if that capability
// is currently enabled
//
// # Inspecting a capability
//
//   - [AVAudioSessionCapability.Enabled]: A Boolean value that indicates whether the capability is enabled.
//   - [AVAudioSessionCapability.Supported]: A Boolean value that indicates whether the capability is supported.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability
type AVAudioSessionCapability struct {
	objectivec.Object
}

// AVAudioSessionCapabilityFromID constructs a [AVAudioSessionCapability] from an objc.ID.
//
// Describes whether a specific capability is supported and if that capability
// is currently enabled
func AVAudioSessionCapabilityFromID(id objc.ID) AVAudioSessionCapability {
	return AVAudioSessionCapability{objectivec.Object{ID: id}}
}

// NOTE: AVAudioSessionCapability adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioSessionCapability] class.
//
// # Inspecting a capability
//
//   - [IAVAudioSessionCapability.Enabled]: A Boolean value that indicates whether the capability is enabled.
//   - [IAVAudioSessionCapability.Supported]: A Boolean value that indicates whether the capability is supported.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability
type IAVAudioSessionCapability interface {
	objectivec.IObject

	// Topic: Inspecting a capability

	// A Boolean value that indicates whether the capability is enabled.
	Enabled() bool
	// A Boolean value that indicates whether the capability is supported.
	Supported() bool

	// An optional port extension that describes capabilities relevant to Bluetooth microphone ports.
	BluetoothMicrophoneExtension() objc.ID
	SetBluetoothMicrophoneExtension(value objc.ID)
}

// Init initializes the instance.
func (a AVAudioSessionCapability) Init() AVAudioSessionCapability {
	rv := objc.Send[AVAudioSessionCapability](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSessionCapability) Autorelease() AVAudioSessionCapability {
	rv := objc.Send[AVAudioSessionCapability](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSessionCapability creates a new AVAudioSessionCapability instance.
func NewAVAudioSessionCapability() AVAudioSessionCapability {
	class := getAVAudioSessionCapabilityClass()
	rv := objc.Send[AVAudioSessionCapability](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the capability is enabled.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isEnabled
func (a AVAudioSessionCapability) Enabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the capability is supported.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isSupported
func (a AVAudioSessionCapability) Supported() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSupported"))
	return rv
}

// An optional port extension that describes capabilities relevant to
// Bluetooth microphone ports.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/bluetoothmicrophoneextension
func (a AVAudioSessionCapability) BluetoothMicrophoneExtension() objc.ID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bluetoothMicrophoneExtension"))
	return rv
}
func (a AVAudioSessionCapability) SetBluetoothMicrophoneExtension(value objc.ID) {
	objc.Send[struct{}](a.ID, objc.Sel("setBluetoothMicrophoneExtension:"), value)
}
