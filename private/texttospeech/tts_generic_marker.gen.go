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
	rv := objc.SendIfResponds[TTSGenericMarker](objc.ID(tc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (t TTSGenericMarker) Init() TTSGenericMarker {
	rv := objc.SendIfResponds[TTSGenericMarker](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSGenericMarker) Autorelease() TTSGenericMarker {
	rv := objc.SendIfResponds[TTSGenericMarker](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSGenericMarker creates a new TTSGenericMarker instance.
func NewTTSGenericMarker() TTSGenericMarker {
	class := getTTSGenericMarkerClass()
	rv := objc.SendIfResponds[TTSGenericMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSGenericMarker) AvMark() avfaudio.AVSpeechSynthesisMarker {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("avMark"))
	return avfaudio.AVSpeechSynthesisMarkerFromID(objc.ID(rv))
}
func (t TTSGenericMarker) ByteOffset() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("byteOffset"))
	return rv
}
func (t TTSGenericMarker) SetByteOffset(value int64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setByteOffset:"), value)
}
func (t TTSGenericMarker) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSGenericMarker) Description() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSGenericMarker) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("hash"))
	return rv
}
func (t TTSGenericMarker) MarkType() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("markType"))
	return rv
}
func (t TTSGenericMarker) Name() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSGenericMarker) SetName(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}
func (t TTSGenericMarker) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](t.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
