// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// TTSMarker protocol.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker
type TTSMarker interface {
	objectivec.IObject

	// ByteOffset protocol.
	//
	// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/byteOffset
	ByteOffset() int64

	// MarkType protocol.
	//
	// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/markType
	MarkType() int64

	// SetByteOffset protocol.
	//
	// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/setByteOffset:
	SetByteOffset(offset int64)
}

// TTSMarkerObject wraps an existing Objective-C object that conforms to the TTSMarker protocol.
type TTSMarkerObject struct {
	objectivec.Object
}

func (o TTSMarkerObject) BaseObject() objectivec.Object {
	return o.Object
}

// TTSMarkerObjectFromID constructs a [TTSMarkerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TTSMarkerObjectFromID(id objc.ID) TTSMarkerObject {
	return TTSMarkerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/avMark
func (o TTSMarkerObject) AvMark() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("avMark"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/byteOffset
func (o TTSMarkerObject) ByteOffset() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("byteOffset"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/markType
func (o TTSMarkerObject) MarkType() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("markType"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSMarker/setByteOffset:
func (o TTSMarkerObject) SetByteOffset(offset int64) {
	objc.Send[struct{}](o.ID, objc.Sel("setByteOffset:"), offset)
}
