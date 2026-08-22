// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSPhonemeMarker] class.
var (
	_TTSPhonemeMarkerClass     TTSPhonemeMarkerClass
	_TTSPhonemeMarkerClassOnce sync.Once
)

func getTTSPhonemeMarkerClass() TTSPhonemeMarkerClass {
	_TTSPhonemeMarkerClassOnce.Do(func() {
		_TTSPhonemeMarkerClass = TTSPhonemeMarkerClass{class: objc.GetClass("TTSPhonemeMarker")}
	})
	return _TTSPhonemeMarkerClass
}

// GetTTSPhonemeMarkerClass returns the class object for TTSPhonemeMarker.
func GetTTSPhonemeMarkerClass() TTSPhonemeMarkerClass {
	return getTTSPhonemeMarkerClass()
}

type TTSPhonemeMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSPhonemeMarkerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSPhonemeMarkerClass) Alloc() TTSPhonemeMarker {
	rv := objc.SendIfResponds[TTSPhonemeMarker](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSPhonemeMarker.Alphabet]
//   - [TTSPhonemeMarker.SetAlphabet]
//   - [TTSPhonemeMarker.AvMark]
//   - [TTSPhonemeMarker.ByteOffset]
//   - [TTSPhonemeMarker.SetByteOffset]
//   - [TTSPhonemeMarker.MarkType]
//   - [TTSPhonemeMarker.Phoneme]
//   - [TTSPhonemeMarker.SetPhoneme]
//   - [TTSPhonemeMarker.ToAVMarkAtOffset]
//   - [TTSPhonemeMarker.DebugDescription]
//   - [TTSPhonemeMarker.Description]
//   - [TTSPhonemeMarker.Hash]
//   - [TTSPhonemeMarker.Superclass]
type TTSPhonemeMarker struct {
	objectivec.Object
}

// TTSPhonemeMarkerFromID constructs a [TTSPhonemeMarker] from an objc.ID.
func TTSPhonemeMarkerFromID(id objc.ID) TTSPhonemeMarker {
	return TTSPhonemeMarker{objectivec.Object{ID: id}}
}

// Ensure TTSPhonemeMarker implements ITTSPhonemeMarker.
var _ ITTSPhonemeMarker = TTSPhonemeMarker{}

// An interface definition for the [TTSPhonemeMarker] class.
//
// # Methods
//
//   - [ITTSPhonemeMarker.Alphabet]
//   - [ITTSPhonemeMarker.SetAlphabet]
//   - [ITTSPhonemeMarker.AvMark]
//   - [ITTSPhonemeMarker.ByteOffset]
//   - [ITTSPhonemeMarker.SetByteOffset]
//   - [ITTSPhonemeMarker.MarkType]
//   - [ITTSPhonemeMarker.Phoneme]
//   - [ITTSPhonemeMarker.SetPhoneme]
//   - [ITTSPhonemeMarker.ToAVMarkAtOffset]
//   - [ITTSPhonemeMarker.DebugDescription]
//   - [ITTSPhonemeMarker.Description]
//   - [ITTSPhonemeMarker.Hash]
//   - [ITTSPhonemeMarker.Superclass]
type ITTSPhonemeMarker interface {
	objectivec.IObject

	// Topic: Methods

	Alphabet() int64
	SetAlphabet(value int64)
	AvMark() avfaudio.AVSpeechSynthesisMarker
	ByteOffset() int64
	SetByteOffset(value int64)
	MarkType() int64
	Phoneme() string
	SetPhoneme(value string)
	ToAVMarkAtOffset(offset int64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (t TTSPhonemeMarker) Init() TTSPhonemeMarker {
	rv := objc.SendIfResponds[TTSPhonemeMarker](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSPhonemeMarker) Autorelease() TTSPhonemeMarker {
	rv := objc.SendIfResponds[TTSPhonemeMarker](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSPhonemeMarker creates a new TTSPhonemeMarker instance.
func NewTTSPhonemeMarker() TTSPhonemeMarker {
	class := getTTSPhonemeMarkerClass()
	rv := objc.SendIfResponds[TTSPhonemeMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSPhonemeMarker) ToAVMarkAtOffset(offset int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("toAVMarkAtOffset:"), offset)
	return objectivec.Object{ID: rv}
}

func (t TTSPhonemeMarker) Alphabet() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("alphabet"))
	return rv
}
func (t TTSPhonemeMarker) SetAlphabet(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAlphabet:"), value)
}
func (t TTSPhonemeMarker) AvMark() avfaudio.AVSpeechSynthesisMarker {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("avMark"))
	return avfaudio.AVSpeechSynthesisMarkerFromID(objc.ID(rv))
}
func (t TTSPhonemeMarker) ByteOffset() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("byteOffset"))
	return rv
}
func (t TTSPhonemeMarker) SetByteOffset(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setByteOffset:"), value)
}
func (t TTSPhonemeMarker) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSPhonemeMarker) Description() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSPhonemeMarker) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("hash"))
	return rv
}
func (t TTSPhonemeMarker) MarkType() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("markType"))
	return rv
}
func (t TTSPhonemeMarker) Phoneme() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("phoneme"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSPhonemeMarker) SetPhoneme(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setPhoneme:"), objc.String(value))
}
func (t TTSPhonemeMarker) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](t.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
