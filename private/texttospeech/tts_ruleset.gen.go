// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (tc TTSRulesetClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRulesetClass) Alloc() TTSRuleset {
	rv := objc.SendIfResponds[TTSRuleset](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

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
	rv := objc.SendIfResponds[TTSRuleset](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRuleset) Autorelease() TTSRuleset {
	rv := objc.SendIfResponds[TTSRuleset](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRuleset creates a new TTSRuleset instance.
func NewTTSRuleset() TTSRuleset {
	class := getTTSRulesetClass()
	rv := objc.SendIfResponds[TTSRuleset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSRuleset) AddRuleReplacement(replacement objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("addRuleReplacement:"), replacement)
}
func (t TTSRuleset) AddRuleString(string_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("addRuleString:"), string_)
}
func (t TTSRuleset) AddRules(rules objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("addRules:"), rules)
}
func (t TTSRuleset) AddRulesFromData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("addRulesFromData:"), data)
}

func (_TTSRulesetClass TTSRulesetClass) ProcessReplacementStringForSpecialCharacters(characters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSRulesetClass.class), objc.Sel("processReplacementStringForSpecialCharacters:"), characters)
	return objectivec.Object{ID: rv}
}
func (_TTSRulesetClass TTSRulesetClass) RulesetWithDataIdentifierPriority(data objectivec.IObject, identifier objectivec.IObject, priority uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSRulesetClass.class), objc.Sel("rulesetWithData:identifier:priority:"), data, identifier, priority)
	return objectivec.Object{ID: rv}
}

func (t TTSRuleset) ActivationRegex() ITTSRegex {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("activationRegex"))
	return TTSRegexFromID(objc.ID(rv))
}
func (t TTSRuleset) SetActivationRegex(value ITTSRegex) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setActivationRegex:"), value)
}
func (t TTSRuleset) Identifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSRuleset) SetIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
func (t TTSRuleset) Priority() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("priority"))
	return rv
}
func (t TTSRuleset) SetPriority(value uint64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setPriority:"), value)
}
func (t TTSRuleset) RuleCount() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("ruleCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSRuleset) RuleReplacements() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("ruleReplacements"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSRuleset) SetRuleReplacements(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setRuleReplacements:"), value)
}
