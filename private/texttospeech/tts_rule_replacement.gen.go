// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRuleReplacement] class.
var (
	_TTSRuleReplacementClass     TTSRuleReplacementClass
	_TTSRuleReplacementClassOnce sync.Once
)

func getTTSRuleReplacementClass() TTSRuleReplacementClass {
	_TTSRuleReplacementClassOnce.Do(func() {
		_TTSRuleReplacementClass = TTSRuleReplacementClass{class: objc.GetClass("TTSRuleReplacement")}
	})
	return _TTSRuleReplacementClass
}

// GetTTSRuleReplacementClass returns the class object for TTSRuleReplacement.
func GetTTSRuleReplacementClass() TTSRuleReplacementClass {
	return getTTSRuleReplacementClass()
}

type TTSRuleReplacementClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRuleReplacementClass) Alloc() TTSRuleReplacement {
	rv := objc.Send[TTSRuleReplacement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRuleReplacement.EffectiveIndex]
//   - [TTSRuleReplacement.Group]
//   - [TTSRuleReplacement.SetGroup]
//   - [TTSRuleReplacement.Identifier]
//   - [TTSRuleReplacement.IsTerminalRule]
//   - [TTSRuleReplacement.SetIsTerminalRule]
//   - [TTSRuleReplacement.OriginalRulesetIndex]
//   - [TTSRuleReplacement.SetOriginalRulesetIndex]
//   - [TTSRuleReplacement.Regex]
//   - [TTSRuleReplacement.SetRegex]
//   - [TTSRuleReplacement.Replacement]
//   - [TTSRuleReplacement.SetReplacement]
//   - [TTSRuleReplacement.Ruleset]
//   - [TTSRuleReplacement.SetRuleset]
//   - [TTSRuleReplacement.SetIndex]
//   - [TTSRuleReplacement.SetPostMatch]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement
type TTSRuleReplacement struct {
	objectivec.Object
}

// TTSRuleReplacementFromID constructs a [TTSRuleReplacement] from an objc.ID.
func TTSRuleReplacementFromID(id objc.ID) TTSRuleReplacement {
	return TTSRuleReplacement{objectivec.Object{ID: id}}
}
// Ensure TTSRuleReplacement implements ITTSRuleReplacement.
var _ ITTSRuleReplacement = TTSRuleReplacement{}

// An interface definition for the [TTSRuleReplacement] class.
//
// # Methods
//
//   - [ITTSRuleReplacement.EffectiveIndex]
//   - [ITTSRuleReplacement.Group]
//   - [ITTSRuleReplacement.SetGroup]
//   - [ITTSRuleReplacement.Identifier]
//   - [ITTSRuleReplacement.IsTerminalRule]
//   - [ITTSRuleReplacement.SetIsTerminalRule]
//   - [ITTSRuleReplacement.OriginalRulesetIndex]
//   - [ITTSRuleReplacement.SetOriginalRulesetIndex]
//   - [ITTSRuleReplacement.Regex]
//   - [ITTSRuleReplacement.SetRegex]
//   - [ITTSRuleReplacement.Replacement]
//   - [ITTSRuleReplacement.SetReplacement]
//   - [ITTSRuleReplacement.Ruleset]
//   - [ITTSRuleReplacement.SetRuleset]
//   - [ITTSRuleReplacement.SetIndex]
//   - [ITTSRuleReplacement.SetPostMatch]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement
type ITTSRuleReplacement interface {
	objectivec.IObject

	// Topic: Methods

	EffectiveIndex() uint64
	Group() ITTSRuleGroup
	SetGroup(value ITTSRuleGroup)
	Identifier() string
	IsTerminalRule() bool
	SetIsTerminalRule(value bool)
	OriginalRulesetIndex() uint32
	SetOriginalRulesetIndex(value uint32)
	Regex() ITTSRegex
	SetRegex(value ITTSRegex)
	Replacement() string
	SetReplacement(value string)
	Ruleset() ITTSRuleset
	SetRuleset(value ITTSRuleset)
	SetIndex(index uint64)
	SetPostMatch(match VoidHandler)
}

// Init initializes the instance.
func (t TTSRuleReplacement) Init() TTSRuleReplacement {
	rv := objc.Send[TTSRuleReplacement](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRuleReplacement) Autorelease() TTSRuleReplacement {
	rv := objc.Send[TTSRuleReplacement](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRuleReplacement creates a new TTSRuleReplacement instance.
func NewTTSRuleReplacement() TTSRuleReplacement {
	class := getTTSRuleReplacementClass()
	rv := objc.Send[TTSRuleReplacement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/effectiveIndex
func (t TTSRuleReplacement) EffectiveIndex() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("effectiveIndex"))
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/setIndex:
func (t TTSRuleReplacement) SetIndex(index uint64) {
	objc.Send[objc.ID](t.ID, objc.Sel("setIndex:"), index)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/setPostMatch:
func (t TTSRuleReplacement) SetPostMatch(match VoidHandler) {
_block0, _ := NewVoidBlock(match)
	objc.Send[objc.ID](t.ID, objc.Sel("setPostMatch:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/group
func (t TTSRuleReplacement) Group() ITTSRuleGroup {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("group"))
	return TTSRuleGroupFromID(objc.ID(rv))
}
func (t TTSRuleReplacement) SetGroup(value ITTSRuleGroup) {
	objc.Send[struct{}](t.ID, objc.Sel("setGroup:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/identifier
func (t TTSRuleReplacement) Identifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/isTerminalRule
func (t TTSRuleReplacement) IsTerminalRule() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isTerminalRule"))
	return rv
}
func (t TTSRuleReplacement) SetIsTerminalRule(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIsTerminalRule:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/originalRulesetIndex
func (t TTSRuleReplacement) OriginalRulesetIndex() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("originalRulesetIndex"))
	return rv
}
func (t TTSRuleReplacement) SetOriginalRulesetIndex(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setOriginalRulesetIndex:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/regex
func (t TTSRuleReplacement) Regex() ITTSRegex {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regex"))
	return TTSRegexFromID(objc.ID(rv))
}
func (t TTSRuleReplacement) SetRegex(value ITTSRegex) {
	objc.Send[struct{}](t.ID, objc.Sel("setRegex:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/replacement
func (t TTSRuleReplacement) Replacement() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("replacement"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSRuleReplacement) SetReplacement(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setReplacement:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleReplacement/ruleset
func (t TTSRuleReplacement) Ruleset() ITTSRuleset {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleset"))
	return TTSRulesetFromID(objc.ID(rv))
}
func (t TTSRuleReplacement) SetRuleset(value ITTSRuleset) {
	objc.Send[struct{}](t.ID, objc.Sel("setRuleset:"), value)
}

// SetPostMatchSync is a synchronous wrapper around [TTSRuleReplacement.SetPostMatch].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRuleReplacement) SetPostMatchSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetPostMatch(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

