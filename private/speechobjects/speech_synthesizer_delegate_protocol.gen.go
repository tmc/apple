// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSSpeechSynthesizerDelegate protocol.
type NSSpeechSynthesizerDelegate interface {
	objectivec.IObject
}

// NSSpeechSynthesizerDelegateObject wraps an existing Objective-C object that conforms to the NSSpeechSynthesizerDelegate protocol.
type NSSpeechSynthesizerDelegateObject struct {
	objectivec.Object
}

func (o NSSpeechSynthesizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSpeechSynthesizerDelegateObjectFromID constructs a [NSSpeechSynthesizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSpeechSynthesizerDelegateObjectFromID(id objc.ID) NSSpeechSynthesizerDelegateObject {
	return NSSpeechSynthesizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(synthesizer objectivec.IObject, index uint64, string_ objectivec.IObject, message objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterErrorAtIndex:ofString:message:"), synthesizer, index, string_, message)
}
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterSyncMessage(synthesizer objectivec.IObject, message objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterSyncMessage:"), synthesizer, message)
}
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), synthesizer, speaking)
}
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakPhoneme(synthesizer objectivec.IObject, phoneme int16) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakPhoneme:"), synthesizer, phoneme)
}
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakWordOfString(synthesizer objectivec.IObject, word foundation.NSRange, string_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakWord:ofString:"), synthesizer, word, string_)
}
