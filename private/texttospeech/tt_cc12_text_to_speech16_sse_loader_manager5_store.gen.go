// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SSELoaderManagerStore] class.
var (
	_SSELoaderManagerStoreClass     SSELoaderManagerStoreClass
	_SSELoaderManagerStoreClassOnce sync.Once
)

func getSSELoaderManagerStoreClass() SSELoaderManagerStoreClass {
	_SSELoaderManagerStoreClassOnce.Do(func() {
		_SSELoaderManagerStoreClass = SSELoaderManagerStoreClass{class: objc.GetClass("_TtCC12TextToSpeech16SSELoaderManager5Store")}
	})
	return _SSELoaderManagerStoreClass
}

// GetSSELoaderManagerStoreClass returns the class object for _TtCC12TextToSpeech16SSELoaderManager5Store.
func GetSSELoaderManagerStoreClass() SSELoaderManagerStoreClass {
	return getSSELoaderManagerStoreClass()
}

type SSELoaderManagerStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SSELoaderManagerStoreClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SSELoaderManagerStoreClass) Alloc() SSELoaderManagerStore {
	rv := objc.SendIfResponds[SSELoaderManagerStore](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SSELoaderManagerStore struct {
	objectivec.Object
}

// SSELoaderManagerStoreFromID constructs a [SSELoaderManagerStore] from an objc.ID.
func SSELoaderManagerStoreFromID(id objc.ID) SSELoaderManagerStore {
	return SSELoaderManagerStore{objectivec.Object{ID: id}}
}

// Ensure SSELoaderManagerStore implements ISSELoaderManagerStore.
var _ ISSELoaderManagerStore = SSELoaderManagerStore{}

// An interface definition for the [SSELoaderManagerStore] class.
type ISSELoaderManagerStore interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SSELoaderManagerStore) Init() SSELoaderManagerStore {
	rv := objc.SendIfResponds[SSELoaderManagerStore](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SSELoaderManagerStore) Autorelease() SSELoaderManagerStore {
	rv := objc.SendIfResponds[SSELoaderManagerStore](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSSELoaderManagerStore creates a new SSELoaderManagerStore instance.
func NewSSELoaderManagerStore() SSELoaderManagerStore {
	class := getSSELoaderManagerStoreClass()
	rv := objc.SendIfResponds[SSELoaderManagerStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}
