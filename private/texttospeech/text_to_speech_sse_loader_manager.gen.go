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

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSSELoaderManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSELoaderManagerClass) Alloc() TextToSpeechSSELoaderManager {
	rv := objc.SendIfResponds[TextToSpeechSSELoaderManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

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
type ITextToSpeechSSELoaderManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSSELoaderManager) Init() TextToSpeechSSELoaderManager {
	rv := objc.SendIfResponds[TextToSpeechSSELoaderManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSELoaderManager) Autorelease() TextToSpeechSSELoaderManager {
	rv := objc.SendIfResponds[TextToSpeechSSELoaderManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSELoaderManager creates a new TextToSpeechSSELoaderManager instance.
func NewTextToSpeechSSELoaderManager() TextToSpeechSSELoaderManager {
	class := getTextToSpeechSSELoaderManagerClass()
	rv := objc.SendIfResponds[TextToSpeechSSELoaderManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
