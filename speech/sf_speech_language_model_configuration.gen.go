// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechLanguageModelConfiguration] class.
var (
	_SFSpeechLanguageModelConfigurationClass     SFSpeechLanguageModelConfigurationClass
	_SFSpeechLanguageModelConfigurationClassOnce sync.Once
)

func getSFSpeechLanguageModelConfigurationClass() SFSpeechLanguageModelConfigurationClass {
	_SFSpeechLanguageModelConfigurationClassOnce.Do(func() {
		_SFSpeechLanguageModelConfigurationClass = SFSpeechLanguageModelConfigurationClass{class: objc.GetClass("SFSpeechLanguageModelConfiguration")}
	})
	return _SFSpeechLanguageModelConfigurationClass
}

// GetSFSpeechLanguageModelConfigurationClass returns the class object for SFSpeechLanguageModelConfiguration.
func GetSFSpeechLanguageModelConfigurationClass() SFSpeechLanguageModelConfigurationClass {
	return getSFSpeechLanguageModelConfigurationClass()
}

type SFSpeechLanguageModelConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechLanguageModelConfigurationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechLanguageModelConfigurationClass) Alloc() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object describing the location of a custom language model and
// specialized vocabulary.
//
// # Overview
//
// Pass this object to
// [SFSpeechLanguageModelClass.PrepareCustomLanguageModelForUrlConfigurationCompletion]
// to indicate where that method should create the custom language model file,
// and to [SFSpeechRecognitionRequest.CustomizedLanguageModel] or
// [customizedLanguage(modelConfiguration:)] to indicate where the system
// should find that model to use.
//
// # Creating a language model configuration
//
//   - [SFSpeechLanguageModelConfiguration.InitWithLanguageModel]: Creates a configuration with the location of a language model file.
//   - [SFSpeechLanguageModelConfiguration.InitWithLanguageModelVocabulary]: Creates a configuration with the locations of language model and vocabulary files.
//   - [SFSpeechLanguageModelConfiguration.InitWithLanguageModelVocabularyWeight]: Creates a configuration with the locations of language model and vocabulary files, and custom weight.
//
// # Inspecting a language model
//
//   - [SFSpeechLanguageModelConfiguration.LanguageModel]: The location of a compiled language model file.
//   - [SFSpeechLanguageModelConfiguration.Vocabulary]: The location of a compiled vocabulary file.
//   - [SFSpeechLanguageModelConfiguration.Weight]: The relative weight of the language model customization. Value must be between 0.0 and 1.0 inclusive.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration
//
// [customizedLanguage(modelConfiguration:)]: https://developer.apple.com/documentation/Speech/DictationTranscriber/ContentHint/customizedLanguage(modelConfiguration:)
type SFSpeechLanguageModelConfiguration struct {
	objectivec.Object
}

// SFSpeechLanguageModelConfigurationFromID constructs a [SFSpeechLanguageModelConfiguration] from an objc.ID.
//
// An object describing the location of a custom language model and
// specialized vocabulary.
func SFSpeechLanguageModelConfigurationFromID(id objc.ID) SFSpeechLanguageModelConfiguration {
	return SFSpeechLanguageModelConfiguration{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechLanguageModelConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechLanguageModelConfiguration] class.
//
// # Creating a language model configuration
//
//   - [ISFSpeechLanguageModelConfiguration.InitWithLanguageModel]: Creates a configuration with the location of a language model file.
//   - [ISFSpeechLanguageModelConfiguration.InitWithLanguageModelVocabulary]: Creates a configuration with the locations of language model and vocabulary files.
//   - [ISFSpeechLanguageModelConfiguration.InitWithLanguageModelVocabularyWeight]: Creates a configuration with the locations of language model and vocabulary files, and custom weight.
//
// # Inspecting a language model
//
//   - [ISFSpeechLanguageModelConfiguration.LanguageModel]: The location of a compiled language model file.
//   - [ISFSpeechLanguageModelConfiguration.Vocabulary]: The location of a compiled vocabulary file.
//   - [ISFSpeechLanguageModelConfiguration.Weight]: The relative weight of the language model customization. Value must be between 0.0 and 1.0 inclusive.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration
type ISFSpeechLanguageModelConfiguration interface {
	objectivec.IObject

	// Topic: Creating a language model configuration

	// Creates a configuration with the location of a language model file.
	InitWithLanguageModel(languageModel foundation.NSURL) SFSpeechLanguageModelConfiguration
	// Creates a configuration with the locations of language model and vocabulary files.
	InitWithLanguageModelVocabulary(languageModel foundation.NSURL, vocabulary foundation.NSURL) SFSpeechLanguageModelConfiguration
	// Creates a configuration with the locations of language model and vocabulary files, and custom weight.
	InitWithLanguageModelVocabularyWeight(languageModel foundation.NSURL, vocabulary foundation.NSURL, weight foundation.NSNumber) SFSpeechLanguageModelConfiguration

	// Topic: Inspecting a language model

	// The location of a compiled language model file.
	LanguageModel() foundation.NSURL
	// The location of a compiled vocabulary file.
	Vocabulary() foundation.NSURL
	// The relative weight of the language model customization. Value must be between 0.0 and 1.0 inclusive.
	Weight() foundation.NSNumber

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s SFSpeechLanguageModelConfiguration) Init() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechLanguageModelConfiguration) Autorelease() SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechLanguageModelConfiguration creates a new SFSpeechLanguageModelConfiguration instance.
func NewSFSpeechLanguageModelConfiguration() SFSpeechLanguageModelConfiguration {
	class := getSFSpeechLanguageModelConfigurationClass()
	rv := objc.Send[SFSpeechLanguageModelConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a configuration with the location of a language model file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:)
func NewSpeechLanguageModelConfigurationWithLanguageModel(languageModel foundation.NSURL) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLanguageModel:"), languageModel)
	return SFSpeechLanguageModelConfigurationFromID(rv)
}

// Creates a configuration with the locations of language model and vocabulary
// files.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:)
func NewSpeechLanguageModelConfigurationWithLanguageModelVocabulary(languageModel foundation.NSURL, vocabulary foundation.NSURL) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:"), languageModel, vocabulary)
	return SFSpeechLanguageModelConfigurationFromID(rv)
}

// Creates a configuration with the locations of language model and vocabulary
// files, and custom weight.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:weight:)
func NewSpeechLanguageModelConfigurationWithLanguageModelVocabularyWeight(languageModel foundation.NSURL, vocabulary foundation.NSURL, weight foundation.NSNumber) SFSpeechLanguageModelConfiguration {
	instance := getSFSpeechLanguageModelConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLanguageModel:vocabulary:weight:"), languageModel, vocabulary, weight)
	return SFSpeechLanguageModelConfigurationFromID(rv)
}

// Creates a configuration with the location of a language model file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:)
func (s SFSpeechLanguageModelConfiguration) InitWithLanguageModel(languageModel foundation.NSURL) SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s.ID, objc.Sel("initWithLanguageModel:"), languageModel)
	return rv
}

// Creates a configuration with the locations of language model and vocabulary
// files.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:)
func (s SFSpeechLanguageModelConfiguration) InitWithLanguageModelVocabulary(languageModel foundation.NSURL, vocabulary foundation.NSURL) SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s.ID, objc.Sel("initWithLanguageModel:vocabulary:"), languageModel, vocabulary)
	return rv
}

// Creates a configuration with the locations of language model and vocabulary
// files, and custom weight.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/init(languageModel:vocabulary:weight:)
func (s SFSpeechLanguageModelConfiguration) InitWithLanguageModelVocabularyWeight(languageModel foundation.NSURL, vocabulary foundation.NSURL, weight foundation.NSNumber) SFSpeechLanguageModelConfiguration {
	rv := objc.Send[SFSpeechLanguageModelConfiguration](s.ID, objc.Sel("initWithLanguageModel:vocabulary:weight:"), languageModel, vocabulary, weight)
	return rv
}
func (s SFSpeechLanguageModelConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The location of a compiled language model file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/languageModel
func (s SFSpeechLanguageModelConfiguration) LanguageModel() foundation.NSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageModel"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The location of a compiled vocabulary file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/vocabulary
func (s SFSpeechLanguageModelConfiguration) Vocabulary() foundation.NSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("vocabulary"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The relative weight of the language model customization. Value must be
// between 0.0 and 1.0 inclusive.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/Configuration/weight
func (s SFSpeechLanguageModelConfiguration) Weight() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("weight"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
