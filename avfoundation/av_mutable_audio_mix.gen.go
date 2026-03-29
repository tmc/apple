// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableAudioMix] class.
var (
	_AVMutableAudioMixClass     AVMutableAudioMixClass
	_AVMutableAudioMixClassOnce sync.Once
)

func getAVMutableAudioMixClass() AVMutableAudioMixClass {
	_AVMutableAudioMixClassOnce.Do(func() {
		_AVMutableAudioMixClass = AVMutableAudioMixClass{class: objc.GetClass("AVMutableAudioMix")}
	})
	return _AVMutableAudioMixClass
}

// GetAVMutableAudioMixClass returns the class object for AVMutableAudioMix.
func GetAVMutableAudioMixClass() AVMutableAudioMixClass {
	return getAVMutableAudioMixClass()
}

type AVMutableAudioMixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableAudioMixClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableAudioMixClass) Alloc() AVMutableAudioMix {
	rv := objc.Send[AVMutableAudioMix](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the input parameters for mixing audio tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix
type AVMutableAudioMix struct {
	AVAudioMix
}

// AVMutableAudioMixFromID constructs a [AVMutableAudioMix] from an objc.ID.
//
// An object that manages the input parameters for mixing audio tracks.
func AVMutableAudioMixFromID(id objc.ID) AVMutableAudioMix {
	return AVMutableAudioMix{AVAudioMix: AVAudioMixFromID(id)}
}
// NOTE: AVMutableAudioMix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableAudioMix] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix
type IAVMutableAudioMix interface {
	IAVAudioMix
}

// Init initializes the instance.
func (m AVMutableAudioMix) Init() AVMutableAudioMix {
	rv := objc.Send[AVMutableAudioMix](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableAudioMix) Autorelease() AVMutableAudioMix {
	rv := objc.Send[AVMutableAudioMix](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableAudioMix creates a new AVMutableAudioMix instance.
func NewAVMutableAudioMix() AVMutableAudioMix {
	class := getAVMutableAudioMixClass()
	rv := objc.Send[AVMutableAudioMix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a new mutable audio mix.
//
// # Return Value
// 
// A new mutable audio mix.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/audioMix
func (_AVMutableAudioMixClass AVMutableAudioMixClass) AudioMix() AVMutableAudioMix {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableAudioMixClass.class), objc.Sel("audioMix"))
	return AVMutableAudioMixFromID(rv)
}

