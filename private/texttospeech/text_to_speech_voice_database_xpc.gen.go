// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceDatabaseXPC] class.
var (
	_TextToSpeechVoiceDatabaseXPCClass     TextToSpeechVoiceDatabaseXPCClass
	_TextToSpeechVoiceDatabaseXPCClassOnce sync.Once
)

func getTextToSpeechVoiceDatabaseXPCClass() TextToSpeechVoiceDatabaseXPCClass {
	_TextToSpeechVoiceDatabaseXPCClassOnce.Do(func() {
		_TextToSpeechVoiceDatabaseXPCClass = TextToSpeechVoiceDatabaseXPCClass{class: objc.GetClass("TextToSpeech.VoiceDatabaseXPC")}
	})
	return _TextToSpeechVoiceDatabaseXPCClass
}

// GetTextToSpeechVoiceDatabaseXPCClass returns the class object for TextToSpeech.VoiceDatabaseXPC.
func GetTextToSpeechVoiceDatabaseXPCClass() TextToSpeechVoiceDatabaseXPCClass {
	return getTextToSpeechVoiceDatabaseXPCClass()
}

type TextToSpeechVoiceDatabaseXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceDatabaseXPCClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceDatabaseXPCClass) Alloc() TextToSpeechVoiceDatabaseXPC {
	rv := objc.Send[TextToSpeechVoiceDatabaseXPC](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseXPC
type TextToSpeechVoiceDatabaseXPC struct {
	objectivec.Object
}

// TextToSpeechVoiceDatabaseXPCFromID constructs a [TextToSpeechVoiceDatabaseXPC] from an objc.ID.
func TextToSpeechVoiceDatabaseXPCFromID(id objc.ID) TextToSpeechVoiceDatabaseXPC {
	return TextToSpeechVoiceDatabaseXPC{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechVoiceDatabaseXPC struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechVoiceDatabaseXPC embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechVoiceDatabaseXPC] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseXPC
type ITextToSpeechVoiceDatabaseXPC interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceDatabaseXPC) Init() TextToSpeechVoiceDatabaseXPC {
	rv := objc.Send[TextToSpeechVoiceDatabaseXPC](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceDatabaseXPC) Autorelease() TextToSpeechVoiceDatabaseXPC {
	rv := objc.Send[TextToSpeechVoiceDatabaseXPC](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceDatabaseXPC creates a new TextToSpeechVoiceDatabaseXPC instance.
func NewTextToSpeechVoiceDatabaseXPC() TextToSpeechVoiceDatabaseXPC {
	class := getTextToSpeechVoiceDatabaseXPCClass()
	rv := objc.Send[TextToSpeechVoiceDatabaseXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

