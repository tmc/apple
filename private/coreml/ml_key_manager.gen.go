// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLKeyManager] class.
var (
	_MLKeyManagerClass     MLKeyManagerClass
	_MLKeyManagerClassOnce sync.Once
)

func getMLKeyManagerClass() MLKeyManagerClass {
	_MLKeyManagerClassOnce.Do(func() {
		_MLKeyManagerClass = MLKeyManagerClass{class: objc.GetClass("MLKeyManager")}
	})
	return _MLKeyManagerClass
}

// GetMLKeyManagerClass returns the class object for MLKeyManager.
func GetMLKeyManagerClass() MLKeyManagerClass {
	return getMLKeyManagerClass()
}

type MLKeyManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLKeyManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLKeyManagerClass) Alloc() MLKeyManager {
	rv := objc.Send[MLKeyManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager
type MLKeyManager struct {
	objectivec.Object
}

// MLKeyManagerFromID constructs a [MLKeyManager] from an objc.ID.
func MLKeyManagerFromID(id objc.ID) MLKeyManager {
	return MLKeyManager{objectivec.Object{ID: id}}
}

// Ensure MLKeyManager implements IMLKeyManager.
var _ IMLKeyManager = MLKeyManager{}

// An interface definition for the [MLKeyManager] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLKeyManager
type IMLKeyManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (k MLKeyManager) Init() MLKeyManager {
	rv := objc.Send[MLKeyManager](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MLKeyManager) Autorelease() MLKeyManager {
	rv := objc.Send[MLKeyManager](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKeyManager creates a new MLKeyManager instance.
func NewMLKeyManager() MLKeyManager {
	class := getMLKeyManagerClass()
	rv := objc.Send[MLKeyManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/createPersistentKeyBlobForKeyID:usesCodeSigningIdentityForEncryption:error:
func (_MLKeyManagerClass MLKeyManagerClass) CreatePersistentKeyBlobForKeyIDUsesCodeSigningIdentityForEncryptionError(id objectivec.IObject, encryption bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKeyManagerClass.class), objc.Sel("createPersistentKeyBlobForKeyID:usesCodeSigningIdentityForEncryption:error:"), id, encryption, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/decryptSessionForModelAtURL:error:
func (_MLKeyManagerClass MLKeyManagerClass) DecryptSessionForModelAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKeyManagerClass.class), objc.Sel("decryptSessionForModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/extractKeyIdentifierFromModelAtURL:usesCodeSigningIdentityForEncryption:error:
func (_MLKeyManagerClass MLKeyManagerClass) ExtractKeyIdentifierFromModelAtURLUsesCodeSigningIdentityForEncryptionError(url foundation.INSURL, encryption unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKeyManagerClass.class), objc.Sel("extractKeyIdentifierFromModelAtURL:usesCodeSigningIdentityForEncryption:error:"), url, encryption, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/isModelEncrypted:
func (_MLKeyManagerClass MLKeyManagerClass) IsModelEncrypted(encrypted objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLKeyManagerClass.class), objc.Sel("isModelEncrypted:"), encrypted)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/loadDecryptionKeyForModelAtURL:retUsesCodeSigningIdentityForEncryption:error:
func (_MLKeyManagerClass MLKeyManagerClass) LoadDecryptionKeyForModelAtURLRetUsesCodeSigningIdentityForEncryptionError(url foundation.INSURL, encryption unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKeyManagerClass.class), objc.Sel("loadDecryptionKeyForModelAtURL:retUsesCodeSigningIdentityForEncryption:error:"), url, encryption, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKeyManager/syncQueue
func (_MLKeyManagerClass MLKeyManagerClass) SyncQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLKeyManagerClass.class), objc.Sel("syncQueue"))
	return objectivec.Object{ID: rv}
}
