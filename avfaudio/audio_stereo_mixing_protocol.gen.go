// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines stereo mixing properties a mixer uses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing
type AVAudioStereoMixing interface {
	objectivec.IObject

	// The bus’s stereo pan.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
	Pan() float32

	// The bus’s stereo pan.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
	SetPan(value float32)
}

// AVAudioStereoMixingObject wraps an existing Objective-C object that conforms to the AVAudioStereoMixing protocol.
type AVAudioStereoMixingObject struct {
	objectivec.Object
}

func (o AVAudioStereoMixingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioStereoMixingObjectFromID constructs a [AVAudioStereoMixingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioStereoMixingObjectFromID(id objc.ID) AVAudioStereoMixingObject {
	return AVAudioStereoMixingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The bus’s stereo pan.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (o AVAudioStereoMixingObject) Pan() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("pan"))
	return rv
}

// The bus’s stereo pan.
//
// # Discussion
//
// The default value is `0.0`, and the range of valid values is `-1.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (o AVAudioStereoMixingObject) SetPan(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPan:"), value)
}
