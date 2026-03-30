// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPersistentKeyStorage] class.
var (
	_MLPersistentKeyStorageClass     MLPersistentKeyStorageClass
	_MLPersistentKeyStorageClassOnce sync.Once
)

func getMLPersistentKeyStorageClass() MLPersistentKeyStorageClass {
	_MLPersistentKeyStorageClassOnce.Do(func() {
		_MLPersistentKeyStorageClass = MLPersistentKeyStorageClass{class: objc.GetClass("MLPersistentKeyStorage")}
	})
	return _MLPersistentKeyStorageClass
}

// GetMLPersistentKeyStorageClass returns the class object for MLPersistentKeyStorage.
func GetMLPersistentKeyStorageClass() MLPersistentKeyStorageClass {
	return getMLPersistentKeyStorageClass()
}

type MLPersistentKeyStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPersistentKeyStorageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPersistentKeyStorageClass) Alloc() MLPersistentKeyStorage {
	rv := objc.Send[MLPersistentKeyStorage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage
type MLPersistentKeyStorage struct {
	objectivec.Object
}

// MLPersistentKeyStorageFromID constructs a [MLPersistentKeyStorage] from an objc.ID.
func MLPersistentKeyStorageFromID(id objc.ID) MLPersistentKeyStorage {
	return MLPersistentKeyStorage{objectivec.Object{ID: id}}
}

// Ensure MLPersistentKeyStorage implements IMLPersistentKeyStorage.
var _ IMLPersistentKeyStorage = MLPersistentKeyStorage{}

// An interface definition for the [MLPersistentKeyStorage] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage
type IMLPersistentKeyStorage interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p MLPersistentKeyStorage) Init() MLPersistentKeyStorage {
	rv := objc.Send[MLPersistentKeyStorage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPersistentKeyStorage) Autorelease() MLPersistentKeyStorage {
	rv := objc.Send[MLPersistentKeyStorage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPersistentKeyStorage creates a new MLPersistentKeyStorage instance.
func NewMLPersistentKeyStorage() MLPersistentKeyStorage {
	class := getMLPersistentKeyStorageClass()
	rv := objc.Send[MLPersistentKeyStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage/persistentKeyStorageURL
func (_MLPersistentKeyStorageClass MLPersistentKeyStorageClass) PersistentKeyStorageURL() foundation.NSURL {
	rv := objc.Send[objc.ID](objc.ID(_MLPersistentKeyStorageClass.class), objc.Sel("persistentKeyStorageURL"))
	return foundation.NSURLFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage/retrieveKeyBlobForKeyIdentifier:
func (_MLPersistentKeyStorageClass MLPersistentKeyStorageClass) RetrieveKeyBlobForKeyIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLPersistentKeyStorageClass.class), objc.Sel("retrieveKeyBlobForKeyIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage/storeKeyBlob:forKeyIdentifier:error:
func (_MLPersistentKeyStorageClass MLPersistentKeyStorageClass) StoreKeyBlobForKeyIdentifierError(blob objectivec.IObject, identifier objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLPersistentKeyStorageClass.class), objc.Sel("storeKeyBlob:forKeyIdentifier:error:"), blob, identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("storeKeyBlob:forKeyIdentifier:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPersistentKeyStorage/syncQueue
func (_MLPersistentKeyStorageClass MLPersistentKeyStorageClass) SyncQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLPersistentKeyStorageClass.class), objc.Sel("syncQueue"))
	return objectivec.Object{ID: rv}
}
