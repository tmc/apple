// Code generated from Apple documentation. DO NOT EDIT.

package speech

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles Called when the language model has been created.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlClientIdentifierConfigurationCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlClientIdentifierConfigurationIgnoresCacheCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlClientIdentifierConfigurationCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlClientIdentifierConfigurationIgnoresCacheCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationCompletion]
//   - [SFSpeechLanguageModel.PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// SFSpeechRecognitionResultErrorHandler handles The block to call when partial or final results are available, or when an error occurs.
//   - result: A [SFSpeechRecognitionResult](<doc://com.apple.speech/documentation/Speech/SFSpeechRecognitionResult>) containing the partial or final transcriptions of the audio content.
//   - error: An error object if a problem occurred. This parameter is `nil` if speech recognition was successful.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler]
type SFSpeechRecognitionResultErrorHandler = func(*SFSpeechRecognitionResult, error)

// NewSFSpeechRecognitionResultErrorBlock wraps a Go [SFSpeechRecognitionResultErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler]
func NewSFSpeechRecognitionResultErrorBlock(handler SFSpeechRecognitionResultErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *SFSpeechRecognitionResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SFSpeechRecognitionResultFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SFSpeechRecognizerAuthorizationStatusHandler handles The block to execute when your app’s authorization status is known.
//
// Used by:
//   - [SFSpeechRecognizer.RequestAuthorization]
type SFSpeechRecognizerAuthorizationStatusHandler = func(SFSpeechRecognizerAuthorizationStatus)

// NewSFSpeechRecognizerAuthorizationStatusBlock wraps a Go [SFSpeechRecognizerAuthorizationStatusHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SFSpeechRecognizer.RequestAuthorization]
func NewSFSpeechRecognizerAuthorizationStatusBlock(handler SFSpeechRecognizerAuthorizationStatusHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal SFSpeechRecognizerAuthorizationStatus) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
