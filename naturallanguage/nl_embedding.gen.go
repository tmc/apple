// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLEmbedding] class.
var (
	_NLEmbeddingClass     NLEmbeddingClass
	_NLEmbeddingClassOnce sync.Once
)

func getNLEmbeddingClass() NLEmbeddingClass {
	_NLEmbeddingClassOnce.Do(func() {
		_NLEmbeddingClass = NLEmbeddingClass{class: objc.GetClass("NLEmbedding")}
	})
	return _NLEmbeddingClass
}

// GetNLEmbeddingClass returns the class object for NLEmbedding.
func GetNLEmbeddingClass() NLEmbeddingClass {
	return getNLEmbeddingClass()
}

type NLEmbeddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLEmbeddingClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLEmbeddingClass) Alloc() NLEmbedding {
	rv := objc.Send[NLEmbedding](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A map of strings to vectors, which locates neighboring, similar strings.
//
// # Overview
// 
// Use an [NLEmbedding] to find similar strings based on the proximity of
// their vectors. The is the entire set of strings in an embedding. Each
// string in the vocabulary has a vector, which is an array of doubles, and
// each double corresponds to a dimension in the embedding. An [NLEmbedding]
// uses these vectors to determine the distance between two strings, or to
// find the nearest neighbors of a string in the vocabulary. The higher the
// similarity of any two strings, the smaller the distance is between them.
// 
// [Natural Language] provides built-in word embeddings that you can retrieve
// by using the [NLEmbedding.WordEmbeddingForLanguage] method. You can also compile your
// own custom embedding into an efficient, searchable, on-disk representation.
// Typically, you compile an embedding by using Create ML’s
// [MLWordEmbedding] and save it as a file for your Xcode project at
// development time. Alternatively, you can compile an embedding at runtime by
// using Natural Language’s
// [NLEmbedding.WriteEmbeddingForDictionaryLanguageRevisionToURLError] method.
// 
// Your custom embedding can use any kind of string that’s useful to your
// app, such as phrases, brand names, serial numbers, and so on. For example,
// you could make an embedding of movie titles. Each movie title could have a
// vector that places similar movies close together in the embedding.
//
// [MLWordEmbedding]: https://developer.apple.com/documentation/CreateML/MLWordEmbedding
// [Natural Language]: https://developer.apple.com/documentation/NaturalLanguage
//
// # Inspecting the vocabulary of an embedding
//
//   - [NLEmbedding.Dimension]: The number of dimensions in the vocabulary’s vector space.
//   - [NLEmbedding.VocabularySize]: The number of words in the vocabulary.
//   - [NLEmbedding.Language]: The language of the text in the word embedding.
//   - [NLEmbedding.ContainsString]: Requests a Boolean value that indicates whether the term is in the vocabulary.
//   - [NLEmbedding.Revision]: The revision of the word embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding
type NLEmbedding struct {
	objectivec.Object
}

// NLEmbeddingFromID constructs a [NLEmbedding] from an objc.ID.
//
// A map of strings to vectors, which locates neighboring, similar strings.
func NLEmbeddingFromID(id objc.ID) NLEmbedding {
	return NLEmbedding{objectivec.Object{ID: id}}
}
// NOTE: NLEmbedding adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLEmbedding] class.
//
// # Inspecting the vocabulary of an embedding
//
//   - [INLEmbedding.Dimension]: The number of dimensions in the vocabulary’s vector space.
//   - [INLEmbedding.VocabularySize]: The number of words in the vocabulary.
//   - [INLEmbedding.Language]: The language of the text in the word embedding.
//   - [INLEmbedding.ContainsString]: Requests a Boolean value that indicates whether the term is in the vocabulary.
//   - [INLEmbedding.Revision]: The revision of the word embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding
type INLEmbedding interface {
	objectivec.IObject

	// Topic: Inspecting the vocabulary of an embedding

	// The number of dimensions in the vocabulary’s vector space.
	Dimension() uint
	// The number of words in the vocabulary.
	VocabularySize() uint
	// The language of the text in the word embedding.
	Language() NLLanguage
	// Requests a Boolean value that indicates whether the term is in the vocabulary.
	ContainsString(string_ string) bool
	// The revision of the word embedding.
	Revision() uint

	// Calculates the distance between two strings in the vocabulary space.
	DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType NLDistanceType) NLDistance
	// Copies a vector into the given a pointer to a float array.
	GetVectorForString(string_ string) (float32, bool)
	// Retrieves a limited number of strings near a string in the vocabulary.
	NeighborsForStringMaximumCountDistanceType(string_ string, maxCount uint, distanceType NLDistanceType) []string
	// Retrieves a limited number of strings, within a radius of a string, in the vocabulary.
	NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ string, maxCount uint, maxDistance NLDistance, distanceType NLDistanceType) []string
	// Retrieves a limited number of strings near a location in the vocabulary space.
	NeighborsForVectorMaximumCountDistanceType(vector []foundation.NSNumber, maxCount uint, distanceType NLDistanceType) []string
	// Retrieves a limited number of strings within a radius of a location in the vocabulary space.
	NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector []foundation.NSNumber, maxCount uint, maxDistance NLDistance, distanceType NLDistanceType) []string
	// Requests the vector for the given term.
	VectorForString(string_ string) []foundation.NSNumber
}

// Init initializes the instance.
func (e NLEmbedding) Init() NLEmbedding {
	rv := objc.Send[NLEmbedding](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NLEmbedding) Autorelease() NLEmbedding {
	rv := objc.Send[NLEmbedding](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLEmbedding creates a new NLEmbedding instance.
func NewNLEmbedding() NLEmbedding {
	class := getNLEmbeddingClass()
	rv := objc.Send[NLEmbedding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a word embedding from a model file.
//
// url: The location of the .`mlmodel` file that contains a word embedding.
//
// # Discussion
// 
// Use this initializer to create a word embedding from an
// `XCUIElementTypeMlmodel` file saved by Create ML’s [MLWordEmbedding].
//
// [MLWordEmbedding]: https://developer.apple.com/documentation/CreateML/MLWordEmbedding
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func NewEmbeddingWithContentsOfURLError(url foundation.INSURL) (NLEmbedding, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getNLEmbeddingClass().class), objc.Sel("embeddingWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLEmbeddingFromID(rv), nil
}

// Requests a Boolean value that indicates whether the term is in the
// vocabulary.
//
// string: The term to search for in the word embedding.
//
// # Return Value
// 
// `true` if the term is in the word embedding’s vocabulary, otherwise
// `false`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/contains(_:)
func (e NLEmbedding) ContainsString(string_ string) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("containsString:"), objc.String(string_))
	return rv
}
// Calculates the distance between two strings in the vocabulary space.
//
// firstString: A string in the embedding vocabulary.
//
// secondString: Another string in the embedding vocabulary.
//
// distanceType: A means of calculating distance that determines which formula the method
// uses to evaluate the distance between `firstString` and `secondString`.
//
// # Return Value
// 
// The distance associated with `distanceType`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/distanceBetweenString:andString:distanceType:
func (e NLEmbedding) DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType NLDistanceType) NLDistance {
	rv := objc.Send[NLDistance](e.ID, objc.Sel("distanceBetweenString:andString:distanceType:"), objc.String(firstString), objc.String(secondString), distanceType)
	return NLDistance(rv)
}
// Copies a vector into the given a pointer to a float array.
//
// vector: An array of floats the method copies the vector to. The array’s capacity
// must be at least [Dimension].
//
// string: The term to find in the word embedding.
//
// # Return Value
// 
// A Boolean indicating whether the method copied the vector.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/getVector:forString:
func (e NLEmbedding) GetVectorForString(string_ string) (float32, bool) {
	var vector float32
	rv := objc.Send[bool](e.ID, objc.Sel("getVector:forString:"), unsafe.Pointer(&vector), objc.String(string_))
	return vector, rv
}
// Retrieves a limited number of strings near a string in the vocabulary.
//
// string: A string in the embedding vocabulary.
//
// maxCount: The largest number of neighboring strings that the method can return in an
// array.
//
// distanceType: A means of calculating distance that determines which formula the method
// uses to evaluate a neighbor’s distance from `string`.
//
// # Return Value
// 
// An array of neighboring strings.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:distanceType:
func (e NLEmbedding) NeighborsForStringMaximumCountDistanceType(string_ string, maxCount uint, distanceType NLDistanceType) []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("neighborsForString:maximumCount:distanceType:"), objc.String(string_), maxCount, distanceType)
	return objc.ConvertSliceToStrings(rv)
}
// Retrieves a limited number of strings, within a radius of a string, in the
// vocabulary.
//
// string: A string in the embedding vocabulary.
//
// maxCount: The largest number of neighboring strings that the method can return in an
// array.
//
// maxDistance: The largest distance a neighbor can be from `string`.
//
// distanceType: A means of calculating distance that determines which formula the method
// uses to evaluate a neighbor’s distance from `string`.
//
// # Return Value
// 
// An array of neighboring strings.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:maximumDistance:distanceType:
func (e NLEmbedding) NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ string, maxCount uint, maxDistance NLDistance, distanceType NLDistanceType) []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("neighborsForString:maximumCount:maximumDistance:distanceType:"), objc.String(string_), maxCount, maxDistance, distanceType)
	return objc.ConvertSliceToStrings(rv)
}
// Retrieves a limited number of strings near a location in the vocabulary
// space.
//
// vector: A location in the vocabulary space.
//
// maxCount: The largest number of neighboring strings that the method can return in an
// array.
//
// distanceType: A means of calculating distance that determines which formula the method
// uses to evaluate a neighbor’s distance from `vector`.
//
// # Return Value
// 
// An array of tuples that contain neighboring strings and their distances. In
// Objective-C, this returns an array of neighboring strings.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:distanceType:
func (e NLEmbedding) NeighborsForVectorMaximumCountDistanceType(vector []foundation.NSNumber, maxCount uint, distanceType NLDistanceType) []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("neighborsForVector:maximumCount:distanceType:"), objectivec.IObjectSliceToNSArray(vector), maxCount, distanceType)
	return objc.ConvertSliceToStrings(rv)
}
// Retrieves a limited number of strings within a radius of a location in the
// vocabulary space.
//
// vector: A location in the vocabulary space.
//
// maxCount: The largest number of neighboring strings that the method can return in an
// array.
//
// maxDistance: The largest distance a neighbor can be from `vector`.
//
// distanceType: A means of calculating distance that determines which formula the method
// uses to evaluate a neighbor’s distance from `vector`.
//
// # Return Value
// 
// An array of strings.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:maximumDistance:distanceType:
func (e NLEmbedding) NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector []foundation.NSNumber, maxCount uint, maxDistance NLDistance, distanceType NLDistanceType) []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("neighborsForVector:maximumCount:maximumDistance:distanceType:"), objectivec.IObjectSliceToNSArray(vector), maxCount, maxDistance, distanceType)
	return objc.ConvertSliceToStrings(rv)
}
// Requests the vector for the given term.
//
// string: The term to find in the word embedding.
//
// # Return Value
// 
// A vector represented as an array of doubles if present in the word
// embedding, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vectorForString:
func (e NLEmbedding) VectorForString(string_ string) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("vectorForString:"), objc.String(string_))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Retrieves a word embedding for a given language.
//
// language: The language of the word embedding, such as [french].
// //
// [french]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguage/french
//
// # Return Value
// 
// An [NLEmbedding] if available, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:)
func (_NLEmbeddingClass NLEmbeddingClass) WordEmbeddingForLanguage(language NLLanguage) NLEmbedding {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("wordEmbeddingForLanguage:"), objc.String(string(language)))
	return NLEmbeddingFromID(rv)
}
// Retrieves a word embedding for a given language and revision.
//
// language: The language of the word embedding, such as [french].
// //
// [french]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguage/french
//
// revision: The revision of the word embedding.
//
// # Return Value
// 
// An [NLEmbedding] if available, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:revision:)
func (_NLEmbeddingClass NLEmbeddingClass) WordEmbeddingForLanguageRevision(language NLLanguage, revision uint) NLEmbedding {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("wordEmbeddingForLanguage:revision:"), objc.String(string(language)), revision)
	return NLEmbeddingFromID(rv)
}
// Retrieves a sentence embedding for a given language.
//
// language: The language of the sentence embedding, such as [french]. For possible
// values, see [NLLanguage].
// //
// [french]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguage/french
//
// # Return Value
// 
// An [NLEmbedding] if available, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:)
func (_NLEmbeddingClass NLEmbeddingClass) SentenceEmbeddingForLanguage(language NLLanguage) NLEmbedding {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("sentenceEmbeddingForLanguage:"), objc.String(string(language)))
	return NLEmbeddingFromID(rv)
}
// Retrieves a sentence embedding for a given language and revision.
//
// language: The language of the sentence embedding, such as [french]. For possible
// values, see [NLLanguage].
// //
// [french]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguage/french
//
// revision: The revision of the sentence embedding.
//
// # Return Value
// 
// An [NLEmbedding] if available, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:revision:)
func (_NLEmbeddingClass NLEmbeddingClass) SentenceEmbeddingForLanguageRevision(language NLLanguage, revision uint) NLEmbedding {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("sentenceEmbeddingForLanguage:revision:"), objc.String(string(language)), revision)
	return NLEmbeddingFromID(rv)
}
// Retrieves the current version of a word embedding for the given language.
//
// language: A language supported by the Natural Language framework.
//
// # Return Value
// 
// An integer.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentRevision(for:)
func (_NLEmbeddingClass NLEmbeddingClass) CurrentRevisionForLanguage(language NLLanguage) uint {
	rv := objc.Send[uint](objc.ID(_NLEmbeddingClass.class), objc.Sel("currentRevisionForLanguage:"), objc.String(string(language)))
	return rv
}
// Retrieves all version numbers of a word embedding for the given language.
//
// language: A language supported by the Natural Language framework.
//
// # Return Value
// 
// An index set.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedRevisions(for:)
func (_NLEmbeddingClass NLEmbeddingClass) SupportedRevisionsForLanguage(language NLLanguage) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("supportedRevisionsForLanguage:"), objc.String(string(language)))
	return foundation.NSIndexSetFromID(rv)
}
// Retrieves the current version of a sentence embedding for the given
// language.
//
// language: A language supported by the Natural Language framework. For possible
// values, see [NLLanguage].
//
// # Return Value
// 
// An integer representing the current version number of a sentence embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentSentenceEmbeddingRevision(for:)
func (_NLEmbeddingClass NLEmbeddingClass) CurrentSentenceEmbeddingRevisionForLanguage(language NLLanguage) uint {
	rv := objc.Send[uint](objc.ID(_NLEmbeddingClass.class), objc.Sel("currentSentenceEmbeddingRevisionForLanguage:"), objc.String(string(language)))
	return rv
}
// Retrieves all version numbers of a sentence embedding for the given
// language.
//
// language: A language supported by the Natural Language framework. For possible
// values, see [NLLanguage].
//
// # Return Value
// 
// An index set representing all of the supported version numbers of the
// sentence embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedSentenceEmbeddingRevisions(for:)
func (_NLEmbeddingClass NLEmbeddingClass) SupportedSentenceEmbeddingRevisionsForLanguage(language NLLanguage) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NLEmbeddingClass.class), objc.Sel("supportedSentenceEmbeddingRevisionsForLanguage:"), objc.String(string(language)))
	return foundation.NSIndexSetFromID(rv)
}
// Exports the word embedding contained within a Core ML model file at the
// given URL.
//
// dictionary: A dictionary of terms, and their vectors, which are represented by an array
// of doubles.
//
// language: The language of the text in the word embedding.
//
// revision: The revision of the word embedding.
//
// url: The location in the file system to write the file to.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/writeEmbeddingForDictionary:language:revision:toURL:error:
func (_NLEmbeddingClass NLEmbeddingClass) WriteEmbeddingForDictionaryLanguageRevisionToURLError(dictionary foundation.INSDictionary, language NLLanguage, revision uint, url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NLEmbeddingClass.class), objc.Sel("writeEmbeddingForDictionary:language:revision:toURL:error:"), dictionary, objc.String(string(language)), revision, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeEmbeddingForDictionary:language:revision:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The number of dimensions in the vocabulary’s vector space.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/dimension
func (e NLEmbedding) Dimension() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("dimension"))
	return rv
}
// The number of words in the vocabulary.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vocabularySize
func (e NLEmbedding) VocabularySize() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("vocabularySize"))
	return rv
}
// The language of the text in the word embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/language
func (e NLEmbedding) Language() NLLanguage {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("language"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}
// The revision of the word embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/revision
func (e NLEmbedding) Revision() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("revision"))
	return rv
}

