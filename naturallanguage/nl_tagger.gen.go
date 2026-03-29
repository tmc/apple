// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLTagger] class.
var (
	_NLTaggerClass     NLTaggerClass
	_NLTaggerClassOnce sync.Once
)

func getNLTaggerClass() NLTaggerClass {
	_NLTaggerClassOnce.Do(func() {
		_NLTaggerClass = NLTaggerClass{class: objc.GetClass("NLTagger")}
	})
	return _NLTaggerClass
}

// GetNLTaggerClass returns the class object for NLTagger.
func GetNLTaggerClass() NLTaggerClass {
	return getNLTaggerClass()
}

type NLTaggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLTaggerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLTaggerClass) Alloc() NLTagger {
	rv := objc.Send[NLTagger](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A tagger that analyzes natural language text.
//
// # Overview
// 
// [NLTagger] supports many different languages and scripts. Use it to segment
// natural language text into paragraph, sentence, or word units and to tag
// each unit with information like part of speech, lexical class, lemma,
// script, and language.
// 
// When you create a linguistic tagger, you specify what kind of information
// you’re interested in by passing one or more [NLTagScheme] values. Set the
// [NLTagger.String] property to the natural language text you want to analyze, and the
// linguistic tagger processes it according to the specified tag schemes. You
// can then enumerate over the tags in a specified range, using the methods
// described in Enumerating linguistic tags, to get the information requested
// for a given scheme and unit.
//
// # Creating a tagger
//
//   - [NLTagger.InitWithTagSchemes]: Creates a linguistic tagger instance using the specified tag schemes and options.
//   - [NLTagger.String]: The string being analyzed by the linguistic tagger.
//   - [NLTagger.SetString]
//
// # Getting the tag schemes
//
//   - [NLTagger.TagSchemes]: The tag schemes configured for this linguistic tagger.
//
// # Determining the dominant language and orthography
//
//   - [NLTagger.DominantLanguage]: The dominant language of the string set for the linguistic tagger.
//
// # Using models with a tagger
//
//   - [NLTagger.SetModelsForTagScheme]: Assigns models for a tag scheme.
//   - [NLTagger.ModelsForTagScheme]: Returns the models that apply to the given tag scheme.
//
// # Using gazetteers with a tagger
//
//   - [NLTagger.SetGazetteersForTagScheme]: Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer.
//   - [NLTagger.GazetteersForTagScheme]: Retrieves the gazetteers attached to a tag scheme.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger
type NLTagger struct {
	objectivec.Object
}

// NLTaggerFromID constructs a [NLTagger] from an objc.ID.
//
// A tagger that analyzes natural language text.
func NLTaggerFromID(id objc.ID) NLTagger {
	return NLTagger{objectivec.Object{ID: id}}
}
// NOTE: NLTagger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLTagger] class.
//
// # Creating a tagger
//
//   - [INLTagger.InitWithTagSchemes]: Creates a linguistic tagger instance using the specified tag schemes and options.
//   - [INLTagger.String]: The string being analyzed by the linguistic tagger.
//   - [INLTagger.SetString]
//
// # Getting the tag schemes
//
//   - [INLTagger.TagSchemes]: The tag schemes configured for this linguistic tagger.
//
// # Determining the dominant language and orthography
//
//   - [INLTagger.DominantLanguage]: The dominant language of the string set for the linguistic tagger.
//
// # Using models with a tagger
//
//   - [INLTagger.SetModelsForTagScheme]: Assigns models for a tag scheme.
//   - [INLTagger.ModelsForTagScheme]: Returns the models that apply to the given tag scheme.
//
// # Using gazetteers with a tagger
//
//   - [INLTagger.SetGazetteersForTagScheme]: Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer.
//   - [INLTagger.GazetteersForTagScheme]: Retrieves the gazetteers attached to a tag scheme.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger
type INLTagger interface {
	objectivec.IObject

	// Topic: Creating a tagger

	// Creates a linguistic tagger instance using the specified tag schemes and options.
	InitWithTagSchemes(tagSchemes []string) NLTagger
	// The string being analyzed by the linguistic tagger.
	String() string
	SetString(value string)

	// Topic: Getting the tag schemes

	// The tag schemes configured for this linguistic tagger.
	TagSchemes() []string

	// Topic: Determining the dominant language and orthography

	// The dominant language of the string set for the linguistic tagger.
	DominantLanguage() NLLanguage

	// Topic: Using models with a tagger

	// Assigns models for a tag scheme.
	SetModelsForTagScheme(models []NLModel, tagScheme NLTagScheme)
	// Returns the models that apply to the given tag scheme.
	ModelsForTagScheme(tagScheme NLTagScheme) []NLModel

	// Topic: Using gazetteers with a tagger

	// Attaches gazetteers to a tag scheme, typically one gazetteer per language or one language-independent gazetteer.
	SetGazetteersForTagScheme(gazetteers []NLGazetteer, tagScheme NLTagScheme)
	// Retrieves the gazetteers attached to a tag scheme.
	GazetteersForTagScheme(tagScheme NLTagScheme) []NLGazetteer

	// Enumerates a block over the tagger’s string, given a range, token unit, and tag scheme.
	EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ foundation.NSRange, unit NLTokenUnit, scheme NLTagScheme, options NLTaggerOptions, block func(*string, unsafe.Pointer, *bool))
	// Sets the language for a range of text within the tagger’s string.
	SetLanguageRange(language NLLanguage, range_ foundation.NSRange)
	// Sets the orthography for the specified range.
	SetOrthographyRange(orthography foundation.NSOrthography, range_ foundation.NSRange)
	// Finds a tag for a given linguistic unit, for a single scheme, at the specified character position.
	TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit NLTokenUnit, scheme NLTagScheme, tokenRange foundation.NSRange) NLTag
	// Finds multiple possible tags for a given linguistic unit, for a single scheme, at the specified character position.
	TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit NLTokenUnit, scheme NLTagScheme, maximumCount uint, tokenRange foundation.NSRange) foundation.INSDictionary
	// Finds an array of linguistic tags and token ranges for a given string range and linguistic unit.
	TagsInRangeUnitSchemeOptionsTokenRanges(range_ foundation.NSRange, unit NLTokenUnit, scheme NLTagScheme, options NLTaggerOptions, tokenRanges []foundation.NSValue) []string
	// Returns the range of the linguistic unit containing the specified character index.
	TokenRangeAtIndexUnit(characterIndex uint, unit NLTokenUnit) foundation.NSRange
	// Finds the entire range of all tokens of the specified linguistic unit contained completely or partially within the specified range.
	TokenRangeForRangeUnit(range_ foundation.NSRange, unit NLTokenUnit) foundation.NSRange
}

// Init initializes the instance.
func (t NLTagger) Init() NLTagger {
	rv := objc.Send[NLTagger](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NLTagger) Autorelease() NLTagger {
	rv := objc.Send[NLTagger](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLTagger creates a new NLTagger instance.
func NewNLTagger() NLTagger {
	class := getNLTaggerClass()
	rv := objc.Send[NLTagger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a linguistic tagger instance using the specified tag schemes and
// options.
//
// tagSchemes: An array of tag schemes to be used. See [NLTagScheme] for the possible
// values.
//
// # Discussion
// 
// Pass any tag schemes to tagSchemes that you intend to use with the methods
// described in Enumerating linguistic tags and Getting linguistic tags in
// [NLTagger].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/init(tagSchemes:)
func NewTaggerWithTagSchemes(tagSchemes []string) NLTagger {
	instance := getNLTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTagSchemes:"), objectivec.StringSliceToNSArray(tagSchemes))
	return NLTaggerFromID(rv)
}

// Creates a linguistic tagger instance using the specified tag schemes and
// options.
//
// tagSchemes: An array of tag schemes to be used. See [NLTagScheme] for the possible
// values.
//
// # Discussion
// 
// Pass any tag schemes to tagSchemes that you intend to use with the methods
// described in Enumerating linguistic tags and Getting linguistic tags in
// [NLTagger].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/init(tagSchemes:)
func (t NLTagger) InitWithTagSchemes(tagSchemes []string) NLTagger {
	rv := objc.Send[NLTagger](t.ID, objc.Sel("initWithTagSchemes:"), objectivec.StringSliceToNSArray(tagSchemes))
	return rv
}
// Assigns models for a tag scheme.
//
// models: Array of [NLModel] objects to be used with this tagger.
//
// tagScheme: The tag scheme the models would be used with.
//
// # Discussion
// 
// An array of models is allowed for the case where multiple models need to be
// used. For example, when models were trained on specific languages.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setModels(_:forTagScheme:)
func (t NLTagger) SetModelsForTagScheme(models []NLModel, tagScheme NLTagScheme) {
	objc.Send[objc.ID](t.ID, objc.Sel("setModels:forTagScheme:"), objectivec.IObjectSliceToNSArray(models), objc.String(string(tagScheme)))
}
// Returns the models that apply to the given tag scheme.
//
// tagScheme: The tag scheme to filter the list of models with.
//
// # Return Value
// 
// The array of models that apply to the given tag scheme.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/models(forTagScheme:)
func (t NLTagger) ModelsForTagScheme(tagScheme NLTagScheme) []NLModel {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("modelsForTagScheme:"), objc.String(string(tagScheme)))
	return objc.ConvertSlice(rv, func(id objc.ID) NLModel {
		return NLModelFromID(id)
	})
}
// Attaches gazetteers to a tag scheme, typically one gazetteer per language
// or one language-independent gazetteer.
//
// gazetteers: The gazetteers to attach to a tag scheme.
//
// tagScheme: The tag scheme for the gazetteers.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setGazetteers(_:for:)
func (t NLTagger) SetGazetteersForTagScheme(gazetteers []NLGazetteer, tagScheme NLTagScheme) {
	objc.Send[objc.ID](t.ID, objc.Sel("setGazetteers:forTagScheme:"), objectivec.IObjectSliceToNSArray(gazetteers), objc.String(string(tagScheme)))
}
// Retrieves the gazetteers attached to a tag scheme.
//
// tagScheme: The tag scheme for the gazetteers.
//
// # Return Value
// 
// An array of [NLGazetteer].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/gazetteers(for:)
func (t NLTagger) GazetteersForTagScheme(tagScheme NLTagScheme) []NLGazetteer {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("gazetteersForTagScheme:"), objc.String(string(tagScheme)))
	return objc.ConvertSlice(rv, func(id objc.ID) NLGazetteer {
		return NLGazetteerFromID(id)
	})
}
// Enumerates a block over the tagger’s string, given a range, token unit,
// and tag scheme.
//
// range: The range of the string you want the tagger to analyze.
//
// unit: The linguistic unit of scale you’re interested in, such as
// [TokenUnitWord], [TokenUnitSentence], [TokenUnitParagraph], or
// [TokenUnitDocument].
//
// scheme: The tag scheme the tagger uses to tag the string, such as [language] or
// [script]. This scheme determines which types of [NLTag] the method passes
// to your block. For other tag schemes, see [NLTagScheme].
// //
// [language]: https://developer.apple.com/documentation/NaturalLanguage/NLTagScheme/language
// [script]: https://developer.apple.com/documentation/NaturalLanguage/NLTagScheme/script
//
// options: The set of linguistic tagger options to use, such as
// [TaggerOmitWhitespace]. For all available options, see [NLTagger.Options].
// //
// [NLTagger.Options]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options
//
// block: The block this method uses to iterate over the tagger’s [String]
// property. The block has the following parameters:
// 
// tag: The tag of the token. tokenRange: The range of the token. stop: A
// reference to a Boolean value. The block can set the value to `true` to stop
// further processing of the set. The `stop` argument is an out-only argument.
// You should only ever set this Boolean to `true` within the block.
//
// # Discussion
// 
// Use this method to iterate your block over the given range of a string. The
// method divides up the string with the given [NLTokenUnit] and [NLTagScheme]
// and then calls your block. For example, use the [lexicalClass] tag scheme
// to identify which tokens are parts of speech, types of whitespace, or types
// of punctuation. Use the [lemma] tag scheme to identify the stem form of
// each word token, if known.
//
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
// [lemma]: https://developer.apple.com/documentation/NaturalLanguage/NLTagScheme/lemma
// [lexicalClass]: https://developer.apple.com/documentation/NaturalLanguage/NLTagScheme/lexicalClass
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/enumerateTagsInRange:unit:scheme:options:usingBlock:
func (t NLTagger) EnumerateTagsInRangeUnitSchemeOptionsUsingBlock(range_ foundation.NSRange, unit NLTokenUnit, scheme NLTagScheme, options NLTaggerOptions, block func(*string, unsafe.Pointer, *bool)) {
	_block4 := objc.NewBlock(func(_ objc.Block, arg0 *string, arg1 unsafe.Pointer, arg2 *bool) { block(arg0, arg1, arg2) })
	defer _block4.Release()
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateTagsInRange:unit:scheme:options:usingBlock:"), range_, unit, scheme, options, objc.ID(_block4))
}
// Sets the language for a range of text within the tagger’s string.
//
// language: The language of the text range.
//
// range: The range of the string that is being assigned a language.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setLanguage:range:
func (t NLTagger) SetLanguageRange(language NLLanguage, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLanguage:range:"), objc.String(string(language)), range_)
}
// Sets the orthography for the specified range.
//
// orthography: The orthography for the given range.
//
// range: The range of the string that is being assigned an orthography.
//
// # Discussion
// 
// If the orthography of the linguistic tagger is not set, it will determine
// it automatically from the contents of the text. You should call this method
// only if you know the orthography of the text by some other means.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/setOrthography:range:
func (t NLTagger) SetOrthographyRange(orthography foundation.NSOrthography, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setOrthography:range:"), orthography, range_)
}
// Finds a tag for a given linguistic unit, for a single scheme, at the
// specified character position.
//
// characterIndex: The position of the initial character.
//
// unit: The linguistic unit. See [NLTokenUnit] for possible values.
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// scheme: The tag scheme. See [NLTagScheme] for possible values.
//
// tokenRange: A pointer to the token range.
//
// # Return Value
// 
// The tag for the requested tag scheme and linguistic unit, or `nil`. If a
// tag is returned, this function returns by reference the range of the token
// to `tokenRange`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagAtIndex:unit:scheme:tokenRange:
func (t NLTagger) TagAtIndexUnitSchemeTokenRange(characterIndex uint, unit NLTokenUnit, scheme NLTagScheme, tokenRange foundation.NSRange) NLTag {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tagAtIndex:unit:scheme:tokenRange:"), characterIndex, unit, objc.String(string(scheme)), tokenRange)
	return NLTag(foundation.NSStringFromID(rv).String())
}
// Finds multiple possible tags for a given linguistic unit, for a single
// scheme, at the specified character position.
//
// unit: The linguistic unit. See [NLTokenUnit] for possible values.
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// scheme: The tag scheme. See [NLTagScheme] for possible values. Not all tag schemes
// produce more than one prediction.
//
// maximumCount: The maximum number of tag predictions to return.
//
// tokenRange: The range of the token for which the tags were produced.
//
// # Return Value
// 
// A tuple containing a dictionary and a range.
//
// # Discussion
// 
// Each dictionary entry is a predicted tag with its associated probability
// score. These tags are the top candidates proposed as possible tags for the
// token. The dictionary contains up to `maximumCount` entries.
// 
// The range contains the range of the individual token for which these tags
// were produced.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:
func (t NLTagger) TagHypothesesAtIndexUnitSchemeMaximumCountTokenRange(characterIndex uint, unit NLTokenUnit, scheme NLTagScheme, maximumCount uint, tokenRange foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tagHypothesesAtIndex:unit:scheme:maximumCount:tokenRange:"), characterIndex, unit, objc.String(string(scheme)), maximumCount, tokenRange)
	return foundation.NSDictionaryFromID(rv)
}
// Finds an array of linguistic tags and token ranges for a given string range
// and linguistic unit.
//
// range: The range from which to return tags.
//
// unit: The linguistic unit. See [NLTokenUnit] for possible values.
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// scheme: The tag scheme. See [NLTagScheme] for possible values.
//
// options: The linguistic tagger options to use. See [NLTagger.Options] for possible
// values.
// //
// [NLTagger.Options]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options
//
// tokenRanges: Returns by reference an array of token ranges.
//
// # Return Value
// 
// An array of the tags in the requested range.
//
// # Discussion
// 
// When the returned array contains an entry that doesn’t have a
// corresponding tag scheme, that entry is an empty string (””).
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagsInRange:unit:scheme:options:tokenRanges:
func (t NLTagger) TagsInRangeUnitSchemeOptionsTokenRanges(range_ foundation.NSRange, unit NLTokenUnit, scheme NLTagScheme, options NLTaggerOptions, tokenRanges []foundation.NSValue) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tagsInRange:unit:scheme:options:tokenRanges:"), range_, unit, objc.String(string(scheme)), options, objectivec.IObjectSliceToNSArray(tokenRanges))
	return objc.ConvertSliceToStrings(rv)
}
// Returns the range of the linguistic unit containing the specified character
// index.
//
// characterIndex: The character index to begin examination.
//
// unit: The linguistic unit. For possible values, see [NLTokenUnit].
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// # Return Value
// 
// The range of the substring for the linguistic unit.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeAtIndex:unit:
func (t NLTagger) TokenRangeAtIndexUnit(characterIndex uint, unit NLTokenUnit) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("tokenRangeAtIndex:unit:"), characterIndex, unit)
	return foundation.NSRange(rv)
}
// Finds the entire range of all tokens of the specified linguistic unit
// contained completely or partially within the specified range.
//
// range: The range within the string to search for tokens.
//
// unit: The linguistic unit. For possible values, see [NLTokenUnit].
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// # Return Value
// 
// The smallest possible range that contains all of the tokens of the
// specified linguistic unit within the range specified in `range`. This
// result includes a token’s entire range if any part of that token is
// included within `range`. If the length of `range` is 0, this return value
// is equivalent to [TokenRangeAtIndexUnit].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tokenRangeForRange:unit:
func (t NLTagger) TokenRangeForRangeUnit(range_ foundation.NSRange, unit NLTokenUnit) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("tokenRangeForRange:unit:"), range_, unit)
	return foundation.NSRange(rv)
}

// Retrieves the tag schemes available for a particular unit (like word or
// sentence) and language on the current device.
//
// unit: The linguistic unit. For possible values, see [NLTokenUnit].
// //
// [NLTokenUnit]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
//
// language: The [NLLanguage] identifying the language.
//
// # Return Value
// 
// The supported tag schemes. For possible values, see [NLTagScheme].
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/availableTagSchemes(for:language:)
func (_NLTaggerClass NLTaggerClass) AvailableTagSchemesForUnitLanguage(unit NLTokenUnit, language NLLanguage) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NLTaggerClass.class), objc.Sel("availableTagSchemesForUnit:language:"), unit, objc.String(string(language)))
	return objc.ConvertSliceToStrings(rv)
}
// Asks the Natural Language framework to load any missing assets for a tag
// scheme onto the device for the given language.
//
// language: The language of the tag scheme that you’re asking the framework to load
// onto the device.
//
// tagScheme: The tag scheme that you’re asking the framework to load onto the device.
//
// completionHandler: A closure the framework uses to notify your app when the tag scheme request
// has completed.
//
// # Discussion
// 
// Before using this method, use [AvailableTagSchemesForUnitLanguage] to check
// whether a tag scheme is available on the device. When a tag scheme is
// unavailable for a specific language, it may be because the framework
// hasn’t loaded the support for that language.
// 
// Use this method to ask the [Natural Language] framework to load any missing
// assets for that tag scheme. This method returns immediately but the
// framework may need time to complete the request. When the framework
// completes the request, it notifies your app with the `completionHandler`
// you provided to the method. In your completion handler, use the `result`
// parameter to check whether the tag scheme is now available.
// 
// The [Natural Language] framework may call your `completionHandler`
// immediately if it knows the state of the tag scheme’s assets or if it
// experiences an error.
//
// [Natural Language]: https://developer.apple.com/documentation/NaturalLanguage
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/requestAssets(for:tagScheme:completionHandler:)
func (_NLTaggerClass NLTaggerClass) RequestAssetsForLanguageTagSchemeCompletionHandler(language NLLanguage, tagScheme NLTagScheme, completionHandler NLTaggerAssetsResultErrorHandler) {
_block2, _ := NewNLTaggerAssetsResultErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NLTaggerClass.class), objc.Sel("requestAssetsForLanguage:tagScheme:completionHandler:"), language, tagScheme, _block2)
}

// The string being analyzed by the linguistic tagger.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/string
func (t NLTagger) String() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
func (t NLTagger) SetString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setString:"), objc.String(value))
}
// The tag schemes configured for this linguistic tagger.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/tagSchemes
func (t NLTagger) TagSchemes() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tagSchemes"))
	return objc.ConvertSliceToStrings(rv)
}
// The dominant language of the string set for the linguistic tagger.
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
// In the example, [german] is the dominant language, indicating that the text
// is in German.
//
// [german]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguage/german
// [language]: https://developer.apple.com/documentation/NaturalLanguage/NLTagScheme/language
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/dominantLanguage
func (t NLTagger) DominantLanguage() NLLanguage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dominantLanguage"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}

// RequestAssetsForLanguageTagScheme is a synchronous wrapper around [NLTagger.RequestAssetsForLanguageTagSchemeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (tc NLTaggerClass) RequestAssetsForLanguageTagScheme(ctx context.Context, language NLLanguage, tagScheme NLTagScheme) (NLTaggerAssetsResult, error) {
	type result struct {
		val NLTaggerAssetsResult
		err error
	}
	done := make(chan result, 1)
	tc.RequestAssetsForLanguageTagSchemeCompletionHandler(language, tagScheme, func(val NLTaggerAssetsResult, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

