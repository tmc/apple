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
	rv := objc.SendIfResponds[AVAudioSessionCapability](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioSessionCapability.InitWithIsSupportedIsEnabled]
//   - [AVAudioSessionCapability.Enabled]
//   - [AVAudioSessionCapability.Supported]
type AVAudioSessionCapability struct {
	objectivec.Object
}

// AVAudioSessionCapabilityFromID constructs a [AVAudioSessionCapability] from an objc.ID.
func AVAudioSessionCapabilityFromID(id objc.ID) AVAudioSessionCapability {
	return AVAudioSessionCapability{objectivec.Object{ID: id}}
}

// Ensure AVAudioSessionCapability implements IAVAudioSessionCapability.
var _ IAVAudioSessionCapability = AVAudioSessionCapability{}

// An interface definition for the [AVAudioSessionCapability] class.
//
// # Methods
//
//   - [IAVAudioSessionCapability.InitWithIsSupportedIsEnabled]
//   - [IAVAudioSessionCapability.Enabled]
//   - [IAVAudioSessionCapability.Supported]
type IAVAudioSessionCapability interface {
	objectivec.IObject

	// Topic: Methods

	InitWithIsSupportedIsEnabled(supported bool, enabled bool) AVAudioSessionCapability
	Enabled() bool
	Supported() bool
}

// Init initializes the instance.
func (a AVAudioSessionCapability) Init() AVAudioSessionCapability {
	rv := objc.SendIfResponds[AVAudioSessionCapability](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSessionCapability) Autorelease() AVAudioSessionCapability {
	rv := objc.SendIfResponds[AVAudioSessionCapability](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSessionCapability creates a new AVAudioSessionCapability instance.
func NewAVAudioSessionCapability() AVAudioSessionCapability {
	class := getAVAudioSessionCapabilityClass()
	rv := objc.SendIfResponds[AVAudioSessionCapability](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioSessionCapabilityWithIsSupportedIsEnabled(supported bool, enabled bool) AVAudioSessionCapability {
	instance := getAVAudioSessionCapabilityClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIsSupported:isEnabled:"), supported, enabled)
	return AVAudioSessionCapabilityFromID(rv)
}

func (a AVAudioSessionCapability) InitWithIsSupportedIsEnabled(supported bool, enabled bool) AVAudioSessionCapability {
	rv := objc.SendIfResponds[AVAudioSessionCapability](a.ID, objc.Sel("initWithIsSupported:isEnabled:"), supported, enabled)
	return rv
}

func (a AVAudioSessionCapability) Enabled() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("enabled"))
	return rv
}
func (a AVAudioSessionCapability) Supported() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("supported"))
	return rv
}
