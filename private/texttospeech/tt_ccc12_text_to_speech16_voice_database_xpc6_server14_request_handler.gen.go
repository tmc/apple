// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler] class.
var (
	_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass     TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass
	_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClassOnce sync.Once
)

func getTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass {
	_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClassOnce.Do(func() {
		_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass = TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass{class: objc.GetClass("_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler")}
	})
	return _TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass
}

// GetTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass returns the class object for _TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler.
func GetTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass {
	return getTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass()
}

type TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass) Alloc() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler {
	rv := objc.Send[TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler
type TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler struct {
	objectivec.Object
}

// TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerFromID constructs a [TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler] from an objc.ID.
func TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerFromID(id objc.ID) TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler {
	return TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler{objectivec.Object{ID: id}}
}

// NOTE: TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler struct embeds objectivec.Object (parent type unavailable) but
// ITtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler
type ITtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler) Init() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler {
	rv := objc.Send[TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler) Autorelease() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler {
	rv := objc.Send[TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler creates a new TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler instance.
func NewTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler() TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler {
	class := getTtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandlerClass()
	rv := objc.Send[TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}
