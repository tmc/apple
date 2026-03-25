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

// The class instance for the [NLGazetteer] class.
var (
	_NLGazetteerClass     NLGazetteerClass
	_NLGazetteerClassOnce sync.Once
)

func getNLGazetteerClass() NLGazetteerClass {
	_NLGazetteerClassOnce.Do(func() {
		_NLGazetteerClass = NLGazetteerClass{class: objc.GetClass("NLGazetteer")}
	})
	return _NLGazetteerClass
}

// GetNLGazetteerClass returns the class object for NLGazetteer.
func GetNLGazetteerClass() NLGazetteerClass {
	return getNLGazetteerClass()
}

type NLGazetteerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLGazetteerClass) Alloc() NLGazetteer {
	rv := objc.Send[NLGazetteer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A collection of terms and their labels, which take precedence over a word
// tagger.
//
// # Overview
// 
// Use an [NLGazetteer] to augment an [NLTagger] when you need to tag a
// specific set of terms (single words or short phrases) with a label.
// Typically, you add one gazetteer per language, or one language-independent
// gazetteer, to an [NLTagger] with its [SetGazetteersForTagScheme] method.
// The tagger uses its gazetteers to look up each term it processes. If a
// gazetteer has a label for a term, the tagger uses that label to tag the
// term, instead of inferring a tag itself.
// 
// Typically, you create a gazetteer at development time, such as in a macOS
// playground, with Create ML’s [MLGazetteer]. Alternatively, you can create
// an [NLGazetteer] at runtime by using [NLGazetteer.InitWithDictionaryLanguageError].
//
// [MLGazetteer]: https://developer.apple.com/documentation/CreateML/MLGazetteer
//
// # Creating a Gazetteer
//
//   - [NLGazetteer.InitWithContentsOfURLError]: Creates a Natural Language gazetteer from a model created with the Create ML framework.
//   - [NLGazetteer.InitWithDataError]: Creates a gazetteer from a data instance.
//   - [NLGazetteer.InitWithDictionaryLanguageError]: Creates a gazetteer from a set of labels for terms represented by a dictionary.
//
// # Looking Up Labels for Terms
//
//   - [NLGazetteer.LabelForString]: Retrieves the label for the given term.
//
// # Inspecting a Gazetteer
//
//   - [NLGazetteer.Data]: The gazetteer represented as a data instance.
//   - [NLGazetteer.Language]: The language of the gazetteer.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer
type NLGazetteer struct {
	objectivec.Object
}

// NLGazetteerFromID constructs a [NLGazetteer] from an objc.ID.
//
// A collection of terms and their labels, which take precedence over a word
// tagger.
func NLGazetteerFromID(id objc.ID) NLGazetteer {
	return NLGazetteer{objectivec.Object{ID: id}}
}
// NOTE: NLGazetteer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLGazetteer] class.
//
// # Creating a Gazetteer
//
//   - [INLGazetteer.InitWithContentsOfURLError]: Creates a Natural Language gazetteer from a model created with the Create ML framework.
//   - [INLGazetteer.InitWithDataError]: Creates a gazetteer from a data instance.
//   - [INLGazetteer.InitWithDictionaryLanguageError]: Creates a gazetteer from a set of labels for terms represented by a dictionary.
//
// # Looking Up Labels for Terms
//
//   - [INLGazetteer.LabelForString]: Retrieves the label for the given term.
//
// # Inspecting a Gazetteer
//
//   - [INLGazetteer.Data]: The gazetteer represented as a data instance.
//   - [INLGazetteer.Language]: The language of the gazetteer.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer
type INLGazetteer interface {
	objectivec.IObject

	// Topic: Creating a Gazetteer

	// Creates a Natural Language gazetteer from a model created with the Create ML framework.
	InitWithContentsOfURLError(url foundation.INSURL) (NLGazetteer, error)
	// Creates a gazetteer from a data instance.
	InitWithDataError(data foundation.INSData) (NLGazetteer, error)
	// Creates a gazetteer from a set of labels for terms represented by a dictionary.
	InitWithDictionaryLanguageError(dictionary foundation.INSDictionary, language NLLanguage) (NLGazetteer, error)

	// Topic: Looking Up Labels for Terms

	// Retrieves the label for the given term.
	LabelForString(string_ string) string

	// Topic: Inspecting a Gazetteer

	// The gazetteer represented as a data instance.
	Data() foundation.INSData
	// The language of the gazetteer.
	Language() NLLanguage
}

// Init initializes the instance.
func (g NLGazetteer) Init() NLGazetteer {
	rv := objc.Send[NLGazetteer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NLGazetteer) Autorelease() NLGazetteer {
	rv := objc.Send[NLGazetteer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLGazetteer creates a new NLGazetteer instance.
func NewNLGazetteer() NLGazetteer {
	class := getNLGazetteerClass()
	rv := objc.Send[NLGazetteer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Natural Language gazetteer from a model created with the Create
// ML framework.
//
// url: The location of the .`mlmodel` file that contains a gazetteer.
//
// # Discussion
// 
// Use this initializer to create an [NLGazetteer] from an
// `XCUIElementTypeMlmodel` file saved by [MLGazetteer] in the `Create ML`
// framework.
//
// [MLGazetteer]: https://developer.apple.com/documentation/CreateML/MLGazetteer
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(contentsOf:)
func NewGazetteerWithContentsOfURLError(url foundation.INSURL) (NLGazetteer, error) {
	var errorPtr objc.ID
	instance := getNLGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil
}

// Creates a gazetteer from a data instance.
//
// data: A gazetteer contained in a data instance.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(data:)
func NewGazetteerWithDataError(data foundation.INSData) (NLGazetteer, error) {
	var errorPtr objc.ID
	instance := getNLGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil
}

// Creates a gazetteer from a set of labels for terms represented by a
// dictionary.
//
// dictionary: The dictionary of labels and an array of terms for each label.
//
// language: The language of the text in the dictionary.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(dictionary:language:)
func NewGazetteerWithDictionaryLanguageError(dictionary foundation.INSDictionary, language NLLanguage) (NLGazetteer, error) {
	var errorPtr objc.ID
	instance := getNLGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:language:error:"), dictionary, objc.String(string(language)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil
}

// Creates a Natural Language gazetteer from a model created with the Create
// ML framework.
//
// url: The location of the .`mlmodel` file that contains a gazetteer.
//
// # Discussion
// 
// Use this initializer to create an [NLGazetteer] from an
// `XCUIElementTypeMlmodel` file saved by [MLGazetteer] in the `Create ML`
// framework.
//
// [MLGazetteer]: https://developer.apple.com/documentation/CreateML/MLGazetteer
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(contentsOf:)
func (g NLGazetteer) InitWithContentsOfURLError(url foundation.INSURL) (NLGazetteer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil

}
// Creates a gazetteer from a data instance.
//
// data: A gazetteer contained in a data instance.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(data:)
func (g NLGazetteer) InitWithDataError(data foundation.INSData) (NLGazetteer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil

}
// Creates a gazetteer from a set of labels for terms represented by a
// dictionary.
//
// dictionary: The dictionary of labels and an array of terms for each label.
//
// language: The language of the text in the dictionary.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/init(dictionary:language:)
func (g NLGazetteer) InitWithDictionaryLanguageError(dictionary foundation.INSDictionary, language NLLanguage) (NLGazetteer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithDictionary:language:error:"), dictionary, objc.String(string(language)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil

}
// Retrieves the label for the given term.
//
// string: The term used to find a label.
//
// # Return Value
// 
// A string if the term is in the vocabulary; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/label(for:)
func (g NLGazetteer) LabelForString(string_ string) string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("labelForString:"), objc.String(string_))
	return foundation.NSStringFromID(rv).String()
}

// Creates a gazetteer from a set of labels for terms represented by a
// dictionary and saves the gazetteer to a file.
//
// dictionary: The dictionary of labels and an array of terms for each label.
//
// language: The language of the text in the dictionary.
//
// url: The location in the file system to which the file should be written.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/write(_:language:to:)
func (_NLGazetteerClass NLGazetteerClass) WriteGazetteerForDictionaryLanguageToURLError(dictionary foundation.INSDictionary, language NLLanguage, url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NLGazetteerClass.class), objc.Sel("writeGazetteerForDictionary:language:toURL:error:"), dictionary, objc.String(string(language)), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeGazetteerForDictionary:language:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Creates a Natural Language gazetteer from a model created with the Create
// ML framework.
//
// url: The location of the .`mlmodel` file that contains a gazetteer.
//
// # Discussion
// 
// Use this initializer to create an [NLGazetteer] from an
// `XCUIElementTypeMlmodel` file saved by [MLGazetteer] in the `Create ML`
// framework.
//
// [MLGazetteer]: https://developer.apple.com/documentation/CreateML/MLGazetteer
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/gazetteerWithContentsOfURL:error:
func (_NLGazetteerClass NLGazetteerClass) GazetteerWithContentsOfURLError(url foundation.INSURL) (NLGazetteer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NLGazetteerClass.class), objc.Sel("gazetteerWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLGazetteerFromID(rv), nil

}

// The gazetteer represented as a data instance.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/data
func (g NLGazetteer) Data() foundation.INSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The language of the gazetteer.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer/language
func (g NLGazetteer) Language() NLLanguage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("language"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}

