// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRuleset] class.
var (
	_TTSRulesetClass     TTSRulesetClass
	_TTSRulesetClassOnce sync.Once
)

func getTTSRulesetClass() TTSRulesetClass {
	_TTSRulesetClassOnce.Do(func() {
		_TTSRulesetClass = TTSRulesetClass{class: objc.GetClass("TTSRuleset")}
	})
	return _TTSRulesetClass
}

// GetTTSRulesetClass returns the class object for TTSRuleset.
func GetTTSRulesetClass() TTSRulesetClass {
	return getTTSRulesetClass()
}

type TTSRulesetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRulesetClass) Alloc() TTSRuleset {
	rv := objc.Send[TTSRuleset](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRuleset.ActivationRegex]
//   - [TTSRuleset.SetActivationRegex]
//   - [TTSRuleset.AddRuleReplacement]
//   - [TTSRuleset.AddRuleString]
//   - [TTSRuleset.AddRules]
//   - [TTSRuleset.AddRulesFromData]
//   - [TTSRuleset.Identifier]
//   - [TTSRuleset.SetIdentifier]
//   - [TTSRuleset.Priority]
//   - [TTSRuleset.SetPriority]
//   - [TTSRuleset.RuleCount]
//   - [TTSRuleset.RuleReplacements]
//   - [TTSRuleset.SetRuleReplacements]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset
type TTSRuleset struct {
	objectivec.Object
}

// TTSRulesetFromID constructs a [TTSRuleset] from an objc.ID.
func TTSRulesetFromID(id objc.ID) TTSRuleset {
	return TTSRuleset{objectivec.Object{ID: id}}
}
// Ensure TTSRuleset implements ITTSRuleset.
var _ ITTSRuleset = TTSRuleset{}

// An interface definition for the [TTSRuleset] class.
//
// # Methods
//
//   - [ITTSRuleset.ActivationRegex]
//   - [ITTSRuleset.SetActivationRegex]
//   - [ITTSRuleset.AddRuleReplacement]
//   - [ITTSRuleset.AddRuleString]
//   - [ITTSRuleset.AddRules]
//   - [ITTSRuleset.AddRulesFromData]
//   - [ITTSRuleset.Identifier]
//   - [ITTSRuleset.SetIdentifier]
//   - [ITTSRuleset.Priority]
//   - [ITTSRuleset.SetPriority]
//   - [ITTSRuleset.RuleCount]
//   - [ITTSRuleset.RuleReplacements]
//   - [ITTSRuleset.SetRuleReplacements]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset
type ITTSRuleset interface {
	objectivec.IObject

	// Topic: Methods

	ActivationRegex() ITTSRegex
	SetActivationRegex(value ITTSRegex)
	AddRuleReplacement(replacement objectivec.IObject)
	AddRuleString(string_ objectivec.IObject)
	AddRules(rules objectivec.IObject)
	AddRulesFromData(data objectivec.IObject)
	Identifier() string
	SetIdentifier(value string)
	Priority() uint64
	SetPriority(value uint64)
	RuleCount() foundation.NSNumber
	RuleReplacements() foundation.INSArray
	SetRuleReplacements(value foundation.INSArray)
}

// Init initializes the instance.
func (t TTSRuleset) Init() TTSRuleset {
	rv := objc.Send[TTSRuleset](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRuleset) Autorelease() TTSRuleset {
	rv := objc.Send[TTSRuleset](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRuleset creates a new TTSRuleset instance.
func NewTTSRuleset() TTSRuleset {
	class := getTTSRulesetClass()
	rv := objc.Send[TTSRuleset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/addRuleReplacement:
func (t TTSRuleset) AddRuleReplacement(replacement objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addRuleReplacement:"), replacement)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/addRuleString:
func (t TTSRuleset) AddRuleString(string_ objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addRuleString:"), string_)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/addRules:
func (t TTSRuleset) AddRules(rules objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addRules:"), rules)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/addRulesFromData:
func (t TTSRuleset) AddRulesFromData(data objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addRulesFromData:"), data)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/processReplacementStringForSpecialCharacters:
func (_TTSRulesetClass TTSRulesetClass) ProcessReplacementStringForSpecialCharacters(characters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSRulesetClass.class), objc.Sel("processReplacementStringForSpecialCharacters:"), characters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/rulesetWithData:identifier:priority:
func (_TTSRulesetClass TTSRulesetClass) RulesetWithDataIdentifierPriority(data objectivec.IObject, identifier objectivec.IObject, priority uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSRulesetClass.class), objc.Sel("rulesetWithData:identifier:priority:"), data, identifier, priority)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/activationRegex
func (t TTSRuleset) ActivationRegex() ITTSRegex {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("activationRegex"))
	return TTSRegexFromID(objc.ID(rv))
}
func (t TTSRuleset) SetActivationRegex(value ITTSRegex) {
	objc.Send[struct{}](t.ID, objc.Sel("setActivationRegex:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/identifier
func (t TTSRuleset) Identifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSRuleset) SetIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/priority
func (t TTSRuleset) Priority() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("priority"))
	return rv
}
func (t TTSRuleset) SetPriority(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setPriority:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/ruleCount
func (t TTSRuleset) RuleCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRuleset/ruleReplacements
func (t TTSRuleset) RuleReplacements() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleReplacements"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSRuleset) SetRuleReplacements(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setRuleReplacements:"), value)
}

