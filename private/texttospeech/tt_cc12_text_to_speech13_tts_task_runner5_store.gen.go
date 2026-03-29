// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC12TextToSpeech13TTSTaskRunner5Store] class.
var (
	_TtCC12TextToSpeech13TTSTaskRunner5StoreClass     TtCC12TextToSpeech13TTSTaskRunner5StoreClass
	_TtCC12TextToSpeech13TTSTaskRunner5StoreClassOnce sync.Once
)

func getTtCC12TextToSpeech13TTSTaskRunner5StoreClass() TtCC12TextToSpeech13TTSTaskRunner5StoreClass {
	_TtCC12TextToSpeech13TTSTaskRunner5StoreClassOnce.Do(func() {
		_TtCC12TextToSpeech13TTSTaskRunner5StoreClass = TtCC12TextToSpeech13TTSTaskRunner5StoreClass{class: objc.GetClass("_TtCC12TextToSpeech13TTSTaskRunner5Store")}
	})
	return _TtCC12TextToSpeech13TTSTaskRunner5StoreClass
}

// GetTtCC12TextToSpeech13TTSTaskRunner5StoreClass returns the class object for _TtCC12TextToSpeech13TTSTaskRunner5Store.
func GetTtCC12TextToSpeech13TTSTaskRunner5StoreClass() TtCC12TextToSpeech13TTSTaskRunner5StoreClass {
	return getTtCC12TextToSpeech13TTSTaskRunner5StoreClass()
}

type TtCC12TextToSpeech13TTSTaskRunner5StoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC12TextToSpeech13TTSTaskRunner5StoreClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC12TextToSpeech13TTSTaskRunner5StoreClass) Alloc() TtCC12TextToSpeech13TTSTaskRunner5Store {
	rv := objc.Send[TtCC12TextToSpeech13TTSTaskRunner5Store](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech13TTSTaskRunner5Store
type TtCC12TextToSpeech13TTSTaskRunner5Store struct {
	objectivec.Object
}

// TtCC12TextToSpeech13TTSTaskRunner5StoreFromID constructs a [TtCC12TextToSpeech13TTSTaskRunner5Store] from an objc.ID.
func TtCC12TextToSpeech13TTSTaskRunner5StoreFromID(id objc.ID) TtCC12TextToSpeech13TTSTaskRunner5Store {
	return TtCC12TextToSpeech13TTSTaskRunner5Store{objectivec.Object{ID: id}}
}
// NOTE: TtCC12TextToSpeech13TTSTaskRunner5Store struct embeds objectivec.Object (parent type unavailable) but
// ITtCC12TextToSpeech13TTSTaskRunner5Store embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC12TextToSpeech13TTSTaskRunner5Store] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech13TTSTaskRunner5Store
type ITtCC12TextToSpeech13TTSTaskRunner5Store interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC12TextToSpeech13TTSTaskRunner5Store) Init() TtCC12TextToSpeech13TTSTaskRunner5Store {
	rv := objc.Send[TtCC12TextToSpeech13TTSTaskRunner5Store](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC12TextToSpeech13TTSTaskRunner5Store) Autorelease() TtCC12TextToSpeech13TTSTaskRunner5Store {
	rv := objc.Send[TtCC12TextToSpeech13TTSTaskRunner5Store](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC12TextToSpeech13TTSTaskRunner5Store creates a new TtCC12TextToSpeech13TTSTaskRunner5Store instance.
func NewTtCC12TextToSpeech13TTSTaskRunner5Store() TtCC12TextToSpeech13TTSTaskRunner5Store {
	class := getTtCC12TextToSpeech13TTSTaskRunner5StoreClass()
	rv := objc.Send[TtCC12TextToSpeech13TTSTaskRunner5Store](objc.ID(class.class), objc.Sel("new"))
	return rv
}

