// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSOrthography] class.
var (
	_NSOrthographyClass     NSOrthographyClass
	_NSOrthographyClassOnce sync.Once
)

func getNSOrthographyClass() NSOrthographyClass {
	_NSOrthographyClassOnce.Do(func() {
		_NSOrthographyClass = NSOrthographyClass{class: objc.GetClass("NSOrthography")}
	})
	return _NSOrthographyClass
}

// GetNSOrthographyClass returns the class object for NSOrthography.
func GetNSOrthographyClass() NSOrthographyClass {
	return getNSOrthographyClass()
}

type NSOrthographyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSOrthographyClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOrthographyClass) Alloc() NSOrthography {
	rv := objc.Send[NSOrthography](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of the linguistic content of natural language text, typically
// used for spelling and grammar checking.
//
// # Overview
// 
// Use [NSOrthography] objects to describe the linguistic content of a piece
// of text, including which scripts the text contains, a dominant language
// (and possibly other languages) for each script, and a dominant script and
// language for the text as a whole.
// 
// Scripts are uniformly described by four-letter ISO 15924 script codes, such
// as `"Latn"`, `"Grek"`, and `"Cyrl"`. The supertags `"Jpan"` and `"Kore"`
// are typically used for Japanese and Korean text, and `"Hans"` and `"Hant"`
// are typically used for Chinese text. The tag `"Zyyy"` is used if a specific
// script cannot be identified. See [Internationalization and Localization
// Guide] for more information.
// 
// Languages are uniformly described by BCP-47 tags (preferably in canonical
// form). The tag `"und"` is used if a specific language cannot be determined.
// 
// You typically work with orthography objects returned from methods and
// properties for classes like [NSLinguisticTagger] and [NSSpellChecker].
// 
// # Subclassing Notes
// 
// Subclasses must override the [NSOrthography.DominantScript] and [NSOrthography.LanguageMap] properties.
// These properties are set using [NSOrthography.InitWithDominantScriptLanguageMap] or
// [NSOrthography.OrthographyWithDominantScriptLanguageMap] in Objective-C.
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
// [NSSpellChecker]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
//
// # Creating Orthography Objects
//
//   - [NSOrthography.InitWithDominantScriptLanguageMap]: Creates an orthography object with the specified dominant script and language map.
//
// # Determining Correspondences Between Languages and Scripts
//
//   - [NSOrthography.LanguageMap]: A dictionary that maps script tags to arrays of language tags.
//   - [NSOrthography.DominantLanguage]: The first language in the list of languages for the dominant script.
//   - [NSOrthography.DominantScript]: The dominant script for the text.
//   - [NSOrthography.DominantLanguageForScript]: Returns the dominant language for the specified script.
//   - [NSOrthography.LanguagesForScript]: Returns the list of languages for the specified script.
//   - [NSOrthography.AllScripts]: The scripts appearing as keys in the language map.
//   - [NSOrthography.AllLanguages]: The languages appearing in values of the language map.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography
type NSOrthography struct {
	objectivec.Object
}

// NSOrthographyFromID constructs a [NSOrthography] from an objc.ID.
//
// A description of the linguistic content of natural language text, typically
// used for spelling and grammar checking.
func NSOrthographyFromID(id objc.ID) NSOrthography {
	return NSOrthography{objectivec.Object{ID: id}}
}
// NOTE: NSOrthography adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOrthography] class.
//
// # Creating Orthography Objects
//
//   - [INSOrthography.InitWithDominantScriptLanguageMap]: Creates an orthography object with the specified dominant script and language map.
//
// # Determining Correspondences Between Languages and Scripts
//
//   - [INSOrthography.LanguageMap]: A dictionary that maps script tags to arrays of language tags.
//   - [INSOrthography.DominantLanguage]: The first language in the list of languages for the dominant script.
//   - [INSOrthography.DominantScript]: The dominant script for the text.
//   - [INSOrthography.DominantLanguageForScript]: Returns the dominant language for the specified script.
//   - [INSOrthography.LanguagesForScript]: Returns the list of languages for the specified script.
//   - [INSOrthography.AllScripts]: The scripts appearing as keys in the language map.
//   - [INSOrthography.AllLanguages]: The languages appearing in values of the language map.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography
type INSOrthography interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Orthography Objects

	// Creates an orthography object with the specified dominant script and language map.
	InitWithDominantScriptLanguageMap(script string, map_ INSDictionary) NSOrthography

	// Topic: Determining Correspondences Between Languages and Scripts

	// A dictionary that maps script tags to arrays of language tags.
	LanguageMap() INSDictionary
	// The first language in the list of languages for the dominant script.
	DominantLanguage() string
	// The dominant script for the text.
	DominantScript() string
	// Returns the dominant language for the specified script.
	DominantLanguageForScript(script string) string
	// Returns the list of languages for the specified script.
	LanguagesForScript(script string) []string
	// The scripts appearing as keys in the language map.
	AllScripts() []string
	// The languages appearing in values of the language map.
	AllLanguages() []string
}

// Init initializes the instance.
func (o NSOrthography) Init() NSOrthography {
	rv := objc.Send[NSOrthography](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOrthography) Autorelease() NSOrthography {
	rv := objc.Send[NSOrthography](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOrthography creates a new NSOrthography instance.
func NewNSOrthography() NSOrthography {
	class := getNSOrthographyClass()
	rv := objc.Send[NSOrthography](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/init(coder:)
func NewOrthographyWithCoder(coder INSCoder) NSOrthography {
	instance := getNSOrthographyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSOrthographyFromID(rv)
}

// Creates an orthography object with the specified dominant script and
// language map.
//
// script: The dominant script.
//
// map: A dictionary mapping ISO 15924 script codes to arrays of BCP-47 language
// tags.
//
// # Return Value
// 
// An orthography object initialized with the specified script and language
// map.
//
// # Discussion
// 
// You typically use the [DefaultOrthographyForLanguage] method to create
// orthography objects with automatic language mapping. Use this initializer
// only if you need to override the script associated with one or more
// languages.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/init(dominantScript:languageMap:)
func NewOrthographyWithDominantScriptLanguageMap(script string, map_ INSDictionary) NSOrthography {
	instance := getNSOrthographyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDominantScript:languageMap:"), objc.String(script), map_)
	return NSOrthographyFromID(rv)
}

// Creates an orthography object with the specified dominant script and
// language map.
//
// script: The dominant script.
//
// map: A dictionary mapping ISO 15924 script codes to arrays of BCP-47 language
// tags.
//
// # Return Value
// 
// An orthography object initialized with the specified script and language
// map.
//
// # Discussion
// 
// You typically use the [DefaultOrthographyForLanguage] method to create
// orthography objects with automatic language mapping. Use this initializer
// only if you need to override the script associated with one or more
// languages.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/init(dominantScript:languageMap:)
func (o NSOrthography) InitWithDominantScriptLanguageMap(script string, map_ INSDictionary) NSOrthography {
	rv := objc.Send[NSOrthography](o.ID, objc.Sel("initWithDominantScript:languageMap:"), objc.String(script), map_)
	return rv
}
// Returns the dominant language for the specified script.
//
// script: The specified script.
//
// # Discussion
// 
// The value of this property is a BCP-47 language tag, such as `"en"` or
// `"fr"`, that identifies the dominant language.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage(forScript:)
func (o NSOrthography) DominantLanguageForScript(script string) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dominantLanguageForScript:"), objc.String(script))
	return NSStringFromID(rv).String()
}
// Returns the list of languages for the specified script.
//
// script: The specified script.
//
// # Discussion
// 
// The value of this property is an array of BCP-47 language tags, such as
// `"en"` or `"fr"`, that identify the languages.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/languages(forScript:)
func (o NSOrthography) LanguagesForScript(script string) []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("languagesForScript:"), objc.String(script))
	return objc.ConvertSliceToStrings(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/init(coder:)
func (o NSOrthography) InitWithCoder(coder INSCoder) NSOrthography {
	rv := objc.Send[NSOrthography](o.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (o NSOrthography) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an orthography object with the default language map for
// the specified language.
//
// language: A BCP-47 tag identifying the language.
//
// # Discussion
// 
// This method automatically determines the script for the specified language.
// For example, the default orthography for the Hindi language has a language
// map with a single key, `"Deva"` (the ISO 15924 script code for Devanagari),
// that has a corresponding value of an array containing the element `"hi"`
// (the BCP-47 identifier for Hindi).
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/defaultOrthography(forLanguage:)
func (_NSOrthographyClass NSOrthographyClass) DefaultOrthographyForLanguage(language string) NSOrthography {
	rv := objc.Send[objc.ID](objc.ID(_NSOrthographyClass.class), objc.Sel("defaultOrthographyForLanguage:"), objc.String(language))
	return NSOrthographyFromID(rv)
}
// Creates and returns an orthography object with the specified dominant
// script and language map.
//
// script: The dominant script.
//
// map: A dictionary mapping ISO 15924 script codes to arrays of BCP-47 language
// tags.
//
// # Return Value
// 
// An orthography object initialized with the specified script and language
// map.
//
// # Discussion
// 
// You typically use the [DefaultOrthographyForLanguage] method to create
// orthography objects with automatic language mapping. Use this initializer
// only if you need to override the script associated with one or more
// languages.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/orthographyWithDominantScript:languageMap:
func (_NSOrthographyClass NSOrthographyClass) OrthographyWithDominantScriptLanguageMap(script string, map_ INSDictionary) NSOrthography {
	rv := objc.Send[objc.ID](objc.ID(_NSOrthographyClass.class), objc.Sel("orthographyWithDominantScript:languageMap:"), objc.String(script), map_)
	return NSOrthographyFromID(rv)
}

// A dictionary that maps script tags to arrays of language tags.
//
// # Discussion
// 
// The dictionary’s keys are ISO 15924 script codes (such as `"Latn"` or
// `"Cyrl"`) and its values are arrays of BCP-47 language tags (such as
// `"en"`, `"fr"`, or `"de"`).
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/languageMap
func (o NSOrthography) LanguageMap() INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("languageMap"))
	return NSDictionaryFromID(objc.ID(rv))
}
// The first language in the list of languages for the dominant script.
//
// # Discussion
// 
// The value of this property is a BCP-47 language tag, such as `"en"` or
// `"fr"`, that identifies the dominant language.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantLanguage
func (o NSOrthography) DominantLanguage() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dominantLanguage"))
	return NSStringFromID(rv).String()
}
// The dominant script for the text.
//
// # Discussion
// 
// The value of this property is an ISO 15924 script code, such as `"Latn"` or
// `"Cyrl"`, that identifies the dominant script.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/dominantScript
func (o NSOrthography) DominantScript() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dominantScript"))
	return NSStringFromID(rv).String()
}
// The scripts appearing as keys in the language map.
//
// # Discussion
// 
// The value of this property is an array of ISO 15924 script codes, such as
// `"Latn"` or `"Cyrl"`, that identify the scripts.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/allScripts
func (o NSOrthography) AllScripts() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("allScripts"))
	return objc.ConvertSliceToStrings(rv)
}
// The languages appearing in values of the language map.
//
// # Discussion
// 
// The value of this property is an array of BCP-47 language tags, such as
// `"en"` or `"fr"`, that identify the languages.
//
// See: https://developer.apple.com/documentation/Foundation/NSOrthography/allLanguages
func (o NSOrthography) AllLanguages() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("allLanguages"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

