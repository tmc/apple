// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// TTSSpeechSynthesizerDelegate protocol.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate
type TTSSpeechSynthesizerDelegate interface {
	objectivec.IObject
}

// TTSSpeechSynthesizerDelegateObject wraps an existing Objective-C object that conforms to the TTSSpeechSynthesizerDelegate protocol.
type TTSSpeechSynthesizerDelegateObject struct {
	objectivec.Object
}
func (o TTSSpeechSynthesizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TTSSpeechSynthesizerDelegateObjectFromID constructs a [TTSSpeechSynthesizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TTSSpeechSynthesizerDelegateObjectFromID(id objc.ID) TTSSpeechSynthesizerDelegateObject {
	return TTSSpeechSynthesizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didContinueSpeakingRequest:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didContinueSpeakingRequest:"), synthesizer, request)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didEncounterMarker:forRequest:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterMarker:forRequest:"), synthesizer, marker, request)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didFinishSpeakingRequest:successfully:phonemesSpoken:withError:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, spoken objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:phonemesSpoken:withError:"), synthesizer, request, successfully, spoken, error_)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didFinishSpeakingRequest:successfully:withError:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:withError:"), synthesizer, request, successfully, error_)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didPauseSpeakingRequest:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didPauseSpeakingRequest:"), synthesizer, request)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:didStartSpeakingRequest:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidStartSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didStartSpeakingRequest:"), synthesizer, request)
	}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechSynthesizerDelegate/speechSynthesizer:willSpeakRangeOfSpeechString:forRequest:
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakRangeOfSpeechStringForRequest(synthesizer objectivec.IObject, string_ foundation.NSRange, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakRangeOfSpeechString:forRequest:"), synthesizer, string_, request)
	}

