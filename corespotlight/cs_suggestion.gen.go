// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSSuggestion] class.
var (
	_CSSuggestionClass     CSSuggestionClass
	_CSSuggestionClassOnce sync.Once
)

func getCSSuggestionClass() CSSuggestionClass {
	_CSSuggestionClassOnce.Do(func() {
		_CSSuggestionClass = CSSuggestionClass{class: objc.GetClass("CSSuggestion")}
	})
	return _CSSuggestionClass
}

// GetCSSuggestionClass returns the class object for CSSuggestion.
func GetCSSuggestionClass() CSSuggestionClass {
	return getCSSuggestionClass()
}

type CSSuggestionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSuggestionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSuggestionClass) Alloc() CSSuggestion {
	rv := objc.Send[CSSuggestion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The kind of suggestion to use in a query.
//
// # Overview
//
// Your app uses [CSSuggestion] objects to populate a contextual menu of
// suggestions.
//
// # Setting suggestion attributes
//
//   - [CSSuggestion.LocalizedAttributedSuggestion]: An attributed string for the localized suggestion.
//   - [CSSuggestion.SetLocalizedAttributedSuggestion]
//   - [CSSuggestion.SuggestionKind]: The type of suggestion.
//
// # Comparing suggestions
//
//   - [CSSuggestion.Compare]: Compares the suggestion with a second specified suggestion.
//   - [CSSuggestion.CompareByRank]
//
// # Initializers
//
//   - [CSSuggestion.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion
type CSSuggestion struct {
	objectivec.Object
}

// CSSuggestionFromID constructs a [CSSuggestion] from an objc.ID.
//
// The kind of suggestion to use in a query.
func CSSuggestionFromID(id objc.ID) CSSuggestion {
	return CSSuggestion{objectivec.Object{ID: id}}
}

// NOTE: CSSuggestion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSuggestion] class.
//
// # Setting suggestion attributes
//
//   - [ICSSuggestion.LocalizedAttributedSuggestion]: An attributed string for the localized suggestion.
//   - [ICSSuggestion.SetLocalizedAttributedSuggestion]
//   - [ICSSuggestion.SuggestionKind]: The type of suggestion.
//
// # Comparing suggestions
//
//   - [ICSSuggestion.Compare]: Compares the suggestion with a second specified suggestion.
//   - [ICSSuggestion.CompareByRank]
//
// # Initializers
//
//   - [ICSSuggestion.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion
type ICSSuggestion interface {
	objectivec.IObject

	// Topic: Setting suggestion attributes

	// An attributed string for the localized suggestion.
	LocalizedAttributedSuggestion() foundation.NSAttributedString
	SetLocalizedAttributedSuggestion(value foundation.NSAttributedString)
	// The type of suggestion.
	SuggestionKind() CSSuggestionKind

	// Topic: Comparing suggestions

	// Compares the suggestion with a second specified suggestion.
	Compare(other ICSSuggestion) foundation.ComparisonResult
	CompareByRank(other ICSSuggestion) foundation.ComparisonResult

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CSSuggestion

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSSuggestion) Init() CSSuggestion {
	rv := objc.Send[CSSuggestion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSuggestion) Autorelease() CSSuggestion {
	rv := objc.Send[CSSuggestion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSuggestion creates a new CSSuggestion instance.
func NewCSSuggestion() CSSuggestion {
	class := getCSSuggestionClass()
	rv := objc.Send[CSSuggestion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/init(coder:)
func NewCSSuggestionWithCoder(coder foundation.INSCoder) CSSuggestion {
	instance := getCSSuggestionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSSuggestionFromID(rv)
}

// Compares the suggestion with a second specified suggestion.
//
// other: The suggestion to compare to this suggestion.
//
// # Return Value
//
// Returns an [NSComparisonResult].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(_:)
func (c CSSuggestion) Compare(other ICSSuggestion) foundation.ComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](c.ID, objc.Sel("compare:"), other)
	return foundation.ComparisonResult(rv)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/compare(byRank:)
func (c CSSuggestion) CompareByRank(other ICSSuggestion) foundation.ComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](c.ID, objc.Sel("compareByRank:"), other)
	return foundation.ComparisonResult(rv)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/init(coder:)
func (c CSSuggestion) InitWithCoder(coder foundation.INSCoder) CSSuggestion {
	rv := objc.Send[CSSuggestion](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CSSuggestion) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An attributed string for the localized suggestion.
//
// See: https://developer.apple.com/documentation/corespotlight/cssuggestion/localizedattributedsuggestion-3ssly
func (c CSSuggestion) LocalizedAttributedSuggestion() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedAttributedSuggestion"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (c CSSuggestion) SetLocalizedAttributedSuggestion(value foundation.NSAttributedString) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocalizedAttributedSuggestion:"), value)
}

// The type of suggestion.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/suggestionKind-swift.property
func (c CSSuggestion) SuggestionKind() CSSuggestionKind {
	rv := objc.Send[CSSuggestionKind](c.ID, objc.Sel("suggestionKind"))
	return CSSuggestionKind(rv)
}
