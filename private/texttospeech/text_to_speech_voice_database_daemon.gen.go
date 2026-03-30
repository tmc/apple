// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceDatabaseDaemon] class.
var (
	_TextToSpeechVoiceDatabaseDaemonClass     TextToSpeechVoiceDatabaseDaemonClass
	_TextToSpeechVoiceDatabaseDaemonClassOnce sync.Once
)

func getTextToSpeechVoiceDatabaseDaemonClass() TextToSpeechVoiceDatabaseDaemonClass {
	_TextToSpeechVoiceDatabaseDaemonClassOnce.Do(func() {
		_TextToSpeechVoiceDatabaseDaemonClass = TextToSpeechVoiceDatabaseDaemonClass{class: objc.GetClass("TextToSpeech.VoiceDatabaseDaemon")}
	})
	return _TextToSpeechVoiceDatabaseDaemonClass
}

// GetTextToSpeechVoiceDatabaseDaemonClass returns the class object for TextToSpeech.VoiceDatabaseDaemon.
func GetTextToSpeechVoiceDatabaseDaemonClass() TextToSpeechVoiceDatabaseDaemonClass {
	return getTextToSpeechVoiceDatabaseDaemonClass()
}

type TextToSpeechVoiceDatabaseDaemonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceDatabaseDaemonClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceDatabaseDaemonClass) Alloc() TextToSpeechVoiceDatabaseDaemon {
	rv := objc.Send[TextToSpeechVoiceDatabaseDaemon](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseDaemon
type TextToSpeechVoiceDatabaseDaemon struct {
	objectivec.Object
}

// TextToSpeechVoiceDatabaseDaemonFromID constructs a [TextToSpeechVoiceDatabaseDaemon] from an objc.ID.
func TextToSpeechVoiceDatabaseDaemonFromID(id objc.ID) TextToSpeechVoiceDatabaseDaemon {
	return TextToSpeechVoiceDatabaseDaemon{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechVoiceDatabaseDaemon struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechVoiceDatabaseDaemon embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechVoiceDatabaseDaemon] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceDatabaseDaemon
type ITextToSpeechVoiceDatabaseDaemon interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceDatabaseDaemon) Init() TextToSpeechVoiceDatabaseDaemon {
	rv := objc.Send[TextToSpeechVoiceDatabaseDaemon](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceDatabaseDaemon) Autorelease() TextToSpeechVoiceDatabaseDaemon {
	rv := objc.Send[TextToSpeechVoiceDatabaseDaemon](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceDatabaseDaemon creates a new TextToSpeechVoiceDatabaseDaemon instance.
func NewTextToSpeechVoiceDatabaseDaemon() TextToSpeechVoiceDatabaseDaemon {
	class := getTextToSpeechVoiceDatabaseDaemonClass()
	rv := objc.Send[TextToSpeechVoiceDatabaseDaemon](objc.ID(class.class), objc.Sel("new"))
	return rv
}
