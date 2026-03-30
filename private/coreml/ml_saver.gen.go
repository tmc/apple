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

// The class instance for the [MLSaver] class.
var (
	_MLSaverClass     MLSaverClass
	_MLSaverClassOnce sync.Once
)

func getMLSaverClass() MLSaverClass {
	_MLSaverClassOnce.Do(func() {
		_MLSaverClass = MLSaverClass{class: objc.GetClass("MLSaver")}
	})
	return _MLSaverClass
}

// GetMLSaverClass returns the class object for MLSaver.
func GetMLSaverClass() MLSaverClass {
	return getMLSaverClass()
}

type MLSaverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSaverClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSaverClass) Alloc() MLSaver {
	rv := objc.Send[MLSaver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSaver
type MLSaver struct {
	objectivec.Object
}

// MLSaverFromID constructs a [MLSaver] from an objc.ID.
func MLSaverFromID(id objc.ID) MLSaver {
	return MLSaver{objectivec.Object{ID: id}}
}

// Ensure MLSaver implements IMLSaver.
var _ IMLSaver = MLSaver{}

// An interface definition for the [MLSaver] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLSaver
type IMLSaver interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s MLSaver) Init() MLSaver {
	rv := objc.Send[MLSaver](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSaver) Autorelease() MLSaver {
	rv := objc.Send[MLSaver](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSaver creates a new MLSaver instance.
func NewMLSaver() MLSaver {
	class := getMLSaverClass()
	rv := objc.Send[MLSaver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSaver/copyModelAtURL:toURL:error:
func (_MLSaverClass MLSaverClass) CopyModelAtURLToURLError(url foundation.INSURL, url2 foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLSaverClass.class), objc.Sel("copyModelAtURL:toURL:error:"), url, url2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyModelAtURL:toURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSaver/saveModelToArchive:model:compilerOptions:error:
func (_MLSaverClass MLSaverClass) SaveModelToArchiveModelCompilerOptionsError(archive unsafe.Pointer, model objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLSaverClass.class), objc.Sel("saveModelToArchive:model:compilerOptions:error:"), archive, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveModelToArchive:model:compilerOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSaver/saveModelToArchive:model:error:
func (_MLSaverClass MLSaverClass) SaveModelToArchiveModelError(archive unsafe.Pointer, model objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLSaverClass.class), objc.Sel("saveModelToArchive:model:error:"), archive, model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveModelToArchive:model:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSaver/saveModelToAssetAtURL:model:error:
func (_MLSaverClass MLSaverClass) SaveModelToAssetAtURLModelError(url foundation.INSURL, model objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLSaverClass.class), objc.Sel("saveModelToAssetAtURL:model:error:"), url, model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveModelToAssetAtURL:model:error: returned NO with nil NSError")
	}
	return rv, nil

}
