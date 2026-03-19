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

	// AutocorrectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/autocorrectionType
	AutocorrectionType() NSTextInputTraitType

	// DataDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/dataDetectionType
	DataDetectionType() NSTextInputTraitType

	// GrammarCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/grammarCheckingType
	GrammarCheckingType() NSTextInputTraitType

	// InlinePredictionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/inlinePredictionType
	InlinePredictionType() NSTextInputTraitType

	// LinkDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/linkDetectionType
	LinkDetectionType() NSTextInputTraitType

	// SmartDashesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartDashesType
	SmartDashesType() NSTextInputTraitType

	// SmartInsertDeleteType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartInsertDeleteType
	SmartInsertDeleteType() NSTextInputTraitType

	// SmartQuotesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartQuotesType
	SmartQuotesType() NSTextInputTraitType

	// SpellCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/spellCheckingType
	SpellCheckingType() NSTextInputTraitType

	// TextCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textCompletionType
	TextCompletionType() NSTextInputTraitType

	// TextReplacementType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textReplacementType
	TextReplacementType() NSTextInputTraitType

	// AllowedWritingToolsResultOptions protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/allowedWritingToolsResultOptions
	AllowedWritingToolsResultOptions() NSWritingToolsResultOptions

	// MathExpressionCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/mathExpressionCompletionType
	MathExpressionCompletionType() NSTextInputTraitType

	// WritingToolsBehavior protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/writingToolsBehavior
	WritingToolsBehavior() NSWritingToolsBehavior

	// autocorrectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/autocorrectionType
	SetAutocorrectionType(value NSTextInputTraitType)

	// dataDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/dataDetectionType
	SetDataDetectionType(value NSTextInputTraitType)

	// grammarCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/grammarCheckingType
	SetGrammarCheckingType(value NSTextInputTraitType)

	// inlinePredictionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/inlinePredictionType
	SetInlinePredictionType(value NSTextInputTraitType)

	// linkDetectionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/linkDetectionType
	SetLinkDetectionType(value NSTextInputTraitType)

	// smartDashesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartDashesType
	SetSmartDashesType(value NSTextInputTraitType)

	// smartInsertDeleteType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartInsertDeleteType
	SetSmartInsertDeleteType(value NSTextInputTraitType)

	// smartQuotesType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartQuotesType
	SetSmartQuotesType(value NSTextInputTraitType)

	// spellCheckingType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/spellCheckingType
	SetSpellCheckingType(value NSTextInputTraitType)

	// textCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textCompletionType
	SetTextCompletionType(value NSTextInputTraitType)

	// textReplacementType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textReplacementType
	SetTextReplacementType(value NSTextInputTraitType)

	// allowedWritingToolsResultOptions protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/allowedWritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions)

	// mathExpressionCompletionType protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/mathExpressionCompletionType
	SetMathExpressionCompletionType(value NSTextInputTraitType)

	// writingToolsBehavior protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/writingToolsBehavior
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
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/dataDetectionType
func (o NSTextInputTraitsObject) DataDetectionType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("dataDetectionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/grammarCheckingType
func (o NSTextInputTraitsObject) GrammarCheckingType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("grammarCheckingType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/inlinePredictionType
func (o NSTextInputTraitsObject) InlinePredictionType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("inlinePredictionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/linkDetectionType
func (o NSTextInputTraitsObject) LinkDetectionType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("linkDetectionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartDashesType
func (o NSTextInputTraitsObject) SmartDashesType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartDashesType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartInsertDeleteType
func (o NSTextInputTraitsObject) SmartInsertDeleteType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartInsertDeleteType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/smartQuotesType
func (o NSTextInputTraitsObject) SmartQuotesType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("smartQuotesType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/spellCheckingType
func (o NSTextInputTraitsObject) SpellCheckingType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("spellCheckingType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textCompletionType
func (o NSTextInputTraitsObject) TextCompletionType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("textCompletionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/textReplacementType
func (o NSTextInputTraitsObject) TextReplacementType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("textReplacementType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/allowedWritingToolsResultOptions
func (o NSTextInputTraitsObject) AllowedWritingToolsResultOptions() NSWritingToolsResultOptions {
	
	rv := objc.Send[NSWritingToolsResultOptions](o.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/mathExpressionCompletionType
func (o NSTextInputTraitsObject) MathExpressionCompletionType() NSTextInputTraitType {
	
	rv := objc.Send[NSTextInputTraitType](o.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputTraits/writingToolsBehavior
func (o NSTextInputTraitsObject) WritingToolsBehavior() NSWritingToolsBehavior {
	
	rv := objc.Send[NSWritingToolsBehavior](o.ID, objc.Sel("writingToolsBehavior"))
	return rv
	}

func (o NSTextInputTraitsObject) SetAutocorrectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutocorrectionType:"), value)
}

func (o NSTextInputTraitsObject) SetDataDetectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setDataDetectionType:"), value)
}

func (o NSTextInputTraitsObject) SetGrammarCheckingType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setGrammarCheckingType:"), value)
}

func (o NSTextInputTraitsObject) SetInlinePredictionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setInlinePredictionType:"), value)
}

func (o NSTextInputTraitsObject) SetLinkDetectionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setLinkDetectionType:"), value)
}

func (o NSTextInputTraitsObject) SetSmartDashesType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartDashesType:"), value)
}

func (o NSTextInputTraitsObject) SetSmartInsertDeleteType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartInsertDeleteType:"), value)
}

func (o NSTextInputTraitsObject) SetSmartQuotesType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmartQuotesType:"), value)
}

func (o NSTextInputTraitsObject) SetSpellCheckingType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setSpellCheckingType:"), value)
}

func (o NSTextInputTraitsObject) SetTextCompletionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextCompletionType:"), value)
}

func (o NSTextInputTraitsObject) SetTextReplacementType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextReplacementType:"), value)
}

func (o NSTextInputTraitsObject) SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}

func (o NSTextInputTraitsObject) SetMathExpressionCompletionType(value NSTextInputTraitType) {
	objc.Send[struct{}](o.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}

func (o NSTextInputTraitsObject) SetWritingToolsBehavior(value NSWritingToolsBehavior) {
	objc.Send[struct{}](o.ID, objc.Sel("setWritingToolsBehavior:"), value)
}

