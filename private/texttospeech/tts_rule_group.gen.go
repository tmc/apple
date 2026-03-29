// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRuleGroup] class.
var (
	_TTSRuleGroupClass     TTSRuleGroupClass
	_TTSRuleGroupClassOnce sync.Once
)

func getTTSRuleGroupClass() TTSRuleGroupClass {
	_TTSRuleGroupClassOnce.Do(func() {
		_TTSRuleGroupClass = TTSRuleGroupClass{class: objc.GetClass("TTSRuleGroup")}
	})
	return _TTSRuleGroupClass
}

// GetTTSRuleGroupClass returns the class object for TTSRuleGroup.
func GetTTSRuleGroupClass() TTSRuleGroupClass {
	return getTTSRuleGroupClass()
}

type TTSRuleGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSRuleGroupClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRuleGroupClass) Alloc() TTSRuleGroup {
	rv := objc.Send[TTSRuleGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRuleGroup.EndIndex]
//   - [TTSRuleGroup.SetEndIndex]
//   - [TTSRuleGroup.Key]
//   - [TTSRuleGroup.SetKey]
//   - [TTSRuleGroup.StartIndex]
//   - [TTSRuleGroup.SetStartIndex]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleGroup
type TTSRuleGroup struct {
	objectivec.Object
}

// TTSRuleGroupFromID constructs a [TTSRuleGroup] from an objc.ID.
func TTSRuleGroupFromID(id objc.ID) TTSRuleGroup {
	return TTSRuleGroup{objectivec.Object{ID: id}}
}
// Ensure TTSRuleGroup implements ITTSRuleGroup.
var _ ITTSRuleGroup = TTSRuleGroup{}

// An interface definition for the [TTSRuleGroup] class.
//
// # Methods
//
//   - [ITTSRuleGroup.EndIndex]
//   - [ITTSRuleGroup.SetEndIndex]
//   - [ITTSRuleGroup.Key]
//   - [ITTSRuleGroup.SetKey]
//   - [ITTSRuleGroup.StartIndex]
//   - [ITTSRuleGroup.SetStartIndex]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleGroup
type ITTSRuleGroup interface {
	objectivec.IObject

	// Topic: Methods

	EndIndex() uint64
	SetEndIndex(value uint64)
	Key() string
	SetKey(value string)
	StartIndex() uint64
	SetStartIndex(value uint64)
}

// Init initializes the instance.
func (t TTSRuleGroup) Init() TTSRuleGroup {
	rv := objc.Send[TTSRuleGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRuleGroup) Autorelease() TTSRuleGroup {
	rv := objc.Send[TTSRuleGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRuleGroup creates a new TTSRuleGroup instance.
func NewTTSRuleGroup() TTSRuleGroup {
	class := getTTSRuleGroupClass()
	rv := objc.Send[TTSRuleGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleGroup/endIndex
func (t TTSRuleGroup) EndIndex() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("endIndex"))
	return rv
}
func (t TTSRuleGroup) SetEndIndex(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setEndIndex:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleGroup/key
func (t TTSRuleGroup) Key() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSRuleGroup) SetKey(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleGroup/startIndex
func (t TTSRuleGroup) StartIndex() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("startIndex"))
	return rv
}
func (t TTSRuleGroup) SetStartIndex(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setStartIndex:"), value)
}

