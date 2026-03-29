// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSSoundDelegate protocol.
//
// See: https://developer.apple.com/documentation/SpeechObjects/NSSoundDelegate
type NSSoundDelegate interface {
	objectivec.IObject
}

// NSSoundDelegateObject wraps an existing Objective-C object that conforms to the NSSoundDelegate protocol.
type NSSoundDelegateObject struct {
	objectivec.Object
}
func (o NSSoundDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSoundDelegateObjectFromID constructs a [NSSoundDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSoundDelegateObjectFromID(id objc.ID) NSSoundDelegateObject {
	return NSSoundDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/NSSoundDelegate/sound:didFinishPlaying:
func (o NSSoundDelegateObject) SoundDidFinishPlaying(sound objectivec.IObject, playing bool) {
	objc.Send[struct{}](o.ID, objc.Sel("sound:didFinishPlaying:"), sound, playing)
	}

