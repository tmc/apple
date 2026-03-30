// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLinguisticTagger] class.
var (
	_NSLinguisticTaggerClass     NSLinguisticTaggerClass
	_NSLinguisticTaggerClassOnce sync.Once
)

func getNSLinguisticTaggerClass() NSLinguisticTaggerClass {
	_NSLinguisticTaggerClassOnce.Do(func() {
		_NSLinguisticTaggerClass = NSLinguisticTaggerClass{class: objc.GetClass("NSLinguisticTagger")}
	})
	return _NSLinguisticTaggerClass
}

// GetNSLinguisticTaggerClass returns the class object for NSLinguisticTagger.
func GetNSLinguisticTaggerClass() NSLinguisticTaggerClass {
	return getNSLinguisticTaggerClass()
}

type NSLinguisticTaggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLinguisticTaggerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLinguisticTaggerClass) Alloc() NSLinguisticTagger {
	rv := objc.Send[NSLinguisticTagger](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Analyze natural language text to tag part of speech and lexical class,
// identify names, perform lemmatization, and determine the language and
// script.
//
// # Overview
//
// [NSLinguisticTagger] provides a uniform interface to a variety of natural
// language processing functionality with support for many different languages
// and scripts. You can use this class to segment natural language text into
// paragraphs, sentences, or words, and tag information about those segments,
// such as part of speech, lexical class, lemma, script, and language.
//
// When you create a linguistic tagger, you specify what kind of information
// you’re interested in by passing one or more [NSLinguisticTagScheme]
// values. Set the [String] property to the natural language text you want to
// analyze, and the linguistic tagger processes it according to the specified
// tag schemes. You can then enumerate over the tags in a specified range,
// using the methods described in Enumerating Linguistic Tags, to get the
// information requested for a given scheme and unit.
//
// # Thread Safety
//
// A single instance of [NSLinguisticTagger] should not be used simultaneously
// from multiple threads.
//
// # First Steps
//
//   - [NSLinguisticTagger.String]: The string being analyzed by the linguistic tagger.
//   - [NSLinguisticTagger.SetString]
//
// # Getting the Tag Schemes
//
//   - [NSLinguisticTagger.TagSchemes]: Returns the tag schemes configured for this linguistic tagger. For possible values, see [NSLinguisticTagScheme](<doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagScheme>).
//
// # Determining the Dominant Language and Orthography
//
//   - [NSLinguisticTagger.DominantLanguage]: Returns the dominant language of the string set for the linguistic tagger.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger
type NSLinguisticTagger struct {
	objectivec.Object
}

// NSLinguisticTaggerFromID constructs a [NSLinguisticTagger] from an objc.ID.
//
// Analyze natural language text to tag part of speech and lexical class,
// identify names, perform lemmatization, and determine the language and
// script.
func NSLinguisticTaggerFromID(id objc.ID) NSLinguisticTagger {
	return NSLinguisticTagger{objectivec.Object{ID: id}}
}

// NOTE: NSLinguisticTagger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLinguisticTagger] class.
//
// # First Steps
//
//   - [INSLinguisticTagger.String]: The string being analyzed by the linguistic tagger.
//   - [INSLinguisticTagger.SetString]
//
// # Getting the Tag Schemes
//
//   - [INSLinguisticTagger.TagSchemes]: Returns the tag schemes configured for this linguistic tagger. For possible values, see [NSLinguisticTagScheme](<doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagScheme>).
//
// # Determining the Dominant Language and Orthography
//
//   - [INSLinguisticTagger.DominantLanguage]: Returns the dominant language of the string set for the linguistic tagger.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger
type INSLinguisticTagger interface {
	objectivec.IObject

	// Topic: First Steps

	// The string being analyzed by the linguistic tagger.
	String() string
	SetString(value string)

	// Topic: Getting the Tag Schemes

	// Returns the tag schemes configured for this linguistic tagger. For possible values, see [NSLinguisticTagScheme](<doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagScheme>).
	TagSchemes() []string

	// Topic: Determining the Dominant Language and Orthography

	// Returns the dominant language of the string set for the linguistic tagger.
	DominantLanguage() string
}

// Init initializes the instance.
func (l NSLinguisticTagger) Init() NSLinguisticTagger {
	rv := objc.Send[NSLinguisticTagger](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLinguisticTagger) Autorelease() NSLinguisticTagger {
	rv := objc.Send[NSLinguisticTagger](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLinguisticTagger creates a new NSLinguisticTagger instance.
func NewNSLinguisticTagger() NSLinguisticTagger {
	class := getNSLinguisticTaggerClass()
	rv := objc.Send[NSLinguisticTagger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a linguistic tagger instance using the specified tag schemes and
// options.
//
// tagSchemes: An array of tag schemes to be used. See [NSLinguisticTagScheme] for the
// possible values.
//
// opts: Reserved for future use. Specify `0` for this parameter.
//
// # Return Value
//
// An initialized linguistic tagger.
//
// # Discussion
//
// Pass any tag schemes to `tagSchemes` that you intend to use with the
// methods described in Enumerating Linguistic Tags and Getting Linguistic
// Tags.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/init(tagSchemes:options:)
func NewLinguisticTaggerWithTagSchemesOptions(tagSchemes []string, opts uint) NSLinguisticTagger {
	instance := getNSLinguisticTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTagSchemes:options:"), objectivec.StringSliceToNSArray(tagSchemes), opts)
	return NSLinguisticTaggerFromID(rv)
}

// The string being analyzed by the linguistic tagger.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/string
func (l NSLinguisticTagger) String() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("string"))
	return NSStringFromID(rv).String()
}
func (l NSLinguisticTagger) SetString(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setString:"), objc.String(value))
}

// Returns the tag schemes configured for this linguistic tagger. For possible
// values, see [NSLinguisticTagScheme].
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/tagSchemes
func (l NSLinguisticTagger) TagSchemes() []string {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("tagSchemes"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the dominant language of the string set for the linguistic tagger.
//
// # Return Value
//
// The BCP-47 tag identifying the dominant language of the string, or the tag
// “und” if a specific language cannot be determined.
//
// # Discussion
//
// If you want to know the dominant language of a string that you’re
// analyzing with a linguistic tagger (for example, identifying part of speech
// for each word), specify the [language] tag scheme in the initializer. After
// you set the [String] property of the linguistic tagger, the dominant
// language can be determined with the [DominantLanguage] property, as shown
// in this example:
//
// In the example, the BCP-47 language tag “de” is returned as the
// dominant language, indicating that the text is in German.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/dominantLanguage
//
// [language]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagScheme/language
func (l NSLinguisticTagger) DominantLanguage() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("dominantLanguage"))
	return NSStringFromID(rv).String()
}
