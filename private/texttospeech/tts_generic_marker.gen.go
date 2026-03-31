// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSGenericMarker] class.
var (
	_TTSGenericMarkerClass     TTSGenericMarkerClass
	_TTSGenericMarkerClassOnce sync.Once
)

func getTTSGenericMarkerClass() TTSGenericMarkerClass {
	_TTSGenericMarkerClassOnce.Do(func() {
		_TTSGenericMarkerClass = TTSGenericMarkerClass{class: objc.GetClass("TTSGenericMarker")}
	})
	return _TTSGenericMarkerClass
}

// GetTTSGenericMarkerClass returns the class object for TTSGenericMarker.
func GetTTSGenericMarkerClass() TTSGenericMarkerClass {
	return getTTSGenericMarkerClass()
}

type TTSGenericMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSGenericMarkerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSGenericMarkerClass) Alloc() TTSGenericMarker {
	rv := objc.Send[TTSGenericMarker](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSGenericMarker.AvMark]
//   - [TTSGenericMarker.ByteOffset]
//   - [TTSGenericMarker.SetByteOffset]
//   - [TTSGenericMarker.MarkType]
//   - [TTSGenericMarker.Name]
//   - [TTSGenericMarker.SetName]
//   - [TTSGenericMarker.DebugDescription]
//   - [TTSGenericMarker.Description]
//   - [TTSGenericMarker.Hash]
//   - [TTSGenericMarker.Superclass]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker
type TTSGenericMarker struct {
	objectivec.Object
}

// TTSGenericMarkerFromID constructs a [TTSGenericMarker] from an objc.ID.
func TTSGenericMarkerFromID(id objc.ID) TTSGenericMarker {
	return TTSGenericMarker{objectivec.Object{ID: id}}
}

// Ensure TTSGenericMarker implements ITTSGenericMarker.
var _ ITTSGenericMarker = TTSGenericMarker{}

// An interface definition for the [TTSGenericMarker] class.
//
// # Methods
//
//   - [ITTSGenericMarker.AvMark]
//   - [ITTSGenericMarker.ByteOffset]
//   - [ITTSGenericMarker.SetByteOffset]
//   - [ITTSGenericMarker.MarkType]
//   - [ITTSGenericMarker.Name]
//   - [ITTSGenericMarker.SetName]
//   - [ITTSGenericMarker.DebugDescription]
//   - [ITTSGenericMarker.Description]
//   - [ITTSGenericMarker.Hash]
//   - [ITTSGenericMarker.Superclass]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker
type ITTSGenericMarker interface {
	objectivec.IObject

	// Topic: Methods

	AvMark() avfaudio.AVSpeechSynthesisMarker
	ByteOffset() int64
	SetByteOffset(value int64)
	MarkType() int64
	Name() string
	SetName(value string)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (t TTSGenericMarker) Init() TTSGenericMarker {
	rv := objc.Send[TTSGenericMarker](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSGenericMarker) Autorelease() TTSGenericMarker {
	rv := objc.Send[TTSGenericMarker](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSGenericMarker creates a new TTSGenericMarker instance.
func NewTTSGenericMarker() TTSGenericMarker {
	class := getTTSGenericMarkerClass()
	rv := objc.Send[TTSGenericMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/avMark
func (t TTSGenericMarker) AvMark() avfaudio.AVSpeechSynthesisMarker {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("avMark"))
	return avfaudio.AVSpeechSynthesisMarkerFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/byteOffset
func (t TTSGenericMarker) ByteOffset() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("byteOffset"))
	return rv
}
func (t TTSGenericMarker) SetByteOffset(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setByteOffset:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/debugDescription
func (t TTSGenericMarker) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/description
func (t TTSGenericMarker) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/hash
func (t TTSGenericMarker) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/markType
func (t TTSGenericMarker) MarkType() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("markType"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/name
func (t TTSGenericMarker) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSGenericMarker) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSGenericMarker/superclass
func (t TTSGenericMarker) Superclass() objc.Class {
	rv := objc.Send[objc.Class](t.ID, objc.Sel("superclass"))
	return rv
}
