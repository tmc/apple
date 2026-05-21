// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// TTSSpeechSynthesizerDelegate protocol.
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

func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didContinueSpeakingRequest:"), synthesizer, request)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterMarker:forRequest:"), synthesizer, marker, request)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, spoken objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:phonemesSpoken:withError:"), synthesizer, request, successfully, spoken, error_)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:withError:"), synthesizer, request, successfully, error_)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didPauseSpeakingRequest:"), synthesizer, request)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidStartSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didStartSpeakingRequest:"), synthesizer, request)
}
func (o TTSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakRangeOfSpeechStringForRequest(synthesizer objectivec.IObject, string_ foundation.NSRange, request objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakRangeOfSpeechString:forRequest:"), synthesizer, string_, request)
}
