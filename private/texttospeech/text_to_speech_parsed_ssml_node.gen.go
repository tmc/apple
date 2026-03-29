// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechParsedSSMLNode] class.
var (
	_TextToSpeechParsedSSMLNodeClass     TextToSpeechParsedSSMLNodeClass
	_TextToSpeechParsedSSMLNodeClassOnce sync.Once
)

func getTextToSpeechParsedSSMLNodeClass() TextToSpeechParsedSSMLNodeClass {
	_TextToSpeechParsedSSMLNodeClassOnce.Do(func() {
		_TextToSpeechParsedSSMLNodeClass = TextToSpeechParsedSSMLNodeClass{class: objc.GetClass("TextToSpeech.ParsedSSMLNode")}
	})
	return _TextToSpeechParsedSSMLNodeClass
}

// GetTextToSpeechParsedSSMLNodeClass returns the class object for TextToSpeech.ParsedSSMLNode.
func GetTextToSpeechParsedSSMLNodeClass() TextToSpeechParsedSSMLNodeClass {
	return getTextToSpeechParsedSSMLNodeClass()
}

type TextToSpeechParsedSSMLNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechParsedSSMLNodeClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechParsedSSMLNodeClass) Alloc() TextToSpeechParsedSSMLNode {
	rv := objc.Send[TextToSpeechParsedSSMLNode](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechParsedSSMLNode.Description]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.ParsedSSMLNode
type TextToSpeechParsedSSMLNode struct {
	objectivec.Object
}

// TextToSpeechParsedSSMLNodeFromID constructs a [TextToSpeechParsedSSMLNode] from an objc.ID.
func TextToSpeechParsedSSMLNodeFromID(id objc.ID) TextToSpeechParsedSSMLNode {
	return TextToSpeechParsedSSMLNode{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechParsedSSMLNode implements ITextToSpeechParsedSSMLNode.
var _ ITextToSpeechParsedSSMLNode = TextToSpeechParsedSSMLNode{}

// An interface definition for the [TextToSpeechParsedSSMLNode] class.
//
// # Methods
//
//   - [ITextToSpeechParsedSSMLNode.Description]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.ParsedSSMLNode
type ITextToSpeechParsedSSMLNode interface {
	objectivec.IObject

	// Topic: Methods

	Description() string
}

// Init initializes the instance.
func (t TextToSpeechParsedSSMLNode) Init() TextToSpeechParsedSSMLNode {
	rv := objc.Send[TextToSpeechParsedSSMLNode](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechParsedSSMLNode) Autorelease() TextToSpeechParsedSSMLNode {
	rv := objc.Send[TextToSpeechParsedSSMLNode](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechParsedSSMLNode creates a new TextToSpeechParsedSSMLNode instance.
func NewTextToSpeechParsedSSMLNode() TextToSpeechParsedSSMLNode {
	class := getTextToSpeechParsedSSMLNodeClass()
	rv := objc.Send[TextToSpeechParsedSSMLNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.ParsedSSMLNode/description
func (t TextToSpeechParsedSSMLNode) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

