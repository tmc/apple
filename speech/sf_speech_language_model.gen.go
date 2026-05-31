// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechLanguageModel] class.
var (
	_SFSpeechLanguageModelClass     SFSpeechLanguageModelClass
	_SFSpeechLanguageModelClassOnce sync.Once
)

func getSFSpeechLanguageModelClass() SFSpeechLanguageModelClass {
	_SFSpeechLanguageModelClassOnce.Do(func() {
		_SFSpeechLanguageModelClass = SFSpeechLanguageModelClass{class: objc.GetClass("SFSpeechLanguageModel")}
	})
	return _SFSpeechLanguageModelClass
}

// GetSFSpeechLanguageModelClass returns the class object for SFSpeechLanguageModel.
func GetSFSpeechLanguageModelClass() SFSpeechLanguageModelClass {
	return getSFSpeechLanguageModelClass()
}

type SFSpeechLanguageModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechLanguageModelClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechLanguageModelClass) Alloc() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A language model built from custom training data.
//
// # Overview
//
// Create this object using
// [SFSpeechLanguageModelClass.PrepareCustomLanguageModelForUrlConfigurationCompletion]
// or
// [SFSpeechLanguageModelClass.PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel
type SFSpeechLanguageModel struct {
	objectivec.Object
}

// SFSpeechLanguageModelFromID constructs a [SFSpeechLanguageModel] from an objc.ID.
//
// A language model built from custom training data.
func SFSpeechLanguageModelFromID(id objc.ID) SFSpeechLanguageModel {
	return SFSpeechLanguageModel{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechLanguageModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechLanguageModel] class.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel
type ISFSpeechLanguageModel interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SFSpeechLanguageModel) Init() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechLanguageModel) Autorelease() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechLanguageModel creates a new SFSpeechLanguageModel instance.
func NewSFSpeechLanguageModel() SFSpeechLanguageModel {
	class := getSFSpeechLanguageModelClass()
	rv := objc.Send[SFSpeechLanguageModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a language model from custom training data.
//
// asset: The URL of a file containing custom training data. Create this file with
// [export(to:)].
//
// configuration: An object listing the URLs at which this method should create the language
// model and compiled vocabulary from the training data.
//
// completion: Called when the language model has been created.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:configuration:completion:)
//
// [export(to:)]: https://developer.apple.com/documentation/Speech/SFCustomLanguageModelData/export(to:)
func (_SFSpeechLanguageModelClass SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationCompletion(asset foundation.NSURL, configuration ISFSpeechLanguageModelConfiguration, completion ErrorHandler) {
	_block2, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](objc.ID(_SFSpeechLanguageModelClass.class), objc.Sel("prepareCustomLanguageModelForUrl:configuration:completion:"), asset, configuration, _block2)
}

// Creates a language model from custom training data.
//
// asset: The URL of a file containing custom training data. Create this file with
// [export(to:)].
//
// configuration: An object listing the URLs at which this method should create the language
// model and compiled vocabulary from the training data.
//
// ignoresCache: If `true`, the language model identified by the configuration will be
// recreated even if the `asset` file is unchanged.
//
// completion: Called when the language model has been created.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:configuration:ignoresCache:completion:)
//
// [export(to:)]: https://developer.apple.com/documentation/Speech/SFCustomLanguageModelData/export(to:)
func (_SFSpeechLanguageModelClass SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion(asset foundation.NSURL, configuration ISFSpeechLanguageModelConfiguration, ignoresCache bool, completion ErrorHandler) {
	_block3, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](objc.ID(_SFSpeechLanguageModelClass.class), objc.Sel("prepareCustomLanguageModelForUrl:configuration:ignoresCache:completion:"), asset, configuration, ignoresCache, _block3)
}

// PrepareCustomLanguageModelForUrlConfigurationCompletionSync is a synchronous wrapper around [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationCompletionSync(ctx context.Context, asset foundation.NSURL, configuration ISFSpeechLanguageModelConfiguration) error {
	done := make(chan error, 1)
	sc.PrepareCustomLanguageModelForUrlConfigurationCompletion(asset, configuration, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletionSync is a synchronous wrapper around [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletionSync(ctx context.Context, asset foundation.NSURL, configuration ISFSpeechLanguageModelConfiguration, ignoresCache bool) error {
	done := make(chan error, 1)
	sc.PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion(asset, configuration, ignoresCache, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
