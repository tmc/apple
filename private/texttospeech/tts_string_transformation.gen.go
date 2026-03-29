// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSStringTransformation] class.
var (
	_TTSStringTransformationClass     TTSStringTransformationClass
	_TTSStringTransformationClassOnce sync.Once
)

func getTTSStringTransformationClass() TTSStringTransformationClass {
	_TTSStringTransformationClassOnce.Do(func() {
		_TTSStringTransformationClass = TTSStringTransformationClass{class: objc.GetClass("TTSStringTransformation")}
	})
	return _TTSStringTransformationClass
}

// GetTTSStringTransformationClass returns the class object for TTSStringTransformation.
func GetTTSStringTransformationClass() TTSStringTransformationClass {
	return getTTSStringTransformationClass()
}

type TTSStringTransformationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSStringTransformationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSStringTransformationClass) Alloc() TTSStringTransformation {
	rv := objc.Send[TTSStringTransformation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSStringTransformation.FinalRange]
//   - [TTSStringTransformation.SetFinalRange]
//   - [TTSStringTransformation.OffsetFromEnd]
//   - [TTSStringTransformation.SetOffsetFromEnd]
//   - [TTSStringTransformation.Range]
//   - [TTSStringTransformation.SetRange]
//   - [TTSStringTransformation.Replacement]
//   - [TTSStringTransformation.SetReplacement]
//   - [TTSStringTransformation.SizeDelta]
//   - [TTSStringTransformation.InitWithRangeAndReplacement]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation
type TTSStringTransformation struct {
	objectivec.Object
}

// TTSStringTransformationFromID constructs a [TTSStringTransformation] from an objc.ID.
func TTSStringTransformationFromID(id objc.ID) TTSStringTransformation {
	return TTSStringTransformation{objectivec.Object{ID: id}}
}
// Ensure TTSStringTransformation implements ITTSStringTransformation.
var _ ITTSStringTransformation = TTSStringTransformation{}

// An interface definition for the [TTSStringTransformation] class.
//
// # Methods
//
//   - [ITTSStringTransformation.FinalRange]
//   - [ITTSStringTransformation.SetFinalRange]
//   - [ITTSStringTransformation.OffsetFromEnd]
//   - [ITTSStringTransformation.SetOffsetFromEnd]
//   - [ITTSStringTransformation.Range]
//   - [ITTSStringTransformation.SetRange]
//   - [ITTSStringTransformation.Replacement]
//   - [ITTSStringTransformation.SetReplacement]
//   - [ITTSStringTransformation.SizeDelta]
//   - [ITTSStringTransformation.InitWithRangeAndReplacement]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation
type ITTSStringTransformation interface {
	objectivec.IObject

	// Topic: Methods

	FinalRange() foundation.NSRange
	SetFinalRange(value foundation.NSRange)
	OffsetFromEnd() uint64
	SetOffsetFromEnd(value uint64)
	Range() foundation.NSRange
	SetRange(value foundation.NSRange)
	Replacement() string
	SetReplacement(value string)
	SizeDelta() int64
	InitWithRangeAndReplacement(range_ foundation.NSRange, replacement objectivec.IObject) TTSStringTransformation
}

// Init initializes the instance.
func (t TTSStringTransformation) Init() TTSStringTransformation {
	rv := objc.Send[TTSStringTransformation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSStringTransformation) Autorelease() TTSStringTransformation {
	rv := objc.Send[TTSStringTransformation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSStringTransformation creates a new TTSStringTransformation instance.
func NewTTSStringTransformation() TTSStringTransformation {
	class := getTTSStringTransformationClass()
	rv := objc.Send[TTSStringTransformation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/initWithRange:andReplacement:
func NewTTSStringTransformationWithRangeAndReplacement(range_ foundation.NSRange, replacement objectivec.IObject) TTSStringTransformation {
	instance := getTTSStringTransformationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRange:andReplacement:"), range_, replacement)
	return TTSStringTransformationFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/sizeDelta
func (t TTSStringTransformation) SizeDelta() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("sizeDelta"))
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/initWithRange:andReplacement:
func (t TTSStringTransformation) InitWithRangeAndReplacement(range_ foundation.NSRange, replacement objectivec.IObject) TTSStringTransformation {
	rv := objc.Send[TTSStringTransformation](t.ID, objc.Sel("initWithRange:andReplacement:"), range_, replacement)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/finalRange
func (t TTSStringTransformation) FinalRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("finalRange"))
	return foundation.NSRange(rv)
}
func (t TTSStringTransformation) SetFinalRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setFinalRange:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/offsetFromEnd
func (t TTSStringTransformation) OffsetFromEnd() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("offsetFromEnd"))
	return rv
}
func (t TTSStringTransformation) SetOffsetFromEnd(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setOffsetFromEnd:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/range
func (t TTSStringTransformation) Range() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("range"))
	return foundation.NSRange(rv)
}
func (t TTSStringTransformation) SetRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setRange:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStringTransformation/replacement
func (t TTSStringTransformation) Replacement() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("replacement"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSStringTransformation) SetReplacement(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setReplacement:"), objc.String(value))
}

