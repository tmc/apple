// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// TTSAUMessaging protocol.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging
type TTSAUMessaging interface {
	objectivec.IObject
}

// TTSAUMessagingObject wraps an existing Objective-C object that conforms to the TTSAUMessaging protocol.
type TTSAUMessagingObject struct {
	objectivec.Object
}

func (o TTSAUMessagingObject) BaseObject() objectivec.Object {
	return o.Object
}

// TTSAUMessagingObjectFromID constructs a [TTSAUMessagingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TTSAUMessagingObjectFromID(id objc.ID) TTSAUMessagingObject {
	return TTSAUMessagingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging/defaultSettingsForVoice:
func (o TTSAUMessagingObject) DefaultSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("defaultSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging/echo:
func (o TTSAUMessagingObject) Echo(echo objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("echo:"), echo)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging/prewarmWithVoice:
func (o TTSAUMessagingObject) PrewarmWithVoice(voice objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("prewarmWithVoice:"), voice)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging/requireFirstUnlockForVoiceLoad
func (o TTSAUMessagingObject) RequireFirstUnlockForVoiceLoad() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("requireFirstUnlockForVoiceLoad"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessaging/voicesExternallyManaged
func (o TTSAUMessagingObject) VoicesExternallyManaged() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("voicesExternallyManaged"))
	return objectivec.Object{ID: rv}
}
