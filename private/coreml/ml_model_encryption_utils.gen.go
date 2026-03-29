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

// The class instance for the [MLModelEncryptionUtils] class.
var (
	_MLModelEncryptionUtilsClass     MLModelEncryptionUtilsClass
	_MLModelEncryptionUtilsClassOnce sync.Once
)

func getMLModelEncryptionUtilsClass() MLModelEncryptionUtilsClass {
	_MLModelEncryptionUtilsClassOnce.Do(func() {
		_MLModelEncryptionUtilsClass = MLModelEncryptionUtilsClass{class: objc.GetClass("MLModelEncryptionUtils")}
	})
	return _MLModelEncryptionUtilsClass
}

// GetMLModelEncryptionUtilsClass returns the class object for MLModelEncryptionUtils.
func GetMLModelEncryptionUtilsClass() MLModelEncryptionUtilsClass {
	return getMLModelEncryptionUtilsClass()
}

type MLModelEncryptionUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelEncryptionUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelEncryptionUtilsClass) Alloc() MLModelEncryptionUtils {
	rv := objc.Send[MLModelEncryptionUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEncryptionUtils
type MLModelEncryptionUtils struct {
	objectivec.Object
}

// MLModelEncryptionUtilsFromID constructs a [MLModelEncryptionUtils] from an objc.ID.
func MLModelEncryptionUtilsFromID(id objc.ID) MLModelEncryptionUtils {
	return MLModelEncryptionUtils{objectivec.Object{ID: id}}
}
// Ensure MLModelEncryptionUtils implements IMLModelEncryptionUtils.
var _ IMLModelEncryptionUtils = MLModelEncryptionUtils{}

// An interface definition for the [MLModelEncryptionUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelEncryptionUtils
type IMLModelEncryptionUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelEncryptionUtils) Init() MLModelEncryptionUtils {
	rv := objc.Send[MLModelEncryptionUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelEncryptionUtils) Autorelease() MLModelEncryptionUtils {
	rv := objc.Send[MLModelEncryptionUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelEncryptionUtils creates a new MLModelEncryptionUtils instance.
func NewMLModelEncryptionUtils() MLModelEncryptionUtils {
	class := getMLModelEncryptionUtilsClass()
	rv := objc.Send[MLModelEncryptionUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEncryptionUtils/addEncryptionHeaderToUnencryptedFile:saveToFile:error:
func (_MLModelEncryptionUtilsClass MLModelEncryptionUtilsClass) AddEncryptionHeaderToUnencryptedFileSaveToFileError(file objectivec.IObject, file2 objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelEncryptionUtilsClass.class), objc.Sel("addEncryptionHeaderToUnencryptedFile:saveToFile:error:"), file, file2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addEncryptionHeaderToUnencryptedFile:saveToFile:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelEncryptionUtils/encryptFile:withKey:iv:saveToFile:error:
func (_MLModelEncryptionUtilsClass MLModelEncryptionUtilsClass) EncryptFileWithKeyIvSaveToFileError(file objectivec.IObject, key objectivec.IObject, iv objectivec.IObject, file2 objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelEncryptionUtilsClass.class), objc.Sel("encryptFile:withKey:iv:saveToFile:error:"), file, key, iv, file2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("encryptFile:withKey:iv:saveToFile:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModelEncryptionUtils/sinfData
func (_MLModelEncryptionUtilsClass MLModelEncryptionUtilsClass) SinfData() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelEncryptionUtilsClass.class), objc.Sel("sinfData"))
	return objectivec.Object{ID: rv}
}

