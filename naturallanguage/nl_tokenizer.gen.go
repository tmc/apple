// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLTokenizer] class.
var (
	_NLTokenizerClass     NLTokenizerClass
	_NLTokenizerClassOnce sync.Once
)

func getNLTokenizerClass() NLTokenizerClass {
	_NLTokenizerClassOnce.Do(func() {
		_NLTokenizerClass = NLTokenizerClass{class: objc.GetClass("NLTokenizer")}
	})
	return _NLTokenizerClass
}

// GetNLTokenizerClass returns the class object for NLTokenizer.
func GetNLTokenizerClass() NLTokenizerClass {
	return getNLTokenizerClass()
}

type NLTokenizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLTokenizerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLTokenizerClass) Alloc() NLTokenizer {
	rv := objc.Send[NLTokenizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A tokenizer that segments natural language text into semantic units.
//
// # Overview
//
// [NLTokenizer] creates individual units from natural language text. Define
// the desired unit (word, sentence, paragraph, or document as declared in the
// [NLTokenUnit]) for tokenization, and then assign a string to tokenize. The
// [NLTokenizer.EnumerateTokensInRangeUsingBlock] method provides the ranges of the tokens
// in the string based on the tokenization unit.
//
// For more information, see [Tokenizing natural language text].
//
// # Creating a tokenizer
//
//   - [NLTokenizer.InitWithUnit]: Creates a tokenizer with the specified unit.
//
// # Configuring a tokenizer
//
//   - [NLTokenizer.String]: The text to be tokenized.
//   - [NLTokenizer.SetString]
//   - [NLTokenizer.SetLanguage]: Sets the language of the text to be tokenized.
//   - [NLTokenizer.Unit]: The linguistic unit that this tokenizer uses.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer
//
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
// [Tokenizing natural language text]: https://developer.apple.com/documentation/NaturalLanguage/tokenizing-natural-language-text
type NLTokenizer struct {
	objectivec.Object
}

// NLTokenizerFromID constructs a [NLTokenizer] from an objc.ID.
//
// A tokenizer that segments natural language text into semantic units.
func NLTokenizerFromID(id objc.ID) NLTokenizer {
	return NLTokenizer{objectivec.Object{ID: id}}
}

// NOTE: NLTokenizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLTokenizer] class.
//
// # Creating a tokenizer
//
//   - [INLTokenizer.InitWithUnit]: Creates a tokenizer with the specified unit.
//
// # Configuring a tokenizer
//
//   - [INLTokenizer.String]: The text to be tokenized.
//   - [INLTokenizer.SetString]
//   - [INLTokenizer.SetLanguage]: Sets the language of the text to be tokenized.
//   - [INLTokenizer.Unit]: The linguistic unit that this tokenizer uses.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer
type INLTokenizer interface {
	objectivec.IObject

	// Topic: Creating a tokenizer

	// Creates a tokenizer with the specified unit.
	InitWithUnit(unit NLTokenUnit) NLTokenizer

	// Topic: Configuring a tokenizer

	// The text to be tokenized.
	String() string
	SetString(value string)
	// Sets the language of the text to be tokenized.
	SetLanguage(language NLLanguage)
	// The linguistic unit that this tokenizer uses.
	Unit() NLTokenUnit

	// Enumerates over a given range of the string and calls the specified block for each token.
	EnumerateTokensInRangeUsingBlock(range_ foundation.NSRange, block func(unsafe.Pointer, uint64, *bool))
	// Finds the range of the token at the given index.
	TokenRangeAtIndex(characterIndex uint) foundation.NSRange
	// Finds the entire range of all tokens contained completely or partially within the specified range.
	TokenRangeForRange(range_ foundation.NSRange) foundation.NSRange
	// Tokenizes the string within the provided range.
	TokensForRange(range_ foundation.NSRange) []foundation.NSValue
}

// Init initializes the instance.
func (t NLTokenizer) Init() NLTokenizer {
	rv := objc.Send[NLTokenizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NLTokenizer) Autorelease() NLTokenizer {
	rv := objc.Send[NLTokenizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLTokenizer creates a new NLTokenizer instance.
func NewNLTokenizer() NLTokenizer {
	class := getNLTokenizerClass()
	rv := objc.Send[NLTokenizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a tokenizer with the specified unit.
//
// unit: The level of granularity that the tokenizer operates at.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/init(unit:)
func NewTokenizerWithUnit(unit NLTokenUnit) NLTokenizer {
	instance := getNLTokenizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnit:"), unit)
	return NLTokenizerFromID(rv)
}

// Creates a tokenizer with the specified unit.
//
// unit: The level of granularity that the tokenizer operates at.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/init(unit:)
func (t NLTokenizer) InitWithUnit(unit NLTokenUnit) NLTokenizer {
	rv := objc.Send[NLTokenizer](t.ID, objc.Sel("initWithUnit:"), unit)
	return rv
}

// Sets the language of the text to be tokenized.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/setLanguage(_:)
func (t NLTokenizer) SetLanguage(language NLLanguage) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLanguage:"), objc.String(string(language)))
}

// Enumerates over a given range of the string and calls the specified block
// for each token.
//
// range: The range of the string to tokenize.
//
// block: The closure to call after each token; return false if processing should
// stop.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/enumerateTokensInRange:usingBlock:
func (t NLTokenizer) EnumerateTokensInRangeUsingBlock(range_ foundation.NSRange, block func(unsafe.Pointer, uint64, *bool)) {
	_block1 := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer, arg1 uint64, arg2 *bool) { block(arg0, arg1, arg2) })
	defer _block1.Release()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateTokensInRange:usingBlock:"), range_, objc.ID(_block1))
}

// Finds the range of the token at the given index.
//
// # Return Value
//
// The range of the token at the given location.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeAtIndex:
func (t NLTokenizer) TokenRangeAtIndex(characterIndex uint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("tokenRangeAtIndex:"), characterIndex)
	return foundation.NSRange(rv)
}

// Finds the entire range of all tokens contained completely or partially
// within the specified range.
//
// range: The range within the string to search for tokens.
//
// # Return Value
//
// The smallest possible range that contains all of the tokens within the
// range specified in `range`. This result includes a token’s entire range
// if any part of that token is included within `range`. If the length of
// `range` is 0, this return value is equivalent to [TokenRangeAtIndex].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeForRange:
func (t NLTokenizer) TokenRangeForRange(range_ foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("tokenRangeForRange:"), range_)
	return foundation.NSRange(rv)
}

// Tokenizes the string within the provided range.
//
// range: The range within the string that should be tokenzied.
//
// # Return Value
//
// Returns the ranges corresponding to the tokens for the tokenizer’s unit
// that intersect the given range.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokensForRange:
func (t NLTokenizer) TokensForRange(range_ foundation.NSRange) []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tokensForRange:"), range_)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

// The text to be tokenized.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/string
func (t NLTokenizer) String() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
func (t NLTokenizer) SetString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setString:"), objc.String(value))
}

// The linguistic unit that this tokenizer uses.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/unit
func (t NLTokenizer) Unit() NLTokenUnit {
	rv := objc.Send[NLTokenUnit](t.ID, objc.Sel("unit"))
	return NLTokenUnit(rv)
}
