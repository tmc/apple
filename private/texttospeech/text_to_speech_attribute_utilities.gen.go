// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAttributeUtilities] class.
var (
	_TextToSpeechAttributeUtilitiesClass     TextToSpeechAttributeUtilitiesClass
	_TextToSpeechAttributeUtilitiesClassOnce sync.Once
)

func getTextToSpeechAttributeUtilitiesClass() TextToSpeechAttributeUtilitiesClass {
	_TextToSpeechAttributeUtilitiesClassOnce.Do(func() {
		_TextToSpeechAttributeUtilitiesClass = TextToSpeechAttributeUtilitiesClass{class: objc.GetClass("TextToSpeech.AttributeUtilities")}
	})
	return _TextToSpeechAttributeUtilitiesClass
}

// GetTextToSpeechAttributeUtilitiesClass returns the class object for TextToSpeech.AttributeUtilities.
func GetTextToSpeechAttributeUtilitiesClass() TextToSpeechAttributeUtilitiesClass {
	return getTextToSpeechAttributeUtilitiesClass()
}

type TextToSpeechAttributeUtilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechAttributeUtilitiesClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAttributeUtilitiesClass) Alloc() TextToSpeechAttributeUtilities {
	rv := objc.SendIfResponds[TextToSpeechAttributeUtilities](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechAttributeUtilities struct {
	objectivec.Object
}

// TextToSpeechAttributeUtilitiesFromID constructs a [TextToSpeechAttributeUtilities] from an objc.ID.
func TextToSpeechAttributeUtilitiesFromID(id objc.ID) TextToSpeechAttributeUtilities {
	return TextToSpeechAttributeUtilities{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechAttributeUtilities implements ITextToSpeechAttributeUtilities.
var _ ITextToSpeechAttributeUtilities = TextToSpeechAttributeUtilities{}

// An interface definition for the [TextToSpeechAttributeUtilities] class.
type ITextToSpeechAttributeUtilities interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAttributeUtilities) Init() TextToSpeechAttributeUtilities {
	rv := objc.SendIfResponds[TextToSpeechAttributeUtilities](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAttributeUtilities) Autorelease() TextToSpeechAttributeUtilities {
	rv := objc.SendIfResponds[TextToSpeechAttributeUtilities](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAttributeUtilities creates a new TextToSpeechAttributeUtilities instance.
func NewTextToSpeechAttributeUtilities() TextToSpeechAttributeUtilities {
	class := getTextToSpeechAttributeUtilitiesClass()
	rv := objc.SendIfResponds[TextToSpeechAttributeUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}
