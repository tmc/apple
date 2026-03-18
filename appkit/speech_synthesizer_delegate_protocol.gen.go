// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSSpeechSynthesizer](<doc://com.apple.appkit/documentation/AppKit/NSSpeechSynthesizer>) objects.
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
	// DidFinishSpeaking — Sent when an [NSSpeechSynthesizer](<doc://com.apple.appkit/documentation/AppKit/NSSpeechSynthesizer>) object finishes speaking through the sound output device.
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
				sender := NSSpeechSynthesizerFromID(senderID)
				fn(sender, phonemeOpcode)
			},
		})
	}

	if config.DidFinishSpeaking != nil {
		fn := config.DidFinishSpeaking
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechSynthesizer:didFinishSpeaking:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, finishedSpeaking bool) {
				sender := NSSpeechSynthesizerFromID(senderID)
				fn(sender, finishedSpeaking)
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

