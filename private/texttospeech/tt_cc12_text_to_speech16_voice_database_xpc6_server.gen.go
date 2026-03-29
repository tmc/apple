// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC12TextToSpeech16VoiceDatabaseXPC6Server] class.
var (
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass     TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClassOnce sync.Once
)

func getTtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass() TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass {
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClassOnce.Do(func() {
		_TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass = TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass{class: objc.GetClass("_TtCC12TextToSpeech16VoiceDatabaseXPC6Server")}
	})
	return _TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass
}

// GetTtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass returns the class object for _TtCC12TextToSpeech16VoiceDatabaseXPC6Server.
func GetTtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass() TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass {
	return getTtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass()
}

type TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass) Alloc() TtCC12TextToSpeech16VoiceDatabaseXPC6Server {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Server](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech16VoiceDatabaseXPC6Server
type TtCC12TextToSpeech16VoiceDatabaseXPC6Server struct {
	objectivec.Object
}

// TtCC12TextToSpeech16VoiceDatabaseXPC6ServerFromID constructs a [TtCC12TextToSpeech16VoiceDatabaseXPC6Server] from an objc.ID.
func TtCC12TextToSpeech16VoiceDatabaseXPC6ServerFromID(id objc.ID) TtCC12TextToSpeech16VoiceDatabaseXPC6Server {
	return TtCC12TextToSpeech16VoiceDatabaseXPC6Server{objectivec.Object{ID: id}}
}
// NOTE: TtCC12TextToSpeech16VoiceDatabaseXPC6Server struct embeds objectivec.Object (parent type unavailable) but
// ITtCC12TextToSpeech16VoiceDatabaseXPC6Server embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC12TextToSpeech16VoiceDatabaseXPC6Server] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech16VoiceDatabaseXPC6Server
type ITtCC12TextToSpeech16VoiceDatabaseXPC6Server interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC12TextToSpeech16VoiceDatabaseXPC6Server) Init() TtCC12TextToSpeech16VoiceDatabaseXPC6Server {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Server](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC12TextToSpeech16VoiceDatabaseXPC6Server) Autorelease() TtCC12TextToSpeech16VoiceDatabaseXPC6Server {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Server](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC12TextToSpeech16VoiceDatabaseXPC6Server creates a new TtCC12TextToSpeech16VoiceDatabaseXPC6Server instance.
func NewTtCC12TextToSpeech16VoiceDatabaseXPC6Server() TtCC12TextToSpeech16VoiceDatabaseXPC6Server {
	class := getTtCC12TextToSpeech16VoiceDatabaseXPC6ServerClass()
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Server](objc.ID(class.class), objc.Sel("new"))
	return rv
}

