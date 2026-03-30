// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLLanguageRecognizer] class.
var (
	_NLLanguageRecognizerClass     NLLanguageRecognizerClass
	_NLLanguageRecognizerClassOnce sync.Once
)

func getNLLanguageRecognizerClass() NLLanguageRecognizerClass {
	_NLLanguageRecognizerClassOnce.Do(func() {
		_NLLanguageRecognizerClass = NLLanguageRecognizerClass{class: objc.GetClass("NLLanguageRecognizer")}
	})
	return _NLLanguageRecognizerClass
}

// GetNLLanguageRecognizerClass returns the class object for NLLanguageRecognizer.
func GetNLLanguageRecognizerClass() NLLanguageRecognizerClass {
	return getNLLanguageRecognizerClass()
}

type NLLanguageRecognizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLLanguageRecognizerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLLanguageRecognizerClass) Alloc() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The language of a body of text.
//
// # Overview
//
// An [NLLanguageRecognizer] object automatically detects the language of a
// piece of text. It performs language identification by:
//
// - Identifying the dominant script of a piece of text. Some languages have a
// unique script (like Greek), but others share the same script (like English,
// French, and German, which all share the Latin script). - Identifying the
// language itself.
//
// The identification obtained from an [NLLanguageRecognizer] object can be
// either a single most likely language, access through [NLLanguageRecognizer.DominantLanguage], or
// a set of language candidates with probabilities, using
// [NLLanguageRecognizer.LanguageHypothesesWithMaximum]. You can reset the recognizer to its
// initial state, to be reused for new analysis.
//
// Use the convenience method, [NLLanguageRecognizer.DominantLanguageForString], to get the most
// likely language without creating an [NLLanguageRecognizer].
//
// # Determining the language
//
//   - [NLLanguageRecognizer.ProcessString]: Analyzes the piece of text to determine its dominant language.
//   - [NLLanguageRecognizer.DominantLanguage]: The most likely language for the processed text.
//   - [NLLanguageRecognizer.Reset]: Resets the recognizer to its initial state.
//
// # Guiding the recognizer
//
//   - [NLLanguageRecognizer.LanguageConstraints]: Limits the set of possible languages that the recognizer will return.
//   - [NLLanguageRecognizer.SetLanguageConstraints]
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer
type NLLanguageRecognizer struct {
	objectivec.Object
}

// NLLanguageRecognizerFromID constructs a [NLLanguageRecognizer] from an objc.ID.
//
// The language of a body of text.
func NLLanguageRecognizerFromID(id objc.ID) NLLanguageRecognizer {
	return NLLanguageRecognizer{objectivec.Object{ID: id}}
}

// NOTE: NLLanguageRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLLanguageRecognizer] class.
//
// # Determining the language
//
//   - [INLLanguageRecognizer.ProcessString]: Analyzes the piece of text to determine its dominant language.
//   - [INLLanguageRecognizer.DominantLanguage]: The most likely language for the processed text.
//   - [INLLanguageRecognizer.Reset]: Resets the recognizer to its initial state.
//
// # Guiding the recognizer
//
//   - [INLLanguageRecognizer.LanguageConstraints]: Limits the set of possible languages that the recognizer will return.
//   - [INLLanguageRecognizer.SetLanguageConstraints]
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer
type INLLanguageRecognizer interface {
	objectivec.IObject

	// Topic: Determining the language

	// Analyzes the piece of text to determine its dominant language.
	ProcessString(string_ string)
	// The most likely language for the processed text.
	DominantLanguage() NLLanguage
	// Resets the recognizer to its initial state.
	Reset()

	// Topic: Guiding the recognizer

	// Limits the set of possible languages that the recognizer will return.
	LanguageConstraints() []string
	SetLanguageConstraints(value []string)

	// A dictionary that maps languages to their probabilities in the language identification process.
	LanguageHints() foundation.INSDictionary
	SetLanguageHints(value foundation.INSDictionary)
	// Generates the probabilities of possible languages for the processed text.
	LanguageHypothesesWithMaximum(maxHypotheses uint) foundation.INSDictionary
}

// Init initializes the instance.
func (l NLLanguageRecognizer) Init() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NLLanguageRecognizer) Autorelease() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLLanguageRecognizer creates a new NLLanguageRecognizer instance.
func NewNLLanguageRecognizer() NLLanguageRecognizer {
	class := getNLLanguageRecognizerClass()
	rv := objc.Send[NLLanguageRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Analyzes the piece of text to determine its dominant language.
//
// # Discussion
//
// Use this method to process the provided text and to update the
// [DominantLanguage] result and `languageHypotheses()` probabilities.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/processString(_:)
func (l NLLanguageRecognizer) ProcessString(string_ string) {
	objc.Send[objc.ID](l.ID, objc.Sel("processString:"), objc.String(string_))
}

// Resets the recognizer to its initial state.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/reset()
func (l NLLanguageRecognizer) Reset() {
	objc.Send[objc.ID](l.ID, objc.Sel("reset"))
}

// Generates the probabilities of possible languages for the processed text.
//
// maxHypotheses: The maximum number of languages to return.
//
// # Return Value
//
// A dictionary mapping languages with their probabilities, up to
// `maxHypotheses` languages.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHypothesesWithMaximum:
func (l NLLanguageRecognizer) LanguageHypothesesWithMaximum(maxHypotheses uint) foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("languageHypothesesWithMaximum:"), maxHypotheses)
	return foundation.NSDictionaryFromID(rv)
}

// Finds the most likely language of a piece of text.
//
// string: The text to analyze.
//
// # Return Value
//
// The most probable language of the piece of text.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/dominantLanguage(for:)
func (_NLLanguageRecognizerClass NLLanguageRecognizerClass) DominantLanguageForString(string_ string) NLLanguage {
	rv := objc.Send[objc.ID](objc.ID(_NLLanguageRecognizerClass.class), objc.Sel("dominantLanguageForString:"), objc.String(string_))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}

// The most likely language for the processed text.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/dominantLanguage
func (l NLLanguageRecognizer) DominantLanguage() NLLanguage {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("dominantLanguage"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}

// Limits the set of possible languages that the recognizer will return.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageConstraints
func (l NLLanguageRecognizer) LanguageConstraints() []string {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("languageConstraints"))
	return objc.ConvertSliceToStrings(rv)
}
func (l NLLanguageRecognizer) SetLanguageConstraints(value []string) {
	objc.Send[struct{}](l.ID, objc.Sel("setLanguageConstraints:"), objectivec.StringSliceToNSArray(value))
}

// A dictionary that maps languages to their probabilities in the language
// identification process.
//
// # Discussion
//
// This is a dictionary mapping languages to their probabilities and used by
// [ProcessString].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHints-3gy00
func (l NLLanguageRecognizer) LanguageHints() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("languageHints"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (l NLLanguageRecognizer) SetLanguageHints(value foundation.INSDictionary) {
	objc.Send[struct{}](l.ID, objc.Sel("setLanguageHints:"), value)
}
