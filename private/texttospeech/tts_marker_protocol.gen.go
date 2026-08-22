// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// TTSMarker protocol.
type TTSMarker interface {
	objectivec.IObject

	// AvMark protocol.
	AvMark() objectivec.IObject

	// ByteOffset protocol.
	ByteOffset() int64

	// MarkType protocol.
	MarkType() int64

	// SetByteOffset protocol.
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

func (o TTSMarkerObject) AvMark() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("avMark"))
	return objectivec.Object{ID: rv}
}
func (o TTSMarkerObject) ByteOffset() int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("byteOffset"))
	return rv
}
func (o TTSMarkerObject) MarkType() int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("markType"))
	return rv
}
func (o TTSMarkerObject) SetByteOffset(offset int64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setByteOffset:"), offset)
}
