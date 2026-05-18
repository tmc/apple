// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSTextInputTraits protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits
type NSTextInputTraits interface {
	objectivec.IObject

	// autocorrectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/autocorrectionType
	AutocorrectionType() NSTextInputTraitType
	SetAutocorrectionType(value NSTextInputTraitType)

	// dataDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/dataDetectionType
	DataDetectionType() NSTextInputTraitType
	SetDataDetectionType(value NSTextInputTraitType)

	// grammarCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/grammarCheckingType
	GrammarCheckingType() NSTextInputTraitType
	SetGrammarCheckingType(value NSTextInputTraitType)

	// inlinePredictionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/inlinePredictionType
	InlinePredictionType() NSTextInputTraitType
	SetInlinePredictionType(value NSTextInputTraitType)

	// linkDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/linkDetectionType
	LinkDetectionType() NSTextInputTraitType
	SetLinkDetectionType(value NSTextInputTraitType)

	// smartDashesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartDashesType
	SmartDashesType() NSTextInputTraitType
	SetSmartDashesType(value NSTextInputTraitType)

	// smartInsertDeleteType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartInsertDeleteType
	SmartInsertDeleteType() NSTextInputTraitType
	SetSmartInsertDeleteType(value NSTextInputTraitType)

	// smartQuotesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartQuotesType
	SmartQuotesType() NSTextInputTraitType
	SetSmartQuotesType(value NSTextInputTraitType)

	// spellCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/spellCheckingType
	SpellCheckingType() NSTextInputTraitType
	SetSpellCheckingType(value NSTextInputTraitType)

	// textCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textCompletionType
	TextCompletionType() NSTextInputTraitType
	SetTextCompletionType(value NSTextInputTraitType)

	// textReplacementType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textReplacementType
	TextReplacementType() NSTextInputTraitType
	SetTextReplacementType(value NSTextInputTraitType)

	// allowedWritingToolsResultOptions protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/allowedWritingToolsResultOptions
	AllowedWritingToolsResultOptions() NSWritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions)

	// mathExpressionCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/mathExpressionCompletionType
	MathExpressionCompletionType() NSTextInputTraitType
	SetMathExpressionCompletionType(value NSTextInputTraitType)

	// writingToolsBehavior protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/writingToolsBehavior
	WritingToolsBehavior() NSWritingToolsBehavior
	SetWritingToolsBehavior(value NSWritingToolsBehavior)
}

// NSTextInputTraitsObject wraps an existing Objective-C object that conforms to the NSTextInputTraits protocol.
type NSTextInputTraitsObject struct {
	objectivec.Object
}

func (o NSTextInputTraitsObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextInputTraitsObjectFromID constructs a [NSTextInputTraitsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextInputTraitsObjectFromID(id objc.ID) NSTextInputTraitsObject {
	return NSTextInputTraitsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/autocorrectionType
func (o NSTextInputTraitsObject) AutocorrectionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("autocorrectionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetAutocorrectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutocorrectionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/dataDetectionType
func (o NSTextInputTraitsObject) DataDetectionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("dataDetectionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetDataDetectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setDataDetectionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/grammarCheckingType
func (o NSTextInputTraitsObject) GrammarCheckingType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("grammarCheckingType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetGrammarCheckingType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setGrammarCheckingType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/inlinePredictionType
func (o NSTextInputTraitsObject) InlinePredictionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("inlinePredictionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetInlinePredictionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setInlinePredictionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/linkDetectionType
func (o NSTextInputTraitsObject) LinkDetectionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("linkDetectionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetLinkDetectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setLinkDetectionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartDashesType
func (o NSTextInputTraitsObject) SmartDashesType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartDashesType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetSmartDashesType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartDashesType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartInsertDeleteType
func (o NSTextInputTraitsObject) SmartInsertDeleteType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartInsertDeleteType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetSmartInsertDeleteType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartInsertDeleteType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartQuotesType
func (o NSTextInputTraitsObject) SmartQuotesType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartQuotesType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetSmartQuotesType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartQuotesType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/spellCheckingType
func (o NSTextInputTraitsObject) SpellCheckingType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("spellCheckingType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetSpellCheckingType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSpellCheckingType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textCompletionType
func (o NSTextInputTraitsObject) TextCompletionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("textCompletionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetTextCompletionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextCompletionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textReplacementType
func (o NSTextInputTraitsObject) TextReplacementType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("textReplacementType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetTextReplacementType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextReplacementType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/allowedWritingToolsResultOptions
func (o NSTextInputTraitsObject) AllowedWritingToolsResultOptions() NSWritingToolsResultOptions {
	rv := objc.Send[NSWritingToolsResultOptions](o.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return NSWritingToolsResultOptions(rv)
}

func (o NSTextInputTraitsObject) SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/mathExpressionCompletionType
func (o NSTextInputTraitsObject) MathExpressionCompletionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("mathExpressionCompletionType"))
	return NSTextInputTraitType(rv)
}

func (o NSTextInputTraitsObject) SetMathExpressionCompletionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/writingToolsBehavior
func (o NSTextInputTraitsObject) WritingToolsBehavior() NSWritingToolsBehavior {
	rv := objc.Send[NSWritingToolsBehavior](o.ID, objc.Sel("writingToolsBehavior"))
	return NSWritingToolsBehavior(rv)
}

func (o NSTextInputTraitsObject) SetWritingToolsBehavior(value NSWritingToolsBehavior) {
	objc.Send[struct{}](o.ID, objc.Sel("setWritingToolsBehavior:"), value)
}
