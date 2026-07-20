// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Store] class.
var (
	_StoreClass     StoreClass
	_StoreClassOnce sync.Once
)

func getStoreClass() StoreClass {
	_StoreClassOnce.Do(func() {
		_StoreClass = StoreClass{class: objc.GetClass("_TtCC12TextToSpeech13TTSTaskRunner5Store")}
	})
	return _StoreClass
}

// GetStoreClass returns the class object for _TtCC12TextToSpeech13TTSTaskRunner5Store.
func GetStoreClass() StoreClass {
	return getStoreClass()
}

type StoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc StoreClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc StoreClass) Alloc() Store {
	rv := objc.Send[Store](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type Store struct {
	objectivec.Object
}

// StoreFromID constructs a [Store] from an objc.ID.
func StoreFromID(id objc.ID) Store {
	return Store{objectivec.Object{ID: id}}
}

// NOTE: Store struct embeds objectivec.Object (parent type unavailable) but
// IStore embeds the parent interface; skip compile-time assertion.

// An interface definition for the [Store] class.
type IStore interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s Store) Init() Store {
	rv := objc.Send[Store](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Store) Autorelease() Store {
	rv := objc.Send[Store](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewStore creates a new Store instance.
func NewStore() Store {
	class := getStoreClass()
	rv := objc.Send[Store](objc.ID(class.class), objc.Sel("new"))
	return rv
}
