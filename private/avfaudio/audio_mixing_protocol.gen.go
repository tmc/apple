// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioMixing protocol.
type AVAudioMixing interface {
	objectivec.IObject

	// SetVolume protocol.
	SetVolume(volume float32)

	// Volume protocol.
	Volume() float32
}

// AVAudioMixingObject wraps an existing Objective-C object that conforms to the AVAudioMixing protocol.
type AVAudioMixingObject struct {
	objectivec.Object
}

func (o AVAudioMixingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioMixingObjectFromID constructs a [AVAudioMixingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioMixingObjectFromID(id objc.ID) AVAudioMixingObject {
	return AVAudioMixingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o AVAudioMixingObject) DestinationForMixerBus(mixer objectivec.IObject, bus uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return objectivec.Object{ID: rv}
}
func (o AVAudioMixingObject) SetVolume(volume float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setVolume:"), volume)
}
func (o AVAudioMixingObject) Volume() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("volume"))
	return rv
}
