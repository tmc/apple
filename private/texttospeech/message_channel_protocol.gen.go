// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AUMessageChannel protocol.
//
// See: https://developer.apple.com/documentation/TextToSpeech/AUMessageChannel
type AUMessageChannel interface {
	objectivec.IObject

	// CallHostBlock protocol.
	//
	// See: https://developer.apple.com/documentation/TextToSpeech/AUMessageChannel/callHostBlock
	CallHostBlock() objectivec.IObject
}

// AUMessageChannelObject wraps an existing Objective-C object that conforms to the AUMessageChannel protocol.
type AUMessageChannelObject struct {
	objectivec.Object
}

func (o AUMessageChannelObject) BaseObject() objectivec.Object {
	return o.Object
}

// AUMessageChannelObjectFromID constructs a [AUMessageChannelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUMessageChannelObjectFromID(id objc.ID) AUMessageChannelObject {
	return AUMessageChannelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/TextToSpeech/AUMessageChannel/callAudioUnit:
func (o AUMessageChannelObject) CallAudioUnit(unit objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("callAudioUnit:"), unit)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/AUMessageChannel/callHostBlock
func (o AUMessageChannelObject) CallHostBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("callHostBlock"))
	return objectivec.Object{ID: rv}
}
