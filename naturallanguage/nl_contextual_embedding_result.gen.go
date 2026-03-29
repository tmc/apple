// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLContextualEmbeddingResult] class.
var (
	_NLContextualEmbeddingResultClass     NLContextualEmbeddingResultClass
	_NLContextualEmbeddingResultClassOnce sync.Once
)

func getNLContextualEmbeddingResultClass() NLContextualEmbeddingResultClass {
	_NLContextualEmbeddingResultClassOnce.Do(func() {
		_NLContextualEmbeddingResultClass = NLContextualEmbeddingResultClass{class: objc.GetClass("NLContextualEmbeddingResult")}
	})
	return _NLContextualEmbeddingResultClass
}

// GetNLContextualEmbeddingResultClass returns the class object for NLContextualEmbeddingResult.
func GetNLContextualEmbeddingResultClass() NLContextualEmbeddingResultClass {
	return getNLContextualEmbeddingResultClass()
}

type NLContextualEmbeddingResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLContextualEmbeddingResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLContextualEmbeddingResultClass) Alloc() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the embedding vector result from applying a
// contextual embedding to a string.
//
// # Inspecting the result
//
//   - [NLContextualEmbeddingResult.Language]: The resulting language.
//   - [NLContextualEmbeddingResult.SequenceLength]: The number of embedding vectors the request generates.
//   - [NLContextualEmbeddingResult.String]: The string value.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult
type NLContextualEmbeddingResult struct {
	objectivec.Object
}

// NLContextualEmbeddingResultFromID constructs a [NLContextualEmbeddingResult] from an objc.ID.
//
// An object that represents the embedding vector result from applying a
// contextual embedding to a string.
func NLContextualEmbeddingResultFromID(id objc.ID) NLContextualEmbeddingResult {
	return NLContextualEmbeddingResult{objectivec.Object{ID: id}}
}
// NOTE: NLContextualEmbeddingResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLContextualEmbeddingResult] class.
//
// # Inspecting the result
//
//   - [INLContextualEmbeddingResult.Language]: The resulting language.
//   - [INLContextualEmbeddingResult.SequenceLength]: The number of embedding vectors the request generates.
//   - [INLContextualEmbeddingResult.String]: The string value.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult
type INLContextualEmbeddingResult interface {
	objectivec.IObject

	// Topic: Inspecting the result

	// The resulting language.
	Language() NLLanguage
	// The number of embedding vectors the request generates.
	SequenceLength() uint
	// The string value.
	String() string

	// Returns a token vector at the specified character index.
	TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange foundation.NSRange) []foundation.NSNumber
}

// Init initializes the instance.
func (c NLContextualEmbeddingResult) Init() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NLContextualEmbeddingResult) Autorelease() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLContextualEmbeddingResult creates a new NLContextualEmbeddingResult instance.
func NewNLContextualEmbeddingResult() NLContextualEmbeddingResult {
	class := getNLContextualEmbeddingResultClass()
	rv := objc.Send[NLContextualEmbeddingResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a token vector at the specified character index.
//
// characterIndex: The index to get the token vector at.
//
// tokenRange: The character range of the token in the input string.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/tokenVectorAtIndex:tokenRange:
func (c NLContextualEmbeddingResult) TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange foundation.NSRange) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("tokenVectorAtIndex:tokenRange:"), characterIndex, tokenRange)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The resulting language.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/language
func (c NLContextualEmbeddingResult) Language() NLLanguage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("language"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}
// The number of embedding vectors the request generates.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/sequenceLength
func (c NLContextualEmbeddingResult) SequenceLength() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sequenceLength"))
	return rv
}
// The string value.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/string
func (c NLContextualEmbeddingResult) String() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

