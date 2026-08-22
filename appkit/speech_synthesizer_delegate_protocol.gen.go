// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSSpeechSynthesizer](<https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate
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

// Sent just before a synthesized word is spoken through the sound output
// device.
//
// sender: An [NSSpeechSynthesizer] object that’s synthesizing text into speech.
//
// characterRange: Word that `sender` is about to speak into the sound output device.
//
// string: Text that is being synthesized by `sender`.
//
// # Discussion
//
// One use of this method might be to visually highlight the word being
// spoken.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate/speechSynthesizer(_:willSpeakWord:of:)
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakWordOfString(sender INSSpeechSynthesizer, characterRange foundation.NSRange, string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakWord:ofString:"), sender, characterRange, objc.String(string_))
}

// Sent just before a synthesized phoneme is spoken through the sound output
// device.
//
// sender: An [NSSpeechSynthesizer] object that’s synthesizing text into speech.
//
// phonemeOpcode: Phoneme that `sender` is about to speak into the sound output device.
//
// # Discussion
//
// One use of this method might be to animate a mouth on screen to match the
// generated speech.
//
// # Special Considerations
//
// This method is not sent for modern voices. It is only supported for
// MacinTalk voices.
//
// In OS X v10.4 and earlier, the delegate is not sent this message when the
// [NSSpeechSynthesizer] object is synthesizing speech to a file
// ([NSSpeechSynthesizer.StartSpeakingStringToURL]).
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate/speechSynthesizer(_:willSpeakPhoneme:)
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakPhoneme(sender INSSpeechSynthesizer, phonemeOpcode int16) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakPhoneme:"), sender, phonemeOpcode)
}

// Sent to the delegate when a speech synthesizer encounters an error in text
// being synthesized.
//
// sender: Speech synthesizer informing its delegate of an error.
//
// characterIndex: Location in text where the receiver encountered the error.
//
// string: Text the receiver was synthesizing when the error occurred.
//
// message: Error message.
//
// # Discussion
//
// The synthesizer sends an error delegate message whenever it encounters a
// syntax error within a command embedded in the string it is processing. This
// can be useful during application debugging, to detect problems with
// commands that you have embedded in strings that your application speaks. It
// can also be useful if your application allows users to embed commands
// within strings. Your application might display an alert indicating that the
// synthesizer encountered a problem in processing an embedded command.
//
// If your application needs information about errors that occurred prior to
// calling your error delegate method, the application (including the error
// delegate method) can call the sender’s
// [NSSpeechSynthesizer.ObjectForPropertyError] method with the [errors]
// constant.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate/speechSynthesizer(_:didEncounterErrorAt:of:message:)
//
// [errors]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/errors
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender INSSpeechSynthesizer, characterIndex uint, string_ string, message string) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterErrorAtIndex:ofString:message:"), sender, characterIndex, objc.String(string_), objc.String(message))
}

// Sent to the delegate when a speech synthesizer encounters a synchronization
// error.
//
// sender: Speech synthesizer informing its delegate of an error.
//
// message: Error message.
//
// # Discussion
//
// The synthesizer calls your synchronization delegate method whenever it
// encounters a synchronization command embedded in a string. You might use
// the synchronization delegate method to provide a callback not ordinarily
// provided.
//
// For example, you might insert synchronization commands at the end of every
// sentence in a string, or you might enter synchronization commands after
// every numeric value in the text.
//
// However, to synchronize your application with phonemes or words, it makes
// more sense to use the built-in phoneme and word delegate methods:
// [SpeechSynthesizerWillSpeakPhoneme] and
// [SpeechSynthesizerWillSpeakWordOfString].
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate/speechSynthesizer(_:didEncounterSyncMessage:)
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterSyncMessage(sender INSSpeechSynthesizer, message string) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didEncounterSyncMessage:"), sender, objc.String(message))
}

// Sent when an [NSSpeechSynthesizer] object finishes speaking through the
// sound output device.
//
// sender: An [NSSpeechSynthesizer] object that has stopped speaking into the sound
// output device.
//
// finishedSpeaking: true when speaking completed normally, false if speaking is stopped
// prematurely for any reason.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizerDelegate/speechSynthesizer(_:didFinishSpeaking:)
func (o NSSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeaking(sender INSSpeechSynthesizer, finishedSpeaking bool) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), sender, finishedSpeaking)
}

// NSSpeechSynthesizerDelegateConfig holds optional typed callbacks for [NSSpeechSynthesizerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizerdelegate
type NSSpeechSynthesizerDelegateConfig struct {

	// Synthesizing Speech
	// WillSpeakPhoneme — Sent just before a synthesized phoneme is spoken through the sound output device.
	WillSpeakPhoneme func(sender NSSpeechSynthesizer, phonemeOpcode int16)
	// DidFinishSpeaking — Sent when an [NSSpeechSynthesizer](<https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer>) object finishes speaking through the sound output device.
	DidFinishSpeaking func(sender NSSpeechSynthesizer, finishedSpeaking bool)
}

// NewNSSpeechSynthesizerDelegate creates an Objective-C object implementing the [NSSpeechSynthesizerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSpeechSynthesizerDelegateObject] satisfies the [NSSpeechSynthesizerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizerdelegate
func NewNSSpeechSynthesizerDelegate(config NSSpeechSynthesizerDelegateConfig) NSSpeechSynthesizerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSpeechSynthesizerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WillSpeakPhoneme != nil {
		fn := config.WillSpeakPhoneme
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:willSpeakPhoneme:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, phonemeOpcode int16) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSSpeechSynthesizerDelegate", "speechSynthesizer:willSpeakPhoneme:")
					}
				}()
				sender := NSSpeechSynthesizerFromID(senderID)
				fn(sender, phonemeOpcode)
				_delegateDone = true
			},
		})
	}

	if config.DidFinishSpeaking != nil {
		fn := config.DidFinishSpeaking
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didFinishSpeaking:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, finishedSpeaking bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSSpeechSynthesizerDelegate", "speechSynthesizer:didFinishSpeaking:")
					}
				}()
				sender := NSSpeechSynthesizerFromID(senderID)
				fn(sender, finishedSpeaking)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSpeechSynthesizerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSpeechSynthesizerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSpeechSynthesizerDelegateObjectFromID(instance)
}
