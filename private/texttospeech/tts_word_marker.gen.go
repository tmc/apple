// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSWordMarker] class.
var (
	_TTSWordMarkerClass     TTSWordMarkerClass
	_TTSWordMarkerClassOnce sync.Once
)

func getTTSWordMarkerClass() TTSWordMarkerClass {
	_TTSWordMarkerClassOnce.Do(func() {
		_TTSWordMarkerClass = TTSWordMarkerClass{class: objc.GetClass("TTSWordMarker")}
	})
	return _TTSWordMarkerClass
}

// GetTTSWordMarkerClass returns the class object for TTSWordMarker.
func GetTTSWordMarkerClass() TTSWordMarkerClass {
	return getTTSWordMarkerClass()
}

type TTSWordMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSWordMarkerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSWordMarkerClass) Alloc() TTSWordMarker {
	rv := objc.Send[TTSWordMarker](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSWordMarker.AvMark]
//   - [TTSWordMarker.ByteOffset]
//   - [TTSWordMarker.SetByteOffset]
//   - [TTSWordMarker.MarkType]
//   - [TTSWordMarker.WordRange]
//   - [TTSWordMarker.SetWordRange]
//   - [TTSWordMarker.DebugDescription]
//   - [TTSWordMarker.Description]
//   - [TTSWordMarker.Hash]
//   - [TTSWordMarker.Superclass]
type TTSWordMarker struct {
	objectivec.Object
}

// TTSWordMarkerFromID constructs a [TTSWordMarker] from an objc.ID.
func TTSWordMarkerFromID(id objc.ID) TTSWordMarker {
	return TTSWordMarker{objectivec.Object{ID: id}}
}

// Ensure TTSWordMarker implements ITTSWordMarker.
var _ ITTSWordMarker = TTSWordMarker{}

// An interface definition for the [TTSWordMarker] class.
//
// # Methods
//
//   - [ITTSWordMarker.AvMark]
//   - [ITTSWordMarker.ByteOffset]
//   - [ITTSWordMarker.SetByteOffset]
//   - [ITTSWordMarker.MarkType]
//   - [ITTSWordMarker.WordRange]
//   - [ITTSWordMarker.SetWordRange]
//   - [ITTSWordMarker.DebugDescription]
//   - [ITTSWordMarker.Description]
//   - [ITTSWordMarker.Hash]
//   - [ITTSWordMarker.Superclass]
type ITTSWordMarker interface {
	objectivec.IObject

	// Topic: Methods

	AvMark() avfaudio.AVSpeechSynthesisMarker
	ByteOffset() int64
	SetByteOffset(value int64)
	MarkType() int64
	WordRange() foundation.NSRange
	SetWordRange(value foundation.NSRange)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (t TTSWordMarker) Init() TTSWordMarker {
	rv := objc.Send[TTSWordMarker](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSWordMarker) Autorelease() TTSWordMarker {
	rv := objc.Send[TTSWordMarker](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSWordMarker creates a new TTSWordMarker instance.
func NewTTSWordMarker() TTSWordMarker {
	class := getTTSWordMarkerClass()
	rv := objc.Send[TTSWordMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSWordMarker) AvMark() avfaudio.AVSpeechSynthesisMarker {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("avMark"))
	return avfaudio.AVSpeechSynthesisMarkerFromID(objc.ID(rv))
}
func (t TTSWordMarker) ByteOffset() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("byteOffset"))
	return rv
}
func (t TTSWordMarker) SetByteOffset(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setByteOffset:"), value)
}
func (t TTSWordMarker) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSWordMarker) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSWordMarker) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}
func (t TTSWordMarker) MarkType() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("markType"))
	return rv
}
func (t TTSWordMarker) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](t.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (t TTSWordMarker) WordRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("wordRange"))
	return foundation.NSRange(rv)
}
func (t TTSWordMarker) SetWordRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setWordRange:"), value)
}
