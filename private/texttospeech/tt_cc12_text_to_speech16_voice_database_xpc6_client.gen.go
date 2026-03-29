// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC12TextToSpeech16VoiceDatabaseXPC6Client] class.
var (
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass     TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClassOnce sync.Once
)

func getTtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass() TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass {
	_TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClassOnce.Do(func() {
		_TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass = TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass{class: objc.GetClass("_TtCC12TextToSpeech16VoiceDatabaseXPC6Client")}
	})
	return _TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass
}

// GetTtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass returns the class object for _TtCC12TextToSpeech16VoiceDatabaseXPC6Client.
func GetTtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass() TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass {
	return getTtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass()
}

type TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass) Alloc() TtCC12TextToSpeech16VoiceDatabaseXPC6Client {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Client](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech16VoiceDatabaseXPC6Client
type TtCC12TextToSpeech16VoiceDatabaseXPC6Client struct {
	objectivec.Object
}

// TtCC12TextToSpeech16VoiceDatabaseXPC6ClientFromID constructs a [TtCC12TextToSpeech16VoiceDatabaseXPC6Client] from an objc.ID.
func TtCC12TextToSpeech16VoiceDatabaseXPC6ClientFromID(id objc.ID) TtCC12TextToSpeech16VoiceDatabaseXPC6Client {
	return TtCC12TextToSpeech16VoiceDatabaseXPC6Client{objectivec.Object{ID: id}}
}
// NOTE: TtCC12TextToSpeech16VoiceDatabaseXPC6Client struct embeds objectivec.Object (parent type unavailable) but
// ITtCC12TextToSpeech16VoiceDatabaseXPC6Client embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC12TextToSpeech16VoiceDatabaseXPC6Client] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech16VoiceDatabaseXPC6Client
type ITtCC12TextToSpeech16VoiceDatabaseXPC6Client interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC12TextToSpeech16VoiceDatabaseXPC6Client) Init() TtCC12TextToSpeech16VoiceDatabaseXPC6Client {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Client](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC12TextToSpeech16VoiceDatabaseXPC6Client) Autorelease() TtCC12TextToSpeech16VoiceDatabaseXPC6Client {
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Client](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC12TextToSpeech16VoiceDatabaseXPC6Client creates a new TtCC12TextToSpeech16VoiceDatabaseXPC6Client instance.
func NewTtCC12TextToSpeech16VoiceDatabaseXPC6Client() TtCC12TextToSpeech16VoiceDatabaseXPC6Client {
	class := getTtCC12TextToSpeech16VoiceDatabaseXPC6ClientClass()
	rv := objc.Send[TtCC12TextToSpeech16VoiceDatabaseXPC6Client](objc.ID(class.class), objc.Sel("new"))
	return rv
}

