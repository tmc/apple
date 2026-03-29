// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLContextualEmbedding] class.
var (
	_NLContextualEmbeddingClass     NLContextualEmbeddingClass
	_NLContextualEmbeddingClassOnce sync.Once
)

func getNLContextualEmbeddingClass() NLContextualEmbeddingClass {
	_NLContextualEmbeddingClassOnce.Do(func() {
		_NLContextualEmbeddingClass = NLContextualEmbeddingClass{class: objc.GetClass("NLContextualEmbedding")}
	})
	return _NLContextualEmbeddingClass
}

// GetNLContextualEmbeddingClass returns the class object for NLContextualEmbedding.
func GetNLContextualEmbeddingClass() NLContextualEmbeddingClass {
	return getNLContextualEmbeddingClass()
}

type NLContextualEmbeddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLContextualEmbeddingClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLContextualEmbeddingClass) Alloc() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A model that computes sequences of embedding vectors for natural language
// utterances.
//
// # Overview
// 
// Starting in iOS 17 and macOS 14, the framework supports 27 languages across
// three models:
// 
// - Latin — including Croatian, Czech, Danish, Dutch, English, Finnish,
// French, German, Hungarian, Indonesian, Italian, Norwegian, Polish,
// Portuguese, Romanian, Slovak, Swedish, Spanish, Turkish, and Vietnamese -
// Cyrillic — including Bulgarian, Kazakh, Russian, and Ukrainian - Chinese,
// Japanese, and Korean
// 
// In iOS 18 and macOS 15, the framework expands language support to include
// three additional models:
// 
// - Arabic - Thai - Indic — including Hindi, Marathi, Bangla, Urdu,
// Punjabi, Gujarati, Tamil, Telugu, Kannada, and Malayalam
//
// # Inspecting the contextual embedding
//
//   - [NLContextualEmbedding.Dimension]: The number of dimensions in the script’s vector space.
//   - [NLContextualEmbedding.HasAvailableAssets]: A Boolean value that indicates whether assets are available to load.
//   - [NLContextualEmbedding.Languages]: The languages of the text in the contextual embedding.
//   - [NLContextualEmbedding.MaximumSequenceLength]: The maximum number of embedding vectors the model generates, in sequence.
//   - [NLContextualEmbedding.ModelIdentifier]: The model identifier.
//   - [NLContextualEmbedding.Revision]: The revision of the contextual embedding.
//   - [NLContextualEmbedding.Scripts]: The scripts of the text in the contextual embedding.
//
// # Requesting assets
//
//   - [NLContextualEmbedding.RequestEmbeddingAssetsWithCompletionHandler]: Requests assets for an embedding, if available.
//
// # Loading and unloading assets
//
//   - [NLContextualEmbedding.LoadWithError]: Loads the embedding model.
//   - [NLContextualEmbedding.Unload]: Unloads the embedding model.
//
// # Applying an embedding
//
//   - [NLContextualEmbedding.EmbeddingResultForStringLanguageError]: Applies an embedding to a string and obtains the resulting embedding vectors.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding
type NLContextualEmbedding struct {
	objectivec.Object
}

// NLContextualEmbeddingFromID constructs a [NLContextualEmbedding] from an objc.ID.
//
// A model that computes sequences of embedding vectors for natural language
// utterances.
func NLContextualEmbeddingFromID(id objc.ID) NLContextualEmbedding {
	return NLContextualEmbedding{objectivec.Object{ID: id}}
}
// NOTE: NLContextualEmbedding adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLContextualEmbedding] class.
//
// # Inspecting the contextual embedding
//
//   - [INLContextualEmbedding.Dimension]: The number of dimensions in the script’s vector space.
//   - [INLContextualEmbedding.HasAvailableAssets]: A Boolean value that indicates whether assets are available to load.
//   - [INLContextualEmbedding.Languages]: The languages of the text in the contextual embedding.
//   - [INLContextualEmbedding.MaximumSequenceLength]: The maximum number of embedding vectors the model generates, in sequence.
//   - [INLContextualEmbedding.ModelIdentifier]: The model identifier.
//   - [INLContextualEmbedding.Revision]: The revision of the contextual embedding.
//   - [INLContextualEmbedding.Scripts]: The scripts of the text in the contextual embedding.
//
// # Requesting assets
//
//   - [INLContextualEmbedding.RequestEmbeddingAssetsWithCompletionHandler]: Requests assets for an embedding, if available.
//
// # Loading and unloading assets
//
//   - [INLContextualEmbedding.LoadWithError]: Loads the embedding model.
//   - [INLContextualEmbedding.Unload]: Unloads the embedding model.
//
// # Applying an embedding
//
//   - [INLContextualEmbedding.EmbeddingResultForStringLanguageError]: Applies an embedding to a string and obtains the resulting embedding vectors.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding
type INLContextualEmbedding interface {
	objectivec.IObject

	// Topic: Inspecting the contextual embedding

	// The number of dimensions in the script’s vector space.
	Dimension() uint
	// A Boolean value that indicates whether assets are available to load.
	HasAvailableAssets() bool
	// The languages of the text in the contextual embedding.
	Languages() []string
	// The maximum number of embedding vectors the model generates, in sequence.
	MaximumSequenceLength() uint
	// The model identifier.
	ModelIdentifier() string
	// The revision of the contextual embedding.
	Revision() uint
	// The scripts of the text in the contextual embedding.
	Scripts() []string

	// Topic: Requesting assets

	// Requests assets for an embedding, if available.
	RequestEmbeddingAssetsWithCompletionHandler(completionHandler NLContextualEmbeddingAssetsResultErrorHandler)

	// Topic: Loading and unloading assets

	// Loads the embedding model.
	LoadWithError() (bool, error)
	// Unloads the embedding model.
	Unload()

	// Topic: Applying an embedding

	// Applies an embedding to a string and obtains the resulting embedding vectors.
	EmbeddingResultForStringLanguageError(string_ string, language NLLanguage) (INLContextualEmbeddingResult, error)
}

// Init initializes the instance.
func (c NLContextualEmbedding) Init() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NLContextualEmbedding) Autorelease() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLContextualEmbedding creates a new NLContextualEmbedding instance.
func NewNLContextualEmbedding() NLContextualEmbedding {
	class := getNLContextualEmbeddingClass()
	rv := objc.Send[NLContextualEmbedding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a contextual embedding from a language.
//
// language: The language the framework uses to find the most recent embedding suitable
// for the value you specify.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func NewContextualEmbeddingWithLanguage(language NLLanguage) NLContextualEmbedding {
	rv := objc.Send[objc.ID](objc.ID(getNLContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithLanguage:"), objc.String(string(language)))
	return NLContextualEmbeddingFromID(rv)
}

// Creates a contextual embedding from a model identifier.
//
// modelIdentifier: The model identifier that identifies the embedding model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func NewContextualEmbeddingWithModelIdentifier(modelIdentifier string) NLContextualEmbedding {
	rv := objc.Send[objc.ID](objc.ID(getNLContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), objc.String(modelIdentifier))
	return NLContextualEmbeddingFromID(rv)
}

// Creates a contextual embedding from a script.
//
// script: The script the framework uses to find the most recent embedding suitable
// for the value you specify.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func NewContextualEmbeddingWithScript(script NLScript) NLContextualEmbedding {
	rv := objc.Send[objc.ID](objc.ID(getNLContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithScript:"), objc.String(string(script)))
	return NLContextualEmbeddingFromID(rv)
}

// Requests assets for an embedding, if available.
//
// completionHandler: A completion handler the system calls after it finishes the request.
//
// # Discussion
// 
// You use a contextual embedding after loading the necessary assets onto the
// device. Use [HasAvailableAssets] to determine whether assets are available.
// 
// This method returns immediately if the framework knows the state of the
// assets or if an error occurs.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/requestAssets(completionHandler:)
func (c NLContextualEmbedding) RequestEmbeddingAssetsWithCompletionHandler(completionHandler NLContextualEmbeddingAssetsResultErrorHandler) {
_block0, _ := NewNLContextualEmbeddingAssetsResultErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("requestEmbeddingAssetsWithCompletionHandler:"), _block0)
}
// Loads the embedding model.
//
// # Discussion
// 
// When you create a contextual embedding, the model isn’t loaded until you
// need it, by default. Use [LoadWithError] and [Unload] to control when to
// load and unload the model.
// 
// The method fails if the necessary assets — for the model you specify —
// aren’t on device. Use [HasAvailableAssets] and
// [RequestEmbeddingAssetsWithCompletionHandler] to manage the assets.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/load()
func (c NLContextualEmbedding) LoadWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("loadWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadWithError: returned NO with nil NSError")
	}
	return rv, nil

}
// Unloads the embedding model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/unload()
func (c NLContextualEmbedding) Unload() {
	objc.Send[objc.ID](c.ID, objc.Sel("unload"))
}
// Applies an embedding to a string and obtains the resulting embedding
// vectors.
//
// string: The string to apply an embedding to.
//
// language: The language of the string.
//
// # Return Value
// 
// An embedding result.
//
// # Discussion
// 
// If the language of the string is unknown, the framework infers it from the
// string you specify.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/embeddingResult(for:language:)
func (c NLContextualEmbedding) EmbeddingResultForStringLanguageError(string_ string, language NLLanguage) (INLContextualEmbeddingResult, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("embeddingResultForString:language:error:"), objc.String(string_), objc.String(string(language)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLContextualEmbeddingResult{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLContextualEmbeddingResultFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/contextualEmbeddings(forValues:)
func (_NLContextualEmbeddingClass NLContextualEmbeddingClass) ContextualEmbeddingsForValues(valuesDictionary foundation.INSDictionary) []NLContextualEmbedding {
	rv := objc.Send[[]objc.ID](objc.ID(_NLContextualEmbeddingClass.class), objc.Sel("contextualEmbeddingsForValues:"), valuesDictionary)
	return objc.ConvertSlice(rv, func(id objc.ID) NLContextualEmbedding {
		return NLContextualEmbeddingFromID(id)
	})
}

// The number of dimensions in the script’s vector space.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/dimension
func (c NLContextualEmbedding) Dimension() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dimension"))
	return rv
}
// A Boolean value that indicates whether assets are available to load.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/hasAvailableAssets
func (c NLContextualEmbedding) HasAvailableAssets() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasAvailableAssets"))
	return rv
}
// The languages of the text in the contextual embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/languages
func (c NLContextualEmbedding) Languages() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("languages"))
	return objc.ConvertSliceToStrings(rv)
}
// The maximum number of embedding vectors the model generates, in sequence.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/maximumSequenceLength
func (c NLContextualEmbedding) MaximumSequenceLength() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maximumSequenceLength"))
	return rv
}
// The model identifier.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/modelIdentifier
func (c NLContextualEmbedding) ModelIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The revision of the contextual embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/revision
func (c NLContextualEmbedding) Revision() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("revision"))
	return rv
}
// The scripts of the text in the contextual embedding.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/scripts
func (c NLContextualEmbedding) Scripts() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("scripts"))
	return objc.ConvertSliceToStrings(rv)
}

// RequestEmbeddingAssets is a synchronous wrapper around [NLContextualEmbedding.RequestEmbeddingAssetsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c NLContextualEmbedding) RequestEmbeddingAssets(ctx context.Context) (NLContextualEmbeddingAssetsResult, error) {
	type result struct {
		val NLContextualEmbeddingAssetsResult
		err error
	}
	done := make(chan result, 1)
	c.RequestEmbeddingAssetsWithCompletionHandler(func(val NLContextualEmbeddingAssetsResult, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

