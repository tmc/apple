// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A delegate protocol that contains optional methods you can implement to respond to events that occur during speech synthesis.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate
type AVSpeechSynthesizerDelegate interface {
	objectivec.IObject
}

// AVSpeechSynthesizerDelegateObject wraps an existing Objective-C object that conforms to the AVSpeechSynthesizerDelegate protocol.
type AVSpeechSynthesizerDelegateObject struct {
	objectivec.Object
}

func (o AVSpeechSynthesizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVSpeechSynthesizerDelegateObjectFromID constructs a [AVSpeechSynthesizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVSpeechSynthesizerDelegateObjectFromID(id objc.ID) AVSpeechSynthesizerDelegateObject {
	return AVSpeechSynthesizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when the synthesizer begins speaking an utterance.
//
// synthesizer: The speech synthesizer that starts speaking the utterance.
//
// utterance: The utterance that the speech synthesizer starts speaking.
//
// # Discussion
//
// If the utterance’s [PreUtteranceDelay] property is greater than zero, the
// system calls this method after the delay completes and speech begins.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:didStart:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerDidStartSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didStartSpeechUtterance:"), synthesizer, utterance)
}

// Tells the delegate when the synthesizer is about to speak a portion of an
// utterance’s text.
//
// synthesizer: The speech synthesizer that’s about to speak an utterance.
//
// characterRange: The range of characters in the utterance’s [SpeechString] that correspond
// to the unit of speech the synthesizer is about to speak.
//
// utterance: The utterance that the speech synthesizer is about to speak.
//
// # Discussion
//
// The system calls this method once for each unit of speech in the
// utterance’s text, which is generally a word.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:willSpeakRangeOfSpeechString:utterance:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance(synthesizer IAVSpeechSynthesizer, characterRange foundation.NSRange, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakRangeOfSpeechString:utterance:"), synthesizer, characterRange, utterance)
}

// Tells the delegate when the synthesizer is about to speak a marker of an
// utterance’s text.
//
// synthesizer: The speech synthesizer that’s about to speak a marker of an utterance.
//
// marker: The synthesized audio that the speech synthesizer is about to speak.
//
// utterance: The utterance that the speech synthesizer pauses speaking.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:willSpeak:utterance:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakMarkerUtterance(synthesizer IAVSpeechSynthesizer, marker IAVSpeechSynthesisMarker, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:willSpeakMarker:utterance:"), synthesizer, marker, utterance)
}

// Tells the delegate when the synthesizer pauses while speaking an utterance.
//
// synthesizer: The speech synthesizer that pauses speaking the utterance.
//
// utterance: The utterance that the speech synthesizer pauses speaking.
//
// # Discussion
//
// The system only calls this method if a speech synthesizer is speaking an
// utterance and the system calls its [PauseSpeakingAtBoundary] method. The
// system doesn’t call this method if the synthesizer is in a delay between
// utterances when speech pauses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:didPause:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerDidPauseSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didPauseSpeechUtterance:"), synthesizer, utterance)
}

// Tells the delegate when the synthesizer resumes speaking an utterance after
// pausing.
//
// synthesizer: The speech synthesizer that resumes speaking the utterance.
//
// utterance: The utterance that the speech synthesizer resumes speaking.
//
// # Discussion
//
// The system only calls this method if a speech synthesizer pauses speaking
// and the system calls its [PauseSpeakingAtBoundary] method. The system
// doesn’t call this method if the synthesizer pauses while in a delay
// between utterances.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:didContinue:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerDidContinueSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didContinueSpeechUtterance:"), synthesizer, utterance)
}

// Tells the delegate when the synthesizer finishes speaking an utterance.
//
// synthesizer: The speech synthesizer that finishes speaking the utterance.
//
// utterance: The utterance that the speech synthesizer finishes speaking.
//
// # Discussion
//
// The system ignores the final utterance’s [PostUtteranceDelay] and calls
// this method immediately when speech ends.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:didFinish:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didFinishSpeechUtterance:"), synthesizer, utterance)
}

// Tells the delegate when the synthesizer cancels speaking an utterance.
//
// synthesizer: The speech synthesizer that cancels speaking the utterance.
//
// utterance: The utterance that the speech synthesizer cancels speaking.
//
// # Discussion
//
// The system only calls this method if a speech synthesizer is speaking an
// utterance and the system calls its [StopSpeakingAtBoundary] method. The
// system doesn’t call this method if the synthesizer is in a delay between
// utterances when speech stops, and it doesn’t call it for unspoken
// utterances.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizerDelegate/speechSynthesizer(_:didCancel:)
func (o AVSpeechSynthesizerDelegateObject) SpeechSynthesizerDidCancelSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	objc.Send[struct{}](o.ID, objc.Sel("speechSynthesizer:didCancelSpeechUtterance:"), synthesizer, utterance)
}

// AVSpeechSynthesizerDelegateConfig holds optional typed callbacks for [AVSpeechSynthesizerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizerdelegate
type AVSpeechSynthesizerDelegateConfig struct {

	// Responding to speech synthesis events
	// SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance — Tells the delegate when the synthesizer is about to speak a portion of an utterance’s text.
	SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance func(synthesizer AVSpeechSynthesizer, characterRange foundation.NSRange, utterance AVSpeechUtterance)

	// Other Methods
	// SpeechSynthesizerDidStartSpeechUtterance — Tells the delegate when the synthesizer begins speaking an utterance.
	SpeechSynthesizerDidStartSpeechUtterance func(synthesizer AVSpeechSynthesizer, utterance AVSpeechUtterance)
	// SpeechSynthesizerWillSpeakMarkerUtterance — Tells the delegate when the synthesizer is about to speak a marker of an utterance’s text.
	SpeechSynthesizerWillSpeakMarkerUtterance func(synthesizer AVSpeechSynthesizer, marker AVSpeechSynthesisMarker, utterance AVSpeechUtterance)
	// SpeechSynthesizerDidPauseSpeechUtterance — Tells the delegate when the synthesizer pauses while speaking an utterance.
	SpeechSynthesizerDidPauseSpeechUtterance func(synthesizer AVSpeechSynthesizer, utterance AVSpeechUtterance)
	// SpeechSynthesizerDidContinueSpeechUtterance — Tells the delegate when the synthesizer resumes speaking an utterance after pausing.
	SpeechSynthesizerDidContinueSpeechUtterance func(synthesizer AVSpeechSynthesizer, utterance AVSpeechUtterance)
	// SpeechSynthesizerDidFinishSpeechUtterance — Tells the delegate when the synthesizer finishes speaking an utterance.
	SpeechSynthesizerDidFinishSpeechUtterance func(synthesizer AVSpeechSynthesizer, utterance AVSpeechUtterance)
	// SpeechSynthesizerDidCancelSpeechUtterance — Tells the delegate when the synthesizer cancels speaking an utterance.
	SpeechSynthesizerDidCancelSpeechUtterance func(synthesizer AVSpeechSynthesizer, utterance AVSpeechUtterance)
}

// NewAVSpeechSynthesizerDelegate creates an Objective-C object implementing the [AVSpeechSynthesizerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVSpeechSynthesizerDelegateObject] satisfies the [AVSpeechSynthesizerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesizerdelegate
func NewAVSpeechSynthesizerDelegate(config AVSpeechSynthesizerDelegateConfig) AVSpeechSynthesizerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVSpeechSynthesizerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SpeechSynthesizerDidStartSpeechUtterance != nil {
		fn := config.SpeechSynthesizerDidStartSpeechUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didStartSpeechUtterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, utterance)
			},
		})
	}

	if config.SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance != nil {
		fn := config.SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:willSpeakRangeOfSpeechString:utterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, characterRange foundation.NSRange, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, characterRange, utterance)
			},
		})
	}

	if config.SpeechSynthesizerWillSpeakMarkerUtterance != nil {
		fn := config.SpeechSynthesizerWillSpeakMarkerUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:willSpeakMarker:utterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, markerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				marker := AVSpeechSynthesisMarkerFromID(markerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, marker, utterance)
			},
		})
	}

	if config.SpeechSynthesizerDidPauseSpeechUtterance != nil {
		fn := config.SpeechSynthesizerDidPauseSpeechUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didPauseSpeechUtterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, utterance)
			},
		})
	}

	if config.SpeechSynthesizerDidContinueSpeechUtterance != nil {
		fn := config.SpeechSynthesizerDidContinueSpeechUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didContinueSpeechUtterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, utterance)
			},
		})
	}

	if config.SpeechSynthesizerDidFinishSpeechUtterance != nil {
		fn := config.SpeechSynthesizerDidFinishSpeechUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didFinishSpeechUtterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, utterance)
			},
		})
	}

	if config.SpeechSynthesizerDidCancelSpeechUtterance != nil {
		fn := config.SpeechSynthesizerDidCancelSpeechUtterance
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didCancelSpeechUtterance:"),
			Fn: func(self objc.ID, _cmd objc.SEL, synthesizerID objc.ID, utteranceID objc.ID) {
				synthesizer := AVSpeechSynthesizerFromID(synthesizerID)
				utterance := AVSpeechUtteranceFromID(utteranceID)
				fn(synthesizer, utterance)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVSpeechSynthesizerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVSpeechSynthesizerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVSpeechSynthesizerDelegateObjectFromID(instance)
}
