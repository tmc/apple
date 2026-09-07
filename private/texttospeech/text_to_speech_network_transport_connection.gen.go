// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechNetworkTransportConnection] class.
var (
	_TextToSpeechNetworkTransportConnectionClass     TextToSpeechNetworkTransportConnectionClass
	_TextToSpeechNetworkTransportConnectionClassOnce sync.Once
)

func getTextToSpeechNetworkTransportConnectionClass() TextToSpeechNetworkTransportConnectionClass {
	_TextToSpeechNetworkTransportConnectionClassOnce.Do(func() {
		_TextToSpeechNetworkTransportConnectionClass = TextToSpeechNetworkTransportConnectionClass{class: objc.GetClass("TextToSpeech.NetworkTransportConnection")}
	})
	return _TextToSpeechNetworkTransportConnectionClass
}

// GetTextToSpeechNetworkTransportConnectionClass returns the class object for TextToSpeech.NetworkTransportConnection.
func GetTextToSpeechNetworkTransportConnectionClass() TextToSpeechNetworkTransportConnectionClass {
	return getTextToSpeechNetworkTransportConnectionClass()
}

type TextToSpeechNetworkTransportConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechNetworkTransportConnectionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechNetworkTransportConnectionClass) Alloc() TextToSpeechNetworkTransportConnection {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportConnection](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechNetworkTransportConnection struct {
	objectivec.Object
}

// TextToSpeechNetworkTransportConnectionFromID constructs a [TextToSpeechNetworkTransportConnection] from an objc.ID.
func TextToSpeechNetworkTransportConnectionFromID(id objc.ID) TextToSpeechNetworkTransportConnection {
	return TextToSpeechNetworkTransportConnection{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechNetworkTransportConnection implements ITextToSpeechNetworkTransportConnection.
var _ ITextToSpeechNetworkTransportConnection = TextToSpeechNetworkTransportConnection{}

// An interface definition for the [TextToSpeechNetworkTransportConnection] class.
type ITextToSpeechNetworkTransportConnection interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechNetworkTransportConnection) Init() TextToSpeechNetworkTransportConnection {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportConnection](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechNetworkTransportConnection) Autorelease() TextToSpeechNetworkTransportConnection {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportConnection](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechNetworkTransportConnection creates a new TextToSpeechNetworkTransportConnection instance.
func NewTextToSpeechNetworkTransportConnection() TextToSpeechNetworkTransportConnection {
	class := getTextToSpeechNetworkTransportConnectionClass()
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}
