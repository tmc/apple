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
	rv := objc.SendIfResponds[TextToSpeechVoiceDatabaseClient](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechVoiceDatabaseClient struct {
	objectivec.Object
}

// TextToSpeechVoiceDatabaseClientFromID constructs a [TextToSpeechVoiceDatabaseClient] from an objc.ID.
func TextToSpeechVoiceDatabaseClientFromID(id objc.ID) TextToSpeechVoiceDatabaseClient {
	return TextToSpeechVoiceDatabaseClient{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechVoiceDatabaseClient implements ITextToSpeechVoiceDatabaseClient.
var _ ITextToSpeechVoiceDatabaseClient = TextToSpeechVoiceDatabaseClient{}

// An interface definition for the [TextToSpeechVoiceDatabaseClient] class.
type ITextToSpeechVoiceDatabaseClient interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceDatabaseClient) Init() TextToSpeechVoiceDatabaseClient {
	rv := objc.SendIfResponds[TextToSpeechVoiceDatabaseClient](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceDatabaseClient) Autorelease() TextToSpeechVoiceDatabaseClient {
	rv := objc.SendIfResponds[TextToSpeechVoiceDatabaseClient](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceDatabaseClient creates a new TextToSpeechVoiceDatabaseClient instance.
func NewTextToSpeechVoiceDatabaseClient() TextToSpeechVoiceDatabaseClient {
	class := getTextToSpeechVoiceDatabaseClientClass()
	rv := objc.SendIfResponds[TextToSpeechVoiceDatabaseClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}
