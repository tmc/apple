// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSMatchedRuleReplacement] class.
var (
	_TTSMatchedRuleReplacementClass     TTSMatchedRuleReplacementClass
	_TTSMatchedRuleReplacementClassOnce sync.Once
)

func getTTSMatchedRuleReplacementClass() TTSMatchedRuleReplacementClass {
	_TTSMatchedRuleReplacementClassOnce.Do(func() {
		_TTSMatchedRuleReplacementClass = TTSMatchedRuleReplacementClass{class: objc.GetClass("TTSMatchedRuleReplacement")}
	})
	return _TTSMatchedRuleReplacementClass
}

// GetTTSMatchedRuleReplacementClass returns the class object for TTSMatchedRuleReplacement.
func GetTTSMatchedRuleReplacementClass() TTSMatchedRuleReplacementClass {
	return getTTSMatchedRuleReplacementClass()
}

type TTSMatchedRuleReplacementClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSMatchedRuleReplacementClass) Alloc() TTSMatchedRuleReplacement {
	rv := objc.Send[TTSMatchedRuleReplacement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSMatchedRuleReplacement.Match]
//   - [TTSMatchedRuleReplacement.SetMatch]
//   - [TTSMatchedRuleReplacement.Replacement]
//   - [TTSMatchedRuleReplacement.SetReplacement]
//   - [TTSMatchedRuleReplacement.RuleReplacement]
//   - [TTSMatchedRuleReplacement.SetRuleReplacement]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSMatchedRuleReplacement
type TTSMatchedRuleReplacement struct {
	objectivec.Object
}

// TTSMatchedRuleReplacementFromID constructs a [TTSMatchedRuleReplacement] from an objc.ID.
func TTSMatchedRuleReplacementFromID(id objc.ID) TTSMatchedRuleReplacement {
	return TTSMatchedRuleReplacement{objectivec.Object{ID: id}}
}
// Ensure TTSMatchedRuleReplacement implements ITTSMatchedRuleReplacement.
var _ ITTSMatchedRuleReplacement = TTSMatchedRuleReplacement{}

// An interface definition for the [TTSMatchedRuleReplacement] class.
//
// # Methods
//
//   - [ITTSMatchedRuleReplacement.Match]
//   - [ITTSMatchedRuleReplacement.SetMatch]
//   - [ITTSMatchedRuleReplacement.Replacement]
//   - [ITTSMatchedRuleReplacement.SetReplacement]
//   - [ITTSMatchedRuleReplacement.RuleReplacement]
//   - [ITTSMatchedRuleReplacement.SetRuleReplacement]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSMatchedRuleReplacement
type ITTSMatchedRuleReplacement interface {
	objectivec.IObject

	// Topic: Methods

	Match() ITTSRegexMatch
	SetMatch(value ITTSRegexMatch)
	Replacement() string
	SetReplacement(value string)
	RuleReplacement() ITTSRuleReplacement
	SetRuleReplacement(value ITTSRuleReplacement)
}

// Init initializes the instance.
func (t TTSMatchedRuleReplacement) Init() TTSMatchedRuleReplacement {
	rv := objc.Send[TTSMatchedRuleReplacement](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSMatchedRuleReplacement) Autorelease() TTSMatchedRuleReplacement {
	rv := objc.Send[TTSMatchedRuleReplacement](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSMatchedRuleReplacement creates a new TTSMatchedRuleReplacement instance.
func NewTTSMatchedRuleReplacement() TTSMatchedRuleReplacement {
	class := getTTSMatchedRuleReplacementClass()
	rv := objc.Send[TTSMatchedRuleReplacement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSMatchedRuleReplacement/match
func (t TTSMatchedRuleReplacement) Match() ITTSRegexMatch {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("match"))
	return TTSRegexMatchFromID(objc.ID(rv))
}
func (t TTSMatchedRuleReplacement) SetMatch(value ITTSRegexMatch) {
	objc.Send[struct{}](t.ID, objc.Sel("setMatch:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSMatchedRuleReplacement/replacement
func (t TTSMatchedRuleReplacement) Replacement() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("replacement"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSMatchedRuleReplacement) SetReplacement(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setReplacement:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSMatchedRuleReplacement/ruleReplacement
func (t TTSMatchedRuleReplacement) RuleReplacement() ITTSRuleReplacement {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleReplacement"))
	return TTSRuleReplacementFromID(objc.ID(rv))
}
func (t TTSMatchedRuleReplacement) SetRuleReplacement(value ITTSRuleReplacement) {
	objc.Send[struct{}](t.ID, objc.Sel("setRuleReplacement:"), value)
}

