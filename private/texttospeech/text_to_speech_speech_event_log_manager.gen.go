// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSpeechEventLogManager] class.
var (
	_TextToSpeechSpeechEventLogManagerClass     TextToSpeechSpeechEventLogManagerClass
	_TextToSpeechSpeechEventLogManagerClassOnce sync.Once
)

func getTextToSpeechSpeechEventLogManagerClass() TextToSpeechSpeechEventLogManagerClass {
	_TextToSpeechSpeechEventLogManagerClassOnce.Do(func() {
		_TextToSpeechSpeechEventLogManagerClass = TextToSpeechSpeechEventLogManagerClass{class: objc.GetClass("TextToSpeech.SpeechEventLogManager")}
	})
	return _TextToSpeechSpeechEventLogManagerClass
}

// GetTextToSpeechSpeechEventLogManagerClass returns the class object for TextToSpeech.SpeechEventLogManager.
func GetTextToSpeechSpeechEventLogManagerClass() TextToSpeechSpeechEventLogManagerClass {
	return getTextToSpeechSpeechEventLogManagerClass()
}

type TextToSpeechSpeechEventLogManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSpeechEventLogManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSpeechEventLogManagerClass) Alloc() TextToSpeechSpeechEventLogManager {
	rv := objc.SendIfResponds[TextToSpeechSpeechEventLogManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechSpeechEventLogManager struct {
	objectivec.Object
}

// TextToSpeechSpeechEventLogManagerFromID constructs a [TextToSpeechSpeechEventLogManager] from an objc.ID.
func TextToSpeechSpeechEventLogManagerFromID(id objc.ID) TextToSpeechSpeechEventLogManager {
	return TextToSpeechSpeechEventLogManager{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechSpeechEventLogManager implements ITextToSpeechSpeechEventLogManager.
var _ ITextToSpeechSpeechEventLogManager = TextToSpeechSpeechEventLogManager{}

// An interface definition for the [TextToSpeechSpeechEventLogManager] class.
type ITextToSpeechSpeechEventLogManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSpeechEventLogManager) Init() TextToSpeechSpeechEventLogManager {
	rv := objc.SendIfResponds[TextToSpeechSpeechEventLogManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSpeechEventLogManager) Autorelease() TextToSpeechSpeechEventLogManager {
	rv := objc.SendIfResponds[TextToSpeechSpeechEventLogManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSpeechEventLogManager creates a new TextToSpeechSpeechEventLogManager instance.
func NewTextToSpeechSpeechEventLogManager() TextToSpeechSpeechEventLogManager {
	class := getTextToSpeechSpeechEventLogManagerClass()
	rv := objc.SendIfResponds[TextToSpeechSpeechEventLogManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
