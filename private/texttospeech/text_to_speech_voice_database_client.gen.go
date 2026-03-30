// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceDatabaseClient] class.
var (
	_TextToSpeechVoiceDatabaseClientClass     TextToSpeechVoiceDatabaseClientClass
	_TextToSpeechVoiceDatabaseClientClassOnce sync.Once
)

func getTextToSpeechVoiceDatabaseClientClass() TextToSpeechVoiceDatabaseClientClass {
	_TextToSpeechVoiceDatabaseClientClassOnce.Do(func() {
		_TextToSpeechVoiceDatabaseClientClass = TextToSpeechVoiceDatabaseClientClass{class: objc.GetClass("TextToSpeech.VoiceDatabaseClient")}
	})
	return _TextToSpeechVoiceDatabaseClientClass
}

// GetTextToSpeechVoiceDatabaseClientClass returns the class object for TextToSpeech.VoiceDatabaseClient.
func GetTextToSpeechVoiceDatabaseClientClass() TextToSpeechVoiceDatabaseClientClass {
	return getTextToSpeechVoiceDatabaseClientClass()
}

type TextToSpeechVoiceDatabaseClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceDatabaseClientClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceDatabaseClientClass) Alloc() TextToSpeechVoiceDatabaseClient {
	rv := objc.Send[TextToSpeechVoiceDatabaseClient](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseClient
type TextToSpeechVoiceDatabaseClient struct {
	objectivec.Object
}

// TextToSpeechVoiceDatabaseClientFromID constructs a [TextToSpeechVoiceDatabaseClient] from an objc.ID.
func TextToSpeechVoiceDatabaseClientFromID(id objc.ID) TextToSpeechVoiceDatabaseClient {
	return TextToSpeechVoiceDatabaseClient{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechVoiceDatabaseClient struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechVoiceDatabaseClient embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechVoiceDatabaseClient] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseClient
type ITextToSpeechVoiceDatabaseClient interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceDatabaseClient) Init() TextToSpeechVoiceDatabaseClient {
	rv := objc.Send[TextToSpeechVoiceDatabaseClient](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceDatabaseClient) Autorelease() TextToSpeechVoiceDatabaseClient {
	rv := objc.Send[TextToSpeechVoiceDatabaseClient](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceDatabaseClient creates a new TextToSpeechVoiceDatabaseClient instance.
func NewTextToSpeechVoiceDatabaseClient() TextToSpeechVoiceDatabaseClient {
	class := getTextToSpeechVoiceDatabaseClientClass()
	rv := objc.Send[TextToSpeechVoiceDatabaseClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}
