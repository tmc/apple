// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that you adopt in your objects to track the availability of a speech recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerDelegate
type SFSpeechRecognizerDelegate interface {
	objectivec.IObject
}

// SFSpeechRecognizerDelegateObject wraps an existing Objective-C object that conforms to the SFSpeechRecognizerDelegate protocol.
type SFSpeechRecognizerDelegateObject struct {
	objectivec.Object
}

func (o SFSpeechRecognizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SFSpeechRecognizerDelegateObjectFromID constructs a [SFSpeechRecognizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SFSpeechRecognizerDelegateObjectFromID(id objc.ID) SFSpeechRecognizerDelegateObject {
	return SFSpeechRecognizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the availability of its associated speech
// recognizer changed.
//
// speechRecognizer: The [SFSpeechRecognizer] object whose availability changed.
//
// available: A Boolean value that indicates the new availability of the speech
// recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerDelegate/speechRecognizer(_:availabilityDidChange:)
func (o SFSpeechRecognizerDelegateObject) SpeechRecognizerAvailabilityDidChange(speechRecognizer ISFSpeechRecognizer, available bool) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognizer:availabilityDidChange:"), speechRecognizer, available)
}

// SFSpeechRecognizerDelegateConfig holds optional typed callbacks for [SFSpeechRecognizerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/speech/sfspeechrecognizerdelegate
type SFSpeechRecognizerDelegateConfig struct {

	// Monitoring speech recognizer availability
	// SpeechRecognizerAvailabilityDidChange — Tells the delegate that the availability of its associated speech recognizer changed.
	SpeechRecognizerAvailabilityDidChange func(speechRecognizer SFSpeechRecognizer, available bool)
}

// NewSFSpeechRecognizerDelegate creates an Objective-C object implementing the [SFSpeechRecognizerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SFSpeechRecognizerDelegateObject] satisfies the [SFSpeechRecognizerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/speech/sfspeechrecognizerdelegate
func NewSFSpeechRecognizerDelegate(config SFSpeechRecognizerDelegateConfig) SFSpeechRecognizerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSFSpeechRecognizerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SpeechRecognizerAvailabilityDidChange != nil {
		fn := config.SpeechRecognizerAvailabilityDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognizer:availabilityDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, speechRecognizerID objc.ID, available bool) {
				speechRecognizer := SFSpeechRecognizerFromID(speechRecognizerID)
				fn(speechRecognizer, available)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SFSpeechRecognizerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSFSpeechRecognizerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SFSpeechRecognizerDelegateObjectFromID(instance)
}
