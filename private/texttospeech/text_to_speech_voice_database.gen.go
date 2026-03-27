// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceDatabase] class.
var (
	_TextToSpeechVoiceDatabaseClass     TextToSpeechVoiceDatabaseClass
	_TextToSpeechVoiceDatabaseClassOnce sync.Once
)

func getTextToSpeechVoiceDatabaseClass() TextToSpeechVoiceDatabaseClass {
	_TextToSpeechVoiceDatabaseClassOnce.Do(func() {
		_TextToSpeechVoiceDatabaseClass = TextToSpeechVoiceDatabaseClass{class: objc.GetClass("TextToSpeech.VoiceDatabase")}
	})
	return _TextToSpeechVoiceDatabaseClass
}

// GetTextToSpeechVoiceDatabaseClass returns the class object for TextToSpeech.VoiceDatabase.
func GetTextToSpeechVoiceDatabaseClass() TextToSpeechVoiceDatabaseClass {
	return getTextToSpeechVoiceDatabaseClass()
}

type TextToSpeechVoiceDatabaseClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceDatabaseClass) Alloc() TextToSpeechVoiceDatabase {
	rv := objc.Send[TextToSpeechVoiceDatabase](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabase
type TextToSpeechVoiceDatabase struct {
	objectivec.Object
}

// TextToSpeechVoiceDatabaseFromID constructs a [TextToSpeechVoiceDatabase] from an objc.ID.
func TextToSpeechVoiceDatabaseFromID(id objc.ID) TextToSpeechVoiceDatabase {
	return TextToSpeechVoiceDatabase{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechVoiceDatabase implements ITextToSpeechVoiceDatabase.
var _ ITextToSpeechVoiceDatabase = TextToSpeechVoiceDatabase{}

// An interface definition for the [TextToSpeechVoiceDatabase] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabase
type ITextToSpeechVoiceDatabase interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceDatabase) Init() TextToSpeechVoiceDatabase {
	rv := objc.Send[TextToSpeechVoiceDatabase](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceDatabase) Autorelease() TextToSpeechVoiceDatabase {
	rv := objc.Send[TextToSpeechVoiceDatabase](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceDatabase creates a new TextToSpeechVoiceDatabase instance.
func NewTextToSpeechVoiceDatabase() TextToSpeechVoiceDatabase {
	class := getTextToSpeechVoiceDatabaseClass()
	rv := objc.Send[TextToSpeechVoiceDatabase](objc.ID(class.class), objc.Sel("new"))
	return rv
}

