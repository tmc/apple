// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFairPlayDecryptSessionManager] class.
var (
	_MLFairPlayDecryptSessionManagerClass     MLFairPlayDecryptSessionManagerClass
	_MLFairPlayDecryptSessionManagerClassOnce sync.Once
)

func getMLFairPlayDecryptSessionManagerClass() MLFairPlayDecryptSessionManagerClass {
	_MLFairPlayDecryptSessionManagerClassOnce.Do(func() {
		_MLFairPlayDecryptSessionManagerClass = MLFairPlayDecryptSessionManagerClass{class: objc.GetClass("MLFairPlayDecryptSessionManager")}
	})
	return _MLFairPlayDecryptSessionManagerClass
}

// GetMLFairPlayDecryptSessionManagerClass returns the class object for MLFairPlayDecryptSessionManager.
func GetMLFairPlayDecryptSessionManagerClass() MLFairPlayDecryptSessionManagerClass {
	return getMLFairPlayDecryptSessionManagerClass()
}

type MLFairPlayDecryptSessionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFairPlayDecryptSessionManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFairPlayDecryptSessionManagerClass) Alloc() MLFairPlayDecryptSessionManager {
	rv := objc.Send[MLFairPlayDecryptSessionManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLFairPlayDecryptSessionManager.ModelPathToSessionID]
//   - [MLFairPlayDecryptSessionManager.SessionContext]
//   - [MLFairPlayDecryptSessionManager.StartDecryptionOfModelAtPathUsingKeyBlobTeamIdentifierError]
//   - [MLFairPlayDecryptSessionManager.StopDecryptionOfModelAtPath]
//   - [MLFairPlayDecryptSessionManager.StopDecryptionOfModelAtPathError]
//   - [MLFairPlayDecryptSessionManager.SyncQueue]
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager
type MLFairPlayDecryptSessionManager struct {
	objectivec.Object
}

// MLFairPlayDecryptSessionManagerFromID constructs a [MLFairPlayDecryptSessionManager] from an objc.ID.
func MLFairPlayDecryptSessionManagerFromID(id objc.ID) MLFairPlayDecryptSessionManager {
	return MLFairPlayDecryptSessionManager{objectivec.Object{ID: id}}
}
// Ensure MLFairPlayDecryptSessionManager implements IMLFairPlayDecryptSessionManager.
var _ IMLFairPlayDecryptSessionManager = MLFairPlayDecryptSessionManager{}

// An interface definition for the [MLFairPlayDecryptSessionManager] class.
//
// # Methods
//
//   - [IMLFairPlayDecryptSessionManager.ModelPathToSessionID]
//   - [IMLFairPlayDecryptSessionManager.SessionContext]
//   - [IMLFairPlayDecryptSessionManager.StartDecryptionOfModelAtPathUsingKeyBlobTeamIdentifierError]
//   - [IMLFairPlayDecryptSessionManager.StopDecryptionOfModelAtPath]
//   - [IMLFairPlayDecryptSessionManager.StopDecryptionOfModelAtPathError]
//   - [IMLFairPlayDecryptSessionManager.SyncQueue]
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager
type IMLFairPlayDecryptSessionManager interface {
	objectivec.IObject

	// Topic: Methods

	ModelPathToSessionID() foundation.INSDictionary
	SessionContext() objectivec.IObject
	StartDecryptionOfModelAtPathUsingKeyBlobTeamIdentifierError(path objectivec.IObject, blob objectivec.IObject, identifier objectivec.IObject) (bool, error)
	StopDecryptionOfModelAtPath(path objectivec.IObject) int
	StopDecryptionOfModelAtPathError(path objectivec.IObject) (bool, error)
	SyncQueue() objectivec.Object
}

// Init initializes the instance.
func (f MLFairPlayDecryptSessionManager) Init() MLFairPlayDecryptSessionManager {
	rv := objc.Send[MLFairPlayDecryptSessionManager](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFairPlayDecryptSessionManager) Autorelease() MLFairPlayDecryptSessionManager {
	rv := objc.Send[MLFairPlayDecryptSessionManager](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFairPlayDecryptSessionManager creates a new MLFairPlayDecryptSessionManager instance.
func NewMLFairPlayDecryptSessionManager() MLFairPlayDecryptSessionManager {
	class := getMLFairPlayDecryptSessionManagerClass()
	rv := objc.Send[MLFairPlayDecryptSessionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/startDecryptionOfModelAtPath:usingKeyBlob:teamIdentifier:error:
func (f MLFairPlayDecryptSessionManager) StartDecryptionOfModelAtPathUsingKeyBlobTeamIdentifierError(path objectivec.IObject, blob objectivec.IObject, identifier objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("startDecryptionOfModelAtPath:usingKeyBlob:teamIdentifier:error:"), path, blob, identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startDecryptionOfModelAtPath:usingKeyBlob:teamIdentifier:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/stopDecryptionOfModelAtPath:
func (f MLFairPlayDecryptSessionManager) StopDecryptionOfModelAtPath(path objectivec.IObject) int {
	rv := objc.Send[int](f.ID, objc.Sel("stopDecryptionOfModelAtPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/stopDecryptionOfModelAtPath:error:
func (f MLFairPlayDecryptSessionManager) StopDecryptionOfModelAtPathError(path objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("stopDecryptionOfModelAtPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopDecryptionOfModelAtPath:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/modelPathToSessionID
func (f MLFairPlayDecryptSessionManager) ModelPathToSessionID() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("modelPathToSessionID"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/sessionContext
func (f MLFairPlayDecryptSessionManager) SessionContext() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sessionContext"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSessionManager/syncQueue
func (f MLFairPlayDecryptSessionManager) SyncQueue() objectivec.Object {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("syncQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

