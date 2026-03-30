// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAXCatalogUpdater] class.
var (
	_TextToSpeechAXCatalogUpdaterClass     TextToSpeechAXCatalogUpdaterClass
	_TextToSpeechAXCatalogUpdaterClassOnce sync.Once
)

func getTextToSpeechAXCatalogUpdaterClass() TextToSpeechAXCatalogUpdaterClass {
	_TextToSpeechAXCatalogUpdaterClassOnce.Do(func() {
		_TextToSpeechAXCatalogUpdaterClass = TextToSpeechAXCatalogUpdaterClass{class: objc.GetClass("TextToSpeech.AXCatalogUpdater")}
	})
	return _TextToSpeechAXCatalogUpdaterClass
}

// GetTextToSpeechAXCatalogUpdaterClass returns the class object for TextToSpeech.AXCatalogUpdater.
func GetTextToSpeechAXCatalogUpdaterClass() TextToSpeechAXCatalogUpdaterClass {
	return getTextToSpeechAXCatalogUpdaterClass()
}

type TextToSpeechAXCatalogUpdaterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechAXCatalogUpdaterClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAXCatalogUpdaterClass) Alloc() TextToSpeechAXCatalogUpdater {
	rv := objc.Send[TextToSpeechAXCatalogUpdater](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXCatalogUpdater
type TextToSpeechAXCatalogUpdater struct {
	objectivec.Object
}

// TextToSpeechAXCatalogUpdaterFromID constructs a [TextToSpeechAXCatalogUpdater] from an objc.ID.
func TextToSpeechAXCatalogUpdaterFromID(id objc.ID) TextToSpeechAXCatalogUpdater {
	return TextToSpeechAXCatalogUpdater{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechAXCatalogUpdater struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechAXCatalogUpdater embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechAXCatalogUpdater] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXCatalogUpdater
type ITextToSpeechAXCatalogUpdater interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAXCatalogUpdater) Init() TextToSpeechAXCatalogUpdater {
	rv := objc.Send[TextToSpeechAXCatalogUpdater](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAXCatalogUpdater) Autorelease() TextToSpeechAXCatalogUpdater {
	rv := objc.Send[TextToSpeechAXCatalogUpdater](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAXCatalogUpdater creates a new TextToSpeechAXCatalogUpdater instance.
func NewTextToSpeechAXCatalogUpdater() TextToSpeechAXCatalogUpdater {
	class := getTextToSpeechAXCatalogUpdaterClass()
	rv := objc.Send[TextToSpeechAXCatalogUpdater](objc.ID(class.class), objc.Sel("new"))
	return rv
}
