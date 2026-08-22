// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTaggedSSML] class.
var (
	_TextToSpeechTaggedSSMLClass     TextToSpeechTaggedSSMLClass
	_TextToSpeechTaggedSSMLClassOnce sync.Once
)

func getTextToSpeechTaggedSSMLClass() TextToSpeechTaggedSSMLClass {
	_TextToSpeechTaggedSSMLClassOnce.Do(func() {
		_TextToSpeechTaggedSSMLClass = TextToSpeechTaggedSSMLClass{class: objc.GetClass("TextToSpeech.TaggedSSML")}
	})
	return _TextToSpeechTaggedSSMLClass
}

// GetTextToSpeechTaggedSSMLClass returns the class object for TextToSpeech.TaggedSSML.
func GetTextToSpeechTaggedSSMLClass() TextToSpeechTaggedSSMLClass {
	return getTextToSpeechTaggedSSMLClass()
}

type TextToSpeechTaggedSSMLClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTaggedSSMLClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTaggedSSMLClass) Alloc() TextToSpeechTaggedSSML {
	rv := objc.SendIfResponds[TextToSpeechTaggedSSML](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TextToSpeechTaggedSSML.OriginalSSML]
//   - [TextToSpeechTaggedSSML.SetOriginalSSML]
//   - [TextToSpeechTaggedSSML.SsmlSnippets]
//   - [TextToSpeechTaggedSSML.SetSsmlSnippets]
type TextToSpeechTaggedSSML struct {
	objectivec.Object
}

// TextToSpeechTaggedSSMLFromID constructs a [TextToSpeechTaggedSSML] from an objc.ID.
func TextToSpeechTaggedSSMLFromID(id objc.ID) TextToSpeechTaggedSSML {
	return TextToSpeechTaggedSSML{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechTaggedSSML implements ITextToSpeechTaggedSSML.
var _ ITextToSpeechTaggedSSML = TextToSpeechTaggedSSML{}

// An interface definition for the [TextToSpeechTaggedSSML] class.
//
// # Methods
//
//   - [ITextToSpeechTaggedSSML.OriginalSSML]
//   - [ITextToSpeechTaggedSSML.SetOriginalSSML]
//   - [ITextToSpeechTaggedSSML.SsmlSnippets]
//   - [ITextToSpeechTaggedSSML.SetSsmlSnippets]
type ITextToSpeechTaggedSSML interface {
	objectivec.IObject

	// Topic: Methods

	OriginalSSML() string
	SetOriginalSSML(value string)
	SsmlSnippets() foundation.INSArray
	SetSsmlSnippets(value foundation.INSArray)
}

// Init initializes the instance.
func (t TextToSpeechTaggedSSML) Init() TextToSpeechTaggedSSML {
	rv := objc.SendIfResponds[TextToSpeechTaggedSSML](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTaggedSSML) Autorelease() TextToSpeechTaggedSSML {
	rv := objc.SendIfResponds[TextToSpeechTaggedSSML](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTaggedSSML creates a new TextToSpeechTaggedSSML instance.
func NewTextToSpeechTaggedSSML() TextToSpeechTaggedSSML {
	class := getTextToSpeechTaggedSSMLClass()
	rv := objc.SendIfResponds[TextToSpeechTaggedSSML](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TextToSpeechTaggedSSML) OriginalSSML() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("originalSSML"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechTaggedSSML) SetOriginalSSML(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setOriginalSSML:"), objc.String(value))
}
func (t TextToSpeechTaggedSSML) SsmlSnippets() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("ssmlSnippets"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechTaggedSSML) SetSsmlSnippets(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSsmlSnippets:"), value)
}
