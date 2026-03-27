// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSSELoaderManager] class.
var (
	_TextToSpeechSSELoaderManagerClass     TextToSpeechSSELoaderManagerClass
	_TextToSpeechSSELoaderManagerClassOnce sync.Once
)

func getTextToSpeechSSELoaderManagerClass() TextToSpeechSSELoaderManagerClass {
	_TextToSpeechSSELoaderManagerClassOnce.Do(func() {
		_TextToSpeechSSELoaderManagerClass = TextToSpeechSSELoaderManagerClass{class: objc.GetClass("TextToSpeech.SSELoaderManager")}
	})
	return _TextToSpeechSSELoaderManagerClass
}

// GetTextToSpeechSSELoaderManagerClass returns the class object for TextToSpeech.SSELoaderManager.
func GetTextToSpeechSSELoaderManagerClass() TextToSpeechSSELoaderManagerClass {
	return getTextToSpeechSSELoaderManagerClass()
}

type TextToSpeechSSELoaderManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSELoaderManagerClass) Alloc() TextToSpeechSSELoaderManager {
	rv := objc.Send[TextToSpeechSSELoaderManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSELoaderManager
type TextToSpeechSSELoaderManager struct {
	objectivec.Object
}

// TextToSpeechSSELoaderManagerFromID constructs a [TextToSpeechSSELoaderManager] from an objc.ID.
func TextToSpeechSSELoaderManagerFromID(id objc.ID) TextToSpeechSSELoaderManager {
	return TextToSpeechSSELoaderManager{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechSSELoaderManager implements ITextToSpeechSSELoaderManager.
var _ ITextToSpeechSSELoaderManager = TextToSpeechSSELoaderManager{}

// An interface definition for the [TextToSpeechSSELoaderManager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSELoaderManager
type ITextToSpeechSSELoaderManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSSELoaderManager) Init() TextToSpeechSSELoaderManager {
	rv := objc.Send[TextToSpeechSSELoaderManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSELoaderManager) Autorelease() TextToSpeechSSELoaderManager {
	rv := objc.Send[TextToSpeechSSELoaderManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSELoaderManager creates a new TextToSpeechSSELoaderManager instance.
func NewTextToSpeechSSELoaderManager() TextToSpeechSSELoaderManager {
	class := getTextToSpeechSSELoaderManagerClass()
	rv := objc.Send[TextToSpeechSSELoaderManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

