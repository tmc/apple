// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSApplebetMapperPhonemeInfo] class.
var (
	_TTSApplebetMapperPhonemeInfoClass     TTSApplebetMapperPhonemeInfoClass
	_TTSApplebetMapperPhonemeInfoClassOnce sync.Once
)

func getTTSApplebetMapperPhonemeInfoClass() TTSApplebetMapperPhonemeInfoClass {
	_TTSApplebetMapperPhonemeInfoClassOnce.Do(func() {
		_TTSApplebetMapperPhonemeInfoClass = TTSApplebetMapperPhonemeInfoClass{class: objc.GetClass("TTSApplebetMapperPhonemeInfo")}
	})
	return _TTSApplebetMapperPhonemeInfoClass
}

// GetTTSApplebetMapperPhonemeInfoClass returns the class object for TTSApplebetMapperPhonemeInfo.
func GetTTSApplebetMapperPhonemeInfoClass() TTSApplebetMapperPhonemeInfoClass {
	return getTTSApplebetMapperPhonemeInfoClass()
}

type TTSApplebetMapperPhonemeInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSApplebetMapperPhonemeInfoClass) Alloc() TTSApplebetMapperPhonemeInfo {
	rv := objc.Send[TTSApplebetMapperPhonemeInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSApplebetMapperPhonemeInfo.EndTime]
//   - [TTSApplebetMapperPhonemeInfo.SetEndTime]
//   - [TTSApplebetMapperPhonemeInfo.Phoneme]
//   - [TTSApplebetMapperPhonemeInfo.SetPhoneme]
//   - [TTSApplebetMapperPhonemeInfo.StartTime]
//   - [TTSApplebetMapperPhonemeInfo.SetStartTime]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperPhonemeInfo
type TTSApplebetMapperPhonemeInfo struct {
	objectivec.Object
}

// TTSApplebetMapperPhonemeInfoFromID constructs a [TTSApplebetMapperPhonemeInfo] from an objc.ID.
func TTSApplebetMapperPhonemeInfoFromID(id objc.ID) TTSApplebetMapperPhonemeInfo {
	return TTSApplebetMapperPhonemeInfo{objectivec.Object{ID: id}}
}
// Ensure TTSApplebetMapperPhonemeInfo implements ITTSApplebetMapperPhonemeInfo.
var _ ITTSApplebetMapperPhonemeInfo = TTSApplebetMapperPhonemeInfo{}

// An interface definition for the [TTSApplebetMapperPhonemeInfo] class.
//
// # Methods
//
//   - [ITTSApplebetMapperPhonemeInfo.EndTime]
//   - [ITTSApplebetMapperPhonemeInfo.SetEndTime]
//   - [ITTSApplebetMapperPhonemeInfo.Phoneme]
//   - [ITTSApplebetMapperPhonemeInfo.SetPhoneme]
//   - [ITTSApplebetMapperPhonemeInfo.StartTime]
//   - [ITTSApplebetMapperPhonemeInfo.SetStartTime]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperPhonemeInfo
type ITTSApplebetMapperPhonemeInfo interface {
	objectivec.IObject

	// Topic: Methods

	EndTime() foundation.NSNumber
	SetEndTime(value foundation.NSNumber)
	Phoneme() string
	SetPhoneme(value string)
	StartTime() foundation.NSNumber
	SetStartTime(value foundation.NSNumber)
}

// Init initializes the instance.
func (t TTSApplebetMapperPhonemeInfo) Init() TTSApplebetMapperPhonemeInfo {
	rv := objc.Send[TTSApplebetMapperPhonemeInfo](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSApplebetMapperPhonemeInfo) Autorelease() TTSApplebetMapperPhonemeInfo {
	rv := objc.Send[TTSApplebetMapperPhonemeInfo](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSApplebetMapperPhonemeInfo creates a new TTSApplebetMapperPhonemeInfo instance.
func NewTTSApplebetMapperPhonemeInfo() TTSApplebetMapperPhonemeInfo {
	class := getTTSApplebetMapperPhonemeInfoClass()
	rv := objc.Send[TTSApplebetMapperPhonemeInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperPhonemeInfo/endTime
func (t TTSApplebetMapperPhonemeInfo) EndTime() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("endTime"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSApplebetMapperPhonemeInfo) SetEndTime(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setEndTime:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperPhonemeInfo/phoneme
func (t TTSApplebetMapperPhonemeInfo) Phoneme() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("phoneme"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSApplebetMapperPhonemeInfo) SetPhoneme(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPhoneme:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperPhonemeInfo/startTime
func (t TTSApplebetMapperPhonemeInfo) StartTime() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("startTime"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSApplebetMapperPhonemeInfo) SetStartTime(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setStartTime:"), value)
}

