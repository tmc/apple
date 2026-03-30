// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRulesetRunner] class.
var (
	_TTSRulesetRunnerClass     TTSRulesetRunnerClass
	_TTSRulesetRunnerClassOnce sync.Once
)

func getTTSRulesetRunnerClass() TTSRulesetRunnerClass {
	_TTSRulesetRunnerClassOnce.Do(func() {
		_TTSRulesetRunnerClass = TTSRulesetRunnerClass{class: objc.GetClass("TTSRulesetRunner")}
	})
	return _TTSRulesetRunnerClass
}

// GetTTSRulesetRunnerClass returns the class object for TTSRulesetRunner.
func GetTTSRulesetRunnerClass() TTSRulesetRunnerClass {
	return getTTSRulesetRunnerClass()
}

type TTSRulesetRunnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSRulesetRunnerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRulesetRunnerClass) Alloc() TTSRulesetRunner {
	rv := objc.Send[TTSRulesetRunner](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSRulesetRunner._computeActiveRangesWithIgnoreRanges]
//   - [TTSRulesetRunner._ignoreRangesForString]
//   - [TTSRulesetRunner._processSpeechStringStartingAtCurrentRecursionDepth]
//   - [TTSRulesetRunner._processTemplateReplacementTextForTextReplacementCString]
//   - [TTSRulesetRunner._recomputeRuleOrdering]
//   - [TTSRulesetRunner.CancelProcessing]
//   - [TTSRulesetRunner.Executing]
//   - [TTSRulesetRunner.SetExecuting]
//   - [TTSRulesetRunner.LoadRuleSet]
//   - [TTSRulesetRunner.ProcessText]
//   - [TTSRulesetRunner.RegexExecutionQueue]
//   - [TTSRulesetRunner.SetRegexExecutionQueue]
//   - [TTSRulesetRunner.Reset]
//   - [TTSRulesetRunner.RuleCount]
//   - [TTSRulesetRunner.RuleReplacements]
//   - [TTSRulesetRunner.SetRuleReplacements]
//   - [TTSRulesetRunner.RuleSets]
//   - [TTSRulesetRunner.SetRuleSets]
//   - [TTSRulesetRunner.SetMatchLogger]
//   - [TTSRulesetRunner.SetPostRuleWriter]
//   - [TTSRulesetRunner.SetPreRuleWriter]
//   - [TTSRulesetRunner.ShouldAbort]
//   - [TTSRulesetRunner.SetShouldAbort]
//   - [TTSRulesetRunner.UnloadRuleset]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner
type TTSRulesetRunner struct {
	objectivec.Object
}

// TTSRulesetRunnerFromID constructs a [TTSRulesetRunner] from an objc.ID.
func TTSRulesetRunnerFromID(id objc.ID) TTSRulesetRunner {
	return TTSRulesetRunner{objectivec.Object{ID: id}}
}

// Ensure TTSRulesetRunner implements ITTSRulesetRunner.
var _ ITTSRulesetRunner = TTSRulesetRunner{}

// An interface definition for the [TTSRulesetRunner] class.
//
// # Methods
//
//   - [ITTSRulesetRunner._computeActiveRangesWithIgnoreRanges]
//   - [ITTSRulesetRunner._ignoreRangesForString]
//   - [ITTSRulesetRunner._processSpeechStringStartingAtCurrentRecursionDepth]
//   - [ITTSRulesetRunner._processTemplateReplacementTextForTextReplacementCString]
//   - [ITTSRulesetRunner._recomputeRuleOrdering]
//   - [ITTSRulesetRunner.CancelProcessing]
//   - [ITTSRulesetRunner.Executing]
//   - [ITTSRulesetRunner.SetExecuting]
//   - [ITTSRulesetRunner.LoadRuleSet]
//   - [ITTSRulesetRunner.ProcessText]
//   - [ITTSRulesetRunner.RegexExecutionQueue]
//   - [ITTSRulesetRunner.SetRegexExecutionQueue]
//   - [ITTSRulesetRunner.Reset]
//   - [ITTSRulesetRunner.RuleCount]
//   - [ITTSRulesetRunner.RuleReplacements]
//   - [ITTSRulesetRunner.SetRuleReplacements]
//   - [ITTSRulesetRunner.RuleSets]
//   - [ITTSRulesetRunner.SetRuleSets]
//   - [ITTSRulesetRunner.SetMatchLogger]
//   - [ITTSRulesetRunner.SetPostRuleWriter]
//   - [ITTSRulesetRunner.SetPreRuleWriter]
//   - [ITTSRulesetRunner.ShouldAbort]
//   - [ITTSRulesetRunner.SetShouldAbort]
//   - [ITTSRulesetRunner.UnloadRuleset]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner
type ITTSRulesetRunner interface {
	objectivec.IObject

	// Topic: Methods

	_computeActiveRangesWithIgnoreRanges(ranges objectivec.IObject, ranges2 objectivec.IObject) objectivec.IObject
	_ignoreRangesForString(string_ objectivec.IObject) objectivec.IObject
	_processSpeechStringStartingAtCurrentRecursionDepth(string_ objectivec.IObject, at uint64, depth uint64) objectivec.IObject
	_processTemplateReplacementTextForTextReplacementCString(text objectivec.IObject, replacement objectivec.IObject, string_ string) objectivec.IObject
	_recomputeRuleOrdering()
	CancelProcessing()
	Executing() bool
	SetExecuting(value bool)
	LoadRuleSet(set objectivec.IObject)
	ProcessText(text objectivec.IObject) objectivec.IObject
	RegexExecutionQueue() objectivec.Object
	SetRegexExecutionQueue(value objectivec.Object)
	Reset()
	RuleCount() foundation.NSNumber
	RuleReplacements() foundation.INSArray
	SetRuleReplacements(value foundation.INSArray)
	RuleSets() foundation.INSArray
	SetRuleSets(value foundation.INSArray)
	SetMatchLogger(logger VoidHandler)
	SetPostRuleWriter(writer VoidHandler)
	SetPreRuleWriter(writer VoidHandler)
	ShouldAbort() bool
	SetShouldAbort(value bool)
	UnloadRuleset(ruleset objectivec.IObject)
}

// Init initializes the instance.
func (t TTSRulesetRunner) Init() TTSRulesetRunner {
	rv := objc.Send[TTSRulesetRunner](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRulesetRunner) Autorelease() TTSRulesetRunner {
	rv := objc.Send[TTSRulesetRunner](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRulesetRunner creates a new TTSRulesetRunner instance.
func NewTTSRulesetRunner() TTSRulesetRunner {
	class := getTTSRulesetRunnerClass()
	rv := objc.Send[TTSRulesetRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/_computeActiveRanges:withIgnoreRanges:
func (t TTSRulesetRunner) _computeActiveRangesWithIgnoreRanges(ranges objectivec.IObject, ranges2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_computeActiveRanges:withIgnoreRanges:"), ranges, ranges2)
	return objectivec.Object{ID: rv}
}

// ComputeActiveRangesWithIgnoreRanges is an exported wrapper for the private method _computeActiveRangesWithIgnoreRanges.
func (t TTSRulesetRunner) ComputeActiveRangesWithIgnoreRanges(ranges objectivec.IObject, ranges2 objectivec.IObject) objectivec.IObject {
	return t._computeActiveRangesWithIgnoreRanges(ranges, ranges2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/_ignoreRangesForString:
func (t TTSRulesetRunner) _ignoreRangesForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_ignoreRangesForString:"), string_)
	return objectivec.Object{ID: rv}
}

// IgnoreRangesForString is an exported wrapper for the private method _ignoreRangesForString.
func (t TTSRulesetRunner) IgnoreRangesForString(string_ objectivec.IObject) objectivec.IObject {
	return t._ignoreRangesForString(string_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/_processSpeechString:startingAt:currentRecursionDepth:
func (t TTSRulesetRunner) _processSpeechStringStartingAtCurrentRecursionDepth(string_ objectivec.IObject, at uint64, depth uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_processSpeechString:startingAt:currentRecursionDepth:"), string_, at, depth)
	return objectivec.Object{ID: rv}
}

// ProcessSpeechStringStartingAtCurrentRecursionDepth is an exported wrapper for the private method _processSpeechStringStartingAtCurrentRecursionDepth.
func (t TTSRulesetRunner) ProcessSpeechStringStartingAtCurrentRecursionDepth(string_ objectivec.IObject, at uint64, depth uint64) objectivec.IObject {
	return t._processSpeechStringStartingAtCurrentRecursionDepth(string_, at, depth)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/_processTemplateReplacementTextForText:replacement:cString:
func (t TTSRulesetRunner) _processTemplateReplacementTextForTextReplacementCString(text objectivec.IObject, replacement objectivec.IObject, string_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_processTemplateReplacementTextForText:replacement:cString:"), text, replacement, unsafe.Pointer(unsafe.StringData(string_+"\x00")))
	return objectivec.Object{ID: rv}
}

// ProcessTemplateReplacementTextForTextReplacementCString is an exported wrapper for the private method _processTemplateReplacementTextForTextReplacementCString.
func (t TTSRulesetRunner) ProcessTemplateReplacementTextForTextReplacementCString(text objectivec.IObject, replacement objectivec.IObject, string_ string) objectivec.IObject {
	return t._processTemplateReplacementTextForTextReplacementCString(text, replacement, string_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/_recomputeRuleOrdering
func (t TTSRulesetRunner) _recomputeRuleOrdering() {
	objc.Send[objc.ID](t.ID, objc.Sel("_recomputeRuleOrdering"))
}

// RecomputeRuleOrdering is an exported wrapper for the private method _recomputeRuleOrdering.
func (t TTSRulesetRunner) RecomputeRuleOrdering() {
	t._recomputeRuleOrdering()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/cancelProcessing
func (t TTSRulesetRunner) CancelProcessing() {
	objc.Send[objc.ID](t.ID, objc.Sel("cancelProcessing"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/loadRuleSet:
func (t TTSRulesetRunner) LoadRuleSet(set objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("loadRuleSet:"), set)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/processText:
func (t TTSRulesetRunner) ProcessText(text objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("processText:"), text)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/reset
func (t TTSRulesetRunner) Reset() {
	objc.Send[objc.ID](t.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/setMatchLogger:
func (t TTSRulesetRunner) SetMatchLogger(logger VoidHandler) {
	_block0, _ := NewVoidBlock(logger)
	objc.Send[objc.ID](t.ID, objc.Sel("setMatchLogger:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/setPostRuleWriter:
func (t TTSRulesetRunner) SetPostRuleWriter(writer VoidHandler) {
	_block0, _ := NewVoidBlock(writer)
	objc.Send[objc.ID](t.ID, objc.Sel("setPostRuleWriter:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/setPreRuleWriter:
func (t TTSRulesetRunner) SetPreRuleWriter(writer VoidHandler) {
	_block0, _ := NewVoidBlock(writer)
	objc.Send[objc.ID](t.ID, objc.Sel("setPreRuleWriter:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/unloadRuleset:
func (t TTSRulesetRunner) UnloadRuleset(ruleset objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("unloadRuleset:"), ruleset)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/executing
func (t TTSRulesetRunner) Executing() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("executing"))
	return rv
}
func (t TTSRulesetRunner) SetExecuting(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setExecuting:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/regexExecutionQueue
func (t TTSRulesetRunner) RegexExecutionQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regexExecutionQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSRulesetRunner) SetRegexExecutionQueue(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setRegexExecutionQueue:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/ruleCount
func (t TTSRulesetRunner) RuleCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/ruleReplacements
func (t TTSRulesetRunner) RuleReplacements() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleReplacements"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSRulesetRunner) SetRuleReplacements(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setRuleReplacements:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/ruleSets
func (t TTSRulesetRunner) RuleSets() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ruleSets"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSRulesetRunner) SetRuleSets(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setRuleSets:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRulesetRunner/shouldAbort
func (t TTSRulesetRunner) ShouldAbort() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldAbort"))
	return rv
}
func (t TTSRulesetRunner) SetShouldAbort(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldAbort:"), value)
}

// SetMatchLoggerSync is a synchronous wrapper around [TTSRulesetRunner.SetMatchLogger].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRulesetRunner) SetMatchLoggerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetMatchLogger(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPostRuleWriterSync is a synchronous wrapper around [TTSRulesetRunner.SetPostRuleWriter].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRulesetRunner) SetPostRuleWriterSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetPostRuleWriter(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPreRuleWriterSync is a synchronous wrapper around [TTSRulesetRunner.SetPreRuleWriter].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRulesetRunner) SetPreRuleWriterSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetPreRuleWriter(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
