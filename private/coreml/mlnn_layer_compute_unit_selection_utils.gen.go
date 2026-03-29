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

// The class instance for the [MLNNLayerComputeUnitSelectionUtils] class.
var (
	_MLNNLayerComputeUnitSelectionUtilsClass     MLNNLayerComputeUnitSelectionUtilsClass
	_MLNNLayerComputeUnitSelectionUtilsClassOnce sync.Once
)

func getMLNNLayerComputeUnitSelectionUtilsClass() MLNNLayerComputeUnitSelectionUtilsClass {
	_MLNNLayerComputeUnitSelectionUtilsClassOnce.Do(func() {
		_MLNNLayerComputeUnitSelectionUtilsClass = MLNNLayerComputeUnitSelectionUtilsClass{class: objc.GetClass("MLNNLayerComputeUnitSelectionUtils")}
	})
	return _MLNNLayerComputeUnitSelectionUtilsClass
}

// GetMLNNLayerComputeUnitSelectionUtilsClass returns the class object for MLNNLayerComputeUnitSelectionUtils.
func GetMLNNLayerComputeUnitSelectionUtilsClass() MLNNLayerComputeUnitSelectionUtilsClass {
	return getMLNNLayerComputeUnitSelectionUtilsClass()
}

type MLNNLayerComputeUnitSelectionUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNNLayerComputeUnitSelectionUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNNLayerComputeUnitSelectionUtilsClass) Alloc() MLNNLayerComputeUnitSelectionUtils {
	rv := objc.Send[MLNNLayerComputeUnitSelectionUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils
type MLNNLayerComputeUnitSelectionUtils struct {
	objectivec.Object
}

// MLNNLayerComputeUnitSelectionUtilsFromID constructs a [MLNNLayerComputeUnitSelectionUtils] from an objc.ID.
func MLNNLayerComputeUnitSelectionUtilsFromID(id objc.ID) MLNNLayerComputeUnitSelectionUtils {
	return MLNNLayerComputeUnitSelectionUtils{objectivec.Object{ID: id}}
}
// Ensure MLNNLayerComputeUnitSelectionUtils implements IMLNNLayerComputeUnitSelectionUtils.
var _ IMLNNLayerComputeUnitSelectionUtils = MLNNLayerComputeUnitSelectionUtils{}

// An interface definition for the [MLNNLayerComputeUnitSelectionUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils
type IMLNNLayerComputeUnitSelectionUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n MLNNLayerComputeUnitSelectionUtils) Init() MLNNLayerComputeUnitSelectionUtils {
	rv := objc.Send[MLNNLayerComputeUnitSelectionUtils](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNNLayerComputeUnitSelectionUtils) Autorelease() MLNNLayerComputeUnitSelectionUtils {
	rv := objc.Send[MLNNLayerComputeUnitSelectionUtils](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNNLayerComputeUnitSelectionUtils creates a new MLNNLayerComputeUnitSelectionUtils instance.
func NewMLNNLayerComputeUnitSelectionUtils() MLNNLayerComputeUnitSelectionUtils {
	class := getMLNNLayerComputeUnitSelectionUtilsClass()
	rv := objc.Send[MLNNLayerComputeUnitSelectionUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils/getLayerHints:error:
func (_MLNNLayerComputeUnitSelectionUtilsClass MLNNLayerComputeUnitSelectionUtilsClass) GetLayerHintsError(hints objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNNLayerComputeUnitSelectionUtilsClass.class), objc.Sel("getLayerHints:error:"), hints, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils/getLayerTypes:error:
func (_MLNNLayerComputeUnitSelectionUtilsClass MLNNLayerComputeUnitSelectionUtilsClass) GetLayerTypesError(types objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNNLayerComputeUnitSelectionUtilsClass.class), objc.Sel("getLayerTypes:error:"), types, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils/getNetJson:error:
func (_MLNNLayerComputeUnitSelectionUtilsClass MLNNLayerComputeUnitSelectionUtilsClass) GetNetJsonError(json objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNNLayerComputeUnitSelectionUtilsClass.class), objc.Sel("getNetJson:error:"), json, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils/undoLastHintUpdate:error:
func (_MLNNLayerComputeUnitSelectionUtilsClass MLNNLayerComputeUnitSelectionUtilsClass) UndoLastHintUpdateError(update objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLNNLayerComputeUnitSelectionUtilsClass.class), objc.Sel("undoLastHintUpdate:error:"), update, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("undoLastHintUpdate:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNNLayerComputeUnitSelectionUtils/updateHints:hints:error:
func (_MLNNLayerComputeUnitSelectionUtilsClass MLNNLayerComputeUnitSelectionUtilsClass) UpdateHintsHintsError(hints objectivec.IObject, hints2 objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLNNLayerComputeUnitSelectionUtilsClass.class), objc.Sel("updateHints:hints:error:"), hints, hints2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateHints:hints:error: returned NO with nil NSError")
	}
	return rv, nil

}

