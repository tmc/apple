// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSSMLServices] class.
var (
	_TextToSpeechSSMLServicesClass     TextToSpeechSSMLServicesClass
	_TextToSpeechSSMLServicesClassOnce sync.Once
)

func getTextToSpeechSSMLServicesClass() TextToSpeechSSMLServicesClass {
	_TextToSpeechSSMLServicesClassOnce.Do(func() {
		_TextToSpeechSSMLServicesClass = TextToSpeechSSMLServicesClass{class: objc.GetClass("TextToSpeech.SSMLServices")}
	})
	return _TextToSpeechSSMLServicesClass
}

// GetTextToSpeechSSMLServicesClass returns the class object for TextToSpeech.SSMLServices.
func GetTextToSpeechSSMLServicesClass() TextToSpeechSSMLServicesClass {
	return getTextToSpeechSSMLServicesClass()
}

type TextToSpeechSSMLServicesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSSMLServicesClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSMLServicesClass) Alloc() TextToSpeechSSMLServices {
	rv := objc.Send[TextToSpeechSSMLServices](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TextToSpeechSSMLServices.MakeProsodySnippetWithStringRatePitchVolume]
//   - [TextToSpeechSSMLServices.ParseSSMLToPlainTextError]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLServices
type TextToSpeechSSMLServices struct {
	objectivec.Object
}

// TextToSpeechSSMLServicesFromID constructs a [TextToSpeechSSMLServices] from an objc.ID.
func TextToSpeechSSMLServicesFromID(id objc.ID) TextToSpeechSSMLServices {
	return TextToSpeechSSMLServices{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechSSMLServices implements ITextToSpeechSSMLServices.
var _ ITextToSpeechSSMLServices = TextToSpeechSSMLServices{}

// An interface definition for the [TextToSpeechSSMLServices] class.
//
// # Methods
//
//   - [ITextToSpeechSSMLServices.MakeProsodySnippetWithStringRatePitchVolume]
//   - [ITextToSpeechSSMLServices.ParseSSMLToPlainTextError]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLServices
type ITextToSpeechSSMLServices interface {
	objectivec.IObject

	// Topic: Methods

	MakeProsodySnippetWithStringRatePitchVolume(string_ objectivec.IObject, rate objectivec.IObject, pitch objectivec.IObject, volume objectivec.IObject) objectivec.IObject
	ParseSSMLToPlainTextError(text objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (t TextToSpeechSSMLServices) Init() TextToSpeechSSMLServices {
	rv := objc.Send[TextToSpeechSSMLServices](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSMLServices) Autorelease() TextToSpeechSSMLServices {
	rv := objc.Send[TextToSpeechSSMLServices](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSMLServices creates a new TextToSpeechSSMLServices instance.
func NewTextToSpeechSSMLServices() TextToSpeechSSMLServices {
	class := getTextToSpeechSSMLServicesClass()
	rv := objc.Send[TextToSpeechSSMLServices](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLServices/makeProsodySnippetWithString:rate:pitch:volume:
func (t TextToSpeechSSMLServices) MakeProsodySnippetWithStringRatePitchVolume(string_ objectivec.IObject, rate objectivec.IObject, pitch objectivec.IObject, volume objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("makeProsodySnippetWithString:rate:pitch:volume:"), string_, rate, pitch, volume)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLServices/parseSSMLToPlainText:error:
func (t TextToSpeechSSMLServices) ParseSSMLToPlainTextError(text objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("parseSSMLToPlainText:error:"), text, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLServices/setShared:
func (_TextToSpeechSSMLServicesClass TextToSpeechSSMLServicesClass) SetShared(shared objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TextToSpeechSSMLServicesClass.class), objc.Sel("setShared:"), shared)
}
