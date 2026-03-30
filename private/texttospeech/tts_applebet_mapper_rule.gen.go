// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSApplebetMapperRule] class.
var (
	_TTSApplebetMapperRuleClass     TTSApplebetMapperRuleClass
	_TTSApplebetMapperRuleClassOnce sync.Once
)

func getTTSApplebetMapperRuleClass() TTSApplebetMapperRuleClass {
	_TTSApplebetMapperRuleClassOnce.Do(func() {
		_TTSApplebetMapperRuleClass = TTSApplebetMapperRuleClass{class: objc.GetClass("TTSApplebetMapperRule")}
	})
	return _TTSApplebetMapperRuleClass
}

// GetTTSApplebetMapperRuleClass returns the class object for TTSApplebetMapperRule.
func GetTTSApplebetMapperRuleClass() TTSApplebetMapperRuleClass {
	return getTTSApplebetMapperRuleClass()
}

type TTSApplebetMapperRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSApplebetMapperRuleClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSApplebetMapperRuleClass) Alloc() TTSApplebetMapperRule {
	rv := objc.Send[TTSApplebetMapperRule](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSApplebetMapperRule.Left]
//   - [TTSApplebetMapperRule.SetLeft]
//   - [TTSApplebetMapperRule.Match]
//   - [TTSApplebetMapperRule.SetMatch]
//   - [TTSApplebetMapperRule.Right]
//   - [TTSApplebetMapperRule.SetRight]
//   - [TTSApplebetMapperRule.SetMatchRule]
//   - [TTSApplebetMapperRule.Substitution]
//   - [TTSApplebetMapperRule.SetSubstitution]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule
type TTSApplebetMapperRule struct {
	objectivec.Object
}

// TTSApplebetMapperRuleFromID constructs a [TTSApplebetMapperRule] from an objc.ID.
func TTSApplebetMapperRuleFromID(id objc.ID) TTSApplebetMapperRule {
	return TTSApplebetMapperRule{objectivec.Object{ID: id}}
}

// Ensure TTSApplebetMapperRule implements ITTSApplebetMapperRule.
var _ ITTSApplebetMapperRule = TTSApplebetMapperRule{}

// An interface definition for the [TTSApplebetMapperRule] class.
//
// # Methods
//
//   - [ITTSApplebetMapperRule.Left]
//   - [ITTSApplebetMapperRule.SetLeft]
//   - [ITTSApplebetMapperRule.Match]
//   - [ITTSApplebetMapperRule.SetMatch]
//   - [ITTSApplebetMapperRule.Right]
//   - [ITTSApplebetMapperRule.SetRight]
//   - [ITTSApplebetMapperRule.SetMatchRule]
//   - [ITTSApplebetMapperRule.Substitution]
//   - [ITTSApplebetMapperRule.SetSubstitution]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule
type ITTSApplebetMapperRule interface {
	objectivec.IObject

	// Topic: Methods

	Left() foundation.INSArray
	SetLeft(value foundation.INSArray)
	Match() foundation.INSArray
	SetMatch(value foundation.INSArray)
	Right() foundation.INSArray
	SetRight(value foundation.INSArray)
	SetMatchRule(rule VoidHandler)
	Substitution() foundation.INSArray
	SetSubstitution(value foundation.INSArray)
}

// Init initializes the instance.
func (t TTSApplebetMapperRule) Init() TTSApplebetMapperRule {
	rv := objc.Send[TTSApplebetMapperRule](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSApplebetMapperRule) Autorelease() TTSApplebetMapperRule {
	rv := objc.Send[TTSApplebetMapperRule](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSApplebetMapperRule creates a new TTSApplebetMapperRule instance.
func NewTTSApplebetMapperRule() TTSApplebetMapperRule {
	class := getTTSApplebetMapperRuleClass()
	rv := objc.Send[TTSApplebetMapperRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/setMatchRule:
func (t TTSApplebetMapperRule) SetMatchRule(rule VoidHandler) {
	_block0, _ := NewVoidBlock(rule)
	objc.Send[objc.ID](t.ID, objc.Sel("setMatchRule:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/ruleWithHeterogeniousArray:
func (_TTSApplebetMapperRuleClass TTSApplebetMapperRuleClass) RuleWithHeterogeniousArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSApplebetMapperRuleClass.class), objc.Sel("ruleWithHeterogeniousArray:"), array)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/left
func (t TTSApplebetMapperRule) Left() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("left"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSApplebetMapperRule) SetLeft(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setLeft:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/match
func (t TTSApplebetMapperRule) Match() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("match"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSApplebetMapperRule) SetMatch(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setMatch:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/right
func (t TTSApplebetMapperRule) Right() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("right"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSApplebetMapperRule) SetRight(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setRight:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSApplebetMapperRule/substitution
func (t TTSApplebetMapperRule) Substitution() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("substitution"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSApplebetMapperRule) SetSubstitution(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setSubstitution:"), value)
}

// SetMatchRuleSync is a synchronous wrapper around [TTSApplebetMapperRule.SetMatchRule].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSApplebetMapperRule) SetMatchRuleSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetMatchRule(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
