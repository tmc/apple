// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechNetworkTransportListener] class.
var (
	_TextToSpeechNetworkTransportListenerClass     TextToSpeechNetworkTransportListenerClass
	_TextToSpeechNetworkTransportListenerClassOnce sync.Once
)

func getTextToSpeechNetworkTransportListenerClass() TextToSpeechNetworkTransportListenerClass {
	_TextToSpeechNetworkTransportListenerClassOnce.Do(func() {
		_TextToSpeechNetworkTransportListenerClass = TextToSpeechNetworkTransportListenerClass{class: objc.GetClass("TextToSpeech.NetworkTransportListener")}
	})
	return _TextToSpeechNetworkTransportListenerClass
}

// GetTextToSpeechNetworkTransportListenerClass returns the class object for TextToSpeech.NetworkTransportListener.
func GetTextToSpeechNetworkTransportListenerClass() TextToSpeechNetworkTransportListenerClass {
	return getTextToSpeechNetworkTransportListenerClass()
}

type TextToSpeechNetworkTransportListenerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechNetworkTransportListenerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechNetworkTransportListenerClass) Alloc() TextToSpeechNetworkTransportListener {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportListener](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechNetworkTransportListener struct {
	objectivec.Object
}

// TextToSpeechNetworkTransportListenerFromID constructs a [TextToSpeechNetworkTransportListener] from an objc.ID.
func TextToSpeechNetworkTransportListenerFromID(id objc.ID) TextToSpeechNetworkTransportListener {
	return TextToSpeechNetworkTransportListener{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechNetworkTransportListener implements ITextToSpeechNetworkTransportListener.
var _ ITextToSpeechNetworkTransportListener = TextToSpeechNetworkTransportListener{}

// An interface definition for the [TextToSpeechNetworkTransportListener] class.
type ITextToSpeechNetworkTransportListener interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechNetworkTransportListener) Init() TextToSpeechNetworkTransportListener {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportListener](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechNetworkTransportListener) Autorelease() TextToSpeechNetworkTransportListener {
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportListener](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechNetworkTransportListener creates a new TextToSpeechNetworkTransportListener instance.
func NewTextToSpeechNetworkTransportListener() TextToSpeechNetworkTransportListener {
	class := getTextToSpeechNetworkTransportListenerClass()
	rv := objc.SendIfResponds[TextToSpeechNetworkTransportListener](objc.ID(class.class), objc.Sel("new"))
	return rv
}
