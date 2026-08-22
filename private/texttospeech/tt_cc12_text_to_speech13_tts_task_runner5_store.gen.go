// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSTaskRunnerStore] class.
var (
	_TTSTaskRunnerStoreClass     TTSTaskRunnerStoreClass
	_TTSTaskRunnerStoreClassOnce sync.Once
)

func getTTSTaskRunnerStoreClass() TTSTaskRunnerStoreClass {
	_TTSTaskRunnerStoreClassOnce.Do(func() {
		_TTSTaskRunnerStoreClass = TTSTaskRunnerStoreClass{class: objc.GetClass("_TtCC12TextToSpeech13TTSTaskRunner5Store")}
	})
	return _TTSTaskRunnerStoreClass
}

// GetTTSTaskRunnerStoreClass returns the class object for _TtCC12TextToSpeech13TTSTaskRunner5Store.
func GetTTSTaskRunnerStoreClass() TTSTaskRunnerStoreClass {
	return getTTSTaskRunnerStoreClass()
}

type TTSTaskRunnerStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSTaskRunnerStoreClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSTaskRunnerStoreClass) Alloc() TTSTaskRunnerStore {
	rv := objc.SendIfResponds[TTSTaskRunnerStore](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TTSTaskRunnerStore struct {
	objectivec.Object
}

// TTSTaskRunnerStoreFromID constructs a [TTSTaskRunnerStore] from an objc.ID.
func TTSTaskRunnerStoreFromID(id objc.ID) TTSTaskRunnerStore {
	return TTSTaskRunnerStore{objectivec.Object{ID: id}}
}

// Ensure TTSTaskRunnerStore implements ITTSTaskRunnerStore.
var _ ITTSTaskRunnerStore = TTSTaskRunnerStore{}

// An interface definition for the [TTSTaskRunnerStore] class.
type ITTSTaskRunnerStore interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSTaskRunnerStore) Init() TTSTaskRunnerStore {
	rv := objc.SendIfResponds[TTSTaskRunnerStore](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSTaskRunnerStore) Autorelease() TTSTaskRunnerStore {
	rv := objc.SendIfResponds[TTSTaskRunnerStore](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSTaskRunnerStore creates a new TTSTaskRunnerStore instance.
func NewTTSTaskRunnerStore() TTSTaskRunnerStore {
	class := getTTSTaskRunnerStoreClass()
	rv := objc.SendIfResponds[TTSTaskRunnerStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}
